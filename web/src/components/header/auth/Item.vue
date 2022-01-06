<template>
<span>
  <router-link class='header_item' tag='div' v-if="url != ''"  :to='url'>{{title}}</router-link>
  <div class='header_item' v-if="url == ''" v-on:click="signOut()">
    {{title}}
  </div>
</span>
  
</template>

<script>
import { signOutAction } from '../../../actions/auth';
export default {
  name: 'AuthItem',
  props: {
    title: String,
    url: String,
  },
  methods:{
    signOut: () => {
      signOutAction()
    }
  }
}
</script>
