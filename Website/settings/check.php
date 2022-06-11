<?php

include_once "../database.php";

function check(bool $admin): void {
	if($_SERVER['REQUEST_METHOD'] == 'POST') {
		
		// Local Settings
		localSettings();
		
		// Client Settings
		if($admin)
			clientSettings();
		
		// Server Settings
		if($admin)
			serverSettings();
	}
}

function localSettings(): void {

}

function clientSettings(): void {
	if(!empty($_POST['new-timeout'])) {
		if($_POST['new-timeout'] != getSetting('timeout')) {
			updateSetting("timeout", $_POST['new-timeout']);
			?>
			<h2 class="feedback">New Timeout Set</h2>
			<?php
		}
	}
}

function serverSettings(): void {

}
