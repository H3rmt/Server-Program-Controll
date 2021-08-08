<style>
	<?php
	include 'navbar.css';
	?>
</style>

<?php

include_once "../database.php";

foreach(getprogramms() as $col) {
	?>
	<li>
		<a href="../program/program.php?id=<?= $col["ID"] ?>"><?= $col["Name"] ?></a>
	</li>
	<?php
}

?>
<li style="position: absolute; bottom: 0; right: 0;">
	<a href="../settings/settings.php">
		<img src="../Images/settings.png" alt="" onerror="this.onerror=null; this.src='../Images/settings.png'"
		     style="width: 40px; height: 40px;"/>
	</a>
</li>