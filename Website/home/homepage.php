<!DOCTYPE html>
<html lang="en">

<head>
	<meta charset="utf-8"/>
	<meta name="viewport" content="width=device-width">
	<title>Overview</title>
	<link rel="stylesheet" href="homepage.css"/>
	<link rel="stylesheet" href="../mainstyle.css"/>
	<link rel="stylesheet" href="../modal.css"/>
</head>

<body>
<?php

include "../navbar/navbar.php";
?>

<div id="main">
	<div class="top">
		<h1 class="title">Overview</h1>
		<button class="disabled protected" onclick="openmodal()"><b>New Program</b></button>
	</div>
	<div id="boxes">
		<?php
		include "loadBoxes.php";
		?>
	</div>
</div>
<?php
include "newprogram.php";
?>
<script src="../JS/fade.js"></script>
<script src="../JS/disable buttons.js"></script>
</body>

</html>