<template>
  <div id="header_cont">
      <Item title="Main" url="/"/>
      <Item title="Profile" url="/profile"/>
      <Item title="About" url="/about"/>
      <AuthItem title="Sign in" url="/auth" :key="isAuth" v-if="isAuth"/>
      <AuthItem title="Sign out" :key="isAuth" v-if="!isAuth"/>
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
    min-width: 630px;
    padding: 0px 20vw;
}
</style>