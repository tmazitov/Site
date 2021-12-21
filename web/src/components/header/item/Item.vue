<template>
  <div class="header_item" v-on:click="goTo(url)">
      {{title}}
  </div>
</template>

<script>
export default {
  name: 'Item',
  props: {
    title: String,
    url: String,
  },
  methods:{
    goTo: function(url){
      window.location.href = url;
    }
  }
}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
.header_item{
  padding: 20px 26px;
  float: left;
  color: #323134;
  font-weight: 500;
  font-family: 'Baloo Bhaijaan 2';
  font-size:16px;
  cursor: pointer;
}
.header_item:hover{
  background: #588C8E;
  color:wheat;
}
.header_item:active{
  background: #507A7C;
}
</style>