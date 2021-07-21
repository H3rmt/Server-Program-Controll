<!DOCTYPE html>
<html lang="en">
<?php
include_once "database.php";
$info = getprogramm(htmlspecialchars(stripslashes(trim($_GET['id']))));
?>

<head>
    <meta charset="utf-8" />
    <title>
        <?= $info["Name"]; ?>
    </title>
    <link rel="stylesheet" href="mainstyle.css" />
    <link rel="stylesheet" href="program.css" />
    <script src="sha256.js"></script>
    <script src="Websocket.js"></script>
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
            if (error == "no admin permissions") {
                console.log("no admin permissions");
            } else {
                console.log("Error");
            }
            return;
        }
        if (data == true) {
            console.log("success", data);
        } else {
            console.log("not successfull", data);
        }
    }

    function recivestop(data, error) {
        if (error != null) {
            if (error == "no admin permissions") {
                console.log("no admin permissions");
            } else {
                console.log("Error");
            }
            return;
        }
        if (data == true) {
            console.log("success", data);
        } else {
            console.log("not successfull", data);
        }
    }
    </script>
</head>

<body>
    <ul id="navbar">
        <li>
            <a id="Overview" href="index.php">Overview</a>
        </li>
        <?php
    include "navbar.php";
    ?>
    </ul>

    <div id="div1">
        <div class="top">
            <h1><?= $info["Name"]; ?></h1>
            <div>
                <button class="start disabled"><b>Start</b></button>
                <button class="stop disabled"><b>Stop</b></button>
                <button class="delete disabled"><b>Delete</b></button>
            </div>
        </div>
        <div id="boxes">
            <div id="topbar">
                <img src=<?= $info["Imagesource"]; ?> class="boticon"
                    onerror="this.onerror=null; this.src='Images/imgnotfound.png'" alt="" />
                <h2 class="description"><?= $info["Description"]; ?></h2>
            </div>
            <div id="activity">

            </div>
            <div id="logs">
                <canvas id="myChart" width="100" height="100"></canvas>
                <script src="../.idea/chart.js"></script>
                <script>
                var ctx = document.getElementById('myChart').getContext('2d');

                const DATA_COUNT = 7;
                const NUMBER_CFG = {
                    count: DATA_COUNT,
                    min: -100,
                    max: 100
                };

                const labels = ["friday", "friday2", "friday3", "friday4", "friday5", "friday6", "friday7"]
                const data = {
                    labels: labels,
                    datasets: [{
                        label: 'Dataset 1',
                        data: [12, 13, 11, 12, 18, 17],
                        borderColor: 'rgb(255, 255, 132)',
                        backgroundColor: 'rgb(255, 99, 132)'
                    }]
                };

                var myChart = new Chart(ctx, {
                    type: 'line',
                    data: data,
                    options: {
                        responsive: true,
                        plugins: {
                            legend: {
                                position: 'top',
                            },
                            title: {
                                display: true,
                                text: 'Line Chart'
                            }
                        }
                    },
                });
                </script>
            </div>
        </div>
    </div>
    <script src="disable buttons.js"></script>
</body>

</html>