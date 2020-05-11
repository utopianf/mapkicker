<template>
  <div id="app">
    <header class="title">mapkicker</header>
    <!-- <h2 class="title">mapkicker</h2> -->
    <div class="holyGrail">
      <main class="content">
        <h3>Chat</h3>
        <Chat />
        <h3>Mappool</h3>
        <Mappool />
      </main>
      <nav class="side-nav"></nav>
      <aside class="side-bar"></aside>
    </div>
  </div>
</template>

<script>
import Mappool from "./components/Mappool.vue";
import Chat from "./components/Chat.vue";

export default {
  name: "App",
  components: {
    Mappool,
    Chat
  },
  data() {
    return {
      logoSVG: require("./assets/logo.svg")
    };
  }
};
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

  display: flex;
  min-height: 100vh;
  flex-direction: column;
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

header,
footer {
  flex: 0;
}

.holyGrail {
  display: flex;
  flex: 1;
}

.content {
  flex: 1
}

.side-nav,
.side-bar {
  width: 200px;
}

.side-nav {
  order: -1;
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
