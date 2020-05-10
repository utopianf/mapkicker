<template>
  <div>
    <ul class="messages">
      <li v-for="msg in messages" v-bind:key="msg.id">
        <b>{{msg.content}}</b>
      </li>
    </ul>
    <form id="chat_form">
      <input type="text" id="chat_text" />
    </form>
  </div>
</template>

<script>
// import axios from "axios";

export default {
  name: "Chat",
  data() {
    return {
      messages: []
    };
  },
  mounted() {
    const url = "ws://localhost:8080/join";
    const socket = new WebSocket(url);
    socket.onclose = () => {
      console.log(`Websocket connection to ${url} has been closed.`);
    };
    socket.onmessage = e => {
      console.table(e.data);
      this.messages.push({
        id: this.messages.length,
        content: e.data
      });
    };
    socket.onopen = () => {
      console.log(`Websocket connection to ${url} has opened.`);
    };
    socket.onerror = () => {
      console.error(`Websocket connection to ${url} has an error.`);
    };
    const form = document.getElementById("chat_form");
    form.addEventListener("submit", e => {
      e.preventDefault();
      const textInput = document.getElementById("chat_text");
      if (!textInput.value) return false;
      if (!socket) {
        console.error("Error: no websocket connection");
        return false;
      }
      socket.send(textInput.value);
      textInput.value = "";
      return false;
    });
  }
};
</script>

<style scoped>
.mappool {
  margin-top: 5px;
}
</style>

