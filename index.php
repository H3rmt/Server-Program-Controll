<?php
require_once "database.php";

function checkSession(): bool {
	// check if user already has a session
	if(!empty($_COOKIE["username"]) && !empty($_COOKIE["hash"])) {
		$session = getSession($_COOKIE["username"], $_COOKIE["hash"]);

		// no session found
		if(!$session)
			return false;

		list('ID' => $id, 'expire_date' => $expire_date) = $session;


		if($expire_date >= date("Y-m-d H:i:s", time())) {
			// authorise user, because session is valid
			return true;
		} else {
			// drop expired session
			dropSession($id);

			// drop invalid cookies for invalid session
			setcookie("username", "");
			setcookie("hash", "");
			return false;
		}

	}
	return false;
}

function checkLogin(): string {
	if(!empty($_POST["login"])) {
		$username = $_POST["username"];
		$password = $_POST["password"];
		$remember = !empty($_POST["remember"]);

		$member = getMember($username);

		// no member with username found
		if(!$member)
			return "Invalid Login {Username}";


		list('passwd' => $passwd) = $member;
		if(password_verify($password, $passwd)) {
			if($remember) {
				$current_time = time();

				// Set Cookie expiration for 1 month
				$cookie_expiration_time = $current_time + (30 * 24 * 60 * 60);  // for 1 month

				$hash = bin2hex(random_bytes(30));
				setcookie("username", $username, $cookie_expiration_time);
				setcookie("hash", $hash, $cookie_expiration_time);

				$expiry_date = date("Y-m-d H:i:s", $cookie_expiration_time);

				// Start new Session
				createSession($username, $hash, $expiry_date);
			} else {
				// clear cookies if existed
				setcookie("username", "");
				setcookie("password", "");
				setcookie("selector", "");
			}
			header("Location: Website/home/homepage.php");
		} else {
			return "Invalid Login {Password}"; // invalid password
		}
	}
	return "";
}

// check for valid session in cookies
if(checkSession()) {
	header("Location: Website/home/homepage.php");
}

// check login try
$message = checkLogin()


// psw: aILwxKROgvVHXmYdcqhSHugRGQxjqvcoYNFaMpAkEQXyxIAtOxjumfprrJKwqHdIkvgeh
// db:  $2y$10$a9QQCHHEJ7bZ//D1My0oAOMRQXtRZIGV0YK5emhasFR9xusyBvyha
?>

<form method="post">
	<div><?= $message ?></div>
	<div>
		<div>
			<label for="username">Username</label>
		</div>
		<div>
			<label>
				<input name="username" type="text" value="<?php if(isset($_COOKIE["username"])) {
					echo $_COOKIE["username"];
				} ?>">
			</label>
		</div>
	</div>
	<div>
		<div>
			<label for="password">Password</label>
		</div>
		<div>
			<label>
				<input name="password" type="password">
			</label>
		</div>
	</div>
	<div>
		<div>
			<label for="remember">Remember me</label>
			<input type="checkbox" name="remember" id="remember"<?php if(isset($_COOKIE["username"])) { ?> checked<?php } ?> />
		</div>
	</div>
	<div>
		<div>
			<input type="submit" name="login" value="Login">
		</div>
	</div>
</form>