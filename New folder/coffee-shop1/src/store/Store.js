import Vue from 'vue'
import Vuex from 'vuex'
import axios from 'axios'
import router from '@/router/index'


Vue.use(Vuex)
axios.defaults.baseURL = 'http://localhost:5000'

let cart = window.localStorage.getItem('cart')
let fix = window.localStorage.getItem('fix')

export const store = new Vuex.Store({
    state:{
        token :localStorage.getItem('access_token') || null,
        email: localStorage.getItem('email') || '',
        name: '',
        id: '',
        products:[],
        orders:[],
        productsSell:[],
        katalog: [],
        errorStatus : 0,
        images:[],
        orderimages:[],
        profileimages:[],
        editimage:'',
        profiles:[],
        Cek : false,
        cart: cart ? JSON.parse(cart) : [],
        fix: fix ? JSON.parse(fix) : [],
        idfix: '',
        gambarprofile:'',
        dataprofile:'',
        comments:[],
        details:[]
    },
    getters:{
        Editimage(state){
            return state.editimage
        },
        authStatus(state){
            return state.Cek
        },
        loggedIn(state){
            return state.token !== null
        },
        loggedAdmin(state){
            return state.Cek
        },
        loggedEmail(state){
            return state.email !== null
        },
        loggedName(state){
            return state.name !== null
        },
        ProductsFilter(state){
            return state.products
        },
        IdFixFilter(state){
            return state.idfix
        },
        OrdersFilter(state){
            return state.orders
        },
        ProductsSell(state){
            return state.productsSell
        },
        KatalogFilter(state){
            return state.katalog
        },
        getError(state){
            return state.errorStatus
        },
        ImagesFilter(state){
            return state.images
        },
        ProfileFilter(state){
            return state.profiles
        },
        ExportExcel(state){
            return state.exporturl
        },
    },
    mutations:{
        AddIdOrder(state, idorder){
            state.idfix = idorder
        },
        AddImageProfile(state, gambar){
            state.gambarprofile = gambar
        },
        AddDataProfile(state, data){
            state.dataprofile = data
        },
        EditAddImage(state, image){
            state.editimage = image
        },
        AddImage(state, image){
            state.images.push(image)
        },
        AddOrderImage(state, image){
            state.orderimages.push(image)
        },
        AddProfileImage(state,image){
            state.profileimages.push(image)
        },
        retrieveName(state, name){
            state.name = name
        },
        retrieveComments(state,data){
            state.comments = data
        },
        retrieveId(state, id){
            state.id = id
        },
        retrieveEmail(state, email){
            state.email = email
        },
        retrieveToken(state, token){
            state.token = token
        },
        destroyToken(state){
            state.Cek = false
            state.token = null
            state.email = ''
            state.name = ''
            state.id = ''
            state.cart = []
            state.fix = []
            state.katalog = []
        },
        destroycart(state){
            state.cart = []
            state.fix = []
        },
        CekAdmin(state, cekuser){
            state.cek = cekuser
        },
        retrieveByName(state, orders){
            state.katalog = orders
        },
        retrieveOrders(state, orders){
            state.orders = orders
        },
        retrieveProductsSell(state, orders){
            state.productsSell = orders
        },
        retrieveProducts(state, products){
            state.products = products
        },
        addProduct(state, data){
            state.products.push({
                id: data.id,
                name: data.name,
                description : data.description,
                price: data.price,
                image: data.image,
            })
        },
        gantiProduk(state, product){
         const index = state.products.findIndex(item => item.id == product.id);
            state.products[index] = product
        },
        deleteProduct(state, id){
            const index = state.products.findIndex(item => item.id == id);
            state.products.splice(index, 1)
        },
        deleteFile(state, id){
            const index = state.images.findIndex(item => item.id == id);
            state.images.splice(index, 1)
        },
        retrieveProfile(state, profiles){
            state.profiles = profiles
        },
        addProfile(state, profile){
            state.profiles.push({
                id: profile.id,
                name: profile.name,
                phone : profile.phone,
                address: profile.address,
                postcode: profile.postcode,
            })
        },
        addOrder(state, cart){
            state.carts.push({
                name : cart.name,
                image : cart.image,
                price : cart.price
            })
        },
        gantiProfile(state, profile){
            const index = state.profiles.findIndex(item => item.id == profile.id);
            state.profiles[index] = profile
        },
        deleteProfile(state, id){
            const index = state.profiles.findIndex(item => item.id == id);
            state.profiles.splice(index, 1)
        },
        setError(state, id){
            state.errorStatus = id
        },
        setCek(state, data){
            state.Cek = data
        },
        addToCart(state, item){
            // let found = state.cart.find(products => products.id_produk = item.id_produk);
            let found = state.cart.find(function(product){
                return product.id_produk == item.id_produk
            })  

            if(typeof found !== "undefined"){
                found.quantity++;
            }else{
                state.cart.push(item)
            }
            this.commit('saveData')
        },
        fix(state, item){
            // let found = state.cart.find(products => products.id_produk = item.id_produk);
            let found = state.fix.find(function(product){
                return product.id_produk == item.id_produk
            })  

            if(typeof found !== "undefined"){
                found.quantity++;
            }else{
                state.fix.push(item)
            }
            this.commit('saveFix')
        },
        saveFix(state){
            window.localStorage.setItem('fix', JSON.stringify(state.fix))
        },
        saveData(state){
            window.localStorage.setItem('cart', JSON.stringify(state.cart))
        },
        removeFromCart(state,item){
            let index = state.cart.indexOf(item)
            let index1 = state.fix.indexOf(item)
            state.cart.splice(index,1)
            state.fix.splice(index1,1)
            this.commit('saveData')
            this.commit('saveFix')
        },
        increment(state, id){
            let found = state.cart.find(product => product.id_produk == id );
            let found1 = state.fix.find(product => product.id_produk == id );
            
            if(found && found1){
              found.quantity++;
              found1.quantity++;
            }else{
              state.cart.push(item);
              state.fix.push(item);
            }
            this.commit('saveData');
            this.commit('saveFix');
        },
        decrement(state, id){

            let found = state.cart.find(product => product.id_produk == id );
            let found1 = state.fix.find(product => product.id_produk == id );
    
            if(found && found1){
              found.quantity--;
              found1.quantity--;
            }
            if (found.quantity < 1 && found1.quantity < 1) {
              this.removeFromCart(state, found);          
              this.removeFromCart(state, found1);          
            }
            this.commit('saveData');
            this.commit('saveFix');
    
        },
        addDetail(state, detail){
            state.details = detail
        }
        
    },
    actions:{
        exportExcel(context,data){
            var config = {
                headers: {  'Access-Control-Allow-Origin': '*',
                            Authorization: localStorage.getItem('access_token'),
                            'Content-Type': 'application-json',    
                         }
            };
            var dataa={
                datesatu:data.tgl1,
                datetwo:data.tgl2
            }
            var tgl = data.tgl1
            var tglb = data.tgl2
            // console.log(dataa)
            axios.post('/test/'+ tgl +'/'+ tglb, config,dataa)
            .then(response => {
                window.location.href= "http://localhost:5000" + response.data.url
            })
            .catch(error => {
                console.log(error)
            })
        },
        deleteOrder(context, id){
            var config = {
                headers: {  'Access-Control-Allow-Origin': '*',
                            Authorization: localStorage.getItem('access_token'),
                            'Content-Type': 'application-json',    
                         }
            };
            axios.delete('/edit/order/delete/'+ id, config)
            .then(response => {
                // context.commit('deleteProduct', response.data.message)
                console.log(response.data.message)
            })
            .catch(error => {
                console.log(error)
            })
        },
        UbahDataOrder(context,data){
            var config = {
                headers: {  'Access-Control-Allow-Origin': '*',
                            Authorization: localStorage.getItem('access_token'),
                            'Content-Type': 'application-json',    
                         }
            };
            var id = data[3]
            axios.put('/edit/order/'+id,{
                tglorder: data[0],
                total : parseInt(data[2]),
                status : data[1]
            },config)
            .then(response => {
                // console.log(response.data.message)
                // context.commit('retrieveByName', response.data.result)
            })
            .catch(error => {
                console.log(error)
            })
        },
        retrieveComments(context,data){
            var config = {
                headers: {  'Access-Control-Allow-Origin': '*',
                            Authorization: localStorage.getItem('access_token'),
                            'Content-Type': 'application-json',    
                         }
            };
            axios.get('/komentar/get', config)
            .then(response => {
                var data = response.data.result
                // console.log(response.data.result)
                context.commit('retrieveComments', data)
            })
            .catch(error => {
                console.log(error)
            })
        },
        ImageOrder(context, data){
            var config = {
                headers: {  'Access-Control-Allow-Origin': '*',
                            Authorization: localStorage.getItem('access_token'),
                            'Content-Type': 'application-json',    
                         }
            };
            var no = data[0]
            var gambar = data[1]
            
            axios.put('/order/edit/'+ no,{
                status: 'Sudah Upload Bukti Transfer',
                image_order : gambar
            },config)

            .then(response => {
                // context.commit('gantiProduk', product)
                console.log(response)
            })
            .catch(error => {
                console.log(error)
            })
        },
        selectByName(context, data){
            var config = {
                headers: {  'Access-Control-Allow-Origin': '*',
                            Authorization: localStorage.getItem('access_token'),
                            'Content-Type': 'application-json',    
                         }
            };
            // console.log(data)
            axios.get('/order/get/'+ data, config)
            .then(response => {
                // console.log(response.data.result)
                context.commit('retrieveByName', response.data.result)
            })
            .catch(error => {
                console.log(error)
            })
        },
        retrieveProductsSell(context){
            var config = {
                headers: {  'Access-Control-Allow-Origin': '*',
                            Authorization: localStorage.getItem('access_token'),
                            'Content-Type': 'application-json',    
                         }
            };
            axios.get('/order/produksell',config)
            .then(response => {
                context.commit('retrieveProductsSell', response.data.result)
            })
            .catch(error => {
                console.log(error)
            })
        },
        retrieveOrder(context){
            var config = {
                headers: {  'Access-Control-Allow-Origin': '*',
                            Authorization: localStorage.getItem('access_token'),
                            'Content-Type': 'application-json',    
                         }
            };
            axios.get('/order/all',config)
            .then(response => {
                // console.log(response.data.result)
                // console.log(response.data.result)
                context.commit('retrieveOrders', response.data.result)
            })
            .catch(error => {
                console.log(error)
            })
        },
        order(context,data){
            const produk = context.state.fix
            var data = {
                // tglorder: data[1].tglorder,
                iduser: parseInt(data[1].iduser),
                produk : produk,
                status : 'Belum Bayar'
            }
            return new Promise((resolve, reject) => {
                axios.post('/order/add', data)
                .then(response => {
                    
                    const idorder = response.data.idOrder
                    context.commit('AddIdOrder', idorder)

                })
                .catch(error => {
                    reject(error)
                })
            })
        },
        UploadProfileFix(context,data){
            
            var id = parseInt(data[0])
            var gambar1 = data[1]
            console.log(id)
            console.log(gambar1)

                axios.put('/uploadprofilefix/'+ id, {
                    gambar: gambar1})
                .then(response => {
                    const gambar = response.data
                    // console.log(response.data)
                    // context.commit('AddImageProfile', gambar)
                })
                .catch(error => {
                    console.log(error)
                })
            
        },
        GetImageProfile(context,data){
            var id1 = parseInt(data[0])

            axios.get('/getimageprofile/'+ id1)
            .then(response => {
                const gambar = response.data.result[0].image
                // console.log(response.data.result[0].image)
                context.commit('AddImageProfile', gambar)
            })
            .catch(error => {
                console.log(error)
            })
            
        },
        GetDataProfile(context,data){
            var id1 = parseInt(data[0])

            axios.get('/getdataprofile/'+ id1)
            .then(response => {
                const data = response.data.result[0]
                // console.log(response.data.result[0])
                context.commit('AddDataProfile', data)
            })
            .catch(error => {
                console.log(error)
            })
            
        },
        UpdateDataProfile(context,data){
            var id = parseInt(data[0])
            var data = data[1]

            axios.put('/updatedataprofile',{
                id_user : id,
                name : data.name,
                email : data.email,
                password : data.password,
                phone : data.phone,
                address : data.address,
                postcode : data.postcode,
            })
            .then(response => {
                const data = response.data.result
                console.log(response.data)
                // context.commit('AddDataProfile', data)
            })
            .catch(error => {
                console.log(error)
            })
        },
        deleteEditFileImage(context,data){
            var config = {
                headers: {  'Access-Control-Allow-Origin': '*',
                             Authorization: localStorage.getItem('access_token'),
                            'Content-Type': 'application-json',    
                         }
            };
            axios.delete('/image/removeEditImage/'+data,config)
            .then(response => {
                console.log(response.data.message)
                // context.commit('deleteEditFile', response.data.message)
            })
            .catch(error => {
                console.log(error)
            })
        },
        deleteEditUploadImage(context,data){
            var config = {
                headers: {  'Access-Control-Allow-Origin': '*',
                             Authorization: localStorage.getItem('access_token'),
                            'Content-Type': 'application-json',    
                         }
            };
            axios.delete('/removeUploadImage/'+data,config)
            .then(response => {
                // console.log(response.data.message)
                // context.commit('deleteEditFile', response.data.message)
            })
            .catch(error => {
                console.log(error)
            })
        },
        deleteEditProfilImage(context,data){
            var config = {
                headers: {  'Access-Control-Allow-Origin': '*',
                             Authorization: localStorage.getItem('access_token'),
                            'Content-Type': 'application-json',    
                         }
            };
            axios.delete('/removeProfilImage/'+data,config)
            .then(response => {
                // console.log(response.data.message)
                // context.commit('deleteEditFile', response.data.message)
            })
            .catch(error => {
                console.log(error)
            })
        },
        deleteFileImage(context,data){
            var config = {
                headers: {  'Access-Control-Allow-Origin': '*',
                             Authorization: localStorage.getItem('access_token'),
                            'Content-Type': 'application-json',    
                         }
            };
            axios.delete('/image/removeImage/'+data,config)
            .then(response => {
                context.commit('deleteFile', response.data.message)
            })
            .catch(error => {
                console.log(error)
            })
        },
        OrderUpload(context, data){
            var config = {
                headers: {  'Access-Control-Allow-Origin': '*',
                            'Content-Type': 'multipart/form-data',    
                            Authorization: localStorage.getItem('access_token'),
                         }
            };
            axios.post('/order/orderuploadimage', data, config)
            .then(res =>{
                const image = res.data
                context.commit('AddOrderImage', image)
            })
            .catch(error => {
                console.log(error)
            })
        },
        ProfileUpload(context, data){
            var config = {
                headers: {  'Access-Control-Allow-Origin': '*',
                            'Content-Type': 'multipart/form-data',    
                            Authorization: localStorage.getItem('access_token'),
                         }
            };
            axios.post('/profileuploadimage', data, config)
            .then(res =>{
                const image = res.data
                context.commit('AddProfileImage', image)
            })
            .catch(error => {
                console.log(error)
            })
        },
        EditonUpload(context, data){
            var config = {
                headers: {  'Access-Control-Allow-Origin': '*',
                            'Content-Type': 'multipart/form-data',    
                            Authorization: localStorage.getItem('access_token'),
                         }
            };
            axios.post('/product/edituploadimage', data, config)
            .then(res =>{
                const image = res.data.url
                console.log(image)
                // context.commit('AddImage', image)
                context.commit('EditAddImage', image)
            })
            .catch(error => {
                console.log(error)
            })
        },
        onUpload(context, data){
            var config = {
                headers: {  'Access-Control-Allow-Origin': '*',
                            'Content-Type': 'multipart/form-data',    
                            Authorization: localStorage.getItem('access_token'),
                         }
            };
            axios.post('/product/uploadimage', data, config)
            .then(res =>{
                // console.log(res.data)
                const image = res.data
                context.commit('AddImage', image)

            })
            .catch(error => {
                console.log(error)
            })
        },
        retrieveToken(context, credentials){
            return new Promise((resolve, reject) => {
                axios.post('auth/login',{
                    email : credentials.email,
                    password : credentials.password,
                })
                .then(response => {
                    const token = response.data.access_token
                    const email = response.data.access_email
                    const name = response.data.access_name
                    const id = response.data.access_id
                    
                    localStorage.setItem('access_token', token)
                    localStorage.setItem('email', email)
                    localStorage.setItem('name', name)
                    localStorage.setItem('id', id)
                    context.commit('retrieveToken', token)
                    context.commit('retrieveEmail', email)
                    context.commit('retrieveName', name)
                    context.commit('retrieveId', id)
                    resolve(response)
                })
                .catch(error => {
                   context.commit('setError', 1)
                })
            })
        },
        register(context, data){
            return new Promise((resolve, reject) => {
                axios.post('/auth/register',{
                    name: data.name,
                    email: data.email,
                    password: data.password,
                })
                .then(response => {
                    resolve(response)
                })
                .catch(error => {
                    reject(error)
                })
            })
        },
        komentar(context, data){
            return new Promise((resolve, reject) => {
                axios.post('/komentar/add',{
                    name: data.name,
                    email: data.email,
                    phone: data.phone,
                    message: data.message,
                })
                .then(response => {
                    resolve(response)
                })
                .catch(error => {
                    reject(error)
                })
            })
        },
        destroyToken(context){
            localStorage.removeItem('access_token')
            localStorage.setItem('email', '')
            localStorage.setItem('name', '')
            localStorage.setItem('id', '')
            localStorage.setItem('cart', [])
            localStorage.setItem('fix', [])
            context.commit('destroyToken')
        },
        destroycart(context){
            localStorage.setItem('cart', [])
            localStorage.setItem('fix', [])
            context.commit('destroycart')
        },
        saveDataProduk(context, product){
            var config = {
                headers: {  'Access-Control-Allow-Origin': '*',
                            Authorization: localStorage.getItem('access_token'),
                            'Content-Type': 'application-json',    
                         }
            };
            axios.post('/product/add',{
                // id: parseInt(product.id),
                name: product.name,
                description: product.description,
                price: parseInt(product.price),
                image : "http://localhost:5000/image/"+ product.image
            },config)
            .then(response => {
                // console.log(response.data)
                // produk = response.data
                context.commit('addProduct', response.data)
            })
            .catch(error => {
                console.log(error)
            })
        },
        retrieveProducts(context){
            var config = {
                headers: {  'Access-Control-Allow-Origin': '*',
                            Authorization: localStorage.getItem('access_token'),
                            'Content-Type': 'application-json',    
                         }
            };
            axios.get('/product/all',config)
            .then(response => {
                // console.log(response.data.result)
                context.commit('retrieveProducts', response.data.result)
            })
            .catch(error => {
                console.log(error)
            })
        },
        retrieveProductsHome(context){
            var config = {
                headers: {  'Access-Control-Allow-Origin': '*',
                            Authorization: localStorage.getItem('access_token'),
                            'Content-Type': 'application-json',    
                         }
            };
            axios.get('/product/all/home',config)
            .then(response => {
                context.commit('retrieveProducts', response.data.result)
            })
            .catch(error => {
                console.log(error)
            })
        },
        deleteProduct(context, id){
            var config = {
                headers: {  'Access-Control-Allow-Origin': '*',
                            Authorization: localStorage.getItem('access_token'),
                            'Content-Type': 'application-json',    
                         }
            };
            axios.delete('/product/'+ id, config)
            .then(response => {
                context.commit('deleteProduct', response.data.message)
            })
            .catch(error => {
                console.log(error)
            })
        },
        deleteKomen(context, id){
            var config = {
                headers: {  'Access-Control-Allow-Origin': '*',
                            Authorization: localStorage.getItem('access_token'),
                            'Content-Type': 'application-json',    
                         }
            };
            axios.delete('/komentar/delete/'+ id, config)
            .then(response => {
                // context.commit('deleteProduct', response.data.message)
                // console.log(response)
            })
            .catch(error => {
                console.log(error)
            })
        },
        updateDataProduct(context, product){
            var config = {
                headers: {  'Access-Control-Allow-Origin': '*',
                            Authorization: localStorage.getItem('access_token'),
                            'Content-Type': 'application-json',    
                         }
            };
            axios.put('/product/edit/'+ product.id,{
                name: product.name,
                description: product.description,
                price: parseInt(product.price),
                image: product.image
            },config)

            .then(response => {
                context.commit('gantiProduk', product)
            })
            .catch(error => {
                console.log(error)
            })
        },
        retrieveProfile(context){
            var config = {
                headers: {  'Access-Control-Allow-Origin': '*',
                            Authorization: localStorage.getItem('access_token'),
                            'Content-Type': 'application-json',    
                         }
            };
            axios.get('/profile/all',config)
            .then(response => {
                context.commit('retrieveProfile', response.data.result)
            })
            .catch(error => {
                console.log(error)
            })
        },
        updateDataProfile(context, profile){
            var config = {
                headers: {  'Access-Control-Allow-Origin': '*',
                            Authorization: localStorage.getItem('access_token'),
                            'Content-Type': 'application-json',    
                         }
            };
            axios.put('/profile/edit/'+ profile.id,{
                name: profile.name,
                phone: profile.phone,
                address: profile.address,
                postcode: profile.postcode,
                role: profile.role,
            },config)

            .then(response => {
                context.commit('gantiProfile', profile)
            })
            .catch(error => {
                console.log(error)
            })
        },
        TambahDataProfile(context, data){
            var id = data[0]
            var data = data[1]
            console.log(data)
            axios.post('/tambahdataprofile',{
                id_user: parseInt(id),
                name: data.name,
                phone: data.phone ,
                address: data.address,
                postcode: data.postcode,
                email : data.email,
                password : data.password,

            })

            .then(response => {
                console.log(response.result)
                // context.commit('gantiProfile', profile)
            })
            .catch(error => {
                console.log(error)
            })
        },
        resetPasswordEmail(context, profile){
            var config = {
                headers: {  'Access-Control-Allow-Origin': '*',
                            Authorization: localStorage.getItem('access_token'),
                            'Content-Type': 'application-json',    
                         }
            };
            axios.put('/profile/reset/'+ profile.id,{
                email: "",
                password: "",
            },config)

            .then(response => {
                context.commit('gantiProfile', profile)
            })
            .catch(error => {
                console.log(error)
            })
        },
        updateDataAccount(context, profile){
            var config = {
                headers: {  'Access-Control-Allow-Origin': '*',
                            Authorization: localStorage.getItem('access_token'),
                            'Content-Type': 'application-json',    
                         }
            };
            axios.put('/profile/new/'+ profile.id,{
                name: profile.name,
                email: profile.nemail,
                password: profile.npassword,
            },config)

            .then(response => {
                context.commit('gantiProfile', profile)
            })
            .catch(error => {
                console.log(error)
            })
        },
        deleteProfile(context, id){
            var config = {
                headers: {  'Access-Control-Allow-Origin': '*',
                            Authorization: localStorage.getItem('access_token'),
                            'Content-Type': 'application-json',    
                         }
            };
            axios.delete('/profile/'+ id, config)
            .then(response => {
                context.commit('deleteProfile', response.data.message)
            })
            .catch(error => {
                console.log(error)
            })
        },
        CekAdmin(context){
            var config = {
                headers: {  'Access-Control-Allow-Origin': '*',
                            Authorization: localStorage.getItem('access_token'),
                            'Content-Type': 'application-json',    
                         }
            };
            axios.get('/login/cek', config)
            .then(response =>{
                context.commit("setCek", true)
                router.push({
                    name: 'admin'
                })
            }).catch(error =>{
                router.push({
                    name: 'home'
                })
            })
        },
        addToCart(context, cart){
            var config = {
                headers: {  'Access-Control-Allow-Origin': '*',
                            Authorization: localStorage.getItem('access_token'),
                            'Content-Type': 'application-json',    
                         }
            };
            axios.post('/order/add', config,{
                name : cart.item.productName,
                image: cart.item.productImage,
                price: cart.item.productPrice
            })
            .then(response => {
                context.commit('addOrder', cart)
            })
            .catch(error => {
                console.log(error)
            })
        },
        detailsProductSells(context, data){
            var config = {
                headers: {  'Access-Control-Allow-Origin': '*',
                            Authorization: localStorage.getItem('access_token'),
                            'Content-Type': 'application-json',    
                         }
            };
            // console.log(data.name)
            axios.post('/details/product/'+ data.name, config)
            .then(response => {
                // console.log(response.data.result)
                var detail = response.data.result
                context.commit('addDetail', detail)
            })
            .catch(error => {
                console.log(error)
            })
        }
        
    }
});