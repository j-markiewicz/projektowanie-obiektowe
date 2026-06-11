import Fluent

struct CreateProducts: AsyncMigration {
	func prepare(on database: any Database) async throws {
		try await database.schema("products")
			.id()
			.field("name", .string, .required)
			.field("category", .uuid, .references("categories", "id"))
			.field("description", .string, .required)
			.field("price", .uint32, .required)
			.create()
	}

	func revert(on database: any Database) async throws {
		try await database.schema("products").delete()
	}
}
