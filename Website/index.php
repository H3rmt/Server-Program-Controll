<!DOCTYPE html>
<html lang="en">

<head>
    <meta charset="utf-8"/>
    <title>Overview</title>
    <link rel="stylesheet" href="homepage.css"/>
    <link rel="stylesheet" href="mainstyle.css"/>
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
        <h1>Overview</h1>
        <button class="new disabled" onclick="opennewprogramm()"><b>New Programm</b></button>
    </div>
    <div id="boxes">
        <?php
        include_once "loadboxes.php";
        ?>
    </div>
</div>
<?php
include_once "newprogram.php";
?>
<script src="fade.js"></script>
<script src="disable buttons.js"></script>
</body>

</html>