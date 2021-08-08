<?php

if($_SERVER['REQUEST_METHOD'] == 'POST' && array_key_exists('name', $_POST) && array_key_exists('picturesrc', $_POST)) {
	if(!empty($_POST['name']) && !empty($_POST['description'])) {
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
		<div class="modal" style="display: block">
			<form class="Popup" action="homepage.php" method="POST" autocomplete="off">
				<h2>Error creating Program</h2>
				<h2><?php if(empty($_POST['name'])) echo "Name missing"; else if (empty($_POST['description'])) echo "Description missing";  ?></h2>
				<button class="Close"><b>Close</b></button>
			</form>
		</div>
		<?php
	}
} else {
	?>
	<div class="modal" id="closablemodal" style="display: none">
		<form class="Popup" action="homepage.php" method="POST" autocomplete="off">
			<h1>Create a new Program</h1>
			<p class="close" onclick="closenewprogramm()">&times;</p>
			<table>
				<tr>
					<td>
						<label for="name">Program name:</label>
					</td>
					<td>
						<input id="name" type="text" name="name" value="" autocomplete="off">
					</td>
				</tr>
				<tr>
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
?>