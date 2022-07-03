<?php
require_once "../session.php";

$member = checkSession();

if(!$member) {
	redirectToLogin();
	exit();
}
?>

<!DOCTYPE html>
<html lang="en">

<head>
	<meta charset="utf-8"/>
	<meta name="viewport" content="width=device-width">
	<title>Settings</title>
	<link rel="stylesheet" href="settings.css"/>
	<link rel="stylesheet" href="../mainstyle.css"/>
	<link rel="stylesheet" href="../modal.css"/>
	<link rel="stylesheet" href="../navbar/navbar.css"/>

	<script src="../JS/utils.js"></script>
	<script src="../JS/settings.js"></script>
</head>

<body>
<?php
include "../navbar/navbar.php";
displayNavbar($member['ID']);
?>

<div id="main">
	<div class="top">
		<h1 class="title">Settings</h1>
		<div class="topButtonBar">
			<button class="logout danger" onclick="logoutAllSessions()"><b>Logout all Sessions</b></button>
			<button class="logout" onclick="logout()"><b>Logout</b></button>
		</div>
	</div>
	<?php
	include "check.php";
	check($member['admin'], $member['ID']);
	?>
	<div id="boxes">
		<form id="LocalSettings" class="settingsBox" method="POST">
			<div class="topSetting">
				<h1>Local settings</h1>
				<div class="buttons">
					<button type="submit" class="save" onclick="saveSettings('Local settings',true,event)"><b>Save</b></button>
					<button class="reset danger" onclick="rresetSettings('Local settings',true,event)"><b>Restore Defaults</b></button>
				</div>
			</div>
			<table class="settings">
				<tr>
					<td class="setting">
						<div>
							<h2>
								<label for="new_password">New Password</label>
							</h2>
							<p class="additional">New password for current user</p>
							<p class="additional">Changing a user-password logs out all current active sessions</p>
							<p class="additional">Default: password</p>
						</div>
						<input id="new_password" type="password" name="password" autocomplete="new-password">
					</td>
				</tr>
				<tr>
					<td class="setting">
						<div>
							<h2>
								<label for="new_password_2">Repeat New Password</label>
							</h2>
						</div>
						<input id="new_password_2" type="password" name="password_2" autocomplete="new-password">
					</td>
				</tr>
			</table>

			<input style="display: none" id="Local settings_reset" type="checkbox" name="resetSettingsLocal" autocomplete="off" value="false">
		</form>

		<form id="ClientSettings" class="settingsBox <?= $member['admin'] ? '' : 'disabled' ?>" method="POST">
			<div class="topSetting">
				<h1>Client settings</h1>
				<div class="buttons">
					<button type="submit" class="save" onclick="saveSettings('Client settings',<?= $member['admin'] ?>,event);"><b>Save</b>
					</button>
					<button class="reset danger" onclick="rresetSettings('Client settings',<?= $member['admin'] ?>,event);"><b>Restore
							Defaults</b></button>
				</div>
			</div>
			<table class="settings">
				<tr>
					<td class="setting">
						<div>
							<h2>
								<label for="new_timeout">Session Timeout</label>
							</h2>
							<p class="additional">Time in days before user session expires and new
								login is required<br>Updating this setting does not
								affect already created sessions</p>
							<p class="additional">Default: 30</p>
						</div>
						<input id="new_timeout" type="number" name="timeout">
					</td>
				</tr>
			</table>

			<input style="display: none" id="Client settings_reset" type="checkbox" name="resetSettingsClient" autocomplete="off"
					 value="false">
		</form>

		<form id="ServerSettings" class="settingsBox <?= $member['admin'] ? '' : 'disabled' ?>" method="POST">
			<div class="topSetting">
				<h1>Server settings</h1>
				<div class="buttons">
					<button type="submit" class="save" onclick="saveSettings('Server settings',<?= $member['admin'] ?>,event)"><b>Save</b>
					</button>
					<button class="reset danger" onclick="rresetSettings('Server settings',<?= $member['admin'] ?>,event)"><b>Restore
							Defaults</b></button>
				</div>
			</div>
			<table class="settings">

				<!--				<tr>-->
				<!--					<td class="seperator"></td>-->
				<!--				</tr>-->
			</table>

			<input style="display: none" id="Server settings_reset" type="checkbox" name="resetSettingsServer" autocomplete="off"
					 value="false">
		</form>
	</div>
</div>
<script>
	searchModal()
	replaceImages()
</script>
</body>

</html>