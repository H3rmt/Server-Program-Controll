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
</head>

<?php
include "authorise.php";

createcookies()
?>

<body>

<?php

include "../navbar/navbar.php";

?>

<div id="main">
	<div class="top">
		<h1 class="title">Settings</h1>
		<button class="authorise" onclick="openmodal()"><b>Authorise</b></button>
	</div>
	<div id="boxes">
		<div class="settingsbox">
			<div class="topsetting">
				<h1>Local settings</h1>
				<button class="reset" onclick="reset('Refresh settings')"><b>Reset to Default</b></button>
			</div>
			<ul class="settings">
				<li class="setting">
					<b>Refresh Delay</b>
				</li>
				<li class="setting">
					<b>Autosave</b>
				</li>
				<li class="setting">
					<b>Autosave</b>
				</li>
			</ul>
		</div>
		<div class="settingsbox disabled protected">
			<div class="topsetting">
				<h1>Client settings</h1>
				<button class="reset protected" onclick="reset('Connection settings')"><b>Reset to Default</b></button>
			</div>
			<ul class="settings">
				<li class="setting">
					<b>Connection</b>
				</li>
				<li class="setting">
					<b>Connection</b>
				</li>
			</ul>
		</div>
		<div class="settingsbox disabled protected">
			<div class="topsetting">
				<h1>Server settings</h1>
				<button class="reset protected" onclick="reset('Other settings')"><b>Reset to Default</b></button>
			</div>
			<table class="settings">
				<td class="setting">
					<h2>
						<label for="new password">Password</label>
					</h2>
					<input id="new password" type="password" name="password">
				</td>
				<td class="setting">
					<h2>
						<label for="new salt">Salt</label>
					</h2>
					<input id="new salt" type="password" name="salt">
				</td>
			</table>
		</div>
	</div>
</div>
<?php
createmodal()
?>
<script>
	searchmodal()

	if (getAuthorisationCookie() !== "") {
		protect();
	}

	disable();
	replaceImages();
</script>
</body>

</html>