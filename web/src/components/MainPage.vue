<template>
  <div id="main">
    <Header />
    <div id="main_cont">
      <div class="main_page">
        <h2 class="header">The good example of good header</h2>
        <div class="text_cont">
          This page about my way in programming study.
        </div>
      </div>
    </div>
  </div>
</template>

<script>
import Header from './header/Header.vue'

export default {
  name: 'Main',
  props: {
    msg: String
  },
  components: {
    Header,
  }
}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>

</style>
