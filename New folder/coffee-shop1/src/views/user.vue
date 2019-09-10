<template>
  <div class="products">
    <div class="container">
      <div class="intro h-100">
        <div class="row h-100 align-items-center">
          <div class="col-md-6 ml-3">
            <h3>Profile settings</h3>
            <p>
              Change your profile settings here
            </p>
          </div>
          <div class="col-md-5">
              <img src="./image/profile.svg" width="300" alt="" class="img-fluid">
          </div>
        </div>
      </div>
      <hr>
      
      <h3 class="d-inline-block">User Profile</h3>
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

            <template slot="edit" slot-scope="row">
              <button class="btn btn-primary" @click="editProfile(row.item)">Edit</button>
              <button class="btn btn-danger" @click="deleteProfile(row.item.id)">Delete</button>
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
      <table id="dtBasicExample" class="table table-striped table-bordered table-sm" cellspacing="0" width="100%">
        <thead class="thead-dark">
          <tr>
            <th>Nama Lengkap</th>
            <th>No. Telepon</th>
            <th>Alamat</th>
            <th>Postcode</th>
            <th>Role</th>
            <th>Modify</th>
          </tr>
        </thead>
        <tbody>
          <tr v-for="profile in profileFiltered" :key="profile.id">
            <td>
              {{profile.name}}
            </td>
            <td>
              {{profile.phone}}
            </td>
            <td>
              {{profile.address}}
            </td>
            <td>
              {{profile.postcode}}
            </td>
            <td>
              {{profile.role}}
            </td>
            <td>
              <button class="btn btn-primary" @click="editProfile(profile)">Edit</button>
              <button class="btn btn-danger" @click="deleteProfile(profile.id)">Delete</button>
            </td>
          </tr>
        </tbody>
      </table>
      </div> -->

      <div class="modal fade" id="edit" tabindex="-1" role="dialog" aria-labelledby="editLabel" aria-hidden="true">
        <div class="modal-dialog modal-lg" role="document">
          <div class="modal-content">
            <div class="modal-header">
              <h5 class="modal-title" id="editLabel">Edit Profile</h5>
              <button type="button" class="close" data-dismiss="modal" aria-label="Close">
                <span aria-hidden="true">&times;</span>
              </button>
            </div>
            <div class="modal-body">
              <ul class="nav nav-pills ml-3" id="myTab" role="tablist">
                <li class="nav-item">
                  <a class="nav-link active" id="profile-tab" data-toggle="tab" href="#profile" role="tab" aria-controls="profile" aria-selected="true">Profile</a>
                </li>
                <li class="nav-item">
                  <a class="nav-link"  id="account-tab" data-toggle="tab" href="#account" role="tab" aria-controls="account" aria-selected="false">Account settings</a>
                </li>
              </ul>

              <div class="tab-content" id="myTabContent">
                <div class="tab-pane fade show active pt-3" id="profile" role="tabpanel" aria-labelledby="profile-tab">
                  <div class="container">
                    <div class="row">
                      <div class="col-md-6">
                        <div class="form-group">
                          <input type="text" name="" v-model="editObj.name" placeholder="Full name" class="form-control">
                          <input type="hidden" placeholder="nama profile" v-model="editObj.id" class="form-control">
                        </div>
                      </div>
                      <div class="col-md-6">
                        <div class="form-group">
                          <input type="text"  v-model="editObj.phone" placeholder="Phone" class="form-control">
                        </div>
                      </div>
                      <div class="col-md-12">
                        <div class="form-group">
                          <input type="text"  v-model="editObj.address" placeholder="Address" class="form-control">
                        </div>
                      </div>
                      <div class="col-md-4">
                        <div class="form-group">
                          <input type="text"  v-model="editObj.postcode" placeholder="Postcode" class="form-control">
                        </div>
                      </div>
                      <div class="col-md-4">
                        <div class="form-group">
                          <input type="text"  v-model="editObj.role" placeholder="Role" class="form-control">
                        </div>
                      </div>
                      <div class="col-md-4">
                        <div class="form-group">
                          <input type="submit" @click="updateProfile()" value="Save" class="btn btn-primary w-100">
                        </div>
                      </div>
                    </div>
                  </div>
                </div>

                <div class="tab-pane fade pt-3" id="account" role="tabpanel" aria-labelledby="account-tab">
                  <div class="container">
                      <div class="row">
                        <div class="col-md-">
                          <div class="alert alert-info">
                            Please use the Reset password email button for reseting the password. The form doens't work currently
                          </div>
                        </div>
                        <div class="col-md-6">
                          <div class="form-group">
                            <input type="text"  v-model="editObj.name" placeholder="User name" class="form-control">
                            <input type="hidden" placeholder="nama profile" v-model="editObj.id" class="form-control">
                          </div>
                        </div>
                        <div class="col-md-6">
                          <div class="form-group">
                            <div class="input-group">
                              <div class="input-group-prepend">
                                <span class="input-group-text">Email</span>
                              </div>
                              <input type="text" class="form-control" v-model="editObj.email" placeholder="Email">
                            </div>
                          </div>
                        </div>
                        <div class="col-md-6">
                          <div class="form-group">
                            <input type="text"  v-model="editObj.nemail" placeholder="New email" class="form-control">
                          </div>
                        </div>
                        <div class="col-md-6">
                          <div class="form-group">
                            <div class="input-group">
                              <div class="input-group-prepend">
                                <span class="input-group-text">Pass</span>
                              </div>
                              <input type="text" class="form-control" v-model="editObj.password" placeholder="Password">
                            </div>
                          </div>
                        </div>
                        <div class="col-md-6">
                          <div class="form-group">
                            <input type="text"  v-model="editObj.npassword" placeholder="New password" class="form-control">
                          </div>
                        </div>

                        <!-- <div class="col-md-4">
                          <div class="form-group">
                              <input type="file" @change="uploadImage" class="form-control">
                          </div>
                        </div> -->

                        <div class="col-md-6">
                          <div class="form-group">
                            <input type="button" @click="resetPasswordEmail()" value="Reset password email" class="btn btn-success w-100">
                          </div>
                        </div>
                        <div class="col-md-6">
                          <div class="form-group">
                            <input type="submit" @click="updateAccount()" value="Save Changes" class="btn btn-primary w-100">
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
  </div>
</template>

<script>
export default {
  name:'profile',
  data(){
    return{
      id:'',
      name:'',
      email:'',
      nemail:'',
      password:'',
      npassword:'',
      phone:'',
      address:'',
      postcode:'',
      role:'',
      editStatus: 0,
      editObj : [],
      items:[],
      fields: [
        { key: 'no', label: 'No'},
        { key: 'name', label: 'Nama Lengkap', sortable: true, sortDirection: 'desc' },
        { key: 'phone', label: 'No. Telepon', sortable: true, class: 'text-center', sortDirection: 'desc' },
        { key: 'address', label: 'Alamat', sortable: true, class: 'text-center', sortDirection: 'desc' },
        { key: 'postcode', label: 'Postcode', sortable: true, class: 'text-center', sortDirection: 'desc' },
        { key: 'role', label: 'Role', sortable: true, class: 'text-center', sortDirection: 'desc' },
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
    profileFiltered(val){
      this.items = val
      this.totalRows = this.items.length
    }
  },
  created() {
    this.$store.dispatch('retrieveProfile')
  },
  computed:{
    profileFiltered(){
      return this.$store.getters.ProfileFilter
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
  methods:{
    onFiltered(filteredItems) {
      // Trigger pagination to update the number of buttons/pages due to filtering
      this.totalRows = filteredItems.length
      this.currentPage = 1
    },
    editProfile(profile){
      $('#edit').modal('show');
      this.editObj = profile
    },
    updateProfile(){
      this.$store.dispatch('updateDataProfile', {
        id : this.editObj.id,
        name: this.editObj.name,
        phone: this.editObj.phone,
        address: this.editObj.address,
        postcode: this.editObj.postcode,
        role: this.editObj.role,
      })
      .then(response => {
        this.$router.push({ name: 'user' })
        $('#edit').modal('hide');
        Toast.fire({
          type: 'success',
          title: 'Edit  successfully'
        })
      })
    },
    resetPasswordEmail(){
      this.$store.dispatch('resetPasswordEmail', {
        id : this.editObj.id,
        email: this.editObj.email,
        password: this.editObj.password,
      })
      .then(response => {
        this.$router.push({ name: 'profile' })
        $('#edit').modal('hide');
        Toast.fire({
        type: 'success',
        title: 'Reset  successfully'
      })
      })
    },
    updateAccount(){
      this.$store.dispatch('updateDataAccount', {
        id : this.editObj.id,
        name : this.editObj.name,
        nemail: this.editObj.nemail,
        npassword: this.editObj.npassword,
      })
      .then(response => {
        this.$router.push({ name: 'profile' })
        $('#edit').modal('hide');
        Toast.fire({
        type: 'success',
        title: 'Save Email and Password successfully'
      })
      })
    },
    deleteProfile(id){
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
            this.$store.dispatch('deleteProfile', id)
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

<style scoped>
.modal-admin {
    width: 75%;
}
</style>


