let modal =  document.getElementById("closablemodal");

function searchmodal() {
	modal =  document.getElementById("closablemodal");
}

async function openmodal() {
	if (modal.style.display == "block")
		return;
	modal.style.display = "block";
	for (let i = 0; i < 1; i += 0.05) {
		modal.style.opacity = i.toString();
		await sleep(12);
	}
}

async function closemodal() {
	if (modal.style.display == "none")
		return;
	for (let i = 1; i > 0; i -= 0.05) {
		modal.style.opacity = i.toString();
		await sleep(12);
	}
	modal.style.display = "none";
}

window.onclick = async function (event) {
	if (event.target === modal) {
		await closemodal();
	}
};

async function sleep(ms) {
	return new Promise((resolve) => setTimeout(resolve, ms));
}

function disable() {
	Array.from(document.body.getElementsByTagName('*')).forEach((element) => {
		element.disabled = element.classList.contains('disabled')
	})
}

function protect() {
	Array.from(document.body.getElementsByTagName('*')).forEach((element) => {
		if (element.classList.contains("protected"))
			element.classList.remove('disabled')
	})
}

function replaceImages() {
	Array.from(document.body.getElementsByTagName('img')).forEach((img) => {
		img.onerror = function () {
			img.src = "../Images/imgnotfound.png";
		}
	})
}

function getAuthorisationCookie() {
	let name = "authorisation=";
	let decodedCookie = decodeURIComponent(document.cookie);
	let ca = decodedCookie.split(';');
	for (let i = 0; i < ca.length; i++) {
		let c = ca[i];
		while (c.charAt(0) === ' ') {
			c = c.substring(1);
		}
		if (c.indexOf(name) === 0) {
			return c.substring(name.length, c.length);
		}
	}
	return "";
}


// document.cookie = "authorisation=; expires=Thu, 01 Jan 1970 00:00:00 UTC; path=/;";