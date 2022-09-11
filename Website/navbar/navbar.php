<?php

include_once "../database.php";

function displayNavbar(int $id): void { ?>
	<div id="navbar">
		<a id="Overview" href="../home/home.php">
			<img class="icon" src="../Images/home.svg" alt="Homepage"/>
			<h2>Overview</h2>
		</a>
		<div id="navbar_list">
			<?php
			foreach(getProgramms($id) as $col) {
				?>
				<a href="../program/program.php?id=<?= htmlspecialchars($col['program']["ID"]); ?>">
					<img class="icon" src="<?= htmlspecialchars($col['program']["Imagesource"]); ?>" alt="<?= htmlspecialchars($col['program']["Name"]); ?>"/>
					<h2><?= htmlspecialchars($col['program']["Name"]); ?></h2>
				</a>
				<?php
			} ?>
		</div>
		<a id="Settings" href="../settings/settings.php">
			<img class="icon" src="../Images/settings.svg" alt="Settings"/>
			<h2>Settings</h2>
		</a>
	</div>
	<?php
}