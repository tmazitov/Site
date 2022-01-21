<template>

    <div>
        <div class="create_item">
            <div class="sub_header">Title</div>
            <input 
            id="title" 
            class="create_field long"
            placeholder="About your order"
            />
        </div>
        <div class="create_item inline">
            <div class="sub_header">Price ( in $ )</div>
            <input id="price" class="create_field short"/>
        </div>
        <div class="create_item">
            <div class="sub_header">Comment</div>
            <textarea 
            id="comment" 
            class="create_field long large"
            placeholder="Help you get the job done with precision"
            />
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
                this.$swal({
                    title: 'Success',
                    text:'New order was successfully created!',
                    icon: "success",
                });
                title.value = ""
                price.value = ""
                comment.value = ""
            }).catch((error) => {
                if ( error.status == 400){
                    this.$swal({
                    title:'Error',
                    text: 'Invalid credentials',
                    icon: "success",
                    });
                } else if (error.status == 500) {
                    this.$swal({
                    title:'Error',
                    text: 'Server error',
                    icon: "success",
                    }); 
                }

            })
        },
    }
}
</script>

<style scoped>
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