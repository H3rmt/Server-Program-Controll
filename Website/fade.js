let modal = document.getElementById("closablemodal");

async function opennewprogramm() {
    if (modal == null) return;
    if (modal.style.display === "block") {
        modal.style.opacity = "1";
        return;
    }
    modal.style.display = "block";
    modal.style.opacity = "0.15";
    for (let i = 0; i < 1; i += 0.05) {
        modal.style.opacity = i.toString();
        await sleep(12);
    }
    modal.style.opacity = "1";
}

async function closenewprogramm() {
    if (modal == null) return;
    for (let i = 1; i > 0; i -= 0.05) {
        modal.style.opacity = i.toString();
        await sleep(12);
    }
    modal.style.opacity = "0";
    modal.style.display = "none";
}

window.onclick = async function (event) {
    if (event.target === modal) {
        await closenewprogramm();
    }
};

async function sleep(ms) {
    return new Promise((resolve) => setTimeout(resolve, ms));
}
