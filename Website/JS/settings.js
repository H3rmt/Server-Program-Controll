function rresetSettings(name, allowed, e) {
	e.stopPropagation()
	if (!allowed) {
		alert(`Not allowed to reset ${name}`)
		console.log(`Not allowed to reset ${name}`)
		e.preventDefault()
		return
	}
	let reset = confirm(`Reset ${name} Settings to default?`)
	if (!reset) {
		console.error(`Stopped reset ${name}`)
		e.preventDefault()
		return
	}
	document.getElementById(`${name}_reset`).checked = true
	// allow form submit
}

function saveSettings(name, allowed, e) {
	e.stopPropagation()
	if (!allowed) {
		alert(`Not allowed to save ${name}`)
		console.error(`Not allowed to save ${name}`)
		e.preventDefault()
		return
	}
	let reset = confirm(`Save ${name} Settings?`)
	if (!reset)
		if (name === "Server settings") {
			let passwd = document.getElementById("new_password").value
			if (passwd.length !== 0 && passwd.length < 8) {
				alert(`Not allowed to use password < 8  (${passwd.length})`)
				console.error(`Not allowed to use password < 8  (${passwd.length})`)
				e.preventDefault()
				return
			}
		}
	// allow form submit
}


function sleep(time) {
	return new Promise(async (resolve) => {
		setTimeout(resolve, time)
	})
}

(async () => {
		 await sleep(2500)
		 for (let element of document.getElementsByClassName("feedback")) {
			 (async () => {
				 console.log(element)
				 element.style.opacity = '0'
				 await sleep(1000)
				 element.style.padding = '0'
				 await sleep(1000)
				 element.remove()
			 })().then(r => {
			 })
		 }
	 }
)()