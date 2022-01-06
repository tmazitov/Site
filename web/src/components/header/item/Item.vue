<template>
  <router-link tag='div' class='header_item' :to='url'>{{title}}</router-link> 
</template>

<script>
export default {
  name: 'Item',
  props: {
    title: String,
    url: String,
  },
  methods:{
    goTo: function(url){
      window.location.href = url;
    }
  }
}
</script>
