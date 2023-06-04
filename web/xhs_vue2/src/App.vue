<template>
  <div id="app">
    <!-- <img src="./assets/logo.png">
    <div>
      <p>
        这一段来自APP。vue
        <code v-text="'<el-button>'"></code>
        below
      </p>
      <el-button>el-button</el-button>
    </div> -->
    <HelloWorld msg="欢迎来到vue"/>
  </div>
</template>

<script>
//import HelloWorld from './components/HelloWorld.vue'
import HelloWorld from './views/PersonalView.vue'

export default {
  name: 'app',
  components: {
    HelloWorld
  }
}
</script>

<!-- <script>
export default {
  name: 'PersonalView',
  data() {
    return {
      msg: ''
    }
  },
  created() {
        this.$axios.get('/api/host').then(result=>{            
            this.msg = result.data;
        })
    }
}
</script> -->

<style>
#app {
  font-family: 'Avenir', Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  color: #2c3e50;
  margin-top: 60px;
}
</style>
