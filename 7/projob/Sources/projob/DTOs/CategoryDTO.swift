import Fluent
import Vapor

struct CategoryDTO: Content {
	var id: UUID?
	var name: String?

	func toModel() -> Cat {
		let model = Cat()
		
		model.id = self.id

		if let name = self.name {
			model.name = name
		}

		return model
	}
}
