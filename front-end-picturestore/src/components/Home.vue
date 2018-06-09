<template>
  <div class="hello">
    <h1>{{ msg }}</h1>
    <h2>Essential Links</h2>
    <ul>
      <li>
        <a
          href="https://vuejs.org"
          target="_blank"
        >
          Core Docs
        </a>
      </li>
      <li>
        <a
          href="https://forum.vuejs.org"
          target="_blank"
        >
          Forum
        </a>
      </li>
      <li>
        <a
          href="https://chat.vuejs.org"
          target="_blank"
        >
          Community Chat
        </a>
      </li>
      <li>
        <a
          href="https://twitter.com/vuejs"
          target="_blank"
        >
          Twitter
        </a>
      </li>
      <br>
      <li>
        <a
          href="http://vuejs-templates.github.io/webpack/"
          target="_blank"
        >
          Docs for This Template
        </a>
      </li>
    </ul>
    <h2>Ecosystem</h2>
    <ul>
      <li>
        <a
          href="http://router.vuejs.org/"
          target="_blank"
        >
          vue-router
        </a>
      </li>
      <li>
        <a
          href="#/signup"
        >
          Sign up
        </a>
      </li>
      <li>
        <a
          href="#/login"
        >
          Login
        </a>
      </li>
      <li v-if="token">
        <button
          v-on:click="logout"
        >
          Logout
        </button>
      </li>
    </ul>
  </div>
</template>

<script>
export default {
  name: 'Home',
  data () {
    return {
      msg: 'Loading',
      token: localStorage.token
    }
  },
  created: function () {
    this.fetchData()
  },
  methods: {
    logout () {
      delete localStorage.token
      this.token = false
      this.msg = 'Logged out'
      this.flash('Logout successful', 'success')
      this.$router.replace(this.$route.query.redirect || '/')
    },
    fetchData () {
      this.$http.get('http://localhost:4000', {headers: {'x-access-token': localStorage.token}})
        .then(response => this.successful(response))
        .catch(() => this.fetchFailed())
    },
    successful (response) {
      this.msg = response.body
    },
    fetchFailed () {
      this.msg = 'Couldn\'t get data'
    }
  }
}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
h1, h2 {
  font-weight: normal;
}
ul {
  list-style-type: none;
  padding: 0;
}
li {
  display: inline-block;
  margin: 0 10px;
}
a {
  color: #42b983;
}
</style>
