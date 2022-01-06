<template>
  <router-link class='header_item' tag='div'   :to='url'>{{title}}</router-link>
</template>

<script>

export default {
  name: 'SignIn',
  props: {
    title: String,
    url: String,
  },
}
</script>
