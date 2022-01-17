<template>
  <div id="main">
    <div id="main_cont">
      <div id="profile_page">

        <div id="menu_cont">
          <div id="info" class="menu_item active" v-on:click="toPage('info')">Info</div>
          <div id="my" class="menu_item" v-on:click="toPage('my')">My orders</div>
          <div id="other" class="menu_item" v-on:click="toPage('other')">Other orders</div>
          <div id="create" class="menu_item" v-on:click="toPage('create')">Create new</div>
          <div 
          id="userlist" 
          v-if="role == 'Admin'" 
          class="menu_item"
          v-on:click="toPage('userlist')"
          > Userlist
          </div>
        </div>
        <div style="display: grid; margin-left:16px" :key="currentPage">
          <div v-if="currentPage == 'info'" >
            <h2 class="header">Profile</h2>
            <div class="profile_item">Username: {{username}}</div>
            <div class="profile_item">Role: {{role}}</div>
            <div class="profile_item">Email: {{email}}</div>
            <div class="profile_item">Registration date: {{register}}</div>
          </div>

          <div id="orders" v-if="currentPage == 'my'">
            <h2 class="header">My orders</h2>
            <OrderList/>
          </div>
          <div id="orders" v-if="currentPage == 'other'">
            <h2 class="header">Other orders</h2>
            <OrderList/>
          </div>
          <div id="orders" v-if="currentPage == 'create'">
            <h2 class="header">New order</h2>
            <CreateOrder/>
          </div>
          <div id="orders" v-if="currentPage == 'userlist'">
            <h2 class="header">Userlist</h2>
            <UserList/>
          </div>
        </div>
        
      </div>
    </div>
  </div>
</template>

<script>
import UserList from '../components/userlist/Table.vue'
import client from '../client/client.js'
import CreateOrder from '../components/order/Create.vue'
import OrderList from '../components/order/List.vue'

export default {
  name: 'Profile',
  data(){
      return {
        username: null,
        register: null,
        role: null,
        email: null,
        currentPage:"info",
      }
  },
  beforeCreate() {
    client
      .get('/user/profile')
      .then((response) => {
        this.username = response.data["username"]
        this.register = (new Date(response.data["register"]*1000)).toLocaleString("ru")
        this.role = response.data["role"]
        this.email = response.data["email"]
      })
  },
  components: {
    UserList,
    CreateOrder,
    OrderList
  },
  methods: {
    toPage(id){
      let newCurrentPage = document.getElementById(id)
      let nowCurrentPage = document.getElementById(this.currentPage)

      nowCurrentPage.className = newCurrentPage.className
      newCurrentPage.className += " active"
      
      this.currentPage = id
    },
  },
}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
.profile_item{
    margin-top: 7px;
}
#panel_cont{
  padding: 20px 5px;
  height: 16px;
}
.panel_options_item{
  border: 1px solid #5f9ea0;
  border-radius: 7px;
  height: fit-content;
  width: fit-content;
}

#orders {
  height: 450px;
}

.order_menu_item{
  display: grid;
  text-align: center;
  font-family: 'Baloo Bhaijaan 2';
  font-size: 22px;
  font-weight: 500;
  color: #323134;
  cursor: pointer;
}

#order_menu{
  display: grid;
  grid-template-columns: 50% 50%;
}

#line{
  margin-top: 0px;
}

#menu_cont{
  width: fit-content;
  height: fit-content;

  border-radius: 7px;
  background: lightgrey;

  display: grid;
  margin-top: 31px;
}
.menu_item{
  color: #323134;
  font-family: 'Baloo Bhaijaan 2';
  font-weight: 500;
  padding:3px 8px;
  cursor: pointer;
  border-radius: 7px;
}
.menu_item:hover{
  background: #5e5c61;
  color: lightgrey;
}

.active{
  background: #323134;
  color: lightgrey;
}

.active:hover{
  background: #323134;
  color: lightgrey;
}

#profile_page{
  display: grid;
  grid-template-columns: 120px 1fr;
}
</style>