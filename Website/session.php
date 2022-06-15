<?php
require_once "database.php";

function checkSession(bool $dropExpired = false): bool|array|string {
	// check if user already has a session and no message is sent (liste logout)
	if(array_key_exists("username", $_COOKIE) && array_key_exists("hash", $_COOKIE) && !array_key_exists("message", $_GET)) {
		$session = getSession($_COOKIE["username"], $_COOKIE["hash"]);
		
		// no session found
		if(!$session)
			return false;
		
		list('ID' => $id, 'expire_date' => $expire_date) = $session;
		
		if($expire_date >= date("Y-m-d H:i:s", time())) {
			// authorise user, because session is valid
			return getMember($_COOKIE["username"]);
		} else {
			// only drop if index site to display expired info
			if($dropExpired) {
				// drop expired session
				dropSession($id);
				
				// drop invalid cookies for invalid session
				setcookie("hash", "");
				return "Session expired";
			}
			return false;
		}
		
	}
	return false;
}


function checkLogin(): string {
	if(array_key_exists("username", $_POST)) {
		$username = $_POST["username"];
		$password = $_POST["password"];
		
		$member = getMember($username);
		
		// no member with username found
		if(!$member)
//			return "Invalid Login {Username}";
			return "Invalid Login";
		
		list('passwd' => $passwd, 'ID' => $id) = $member;
		if(password_verify($password, $passwd)) {
			// login success
			$current_time = time();
			
			$cookie_expiration_time = $current_time + getSetting('timeout') * 60 * 60 * 24;
			
			$hash = bin2hex(random_bytes(30));
			setcookie("username", $username, $cookie_expiration_time);
			setcookie("hash", $hash, $cookie_expiration_time);
			
			$expiry_date = date("Y-m-d H:i:s", $cookie_expiration_time);
			
			// Start new Session
			createSession($id, $hash, $expiry_date);
			header("Location: home/home.php");
			return "Success";
			
		} else {
//			return "Invalid Login {Password}"; // invalid password
			return "Invalid Login"; // invalid password
		}
	}
	// no login (site just opened)
	return "";
}

function checkActions(): string {
	if(array_key_exists("message", $_GET))
		switch($_GET['message']) {
			case 'logout':
				if(!array_key_exists("username", $_COOKIE) || !array_key_exists("hash", $_COOKIE))
					return "error logging out";
				logout($_COOKIE["username"], $_COOKIE["hash"]);
				return "logged out";
			case 'clearsessions':
				if(!array_key_exists("username", $_COOKIE))
					return "error clearing sessions";
				clearSessions($_COOKIE["username"]);
				return "cleared sessions";
			default:
				return 'unknown action';
		}
	return '';
}

function redirectToLogin(int $seconds = 0): void {
	header("refresh:$seconds;url=../index.php");
}