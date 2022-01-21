<template>
    <div id="order_cont">
        <vue-final-modal v-model="showAbout" name="about"  >
          <div id="modal_cont">
            <div id="close_btn">
              <unicon  
              name="times" 
              fill="red"
              v-on:click="showAbout = false"
              > 
              </unicon>
            </div>
            <h2 class="header">{{about.Title}}</h2>
            <div>Writer: {{about.Writer}}</div>
            <div>Worker: {{about.Worker}}</div>
            <div>Status: {{about.Status}}</div>
            <div>Price: {{about.Price}} $</div>
            <div>Comment: {{about.Comment}}</div>
            <div 
            v-if="about.Writer == username" 
            id="del_btn" 
            v-on:click="deleteOrder()">
            Delete
            </div>
            <div class="date">{{dateFormat(about.Date)}}</div>
          </div>
        </vue-final-modal>
        <vue-final-modal v-model="showUpdate" name="update"  >
          <div id="modal_cont">
            <div id="close_btn">
              <unicon  
              name="times" 
              fill="red"
              v-on:click="showUpdate = false"
              > 
              </unicon>
            </div>
            <h2 class="header">Update order</h2>
            <div class="create_item">
            <div class="sub_header">Title</div>
            <input 
            id="title" 
            class="create_field long"
            :value="about.Title"
            placeholder="About your order"
            />
            </div>
            <div class="create_item">
                <div class="sub_header">Comment</div>
                <textarea 
                id="comment"
                :value="about.Comment" 
                class="create_field long large"
                placeholder="Help you get the job done with precision"
                />
            </div>
            <div id="create_btn_cont">
                <button id="create_btn" v-on:click="udpate()">
                Update
                </button>  
            </div>
          </div>
        </vue-final-modal>
        <div v-for="order in orders" :key="order.UUID" class="order" > 
            <div class="order_header">{{titleFormat(order.Title)}}</div>
            <div>Writer: {{order.Writer}}</div>
            <div>Status: {{order.Status}}</div>
            <div>Price: {{order.Price}} $</div>
            <div 
            v-if="order.Status == 'wait' && role == 'Worker'"
            class="take_btn"
            v-on:click="setWorker(order.UUID)"
            >Take
            </div>
            <div 
            v-if="order.Status == 'in process' && order.Worker == username"
            class="take_btn"
            v-on:click="done(order.UUID)"
            >Done
            </div>
            <div 
            v-if="role == 'Manager' || role == 'Worker' || role == 'Admin'"
            class="info_btn"
            v-on:click="getInfo(order.UUID)"
            >Info
            </div>
            <div 
            v-if="order.Writer == username && (role == 'Manager' || role == 'Admin')"
            class="up_btn"
            v-on:click="setUpdate(order.UUID)"
            >Update
            </div>
            <div class="date">{{dateFormat(order.Date)}}</div>
        </div>
    </div>
</template>

<script>

import client from '../../client/client'

import {readValue } from '../../actions/jwt'
import {refreshTokens} from '../../actions/auth'

import { VueFinalModal } from 'vue-final-modal'
export default {
    name: "OrderList",
    props: {
        username: String,
        role: String,
        mode: String,
    },
    components: {
        VueFinalModal
    },
    beforeMount() {
        client
            .get("/order/list?mode="+ this.mode).then((response) => {
                this.orders = response.data.orders
            })
        this.intervalId = setInterval(() => {
            if (!this.showUpdate) {
                client
            .get("/order/list?mode="+ this.mode).then((response) => {
                this.orders = response.data.orders
            }).catch(() => {
                if ( readValue() === ""){
                    this.$router.push('auth')
                }
                refreshTokens()
            })
            }
        }, 5000)
    },
    destroyed() {
        clearInterval(this.intervalId)
    },
    data() {
        return {
            orders: [],
            intervalId: 0,
            about: {
                Title: "",
                Writer: "",
                Worker: "",
                Date: "",
                Status: "",
                Comment: "",
                Price: 0,
            },
            showAbout: false,
            showUpdate: false,
            UUID: "",
        }
    },
    methods: {
        titleFormat(title){
            if (title.length > 19) {
                return title.slice(0,19) + "..."
            } else {
                return title
            }
        },
        dateFormat(value){
            let d = new Date(value * 1000)
            return d.toLocaleString("ru")
        },
        setWorker(uuid){
            client
            .get("/order/worker?order=" + uuid)
        },
        done(uuid){
            client
            .get("/order/complite?order=" + uuid)
        },
        getInfo(uuid){
            client
            .get("/order/get?order=" + uuid)
            .then((response) => {
                this.about.Writer = response.data.order.Writer
                this.about.Title = response.data.order.Title
                this.about.Worker = response.data.order.Worker
                this.about.Comment = response.data.order.Comment
                this.about.Date = response.data.order.Date
                this.about.Status = response.data.order.Status
                this.about.Price = response.data.order.Price
                this.$vfm.show('about')
            })
        },
        setUpdate(uuid){
            this.UUID = uuid
            client
            .get("/order/get?order=" + uuid).then((response) => {
                this.about.Title = response.data.order.Title
                this.about.Comment = response.data.order.Comment
                this.$vfm.show('update')
            })
            this.$vfm.show('update')
        },
        udpate(){

            var title = document.getElementById("title")
            var comment = document.getElementById("comment")

            client
            .put("/order/update", {
                
                Title: title.value,
                Comment: comment.value,
                UUID: this.UUID,
            }).then(() => {
                this.$swal({
                    title: 'Success',
                    text:'Order was successfully updated!',
                    icon: "success",
                });
            }).catch((error)=> {
                if ( error.status == 400){
                    this.$swal({
                    title:'Error',
                    text: 'Invalid credentials',
                    icon: "error",
                    });
                } else if (error.status == 500) {
                    this.$swal({
                    title:'Error',
                    text: 'Server error',
                    icon: "error",
                    }); 
                }
            })
        },
        deleteOrder(){
            client
            .get("/order/delete?order=" + this.uuid)
            .then(() => {
                this.$swal({
                    title: 'Success',
                    text:'Order was successfully deleted!',
                    icon: "success",
                });
            }).catch((error) => {
                if (error.status == 500) {
                    this.$swal({
                    title:'Error',
                    text: 'Server error',
                    icon: "error",
                    }); 
                }
            })    
        }
    },
}
</script>

<style scoped>

    #order_cont{
        height: 487px;
        overflow-y: auto;
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
    .order{
        border: 1.5px solid #5f9ea0;
        border-radius: 7px;
        padding: 9px 16px;
        margin-top: 16px;
        margin-right: 16px;
        width: 240px;
        display: inline-block;
    }
    .order_header{
        font-size: 22px;
        font-family: "Baloo Bhaijaan 2";
        font-weight: 500;
        color: #323134;
    }
    .date{
        color: rgb(170, 170, 170);
        font-size: 11px;
        margin-top: 10px;
        display: inline-block;
    }
    .complite {
        height: 20px;
        width: 20px;
        background: #5f9ea0;
        border-radius: 7px;
        display: inline-block;
        margin-left: 50px;
    }
    .complite:hover {
        background: #588C8E;
    }
    .complite:active {
        background: #477172;
    }
    .update {
        height: 20px;
        width: 20px;
        background: #ffce44;
        border-radius: 7px;
        display: inline-block;
        margin-left: 50px;
    }
    .update:hover {
        background: #daaf3a;
    }
    .update:active {
        background: #bb9632;
    }
    .take_btn{
        background: #5f9ea0;
        border-radius:7px;
        width: fit-content;
        padding: 3px 8px;
        color: #323134;
        font-family: 'Baloo Bhaijaan 2';
        font-weight: 500;
        cursor: pointer;
        float: right;
    }
    .take_btn:hover{
        background: #588C8E;
        color: wheat;
    }
    .take_btn:active {
        background: #507A7C;
    }
    .up_btn{
        border-radius:7px;
        width: fit-content;
        padding: 3px 8px;
        color: #ffce44;
        font-family: 'Baloo Bhaijaan 2';
        font-weight: 500;
        cursor: pointer;
        float: right;
    }
    .info_btn{
        border-radius:7px;
        width: fit-content;
        padding: 3px 8px;
        color: #5f9ea0;
        font-family: 'Baloo Bhaijaan 2';
        font-weight: 500;
        cursor: pointer;
        float: right;
    }
    #del_btn{
        border-radius:7px;
        width: fit-content;
        padding: 3px 8px;
        color: red;
        font-family: 'Baloo Bhaijaan 2';
        font-weight: 500;
        cursor: pointer;
        float: right;
    }
    .short{
        width:100px;
    }
    .inline{
        display: inline-block;
        margin-right: 18px;
    }
    .long{
        width: 300px;
    }

    .large{
        height: 120px;
        vertical-align: none;
        max-height: 324px;
    }
</style>