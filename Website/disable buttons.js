function disable() {
    elements = document.body.getElementsByTagName("*");
    console.log(elements)
    for (var i = 0; i < elements.length; i++) {
        elements[i].disabled = elements[i].classList.contains('disabled')
    }
}

disable();
