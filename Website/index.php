<?php
require_once "session.php";

// check if user already logged in or just logged in
$check = checkSession(true);

// redirect if logged in
if($check === true)
	header("Location: home/home.php");

//
$actions = checkActions();

$login = checkLogin();

// check login try
if($login !== '')           // message like "Invalid Login"  (never get triggered after redirekt or on first visit)
	$message = $login;
else if($actions !== '')    // message like "logged out" / "cleared sessions"
	$message = $actions;
else                        // message like "Session expired"
	$message = $check;


// psw: aILwxKROgvVHXmYdcqhSHugRGQxjqvcoYNFaMpAkEQXyxIAtOxjumfprrJKwqHdIkvgeh
// db:  $2y$10$a9QQCHHEJ7bZ//D1My0oAOMRQXtRZIGV0YK5emhasFR9xusyBvyha

# password_hash('test', null);
?>

<!DOCTYPE html>
<html lang="en">

<head>
	<meta charset="utf-8"/>
	<meta name="viewport" content="width=device-width">
	<title>Login</title>
	<link rel="stylesheet" href="login.css"/>
	<link rel="stylesheet" href="mainstyle.css"/>
	<link rel="stylesheet" href="modal.css"/>
</head>

<body>
<div class="modal" id="closable-modal">
	<form id="login" class="Popup" method="post" action="index.php">
		<table>
			<tr>
				<td>
					<label for="username">Username</label>
				</td>
				<td>
					<input id="username" name="username" type="text" value="<?= $_COOKIE['username'] ?? '' ?>">
				</td>
			</tr>
			<tr>
				<td>
					<label for="password">Password</label>
				</td>
				<td>
					<input id="password" name="password" type="password">
				</td>
			</tr>
		</table>
		<div id="bottom">
			<h3 id="message"><?= $message ?></h3>
			<button type="submit" class="save"><b>Login</b>
		</div>
	</form>
</div>

</body>

</html>