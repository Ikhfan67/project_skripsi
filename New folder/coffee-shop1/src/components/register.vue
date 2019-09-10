<template>
 <div class="register">
    <h5 class="text-center">Create New Account</h5>
    <form action="#" @submit.prevent="register">
        <div class="form-group">
            <label for="name">Your name</label>
            <input type="text" v-model="name" class="form-control" id="name" placeholder="Your name">
        </div>

        <div class="form-group">
            <label for="email">Email address</label>
            <input type="email"  v-model="email" class="form-control" id="email" aria-describedby="emailHelp" placeholder="Enter email">
        </div>
        <div class="form-group">
            <label for="password">Password</label>
            <input type="password" v-model="password" class="form-control" id="password" placeholder="Password">
        </div>

        <div class="form-group">
            <button class="btn btn-primary" type="submit">Signup</button>
        </div>
    </form>                         
 </div>
</template>

<script>
export default {
  name: 'login',
  data () {
    return {
      name: '',
      email: '',
      password: '',
    }
  },
  methods:{
      register() {
      this.$store.dispatch('register', {
        name: this.name,
        email: this.email,
        password: this.password,
      })
        .then(response => {
        $('#login').modal('hide')
          this.$router.push({ name: 'home' })
        })
    }
  },
}
</script>