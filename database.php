<?php

$db = new PDO('mysql:host=172.17.0.1; port=3308; dbname=Auth', 'Website', '/6uM8qlYUm*NFCef');


function getMember(string $username): array|bool {
	global $db;
	$prep = $db->prepare("SELECT * FROM users WHERE name=:username");
	$prep->execute([
			':username' => $username
	]);
	$ret = $prep->fetchAll(PDO::FETCH_ASSOC);
	return $ret ? $ret[0] : false;
}

function getSession(string $username, string $hash): array|bool {
	global $db;
	$prep = $db->prepare("SELECT ID, username, expire_date, hash FROM sessions WHERE username=:username AND hash=:hash");
	$prep->execute([
			':username' => $username,
			':hash' => $hash
	]);
	$ret = $prep->fetchAll(PDO::FETCH_ASSOC);
	return $ret ? $ret[0] : false;
}

function dropSession(int $id): void {
	global $db;
	$prep = $db->prepare("DELETE FROM sessions WHERE ID=:id");
	$prep->execute([
			':id' => $id,
	]);
}

function createSession($username, $hash, $expiry_date): void {
	global $db;
	$prep = $db->prepare("INSERT INTO sessions (username, hash, expire_date) VALUES (:username, :hash, :expire_date)");
	$prep->execute([
			':username' => $username,
			':hash' => $hash,
			':expire_date' => $expiry_date
	]);
}