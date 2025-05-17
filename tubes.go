package main

import (
	"fmt"
	"math/rand"
)

/* kita ingin membuat sebuah applikasi untuk simulasi pasar saham, dengan tampilan awal dengan beberapa pilihan.
Memperlihatkan daftar saham, beli seham, jual saham, menampilkan portofolio, menampilkan statistik untung/rugi
kita beranggapan bahwa tiap memjalankan applikasi ini adalah orang yang berbeda
*/

type pengguna struct {
	nama, SahamDimiliki            string
	nilaiSahamDimiliki, untungRugi float64
}

type transaksi struct {
	Pembelian, penjualan, lot, lembarSaham int // Lot adlah minimal pembelian saham, 1 Lot = 100 Lembar saham
	namaTransaksi, idsaham                 string
}

type saham struct {
	id, namasaham                  string
	volume, hargaSaham, nilaiSaham float64
}

type tabtr [150]transaksi //array transaksi
type tabpe [15]pengguna   //array pengguna
type tabsa [10]saham      // array saham

var daftarSaham tabsa

var Transaksi tabtr
var jumlahTransaksi int = 0

var portofolio tabpe
var jumlahPortofolio int = 0

var saldo float64 = 10000000 //karena trading saham tanpa resiko nyata, asumsikan kita menyedikan uang virtual sebagai bahan pembelajaran

func inisialisasiDataSaham() {
	daftarSaham[0] = saham{id: "BBCA", namasaham: "Bank Central Asia", hargaSaham: 9275, volume: 10000, nilaiSaham: 6788889} //Harga Saham dalam 1 lebar saham
	daftarSaham[1] = saham{id: "BBRI", namasaham: "Bank rakyat Indonesia", hargaSaham: 4260, volume: 10000, nilaiSaham: 6788889}
	daftarSaham[2] = saham{id: "BMRI", namasaham: "Bank Mandiri", hargaSaham: 5275, volume: 10000, nilaiSaham: 6788889}
	daftarSaham[3] = saham{id: "TLKM", namasaham: "Telkom Indonesia", hargaSaham: 2660, volume: 10000, nilaiSaham: 6788889}
	daftarSaham[4] = saham{id: "ASII", namasaham: "Astra Indonesia", hargaSaham: 4820, volume: 10000, nilaiSaham: 6788889}
	daftarSaham[5] = saham{id: "UNVR", namasaham: "Uniliver Indonesia", hargaSaham: 4500, volume: 10000, nilaiSaham: 6788889}
	daftarSaham[6] = saham{id: "ICBP", namasaham: "Indofood CBP Sukses Makmur", hargaSaham: 10000, volume: 10000, nilaiSaham: 6788889}
	daftarSaham[7] = saham{id: "ADRO", namasaham: "Adaro Energy Indonesia", hargaSaham: 3000, volume: 10000, nilaiSaham: 6788889}
	daftarSaham[8] = saham{id: "AKRA", namasaham: "AKR Corpindo", hargaSaham: 1300, volume: 10000, nilaiSaham: 6788889}
	daftarSaham[9] = saham{id: "ACES", namasaham: "Ace Hardware Indonesia", hargaSaham: 545, volume: 10000, nilaiSaham: 6788889}
}

func tampilkanDaftarSaham() { // Daftar-daftar saham yang dapat dibeli
	var data tabsa
	fmt.Println("-----------------------------------------------------------------------------------")
	fmt.Printf("| %-5s | %-25s | Rp%-15s | %-15s | %-15s |\n", "KODE", "NAMA", "HARGA SAHAM", "VOLUME", "NILAI SAHAM")
	fmt.Println("-----------------------------------------------------------------------------------")
	for i := 0; i < 10; i++ {
		if data[i].id != "" {
			fmt.Printf("| %-5s | %-25s | Rp%-15f | %-15f | %-15f |\n", data[i].id, data[i].namasaham, data[i].hargaSaham, data[i].volume, data[i].nilaiSaham)
		}
	}
	fmt.Println("-----------------------------------------------------------------------------------")
}

func beli() { //beli: nama saham, kode saham, beli berdasarkan lembar saham(bukan nilai dlm rupiah) dan pastikan saldo masih cukup,

}
func jual() {

}
func portofolioPengguna(nama, sahamDimiliki, nilaiSHDimiliki, untungRugi tabpe) { //input nama pengguna lalu buat var saham dimiliki, nilai saham, keuntungan, kerugian

}
func tampilkanPortofolioPengguna() {

}
func tampilkanStatistikUntungRugi() {

}
func hitungUntungRugi() {
	//kurang lebih saldo yang "sudah dibelikan saham" akan dihitung,
	//pertambahan nilai atau pengurangan nilai dari saham setelah melakukan simulasi trading

}
func simulasiTrading() { //n untuk berapa lama menahan
	/* idenya adalah asumsikan bahwa naik turun harga saham memiliki kemungkinan untung 3 : 1 rugi
	anggap kita mempunya angka random 1-5 dengan jika dapat 1 tidak ada perubahan jika dapat 2 Harga turun sekian dalam %
	jika dapat 3-5 harga naik sekian dalam %. anggap tiap kenaikkan adalah dalam rentang 1 bulan sehingga masukkan
	misal 4 jadi kita seperti perkembangan harga saham dalam 4 bulan atau 4 kali naik / turun */
	var n int
	fmt.Println("Simulasi pergerakan harga saham sebanyak 1 perbulan, ingin menahan berapa lama: ")
	fmt.Scan(&n)
	for i := 1; i <= n; i++ {
		// random 1–5
		switch rand.Intn(6) + 1 { // + 1 karena dari nol
		case 1:
			// tidak ada perubahan
		case 2:
			daftarSaham[i].hargaSaham *= 0.98 // turun 2%
			daftarSaham[i].nilaiSaham *= 0.98
		case 3:
			daftarSaham[i].hargaSaham *= 1.03 // naik 3%
			daftarSaham[i].nilaiSaham *= 1.03
		case 4:
			daftarSaham[i].hargaSaham *= 1.04 // naik 4%
			daftarSaham[i].nilaiSaham *= 1.04
		case 5:
			daftarSaham[i].hargaSaham *= 1.05 // naik 5%
			daftarSaham[i].nilaiSaham *= 1.05
		case 6:
			daftarSaham[i].hargaSaham *= 0.97 // turun 3%
			daftarSaham[i].nilaiSaham *= 0.97
		}
	}
	fmt.Println("Harga telah diperbarui.")
}

func urutDataDescedingHarga(data *tabsa, ndata int) { //untuk mengurutkan data saham dari besar ke kecil berdasarkan Harga saham
	var i, idx, pass int
	var temp saham //penyimpanan data sementara
	pass = 1
	for pass < ndata {
		idx = pass - 1
		i = pass
		for i < ndata {
			if (*data)[i].hargaSaham > (*data)[idx].hargaSaham {
				idx = i
			}
			i = i + 1
		}
		temp = (*data)[pass-1]
		(*data)[pass-1] = (*data)[idx]
		(*data)[idx] = temp
		pass = pass + 1
	}
}

func urutDataDescedingVolume(data *tabsa, ndata int) { //untuk mengurutkan data saham dari besar ke kecil berdasarkan volume
	var i, idx, pass int
	var temp saham //penyimpanan data sementara
	pass = 1
	for pass < ndata {
		idx = pass - 1
		i = pass
		for i < ndata {
			if (*data)[i].volume > (*data)[idx].volume {
				idx = i
			}
			i = i + 1
		}
		temp = (*data)[pass-1]
		(*data)[pass-1] = (*data)[idx]
		(*data)[idx] = temp
		pass = pass + 1
	}
}

func pilihan() { //pengguna memilih akan melakukan apa
	for {
		var pilih int
		fmt.Printf("Pilih hal apa yang akan anda lakukan: \n1. Tampilkan daftar saham. \n2. Tampilkan portofolio anda. \n3. Beli saham. \n4. Jual saham. \n5. Tampilkan statistik keuntungan/kerugian. \n6. Melakukan simulasi saham. \n7. Keluar.")
		fmt.Scan(&pilih)
		switch pilih {
		case 1: //memanggil fungsi daftar saham
			tampilkanDaftarSaham()
		case 2: //fungsi porto

		case 3: //

		case 4:

		case 5:

		case 6:
			simulasiTrading()
		default:
			return
		}
		fmt.Println()
	}
}

func main() {
	inisialisasiDataSaham()
	pilihan()

}
