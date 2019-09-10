<template>
  <div class="auth">
       <!-- Modal -->
        <div class="modal fade" id="login" tabindex="-1" role="dialog" aria-labelledby="loginTitle" aria-hidden="true">
        <div class="modal-dialog modal-dialog-centered" role="document">
            <div class="modal-content">

                <div class="modal-body">

                  <ul class="nav nav-fill nav-pills mb-3" id="pills-tab" role="tablist">
                      <li class="nav-item">
                          <a class="nav-link active" id="pills-home-tab" data-toggle="pill" href="#pills-login" role="tab" aria-controls="pills-login" aria-selected="true">Login</a>
                      </li>
                      <li class="nav-item">
                          <a class="nav-link" id="pills-register-tab" data-toggle="pill" href="#pills-register" role="tab" aria-controls="pills-register" aria-selected="false">Signup</a>
                      </li>
                  </ul>

                  <div class="tab-content" id="pills-tabContent">
                    <div class="tab-pane fade show active" id="pills-login" role="tabpanel" aria-labelledby="pills-login-tab">
                      <login></login>
                    </div>
                      
                    <div class="tab-pane fade" id="pills-register" role="tabpanel" aria-labelledby="pills-register-tab">
                      <register></register>
                    </div>
                  </div>
                    
                </div>
           
            </div>
        </div>
        </div>
  </div>
</template>

<script>
import Login from './login'
import Register from './register'
export default {
  name: 'auth',
  data () {
    return {
    }
  },
  components:{
    Login,Register
  },
  methods:{
      
  },
  mounted(){
    var url_string = window.location.href
    var url = new URL(url_string);
    var c = url.searchParams.get("error");
    // console.log(c == "true")
    if(c == "true") {
      // console.log(document.getElementsByClassName("auth")[0])
      $('#login').modal('show');      
    }
  },
}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>

</style>
