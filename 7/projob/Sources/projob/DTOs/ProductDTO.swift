import Fluent
import Vapor

struct ProductDTO: Content {
	var id: UUID?
	var name: String?
	var category: Cat??
	var description: String?
	var price: UInt32?

	func toModel() -> Product {
		let model = Product()
		
		model.id = self.id

		if let name = self.name {
			model.name = name
		}

		if let category = self.category {
			model.category = category
		}

		if let description = self.description {
			model.description = description
		}

		if let price = self.price {
			model.price = price
		}

		return model
	}
}
