<?php

$db = new PDO('mysql:host=172.17.0.1', 'Website', '/6uM8qlYUm*NFCef');
//$db = new PDO('mysql:host=localhost', 'Website', '/6uM8qlYUm*NFCef');


function getprogramms(): array {
	global $db;
	$prep = $db->prepare("SELECT ID, Name, Description, Imagesource, StatechangeTime FROM programs.programs");
	$prep->execute();
	return $prep->fetchAll(PDO::FETCH_ASSOC);
}

function getprogramm(int $id) {
	global $db;
	$prep = $db->prepare("SELECT ID, Name, Description, Imagesource, StatechangeTime FROM programs.programs WHERE ID=:id");
	$prep->execute([':id' => $id]);
	return $prep->fetchAll(PDO::FETCH_ASSOC)[0];
}

function addtoDatabase(string $Name, string $Description, string $imgsrc): array {
	global $db;
	$key = uniqid();
	$ID = getnewID("SELECT ID FROM programs.programs ORDER BY ID");
	$prep = $db->prepare("INSERT INTO programs.programs (ID,Name,Description,Imagesource,APIKey) VALUES (:ID,:Name,:Desc,:Imagesource,:APIKey)");
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
	$prep = $db->prepare("SELECT Value FROM programs.settings WHERE Name=:name");
	$prep->execute([':name' => $name]);
	return $prep->fetchAll(PDO::FETCH_ASSOC)[0]['Value'];
}

function updateSetting(string $name, $value) {
	global $db;
	$prep = $db->prepare("UPDATE programs.settings SET Value=:value WHERE Name=:name");
	$prep->execute([':value' => $value, ':name' => $name]);
}


function testadmincookie(): bool {
	if(isset($_COOKIE['authorisation'])) {
		return getSetting('admincookie') === $_COOKIE['authorisation'];
	} else {
		return false;
	}
}
