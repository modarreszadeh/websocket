
var input = document.getElementById("input");
var output = document.getElementById("output");

var socket = new WebSocket("ws://127.0.0.1:8080/websocket");

socket.onopen = function () {
    output.innerHTML = "";
    output.innerHTML += "Status: Online\n\n";
};

socket.onclose = function () {
    output.innerHTML = "";
    output.innerHTML += "Status: Offline\n\n";
}

socket.onmessage = function (e) {
    output.innerHTML += "[" + navigator.platform + "]: " + e.data + "\n";
};

function send() {
    socket.send(input.value);
    input.value = "";
}