<template>
  <div id="app">
    <div class="container">
      <h1>user details : {{ msg }}</h1><!-- top bar for user details -->
      <flash-message class="myCustomClass"></flash-message>
      <router-view></router-view>
    </div>
  </div>
</template>

<script>
export default {
  name: 'App',
  data () {
    return {
      msg: 'Loading'
    }
  }
}
</script>

<style>
</style>
