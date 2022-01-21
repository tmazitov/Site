<template>
  <div id="main">
    <div id="main_cont">
      <div id="auth_cont" :key="signIn">
          <h2 class="header" v-if="signIn">Sign in</h2>
          <h2 class="header" v-if="!signIn">Sign up</h2>
            <div class="auth_item">
                <div class="sub_header">Login</div>
                <input id="username" type="text" class="auth_field"/>
            </div>
            <div class="auth_item"  v-if="!signIn">
                <div class="sub_header">Email</div>
                <input id="email" type="email" class="auth_field"/>
            </div>
            <div class="auth_item">
                <div class="sub_header">Password</div>
                <input id="password" type="password" class="auth_field"/>
            </div>
            <div class="auth_noacc"  v-if="signIn">
                <a>Missing password? Change them 
                    <span style="color:#5f9ea0;font-weight:500">
                        here
                    </span>!
                </a>   
            </div>
            <div id="auth_err">
            </div> 
            <div id="auth_btn_cont" > 
                <button id="auth_btn" v-if="signIn" v-on:click="doSignIn()">
                    Sign in
                </button>
                <button id="auth_btn" v-if="!signIn" v-on:click="doSignUp()"> 
                    Sign up
                </button>
            </div>
            <div class="auth_noacc" v-if="signIn">
                <a>Don't have account? Create 
                    <span class="auth_link" v-on:click="changeAuth()">
                        here
                    </span>!
                </a>   
            </div>
            <div class="auth_noacc" v-if="!signIn">
                <a>Wait... I have 
                    <span class="auth_link" v-on:click="changeAuth()">
                        account
                    </span>!
                </a>   
            </div>
      </div>
    </div>
  </div>
</template>

<script>
import {signInAction, signUpAction} from '../actions/auth.js'

const validate = function(email) {
   var reg = /^([A-Za-z0-9_\-\.])+\@([A-Za-z0-9_\-\.])+\.([A-Za-z]{2,4})$/;//eslint-disable-line
   return !reg.test(email)
}

export default {
  name: 'Auth',
  props: {
    msg: String
  },
  data() {
    return {
        signIn: true,
    };
  },
  methods:{
    changeAuth: function() {
        this.signIn = !this.signIn
    },
    doSignIn: function(){
        var errField = document.getElementById("auth_err")
        var username = document.getElementById("username").value;
        var password = document.getElementById("password").value;

        signInAction(username, password, errField)
    },
    doSignUp: function(){
        var errField = document.getElementById("auth_err")
        var username = document.getElementById("username").value;
        if (username.length < 5){
            errField.innerHTML = "Too short username"
            return
        }
        var email = document.getElementById("email").value;
        if (validate(email)){
            errField.innerHTML = "Invalid email"
            return
        }
        var password = document.getElementById("password").value;
        if (password.length < 8){
            errField.innerHTML = "Too short password"
            return
        }
        
        signUpAction(username, password, email, errField)
    },
  },
}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
    #auth_err{
        color: red;
    }
    #auth_cont{
        margin-top: 80px;
        margin-left: auto;
        margin-right: auto;
        height: 400px;
        width: 300px;
        border: 2px solid #9BC3C4;
        border-radius: 7px;
        margin-left: auto;
        margin-right: auto;
        padding: 25px 15px ;
        height: fit-content;
    }
    .auth_field{
        margin-top: 3px;
        border: 2px solid #9BC3C4;
        border-radius: 7px;
        height: 20px;
        padding: 3px 8px;
    }
    .auth_field:focus{
        outline: none !important;
        border-color: wheat;
    }
    .auth_item{
        margin-top: 18px;
    }
    .auth_noacc{
        font-size: 14px;
        margin-top: 3px;
    }
    #auth_btn_cont{
        margin: 18px 0px;
    }
    #auth_btn{
        height: 36px;
        border-radius: 7px;
        border: 0px solid;
        background: #5f9ea0;
        font-family: "Baloo Bhaijaan 2";
        font-size: 17px;
        font-weight:500 ;
        width: 200px;
        cursor: pointer;
    }
    #auth_btn:hover{
    background: #588C8E;
    color:wheat;
    }
    #auth_btn:active{
    background: #507A7C;
    }
    .auth_link{
        color:#5f9ea0;
        font-weight:500;
        cursor: pointer;
    }
    
</style>