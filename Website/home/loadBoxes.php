<?php

include_once "../database.php";

foreach(getprogramms() as $col) {
	?>
	<div class="flipboxsensor">
		<div class="flipbox">
			<div class="back">
				<div class="outerbox">
					<h1><?= $col["Name"] ?></h1>
					<p class="ontime"><?= "ontime" ?></p>
					<p class="startdate"><?= $col["StatechangeTime"] ?></p>
					<div>
						<button class="start disabled"><b>Start</b></button>
						<button class="stop disabled"><b>Stop</b></button>
					</div>
					<button class="opensite" onclick="window.open('../program/program.php?id=<?= $col["ID"] ?>');">
						<b>Open website</b>
					</button>
				</div>
			</div>
			<div class="front">
				<div class="outerbox">
					<h1><?= $col["Name"] ?></h1>
					<img src="<?= $col["Imagesource"] ?>" class="boticon" alt=""/>
					<p class="description"><?= $col["Description"] ?></p>
					<div class="active">
						<rect></rect>
						<p><?= "active" ?></p>
					</div>
				</div>
			</div>
		</div>
	</div>
	<?php
}