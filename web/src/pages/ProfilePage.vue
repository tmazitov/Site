<template>
  <div id="main">
    <div id="main_cont">
      <div id="profile_page">
        <vue-final-modal v-model="showModal" name="modal" v-if="role == 'Admin'" >
          <div id="modal_cont">
            <div id="close_btn">
              <unicon  
              name="times" 
              fill="red"
              v-on:click="showModal = false"
              > 
              </unicon>
            </div>
            <AdminPanel/>
          </div>
        </vue-final-modal>
        <div id="menu_none" v-if="role=='User'"></div>
        <div id="menu_cont" v-if="role == 'Worker' || role == 'Manager' || role == 'Admin'">
          <div id="info" class="menu_item active" v-on:click="toPage('info')">Info</div>
          <div 
            id="my" 
            class="menu_item"
            v-if="role == 'Worker' || role == 'Manager'"
            v-on:click="toPage('my')"
            >My orders
          </div>
          <div 
          id="other" 
          class="menu_item" 
          v-on:click="toPage('other')"
          v-if="role == 'Worker' || role == 'Manager' || role == 'Admin'"
          >Order list
          </div>
          <div 
            id="create" 
            class="menu_item" 
            v-on:click="toPage('create')"
            v-if="role == 'Admin' || role == 'Manager'"
            >
            New order
          </div>
          <div 
          id="userlist" 
          v-if="role == 'Admin'" 
          class="menu_item"
          @click="openModalExample"
          > Admin panel
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

          <div id="selected_cont" v-if="currentPage == 'my' && (role == 'Worker' || role == 'Manager')">
            <h2 class="header">My orders</h2> 
            <OrderList :username="username" :role="role" mode="my"/>
          </div>
          <div id="selected_cont" v-if="currentPage == 'other'">
            <h2 class="header">Order list</h2>
            <OrderList :username="username" :role="role" mode="all"/>
          </div>
          <div id="selected_cont" v-if="currentPage == 'create'">
            <h2 class="header">New order</h2>
            <CreateOrder/>
          </div>
        </div>
        
      </div>
    </div>
  </div>
</template>

<script>
import { VueFinalModal } from 'vue-final-modal'

import client from '../client/client.js'
import CreateOrder from '../components/order/Create.vue'
import OrderList from '../components/order/List.vue'
import AdminPanel from '../components/adminpanel/AdminPanel.vue'


export default {
  name: 'Profile',
  data(){
      return {
        username: null,
        register: null,
        role: null,
        email: null,
        currentPage:"info",
        showModal: false,
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
    CreateOrder,
    OrderList,
    VueFinalModal,
    AdminPanel
  },
  methods: {
    toPage(id){
      let newCurrentPage = document.getElementById(id)
      let nowCurrentPage = document.getElementById(this.currentPage)

      nowCurrentPage.className = newCurrentPage.className
      newCurrentPage.className += " active"
      
      this.currentPage = id
    },
    openModalExample() {
      this.$vfm.show('modal')
    }
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

#selected_cont {
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

#modal_cont{
  margin: 0px 23vw;
  margin-top: 50px;
  padding: 20px;
  background: white;
  border-radius: 11px;
      overflow-y: auto;
    max-height: 600px;
}
</style>