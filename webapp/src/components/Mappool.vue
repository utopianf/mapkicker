<template>
  <div>
    <div class="mappool">
      <div v-for="map in mappool" v-bind:key="map">
        <input type="checkbox" :id="map" :value="map" v-model="checkedMaps" />
        <label :for="map">{{map}}</label>
      </div>
    </div>
    <!-- <span>Checked Maps: {{ checkedMaps }}</span>
    <br /> -->
    <button v-on:click="kickMap">Kick!</button>
  </div>
</template>

<script>
import axios from "axios";
import Vue from "vue";

export default Vue.component("Mappool", {
  name: "Mappool",
  data: () => ({
    mappool: [],
    checkedMaps: [],
    socket: null,
    url: "ws://localhost:8080/join"
  }),
  mounted: function() {
    this.socket = new WebSocket(this.url);
    axios.get(`${process.env.VUE_APP_API_URL}/api/mappool`).then(response => {
      return (this.mappool = response.data);
    });

    this.socket.onclose = () => {
      console.log(`Websocket connection to ${this.url} has been closed.`);
    };
    this.socket.onopen = () => {
      console.log(`Websocket connection to ${this.url} has opened.`);
    };
    this.socket.onerror = () => {
      console.error(`Websocket connection to ${this.url} has an error.`);
    };
  },
  methods: {
    kickMap: function() {
      console.log(this.checkedMaps);
      this.socket.send(`I kick ${this.checkedMaps}`);
    }
  }
});
</script>

<style scoped>
.mappool {
  display: flex;
  justify-content: space-around;
  flex-wrap: wrap;
}
</style>
