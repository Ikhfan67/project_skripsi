<template>
  <div class="products">
    <div class="container">   
      <div class="intro h-100">
        <div class="row h-100 justify-content-center align-items-center">
          <div class="col-md-6">
              <h3>Comment Page</h3>
              <p>
                Change your comment settings here
              </p>
          </div>
          <div class="col-md-5">
              <img src="./image/products.svg" alt="" width="300" class="img-fluid">
          </div>
        </div>
      </div>
      <hr>
      <h3 class="d-inline-block">Comments List</h3>
      <div class="table-responsive">
        <table id="dtBasicExample" class="table table-striped table-bordered table-sm" cellspacing="0" width="100%">
          <thead class="thead-dark">
            <tr>
              <th>No</th>
              <th>Nama</th>
              <th>Email</th>
              <th>Phone</th>
              <th>Message</th>
              <th>Modify</th>
            </tr>
          </thead>
          <tbody>
            <tr v-for="(comment, index) in Comments" :key="comment.id">
              <td>
                {{index+1}}
              </td>
              <td>
                {{comment.name}}
              </td>
              <td>
                {{comment.email}}
              </td>
               <td>
                {{comment.phone}}
              </td>
               <td>
                {{comment.message}}
              </td>
              <td>
                <button class="btn btn-danger" @click="deleteKomen(comment.id)">Delete</button>
              </td>
            </tr>
          </tbody>
        </table>
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
      name: '',
      
    }
  },
  created() {
    this.$store.dispatch('retrieveComments')
  },
  watch:{
    
  },
  computed:{
    Comments(){
      return this.$store.state.comments
    },
  },
  methods:{
    deleteKomen(id){
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
            this.$store.dispatch('deleteKomen', id)            
            Toast.fire({
              type: 'success',
              title: 'Deleted  successfully'
            })
          }
          })
    },
  },
}
</script>

<style lang="scss">

</style>


