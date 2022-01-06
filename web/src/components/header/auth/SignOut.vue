<template v-else>
    <div class='header_item' v-on:click="signOut()">
      {{title}}
    </div>
</template>

<script>
import { signOutAction } from '../../../actions/auth';

export default {
    name: 'SignOut',
    props: {
        title: String,
    },
    methods:{
    signOut: () => {
      signOutAction()
    }
  }
}
</script>