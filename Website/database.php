<?php

$db = new PDO('mysql:host=172.17.0.1; port=3308; dbname=Programs', 'Website', '/6uM8qlYUm*NFCef');


function getprogramms(): array {
	global $db;
	$prep = $db->prepare("SELECT ID, Name, Description, Imagesource, StatechangeTime FROM programs");
	$prep->execute();
	return $prep->fetchAll(PDO::FETCH_ASSOC);
}

function getprogramm(int $id) {
	global $db;
	$prep = $db->prepare("SELECT ID, Name, Description, Imagesource, StatechangeTime FROM programs WHERE ID=:id");
	$prep->execute([':id' => $id]);
	return $prep->fetchAll(PDO::FETCH_ASSOC)[0];
}

function addtoDatabase(string $Name, string $Description, string $imgsrc): array {
	global $db;
	$key = uniqid();
	$ID = getnewID("SELECT ID FROM programs ORDER BY ID");
	$prep = $db->prepare("INSERT INTO programs (ID,Name,Description,Imagesource,APIKey) VALUES (:ID,:Name,:Desc,:Imagesource,:APIKey)");
	$prep->execute([':ID' => $ID, ':Name' => $Name, ':Desc' => $Description, ':Imagesource' => $imgsrc, ':APIKey' => $key]);
	return [$ID, $key];
}

function getnewID(string $SQL): int {
	global $db;
	$id = 0;
	$prep = $db->prepare($SQL);
	$prep->execute();
	$IdList = $prep->fetchAll(PDO::FETCH_ASSOC);
	foreach($IdList as $nid) {
		if($nid['ID'] == $id)
			$id++;
		else
			return $id;
	}
	return $id;
}


function getSetting(string $name) {
	global $db;
	$prep = $db->prepare("SELECT Value FROM settings WHERE Name=:name");
	$prep->execute([':name' => $name]);
	return $prep->fetchAll(PDO::FETCH_ASSOC)[0]['Value'];
}

function updateSetting(string $name, $value) {
	global $db;
	$prep = $db->prepare("UPDATE settings SET Value=:value WHERE Name=:name");
	$prep->execute([':value' => $value, ':name' => $name]);
}


function testAdminCookie(): bool {
	if(isset($_COOKIE['authorisation'])) {
		return getSetting('adminCookie') === $_COOKIE['authorisation'];
	} else {
		return false;
	}
}


function getPepper(): string {
	return "uwu";
}