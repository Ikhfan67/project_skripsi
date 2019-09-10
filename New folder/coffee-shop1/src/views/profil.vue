<template>
    <div class="profil">
        <Navbar></Navbar>
        <div class="container mt-5 pt-5">
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
            <div class="card mb-3" style="max-width: 100%;">
            <div class="row">
                <div class="col-md-6">
                    <b-container fluid class="p-4 bg-dark">
                    <b-row>
                        <b-col>
                            <div class="putih">
                                <h5>Gambar Profile</h5>
                            </div>
                            <b-img thumbnail fluid :src="ProfileImageFix" height="160%" width="160%" alt="Image 1"></b-img>
                        </b-col>
                    </b-row>
                    </b-container>
                    <div class="form-group">
                        <input type="file" name="fileorder" @change="ProfileImageChange" multiple class="form-control">
                    </div>
                    <section v-if="this.profileimage">
                        <div class="img-wrapp">
                            <img :src="profileimage" height="20%" width="20%">
                            <span class="delete-img" @click="EditdeleteImage(profileimage)">X</span>
                        </div>
                    </section>
                    <section v-else>
                        <b-alert show variant="danger"><a href="#" class="alert-link">Silahkan Pilih File Gambar</a></b-alert>
                    </section>
                    <button class="btn btn-primary" @click="upload" type="submit">Upload Gambar</button>
                </div>
                <div class="col md-6">
                <div class="card-body">
                    <!-- <div class="form-row tengah"> -->
                        <h5 class="card-title">Detail Profile</h5>
                    <!-- </div> -->
                    <form class="needs-validation" novalidate>
                    <div class="form-row">
                        <div class="col-md-5 mb-3">
                            <label for="validationCustom01">Name</label>
                            <input type="text" class="form-control" id="validationCustom01" v-model="data.name" value="Mark" required>
                            <div class="valid-feedback">
                                Looks good!
                            </div>
                        </div>
                        <div class="col-md-5 mb-3">
                            <label for="validationCustom02">Phone</label>
                            <input type="text" class="form-control" id="validationCustom02" v-model="data.phone" value="Otto" required>
                            <div class="valid-feedback">
                                Looks good!
                            </div>
                        </div>
                        <div class="col-md-5 mb-3">
                            <label for="validationCustomUsername">Email</label>
                            <div class="input-group">
                                <input type="text" class="form-control" id="validationCustomUsername"  v-model="data.email" aria-describedby="inputGroupPrepend" required>
                                <div class="invalid-feedback">
                                Please choose a username.
                                </div>
                            </div>
                        </div>
                        <div class="col-md-5 mb-3">
                            <label for="validationCustomUsername">Password</label>
                            <div class="input-group">
                                <input type="text" class="form-control" id="validationCustomUsername"  v-model="data.password" aria-describedby="inputGroupPrepend" required>
                                <div class="invalid-feedback">
                                Please choose a username.
                                </div>
                            </div>
                        </div>
                    </div>
                    <div class="form-row">
                        <div class="col-md-7 mb-3">
                            <label for="validationCustom03">Address</label>
                            <input type="text" class="form-control" id="validationCustom03" v-model="data.address" required>
                            <div class="invalid-feedback">
                                Please provide a valid city.
                            </div>
                        </div>
                        <div class="col-md-3 mb-3">
                            <label for="validationCustom04">Postcode</label>
                            <input type="text" class="form-control" id="validationCustom04" v-model="data.postcode" required>
                            <div class="invalid-feedback">
                                Please provide a valid state.
                            </div>
                        </div>
                    </div>
                    <div class="form-row tengah">
                    <section v-if="this.datakosong">
                        <button class="btn btn-primary" @click="tambahdata" type="button">Tambah Data</button>
                    </section>
                    <section v-else>
                        <button class="btn btn-primary" @click="editdata" type="button">Edit Data</button>
                    </section>
                    </div>
                    </form>
                </div>
                </div>
            </div>
            </div>
        </div>
        <footer>
          <p>&copy; 2019 <a href="https://ikhfannurali.blogspot.com" title="Cafe 3R.ASA">Cafe 3R.ASA.</a> All Rights Reserved. By : <a href="http://crwebstudio.com" target="_blank">Ikhfan Nur Ali</a></p>
        </footer>
    </div>
</template>
<script>
import '@/assets/profile.css';

export default {
    name: 'Profile',
    data () {
        return{
            profileimage : '',
            profileimagefix : '',
            data:[],
            datakosong:[]
        }
    },
    watch:{
        ProfileImage(val){
            // console.log(val[0].url)
            this.profileimage = val[0].url
        },
        ProfileImageFix(val){
            this.profileimagefix = val
            // console.log(val)
        },
        ProfileDataFix(val){
            this.data = val
            this.datakosong = 0
        }
    },
    computed:{
        ProfileImage(){
            return this.$store.state.profileimages
        },
        ProfileImageFix(){
            return this.$store.state.gambarprofile
        },
        ProfileDataFix(){
            return this.$store.state.dataprofile
        },
        userid(){
            return localStorage.getItem('id')
        },
    },
    created(){
        let data = this.userid
        this.$store.dispatch('GetImageProfile',data)
        this.$store.dispatch('GetDataProfile',data)
    },
    methods:{
        tambahdata(){
            let id = this.userid
            let isi = this.data
            let data =[id, isi]
            this.$store.dispatch('TambahDataProfile',data)
        },
        editdata(){
            let id = this.userid
            let isi = this.data
            let data =[id, isi]
            this.$store.dispatch('UpdateDataProfile',data)
        },
        upload(){
            let id = this.userid
            let gambar = this.profileimage
            let data =[id, gambar]

            this.$store.dispatch('UploadProfileFix',data)
            .then(response => {
                this.profileimage=''
                Toast.fire({
                    type: 'success',
                    title: 'Upload Image successfully'
                })
            })
        },
        EditdeleteImage(profileimage){
            let nama = profileimage
            let path = nama.substring(37) 
            // console.log(path)
            this.$store.dispatch('deleteEditProfilImage',path)
            .then(response => {
                this.profileimage = ''
            })
        },
        ProfileImageChange(e){
            const files = e.target.files || e.dataTransfer.files;

            if (!files.length) {
            console.log('no files');
            }

            let form = new FormData();
            for (var i = 0; i < files.length; i++) {
                let file = files.item(i);
                form.append('profile',file, file.name);
            }
            this.$store.dispatch('ProfileUpload', form)
            .then(response => {
            })
        }
    }
}
</script>

<style scoped>
.profil{
  background: rgb(250, 248, 231);
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
/* .tengah{ */
    /* padding-left: 20px; */
    /* text-align: center;
    vertical-align: middle; */
    /* margin: center; */
     /* position: absolute; */
/* } */
</style>

