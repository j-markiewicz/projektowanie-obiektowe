import Fluent
import Vapor

struct CategoriesController: RouteCollection {
	func boot(routes: any RoutesBuilder) throws {
		let categories = routes.grouped("categories")

		categories.get(use: self.list)
		categories.post(use: self.create)
		categories.group(":pid") { category in
			category.get(use: self.read)
			category.put(use: self.update)
			category.delete(use: self.delete)
		}
	}

	@Sendable
	func list(req: Request) async throws -> [CategoryDTO] {
		try await Cat.query(on: req.db).all().map { $0.toDTO() }
	}

	@Sendable
	func create(req: Request) async throws -> CategoryDTO {
		let category = try req.content.decode(CategoryDTO.self).toModel()

		try await category.save(on: req.db)
		return category.toDTO()
	}

	@Sendable
	func read(req: Request) async throws -> CategoryDTO {
		guard let category = try await Cat.find(req.parameters.get("pid"), on: req.db) else {
			throw Abort(.notFound)
		}

		return category.toDTO()
	}

	@Sendable
	func update(req: Request) async throws -> CategoryDTO {
		let newCategory = try req.content.decode(CategoryDTO.self).toModel()
		try await Cat.find(req.parameters.get("pid"), on: req.db)
			.unwrap(or: Abort(.notFound))
			.flatMap { category in
				category.name = newCategory.name
				return category.save(on: req.db).map { category }
			}

		return newCategory.toDTO()
	}

	@Sendable
	func delete(req: Request) async throws -> HTTPStatus {
		guard let category = try await Cat.find(req.parameters.get("pid"), on: req.db) else {
			throw Abort(.notFound)
		}

		try await category.delete(on: req.db)
		return .noContent
	}
}
