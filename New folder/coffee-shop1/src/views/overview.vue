<template>
  <div class="overview">
      <div class="container h-100">
          <div class="intro h-100">
            <div class="row h-100 justify-content-center align-items-center">
              <div class="col-md-5">
                 <h3>Overview Page</h3>
                  <div class="animated-shadow-quote">
                    <blockquote>
                      <p class="quotation">Dimana Rasa dan Asa Bertemu</p>
                      <cite>꧕r•asa Café</cite>
                    </blockquote>
                  </div>
              </div>
              <div class="col-md-6">
                  <img src="./image/overview.svg" width="300" alt="" class="img-fluid">
              </div>
            </div>
            <br>
            <div class="row">
              <div class="col-md-6 col-xl-3">
                <div class="card bg-c-blue order-card hover">
                  <div class="card-body">
                      <h5 class="text-white">Product Sell</h5>
                      <h2 class="text-right text-white"><i class="fas fa-cart-arrow-down float-left" style="font-size: 1.5em;"></i><span class="feather">{{this.produksell}}</span></h2>
                      <span class="float-right">Jumlah Produk Terjual</span>
                  </div>
                </div>
              </div>
              <div class="col-md-6 col-xl-3">
                <div class="card bg-c-green order-card hover">
                    <div class="card-body">
                        <h5 class="text-white">User</h5>
                        <h2 class="text-right text-white"><i class="far fa-user float-left" style="font-size: 1.5em;"></i><span class="feather">{{this.user}}</span></h2>
                        <span class="float-right">Jumlah User Aktif</span>
                    </div>
                </div>
            </div>
            <div class="col-md-6 col-xl-3">
                <div class="card bg-c-yellow order-card hover">
                    <div class="card-body">
                        <h5 class="text-white">Orders</h5>
                        <h2 class="text-right text-white"><i class="fa fa-chart-line float-left" style="font-size: 1.5em;"></i><span class="feather">{{this.orders}}</span></h2>
                        <span class="float-right">Jumlah Order User</span>
                    </div>
                </div>
            </div>
            <div class="col-md-6 col-xl-3">
                <div class="card bg-c-red order-card hover">
                    <div class="card-body">
                        <h5 class="text-white">Comments</h5>
                        <h2 class="text-right text-white"><i class="fa fa-comments float-left" style="font-size: 1.5em;"></i><span class="feather">{{this.comments}}</span></h2>
                        <span class="float-right">Komentar User</span>
                    </div>
                </div>
            </div>

            </div>
          </div>
      </div>
  </div>
</template>

<script>
import '@/assets/overview.css';

export default {
  name: "overview",
  props: {
    msg: String
  },
  data(){
    return{
      produksell:'',
      user:'',
      orders:'',
      comments:''
    }
  },
  watch:{
    productsSell(val){
      this.produksell = val.length
    },
    profileFiltered(val){
      this.user = val.length
    },
    ordersFiltered(val){
      this.orders = val.length
    },
    commentsFiltered(val){
      this.comments = val.length
    }
  },
  created() {
    this.$store.dispatch('retrieveProductsSell')
    this.$store.dispatch('retrieveProfile')
    this.$store.dispatch('retrieveOrder')
    this.$store.dispatch('retrieveComments')
  },
  computed:{
    productsSell(){
      return this.$store.getters.ProductsSell
    },
    profileFiltered(){
      return this.$store.getters.ProfileFilter
    },
    ordersFiltered(){
      return this.$store.getters.OrdersFilter
    },
    commentsFiltered(){
      return this.$store.state.comments
    },
  }
};
</script>

<style scoped lang="scss">
.animated-shadow-quote {
  display: inline-block;
  margin: 1em;
  max-width: 20em;
  position: relative;
}
.animated-shadow-quote blockquote {
  animation: shadows 2s linear infinite alternate;
  display: inline-block;
  margin: 0;
  padding: 1em;
}
.animated-shadow-quote blockquote cite {
  display: block;
  font-style: italic;
  text-align: right;
}
.animated-shadow-quote blockquote cite:before {
  content: "- ";
}

@keyframes shadows {
  0% {
    box-shadow: 0 2px 4px -2px rgba(0, 0, 0, 0.25);
    transform: scale(0.95);
  }
  100% {
    box-shadow: 0 0 4px 2px rgba(0, 0, 0, 0.25);
    transform: scale(1);
  }
}
.square-brackets-quote {
  display: inline-block;
  font-family: sans-serif;
  margin: 1em;
  max-width: 20em;
}
.square-brackets-quote blockquote {
  border: solid 1em #ccc;
  display: inline-block;
  margin: 0;
  padding: 1em;
  position: relative;
}
.square-brackets-quote blockquote:before {
  background-color: #fff;
  bottom: -1em;
  content: "";
  left: 2em;
  position: absolute;
  right: 2em;
  top: -1em;
}
.square-brackets-quote blockquote cite {
  color: #757575;
  display: block;
  font-size: small;
  font-style: normal;
  text-align: right;
  text-transform: uppercase;
}
.square-brackets-quote blockquote > * {
  position: relative;
  z-index: 1;
}
.giant-quotes-quote {
  display: inline-block;
  font-family: sans-serif;
  margin: 1em;
  max-width: 20em;
}
.giant-quotes-quote blockquote {
  display: inline-block;
  margin: 0;
  padding: 1em;
  position: relative;
}
.giant-quotes-quote blockquote:before {
  color: #f90;
  content: "\201C";
  font-size: 4em;
  left: 0;
  line-height: .75em;
  position: absolute;
  top: 0;
  transform: scale(2, 8);
  transform-origin: top;
}
.giant-quotes-quote blockquote:after {
  color: #f90;
  content: "\201D";
  font-size: 4em;
  line-height: .75em;
  position: absolute;
  right: 0;
  top: 0;
  transform: scale(2, 8);
  transform-origin: top;
}
.giant-quotes-quote blockquote > * {
  padding: 0 2em;
}
.giant-quotes-quote blockquote cite {
  color: #c60;
  display: block;
  font-style: normal;
  text-align: right;
  text-transform: uppercase;
}
</style>