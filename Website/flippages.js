fliplist = [];

async function flipback(id) {
    const dec = document.getElementById(id).getElementsByClassName("front")[0];
    fliplist.push(dec);
    await sleep(260);
    if (fliplist.indexOf(dec) >= 0) {
        dec.style.display = "none";
    }
    fliplist.splice(fliplist.indexOf(dec), 1);
}

async function flipfront(id) {
    const dec = document.getElementById(id).getElementsByClassName("front")[0];
    dec.style.display = "block";
    if (fliplist.indexOf(dec) >= 0) {
        fliplist.splice(fliplist.indexOf(dec), 1);
    }
}

async function sleep(ms) {
    return new Promise((resolve) => setTimeout(resolve, ms));
}
