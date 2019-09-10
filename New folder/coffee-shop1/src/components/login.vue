<template>
 <div class="login">     
        <h5 class="text-center">Login Please</h5>
        <form action="#" @submit.prevent="login">
            <div class="form-group">
              <label for="exampleInputEmail1">Email address</label>
              <input type="email" name="email" v-model="email" class="form-control" id="exampleInputEmail1" aria-describedby="emailHelp" placeholder="Enter email">
            </div>
            <div class="form-group">
              <label for="exampleInputPassword1">Password</label>
              <input type="password" name="password" v-model="password" class="form-control" id="exampleInputPassword1" placeholder="Password">
            </div>
            <div class="form-group">
              <button class="btn btn-primary" type="submit">Login</button>
              <!-- <div v-if="getError">
                Terjadi masalah
              </div> -->
              <hr>
              <div class="alert alert-primary" role="alert" v-if="getError">
                Email dan Password anda salah !
              </div>
            </div>
        </form>
    
 </div>
</template>

<script>
export default {
  name: 'login',
  data () {
    return {
      email: '',
      password: '',
    }
  },
  computed : {
      getError(){
       return  this.$store.getters.getError
      }
  },
  watch : {
      getError(val){
         var that = this
        setTimeout(function(){
          that.$store.commit('setError',0)
        }, 5000)
      }
  },
  methods:{
      login() {
      this.$store.dispatch('retrieveToken', {
        email: this.email,
        password: this.password,
      })
        .then(response => {
        $('#login').modal('hide')
          this.$router.push({ name: 'overview' })
          Toast.fire({
              type: 'success',
              title: 'Login  successfully'
            })
        })
        
    }
  }
}
</script>