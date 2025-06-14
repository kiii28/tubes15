Aplikasi Simulasi Pasar Saham Virtual 


Deskripsi: 

Aplikasi ini memungkinkan pengguna untuk mensimulasikan perdagangan saham 
dengan menggunakan saldo virtual. Data utama yang digunakan adalah daftar 
saham, harga saham yang berubah, dan portofolio pengguna. Pengguna aplikasi 
adalah individu yang ingin belajar cara trading saham tanpa risiko nyata. 
Spesifikasi: ...

a. Pengguna dapat menambahkan, mengubah, dan menghapus transaksi 
pembelian dan penjualan saham. 
Diimplementasikan pada fungsi:

beli(), 
jual(), 
hapusSahamPengguna(). 



b. Sistem menghitung perubahan nilai portofolio berdasarkan fluktuasi harga 
saham. 

diimplementasikan pada fungsi:
simulasiTrading(). 

c. Pengguna dapat mencari saham berdasarkan kode atau nama perusahaan 
menggunakan Sequential dan Binary Search. 
Diimplementasikan pada fungsi:

//sebagai binary search \n
cariKodeSaham(),   
//sebagai sequential search
cariKodeSahamPengguna().    



d. Pengguna dapat mengurutkan saham berdasarkan harga tertinggi atau volume 
transaksi menggunakan Selection dan Insertion Sort. 
Diimplementasikan pada fungsi:

//ascending menggunakan insetion sort   
ascendingInsertionsortHargaSa(), 
ascendingInsertionsortVolume(), 

//descending menggunakan selection sort   
urutDataDescedingHarga(), 
urutDataDescedingVolume(), 

//untuk mengurutkan data kembali ke setingan awal
urutDataAwal().



e. Sistem menampilkan statistik keuntungan dan kerugian pengguna dalam 
simulasi trading.
Diimplementasikan pada fungsi:
tampilkanStatistikUntungRugi().  



