<?php
require_once "session.php";

if(checkSession())
	header("Location: home/home.php");

// check login try
$message = checkLogin()


// psw: aILwxKROgvVHXmYdcqhSHugRGQxjqvcoYNFaMpAkEQXyxIAtOxjumfprrJKwqHdIkvgeh
// db:  $2y$10$a9QQCHHEJ7bZ//D1My0oAOMRQXtRZIGV0YK5emhasFR9xusyBvyha
?>

<!DOCTYPE html>
<html lang="en">

<head>
	<meta charset="utf-8"/>
	<meta name="viewport" content="width=device-width">
	<title>Login</title>
	<link rel="stylesheet" href="login.css"/>
	<link rel="stylesheet" href="mainstyle.css"/>
</head>

<body>

<div id="login">
	<form method="post">
		<div>
			<label for="username">Username</label>
			<input id="username" name="username" type="text">
		</div>
		<div>
			<label for="password">Password</label>
			<input id="password" name="password" type="password">
		</div>
		<div id="bottom">
			<h3 id="message"><?= $message ?></h3>
			<input type="submit" name="login" value="Login">
		</div>
	</form>
</div>

</body>

</html>