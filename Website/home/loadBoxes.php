<?php

include_once "../database.php";


function displayPrograms(int $id): void {
	foreach(getProgramms($id) as $col) { ?>
		<div class="flipboxsensor">
			<div class="flipbox">
				<div class="back">
					<div class="outerbox">
						<h1 class="title"><?= $col["Name"] ?></h1>
						<h2 class="ontime"><?= "ontime" ?></h2>
						<h2 class="startdate"><?= $col["StatechangeTime"] ?></h2>
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
					<div class="outerbox">
						<h1 class="title"><?= $col["Name"] ?></h1>
						<img src="<?= $col["Imagesource"] ?>" class="boticon" alt=""/>
						<p class="description"><?= $col["Description"] ?></p>
						<div class="active">
							<rect></rect>
							<h2><?= "active" ?></h2>
						</div>
					</div>
				</div>
			</div>
		</div>
		<?php
	}
}

?>