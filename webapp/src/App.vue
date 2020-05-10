<template>
  <div id="app">
    <h2 class="title">mapkicker</h2>
    <Mappool />
    <form id="chatbox">
      <textarea id="msgbox"></textarea>
      <input type="submit" value="送信" />
    </form>
  </div>
</template>

<script>
import Mappool from "./components/Mappool.vue";

export default {
  name: "App",
  components: {
    Mappool
  },
  data() {
    return {
      logoSVG: require("./assets/logo.svg")
    };
  }
};

window.onload = () => {
  let socket;
  const msgbox = document.getElementById("msgbox");
  const chatbox = document.getElementById("chatbox");
  chatbox.addEventListener("submit", e => {
    e.preventDefault();
    if (!msgbox.value) return false;
    if (!socket) {
      alert("Error: no websocket connection");
      return false;
    }
    socket.send(msgbox.value);
    msgbox.value = "";
    return false;
  });

  if (!window["WebSocket"]) {
    alert("Error, WebSocket isn't supported.");
  } else {
    const url = "ws://localhost:8080/echo";
    socket = new WebSocket(url);
    socket.onclose = () => {
      console.log(`Websocket connection to ${url} has been closed.`);
    };
    socket.onmessage = e => {
      console.table(e.data);
    };
    socket.onopen = () => {
      console.log("Websocket has opened");
    };
    socket.onerror = e => {
      console.error("Websocket has an error");
      console.table(e);
    };
  }
};

// alert("Hello");
</script>

<style>
body {
  margin-top: 5%;
  padding-right: 5%;
  padding-left: 5%;
  font-size: larger;
  font-family: -apple-system, BlinkMacSystemFont, "Segoe UI", "Roboto", "Oxygen",
    "Ubuntu", "Cantarell", "Fira Sans", "Droid Sans", "Helvetica Neue",
    sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
}

@media screen and (min-width: 800px) {
  body {
    padding-right: 15%;
    padding-left: 15%;
  }
}

@media screen and (min-width: 1600px) {
  body {
    padding-right: 30%;
    padding-left: 30%;
  }
}

code {
  font-family: source-code-pro, Menlo, Monaco, Consolas, "Courier New",
    monospace;
  background-color: #b3e6ff;
}

.title {
  text-align: center;
}

.logo {
  text-align: center;
}
</style>
