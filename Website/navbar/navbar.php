<ul id="navbar">
	
	<style>
		<?php
		include 'navbar.css';
		?>
	</style>
	
	<li>
		<a id="Overview" href="../home/homepage.php"><h2>Overview</h2></a>
	</li>
	
	<?php
	
	include_once "../database.php";
	
	foreach(getprogramms() as $col) {
		?>
		<li>
			<a href="../program/program.php?id=<?= $col["ID"] ?>"><h2><?= $col["Name"] ?></h2></a>
		</li>
		<?php
	}
	
	?>
	
	<li style="position: absolute; bottom: 0; right: 0;">
		<a href="../settings/settings.php" style="align-self: end">
			<img src="../Images/settings.png" alt="" style="width: 40px; height: 40px;"/>
		</a>
	</li>

</ul>