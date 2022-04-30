<?php
require_once "../session.php";

if(!checkSession())
	redirectToLogin();
?>

<!DOCTYPE html>
<html lang="en">

<head>
	<meta charset="utf-8"/>
	<meta name="viewport" content="width=device-width">
	<title>Overview</title>
	<link rel="stylesheet" href="homepage.css"/>
	<link rel="stylesheet" href="../mainstyle.css"/>
	<link rel="stylesheet" href="../modal.css"/>

	<script src="../JS/utils.js"></script>
</head>

<body>
<?php
include "../navbar/navbar.php";
?>

<div id="main">
	<div class="top">
		<h1 class="title">Overview</h1>
		<div class="topbuttonbar">
			<button class="disabled" onclick="openmodal()"><b>New Program</b></button>
		</div>
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

<script>
	searchmodal();
	replaceImages();
	protect();
	disable();
</script>
</body>

</html>