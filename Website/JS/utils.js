let root = "Website"

let modal

function searchModal() {
	modal = document.getElementById("closable-modal");
}

async function openModal() {
	if (modal.style.display === "block")
		return;
	modal.style.display = "block";
	for (let i = 0.0; i <= 1; i += 0.05) {
		modal.style.opacity = i.toString();
		await sleep(15);
	}
	modal.style.opacity = "1";
}

async function closeModal() {
	if (modal.style.display === "none")
		return;
	for (let i = 1.0; i >= 0; i -= 0.05) {
		modal.style.opacity = i.toString();
		await sleep(15);
	}
	modal.style.opacity = "0";
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
	window.location.replace(`${window.location.origin}/${root}?message=logout`);
}

function logoutAllSessions() {
	window.location.replace(`${window.location.origin}/${root}?message=clearsessions`);
}

function getCookie(name) {
	for (let cookie of document.cookie.split(';')) {
		if (cookie.trim().split('=')[0] === name) {
			return cookie.trim().split('=')[1]
		}
	}
}