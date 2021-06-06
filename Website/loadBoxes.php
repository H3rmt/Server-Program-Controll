<?php

$build = array('
<div class="flipboxsensor" onmouseenter="flipback(\'{_ID_}\')" onmouseleave="flipfront(\'{_ID_}\')">
    <div class="flipbox" id="{_ID_}">
        <div class="back">
            <div class="outerbox">
                <h2>{_Name_}</h2>
                <p class="ontime">{-Ontime-}</p>
                <p class="startdate">{_Starttime_}</p>
                <div>
                    <button class="start disabled"><b>Start</b></button>
                    <button class="stop disabled"><b>Stop</b></button>
                </div>
                <button class="opensite" onclick="window.location.href=\'bot.php?id={_ID_}\';">
                    <b>Open website</b>
                </button>
            </div>
        </div>
        <div class="front">
            <div class="outerbox">
                <h2>{_Name_}</h2>
                <img src={_Imagesource_} class="boticon" onerror="this.onerror=null; this.src=\'Images/imgnotfound.png\'" alt=""/>
                <p class="description">{_Description_}</p>
                <div class="active">
                    <rect> </rect>
                    <p>{-Active-}</p>
                </div>
            </div>
        </div>
    </div>
</div>');

include_once "database.php"; 

foreach(getprogramms() as $col) {
    $find = array("{_ID_}", "{_Name_}", "{_Starttime_}", "{_Imagesource_}", "{_Description_}");
    $repl = array($col["ID"], $col["Name"], $col["Starttime"], $col["Imagesource"], $col["Description"],);
    echo str_replace($find, $repl, $build)[0];
}