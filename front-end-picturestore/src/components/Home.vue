<template>
  <div class="hello">
    <ul id="images">
      <li v-if="$parent.msg.name">
        <input type="file" @change="onFileChanged">
        <button @click="onUpload">Upload!</button>
      </li>
      <li class="list-element" v-for="pic in pictures" v-bind:key="pic.imageid">
        <img v-bind:src="pic.url" />
      </li>
    </ul>
    <ul>
      <li v-if="!$parent.msg.name">
        <a
          href="#/signup"
        >
          Sign up
        </a>
      </li>
      <li v-if="!$parent.msg.name">
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
import axios from 'axios'
export default {
  name: 'Home',
  axios: 'axios',
  data () {
    return {
      selectedFile: null,
      pictures: 'Loading pictures',
      token: localStorage.token
    }
  },
  created: function () {
    this.fetchUserData()
    this.fetchPictures()
  },
  methods: {
    logout () {
      delete localStorage.token
      this.token = false
      this.$parent.msg = 'Logged out'
      // this.pictures = 'Logged out'
      this.flash('Logout successful', 'success')
      this.$router.replace(this.$route.query.redirect || '/')
    },
    fetchUserData () {
      this.$http.get('http://localhost:4000', {headers: {'x-access-token': localStorage.token}})
        .then(response => this.fetchUserDataSuccessful(response))
        .catch(() => this.fetchUserDataFailed())
    },
    fetchUserDataSuccessful (response) {
      this.$parent.msg = response.body
    },
    fetchUserDataFailed () {
      this.$parent.msg = 'Couldn\'t get user data from the data service'
    },
    fetchPictures () {
      this.$http.get('http://localhost:5000', {headers: {'x-access-token': localStorage.token}})
        .then(response => this.fetchPicturesSuccessful(response))
        .catch(() => this.fetchPicturesFailed())
    },
    fetchPicturesSuccessful (response) {
      this.pictures = response.body
    },
    fetchPicturesFailed () {
      this.pictures = 'Couldn\'t get pictures from the picture service'
    },
    onFileChanged (event) {
      this.selectedFile = event.target.files[0]
    },
    onUpload () {
      const formData = new FormData()
      formData.append('uploadfile', this.selectedFile, this.selectedFile.name)
      axios.post('http://localhost:5000', formData)
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
