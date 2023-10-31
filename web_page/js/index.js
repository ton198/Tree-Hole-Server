
eButtonContainer = document.getElementById("button_container");

window.addEventListener("resize", function () {
  if (window.innerWidth < 1100) {
    eButtonContainer.style.width = "510px"
  } else {
    eButtonContainer.style.width = "1100px"
  }
});
