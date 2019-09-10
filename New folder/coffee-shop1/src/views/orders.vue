<template>
  <div class="products">
    <div class="container">
      <div class="intro h-100">
        <div class="row h-100 justify-content-center align-items-center">
          <div class="col-md-6">
              <h3>Orders Page</h3>
              <p>Change your orders settings here</p>
          </div>
          <div class="col-md-5">
              <img src="./image/orders.svg" width="300" alt="" class="img-fluid">
          </div>
        </div>
      </div>
      <hr>
      
      <h3 class="d-inline-block">Orders list</h3>
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

            <template slot="tglorder" slot-scope="data">
              {{ data.item.tglorder | formatDate }}
            </template>

            <template slot="produk" slot-scope="row">
              <table>
                <thead class="thead-dark">
                  <tr>
                    <th>Produk</th>
                    <th>Jumlah</th>
                  </tr>
                </thead>
                <tbody class="table-secondary">
                  <tr v-for="item in row.item.produk" :key="item.id">
                    <td>{{item.name}}</td>
                    <td>{{item.quantity}}</td>
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
                <figcaption class="figure-caption">Tanggal Upload {{row.item.tglupload | formatDate}}</figcaption>
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
      
      <!-- <div class="product-test">
        <div class="table-responsive">
          <table id="dtBasicExample" class="table table-striped table-bordered table-sm" cellspacing="0" width="100%">
            <thead class="thead-dark">
              <tr>
                <th>No</th>
                <th>Tanggal Order</th>
                <th>Nama</th>
                <th>Produk Detail</th>
                <th>Total</th>
                <th>Status</th>
                <th>Bukti Pembayaran</th>
                <th >Modify</th>
              </tr>
            </thead>
            <tbody>
            <tr v-for="(order, index) in ordersFiltered" :key="order.id"> 
              <td>{{index}}</td>
              <td>{{order.tglorder}}</td>
              <td>{{order.name}}</td>
              <td>
                  <table>
                    <thead class="thead-dark">
                      <tr>
                        <th>Produk</th>
                        <th>Jumlah</th>
                      </tr>
                    </thead>
                    <tbody class="table-secondary">
                      <tr v-for="item in order.produk" :key="item.id">
                        <td>{{item.name}}</td>
                        <td>{{item.quantity}}</td>
                      </tr>
                    </tbody>
                  </table>
              </td>
              <td>{{order.total | currency('IDR ', 0)}}</td>
              <td>{{order.status}}</td>
              <td>
                <div class="d-flex justify-content-center" >
                    <figure class="figure col-md-3">
                        <img :src="order.image_order" class="img-thumbnail w3-hover-opacity" v-on:click="muncul(index)" alt="...">
                    </figure>
                    <figcaption class="figure-caption">Tanggal Upload {{order.tglupload}}</figcaption>
                </div>
                <div class="upload w3-modal w3-animate-zoom" onclick="this.style.display='none'">
                    <img class="w3-modal-content" :src="order.image_order">
                </div>
              </td>
              <td>
                <button class="btn btn-primary" @click="editOrder(order)">Edit</button>
                <button class="btn btn-danger" @click="deleteOrder(order.idorder)">Delete</button>
              </td>
            </tr>
            </tbody>
          </table>
        </div>
      </div> -->
    </div>
    <div class="modal fade" id="edit" tabindex="-1" role="dialog" aria-labelledby="editLabel" aria-hidden="true">
        <div class="modal-dialog modal-lg" role="document">
          <div class="modal-content">
            <div class="modal-header">
              <h5 class="modal-title" id="editLabel">Edit Order</h5>
              <button type="button" class="close" data-dismiss="modal" aria-label="Close">
                <span aria-hidden="true">&times;</span>
              </button>
            </div>
            <div class="modal-body">
              <div class="row">
                <div class="col-md-8">
                  <div class="form-group">
                    <div class="input-group">
                      <div class="input-group-prepend">
                        <span class="input-group-text">Tanggal Order</span>
                      </div>
                      <input type="text" class="form-control" v-model="editData.tglorder" placeholder="tanggal order">
                    </div>
                  </div>
                  <div class="form-group">
                    <input type="hidden" placeholder="id-user" v-model="editData.iduser" class="form-control">
                  </div>
                  <!-- <div class="form-group">
                    <div class="input-group">
                      <div class="input-group-prepend">
                        <span class="input-group-text">Deskripsi</span>
                      </div>
                    </div>
                    <vue-editor placeholder="deskripsi" v-model="editData.produk"></vue-editor>
                  </div> -->
                  <ul class="list-group">
                    <li class="list-group-item list-group-item-action list-group-item-secondary d-flex justify-content-between align-items-center">
                      Product List
                    </li>
                    <li class="list-group-item d-flex justify-content-between align-items-center" v-for="itema in this.editData.produk" :key="itema.id">
                      {{itema.name}}
                      <span class="badge badge-primary badge-pill">{{itema.quantity}}</span>
                    </li>
                    
                  </ul>
                </div>

                <div class="col-md-4">
                  <h5 class="display-6">Product Details</h5>
                    <hr>
                  <div class="form-group">
                    <div class="input-group">
                      <div class="input-group-prepend">
                        <span class="input-group-text">Total Order</span>
                      </div>
                      <input type="text" class="form-control" v-model="editData.total" placeholder="harga">
                    </div>
                  </div>
                  <div class="form-group">
                    <div class="input-group">
                      <div class="input-group-prepend">
                        <span class="input-group-text">Satus</span>
                      </div>
                      <input type="text" class="form-control" v-model="editData.status" placeholder="harga">
                    </div>
                  </div>
                  <!-- <div class="form-group">
                    <input type="file" name="fileipan" @change="EditImageChange" multiple class="form-control">
                  </div> -->
                  <div class="form-group d-flex" >
                    <div class="p-1">
                      <div class="img-wrapp">
                        <img :src="editData.image_order" alt="" width="80px">
                        <!-- <span class="delete-img" @click="EditdeleteImage(editData)">X</span> -->
                        <div>tgl upload:</div>
                        <div>{{this.editData.tglupload}}</div>
                      </div>
                    </div>
                  </div>
                </div>
              </div>
            </div>
            <div class="modal-footer">
              <button type="button" class="btn btn-secondary" data-dismiss="modal">Close</button>
              <button @click="updateDataOrder()" type="button" class="btn btn-primary">Save</button>
            </div>
          </div>
        </div>
    </div>
  </div>
</template>

<script>
import '@/assets/style3.css';

export default {
  name:'order',
  data(){
    return{
      editData:[],
      items:[],
      fields: [
        { key: 'no', label: 'No'},
        { key: 'tglorder', label: 'Tanggal Order', sortable: true, sortDirection: 'desc' },
        { key: 'name', label: 'Nama', sortable: true, class: 'text-center', sortDirection: 'desc' },
        { key: 'produk', label: 'Produk Detail', sortable: true, class: 'text-center', sortDirection: 'desc' },
        { key: 'total', label: 'Total', sortable: true, class: 'text-center', sortDirection: 'desc' },
        { key: 'status', label: 'Status', sortable: true, class: 'text-center', sortDirection: 'desc' },
        { key: 'bukti', label: 'Bukti Pembayaran', sortable: true, class: 'text-center', sortDirection: 'desc' },
        { key: 'edit', label: 'Modify', sortable: true, class: 'text-center', sortDirection: 'desc' },
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
  watch:{
    ordersFiltered(val){
      this.items = val
      this.totalRows = this.items.length
    },
  },
  created(){
    this.$store.dispatch('retrieveOrder')
  },
  computed:{
    ordersFiltered(){
      return this.$store.getters.OrdersFilter
    },
    sortOptions() {
      // Create an options list from our fields
      return this.fields
        .filter(f => f.sortable)
        .map(f => {
          return { text: f.label, value: f.key }
        })
    }
  },
  mounted(){
    
  },
  methods:{
    onFiltered(filteredItems) {
      // Trigger pagination to update the number of buttons/pages due to filtering
      this.totalRows = filteredItems.length
      this.currentPage = 1
    },
    deleteOrder(id){
       Swal.fire({
        title: 'Are you sure?',
        text: "You won't be able to revert this!",
        type: 'warning',
        showCancelButton: true,
        confirmButtonColor: '#3085d6',
        cancelButtonColor: '#d33',
        confirmButtonText: 'Yes, delete it!'
      }).then((result) => {
          if (result.value) {
            this.$store.dispatch('deleteOrder', id)            
            Toast.fire({
              type: 'success',
              title: 'Deleted  successfully'
            })
          }
          })
    },
    editOrder(order){
      $('#edit').modal('show');
      this.editData = order
    },
    updateDataOrder(){
      var tgl = this.editData.tglorder
      var status = this.editData.status
      var total = this.editData.total
      var id = this.editData.idorder
      var data=[tgl,status,total, id]
      this.$store.dispatch('UbahDataOrder', data)
      $('#edit').modal('hide');
        Toast.fire({
            type: 'success',
            title: 'Edit Successfully'
        })
    },
    muncul(i){
      document.getElementsByClassName('upload')[i].style.display = "block"
    },
  }
}
</script>
