<template>
    <div>
        <div v-for="order in orders" :key="order.UUID" class="order"> 
            <div class="order_header">{{titleFormat(order.Title)}}</div>
            <div>Writer: {{order.Writer}}</div>
            <div>Status: {{order.Status}}</div>
            <div>Price: {{order.Price}} $</div>
            <div class="date">{{dateFormat(order.Date)}}</div>
            <div class="complite" v-on:click="complite(order.UUID)"></div>
            <div class="update" v-on:click="update(order.UUID)"></div> 
        </div>
    </div>
</template>

<script>

import client from '../../client/client'

import {readValue } from '../../actions/jwt'
import {refreshTokens} from '../../actions/auth'

export default {
    name: "OrderList",
    beforeMount() {
        client
            .get("/order/list").then((response) => {
                this.orders = response.data.orders
            })
        this.intervalId = setInterval(() => {
            client
            .get("/order/list").then((response) => {
                this.orders = response.data.orders
            }).catch(() => {
                if ( readValue() === ""){
                    this.$router.push('auth')
                }
                refreshTokens()
            })
        }, 15000)
    },
    destroyed() {

        clearInterval(this.intervalId)
    },
    data() {
        return {
            orders: [],
            intervalId: 0,
        }
    },
    methods: {
        titleFormat(title){
            if (title.length > 18) {
                return title.slice(0,18) + "..."
            } else {
                return title
            }
            
        },
        dateFormat(value){
            let d = new Date(value * 1000)
            return d.toLocaleString("ru")
        },
        complite(uuid){
            client
            .put("/order/complite", {
                UUID: uuid,
            })
        },
        update(uuid){
            client
            .put("/order/update", {
                UUID: uuid,
                Title: "Bebra",
            })
        }
    },
}
</script>

<style scoped>
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
</style>