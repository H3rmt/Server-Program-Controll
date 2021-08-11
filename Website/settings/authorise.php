<?php

include_once "../database.php";

function createcookies() {
	if($_SERVER['REQUEST_METHOD'] == 'POST' && array_key_exists('password', $_POST) && !empty($_POST['password'])) {
		$ret = testpassword(htmlspecialchars(stripslashes(trim($_POST['password']))));
		
		if($ret) {
			$cookie_val = time() / 5645;
			setcookie('authorisation', $cookie_val, time() + (86400 / 2), "/");
			updateSetting('admincookie', $cookie_val);
		}
	}
}

function createmodal() {
	$cookie = testadmincookie();
	if(($_SERVER['REQUEST_METHOD'] == 'POST' && array_key_exists('password', $_POST) && !empty($_POST['password'])) || $cookie) {
		if(!$cookie)
			$ret = testpassword(htmlspecialchars(stripslashes(trim($_POST['password']))));
		else
			$ret = true;
		
		if($ret) {
			?>
			<div class="modal" id="closablemodal" style="display: <?= $cookie ? "none" : "block" ?>">
				<form class="Popup" id="authorise" action="settings.php" method="POST" autocomplete="on">
					<h1>Authorise</h1>
					<p class="close" onclick="closemodal()">&times;</p>
					<h3>Authorised</h3>
					<button class="Close" onclick="closemodal()"><b>Close</b></button>
					<script>
						document.getElementById("authorise").addEventListener("submit", (event) => {
							event.preventDefault()
						})
					</script>
				</form>
			</div>
			<?php
		} else {
			?>
			<div class="modal" style="display: block">
				<form class="Popup" id="authorise" action="settings.php" method="POST" autocomplete="on">
					<h1>Authorise</h1>
					<p class="close" onclick="closemodal()">&times;</p>
					<h3>Wrong Password</h3>
					<button class="Close"><b>Close</b></button>
				</form>
			</div>
			<?php
		}
	} else {
		?>
		<div class="modal" id="closablemodal" style="display: none">
			<form class="Popup" id="authorise" action="settings.php" method="POST" autocomplete="on">
				<h1>Authorise</h1>
				<p class="close" onclick="closemodal()">&times;</p>
				<table>
					<tr>
						<td>
							<label for="password">Password:</label>
						</td>
						<td>
							<input id="password" type="password" name="password" value="" autocomplete="on">
						</td>
					</tr>
				</table>
				<button type="submit" class="add"><b>Check</b></button>
				<script>
					document.getElementById("authorise").addEventListener("submit", (event) => {
						event.target[0].value = SHA256(event.target[0].value)
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

function testadmincookie(): bool {
	if(isset($_COOKIE['authorisation'])) {
		$setting = getSetting('admincookie');
		return $setting == $_COOKIE['authorisation'];
	} else {
		return false;
	}
}

?>