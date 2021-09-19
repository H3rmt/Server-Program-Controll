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
				<div>
					<button class="save protected" onclick="save('Local settings')"><b>Save</b></button>
					<button class="reset protected" onclick="reset('Local settings')"><b>Reset to Default</b></button>
				</div>	
			</div>
			<table class="settings">
				
			</table>
		</div>
		<div class="settingsbox disabled protected">
		<div class="topsetting">
				<h1>Client settings</h1>
				<div>
					<button class="save protected" onclick="save('Client settings')"><b>Save</b></button>
					<button class="reset protected" onclick="reset('Client settings')"><b>Reset to Default</b></button>
				</div>	
			</div>
			<table class="settings">
				
			</table>
		</div>
		<div class="settingsbox disabled protected">
			<div class="topsetting">
				<h1>Server settings</h1>
				<div>
					<button class="save protected" onclick="save('Server settings')"><b>Save</b></button>
					<button class="reset protected" onclick="reset('Server settings')"><b>Reset to Default</b></button>
				</div>	
			</div>
			<table class="settings">
				<tr>
				<td class="setting">
					<div>
					<h2>	
						<label for="new password">Password</label>
					</h2>		
					<p>New password to authorise and gain admin privileges</p>
					</div>
					<input id="new password" type="password" name="password">
				</td>
				</tr>
				<tr><td class="seperator"></td></tr>
				<tr>
				<td class="setting">
					<div>
					<h2>
						<label for="new salt">Pepper</label>
					</h2>
					<p>Pepper is added to the password bevore Hash</p>
					</div>
					<input id="new salt" type="password" name="salt">
				</td>
				</tr>
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