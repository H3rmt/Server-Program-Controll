<?php

include_once "../database.php";

function createcookies() {
	if($_SERVER['REQUEST_METHOD'] == 'POST' && array_key_exists('hashed_password', $_POST) && !empty($_POST['hashed_password'])) {
		$ret = testpassword(htmlspecialchars(stripslashes(trim($_POST['hashed_password']))));

		if($ret) {
			$cookie_val = (int)(rand(69696969, 6969696969) / 420 * 5.0);
			setcookie('authorisation', $cookie_val, time() + (86400 / 2), "/");
			updateSetting('admincookie', $cookie_val);
		}
	}
}

/*
// refresh page on load to prevent resending of post
header("Refresh:0");
 */

function createmodal() {
	$isAdmin = testadmincookie();

	// check if password is sent
	if($_SERVER['REQUEST_METHOD'] == 'POST' && array_key_exists('hashed_password', $_POST) && !empty($_POST['hashed_password'])) {
		$valid = testpassword(htmlspecialchars(stripslashes(trim($_POST['hashed_password']))));

		if($valid) {
			?>
			<div class="modal" style="display: block">
				<form class="Popup" id="authorise">
					<h1>Authorise</h1>
					<p class="close" onclick="document.getElementById('authorise').submit()">&times;</p>
					<h3>Authorisation succeeded</h3>
					<button class="Close"><b>Close</b></button>
				</form>
			</div>
			<?php
		} else {
			?>
			<div class="modal" style="display: block">
				<form class="Popup" id="authorise">
					<h1>Authorise</h1>
					<p class="close" onclick="document.getElementById('authorise').submit()">&times;</p>
					<h3>Authorisation failed</h3>
					<button class="Close"><b>Close</b></button>
				</form>
			</div>
			<?php
		}
	}

	if($isAdmin) {
		?>
		<div class="modal" id="closable-modal" style="display: none">
			<form class="Popup" id="authorise">
				<h1>Authorise</h1>
				<p class="close" onclick="closemodal()">&times;</p>
				<h3>Authorised</h3>
				<button class="Reset danger" onclick="removeAuthorisationCookie();document.getElementById('authorise').submit()"><b>Reset</b></button>
				<button class="Close"><b>Close</b></button>
			</form>
		</div>
		<?php
	} else {
		?>
		<div class="modal" id="closable-modal" style="display: none">
			<form class="Popup" id="authorise" action="settings.php" method="POST">
				<h1>Authorise</h1>
				<p class="close" onclick="closemodal()">&times;</p>
				<table>
					<tr>
						<td>
							<label for="password">Password:</label>
						</td>
						<td>
							<input id="password" aria-label="password" type="password" name="password">
						</td>
					</tr>
				</table>
				<button type="submit" class="add"><b>Check</b></button>

				<input style="display: none" id="hashed_password" type="password" name="hashed_password">
				<script>
					document.getElementById("authorise").addEventListener("submit", () => {
						document.getElementById("hashed_password").value = SHA256(document.getElementById("password").value)
						document.getElementById("password").value = "*".repeat(document.getElementById("password").value.length)
					})
				</script>
			</form>
		</div>
		<?php
	}
}


function testpassword($hash): bool {
	$setting = getSetting('password');
	return $setting == $hash;
}

?>