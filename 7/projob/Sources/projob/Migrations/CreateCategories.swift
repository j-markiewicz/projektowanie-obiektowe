import Fluent

struct CreateCategories: AsyncMigration {
	func prepare(on database: any Database) async throws {
		try await database.schema("categories")
			.id()
			.field("name", .string, .required)
			.create()
	}

	func revert(on database: any Database) async throws {
		try await database.schema("categories").delete()
	}
}
