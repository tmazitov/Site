<template>
    <div id="adminpanel">
        <h2 class="header">Userlist</h2>
        <UserList/>
        <hr class="borderline" size=1.5 color="#9BC3C4"/>
        <h2 class="header">Collaborator register</h2>
        <div class="sub_header">Username</div>
        <input id="username" class="create_field"/>
        <div class="sub_header">Role</div>
        <div id="role_cont">
            <div class="role_item active" id="worker" v-on:click="selectRole('worker')">
                Worker
            </div>
            <div class="role_item" id="manager" v-on:click="selectRole('manager')">
                Manager
            </div>
        </div>
        <div id="create_btn_cont">
            <button id="create_btn" v-on:click="regWorker()">
            Register
            </button>  
        </div>
    </div>
</template>
<script>
import UserList from '../userlist/Table.vue'

import client from '../../client/client'

var roles = {
    worker : 'Worker',
    manager: 'Manager',
}

export default {
    name: 'AdminPanel',
    components: {
        UserList
    },
    data() {
        return {
            currentRole: "worker",
        }
    },
    methods: {
        regWorker(){
            let worker = document.getElementById('username')
            client
            .post('/user/upgrade', {
                Username: worker.value,
                Role: roles[this.currentRole],
            }).then(() => {
                worker.value = ""
                this.$swal({
                    title:'Success',
                    text: 'New collaborator was registered',
                    icon: "success",
                    });
            }).catch((error) => {
                if (error.status == 400){
                    this.$swal({
                    title:'Error',
                    text: 'Invalid credentials',
                    icon: "error",
                    });
                }
                if (error.status == 500){
                    this.$swal({
                    title:'Error',
                    text: 'Server error',
                    icon: "error",
                    });
                }
                
            })
        },
        selectRole(role){
            let currentElem = document.getElementById(this.currentRole)
            let selectedElem = document.getElementById(role)

            currentElem.className = selectedElem.className
            selectedElem.className += " active"

            this.currentRole = role
        }
    }
}
</script>

<style scoped>
#role_cont{
    border-radius: 7px;
    width: fit-content;
}
.role_item{
    border-radius: inherit;
    padding: 3px 16px;
    display: inline-block;
    font-family: "Baloo Bhaijaan 2";
    cursor: pointer;
}

.active{
    background: #9BC3C4;
}

.borderline{
    margin: 25px 0px;
}
</style>