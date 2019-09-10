<template>
  <div class="add-to-cart">

    <a class="tombol1" @click="addToCart">BUY</a>
    
  </div>
</template>

<script>
export default {
  name: "add-to-cart",
  props: {
    name: String,
    price: Number,
    pId: Number,
    image:String,
    description: String,
  },
  data(){
    return {
      item : {
        productName: this.name,
        productDescription: this.description,
        productPrice: this.price,
        productImage: this.image,
        id_produk: this.pId,
        quantity: 1,
      },
      fix:{
        id_produk: this.pId,
        quantity: 1,
        productPrice: this.price,
      }
    }
  },
  methods:{
    addToCart(){
    $('#miniCart').modal('show');
    this.$store.commit('addToCart', this.item)
    this.$store.commit('fix', this.fix)
    },
  }
  
};
</script>

<style scoped lang="scss">
</style>