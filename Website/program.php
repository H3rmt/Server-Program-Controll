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

            </div>
        </div>
    </div>

</body>

</html>