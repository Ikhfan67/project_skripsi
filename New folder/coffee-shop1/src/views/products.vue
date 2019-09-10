<template>
  <div class="products">
    <div class="container">   
      <div class="intro h-100">
        <div class="row h-100 justify-content-center align-items-center">
          <div class="col-md-6">
              <h3>Products Page</h3>
              <p>
                Change your products settings here
              </p>
          </div>
          <div class="col-md-5">
              <img src="./image/products.svg" alt="" width="300" class="img-fluid">
          </div>
        </div>
      </div>
      <hr>
      <h3 class="d-inline-block">Products list</h3>
      <button @click="addProduct()" class="btn btn-primary float-right">Add Product</button>
      <div class="table-responsive">
        <table id="dtBasicExample" class="table table-striped table-bordered table-sm" cellspacing="0" width="100%">
          <thead class="thead-dark">
            <tr>
              <th>Name</th>
              <th>Image</th>
              <th>Price</th>
              <th>Description</th>
              <th>Modify</th>
            </tr>
          </thead>
          <tbody>
            <tr v-for="(product,index) in productsFiltered" :key="product.id">
              <td>
                {{product.name}}
              </td>
              <td>
                <!-- <img :src="product.image" alt="" width="100px"> -->
                <div class="d-flex justify-content-center" >
                    <figure class="figure">
                        <img :src="product.image" width="100px" class="img-thumbnail w3-hover-opacity" v-on:click="muncul(index)">
                    </figure>
                </div>
                <div class="upload w3-modal w3-animate-zoom" onclick="this.style.display='none'">
                    <img class="w3-modal-content" :src="product.image">
                </div>
              </td>
              <td>
                {{product.price | currency('Rp. ', 0)}}
              </td>
              <td v-html="product.description">
              </td>
              <td>
                <button class="btn btn-primary" @click="editProduct(product)">Edit</button>
                <button class="btn btn-danger" @click="deleteProduct(product.id)">Delete</button>
              </td>
            </tr>
          </tbody>
        </table>
      </div>
    </div>

      <!-- Modal -->
      <div class="modal fade" id="product" tabindex="-1" role="dialog" aria-labelledby="productLabel" aria-hidden="true">
        <div class="modal-dialog modal-lg" role="document">
          <div class="modal-content">
            <div class="modal-header">
              <h5 class="modal-title" id="productLabel">Add Product</h5>
              <button type="button" class="close" data-dismiss="modal" aria-label="Close">
                <span aria-hidden="true">&times;</span>
              </button>
            </div>
            <div class="modal-body">
              <div class="row">
                <div class="col-md-8">
                  <div class="form-group">
                    <input type="text" placeholder="nama produk" v-model="name" class="form-control">
                    <input type="hidden" placeholder="nama produk" v-model="id" class="form-control">
                  </div>
                  <div class="form-group">
                    <vue-editor placeholder="deskripsi" v-model="description" ></vue-editor>
                  </div>             
                </div>
                <div class="col-md-4">
                  <h5 class="display-6">Product Details</h5>
                  <hr>
                  <div class="form-group">
                    <input type="text" placeholder="harga" v-model="price" class="form-control">
                  </div>
                  <div class="form-group">
                    <input type="file" name="fileipan" @change="ImageChange" multiple class="form-control">
                  </div>
                  <div class="form-group d-flex" >
                    <div class="p-1">
                      <div class="img-wrapp" v-for="(item, index) in imageFiltered" :key="item.url">
                        <img :src="item.url" alt="" width="80px">
                        <span class="delete-img" @click="deleteImage(item,index)">X</span>
                      </div>
                    </div>
                  </div>
                </div>
              </div>
            </div>
            <div class="modal-footer">
              <button type="button" class="btn btn-secondary" data-dismiss="modal">Close</button>
              <button @click="saveDataProduk" type="button" class="btn btn-primary">Tambah</button>
            </div>
          </div>
        </div>
      </div>

      <!-- Modal -->
      <div class="modal fade" id="edit" tabindex="-1" role="dialog" aria-labelledby="editLabel" aria-hidden="true">
        <div class="modal-dialog modal-lg" role="document">
          <div class="modal-content">
            <div class="modal-header">
              <h5 class="modal-title" id="editLabel">Edit Product</h5>
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
                        <span class="input-group-text">Nama Produk</span>
                      </div>
                      <input type="text" class="form-control" v-model="editObj.name" placeholder="nama produk">
                    </div>
                  </div>
                  <div class="form-group">
                    <input type="hidden" placeholder="nama produk" v-model="editObj.id" class="form-control">
                  </div>
                  <div class="form-group">
                    <div class="input-group">
                      <div class="input-group-prepend">
                        <span class="input-group-text">Deskripsi</span>
                      </div>
                    </div>
                    <vue-editor placeholder="deskripsi" v-model="editObj.description"></vue-editor>
                  </div>
                </div>

                <div class="col-md-4">
                  <h5 class="display-6">Product Details</h5>
                    <hr>
                  <div class="form-group">
                    <div class="input-group">
                      <div class="input-group-prepend">
                        <span class="input-group-text">Harga</span>
                      </div>
                      <input type="text" class="form-control" v-model="editObj.price" placeholder="harga">
                    </div>
                  </div>
                  <div class="form-group">
                    <input type="file" name="fileipan" @change="EditImageChange" multiple class="form-control">
                  </div>
                  <div class="form-group d-flex" >
                    <div class="p-1">
                      <div class="img-wrapp">
                        <img :src="editObj.image" alt="" width="80px">
                        <span class="delete-img" @click="EditdeleteImage(editObj)">X</span>
                      </div>
                    </div>
                  </div>
                </div>
              </div>
            </div>
            <div class="modal-footer">
              <button type="button" class="btn btn-secondary" data-dismiss="modal">Close</button>
              <button @click="updateProduk()" type="button" class="btn btn-primary">Save</button>
            </div>
          </div>
        </div>
      </div>

  </div>
</template>

<script>
import { VueEditor } from "vue2-editor";
import { mdbDatatable } from 'mdbvue';
import '@/assets/style3.css';

export default {
  name: 'products',
  components: {
    VueEditor,
    mdbDatatable
  },
  data () {
    return {
      name: '',
      description:'',
      price: '',
      id: '',
      image:'',
      editStatus: 0,
      editObj: [],
      imagebaru:[]
    }
  },
  created() {
    this.$store.dispatch('retrieveProducts')
  },
  watch:{
    editimage(val){
     this.editObj.image = val
    }
  },
  computed:{
    productsFiltered(){
      return this.$store.getters.ProductsFilter
    },
    imageFiltered(){
      return this.$store.getters.ImagesFilter
    },
    editimage(){
      return this.$store.state.editimage
    }
  },
  methods:{
    muncul(i){
      document.getElementsByClassName('upload')[i].style.display = "block"
    },
    EditImageChange(e){
      const files = e.target.files || e.dataTransfer.files;
      if (!files.length) {
        console.log('no files');
      }
      let form = new FormData();
      for (var i = 0; i < files.length; i++) {
        let file = files.item(i);
          form.append('editfileipan',file, file.name);
      }
      this.$store.dispatch('EditonUpload', form)
      .then(response => {
      })
    },
    ImageChange(e){
      const files = e.target.files || e.dataTransfer.files;
      if (!files.length) {
        console.log('no files');
      }
      this.image = e.target.files[0].name
      let form = new FormData();
      for (var i = 0; i < files.length; i++) {
          let file = files.item(i);
          form.append('fileipan',file, file.name);
      }
      this.$store.dispatch('onUpload', form)
    },
    EditdeleteImage(editObj){
      let nama = this.editObj.image
      let path = nama.substring(28) 
      this.$store.dispatch('deleteEditFileImage',path)
      .then(response =>{
      this.editObj.image = ''
      })    
    },
    deleteImage(item, index){
      const path = item.name
      this.$store.dispatch('deleteFileImage',path)
    },
    addProduct(){
      $('#product').modal('show');
    },
    saveDataProduk() {
      if(this.price.trim().length == 0){
          return
        }
      this.$store.dispatch('saveDataProduk', {
        id: this.id,
        name: this.name,
        description: this.description,
        price: this.price,
        image: this.image,
      })
      .then(response => {
        this.$router.push({ name: 'products' })
        this.reset()
        $('#product').modal('hide');
        Toast.fire({
            type: 'success',
            title: 'Add  successfully'
        })
      })
    },
    reset(){
      Object.assign(this.$data, this.$options.data.apply(this));
    },
    deleteProduct(id){
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
            this.$store.dispatch('deleteProduct', id)            
            Toast.fire({
              type: 'success',
              title: 'Deleted  successfully'
            })
          }
          })
    },
    editProduct(product){
      $('#edit').modal('show');
      this.editObj = product
    },
    updateProduk(){
      this.$store.dispatch('updateDataProduct', {
        id : this.editObj.id,
        name: this.editObj.name,
        description: this.editObj.description,
        price: this.editObj.price,
        image: this.editObj.image
      })
      .then(response => {
        this.$router.push({ name: 'products' })
        this.reset()
        $('#edit').modal('hide');
        Toast.fire({
          type: 'success',
          title: 'Edit  successfully'
        })
      })
    },
    updateImage(val){
      this.editObj.image = val
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
</style>


