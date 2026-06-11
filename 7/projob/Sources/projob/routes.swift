import Fluent
import Vapor

func routes(_ app: Application) throws {
	app.get { req async throws -> View in
		let products = try await Product.query(on: req.db).all().map { $0.toDTO() }

		return try await req.view.render("products", ["products": products])
	}

	app.get(":pid") { req async throws -> View in
		guard let product = try await Product.find(req.parameters.get("pid"), on: req.db) else {
			throw Abort(.notFound)
		}

		return try await req.view.render("product", ["product": product])
	}

	try app.register(collection: ProductsController())
	try app.register(collection: CategoriesController())
}
