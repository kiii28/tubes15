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

beli()
jual()
hapusSahamPengguna(index int)



b. Sistem menghitung perubahan nilai portofolio berdasarkan fluktuasi harga 
saham. 

diimplementasikan pada fungsi:
simulasiTrading() 

c. Pengguna dapat mencari saham berdasarkan kode atau nama perusahaan 
menggunakan Sequential dan Binary Search. 
Diimplementasikan pada fungsi:

cariKodeSaham(a tabsa, n int, x string)   //sebagai binary search
cariKodeSahamPengguna(dataSaham tabpe, searchKey string)   //sebagai sequential search



d. Pengguna dapat mengurutkan saham berdasarkan harga tertinggi atau volume 
transaksi menggunakan Selection dan Insertion Sort. 
Diimplementasikan pada fungsi:

//ascending menggunakan insetion sort
ascendingInsertionsortHargaSa(a *tabsa, n int)
ascendingInsertionsortVolume(a *tabsa, n int)

//descending menggunakan selection sort
urutDataDescedingHarga(data *tabsa, ndata int)
urutDataDescedingVolume(data *tabsa, ndata int)
urutDataAwal(data *tabsa, ndata int)   //untuk mengurutkan data kembali ke setingan awal



e. Sistem menampilkan statistik keuntungan dan kerugian pengguna dalam 
simulasi trading.
Diimplementasikan pada fungsi:
tampilkanStatistikUntungRugi()



