<?php

include_once "../database.php";


function displayPrograms(int $id): void {
	foreach(getProgramms($id) as $col) { ?>
		<div class="flipBoxSensor">
			<div class="flipBox">
				<div class="back">
					<div class="outerBox">
						<h1 class="title"><?= htmlspecialchars($col["Name"]) ?></h1>
						<h2 class="ontime"><?= htmlspecialchars("69 days") ?></h2>
						<h2 class="startdate"><?= htmlspecialchars($col["StatechangeTime"]) ?></h2>
						<div>
							<button class="start disabled add"><b>Start</b></button>
							<button class="stop disabled danger "><b>Stop</b></button>
						</div>
						<button class="opensite" onclick="window.open('../program/program.php?id=<?= $col["ID"] ?>');">
							<b>Open website</b>
						</button>
					</div>
				</div>
				<div class="front">
					<div class="outerBox">
						<h1 class="title"><?= $col["Name"] ?></h1>
						<img src="<?= htmlspecialchars($col["Imagesource"]) ?>" class="botIcon" alt=""/>
						<p class="description"><?= htmlspecialchars($col["Description"]) ?></p>
						<div class="active">
							<rect></rect>
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