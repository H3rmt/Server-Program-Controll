let websocket;

function send(_action, _id, extra = null) {
    if (connection == true) {
        websocket.send(JSON.stringify({ ...{ action: _action, Program_id: _id, code: getcode() }, ...extra }));
        console.log("Message send: " + JSON.stringify({ ...{ action: _action, id: _id, code: getcode() }, ...extra }));
    } else {
        console.log("connection not ready")
    }
}

let code = "test";

function getcode() {
    return SHA256(code);
}

const Action = {
    getactivity: "Activity",
    getlogs: "Logs",
    start: "Start",
    stop: "Stop"
};

function processreceived(evt) {
    console.log("Message received: " + evt.data);
    var data = JSON.parse(evt.data);
    switch (data["action"]) {
        case Action.getlogs:
            recivelogs(data["data"]);
            break;
        case Action.getactivity:
            reciveactivity(data["data"]);
            break;

        case Action.start:
            recivestart(data["data"], data["error"]);
            break;

        case Action.stop:
            recivestop(data["data"], data["error"]);
            break;

        case Action.customaction:
            recivecustomaction(data["data"], data["error"]);
            break;

        default:
            console.log("invalid action: " + data["action"]);
            break;
    }
}

let loading = false;
let connection = false;

function builtWebSocket() {
    loading = true;
    try {
        websocket = new WebSocket("ws://" + window.location.host + ":18769/ws");
        console.log("Connection built");
    } catch (err) {
        alert("Connection invalid");
        loading = false;
        return
    }
    connection = true;
    loading = false;

    websocket.onopen = function () {
        websocket.send('{"connection":"opened"}');
        console.log("connection opened!");
    };

    websocket.onerror = function (error) {
        console.log("WebSocket Error: " + error);
    };

    websocket.onclose = function () {
        console.log("Connection lost");
        alert("Connection lost");
    };

    websocket.onmessage = processreceived;
}

builtWebSocket();
