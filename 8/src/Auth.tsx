import { useState } from "react";
import { API_ROOT } from "./main";
import axios from "axios";

function Auth() {
	const [state, setState] = useState("initial");
	const [loginError, setLoginError] = useState("");
	const [signupError, setSignupError] = useState("");
	const [loginEmail, setLoginEmail] = useState("");
	const [loginName, setLoginName] = useState("");
	const [loginPassword, setLoginPassword] = useState("");
	const [signupEmail, setSignupEmail] = useState("");
	const [signupName, setSignupName] = useState("");
	const [signupPassword, setSignupPassword] = useState("");

	const submitLogin = () => {
		setState("busy");

		axios
			.postForm(API_ROOT + "/login", {
				email: loginEmail,
				password: loginPassword,
			})
			.then((res) => {
				setState("done");
				setLoginName(res.data);
				setSignupError("");
				setSignupEmail("");
				setSignupPassword("");
			})
			.catch((err) => {
				if (err.response) {
					setLoginError(
						`Błąd: ${err.response.data} (${err.response.status})`,
					);
				} else {
					setLoginError(`Błąd: ${err} (${err.status})`);
				}

				setState("error");
			});
	};

	const submitSignup = () => {
		setState("busy");

		axios
			.postForm(API_ROOT + "/signup", {
				email: signupEmail,
				name: signupName,
				password: signupPassword,
			})
			.then(() => {
				setState("intial");
				setSignupError("");
				setSignupEmail("");
				setSignupPassword("");
			})
			.catch((err) => {
				if (err.response) {
					setSignupError(
						`Błąd: ${err.response.data} (${err.response.status})`,
					);
				} else {
					setSignupError(`Błąd: ${err} (${err.status})`);
				}

				setState("error");
			});
	};

	return (
		<>
			{state === "done" ? (
				<p>
					Zalogowano jako <em>{loginName}</em>.
				</p>
			) : (
				<div className="forms">
					<form
						onSubmit={(e) => {
							e.preventDefault();
							submitLogin();
						}}
					>
						<h2>Zaloguj</h2>
						{state === "error" ? <p>{loginError}</p> : null}
						<label>
							Adres email:{" "}
							<input
								disabled={state === "busy"}
								size={25}
								minLength={3}
								type="email"
								required
								value={loginEmail}
								onChange={(e) =>
									setLoginEmail(e.currentTarget.value)
								}
							/>
						</label>
						<label>
							Hasło:{" "}
							<input
								disabled={state === "busy"}
								size={25}
								minLength={1}
								maxLength={128}
								type="password"
								required
								value={loginPassword}
								onChange={(e) =>
									setLoginPassword(e.currentTarget.value)
								}
							/>
						</label>

						<button disabled={state === "busy"} type="submit">
							Zaloguj
						</button>
					</form>
					<form
						onSubmit={(e) => {
							e.preventDefault();
							submitSignup();
						}}
					>
						<h2>Utwórz konto</h2>
						{state === "error" ? <p>{signupError}</p> : null}
						<label>
							Imię:{" "}
							<input
								disabled={state === "busy"}
								size={25}
								minLength={1}
								type="text"
								required
								value={signupName}
								onChange={(e) =>
									setSignupName(e.currentTarget.value)
								}
							/>
						</label>
						<label>
							Adres email:{" "}
							<input
								disabled={state === "busy"}
								size={25}
								minLength={3}
								type="email"
								required
								value={signupEmail}
								onChange={(e) =>
									setSignupEmail(e.currentTarget.value)
								}
							/>
						</label>
						<label>
							Hasło:{" "}
							<input
								disabled={state === "busy"}
								size={25}
								minLength={1}
								maxLength={128}
								type="password"
								required
								value={signupPassword}
								onChange={(e) =>
									setSignupPassword(e.currentTarget.value)
								}
							/>
						</label>

						<button disabled={state === "busy"} type="submit">
							Utwórz konto
						</button>
					</form>
				</div>
			)}
		</>
	);
}

export default Auth;
