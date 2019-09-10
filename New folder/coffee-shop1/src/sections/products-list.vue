<template>
    <div class="products" id="products">
        <div class="container">
          <div class="text-center p-5">
            <div id="coffee-shop">
                <div class="coffee-shop-intro wow flipInX">
                    <div class="border">
                        <div class="example">
                          <h1>Buy The Best Of Indonesian Coffee</h1>
                        </div>
                        <div>
                            <p>Konsep dan misi kami adalah memajukan kopi Indonesia<br/>baik di pasar nasional maupun internasional. Mari coba dan nikmati kopi Indonesia.</p>
                            <p>Our concept and mission is to promote Indonesian coffee<br/>both in the local and overseas markets. Come, try and enjoy Indonesian coffee.</p>
                        </div>
                    </div>
                </div>
                <br>
                <div class="row">
                    <div class="col-md-4" v-for="product in productsFiltered" :key="product.id">
                      <div class="">
                        <div id="coffee-wrapper" class="top group">
                            <div id="container">
                                <div class="product-item1 col wow fadeIn">
                                    <img :src="product.image" alt="">
                                    <div class="product-info1 absolutecenter" >
                                        <h4>{{product.name}}</h4>
                                        <div class="ilang" v-html="product.description"></div>
                                        <h4>{{ product.price | currency('IDR ', 0)}} (250gr)</h4>
                                        <add-to-cart
                                            :p-id="product.id"
                                            :price="product.price"
                                            :name="product.name"
                                            :description="product.description"
                                            :image="product.image">
                                        </add-to-cart>
                                    </div>
                                </div>
                            </div>
                        </div>
                      </div>
                    </div>
                </div>
            </div>
          </div>
        </div>
    </div>
</template>

<script>
export default {
  name: "products-list",
  props: {
    msg: String
  },
  created() {
    this.$store.dispatch('retrieveProductsHome')
  },
  watch:{
    // productsFiltered(val){
    //     console.log(val)
    // }
  },
  computed:{
    productsFiltered(){
        return this.$store.getters.ProductsFilter
    },
  },
  methods: {
    getImage(images) {
        return images[0];
    }
  }
};
</script>

<style scoped lang="scss">
.products{
    background:  url(./image/bg_produk4-01.png) ;
}
</style>