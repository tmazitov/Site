<template>
  <div id="header_cont" :key="isAuth()">
      <Item title="Main" url="/"/>
      <Item title="Profile" url="/profile"/>
      <Item title="About" url="/about"/>
      <SignIn title="Sign in" url="/auth" v-if="!isAuth()"/>
      <SignOut title="Sign out" v-if="isAuth()"/>
  </div>
</template>

<script>
import Item from './item/Item.vue'
import SignIn from './auth/SingIn.vue'
import SignOut from './auth/SignOut.vue'
import {tokenIsValid} from '../../actions/jwt.js'

export default {
  name: 'Header',
  components: {
    Item,
    SignIn,
    SignOut
  },
  data() {
    return {
      isAuth : tokenIsValid,
    }
  },
}
</script>

<style scoped>
#header_cont{
    background: #5f9ea0;
    height: 60px;
    min-width: 680px;
    padding: 0px 21vw;
}
.header_item{
  padding: 16px 26px;
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