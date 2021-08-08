<!DOCTYPE html>
<html lang="en">

<head>
	<meta charset="utf-8"/>
	<meta name="viewport" content="width=device-width">
	<title>Settings</title>
	<link rel="stylesheet" href="settings.css"/>
	<link rel="stylesheet" href="../mainstyle.css"/>
	<script src="../JS/fade.js"></script>
</head>

<body>

<?php
include "../navbar/navbar.php";
?>

<div id="main">
	<div class="top">
		<h1>Settings</h1>
		<button class="authorise" onclick="authorise()"><b>Authorise</b></button>
	</div>
	<div id="boxes">
		<div class="settingsbox">
			<div class="topsetting">
				<h1>Refresh settings</h1>
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
		<div class="settingsbox">
			<div class="topsetting">
				<h1>Connection settings</h1>
				<button class="reset" onclick="reset('Connection settings')"><b>Reset to Default</b></button>
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
		<div class="settingsbox">
			<div class="topsetting">
				<h1>Other settings</h1>
				<button class="reset" onclick="reset('Other settings')"><b>Reset to Default</b></button>
			</div>
			<ul class="settings">
				<li class="setting">
					<b>Other 1</b>
				</li>
				<li class="setting">
					<b>Other 2</b>
				</li>
				<li class="setting">
					<b>Other 3</b>
				</li>
			</ul>
		</div>
	</div>
</div>
<script src="../JS/fade.js"></script>
<script src="../JS/disable%20buttons.js"></script>
</body>

</html>