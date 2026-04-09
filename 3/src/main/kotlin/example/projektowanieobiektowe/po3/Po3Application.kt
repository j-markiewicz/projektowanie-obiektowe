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
import org.springframework.web.bind.annotation.RequestHeader
import org.springframework.http.ResponseEntity
import org.springframework.http.HttpHeaders
import org.springframework.ui.Model
import org.springframework.ui.set
import org.springframework.security.crypto.argon2.Argon2PasswordEncoder
import kotlin.io.encoding.Base64

@SpringBootApplication
class PO3Application

object Auth {
	val hasher = Argon2PasswordEncoder(16, 32, 1, 65536, 5)
	val mock_db = mapOf(
		"username" to $$"$argon2id$v=19$m=65536,t=5,p=1$c2FsdHNhbHQ$py5qptouKWGvaRTZenfK0k68h9m2NghrpApn0ZGO7vI" // "password"
	)

	fun authenticate(auth_header: String?): Boolean {
		try {
			if (auth_header?.startsWith("Basic ") != true) {
				return false
			}

			val auth_token = auth_header?.removePrefix("Basic ") ?: ""
			val user_pass = Base64.Default.decode(auth_token).decodeToString()
			if (!user_pass.contains(":")) {
				return false
			}

			val (username, password) = user_pass.split(":", limit = 2)
			val password_hash = mock_db[username]

			if (password_hash == null) {
				return false
			}

			return hasher.matches(
				password,
				password_hash
			)
		} catch (e: Exception) {
			return false
		}
	}

	fun <T: Any>rejection(): ResponseEntity<T> = ResponseEntity.status(401).header("WWW-Authenticate", "Basic").build()
}

fun main(args: Array<String>) {
	runApplication<PO3Application>(*args)
}

@RestController
class ListController {
	val values: MutableList<String> = mutableListOf("a", "b", "c")

	@GetMapping("/")
	fun index(@RequestHeader(HttpHeaders.AUTHORIZATION, required = false) auth_header: String?): ResponseEntity<List<String>> {
		if (!Auth.authenticate(auth_header)) {
			return Auth.rejection()
		}

		return ResponseEntity.ok(values)
	}
	
	@GetMapping("/{i}")
	fun get(@PathVariable(name = "i") i: String, @RequestHeader(HttpHeaders.AUTHORIZATION, required = false) auth_header: String?): ResponseEntity<String> {
		if (!Auth.authenticate(auth_header)) {
			return Auth.rejection()
		}

		try {
			return ResponseEntity.ok(values[i.toInt()])
		} catch (e: Exception) {
			return ResponseEntity.notFound().build()
		}
	}
	
	@PostMapping("/")
	fun post(@RequestBody string: String, @RequestHeader(HttpHeaders.AUTHORIZATION, required = false) auth_header: String?): ResponseEntity<String> {
		if (!Auth.authenticate(auth_header)) {
			return Auth.rejection()
		}

		values.add(string)
		val uri = ServletUriComponentsBuilder.fromCurrentRequest()
			.path("/{i}")
			.buildAndExpand(values.size)
			.toUri();
		return ResponseEntity.created(uri).body(string)
	}
	
	@PutMapping("/{i}")
	fun put(@PathVariable(name = "i") i: String, @RequestBody string: String, @RequestHeader(HttpHeaders.AUTHORIZATION, required = false) auth_header: String?): ResponseEntity<String> {
		if (!Auth.authenticate(auth_header)) {
			return Auth.rejection()
		}

		try {
			values[i.toInt()] = string
			return ResponseEntity.ok(string)
		} catch (e: Exception) {
			return ResponseEntity.notFound().build()
		}
	}
	
	@PatchMapping("/{i}")
	fun patch(@PathVariable(name = "i") i: String, @RequestBody string: String, @RequestHeader(HttpHeaders.AUTHORIZATION, required = false) auth_header: String?): ResponseEntity<String> {
		if (!Auth.authenticate(auth_header)) {
			return Auth.rejection()
		}

		try {
			values[i.toInt()] = string.replace("_", values[i.toInt()])
			return ResponseEntity.ok(values[i.toInt()])
		} catch (e: Exception) {
			return ResponseEntity.notFound().build()
		}
	}

	@DeleteMapping("/{i}")
	fun delete(@PathVariable(name = "i") i: String, @RequestHeader(HttpHeaders.AUTHORIZATION, required = false) auth_header: String?): ResponseEntity<String> {
		if (!Auth.authenticate(auth_header)) {
			return Auth.rejection()
		}

		try {
			val value = values[i.toInt()]
			values[i.toInt()] = ""
			return ResponseEntity.ok(value)
		} catch (e: Exception) {
			return ResponseEntity.notFound().build()
		}
	}
}
