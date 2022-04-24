<?php

include_once "../database.php";

if($_SERVER['REQUEST_METHOD'] == 'POST') {

	// Local Settings


	$isAdmin = testAdminCookie();

	// Client Settings
	if($isAdmin) {

	}

	// testtesttest

	// Server Settings
	if($isAdmin) {
		if(!empty($_POST['hashed-new-password'])) {
			if(hash('sha256', $_POST['hashed-new-password'] . getPepper()) != getSetting('password')) {
				updateSetting("password", hash('sha256', $_POST['hashed-new-password'] . getPepper()));
				?>
				<h2 class="feedback">New Password Set</h2>
				<?php
			}
		}
		if(!empty($_POST['new-timeout'])) {
			if($_POST['new-timeout'] != getSetting('timeout')) {
				updateSetting("timeout", $_POST['new-timeout']);
				?>
				<h2 class="feedback">New Timeout Set</h2>
				<?php
			}
		}
	}
}