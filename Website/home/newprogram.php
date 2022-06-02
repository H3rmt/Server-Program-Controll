<?php

function success(string $name, int $id, string $key): void {
	?>
	<div class="modal">
		<form class="Popup" action="home.php" method="POST" autocomplete="off">
			<h1><?= htmlspecialchars($name) ?> created</h1>
			<h2>ID: <?= htmlspecialchars($id) ?></h2>
			<h2>APIKey: <?= htmlspecialchars($key) ?></h2>
			<button class="Close" type="submit"><b>Close</b></button>
		</form>
	</div>
	<?php
}

function error(string $error): void {
	?>
	<div class="modal">
		<form class="Popup" action="home.php" method="POST" autocomplete="off">
			<h1>Error creating Program</h1>
			<h2><?= htmlspecialchars($error) ?></h2>
			<button class="Close" type="submit"><b>Close</b></button>
		</form>
	</div>
	<?php
}

function displayModal($isAdmin): void {
	if($_SERVER['REQUEST_METHOD'] == 'POST' && array_key_exists('name', $_POST) && array_key_exists('picturesrc', $_POST)) {
		if($isAdmin) {
			if(!empty($_POST['name']) && !empty($_POST['description'])) {
				$ret = addToDatabase($_POST['name'], $_POST['description'], $_POST['picturesrc']);
				success($_POST['name'], $ret[0], $ret[1]);
			} else {
				if(empty($_POST['name']))
					error("Name missing");
				else if(empty($_POST['description']))
					error("Description missing");;
			}
		} else {
			error("Permissions missing");
		}
	} else {
		?>
		<div class="modal" id="closable-modal" style="display: none">
			<form class="Popup" action="home.php" method="POST" autocomplete="off">
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
}

?>
