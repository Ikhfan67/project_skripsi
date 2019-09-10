<template>
 <div class="hello">
    <nav class="navbar custom-nav fixed-top navbar-expand-lg navbar-light">
     <div class="container">
        <b-img slot="aside" src="http://localhost:5000/gambar_bank/logo.png" width="64" alt="placeholder"></b-img> 
        <router-link class="navbar-brand" to="/">Cafe 3R.ASA</router-link>
        <button class="navbar-toggler" type="button" data-toggle="collapse" data-target="#navbarSupportedContent" aria-controls="navbarSupportedContent" aria-expanded="false" aria-label="Toggle navigation">
          <span class="navbar-toggler-icon"></span>
        </button>

        <div class="collapse navbar-collapse" id="navbarSupportedContent">
          <ul class="navbar-nav mr-auto">
            <li class="nav-item">
              <!-- <router-link to="/#homehero" class="nav-link">Home</router-link> -->
              
                <a href="/#homehero" class="nav-link">Home</a>
              
            </li>
            <li class="nav-item">
              <!-- <router-link :to="{name: 'home', hash:'#products'}" class="nav-link">Products</router-link> -->
              
              <a href="/#products" class="nav-link">Products</a>
              
            </li>
            <li class="nav-item">
              <!-- <router-link to="/about" class="nav-link" href="#">About</router-link> -->
              <a href="/#sales" class="nav-link" enter-active-class="animated jackInTheBox">About</a>
            </li>
            <li class="nav-item">
              <router-link to="/carabeli" class="nav-link" href="#">Cara Membeli</router-link>
            </li>
          </ul>
           
          	 <section v-if="!loggedIn">
						  <form class="form-inline my-2 my-lg-0">
                <!-- <input class="form-control mr-sm-2" type="search" placeholder="Search" aria-label="Search"> -->
                <a class="btn btn-outline-success my-2 my-sm-0" data-toggle="modal" data-target="#login">Login</a>
              </form>
						 </section>
						 <section v-else>
                  <p style="display:inline">Hello, {{loginName}}</p> 
                  <router-link to="/profile" href="#" class="btn btn-outline-info border-0 mx-2 my-2 my-sm-0">
                    <i class="fas fa-user-edit"></i>
                  </router-link>
                  <a class="btn btn-outline-info border-0 mx-2 my-2 my-sm-0" data-toggle="modal" data-target="#miniCart">
                    <i class="fas fa-cart-plus">
                      
                      {{carts.length}}
                       
                    </i>
                  </a>
                  <button @click="logout()" class="btn btn-info"> Log Out</button>
						 </section>
        </div>

     </div>

    </nav>
 </div>   
</template>

<script>
export default {
    name:"navbar",
    data(){
      return{
        
      }
    },
    props: {
        msg: String
    },
    components:{},
    computed:{
      loginName(){
        return localStorage.getItem('name')
      },  
      loggedIn(){
        return this.$store.getters.loggedIn
      },
      carts(){
        return this.$store.state.cart
      }
    },
    methods : {
        logout(){
          this.$store.dispatch('destroyToken')
          .then(response => {
          this.$router.push({ name: 'home' })
          })
        },
    },
  // created(){
  //   var urlString = window.location.href 
  //   var url = new URL(urlString);
  //   var c = url.searchParams.get("error");
  //   if(c){
  //     console.log($('#login'))
  //       $('#login').modal('show');
  //   }
    
  // }
}


</script>

<style scoped lang="scss">
  @media (min-width: 992px) { 
    
    .navbar.custom-nav{
      padding-top:16px;
      padding-bottom:16px;
      background-color: rgb(250, 248, 231) !important;
      // background: rgb(250, 248, 231);
    }
  }
  .router-anim-enter-active{
    animation: coming 1s;
    animation-delay: .5s;
    opacity: 0;
  }
  .router-anim-leave-active{
    animation: going 1s;
  }
  @keyframes going{
    from{
      transform: translateX(0);
    }
    to{
      transform: translateX(-50px);
      opacity: 0;
    }
  }
  @keyframes coming{
    from{
      transform: translateX(-50px);
      opacity: 0;
    }
    to{
      transform: translateX(0px);
      opacity: 1;
    }
  }
</style>

