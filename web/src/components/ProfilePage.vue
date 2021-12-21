<template>
  <div id="main">
    <Header />
    <div id="main_cont">
      <div if="profile_page">
        <h2 class="header">Profile</h2>
        <div class="profile_item">Username: {{username}}</div>
        <div class="profile_item">Role: {{role}}</div>
        <div class="profile_item">Email: {{email}}</div>
        <div class="profile_item">Registration date: {{register}}</div>
        <div :key="role" v-if="role=='Admin'" id="panel_cont">
          <UserList>
          </UserList>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
import Header from './header/Header.vue'
import axios from 'axios'
import {readValue} from '../actions/jwt.js'
import {refreshTokens} from '../actions/auth.js'
import UserList from './userlist/Table.vue'

export default {
  name: 'Profile',
  data(){
      return {
        username: null,
        register: null,
        role: null,
        email: null,
      }
  },
  beforeCreate() {
    axios
      .get('http://localhost:8000/user/profile', {
        headers: {
            "Authorization": "Bearer "+ readValue(),
        },
        withCredentials: true 
    })
      .then((response) => {
        this.username = response.data["username"]
        this.register = (new Date(response.data["register"]*1000)).toLocaleString("ru")
        this.role = response.data["role"]
        this.email = response.data["email"]
      })
      .catch(() => {
        if ( readValue() === ""){
          document.location.href = "/auth"
        }
        refreshTokens()
      });
  },
  components: {
    Header,
    UserList,
  },
}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
.profile_item{
    margin-top: 7px;
}
#panel_cont{
  padding: 20px;
  height: 16px;
}
.panel_options_item{
  border: 1px solid #5f9ea0;
  border-radius: 7px;
  height: fit-content;
  width: fit-content;
}
</style>