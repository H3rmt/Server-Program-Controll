<?php
require_once "authDatabase.php";
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
		
		$member = getMember($username);
		
		// no member with username found
		if(!$member)
//			return "Invalid Login {Username}";
			return "Invalid Login";
		
		
		list('passwd' => $passwd) = $member;
		if(password_verify($password, $passwd)) {
			// login success
			$current_time = time();
			
			$cookie_expiration_time = $current_time + getSetting('timeout') * 24 * 60 * 60;
			
			$hash = bin2hex(random_bytes(30));
			setcookie("username", $username, $cookie_expiration_time);
			setcookie("hash", $hash, $cookie_expiration_time);
			
			$expiry_date = date("Y-m-d H:i:s", $cookie_expiration_time);
			
			// Start new Session
			createSession($username, $hash, $expiry_date);
			header("refresh: 0"); //refresh to let checkSession detect session
		} else {
//			return "Invalid Login {Password}"; // invalid password
			return "Invalid Login"; // invalid password
		}
	}
	// no login / site just opened
	return "";
}


function redirectToLogin(): void {
	header("Location: ..");
}