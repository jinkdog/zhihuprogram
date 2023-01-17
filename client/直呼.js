var cancel = document.getElementById("zheshiyigeanniu");
var idn = document.getElementsByClassName("Popover-content--bottom");
console.log(idn);
idn[0].style.display = "none";
cancel.onclick = function () {
	if (ind[0].style.display == "none") {

		ind[0].style.display = "block";
	}
	else if (ind[0].style.display == 'block') {
		ind[0].style.display = "none";
	}
}