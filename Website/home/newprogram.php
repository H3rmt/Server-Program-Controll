<?php

if($_SERVER['REQUEST_METHOD'] == 'POST' && array_key_exists('name', $_POST) && array_key_exists('picturesrc', $_POST)) {
	if(testAdminCookie()) {
		if(!empty($_POST['name']) && !empty($_POST['description'])) {
			$ret = addToDatabase(htmlspecialchars(stripslashes(trim($_POST['name']))), htmlspecialchars(stripslashes(trim($_POST['description']))), htmlspecialchars(stripslashes(trim($_POST['picturesrc']))));
			?>
			<div class="modal" style="display: block">
				<form class="Popup" action="homepage.php" method="POST" autocomplete="off">
					<h1><?= htmlspecialchars(stripslashes(trim($_POST['name']))) ?> created</h1>
					<h2>ID: <?= $ret[0] ?></h2>
					<h2>APIKey: <?= $ret[1] ?></h2>
					<button class="Close" type="submit"><b>Close</b></button>
				</form>
			</div>
			<?php
		} else {
			?>
			<div class="modal" style="display: block">
				<form class="Popup" action="homepage.php" method="POST" autocomplete="off">
					<h1>Error creating Program</h1>
					<h2><?php
						if(empty($_POST['name'])) echo "Name missing"; else if(empty($_POST['description'])) echo "Description missing"; ?></h2>
					<button class="Close" type="submit"><b>Close</b></button>
				</form>
			</div>
			<?php
		}
	} else {
		?>
		<div class="modal" style="display: block">
			<form class="Popup" action="homepage.php" method="POST" autocomplete="off">
				<h1>Error creating Program</h1>
				<h2>Missing Admin Permissions</h2>
				<button class="Close" type="submit"><b>Close</b></button>
			</form>
		</div>
		<?php
	}
} else {
	?>
	<div class="modal" id="closable-modal" style="display: none">
		<form class="Popup" action="homepage.php" method="POST" autocomplete="off">
			<h1>Create a new Program</h1>
			<p class="close" onclick="closeModal()">&times;</p>
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
			<button class="add" type="submit"><b>Add</b></button>
		</form>
	</div>
	<?php
}
?>
