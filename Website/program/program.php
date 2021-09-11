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
	<!-- <script src="../JS/Websocket.js"></script> -->
	<script src="../JS/utils.js"></script>
	
	<script src="https://cdn.jsdelivr.net/npm/apexcharts"></script>
	
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
		<h1 class="title"><?= $program["Name"]; ?></h1>
		<div>
			<button class="start disabled protected"><b>Start</b></button>
			<button class="stop disabled protected"><b>Stop</b></button>
			<button class="delete disabled protected"><b>Delete</b></button>
		</div>
	</div>
	<div id="boxes">
		<div id="topbar">
			<img src="<?= $program["Imagesource"] ?>" class="boticon"
			     alt=""/>
			<h2 class="description"><?= $program["Description"]; ?></h2>
		</div>
		<div id="activity">
			<div id="activity-chart" style="padding-top: 0.7em;"></div>
			<script>
				const color = '#2596e7'

				const options = {
					chart: {
						type: 'line',
						width: '100%',
						foreColor: color,
						animations: {
							enabled: true,
							easing: 'linear',
							speed: 500
						},
						toolbar: {
							show: true,
							offsetX: -15,
							offsetY: 0,
							tools: {
								download: false,
								selection: false,
								pan: false,

								zoom: true,
								zoomin: true,
								zoomout: true,
								reset: true
							},
							autoSelected: 'zoom'
						}
					},
					tooltip: {
						enabled: true,
						followCursor: true,
						intersect: false,
						fillSeriesColor: false,
						theme: "dark",
						style: {
							fontSize: '1.3em'
						},
						onDatasetHover: {
							highlightDataSeries: false
						}
					},

					stroke: {
						width: 4,
						curve: 'smooth',
						colors: color
					},
					series: [{
						data: [{
							x: new Date('2018-02-12').getTime(),
							y: 76
						}, {
							x: new Date('2018-02-11').getTime(),
							y: 90
						}, {
							x: new Date('2018-02-10').getTime(),
							y: 40
						}, {
							x: new Date('2018-02-09').getTime(),
							y: 20
						}, {
							x: new Date('2018-02-08').getTime(),
							y: 60
						}]
					}],
					xaxis: {
						type: 'datetime'
					}
				};
				const chart = new ApexCharts(document.getElementById("activity-chart"), options);
			</script>
		</div>
		<div id="logs">
			<div id="log-chart" style="padding-top: 0.7em;"></div>
			<script>
				const chart2 = new ApexCharts(document.getElementById("log-chart"), options);
			</script>
		</div>
	</div>
</div>
<script>
	searchmodal()

	if (getAuthorisationCookie() !== "") {
		protect();
	}

	disable();
	replaceImages();
</script>
</body>
<script>
	chart.render();
	chart2.render();
</script>
</script>
</html>