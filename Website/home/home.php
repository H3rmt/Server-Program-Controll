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
	<title>Overview</title>
	<link rel="stylesheet" href="homepage.css"/>
	<link rel="stylesheet" href="../mainstyle.css"/>
	<link rel="stylesheet" href="../modal.css"/>
	<link rel="stylesheet" href="../navbar/navbar.css"/>

	<script src="../JS/utils.js"></script>
</head>

<body>
<?php
include "../navbar/navbar.php";
displayNavbar($member['ID']);
?>

<div id="main">
	<div class="top">
		<h1 class="title">Overview</h1>
		<div class="topButtonBar">
			<button class="<?= $member['admin'] ? '' : 'disabled' ?>" onclick="<?= $member['admin'] ? 'openModal()' : '' ?>"><b>New Program</b></button>
		</div>
	</div>
	<div id="boxes">
		<?php
		include "loadBoxes.php";
		displayPrograms($member['ID'], $member['admin'])
		?>
	</div>
</div>
<?php
include "newprogram.php";
displayModal($member['admin']);
?>

<script>
	searchModal();
	replaceImages();
</script>
</body>

</html>