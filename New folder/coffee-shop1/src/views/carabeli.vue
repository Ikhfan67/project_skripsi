<template>
    <div class="carabeli">
        <Navbar></Navbar>
        <div class="container mt-5 pt-5">
          <div class="pading">
            <div class="row">
              <div class="bawah">
                <b-media>
                  <p class="mb-0">
                    <b-list-group>
                      <b-list-group-item variant="dark">Pembayaran</b-list-group-item>
                      <b-list-group-item disabled>
                        <b-media>
                          <b-img slot="aside" src="http://localhost:5000/gambar_bank/mandiri_edited.jpg" width="64"></b-img>
                          <h6 class="mt-0">Cafe 3R.ASA</h6>
                          <p class="mb-0">
                            5003-5678-9101
                          </p>
                        </b-media>
                      </b-list-group-item>
                      <b-list-group-item>
                        <b-media>
                          <b-img slot="aside" src="http://localhost:5000/gambar_bank/bni.png" width="64"></b-img>
                          <h6 class="mt-0">Cafe 3R.ASA</h6>
                          <p class="mb-0">
                            5003-5678-9101
                          </p>
                        </b-media>
                      </b-list-group-item>
                    </b-list-group>
                  </p>
                </b-media>
              </div>
              <div class="col">
                <div class="kanankiri">
                  <h6>
                    <b-img slot="aside" src="http://localhost:5000/gambar_bank/logo.png" width="64" alt="placeholder"></b-img> 
                    Langkah-Langkah :
                  </h6>
                  <p>1.&nbsp;Buatlah akun dengan melakukan register di tombol login</p>
                  <p>2.&nbsp;Isi seluruh kolom dengan data anda yang sesuai dengan benar</p>
                  <p>3.&nbsp;Bila register telah berhasil, lakukan login untuk mendapatkan akses continue order</p>
                  <p>4.&nbsp;Anda dapat memilih produk di halaman home, dan pilih produk kopi sesuai keinginan anda</p>
                  <p>5.&nbsp;Produk akan masuk ke dalam keranjang belanja</p>
                  <p>6.&nbsp;Silahkan cek keranjang belanja, dan klik tombol continue order untuk ke tahap order</p>
                  <p>7.&nbsp;Batas pembayaran belanja anda maksimal 2 hari setelah anda melakukan order</p>
                  <p>8.&nbsp;Lakukan pembayaran sesuai dengan salah satu rekening bank yang tertera di web</p>
                  <p>9.&nbsp;Bila pembayaran telah di lakukan, upload bukti pembayaran di bawah tombol order </p>
                  <p>10.&nbsp;Bila bukti pembayaran telah di upload, status order akan berubah menjadi proses </p>
                  <p>11.&nbsp;Kami akan melakukan pengecekan dan status order anda akan diubah menjadi barang telah terkirim</p>
                  <p>12.&nbsp;Selamat berbelanja</p>
                </div>
              </div>
            </div>
            <div class="row">
                <b-media>
                  <p class="mb-0">
                    <b-list-group>
                      <b-list-group-item variant="dark">Pengiriman</b-list-group-item>
                      <b-list-group-item disabled>
                        <b-media>
                          <b-img slot="aside" src="http://localhost:5000/gambar_bank/jne.jpg" width="64"></b-img>
                          <h5 class="mt-0">JNE</h5>
                        </b-media>
                      </b-list-group-item>
                      <b-list-group-item>
                        <b-media>
                          <b-img slot="aside" src="http://localhost:5000/gambar_bank/jt_edited.jpg" width="64"></b-img>
                          <h5 class="mt-0">JnT</h5>
                        </b-media>
                      </b-list-group-item>
                    </b-list-group>
                  </p>
                </b-media>
            </div>
          </div>
            
            <!-- <b-col cols="9">
              <div>
                <b-card>
                    <h5 class="mt-0">Cara Pembelian Di Web Cafe 3R.ASA</h5>
                  <b-media>
                    <b-img slot="aside" src="http://localhost:5000/gambar_bank/logo.png" width="64" alt="placeholder"></b-img> -->
                    <!-- <div class="media">
                      <h6>Langkah-Langkah :</h6>
                    </div> -->
                    <!-- <div class="hero-content">
                      <div id="welcome-content" class="wow fadeInUp col span_6"> -->
                        <!-- <p>1.&nbsp;Buatlah akun dengan melakukan register di tombol login</p>
                        <p>2.&nbsp;Isi seluruh kolom dengan data anda yang sesuai dengan benar</p>
                        <p>3.&nbsp;Bila register telah berhasil, lakukan login untuk mendapatkan akses continue order</p>
                        <p>4.&nbsp;Anda dapat memilih produk di halaman home, dan pilihlah produk kopi sesuai keinginan anda</p>
                        <p>5.&nbsp;Produk akan masuk ke dalam keranjang belanja</p>
                        <p>6.&nbsp;Silahkan cek keranjang belanja, dan klik tombol continue order untuk ke tahap order</p>
                        <p>7.&nbsp;Batas pembayaran belanja anda maksimal 2 hari setelah anda melakukan order</p>
                        <p>8.&nbsp;Lakukan pembayaran sesuai dengan salah satu rekening bank yang tertera di web</p>
                        <p>9.&nbsp;Bila pembayaran telah di lakukan, upload bukti pembayaran di bawah tombol order </p>
                        <p>10.&nbsp;Bila bukti pembayaran telah di upload, status order akan berubah menjadi proses </p>
                        <p>11.&nbsp;Kami akan melakukan pengecekan dan status order anda akan diubah menjadi barang telah terkirim</p>
                        <p>12.&nbsp;Selamat berbelanja</p> -->
                      <!-- </div>
                    </div> -->
                  <!-- </b-media>
                </b-card>
              </div>
            </b-col> -->
          
        </div>
        <div><br></div>
        <footer>
          <p>&copy; 2019 <a href="https://ikhfannurali.blogspot.com" title="Cafe 3R.ASA">Cafe 3R.ASA.</a> All Rights Reserved. By : <a href="http://crwebstudio.com" target="_blank">Ikhfan Nur Ali</a></p>
        </footer>
    </div>
</template>
<script>
export default {
  name: 'CaraBeli',
  data () {
    return {
        
    }
  },
  computed:{
    
  },
  created(){
    
  },
  methods:{
    
  }
}
</script>

<style scoped>
.carabeli{
  background: rgb(250, 248, 231);
}
.pading{
  padding: 10px;
}
.kanankiri{
  text-align:justify;
}
.bawah{
  padding-bottom: 10px;
}
</style>

