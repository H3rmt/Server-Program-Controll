function recivelogs(data) {
	data.forEach((log) => {
		console.log("log: ", log);
	});
}

function reciveactivity(data) {
	data.forEach((activity) => {
		console.log("activity: ", activity);
	});
}

function recivestart(data, error) {
	if(error != null) {
		if(error === "no admin permissions") {
			console.log("no admin permissions");
		} else {
			console.log("Error");
		}
		return;
	}
	if(data === true) {
		console.log("success", data);
	} else {
		console.log("not successfull", data);
	}
}

function recivestop(data, error) {
	if(error != null) {
		if(error === "no admin permissions") {
			console.log("no admin permissions");
		} else {
			console.log("Error");
		}
		return;
	}
	if(data === true) {
		console.log("success", data);
	} else {
		console.log("not successfull", data);
	}
}