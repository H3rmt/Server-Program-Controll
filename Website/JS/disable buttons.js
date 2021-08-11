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

protect();
disable();
replaceImages();
