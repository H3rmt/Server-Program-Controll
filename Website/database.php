<?php

$db = new PDO('mysql:host=localhost', 'Website', '/6uM8qlYUm*NFCef');

function getprogramms() {
	global $db;
	$prep = $db->prepare("SELECT ID, Name, Description, Imagesource, StatechangeTime FROM programs.programs");
	$prep->execute();
	return $prep->fetchAll(PDO::FETCH_ASSOC);
}

function getprogramm(int $id) {
	global $db;
	$prep = $db->prepare("SELECT * FROM programs.programs WHERE ID=:id");
	$prep->execute([
		':id' => $id
	]);
	return $prep->fetchAll(PDO::FETCH_ASSOC)[0];
}

function addtoDatabase(string $Name, string $Description, string $imgsrc): array {
	global $db;
	$key = uniqid();
	$ID = getnewID("SELECT ID FROM programs.programs ORDER BY ID ASC");
	$prep = $db->prepare("INSERT INTO programs.programs (ID,Name,Description,Imagesource,APIKey) VALUES (:ID,:Name,:Desc,:Imagesource,:APIKey)");
	$prep->execute([
		':ID' => $ID,
		':Name' => $Name,
		':Desc' => $Description,
		':Imagesource' => $imgsrc,
		':APIKey' => $key
	]);
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
	return ++$id;
}