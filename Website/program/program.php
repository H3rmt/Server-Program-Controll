<!DOCTYPE html>
<html lang="en">
<?php
include_once "../database.php";

$program = getprogramm(htmlspecialchars(stripslashes(trim($_GET['id']))));
?>

<head>
	<meta charset="utf-8"/>
	<meta name="viewport" content="width=device-width">
	<title><?= $program["Name"]; ?></title>
	<link rel="stylesheet" href="../mainstyle.css"/>
	<link rel="stylesheet" href="program.css"/>
	<script src="../JS/sha256.js"></script>
	<script src="../JS/Websocket.js"></script>
	<script>
		function recivelogs(data) {
			data.forEach((log) => {
				console.log("log: ", log);
			});
		}
		
		function reciveactivity(data) {
			data.forEach((activity) => {
				console.log("activity: ", activity);
			});
		}
		
		function recivestart(data, error) {
			if (error != null) {
				if (error === "no admin permissions") {
					console.log("no admin permissions");
				} else {
					console.log("Error");
				}
				return;
			}
			if (data === true) {
				console.log("success", data);
			} else {
				console.log("not successfull", data);
			}
		}
		
		function recivestop(data, error) {
			if (error != null) {
				if (error === "no admin permissions") {
					console.log("no admin permissions");
				} else {
					console.log("Error");
				}
				return;
			}
			if (data === true) {
				console.log("success", data);
			} else {
				console.log("not successfull", data);
			}
		}
	</script>
</head>

<body>
<?php
include "../navbar/navbar.php";
?>


<div id="main">
	<div class="top">
		<h1><?= $program["Name"]; ?></h1>
		<div>
			<button class="start disabled protected"><b>Start</b></button>
			<button class="stop disabled protected"><b>Stop</b></button>
			<button class="delete disabled protected"><b>Delete</b></button>
		</div>
	</div>
	<div id="boxes">
		<div id="topbar">
			<img src="<?= file_exists($program["Imagesource"]) ? $program["Imagesource"] : '../Images/imgnotfound.png'; ?>" class="boticon"
			     alt=""/>
			<h2 class="description"><?= $program["Description"]; ?></h2>
		</div>
		<div id="activity">
		
		</div>
		<div id="logs">
		
		</div>
	</div>
</div>
<script src="../JS/disable buttons.js"></script>
</body>

</html>