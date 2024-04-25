let socket = new WebSocket("ws://" + window.location.host + "/ws");
let person = "";

socket.onopen = function (e) {
  console.log("[open] Connection established");
  person = prompt("Please enter your name", "Harry Potter") || "Guest";
  if (person != null) {
    console.log("hi " + person);
  }
};

socket.onmessage = function (event) {
  let messages = document.getElementById("messages");
  let messageItem = document.createElement("li");
  messageItem.textContent = person + ": " + event.data;
  messages.appendChild(messageItem);
};

socket.onclose = function (event) {
  if (event.wasClean) {
    console.log(
      `[close] Connection closed cleanly, code=${event.code} reason=${event.reason}`
    );
  } else {
    console.log("[close] Connection died");
  }
};

socket.onerror = function (error) {
  console.log(`[error] ${error.message}`);
};

function sendMessage() {
  let messageInput = document.getElementById("messageInput");
  let message = messageInput.value;
  socket.send(message);
  messageInput.value = ""; // Clear input field
}
