import Fluent
import Vapor

struct ProductsController: RouteCollection {
	func boot(routes: any RoutesBuilder) throws {
		let products = routes.grouped("products")

		products.get(use: self.list)
		products.post(use: self.create)
		products.group(":pid") { product in
			product.get(use: self.read)
			product.put(use: self.update)
			product.delete(use: self.delete)
		}
	}

	@Sendable
	func list(req: Request) async throws -> [ProductDTO] {
		try await Product.query(on: req.db).all().map { $0.toDTO() }
	}

	@Sendable
	func create(req: Request) async throws -> ProductDTO {
		let product = try req.content.decode(ProductDTO.self).toModel()

		try await product.save(on: req.db)
		return product.toDTO()
	}

	@Sendable
	func read(req: Request) async throws -> ProductDTO {
		guard let product = try await Product.find(req.parameters.get("pid"), on: req.db) else {
			throw Abort(.notFound)
		}

		return product.toDTO()
	}

	@Sendable
	func update(req: Request) async throws -> ProductDTO {
		let newProduct = try req.content.decode(ProductDTO.self).toModel()
		try await Product.find(req.parameters.get("pid"), on: req.db)
			.unwrap(or: Abort(.notFound))
			.flatMap { product in
				product.name = newProduct.name
				product.category = newProduct.category
				product.description = newProduct.description
				product.price = newProduct.price
				return product.save(on: req.db).map { product }
			}

		return newProduct.toDTO()
	}

	@Sendable
	func delete(req: Request) async throws -> HTTPStatus {
		guard let product = try await Product.find(req.parameters.get("pid"), on: req.db) else {
			throw Abort(.notFound)
		}

		try await product.delete(on: req.db)
		return .noContent
	}
}
