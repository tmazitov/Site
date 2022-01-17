<template>

    <div>
        <div class="create_item">
            <div class="sub_header">Title</div>
            <input id="title" class="create_field long"/>
        </div>
        <div class="create_item inline">
            <div class="sub_header">Price</div>
            <input id="price" class="create_field short"/>
        </div>
        <div class="create_item">
            <div class="sub_header">Comment</div>
            <textarea id="comment" class="create_field long large"/>
        </div>
        <div id="create_btn_cont">
            <button id="create_btn" v-on:click="create()">
            Create
            </button>  
        </div>

    </div>
  
</template>

<script>

import client from '../../client/client'

 
export default {
    name: "CreateOrder",
    methods: {
        create(){

            var title = document.getElementById("title")
            var price = document.getElementById("price")
            var comment = document.getElementById("comment")

            client
            .post("/order/create",{
                Title: title.value,
                Price: Number(price.value),
                Comment: comment.value 
            }).then(() => {
                title.value = ""
                price.value = ""
                comment.value = ""
                // TODO: add the success create sign
            })
        },
    }
}
</script>

<style scoped>
    .create_field{
        margin-top: 3px;
        border: 2px solid #9BC3C4;
        border-radius: 7px;
        height: 20px;
        padding: 3px 8px;
    }
    .create_field:focus{
        outline: none !important;
        border-color: wheat;
    }
    .create_item{
        margin-top: 14px;
    }
    #create_btn_cont{
        margin: 18px 0px;
    }
    #create_btn{
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
    #create_btn:hover{
    background: #588C8E;
    color:wheat;
    }
    #create_btn:active{
    background: #507A7C;
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