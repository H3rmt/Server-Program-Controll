<?php
require_once "../session.php";

$member = checkSession();

if(!$member)
	redirectToLogin();
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

	<script src="../JS/utils.js"></script>
	<script src="../JS/settings.js"></script>
</head>

<?php include "authorise.php"; ?>

<body>
<?php
include "../navbar/navbar.php";
displayNavbar($member['ID']);
?>

<div id="main">
	<div class="top">
		<h1 class="title">Settings</h1>
		<div class="topbuttonbar">
			<button class="logout" onclick="logoutAllSessions()"><b>Logout all Sessions</b></button>
			<button class="logout" onclick="logout()"><b>Logout</b></button>
		</div>
	</div>
	<div class="responses">
		<?php include "check.php"; ?>
	</div>
	<div id="boxes">
		<form id="LocalSettings" class="settingsbox" method="POST">
			<div class="topsetting">
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
								<label for="new_password">Password</label>
							</h2>
							<p>New password for current user</p>
						</div>
						<input id="new_password" type="password" name="password" autocomplete="new-password">
					</td>
				</tr>
			</table>

			<input style="display: none" id="Local settings_reset" type="checkbox" name="resetSettings" autocomplete="off">
		</form>

		<form id="ClientSettings" class="settingsbox disabled" method="POST">
			<div class="topsetting">
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
							<p>Time before the session cookie expires in days</p>
							<p class="additional">Time in s before user sessions expire and new login is required<br>Updating this setting does not
								affect already created sessions</p>
						</div>
						<input id="new_timeout" type="number" name="new-timeout">
					</td>
				</tr>
			</table>

			<input style="display: none" id="Client settings_reset" type="checkbox" name="resetSettings" autocomplete="off">
		</form>

		<form id="ServerSettings" class="settingsbox disabled" method="POST">
			<div class="topsetting">
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

			<input style="display: none" id="Server settings_reset" type="checkbox" name="resetSettings" autocomplete="off">
		</form>
	</div>
</div>
<script>
	searchModal()
	replaceImages()
</script>
</body>

</html>