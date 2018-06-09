<template>
  <div>
  <div>
    <form class="form-signin" @submit.prevent="login">
      <h2 class="form-signin-heading">Please login</h2>
      <label for="inputEmail" class="sr-only">Email address</label>
      <input v-model="email" type="email" id="inputEmail" class="form-control" placeholder="Email address" required autofocus>
      <label for="inputPassword" class="sr-only">Password</label>
      <input v-model="password" type="password" id="inputPassword" class="form-control" placeholder="Password" required>
      <button class="btn btn-lg btn-primary btn-block" type="submit">Login</button>
    </form>
  </div>
  </div>
</template>

<script>
export default {
  name: 'Signup',
  data () {
    return {
      email: 'cbing@gmail.com',
      password: 'cbing'
    }
  },
  methods: {
    login () {
      var formData = new FormData()
      formData.append('email', this.email)
      formData.append('password', this.password)
      this.$http.post('http://localhost:3000/api/auth/login', formData)
        .then(response => this.loginSuccessful(response))
        .catch(() => this.loginFailed())
    },
    loginSuccessful (res) {
      if (!res.body.token) {
        this.loginFailed()
        return
      }
      this.flash('Login successful', 'success')
      localStorage.token = res.body.token
      this.error = false
      this.$router.replace(this.$route.query.redirect || '/')
    },
    loginFailed () {
      this.flash('Login failed', 'error')
      this.error = 'Login failed!'
      delete localStorage.token
    }
  }
}
</script>

<style lang="css">
</style>
