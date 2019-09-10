<template>
  <div class="products">
    <div class="container">   
      <div class="intro h-100">
        <div class="row h-100 justify-content-center align-items-center">
          <div class="col-md-6">
              <h3>Products Sell Page</h3>
              <p>
                Change your products sell settings here
              </p>
          </div>
          <div class="col-md-5">
              <img src="./image/products.svg" alt="" width="300" class="img-fluid">
          </div>
        </div>
      </div>
      <hr>
      <h3 class="d-inline-block">Products Sell list</h3>
      <button @click="addExcel()" class="btn btn-primary float-right"><i class="fas fa-file-excel"></i></button>
      <div class="table-responsive">
        <table id="dtBasicExample" class="table table-striped table-bordered table-sm" cellspacing="0" width="100%">
          <thead class="thead-dark">
            <tr>
              <th>Nama</th>
              <th>Gambar</th>
              <th>Harga</th>
              <th>Terjual</th>
              <th>Total</th>
              <th></th>
            </tr>
          </thead>
          <tbody>
            <tr v-for="product in productsSell" :key="product.id">
              <td>
                {{product.name}}
              </td>
              <td>
                <img :src="product.image" alt="" width="100px">
              </td>
              <td>
                {{product.price | currency('Rp. ', 0)}}
              </td>
               <td>
                {{product.jumlah}}
              </td>
               <td>
                {{product.total | currency('Rp. ', 0)}}
              </td>
              <td><button @click="Details(product)" type="button" class="btn btn-primary">Detail</button></td>
            </tr>
          </tbody>
        </table>
      </div>
    </div>
      
      <!-- Modal -->
      <div class="modal fade" id="excel" tabindex="-1" role="dialog" aria-labelledby="productLabel" aria-hidden="true">
        <div class="modal-dialog modal-lg" role="document">
          <div class="modal-content">
            <div class="modal-header">
              <h5 class="modal-title" id="productLabel">Export To Excel</h5>
              <button type="button" class="close" data-dismiss="modal" aria-label="Close">
                <span aria-hidden="true">&times;</span>
              </button>
            </div>
            <div class="modal-body">
              <div class="row">
                <div class="col-md-6">
                  <div class="form-group">
                    <i class="fa fa-calendar"></i>
                    <label for="dateofbirth">Dari Tanggal&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp; : &nbsp;</label>
                    <input v-model="tglawal" type="date" name="dateofbirth" id="dateofbirth">
                  </div>
                  <div class="form-group">
                      <i class="fa fa-calendar"></i>
                      <label for="dateofbirth">Sampai Tanggal : &nbsp;</label>
                      <input v-model="tglakhir" type="date" name="dateofbirth" id="dateofbirth">
                  </div>           
                </div>
                <div class="col-md-5">
                  <h5 class="display-6">Details</h5>
                  <hr>
                  <div>Laporan dari tanggal {{tglawal | formatDate1}} sampai tanggal {{tglakhir | formatDate1}} </div>
                </div>
              </div>
            </div>
            <div class="modal-footer">
              <button type="button" class="btn btn-secondary" data-dismiss="modal">Close</button>
              <button @click="Export" type="button" class="btn btn-primary">Export</button>
              <!-- <download-excel
                  class   = "btn btn-default"
                  :data   = "json_data"
                  :fields = "json_fields"
                  worksheet = "My Worksheet"
                  name    = "filename.xls"
              >
                  Download Excel
              </download-excel> -->
            </div>
          </div>
        </div>
      </div>


      <!-- Modal -->
      <!-- <div class="modal fade" id="detail" tabindex="-1" role="dialog" aria-labelledby="productLabel" aria-hidden="true">
        <div class="modal-dialog modal-lg" role="document">
          <div class="modal-content">
            <div class="modal-header">
              <h5 class="modal-title" id="productLabel">Detail Product</h5>
              <button type="button" class="close" data-dismiss="modal" aria-label="Close">
                <span aria-hidden="true">&times;</span>
              </button>
            </div>
            <div class="modal-body">
              <table class="table-responsive">
                <thead>
                  <tr>
                    <th scope="col">No</th>
                    <th scope="col" >No Pesanan</th>
                    <th scope="col" >Tgl Order</th>
                    <th scope="col" colspan="3">Detail Produk</th>
                    <th>Total Harga</th>
                  </tr>
                </thead>
                <tbody>
                  <tr>
                    <td>1</td>
                    <td>29</td>
                    <td>Tuesday 13, August 2019 16:08 WIB</td>
                    <td>Aceh Gayo Washed</td>
                    <td>8</td>
                    <td>5000</td>
                    <td>80000</td>
                  </tr>
                  <tr>
                    <th scope="row">2</th>
                    <td>Jacob</td>
                    <td>Thornton</td>
                    <td>@fat</td>
                  </tr>
                  <tr>
                    <th scope="row">3</th>
                    <td colspan="3">Larry the Bird</td>
                    <td>@twitter</td>
                  </tr>
                </tbody>
              </table>
            </div>
            <div class="modal-footer">
            </div>
          </div>
        </div>
      </div> -->


      
      <div class="modal fade bd-example-modal-xl" id="detail" tabindex="-1" role="dialog" aria-labelledby="myExtraLargeModalLabel" aria-hidden="true">
        <div class="modal-dialog modal-xl">
          <div class="modal-content">
            <div class="modal-header">
              <h5 class="modal-title" id="productLabel">Detail Product</h5>
              <button type="button" class="close" data-dismiss="modal" aria-label="Close">
                <span aria-hidden="true">&times;</span>
              </button>
            </div>
            <div class="table-responsive  table-hover ">
              <table class="table tablea table-bordered">
                <thead class="dark">
                <tr>
                  <th rowspan="2">No</th>
                  <th rowspan="2">No Pesanan</th>
                  <th rowspan="2">Tgl Order</th>
                  <th colspan="3">Detail Produk</th>
                  <th rowspan="2">Total Harga</th>
                </tr>
                <tr>
                  <th>Nama Produk</th>
                  <th>Jumlah Produk</th>
                  <th>Harga Produk</th>
                </tr>
                </thead>
                <tbody>
                  <tr v-for="(detail, index) in productDetails" :key="detail.id">
                    <th  class="dark" scope="row">{{index+1}}</th>
                    <td>{{detail.id_pesanan}}</td>
                    <td>{{detail.tglorder | formatDate}}</td>
                    <td>{{detail.name}}</td>
                    <td>{{detail.quantity}}</td>
                    <td>{{detail.price | currency('', 0)}}</td>
                    <td>{{detail.total_price | currency('IDR ', 0)}}</td>
                  </tr>
                </tbody>
              </table>
            </div>
          </div>
        </div>
      </div>
  </div>
</template>

<script>
import { VueEditor } from "vue2-editor";
import { mdbDatatable } from 'mdbvue';

export default {
  name: 'products',
  components: {
    VueEditor,
    mdbDatatable
  },
  data () {
    return {
      tglawal:'',
      tglakhir:'',
      // json_fields: {
      //     'Complete name': 'name',
      //     'City': 'city',
      //     'Telephone': 'phone.mobile',
      //     'Telephone 2' : {
      //       field: 'phone.landline',
      //       callback: (value) => {
      //           return `Landline Phone - ${value}`;
      //       }
      //     },
      // },
      // json_data: [
      //   {
      //       'name': 'Tony Peña',
      //       'city': 'New York',
      //       'country': 'United States',
      //       'birthdate': '1978-03-15',
      //       'phone': {
      //           'mobile': '1-541-754-3010',
      //           'landline': '(541) 754-3010'
      //       }
      //   },
      // ]
    }
  },
  created() {
    this.$store.dispatch('retrieveProductsSell')
  },
  watch:{
  
  },
  computed:{
    productsSell(){
      return this.$store.getters.ProductsSell
    },
    productDetails(){
      return this.$store.state.details
    },
  },
  methods:{
    Export(){
      // console.log(this.tglawal)
      this.$store.dispatch('exportExcel', {
        tgl1 : this.tglawal,
        tgl2 : this.tglakhir
      })
      .then(response => {
        this.$router.push({ name: 'productssell' })
        this.reset()
        $('#excel').modal('hide');
        Toast.fire({
            type: 'success',
            title: 'Export  successfully'
        })
      })
    },
    reset(){
      Object.assign(this.$data, this.$options.data.apply(this));
    },
    addExcel(){
      $('#excel').modal('show');
    },
    Details(product){
      $('#detail').modal('show');
      // console.log(product.name)
      this.$store.dispatch('detailsProductSells', {
        name : product.name
      })
    }
  },
}
</script>

<style lang="scss">
.img-wrapp{
  position: relative;
  padding-bottom: 8px;
    display: block;
}
.img-wrapp span.delete-img{
    position: absolute;
    top: -10px;
    left: -1px;
}
.img-wrapp span.delete-img:hover{
  cursor: pointer;
}
.dark{
  background-color: #dab25c;
}
.tablea {
  border: 1px solid black;
  // border-collapse: collapse;
}
</style>


