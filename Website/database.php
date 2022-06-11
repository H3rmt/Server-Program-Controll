<?php

$authDB = new PDO('mysql:host=172.17.0.1; port=3308; dbname=Auth', 'Website', '/6uM8qlYUm*NFCef');

$db = new PDO('mysql:host=172.17.0.1; port=3308; dbname=Programs', 'Website', '/6uM8qlYUm*NFCef');


function getMember(string $username): array|bool {
	global $authDB;
	$prep = $authDB->prepare("SELECT passwd, ID, admin FROM users WHERE name=:username");
	$prep->execute([
			':username' => $username
	]);
	$ret = $prep->fetchAll(PDO::FETCH_ASSOC);
	return $ret ? $ret[0] : false;
}

function getSession(string $username, string $hash): array|bool {
	global $authDB;
	$prep = $authDB->prepare("SELECT ID, expire_date, hash, user_id FROM sessions WHERE user_id=(SELECT ID FROM users WHERE name=:username) AND hash=:hash");
	$prep->execute([
			':username' => $username,
			':hash' => $hash
	]);
	$ret = $prep->fetchAll(PDO::FETCH_ASSOC);
	return $ret ? $ret[0] : false;
}

function dropSession(int $id): void {
	global $authDB;
	$prep = $authDB->prepare("DELETE FROM sessions WHERE ID=:id");
	$prep->execute([
			':id' => $id,
	]);
}

function createSession(int $id, string $hash, string $expiry_date): void {
	global $authDB;
	$prep = $authDB->prepare("INSERT INTO sessions (user_id, hash, expire_date) VALUES (:id, :hash, :expire_date)");
	$prep->execute([
			':id' => $id,
			':hash' => $hash,
			':expire_date' => $expiry_date
	]);
}

function getSetting(string $name) {
	global $authDB;
	$prep = $authDB->prepare("SELECT Value FROM settings WHERE Name=:name");
	$prep->execute([':name' => $name]);
	return $prep->fetchAll(PDO::FETCH_ASSOC)[0]['Value'];
}

function updateSetting(string $name, $value): void {
	global $authDB;
	$prep = $authDB->prepare("UPDATE settings SET Value=:value WHERE Name=:name");
	$prep->execute([':value' => $value, ':name' => $name]);
}

function logout(string $username, string $hash): void {
	global $authDB;
	$prep = $authDB->prepare("DELETE FROM sessions WHERE user_id=(SELECT ID FROM users WHERE name=:username) AND hash=:hash");
	$prep->execute([
			':username' => $username,
			':hash' => $hash
	]);
}

function clearSessions(string $username): void {
	global $authDB;
	$prep = $authDB->prepare("DELETE FROM sessions WHERE user_id=(SELECT ID FROM users WHERE name=:username)");
	$prep->execute([
			':username' => $username,
	]);
}

function getProgramsForUser(int $id): array {
	global $authDB, $db;
	$prep = $authDB->prepare("SELECT ID, admin FROM users WHERE ID = :id");
	$prep->execute([':id' => $id]);
	if($prep->fetchAll(PDO::FETCH_ASSOC)[0]['admin']) {
		$prep = $db->prepare("SELECT ID AS program_id, 'all' as permission FROM programs");
		$prep->execute();
	} else {
		$prep = $authDB->prepare("SELECT program_id, permission FROM user_programs_permissions WHERE user_id = :id");
		$prep->execute([':id' => $id]);
	}
	
	return $prep->fetchAll(PDO::FETCH_ASSOC);
}

function getProgramms(int $user_id): array {
	global $db;
	$prs = getProgramsForUser($user_id);
	$programs = [];
	foreach($prs as $pr) {
		$prep = $db->prepare('SELECT ID, Name, Description, Imagesource, StatechangeTime FROM programs WHERE ID = :id');
		$prep->execute([":id" => $pr['program_id']]);
		$programs[] = ['program' => $prep->fetchAll(PDO::FETCH_ASSOC)[0], 'permission' => $pr['permission']];
	}
	return $programs;
}

function getProgramm(int $id) {
	global $db;
	$prep = $db->prepare("SELECT ID, Name, Description, Imagesource, StatechangeTime FROM programs WHERE ID=:id");
	$prep->execute([':id' => $id]);
	return $prep->fetchAll(PDO::FETCH_ASSOC)[0];
}

function addToDatabase(string $Name, string $Description, string $imgsrc): array {
	global $db;
	$key = uniqid();
	$prep = $db->prepare("INSERT INTO programs (Name,Description,Imagesource,APIKey) VALUES (:Name,:Desc,:Imagesource,:APIKey)");
	$prep->execute([':Name' => $Name, ':Desc' => $Description, ':Imagesource' => $imgsrc, ':APIKey' => $key]);
	return [$db->lastInsertId(), $key];
}
