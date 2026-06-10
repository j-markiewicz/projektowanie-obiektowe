import { Builder, Browser, By, Key } from "selenium-webdriver";

const TESTS = {
	"accepts obviously valid email": async (drv) => {
		await drv.get("http://localhost:4173/");

		const email = `test${Math.floor(Date.now())}@example.com`;

		await (await drv.findElement(By.css("input"))).click();
		await drv.switchTo().activeElement().sendKeys(Key.TAB);
		await drv.switchTo().activeElement().sendKeys(Key.TAB);
		await drv.switchTo().activeElement().sendKeys(Key.TAB);
		await drv.switchTo().activeElement().sendKeys("test");
		await drv.switchTo().activeElement().sendKeys(Key.TAB);
		await drv.switchTo().activeElement().sendKeys(email);
		await drv.switchTo().activeElement().sendKeys(Key.TAB);
		await drv.switchTo().activeElement().sendKeys("password");
		await drv.switchTo().activeElement().sendKeys(Key.TAB);
		await drv.switchTo().activeElement().click();

		await drv.findElement(By.css("input")).sendKeys(email);
		await drv.switchTo().activeElement().sendKeys(Key.TAB);
		await drv.switchTo().activeElement().sendKeys("password");
		await drv.switchTo().activeElement().sendKeys(Key.TAB);
		await drv.switchTo().activeElement().click();

		if (
			!(
				await (await drv.findElement(By.tagName("body"))).getText()
			).includes("Zalogowano")
		) {
			throw "FAIL";
		}
	},
	"rejects obviously wrong email": async (drv) => {
		await drv.get("http://localhost:4173/");

		const email = `123`;

		await (await drv.findElement(By.css("input"))).click();
		await drv.switchTo().activeElement().sendKeys(Key.TAB);
		await drv.switchTo().activeElement().sendKeys(Key.TAB);
		await drv.switchTo().activeElement().sendKeys(Key.TAB);
		await drv.switchTo().activeElement().sendKeys("test");
		await drv.switchTo().activeElement().sendKeys(Key.TAB);
		await drv.switchTo().activeElement().sendKeys(email);
		await drv.switchTo().activeElement().sendKeys(Key.TAB);
		await drv.switchTo().activeElement().sendKeys("password");
		await drv.switchTo().activeElement().sendKeys(Key.TAB);
		await drv.switchTo().activeElement().click();

		await drv.findElement(By.css("input")).sendKeys(email);
		await drv.switchTo().activeElement().sendKeys(Key.TAB);
		await drv.switchTo().activeElement().sendKeys("password");
		await drv.switchTo().activeElement().sendKeys(Key.TAB);
		await drv.switchTo().activeElement().click();

		if (
			!(
				await (await drv.findElement(By.tagName("body"))).getText()
			).includes("Zaloguj")
		) {
			throw "FAIL";
		}
		if (
			!(
				await (await drv.findElement(By.tagName("body"))).getText()
			).includes("Utwórz konto")
		) {
			throw "FAIL";
		}
	},
	"rejects slightly wrong email": async (drv) => {
		await drv.get("http://localhost:4173/");

		const email = `test..${Math.floor(Date.now())}@example.com`;

		await (await drv.findElement(By.css("input"))).click();
		await drv.switchTo().activeElement().sendKeys(Key.TAB);
		await drv.switchTo().activeElement().sendKeys(Key.TAB);
		await drv.switchTo().activeElement().sendKeys(Key.TAB);
		await drv.switchTo().activeElement().sendKeys("test");
		await drv.switchTo().activeElement().sendKeys(Key.TAB);
		await drv.switchTo().activeElement().sendKeys(email);
		await drv.switchTo().activeElement().sendKeys(Key.TAB);
		await drv.switchTo().activeElement().sendKeys("password");
		await drv.switchTo().activeElement().sendKeys(Key.TAB);
		await drv.switchTo().activeElement().click();

		await drv.findElement(By.css("input")).sendKeys(email);
		await drv.switchTo().activeElement().sendKeys(Key.TAB);
		await drv.switchTo().activeElement().sendKeys("password");
		await drv.switchTo().activeElement().sendKeys(Key.TAB);
		await drv.switchTo().activeElement().click();

		if (
			!(
				await (await drv.findElement(By.tagName("body"))).getText()
			).includes("Zaloguj")
		) {
			throw "FAIL";
		}
		if (
			!(
				await (await drv.findElement(By.tagName("body"))).getText()
			).includes("Utwórz konto")
		) {
			throw "FAIL";
		}
	},
	"doesn't run scripts": async (drv) => {
		await drv.get("http://localhost:4173/");

		const email = `test${Math.floor(Date.now())}@example.com`;

		await (await drv.findElement(By.css("input"))).click();
		await drv.switchTo().activeElement().sendKeys(Key.TAB);
		await drv.switchTo().activeElement().sendKeys(Key.TAB);
		await drv.switchTo().activeElement().sendKeys(Key.TAB);
		await drv
			.switchTo()
			.activeElement()
			.sendKeys("<script>alert(1)</script>");
		await drv.switchTo().activeElement().sendKeys(Key.TAB);
		await drv.switchTo().activeElement().sendKeys(email);
		await drv.switchTo().activeElement().sendKeys(Key.TAB);
		await drv.switchTo().activeElement().sendKeys("password");
		await drv.switchTo().activeElement().sendKeys(Key.TAB);
		await drv.switchTo().activeElement().click();

		await drv.findElement(By.css("input")).sendKeys(email);
		await drv.switchTo().activeElement().sendKeys(Key.TAB);
		await drv.switchTo().activeElement().sendKeys("password");
		await drv.switchTo().activeElement().sendKeys(Key.TAB);
		await drv.switchTo().activeElement().click();

		if (
			!(
				await (await drv.findElement(By.tagName("body"))).getText()
			).includes("Zalogowano")
		) {
			throw "FAIL";
		}
		if (
			!(
				await (await drv.findElement(By.tagName("body"))).getText()
			).includes("<script>")
		) {
			throw "FAIL";
		}
	},
	"doesn't allow text formatting": async (drv) => {
		await drv.get("http://localhost:4173/");

		const email = `test${Math.floor(Date.now())}@example.com`;

		await (await drv.findElement(By.css("input"))).click();
		await drv.switchTo().activeElement().sendKeys(Key.TAB);
		await drv.switchTo().activeElement().sendKeys(Key.TAB);
		await drv.switchTo().activeElement().sendKeys(Key.TAB);
		await drv.switchTo().activeElement().sendKeys("<h1>NAME</h1>");
		await drv.switchTo().activeElement().sendKeys(Key.TAB);
		await drv.switchTo().activeElement().sendKeys(email);
		await drv.switchTo().activeElement().sendKeys(Key.TAB);
		await drv.switchTo().activeElement().sendKeys("password");
		await drv.switchTo().activeElement().sendKeys(Key.TAB);
		await drv.switchTo().activeElement().click();

		await drv.findElement(By.css("input")).sendKeys(email);
		await drv.switchTo().activeElement().sendKeys(Key.TAB);
		await drv.switchTo().activeElement().sendKeys("password");
		await drv.switchTo().activeElement().sendKeys(Key.TAB);
		await drv.switchTo().activeElement().click();

		if (
			!(
				await (await drv.findElement(By.tagName("body"))).getText()
			).includes("Zalogowano")
		) {
			throw "FAIL";
		}
		if (
			!(
				await (await drv.findElement(By.tagName("body"))).getText()
			).includes("Zalogowano")
		) {
			throw "FAIL";
		}
	},
	"doesn't allow images": async (drv) => {
		await drv.get("http://localhost:4173/");

		const email = `test${Math.floor(Date.now())}@example.com`;

		await (await drv.findElement(By.css("input"))).click();
		await drv.switchTo().activeElement().sendKeys(Key.TAB);
		await drv.switchTo().activeElement().sendKeys(Key.TAB);
		await drv.switchTo().activeElement().sendKeys(Key.TAB);
		await drv
			.switchTo()
			.activeElement()
			.sendKeys('<img src="invalid" onerror="alert(1)" />');
		await drv.switchTo().activeElement().sendKeys(Key.TAB);
		await drv.switchTo().activeElement().sendKeys(email);
		await drv.switchTo().activeElement().sendKeys(Key.TAB);
		await drv.switchTo().activeElement().sendKeys("password");
		await drv.switchTo().activeElement().sendKeys(Key.TAB);
		await drv.switchTo().activeElement().click();

		await drv.findElement(By.css("input")).sendKeys(email);
		await drv.switchTo().activeElement().sendKeys(Key.TAB);
		await drv.switchTo().activeElement().sendKeys("password");
		await drv.switchTo().activeElement().sendKeys(Key.TAB);
		await drv.switchTo().activeElement().click();

		if (
			!(
				await (await drv.findElement(By.tagName("body"))).getText()
			).includes("Zalogowano")
		) {
			throw "FAIL";
		}
		if (
			!(
				await (await drv.findElement(By.tagName("body"))).getText()
			).includes("onerror")
		) {
			throw "FAIL";
		}
	},
};

(async function example() {
	let driver = await new Builder().forBrowser(Browser.FIREFOX).build();

	try {
		for (let [name, test] of Object.entries(TESTS)) {
			console.info(`${name}:`);
			await test(driver);
			console.info(`OK`);
		}
	} catch {
		console.info(`FAIL`);
	} finally {
		await driver.quit();
	}
})();
