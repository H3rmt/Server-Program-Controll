<?php

if($_SERVER['REQUEST_METHOD'] == 'POST' && array_key_exists('name', $_POST) && !empty($_POST['name']) && array_key_exists('picturesrc', $_POST) && !empty($_POST['description'])) {
	$ret = addtoDatabase(htmlspecialchars(stripslashes(trim($_POST['name']))), htmlspecialchars(stripslashes(trim($_POST['description']))), htmlspecialchars(stripslashes(trim($_POST['picturesrc']))));
	?>
	<div class="modal" style="display: block">
		<form class="Popup" action="homepage.php" method="POST" autocomplete="off">
			<h2><?= htmlspecialchars(stripslashes(trim($_POST['name']))) ?> created</h2>
			<h2>ID: <?= $ret[0] ?></h2>
			<h2>APIKey: <?= $ret[0] ?></h2>
			<button class="Close"><b>Close</b></button>
		</form>
	</div>
	<?php
} else {
	?>
	<div class="modal" id="closablemodal" style="display: none">
		<form class="Popup" action="homepage.php" method="POST" autocomplete="off">
			<h2>Create a new Program</h2>
			<rect class="close" onclick="closenewprogramm()">&times;</rect>
			<table>
				<tr>
					<td>
						<label for="name">Program name:</label>
					</td>
					<td>
						<input id="name" type="text" name="name" value="" autocomplete="off">
					</td>
				</tr>
				<tr style="height:150%">
					<td>
						<label for="description">Program description:</label>
					</td>
					<td>
						<input id="description" type="text" name="description" value="" autocomplete="off">
					</td>
				</tr>
				<tr>
					<td>
						<label for="picturesrc">Picture Source:</label>
					</td>
					<td>
						<input id="picturesrc" type="text" name="picturesrc" value="" autocomplete="off">
					</td>
				</tr>
			</table>
			<button class="add"><b>Add</b></button>
		</form>
	</div>
	<?php
}