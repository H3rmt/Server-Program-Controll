<?php

$db = new PDO('mysql:dbname=programs;host=localhost', 'Website', '/6uM8qlYUm*NFCef');

function getprogramms() {
    global $db;
    $prep = $db->prepare("SELECT * FROM programs");
    $prep->execute();
    return $prep->fetchAll(PDO::FETCH_ASSOC);
}

function getprogramm($id) {
    global $db;
    $prep = $db->prepare("SELECT * FROM programs WHERE ID=:id");
    $prep->execute([':id'=>$id]);
    return $prep->fetchAll(PDO::FETCH_ASSOC)[0];
}

function addtodatabase($Name,$Description,$imgsrc): string {
    global $db;
    $id = uniqid();
    $prep = $db->prepare("INSERT INTO programs (ID,Name,Description,Imagesource) VALUES (:ID,:Name,:Desc,:Imgsrc);");
    $prep->execute([':ID'=>$id,':Name'=>$Name,':Desc'=>$Description,':Imgsrc'=>$imgsrc,]);
    return $id;
}