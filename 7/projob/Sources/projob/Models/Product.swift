import Fluent
import struct Foundation.UUID

/// Property wrappers interact poorly with `Sendable` checking, causing a warning for the `@ID` property
/// It is recommended you write your model with sendability checking on and then suppress the warning
/// afterwards with `@unchecked Sendable`.
final class Product: Model, @unchecked Sendable {
	static let schema = "products"

	@ID(key: .id)
	var id: UUID?

	@Field(key: "name")
	var name: String

	@OptionalParent(key: "category")
	var category: Cat?

	@Field(key: "description")
	var description: String

	@Field(key: "price")
	var price: UInt32

	init() { }

	init(id: UUID? = nil, name: String, category: Cat? = nil, description: String, price: UInt32) {
		self.id = id
		self.name = name
		self.category = category
		self.description = description
		self.price = price
	}
		
	func toDTO() -> ProductDTO {
		.init(
			id: self.id,
			name: self.$name.value,
			category: self.$category.value,
			description: self.$description.value,
			price: self.$price.value
		)
	}
}
