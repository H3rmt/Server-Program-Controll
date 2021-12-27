function saveSettings(name, event) {
	if(getAuthorisationCookie() === "") {
		event.preventDefault()
		alert(`Not allowed to save ${name}`)
		console.log(`Not allowed to save ${name}`)
		return
	}
}

function resetSettings(name, event) {
	if(getAuthorisationCookie() === "") {
		event.preventDefault()
		alert(`Not allowed to reset ${name}`)
		console.log(`Not allowed to reset ${name}`)
		return
	}
}