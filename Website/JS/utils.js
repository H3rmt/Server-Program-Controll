let root = "Server-Program-Controll/Website"

let modal

function searchModal() {
	modal = document.getElementById("closable-modal");
}

async function openModal() {
	if (modal.style.display === "block")
		return;
	modal.style.display = "block";
	for (let i = 0; i < 1; i += 0.05) {
		modal.style.opacity = i.toString();
		await sleep(12);
	}
}

async function closeModal() {
	if (modal.style.display === "none")
		return;
	for (let i = 1; i > 0; i -= 0.05) {
		modal.style.opacity = i.toString();
		await sleep(12);
	}
	modal.style.display = "none";
}

window.onclick = async function (event) {
	if (event.target === modal) {
		await closeModal();
	}
};

async function sleep(ms) {
	return new Promise((resolve) => setTimeout(resolve, ms));
}

function replaceImages() {
	Array.from(document.body.getElementsByTagName("img")).forEach((img) => {
		let src = img.getAttribute("src");
		if (src === null || src.length === 0) img.src = "../Images/imgnotfound.png";
		else
			fetch(src).then((res) => {
				if (res.status >= 200 && res.status <= 299) {
					img.src = src;
				} else {
					img.src = "../Images/imgnotfound.png";
				}
			});
	});
}

function logout() {
	eraseCookie("username")
	eraseCookie("hash")
	window.location.replace(`${window.location.origin}/${root}`);
}

function eraseCookie(name) {
	document.cookie = `${name}=; path=/${root}; Max-Age=-99999999;`;
}

// function getCookie(name) {
// 	return document.cookie.split(';').some(c => {
// 		return c.trim().startsWith(name + '=');
// 	});
// }

// function setCookie(name, value, days) {
// 	let date = new Date();
// 	date.setTime(date.getTime() + (days * 24 * 60 * 60 * 1000));
// 	const expires = "expires=" + date.toUTCString();
// 	document.cookie = name + "=" + value + "; " + expires + "; path=/;";
// }