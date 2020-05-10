<template>
  <ul class="mappool">
    <li v-for="map in mappool" v-bind:key="map">
      <b>{{map}}</b>
    </li>
  </ul>
</template>

<script>
import axios from "axios";

export default {
  name: "Mappool",
  data() {
    return {
      mappool: []
    };
  },
  mounted() {
    axios.get(`${process.env.VUE_APP_API_URL}/api/mappool`).then(response => {
      return (this.mappool = response.data);
    });
  }
};
</script>

<style scoped>
.mappool {
  margin-top: 5px;
}
</style>
