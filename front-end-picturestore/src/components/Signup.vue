<template>
  <div>
  <div>
    <form class="form-signin" @submit.prevent="login">
      <h2 class="form-signin-heading">Please sign up</h2>
      <label for="inputName" class="sr-only">Name</label>
      <input v-model="name" type="name" id="inputName" class="form-control" placeholder="Chandler Bing" required>
      <label for="inputEmail" class="sr-only">Email address</label>
      <input v-model="email" type="email" id="inputEmail" class="form-control" placeholder="Email address" required autofocus>
      <label for="inputPassword" class="sr-only">Password</label>
      <input v-model="password1" type="password" id="inputPassword" class="form-control" placeholder="Password" required>
      <label for="inputPasswordConfirm" class="sr-only">Password Confirm</label>
      <input v-model="password2" type="password" id="inputPasswordConfirm" class="form-control" placeholder="Password" required>
      <button class="btn btn-lg btn-primary btn-block" type="submit">Sign up</button>
    </form>
    <ul>
      <li>
        <a
          href="#/"
        >
          Home
        </a>
      </li>
      <li>
        <a
          href="#/login"
        >
          Login
        </a>
      </li>
    </ul>
  </div>
  </div>
</template>

<script>
export default {
  name: 'Signup',
  data () {
    return {
      name: 'chandler',
      email: 'cbing@gmail.com',
      password1: 'cbing',
      password2: 'cbing'
    }
  },
  methods: {
    login () {
      var formData = new FormData()
      formData.append('name', this.name)
      formData.append('email', this.email)
      formData.append('password1', this.password1)
      formData.append('password2', this.password2)
      this.$http.post('http://localhost:3000/api/auth/register', formData)
        .then(response => this.loginSuccessful(response))
        .catch(() => this.loginFailed())
    },
    loginSuccessful (res) {
      if (!res.body.token) {
        this.loginFailed()
        return
      }
      this.flash('Signup successful', 'success')
      localStorage.token = res.body.token
      this.error = false
      this.$router.replace(this.$route.query.redirect || '/login')
    },
    loginFailed () {
      this.flash('Signup failed', 'error')
      this.error = 'Signup failed!'
      delete localStorage.token
    }
  }
}
</script>

<style lang="css">
</style>
