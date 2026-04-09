package example.projektowanieobiektowe.po3

import org.springframework.boot.autoconfigure.SpringBootApplication
import org.springframework.boot.runApplication
import org.springframework.boot.SpringApplication
import org.springframework.web.servlet.support.ServletUriComponentsBuilder
import org.springframework.web.bind.annotation.PathVariable
import org.springframework.web.bind.annotation.GetMapping
import org.springframework.web.bind.annotation.PostMapping
import org.springframework.web.bind.annotation.PutMapping
import org.springframework.web.bind.annotation.PatchMapping
import org.springframework.web.bind.annotation.DeleteMapping
import org.springframework.web.bind.annotation.RequestParam
import org.springframework.web.bind.annotation.RequestBody
import org.springframework.web.bind.annotation.RestController
import org.springframework.http.ResponseEntity
import org.springframework.ui.Model
import org.springframework.ui.set

@SpringBootApplication
class PO3Application

fun main(args: Array<String>) {
	runApplication<PO3Application>(*args)
}

@RestController
class ListController {
	val values: MutableList<String> = mutableListOf("a", "b", "c")

	@GetMapping("/")
	fun index(): List<String> {
		return values
	}
	
	@GetMapping("/{i}")
	fun get(@PathVariable(name = "i") i: String): ResponseEntity<String> {
		try {
			return ResponseEntity.ok(values[i.toInt()])
		} catch (e: Exception) {
			return ResponseEntity.notFound().build()
		}
	}
	
	@PostMapping("/")
	fun post(@RequestBody string: String): ResponseEntity<String> {
		values.add(string)
		val uri = ServletUriComponentsBuilder.fromCurrentRequest()
			.path("/{i}")
			.buildAndExpand(values.size)
			.toUri();
		return ResponseEntity.created(uri).body(string)
	}
	
	@PutMapping("/{i}")
	fun put(@PathVariable(name = "i") i: String, @RequestBody string: String): ResponseEntity<String> {
		try {
			values[i.toInt()] = string
			return ResponseEntity.ok(string)
		} catch (e: Exception) {
			return ResponseEntity.notFound().build()
		}
	}
	
	@PatchMapping("/{i}")
	fun patch(@PathVariable(name = "i") i: String, @RequestBody string: String): ResponseEntity<String> {
		try {
			values[i.toInt()] = string.replace("_", values[i.toInt()])
			return ResponseEntity.ok(values[i.toInt()])
		} catch (e: Exception) {
			return ResponseEntity.notFound().build()
		}
	}

	@DeleteMapping("/{i}")
	fun delete(@PathVariable(name = "i") i: String): ResponseEntity<String> {
		try {
			val value = values[i.toInt()]
			values[i.toInt()] = ""
			return ResponseEntity.ok(value)
		} catch (e: Exception) {
			return ResponseEntity.notFound().build()
		}
	}
}
