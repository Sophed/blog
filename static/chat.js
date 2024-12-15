const ws = new WebSocket("ws://localhost:1337");
ws.onopen = () => {
  ws.send("hello world");
};
ws.onmessage = (msg) => {
  console.log(msg.data);
  update(msg.data);
};

const inp = document.getElementById("msg-input");
inp.addEventListener("keyup", ({ key }) => {
  if (key === "Enter") {
    ws.send(inp.value);
    inp.value = "";
  }
});

function update(msg) {
  let ls = document.getElementById("lines");
  ls.getElementsByTagName("li")[0].remove();
  let entry = document.createElement("li");
  entry.appendChild(document.createTextNode(msg));
  ls.appendChild(entry);
}
