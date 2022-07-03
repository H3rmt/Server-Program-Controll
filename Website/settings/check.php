<?php

include_once "../database.php";

function check(bool $admin, int $userId): void {
	if($_SERVER['REQUEST_METHOD'] == 'POST') {
		?>
		<div class="responses"><?php
		
		// Local Settings
		localSettings($userId);
		
		// Client Settings
		if($admin)
			clientSettings($userId);
		
		// Server Settings
		if($admin)
			serverSettings($userId);
		?></div><?php
	}
}


function localSettings(int $userId): void {
	if(array_key_exists('resetSettingsLocal', $_POST)) {
		setPassword($userId, 'password');
		clearSessions($_COOKIE["username"]);
		redirectToLogin(2);
		?>
		<h2 class="feedback">Reset Local Settings</h2>
		<?php
		return;
	}
	if(array_key_exists('password', $_POST) && array_key_exists('password_2', $_POST)) {
		if($_POST['password'] != $_POST['password_2']) {
			?>
			<h2 class="feedback feedbackError">Failed to set new Password</h2>
			<?php
		} else {
			setPassword($userId, $_POST['password']);
			clearSessions($_COOKIE["username"]);
			redirectToLogin(2);
			?>
			<h2 class="feedback">New Password Set</h2>
			<?php
		}
	}
}

function clientSettings(int $userId): void {
	if(array_key_exists('resetSettingsClient', $_POST)) {
		updateSetting("timeout", 30);
		?>
		<h2 class="feedback">Reset Client Settings</h2>
		<?php
		return;
	}
	if(array_key_exists('timeout', $_POST)) {
		if($_POST['timeout'] != getSetting('timeout')) {
			updateSetting("timeout", $_POST['timeout']);
			?>
			<h2 class="feedback">New Timeout Set</h2>
			<?php
		}
	}
}

function serverSettings(int $userId): void {

}
