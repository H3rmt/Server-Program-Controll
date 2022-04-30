let websocket;

function send(_action, _id, extra = null) {
	websocket.send(JSON.stringify({...{action: _action, Program_id: _id, admin: getAuthorisationCookie()}, ...extra}));
	console.log("Message send: " + JSON.stringify({...{action: _action, id: _id, code: getAuthorisationCookie()}, ...extra}));
}

const Action = {
	getActivity: "Activity",
	getLogs: "Logs",
	start: "Start",
	stop: "Stop"
};

function processreceived(evt) {
	console.log("Message received: " + evt.data);
	const data = JSON.parse(evt.data);
	if (data["error"] !== undefined) {
		console.log("error:" + data["error"]);
		return
	}
	switch (data["action"]) {
		case Action.getLogs:
			recivelogs(data["data"]);
			break;
		case Action.getActivity:
			reciveactivity(data["data"]);
			break;
		
		case Action.start:
			recivestart(data["data"], data["error"]);
			break;
		
		case Action.stop:
			recivestop(data["data"], data["error"]);
			break;
		
		default:
			console.log("invalid action: " + data["action"]);
			break;
	}
}

function builtWebSocket() {
	let loading = true;
	try {
		websocket = new WebSocket("ws://" + window.location.host + ":18769/ws");
		// websocket = new WebSocket("ws://192.168.187.11:18769/ws");
		console.log("Connection built");
	} catch (err) {
		console.log("Connection invalid");
		alert("Connection invalid");
		loading = false;
		return
	}
	
	websocket.onopen = function () {
		console.log("connection opened!");
		loading = false;
	};
	
	websocket.onerror = function (error) {
		console.log("WebSocket Error: ", error);
	};
	
	websocket.onclose = function () {
		if (loading) {
			console.log("Connection couldn't get created");
		} else {
			console.log("Connection lost");
			alert("Connection lost");
		}
	};
	
	websocket.onmessage = processreceived;
}