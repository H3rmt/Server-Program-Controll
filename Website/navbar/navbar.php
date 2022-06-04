<?php

include_once "../database.php";

function displayNavbar(int $id): void { ?>
	<ul id="navbar">
		<style>
			<?php
			include 'navbar.css';
			?>
		</style>
		<li>
			<a id="Overview" href="../home/home.php">
				<img class="icon" src="../Images/home.svg" alt="Homepage"/>
				<h2>Overview</h2>
			</a>
		</li>
		<?php
		foreach(getProgramms($id) as $col) {
			?>
			<li>
				<a href="../program/program.php?id=<?= $col['program']["ID"] ?>">
					<img class="icon" src="<?= $col['program']["Imagesource"] ?>" alt="<?= $col['program']["Name"] ?>"/>
					<h2><?= $col['program']["Name"] ?></h2>
				</a>
			</li>
			<?php
		} ?>
		<li>
			<a id="Settings" href="../settings/settings.php">
				<img class="icon" src="../Images/settings.svg" alt="Settings"/>
				<h2 style="padding: 0.5em;">Settings</h2>
			</a>
		</li>
	</ul>
	<?php
}