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
	if(!empty($_POST['resetSettingsLocal'])) {
		setPassword($userId, 'password');
		clearSessions($_COOKIE["username"]);
		redirectToLogin(2);
		?>
		<h2 class="feedback">Reset Local Settings</h2>
		<?php
		return;
	}
	if(!empty($_POST['password']) && !empty($_POST['password_2'])) {
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
	if(!empty($_POST['resetSettingsClient'])) {
		updateSetting("timeout", 30);
		?>
		<h2 class="feedback">Reset Local Settings</h2>
		<?php
		return;
	}
	if(!empty($_POST['timeout'])) {
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
