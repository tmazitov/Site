<template>
  <div id="main">
    <Header />
    <div id="main_cont">
      <div class="main_page">
        <h2 class="header">About</h2>
        <div class="text_cont">
          <a>Hi, my name is Timur. </a>
          <a>I create this site to learn Vue framework and Golang http server. </a>  
          <a>I don't know about future of this site, but I hope this will be amazing!</a> 
        </div>
        <div class="text_cont">
          <a>If you want, you can see my </a> 
          <a id="link" href="https://github.com/tmazitov">github</a>.
          <a> Sometimes I post interesting projects there.</a> 
        </div>
      </div>
    </div>
  </div>
</template>

<script>
import Header from './header/Header.vue'

export default {
  name: 'About',
  props: {
    msg: String
  },
  components: {
    Header,
  }
}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
.text_cont{
    margin-top:7px;
}
#link{
    color:#5f9ea0;
    font-weight:500;
    cursor: pointer;
}
</style>
