<template>
  <div id="app">
    <div>
  <router-view>
    
  </router-view>
  <!-- 可以让它跳转到index.js里 -->
   </div>
  </div>

</template>

<script>

export default {
  name: 'app',
}
</script>

<style>
#app {
  font-family: 'Avenir', Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  color: #2c3e50;
  margin-top: 0px;
}
</style>
