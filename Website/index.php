<?php
require_once "session.php";

if(checkSession())
	header("Location: home");

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
			<input type="submit" name="login" value="Login">
		</div>
	</div>
</form>