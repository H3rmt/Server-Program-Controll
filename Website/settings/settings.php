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
	<script src="../JS/sha256.js"></script>
	<script src="../JS/settings.js"></script>
</head>

<?php include "authorise.php"; ?>

<body>

<?php include "../navbar/navbar.php"; ?>


<div id="main">
	<div class="top">
		<h1 class="title">Settings</h1>
		<div class="topbuttonbar">
			<button class="authorise" onclick="openmodal()"><b>Authorise</b></button>
		</div>
	</div>
	<div class="responses">
		<?php include "check.php"; ?>
	</div>
	<?php
	if(testAdminCookie()) {
		$timeout = getSetting('timeout');
		$pepper = getPepper();
	}
	?>
	<div id="boxes">
		<form id="LocalSettings" class="settingsbox" method="POST">
			<div class="topsetting">
				<h1>Local settings</h1>
				<div>
					<button class="reset danger" onclick="rresetSettings('Local settings',false,event)"><b>Restore Defaults</b></button>
					<button class="save" onclick="saveSettings('Local settings',false,event)"><b>Save</b></button>
				</div>
			</div>
			<table class="settings">

			</table>

			<input style="display: none" id="Local settings_reset" type="checkbox" name="resetSettings" autocomplete="off">

			<script>
				document.getElementById("LocalSettings").addEventListener("submit", () => {
				})
			</script>
		</form>

		<form id="ClientSettings" class="settingsbox disabled" method="POST">
			<div class="topsetting">
				<h1>Client settings</h1>
				<div>
					<button class="reset danger" onclick="rresetSettings('Client settings',true,event);"><b>Restore Defaults</b></button>
					<button class="save" onclick="saveSettings('Client settings',true,event);"><b>Save</b></button>
				</div>
			</div>
			<table class="settings">

			</table>

			<input style="display: none" id="Client settings_reset" type="checkbox" name="resetSettings" autocomplete="off">

			<script>
				document.getElementById("ClientSettings").addEventListener("submit", () => {
				})
			</script>
		</form>

		<form id="ServerSettings" class="settingsbox disabled" method="POST">
			<div class="topsetting">
				<h1>Server settings</h1>
				<div>
					<button class="reset danger" onclick="rresetSettings('Server settings',true,event)"><b>Restore Defaults</b></button>
					<button class="save" onclick="saveSettings('Server settings',true,event)"><b>Save</b></button>
				</div>
			</div>
			<table class="settings">
				<tr>
					<td class="setting">
						<div>
							<h2>
								<label for="new_password">Password</label>
							</h2>
							<p>New password to authorise and gain admin privileges</p>
						</div>
						<input id="new_password" type="password" name="password" autocomplete="new-password">
					</td>
				</tr>
				<tr>
					<td class="setting">
						<div>
							<h2>
								<label for="new_timeout">Login Timeout</label>
							</h2>
							<p>Time bevore the login cookie expires in sec</p>
							<p class="additional">86400 seconds = 24 hours; 345600 seconds = 4 days</p>
						</div>
						<input id="new_timeout" type="number" name="new-timeout" value=<?= $timeout ?? '' ?>>
					</td>
				</tr>
				<tr>
					<td class="seperator"></td>
				</tr>
			</table>

			<input style="display: none" id="Server settings_reset" type="checkbox" name="resetSettings" autocomplete="off">
			<input style="display: none" id="hashed_new_password" type="password" name="hashed-new-password" autocomplete="off">
			<script>
				document.getElementById("ServerSettings").addEventListener("submit", () => {
					if (document.getElementById("new_password").value.length !== 0) {
						document.getElementById("hashed_new_password").value = SHA256(document.getElementById("new_password").value)
						document.getElementById("new_password").value = "*".repeat(document.getElementById("new_password").value.length)
					}
				})
			</script>
		</form>
	</div>
</div>
<?php
createmodal()
?>
<script>
	searchmodal()
	protect()
	disable()
	replaceImages();
</script>
</body>

</html>