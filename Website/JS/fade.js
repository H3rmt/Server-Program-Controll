let modal = document.getElementById("closablemodal");

async function openmodal() {
	modal.style.display = "block";
	for (let i = 0; i < 1; i += 0.05) {
		modal.style.opacity = i.toString();
		await sleep(12);
	}
	modal.style.opacity = "1";
}

async function closemodal() {
	for (let i = 1; i > 0; i -= 0.05) {
		modal.style.opacity = i.toString();
		await sleep(12);
	}
	modal.style.opacity = "0";
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
