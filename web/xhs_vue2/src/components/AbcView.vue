<template>
    <div class="hello">
      <h1>{{ msg }}</h1>
      <p>
        gogogo,{{ msg }}
       
      </p>
  
     
    </div>
  </template>
  
  <script>
  
  export default {
    name: 'AbcView',
    data() {
      return {
        msg: ''
      }
    },
    created() {
          this.$axios.get('/api/hst').then(result=>{            
              this.msg = result.data;
          })
      }
  }
  </script>
  
  <!-- Add "scoped" attribute to limit CSS to this component only -->
  <style scoped>
  h3 {
    margin: 40px 0 0;
  }
  ul {
    list-style-type: none;
    padding: 0;
  }
  li {
    display: inline-block;
    margin: 0 10px;
  }
  a {
    color: #42b983;
  }
  </style>