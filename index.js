let url = "wss://" + window.location.host + window.location.pathname + "/ws";
let ws = new WebSocket(url);

let chat = document.getElementById("chat");
let  text = document.getElementById("text");
let name = document.getElementById("name");
let button = document.getElementById("button");

ws.onmessage = function (msg) {
   var line = msg.data + "\n";
   chat.innerText = line + chat.innerText;
};

function Send() {
   if (text.value !== "") {
      ws.send(name.value + ": " + text.value);
      text.value = "";
   }
}
