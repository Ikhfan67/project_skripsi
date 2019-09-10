<template>
    <div class="checkout">
        <Navbar></Navbar>
        <section class="ftco-section ftco-cart">
			<div class="container">
				<div class="row">
                    <div class="bawah">
                        <div class=""><i class="fa fa-calendar"></i> {{tanggal}}</div>
                    </div>
                    <div class="col-md-12 ftco-animate">
                        <div class="cart-list">
                            <table class="table">
                                <thead class="thead-primary">
                                <tr class="text-center">
                                    <th>&nbsp;</th>
                                    <th>&nbsp;</th>
                                    <th>Product</th>
                                    <th>Price</th>
                                    <th>Quantity</th>
                                    <th>Total</th>
                                </tr>
                                </thead>
                                <tbody>
                                <tr class="text-center" v-for="item in this.$store.state.cart" :key="item.id">
                                    <td class="product-remove"><a class="far fa-trash-alt" @click="$store.commit('removeFromCart',item)"></a></td>
                                    
                                    <td class="image-prod"><div class="img"><img :src="item.productImage" alt style="width:100px"></div></td>
                                    
                                    <td class="product-name">
                                        <h3>{{item.productName}}</h3>
                                        <p v-html="item.productDescription"></p>
                                    </td>
                                    
                                    <td class="price">{{item.productPrice | currency('IDR ', 0)}}</td>
                                    
                                    <td class="quantity">
                                        <div class="input-group mb-3">
                                            <span class="input-group-btn">
                                                <button
                                                    type="button"
                                                    class="btn btn-danger"
                                                    @click="decreaseQty(item.id_produk)"
                                                >
                                                <i class="fa fa-minus"></i>
                                                </button>
                                            </span>
                                            <input
                                            type="text"
                                            :value="item.quantity"
                                            class="form-control input-number"
                                            >
                                            <span class="input-group-btn">
                                                <button
                                                    type="button"
                                                    class="btn btn-success"
                                                    @click="increaseQty(item.id_produk)"
                                                >
                                                <i class="fa fa-plus"></i>
                                                </button>
                                            </span>
                                        </div>
                                    </td>
                                    <td class="total">{{item.quantity * item.productPrice | currency('IDR ', 0)}}</td>
                                </tr>
                                </tbody>
                            </table>
                        </div>
                    </div>
    		    </div>
                
                    <div class="col col-lg-3 mt-5 cart-wrap ftco-animate">
                        <div class="cart-total mb-3">
                            <h3>Cart Totals</h3>
                            <p class="d-flex">
                                <span>Subtotal</span>
                                <span>{{totalharga | currency('', 0)}}</span>
                            </p>
                            <p class="d-flex">
                                <span>Quantity</span>
                                <span>{{totalquantity}}</span>
                            </p>
                            <p class="d-flex">
                                <span>Discount</span>
                                <span>{{discount}} %</span>
                            </p>
                            <hr>
                            <p class="d-flex total-price">
                                <span>Total</span>
                                <span>{{total | currency('IDR ', 0)}}</span>
                            </p>
                        </div>
                        <section v-if="total">
                            <button @click="order" class="btn-primary1">Proceed to Order</button>
                        </section>
                            
                        <section v-else>
                            <b-alert show variant="danger"><a href="#" class="alert-link">Tidak bisa Order. Anda harus memilih produk yang dibeli !</a></b-alert>
                        </section>
                        <br>
                            <div id="modal02" class="w3-modal w3-animate-opacity">
                                <div class="w3-modal-content">
                                    <header class="w3-container w3-teal"> 
                                    <span onclick="document.getElementById('modal02').style.display='none'" 
                                    class="w3-button w3-display-topright">&times;</span>
                                    <div class="tulisan">Upload Bukti Transfer</div>
                                    </header>
                                    <br>
                                    <div class="w3-container">
                                        <div class="row">
                                            <div class="col-md-6">
                                                <div class="form-group">
                                                    <div class="input-group">
                                                        <div class="input-group-prepend">
                                                            Pilih No Pesanan :&nbsp;
                                                            <select v-model="tglorderupload">
                                                                <option v-for="item in KatalogFiltered" :key="item.id">{{item.idorder}}<input type="text" value="item.tglorder"></option>
                                                            </select>
                                                        </div>
                                                    </div>    
                                                </div>
                                                <div class="form-group">
                                                    <input type="file" name="fileorder" @change="OrderImageChange" multiple class="form-control">
                                                </div>
                                            </div>
                                            <div class="col-md-6">
                                                <h5 class="display-6">Upload Details</h5>
                                                <hr>
                                                <div class="form-group">
                                                    <div class="input-group">
                                                        <div class="input-group-prepend">
                                                            Gambar Upload untuk tanggal : {{tglorderupload}}
                                                        </div>
                                                    </div>
                                                </div>
                                                <!-- <div class="form-group">
                                                    <div class="input-group">
                                                        <div class="input-group-prepend">
                                                            Tanggal Upload : {{tanggal}}
                                                        </div>
                                                    </div>
                                                </div> -->
                                            </div>
                                        </div>
                                        <br>
                                        <div class="img-wrapp">
                                            <img :src="imageorder" height="10%" width="10%" onclick="document.getElementById('modal01').style.display='block'" class="w3-hover-opacity">
                                            <span class="delete-img" @click="EditdeleteImage(imageorder)">X</span>
                                        </div>
                                        <div id="modal01" class="w3-modal w3-animate-zoom" onclick="this.style.display='none'">
                                                <img class="w3-modal-content" :src="imageorder">
                                        </div> 
                                    </div>
                                    <br>
                                    <div class="w3-container">
                                        <button @click="uploadorder" onclick="document.getElementById('modal02').style.display='none'" class="btn-primary1">Upload</button>
                                    </div>
                                    <br>
                                </div>
                            </div>
                            <button onclick="document.getElementById('modal02').style.display='block'" class="btn-primary1">Upload Bukti Transfer</button>
                        <div>
                            <br>
                        </div>
                    </div>
                    <b-alert
                        :show="dismissCountDown"
                        dismissible
                        variant="warning"
                        @dismissed="dismissCountDown=0"
                        @dismiss-count-down="countDownChanged"
                        >
                        <p>Anda diharuskan mengupload hasil bukti transfer !!</p>
                        <b-progress
                            variant="warning"
                            :max="dismissSecs"
                            :value="dismissCountDown"
                            height="4px"
                        ></b-progress>
                    </b-alert>
                    <div>
                        <h3><b>Tabel Riwayat Pembelian</b></h3>
                        <div class="row justify-content-end"> 
                            <button @click="lihat" class="btn-primary1">Lihat</button>
                        </div>
                    </div>
                    <br>
                    
                    <template>
                        <b-container fluid>
                        <!-- User Interface controls -->
                        <b-row>
                            <b-col md="6" class="my-1">
                            <b-form-group label-cols-sm="3" label="Filter" class="mb-0">
                                <b-input-group>
                                <b-form-input v-model="filter" placeholder="Type to Search"></b-form-input>
                                <b-input-group-append>
                                    <b-button :disabled="!filter" @click="filter = ''">Clear</b-button>
                                </b-input-group-append>
                                </b-input-group>
                            </b-form-group>
                            </b-col>

                            <b-col md="6" class="my-1">
                            <b-form-group label-cols-sm="3" label="Sort" class="mb-0">
                                <b-input-group>
                                <b-form-select v-model="sortBy" :options="sortOptions">
                                    <option slot="first" :value="null">-- none --</option>
                                </b-form-select>
                                <b-form-select v-model="sortDesc" :disabled="!sortBy" slot="append">
                                    <option :value="false">Asc</option> <option :value="true">Desc</option>
                                </b-form-select>
                                </b-input-group>
                            </b-form-group>
                            </b-col>

                            <b-col md="6" class="my-1">
                            <b-form-group label-cols-sm="3" label="Per page" class="mb-0">
                                <b-form-select v-model="perPage" :options="pageOptions"></b-form-select>
                            </b-form-group>
                            </b-col>
                        </b-row>

                        <!-- Main table element -->
                        <b-table
                            show-empty
                            responsive
                            striped hover
                            bordered
                            head-variant="dark"
                            :items="items"
                            :fields="fields"
                            :current-page="currentPage"
                            :per-page="perPage"
                            :filter="filter"
                            :sort-by.sync="sortBy"
                            :sort-desc.sync="sortDesc"
                            @filtered="onFiltered"
                        >
                            <template slot="no" slot-scope="data">
                            {{ data.index + 1 }}
                            </template>

                            <template slot="tglorder" slot-scope="row">
                            {{ row.item.tglorder | formatDate}}
                            </template>

                            <template slot="produk" slot-scope="row">
                            <table>
                                <thead class="thead-dark">
                                <tr>
                                    <th>Produk</th>
                                    <th>Jumlah</th>
                                    <th>Harga</th>
                                </tr>
                                </thead>
                                <tbody class="table-secondary">
                                <tr v-for="item in row.item.produk" :key="item.id">
                                    <td>{{item.name}}</td>
                                    <td>{{item.quantity}}</td>
                                    <td>{{item.price | currency('', 0)}}</td>
                                </tr>
                                </tbody>
                            </table>
                            </template>

                            <template slot="total" slot-scope="row">
                            {{row.item.total | currency('IDR ', 0)}}
                            </template>

                            <template slot="bukti" slot-scope="row">
                            <div class="d-flex justify-content-center" >
                                <figure class="figure col-md-3">
                                    <img :src="row.item.image_order" class="img-thumbnail w3-hover-opacity" v-on:click="muncul(row.index)" alt="...">
                                </figure>
                                <figcaption class="figure-caption">Tanggal Upload {{ checkDate(row.item.tglupload) | formatDate}}</figcaption>
                            </div>
                            <div class="upload w3-modal w3-animate-zoom" onclick="this.style.display='none'">
                                <img class="w3-modal-content" :src="row.item.image_order">
                            </div>
                            </template>

                            <template slot="edit" slot-scope="row">
                            <button class="btn btn-primary" @click="editOrder(row.item)">Edit</button>
                            <button class="btn btn-danger" @click="deleteOrder(row.item.idorder)">Delete</button>
                            </template>
                        </b-table>

                        <b-row>
                            <b-col md="6" class="my-1">
                            <b-pagination
                                v-model="currentPage"
                                :total-rows="totalRows"
                                :per-page="perPage"
                                class="my-0"
                            ></b-pagination>
                            </b-col>
                        </b-row>

                        </b-container>
                    </template>


                    <!-- <div class="table-responsive">
                        <table class="table table-hover table-warning table-bordered">
                            <thead class="table-dark">
                                <tr>
                                <th scope="col-md-4">No</th>
                                <th scope="col">Tgl Order</th>
                                <th scope="col">Produk Detail</th>
                                <th scope="col">Total</th>
                                <th scope="col">Status</th>
                                <th scope="col">Upload Bukti Transfer</th>
                                </tr>
                            </thead>
                            <tbody>
                                <tr v-for="(item, index) in KatalogFiltered" :key="item.id">
                                <td>{{index +1}}</td>
                                <td>{{item.tglorder | formatDate}}</td>
                                <td>
                                    <table>
                                    <thead class="thead-primary">
                                        <tr>
                                        <th scope="col">Produk</th>
                                        <th scope="col">Quantity</th>
                                        <th scope="col">Harga</th>
                                        </tr>
                                    </thead>
                                    <tbody class="table-dark">
                                        <tr v-for="itema in item.produk" :key="itema.id">
                                        <td>{{itema.name}}</td>
                                        <td>{{itema.quantity}}</td>
                                        <td>{{itema.price | currency('', 0)}}</td>
                                        </tr>
                                    </tbody>
                                    </table>
                                </td>
                                <td>{{item.total | currency('IDR ', 0)}}</td>
                                <td>{{item.status}}</td>
                                <td>
                                    <div class="d-flex justify-content-center" >
                                        <figure class="figure col-md-7">
                                            <img :src="item.image_order" height="30%" width="30%" class="img-thumbnail w3-hover-opacity" v-on:click="muncul(index)" alt="...">
                                        </figure>
                                        <figcaption class="figure-caption">Tanggal Upload {{item.tglupload | formatDate}}</figcaption>
                                    </div>
                                        
                                    <div class="upload w3-modal w3-animate-zoom" onclick="this.style.display='none'">
                                        <img class="w3-modal-content" :src="item.image_order">
                                    </div>
                                        
                                </td>
                                </tr>
                            </tbody>
                        </table>
                    </div> -->
			</div>
		</section>
    </div>
</template>
<script>
import '@/assets/style3.css';
import Datepicker from 'vuejs-datepicker'

var moment = require('moment')
export default {
  name: 'MiniCart',
  data () {
    return {
        loggedIn: Boolean,
        status:'',
        data:{
            item:this.$store.state.cart,
            iduser:'',
            user:'',
            // tglorder:''
        },
        imageorder:'',
        date: new Date(),
        dismissSecs: 10,
        dismissCountDown: 0,
        showDismissibleAlert: false,
        tglorderupload:'',
        idOrder : 1,
        items:[],
        fields: [
            { key: 'no', label: 'No'},
            { key: 'tglorder', label: 'Tanggal Order', sortable: true, sortDirection: 'desc' },
            { key: 'idorder', label: 'No Pesanan', sortable: true, sortDirection: 'desc' },
            { key: 'produk', label: 'Produk Detail', sortable: true, class: 'text-center', sortDirection: 'desc' },
            { key: 'total', label: 'Total', sortable: true, class: 'text-center', sortDirection: 'desc' },
            { key: 'status', label: 'Status', sortable: true, class: 'text-center', sortDirection: 'desc' },
            { key: 'bukti', label: 'Upload Bukti Pembayaran', sortable: true, class: 'text-center', sortDirection: 'desc' },
        ],
        totalRows: 1,
        currentPage: 1,
        perPage: 5,
        pageOptions: [5, 10, 15],
        sortBy: null,
        sortDesc: false,
        sortDirection: 'asc',
        filter: null,
    }
  },
  components:{
      Datepicker
  },
  watch:{
    OrderImage(val){
        // console.log(val[0].url)
        this.imageorder = val[0].url
    },
    KatalogFiltered(val){
      this.items = val
      this.totalRows = this.items.length
    },
  },
  mounted(){
    this.idOrder = this.IdFixFiltered
  },
  computed:{
    sortOptions() {
        // Create an options list from our fields
        return this.fields
        .filter(f => f.sortable)
        .map(f => {
            return { text: f.label, value: f.key }
        })
    },
    IdFixFiltered(){
        return this.$store.getters.IdFixFilter
    },
    OrderImage(){
        return this.$store.state.orderimages
    },
    KatalogFiltered(){
        return this.$store.getters.KatalogFilter
    },
    tanggal(){
        return moment(this.date).format('MMMM Do YYYY, h:mm:ss a');
    },
    statusLoggedIn(){
        return this.$store.getters.loggedIn
    },
    userorder(){
        return localStorage.getItem('name')
    },
    iduserorder(){
        return localStorage.getItem('id')
    },
    itemId(){
        var sum = []
        for (var i = 0; i < this.data.item.length; i++){
        // sum += ' *' + this.data.item[i].quantity + ' buah ' + this.data.item[i].productName;
        sum.push(this.data.item[i].id_produk)
        }
        return sum
    },
    itemJumlah(){
        var quantity = []
        for (var i = 0; i < this.data.item.length; i++){
            quantity.push(this.data.item[i].quantity)
        }
        return quantity
    },
    itemHarga(){
        var harga = []
        for (var i = 0; i < this.data.item.length; i++){
            harga.push(this.data.item[i].quantity * this.data.item[i].productPrice)
        }
        return harga
    },
    totalharga(){
        var sum = 0
        for (var i = 0; i < this.$store.state.cart.length; i++){
        sum += this.$store.state.cart[i].productPrice * this.$store.state.cart[i].quantity;
        }
        return sum
    },
    totalquantity(){
        var sum = 0
        for (var i = 0; i < this.$store.state.cart.length; i++){
        sum += this.$store.state.cart[i].quantity;
        }
        return sum
    },
    discount(){
        var sum = 0
        if(this.totalquantity == 5){
            sum = 10
            if(this.totalquantity < 5){
                sum = 0
            }
        }else{
            if(this.totalquantity > 5){
                sum = 20
            }
        }
        return sum
    },
    total(){
        var sum = 0
        const sum1 = (this.totalharga * (this.discount/100))
        sum = this.totalharga - sum1
        return sum
    },
  },
  created(){
    if(!this.statusLoggedIn){
        this.$router.push({ name: 'home-error' })
    }
    // this.$store.dispatch('retrieveIdfix')
    this.data.user = this.iduserorder
    const iduser = this.data.user
    this.$store.dispatch('selectByName', iduser)
  },
  methods:{
      checkDate(time){
          if(time == "0001-01-01T00:00:00Z"){
              return  ""
          }else{    
              return time
          }
      },
    onFiltered(filteredItems) {
      // Trigger pagination to update the number of buttons/pages due to filtering
      this.totalRows = filteredItems.length
      this.currentPage = 1
    },
    EditdeleteImage(imageorder){
      let nama = imageorder
      let path = nama.substring(33) 
    //   console.log(path)
      this.$store.dispatch('deleteEditUploadImage',path)
      .then(response => {
          this.imageorder = ''
      })
    },
    muncul(i){
    document.getElementsByClassName('upload')[i].style.display = "block"
    },
    uploadorder(){
        const no = this.tglorderupload
        const image = this.imageorder
        // const tglupload = this.tanggal
        
        const upload=[
            no,image
        ]

        this.$store.dispatch('ImageOrder', upload)
        .then(response => {
          $('#modal02').modal('hide');
          Toast.fire({
            type: 'success',
            title: 'Upload successfully'
          })
        })
    },
    OrderImageChange(e){
      const files = e.target.files || e.dataTransfer.files;
      
      if (!files.length) {
        console.log('no files');
      }
      let form = new FormData();
      for (var i = 0; i < files.length; i++) {
        let file = files.item(i);
        form.append('order',file, file.name);
      }
      this.$store.dispatch('OrderUpload', form)
      .then(response => {
      })
    },
    lihat(){
      this.data.user = this.iduserorder
      const iduser = this.data.user
      this.$store.dispatch('selectByName', iduser)
    },
    countDownChanged(dismissCountDown) {
        this.dismissCountDown = dismissCountDown
    },
    increaseQty(id) {   
        this.$store.commit('increment', id)
    },
    decreaseQty(id) {
        this.$store.commit('decrement', id)
    },
    order(){
        const data = this.data
        // data.user = this.userorder
        data.iduser = this.iduserorder
        // data.tglorder = this.tanggal
        this.dismissCountDown = this.dismissSecs
        const ItemId = this.itemId
        const ItemJumlah = this.itemJumlah
        const ItemQuantity = this.totalquantity
        const ItemTotal = this.total
        const IdOrder = this.IdFixFiltered
        const ItemHarga = this.itemHarga
        const order=[
            IdOrder,data,ItemId,ItemJumlah,ItemHarga
        ]
        this.$store.dispatch('order', order)
        
        
        this.$store.dispatch('destroycart')

        .then(response => {
            Toast.fire({
                type: 'success',
                title: 'Order successfully'
            })
        })
    }
  }
}
</script>
<style lang="scss">
.bawah{
    padding-bottom: 10px;
}
.img-wrapp{
  position: relative;
  padding-bottom: 8px;
    display: block;
}
.img-wrapp span.delete-img{
    position: center;
    top: 10px;
    left: 1px;
}
.img-wrapp span.delete-img:hover{
  cursor: pointer;
}
// .tableaa {
//   border-collapse: collapse;
// }

.tableaa {
  border: 1px solid black;
}
</style>
