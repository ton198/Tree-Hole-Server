
console.log(document.cookie)

window.addEventListener("load", function () {
  let i = 0;

  let eRepliesContainer = document.getElementById("replies_container");
  while (true) {
    let userId = getCookie("user_id" + i);
    if (userId === "") {
      break
    }
    let sock = new XMLHttpRequest();
    sock.open("POST", "/reply", false);
    let request_data = {
      user_id: userId
    };
    sock.send(JSON.stringify(request_data));
    let response;
    try {
      response = JSON.parse(sock.responseText);
    } catch {
      i++;
      continue
    }
    let eQuestionTag = document.createElement("h1");
    eQuestionTag.innerText = "Your questions:";
    eQuestionTag.style.fontFamily = "dodge_rabbit,serif"

    let eQuestion = document.createElement("p");
    eQuestion.innerText = response.question;
    eQuestion.style.fontFamily = "dodge_rabbit,serif"
    eQuestion.style.fontSize = "25px"

    let eReplyTag = document.createElement("h1");
    eReplyTag.innerText = "Our replies:";
    eReplyTag.style.fontFamily = "dodge_rabbit,serif"

    let eReply = document.createElement("p");
    eReply.innerText = response.reply;
    eReply.style.fontFamily = "dodge_rabbit,serif"
    eReply.style.fontSize = "25px"

    eRepliesContainer.appendChild(eQuestionTag);
    eRepliesContainer.appendChild(eQuestion);
    eRepliesContainer.appendChild(eReplyTag);
    eRepliesContainer.appendChild(eReply);
    eRepliesContainer.appendChild(document.createElement("hr"));
    i++;
  }

  if (i <= 0) {
    let imgBox = document.createElement("img");
    imgBox.src = "/img/NONE.png";

    eRepliesContainer.appendChild(imgBox);
  }
});

function getCookie(cname){
  const name = cname + "=";
  const ca = document.cookie.split(';');
  for(let i=0; i<ca.length; i++) {
    const c = ca[i].trim();
    if (c.indexOf(name)===0) { return c.substring(name.length,c.length); }
  }
  return "";
}
