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
echo '<li style="position: absolute; bottom: 0; right: 0;"><a href="settings.php"><img src="Images/settings.png"  onerror="this.onerror=null; this.src=\'Images/imgnotfound.png\'" alt="" style="width: 40px; height: 40px;"></img></a></li>';