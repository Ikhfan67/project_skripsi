<template>
  <div class="admin">
      <a id="show-sidebar" @click="CloseMenu" class="letak btn btn-sm btn-dark" href="#">
            <i class="fas fa-bars"></i>
        </a>
    <div class="page-wrapper default-theme sidebar-bg bg1 toggled">
        
        <nav id="sidebar" class="sidebar-wrapper">
            <div class="sidebar-content">
                <!-- sidebar-brand  -->
                <div class="sidebar-item sidebar-brand">
                    <a href="#">Coffee Shop</a>
                    <div id="close-sidebar" @click="CloseMenu">
                        <i class="fas fa-times"></i>
                    </div>
                </div>
                <!-- sidebar-header  -->
                <div class="sidebar-item sidebar-header d-flex flex-nowrap">
                    <div class="user-pic">
                        <img class="img-responsive img-rounded" src="./image/user.png" alt="User picture">
                    </div>
                    <div class="user-info">
                        <span class="user-name">
                            <strong>{{name}}</strong>
                        </span>
                        <span class="user-role">{{email}}</span>
                        <span class="user-status">
                            <i class="fa fa-circle"></i>
                            <span>Online</span>
                        </span>
                    </div>
                </div>
                <!-- sidebar-search  -->
                <!-- <div class="sidebar-item sidebar-search">
                    <div>
                        <div class="input-group">
                            <input type="text" class="form-control search-menu" placeholder="Search...">
                            <div class="input-group-append">
                                <span class="input-group-text">
                                    <i class="fa fa-search" aria-hidden="true"></i>
                                </span>
                            </div>
                        </div>
                    </div>
                </div> -->
                <!-- sidebar-menu  -->
                <div class=" sidebar-item sidebar-menu">
                    <ul>
                        <li class="header-menu">
                            <span>General</span>
                        </li>
                        <li class="sidebar">
                            <router-link :to="{ name:'overview' }">
                                <i class="fa fa-tachometer-alt"></i>
                                <span class="">Overview</span>
                            </router-link>
                            <!-- <div class="sidebar-submenu">
                                <ul>
                                    <li>
                                        <a href="#">Dashboard 1
                                            <span class="badge badge-pill badge-success">Pro</span>
                                        </a>
                                    </li>
                                </ul>
                            </div> -->
                        </li>
                        <li class="sidebar">
                            <router-link :to="{ name:'products' }">
                                <i class="fa fa-shopping-cart"></i>
                                <span class="">Products</span>
                            </router-link>
                        </li>
                        <li class="sidebar">
                            <router-link :to="{ name:'productssell' }">
                                <i class="fas fa-cart-arrow-down"></i>
                                <span class="">Products <span class="badge badge-pill badge-success">Sell</span></span>
                            </router-link>
                        </li>
                        <li class="sidebar">
                            <router-link :to="{ name:'user' }">
                                <i class="far fa-user"></i>
                                <span class="">User</span>
                            </router-link>
                        </li>
                        <li class="sidebar">
                            <router-link :to="{ name:'orders' }">
                                <i class="fa fa-chart-line"></i>
                                <span class="">Orders</span>
                            </router-link>
                        </li>
                        <li class="sidebar">
                            <router-link :to="{ name:'comment' }">
                                <i class="fa fa-comments"></i>
                                <span class="">Comments</span>
                            </router-link>
                        </li>
                    </ul>
                </div>
                <!-- sidebar-menu  -->
            </div>
            <!-- sidebar-footer  -->
            <div class="sidebar-footer">
                <div>
                    <a href="#" @click="logout()">
                        <i class="fa fa-power-off"> Logout</i>
                        <span class="badge-sonar"></span>
                    </a>
                </div>
            </div>
        </nav>
        <!-- page-content" -->
        <main class="page-content">
            <!-- <h1>Selamat datang di halaman Admin</h1> -->
            <router-view/>
        </main>
    </div>
    <!-- page-wrapper -->
  </div>
</template>

<script>
export default {
    name: "admin",
    data(){
      return{
          email:null,
          name:null
      }
    },
    computed:{
      emailFiltered(){
      return this.$store.getters.loggedEmail
      },
    },
    created(){
        let user = this.$store.state.email
        this.email = user
        let nama = this.$store.state.name
        this.name = nama
    },
    methods:{
      CloseMenu(){
        $(".page-wrapper").toggleClass("toggled");
      },
      CloseMenuPin(){
          if ($(".page-wrapper").hasClass("pinned")) {
            // unpin sidebar when hovered
            $(".page-wrapper").removeClass("pinned");
            $("#sidebar").unbind( "hover");
            } else {
                $(".page-wrapper").addClass("pinned");
                $("#sidebar").hover(
                    function () {
                        console.log("mouseenter");
                        $(".page-wrapper").addClass("sidebar-hovered");
                    },
                    function () {
                        console.log("mouseout");
                        $(".page-wrapper").removeClass("sidebar-hovered");
                    }
                )

                }
      },
      DropDown(){
          $(".sidebar-submenu").slideUp(200);
        if ($(this).parent().hasClass("active")) {
            $(".sidebar-dropdown").removeClass("active");
            $(this).parent().removeClass("active");
        } else {
            $(".sidebar-dropdown").removeClass("active");
            $(this).next(".sidebar-submenu").slideDown(200);
            $(this).parent().addClass("active");
        }
      },
      logout(){
          this.$store.dispatch('destroyToken')
          .then(response => {
          this.$router.push({ name: 'home' })
      })
      },
    }
}
</script>

<style lang="scss" scoped>
.letak{
    position: relative;
}
</style>

