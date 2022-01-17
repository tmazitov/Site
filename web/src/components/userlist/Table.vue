<template>
  <div>
    <date-picker 
    v-model="date"
    @change="setDate"
    @clear="delDate"
    placeholder="Select date"
    ></date-picker>
    <table name="userlist" class="table" >
      <thead>
        <tr>
          <th>Username</th>
          <th>Role</th>
          <th>Creation date</th>
        </tr>
      </thead>
      <tbody v-if="users.length != 0">
        <tr v-for="user in users" :key="user.Username">
          <td>{{user.Username}}</td>
          <td>{{user.Role}}</td>
          <td>{{dateFormat(user.Register)}}</td>
        </tr>
      </tbody>
    </table>
    <div v-if="users.length == 0">
        <p style="text-align: center">
          No data about users.
        </p>
    </div>
    <div id="buttons_cont" v-if="buttonsCount < 10">
      <button
      v-for="i in buttonsCount"
      :id="'pag_' + i" 
      v-on:click="setCurrentPage(i)"
      :key="i" 
      class="button_item">
      {{i}}
      </button>
    </div>
    <div id="buttons_cont" v-if="buttonsCount >= 10">
      <button
      v-for="i in 3"
      :id="'pag_' + i" 
      v-on:click="setCurrentPage(i)"
      :key="i" 
      class="button_item">
      {{i}}
      </button>
      <button
      :id="'pag_' + currentPage - 2" 
      v-on:click="setCurrentPage(currentPage - 2)"
      :key="currentPage - 2" 
      class="button_item">
      {{currentPage - 2}}
      </button>
      <button
      :id="'pag_' + currentPage - 1" 
      v-on:click="setCurrentPage(currentPage - 1)"
      :key="currentPage - 1" 
      class="button_item">
      {{currentPage - 1}}
      </button>
      <button
      :id="'pag_' + currentPage" 
      v-on:click="setCurrentPage(currentPage)"
      :key="currentPage" 
      class="button_item">
      {{currentPage}}
      </button>
      <button
      :id="'pag_' + currentPage + 1" 
      v-on:click="setCurrentPage(currentPage + 1)"
      :key="currentPage + 1" 
      class="button_item">
      {{currentPage + 1}}
      </button>

    </div>
  </div>
</template>

<script>
import DatePicker from 'vue2-datepicker';
import 'vue2-datepicker/index.css';

import {readValue} from '../../actions/jwt'
import {refreshTokens} from '../../actions/auth'

import client from '../../client/client'

export default {
  name: "UserList",
  components: {
    DatePicker
  },

  created() {
    this.getUsersData()
  },
  data: function() {
    return {
      time: 0,
      date: null,
      dateOptions:{ 
        weekday: 'long', 
        year: 'numeric', 
        month: 'long', 
        day: 'numeric' 
      },
      users: [],
      currentPage: 1,
      buttonsCount: 1,
      perPage: 6,
    };
  },

  watch: {
    time : function() {
      this.currentPage = 1
      this.getUsersData()
    },
    currentPage: function() {
      this.getUsersData()
    }
  },
  
  methods: {
    setDate(date){
      this.time = Date.parse(date) / 1000;
    },
    delDate(){
      this.time = 0;
    },
    setCurrentPage(num){
      document.getElementById("pag_"+this.currentPage).className = "button_item"
      this.currentPage = num
    },
    getUsersData(){

      let token = readValue()

      client.get(
        "/user/list?"
        + "timestamp=" + this.time + "&"
        + "page=" + this.currentPage + "&"
        + "per_page=" + this.perPage ,
        {
          headers: {
            "Authorization": "Bearer "+ token,
          },
          withCredentials: true 
        }
        ).then((response) => {
          this.users = response.data["users"]
          document.getElementById("pag_"+this.currentPage).className += " activate"

          if (this.users.length == this.perPage){
            this.buttonsCount = this.currentPage + 1
          } else {
            this.buttonsCount = this.currentPage
          }
        }).catch(() => {
            if ( readValue() === ""){
                this.$router.push('auth')
            }
            refreshTokens()
        })
    },
    dateFormat(value){
      let d = new Date(value * 1000)
      return d.toLocaleString("ru")
    }
  }
};
</script>

<style>

  .button_item{
    height: 40px;
    width: 40px;
    background-color: white;
    border: 2px solid #5f9ea0;
    border-radius: 10px;
    font-weight: 500;
    font-family: 'Baloo Bhaijaan 2';
    font-size: 16px;
    margin-left: 10px;
    color:#5f9ea0;
  }
  
  .button_item:hover{
    background: #588C8E;
    color: wheat;
    border-color: #588C8E;
  }

  .button_item.activate{
    background: #588C8E;
    color: wheat;
    border-color: #588C8E;
  }

  .button_item:active{
    background: #507A7C;
    color: wheat;
    border-color: #507A7C;
  }

  .pagination {
    margin-top: 1rem;
  }
  .ui.blue.table {
    border-top: 0.2em solid #5f9ea0;
  }
  .mx-input-wrapper{
    margin-bottom: 12px;
  }
  .mx-input:hover{
    border-color: #5f9ea0;
  }
  .mx-btn:hover {
    border-color: #5f9ea0;
    color: #75c1c4;
  }
  .mx-calendar-content .cell:hover {
    background-color: #9BC3C4;
    color:wheat
  }
  .mx-table-date .today {
    color: #75c1c4;
  }
  .mx-calendar-content .cell.active {
    color: wheat;
    background-color: #5f9ea0;
  }
  .table {
  width: 100%;
  border: none;
  margin-bottom: 10px;
  border-collapse: separate;
  margin-top: 10px;
  }
  .table thead th {
  font-weight: 500;
  font-family: 'Baloo Bhaijaan 2';
  font-size: 15px;
  text-align: left;
  border: none;
  padding: 10px 15px;
  background: #5f9ea0;
  border-top: 1px solid #9BC3C4;
  color: #323134
  }
  .table tr th:first-child, .table tr td:first-child {
  border-left: 1px solid #9BC3C4;
  }
  .table tr th:last-child, .table tr td:last-child {
  border-right: 1px solid #9BC3C4;
  }
  .table thead tr th:first-child {
  border-radius: 20px 0 0 0;
  }
  .table thead tr th:last-child {
  border-radius: 0 20px 0 0;
  }
  .table tbody td {
  text-align: left;
  border: none;
  padding: 10px 15px;
  font-size: 14px;
  vertical-align: top;
  }
  .table tbody tr:nth-child(even) {
  background: #9BC3C4;
  }
  .table tbody tr:last-child td{
  border-bottom: 1px solid #ddd;
  }
  .table tbody tr:last-child td:first-child {
  border-radius: 0 0 0 20px;
  }
  .table tbody tr:last-child td:last-child {
  border-radius: 0 0 20px 0;
  }  
</style>