<?php
$default = '
    <div class="modal" id="closablemodal" style="display: {_Display_}">
        <form class="Popup" action="index.php" method="POST" autocomplete="off">
            <h2>Create a new Program</h2>
            <rect class="close" onclick="closenewprogramm()">&times;</rect>

            <input type="hidden" name="action" value="create">
            
            <div class="parallel">
                <label for="name">Program name:</label>
                <input id="name" type="text" name="name" value="{_name_}" autocomplete="off">
            </div>
            <div class="parallel">
                <label for="description">Program description:</label>
                <input id="description" type="text" name="description" value="{_description_}" autocomplete="off">
            </div>
            <div class="parallel">
                <label for="picturesrc">Picture Source:</label>
                <input id="picturesrc" type="text" name="picturesrc" value="{_picturesrc_}" autocomplete="off">
            </div>
            {_Error_}
            <button class="add"><b>Add</b></button>
        </form>
    </div>';
$created = '
    <div class="modal"  style="display: block">
        <div class="Popup">
            <h2>{_name_} created</h2>
            <h2>ID: {_ID_}</h2>
            <button class="Close" onclick="window.location.href = \'index.php\';"><b>Close</b></button>  
        </div>
    </div>';
if($_SERVER['REQUEST_METHOD'] == 'POST' && array_key_exists('name', $_POST) && array_key_exists('picturesrc', $_POST)) {
    if(empty(htmlspecialchars(stripslashes(trim($_POST['name']))))) { # missing name
        echo str_replace(['{_Error_}', '{_Display_}', '{_name_}', '{_description_}', '{_picturesrc_}'], ['<p class="error">Enter valid name</p>', 'block', '', htmlspecialchars(stripslashes(trim($_POST['description']))), htmlspecialchars(stripslashes(trim($_POST['picturesrc'])))], $default);
    } else if(empty(htmlspecialchars(stripslashes(trim($_POST['description']))))) { # missing description
        echo str_replace(['{_Error_}', '{_Display_}', '{_description_}', '{_name_}', '{_picturesrc_}'], ['<p class="error">Enter valid description</p>', 'block', '', htmlspecialchars(stripslashes(trim($_POST['name']))), htmlspecialchars(stripslashes(trim($_POST['picturesrc'])))], $default);
    } else if(empty(htmlspecialchars(stripslashes(trim($_POST['picturesrc']))))) { # missing src for pic
        echo str_replace(['{_Error_}', '{_Display_}', '{_picturesrc_}', '{_name_}', '{_description_}'], ['<p class="error">Enter valid picturesource or {dynamic}</p>', 'block', '', htmlspecialchars(stripslashes(trim($_POST['name']))), htmlspecialchars(stripslashes(trim($_POST['description'])))], $default);
    } else {
        $id = addtodatabase(htmlspecialchars(stripslashes(trim($_POST['name']))), htmlspecialchars(stripslashes(trim($_POST['description']))), htmlspecialchars(stripslashes(trim($_POST['picturesrc']))));
        echo str_replace(['{_name_}', '{_ID_}'], [htmlspecialchars(stripslashes(trim($_POST['name']))), $id], $created); # successfully created new Program
    }
} else {  # default page
    echo str_replace(['{_Error_}', '{_Display_}', '{_name_}', '{_description_}', '{_picturesrc_}'], ['', 'none', '', '', ''], $default);
}