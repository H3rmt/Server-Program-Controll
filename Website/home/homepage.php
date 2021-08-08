<!DOCTYPE html>
<html lang="en">

<head>
	<meta charset="utf-8"/>
	<title>Overview</title>
	<link rel="stylesheet" href="homepage.css"/>
	<link rel="stylesheet" href="../mainstyle.css"/>
</head>

<body>
<ul id="navbar">
	<li>
		<a id="Overview" href="homepage.php">Overview</a>
	</li>
	<?php
	include "../navbar/navbar.php";
	?>
</ul>

<div id="main">
	<div class="top">
		<h1>Overview</h1>
		<button class="new disabled" onclick="opennewprogramm()"><b>New Program</b></button>
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
<script src="../JS/disable%20buttons.js"></script>
</body>

</html>