<?php
require_once "../session.php";

$member = checkSession();

if(!$member)
	redirectToLogin();

$program = getProgramm($_GET['id']);
$permissions = getPermissionsForUserForProgram($member['ID'], $_GET['id']);

if($permissions == -1) // read not allowed  // TODO add missing permissions page
	header("Location: ../home/home.php")
?>

<!DOCTYPE html>
<html lang="en">

<head>
	<meta charset="utf-8"/>
	<meta name="viewport" content="width=device-width">
	<title><?= $program["Name"]; ?></title>
	<link rel="stylesheet" href="program.css"/>
	<link rel="stylesheet" href="../mainstyle.css"/>
	<link rel="stylesheet" href="../modal.css"/>
	<link rel="stylesheet" href="../navbar/navbar.css"/>

	<script src="../JS/Websocket.js"></script>
	<script src="../JS/diagramm.js"></script>
	<script src="../JS/utils.js"></script>
	<script src="../JS/programm.js"></script>

	<script src="https://cdn.jsdelivr.net/npm/apexcharts"></script>
</head>

<body>
<?php
include "../navbar/navbar.php";
displayNavbar($member['ID']);
?>

<div id="main">
	<div class="top">
		<h1 class="title"><?= $program["Name"]; ?></h1>
		<div class="topButtonBar">
			<button class="start <?= $member['admin'] || $permissions >= 1 ? '' : 'disabled' ?>  add"><b>Start</b></button>
			<button class="stop <?= $member['admin'] || $permissions >= 2 ? '' : 'disabled' ?>  danger"><b>Stop</b></button>
			<button class="delete <?= $member['admin'] ? '' : 'disabled' ?> danger"><b>Delete</b></button>
		</div>
	</div>
	<div id="boxes">
		<div id="topbar">
			<img src="<?= $program["Imagesource"] ?>" class="boticon" alt=""/>
			<h2 class="description"><?= $program["Description"]; ?></h2>
		</div>
		<div id="activity">
			<div id="activity-chart" style="padding-top: 0.7em;"></div>
		</div>
		<script>
			let chart = new Chart("activity-chart", 'activity')
			chart.add('activity', {x: new Date(), y: 55.23})
			chart.updateChart()
			chart.add('activity', {x: new Date().setMinutes(new Date().getMinutes() + 60), y: 27.51})
			chart.updateChart()
			chart.add('activity', {x: new Date().setMinutes(new Date().getMinutes() + 120), y: 88.51})
			chart.updateChart()
		</script>
		<div id="logs">

		</div>
	</div>
</div>
<script>
	searchModal()
	replaceImages()
	builtWebSocket()
</script>
</body>

</html>
