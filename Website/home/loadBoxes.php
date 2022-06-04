<?php

include_once "../database.php";


function displayPrograms(int $id, bool $isAdmin): void {
	foreach(getProgramms($id) as $col) { ?>
		<div class="flipBoxSensor">
			<div class="flipBox">
				<div class="back">
					<div class="outerBox">
						<h1 class="title"><?= htmlspecialchars($col['program']["Name"]) ?></h1>
						<h2 class="ontime"><?= htmlspecialchars("69 days") ?></h2>
						<h2 class="startdate"><?= htmlspecialchars($col['program']["StatechangeTime"]) ?></h2>
						<div>
							<button class="start <?= $isAdmin || $col['permission'] > 0 ? '' : 'disabled' ?> add"><b>Start</b></button>
							<button class="stop <?= $isAdmin || $col['permission'] > 1 ? '' : 'disabled' ?> danger "><b>Stop</b></button>
						</div>
						<button class="opensite"
								  onclick="window.open('../program/program.php?id=<?= htmlspecialchars($col['program']["ID"]) ?>');">
							<b>Open website</b>
						</button>
					</div>
				</div>
				<div class="front">
					<div class="outerBox">
						<h1 class="title"><?= htmlspecialchars($col['program']["Name"]) ?></h1>
						<img src="<?= htmlspecialchars($col['program']["Imagesource"]) ?>" class="PIcon" alt="Icon "/>
						<p class="description"><?= htmlspecialchars($col['program']["Description"]) ?></p>
						<div class="active">
							<h2><?= htmlspecialchars("active") ?></h2>
						</div>
					</div>
				</div>
			</div>
		</div>
		<?php
	}
}

?>