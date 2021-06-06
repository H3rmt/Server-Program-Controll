<?php

$build = array('
<li>
    <a href="program.php?id={_ID_}">{_Name_}</a>
</li>'
);

include_once "database.php";
foreach (getprogramms() as $col) {
    $find = array("{_Name_}","{_ID_}");
    $repl = array($col["Name"],$col["ID"]);

    echo str_replace($find, $repl, $build)[0];
}