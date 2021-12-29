<?php

include_once "../database.php";

if($_SERVER['REQUEST_METHOD'] == 'POST') {

	// Local Settings


	$isAdmin = testAdminCookie();

	// Client Settings
	if($isAdmin) {

	}

	// Server Settings
	if($isAdmin) {
		if(!empty($_POST['hashed_new_password'])) {
			updateSetting("password", hash('sha256', $_POST['hashed_new_password'] . getPepper()));
		}
		if(!empty($_POST['pepper'])) {
			setPepper($_POST['pepper']);
		}
		if(!empty($_POST['new_timeout'])) {
			updateSetting("timeout", $_POST['new_timeout']);
		}
	}
}