const e_text_box = document.getElementById("text_box");
const e_submit_button = document.getElementById("submit_button");

console.log(document.cookie)

window.addEventListener("load", function () {
  if (window.innerWidth < 1000) {
    e_text_box.style.width = window.innerWidth + "px"
  } else {
    e_text_box.style.width = "1000px"
  }
}, false)
window.addEventListener("resize", function () {
  if (window.innerWidth < 1000) {
    e_text_box.style.width = window.innerWidth + "px"
  }
  console.log("run to listener once")
}, false)

e_text_box.addEventListener("click", function (){
  if (e_text_box.value === "Write your problems or any other things that want to share with us there ..."){
    e_text_box.value = "";
  }
});

e_submit_button.addEventListener("click", function (){
  e_submit_button.disabled = true;
  let sock = new XMLHttpRequest();
  sock.open("POST", "/ask", false)

  let id = hashCode(e_text_box.value, 1048576).toString();
  console.log("str:" + e_text_box.value)
  console.log("id:" + id);
  let data = {
    "user_id": id,
    "question": e_text_box.value
  };
  let str_data = JSON.stringify(data);
  console.log(str_data)
  sock.send(str_data);
  let i = 0;
  while (checkCookie("user_id" + i)){
    i += 1;
  }

  console.log("for user_id:" + i);
  setCookie("user_id" + i, id, 1000);

  document.body.innerHTML = "<h1 style=\"text-align: center; font-size: 60px\">THANK YOU</h1>\n" +
      "<h1 style=\"text-align: center; font-size: 40px\">Your questions have been successfully collected by us</h1>\n" +
      "<h1 style=\"text-align: center; font-size: 30px\">click <a href=\"/\">here</a> to back to the index page</h1>"
})

function hashCode(str, size){
  let hashCode = 0;

  for(let i =0; i<str.length; i++){
    hashCode += str.charCodeAt(i)
  }
  //取余操作
  return hashCode % size + new Date().getMilliseconds();
}

function setCookie(cname,cvalue,exdays){
  const d = new Date();
  d.setTime(d.getTime()+(exdays*24*60*60*1000));
  const expires = "expires=" + d.toUTCString();
  document.cookie = cname+"="+cvalue+"; "+expires + ";path=/";
}
function getCookie(cname){
  const name = cname + "=";
  const ca = document.cookie.split(';');
  for(let i=0; i<ca.length; i++) {
    const c = ca[i].trim();
    if (c.indexOf(name)===0) { return c.substring(name.length,c.length); }
  }
  return "";
}
function checkCookie(name){
  return getCookie(name) !== "";
}