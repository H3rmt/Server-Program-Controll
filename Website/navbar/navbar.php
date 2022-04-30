<ul id="navbar">

	<style>
		<?php
		include 'navbar.css';
		?>
	</style>

	<li>
		<a id="Overview" href="../home">
			<img class="icon" src="../Images/home.svg" alt="Homepage"/>
			<h2>Overview</h2>
		</a>
	</li>

	<?php
	include_once "../database.php";


	foreach(getprogramms() as $col) {
		?>
		<li>
			<a href="../program/program.php?id=<?= $col["ID"] ?>">
				<img class="icon" src="<?= $col["Imagesource"] ?>" alt="<?= $col["Name"] ?>"/>
				<h2><?= $col["Name"] ?></h2>
			</a>
		</li>
		<?php
	}

	?>

	<li>
		<a id="Settings" href="../settings/settings.php">
			<img class="icon" src="../Images/settings.svg" alt="Settings"/>
			<h2 style="padding: 0.5em;"><?= testAdminCookie() ? "Authorised" : "Unauthorised"; ?></h2>
		</a>
	</li>

</ul>