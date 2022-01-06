<template>
  <div id="header_cont">
      <Item title="Main" url="/"/>
      <Item title="Profile" url="/profile"/>
      <Item title="About" url="/about"/>
      <AuthItem title="Sign in" url="/auth" :key="isAuth" v-if="!isAuth"/>
      <AuthItem title="Sign out" url="" :key="isAuth" v-if="isAuth"/>
  </div>
</template>

<script>
import Item from './item/Item.vue'
import AuthItem from './auth/Item.vue'
import {hasValue} from '../../actions/jwt.js'

export default {
  name: 'Header',
  components: {
    Item,
    AuthItem,
  },
  data(){
    return {
      isAuth: hasValue(),
    }
  },
  items: [
      {text:'Start'},
      {text:'About'},
  ],
  
}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
#header_cont{
    background: #5f9ea0;
    height: 60px;
    min-width: 680px;
    padding: 0px 20vw;
}
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