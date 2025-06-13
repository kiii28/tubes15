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
	nama, SahamDimiliki                                     string //saham yang disimpan dengan kode
	nilaiSahamDimiliki, totalNilaiSahamDimiliki, untungRugi float64
	lembarSahamDimiliki, indeksAsli                         int //indeks asli untuk shorting kembali ke awal

}

type saham struct {
	id, namasaham                  string
	volume, hargaSaham, nilaiSaham float64
	indeksAsli                     int
}
type statistik struct {
	totalKeuntungan, totalKerugian, naikHarga, turunHarga float64
}

type tabpe [15]pengguna //array pengguna
type tabsa [10]saham    // array saham
type tabst [1]statistik
var nData int = 10
var daftarSaham tabsa
var sahamPunya int = 0

var statstk tabst
var jumlahTransaksi int = 0

var portofolio tabpe
var jumlahPortofolio int = 0
var indeksBenar int = 0
var saldo float64 = 10000000 //karena trading saham tanpa resiko nyata, asumsikan kita menyedikan uang virtual sebagai bahan pembelajaran

func inisialisasiDataSaham() {
	daftarSaham[0] = saham{id: "BBCA", namasaham: "Bank Central Asia", hargaSaham: 9275, volume: 1800000, nilaiSaham: 16695000000000, indeksAsli: 0} //Harga Saham dalam 1 lebar saham
	daftarSaham[1] = saham{id: "BBRI", namasaham: "Bank rakyat Indonesia", hargaSaham: 4260, volume: 2300000, nilaiSaham: 9798000000000, indeksAsli: 1}
	daftarSaham[2] = saham{id: "BMRI", namasaham: "Bank Mandiri", hargaSaham: 5275, volume: 1650000, nilaiSaham: 8703750000000, indeksAsli: 2}
	daftarSaham[3] = saham{id: "TLKM", namasaham: "Telkom Indonesia", hargaSaham: 2660, volume: 2750000, nilaiSaham: 7315000000000, indeksAsli: 3}
	daftarSaham[4] = saham{id: "ASII", namasaham: "Astra Indonesia", hargaSaham: 4820, volume: 2750000, nilaiSaham: 13255000000000, indeksAsli: 4}
	daftarSaham[5] = saham{id: "UNVR", namasaham: "Uniliver Indonesia", hargaSaham: 4500, volume: 800000, nilaiSaham: 3600000000000, indeksAsli: 5}
	daftarSaham[6] = saham{id: "ICBP", namasaham: "Indofood CBP Sukses Makmur", hargaSaham: 10000, volume: 670000, nilaiSaham: 6700000000000, indeksAsli: 6}
	daftarSaham[7] = saham{id: "ADRO", namasaham: "Adaro Energy Indonesia", hargaSaham: 3000, volume: 3450000, nilaiSaham: 10350000000000, indeksAsli: 7}
	daftarSaham[8] = saham{id: "AKRA", namasaham: "AKR Corpindo", hargaSaham: 1300, volume: 1120000, nilaiSaham: 1456000000000, indeksAsli: 8}
	daftarSaham[9] = saham{id: "ACES", namasaham: "Ace Hardware Indonesia", hargaSaham: 545, volume: 1234500, nilaiSaham: 672352500000, indeksAsli: 9}
}

func tampilkanDaftarSaham() { // Daftar-daftar saham yang dapat dibeli
	var pilih int
	fmt.Println("---------------------------------------------------------------------------------------------------")
	fmt.Printf("| %-5s | %-27s | Rp %-15s | %-15s | %-18s |\n", "KODE", "NAMA", "HARGA SAHAM", "VOLUME", "NILAI SAHAM")
	fmt.Println("---------------------------------------------------------------------------------------------------")
	for i := 0; i < 10; i++ {
		if daftarSaham[i].id != "" { //memastikan yang di print hanya array yang memiliki nilai
			fmt.Printf("| %-5s | %-27s | Rp %-15.2f | %-15.2f | %-18.2f |\n", daftarSaham[i].id, daftarSaham[i].namasaham, daftarSaham[i].hargaSaham, daftarSaham[i].volume, daftarSaham[i].nilaiSaham)
		}
	}
	fmt.Println("---------------------------------------------------------------------------------------------------")

	fmt.Printf("Pilih: \n1. Sorting mengecil berdasarkan Harga. \n2. Sorting mengecil berdasarkan Volume. \n3. Sorting membesar berdasarkan Harga.\n4. Sorting membesar berdasarkan Volume. \n5. Kembali.\n")
	fmt.Scan(&pilih)
	if pilih == 1 {
		urutDataDescedingHarga(&daftarSaham, nData)
		tampilkanDaftarSaham()
	} else if pilih == 2 {
		urutDataDescedingVolume(&daftarSaham, nData)
		tampilkanDaftarSaham()
	} else if pilih == 3 {
		ascendingInsertionsortHargaSa(&daftarSaham, nData)
		tampilkanDaftarSaham()
	} else if pilih == 4 {
		ascendingInsertionsortVolume(&daftarSaham, nData)
		tampilkanDaftarSaham()
	} else {
		urutDataAwal(&daftarSaham, nData)

	}

}

/*
func cariKodeSaham(dataSaham tabsa, searchKey string) int { // funsi cari di array saham

	var idx, i int
	idx = -1
	i = 0
	for idx == -1 && i < 10 {
		if dataSaham[i].id == searchKey {
			idx = i

		}
		i = i + 1
	}
	return idx

}
*/
func cariKodeSaham(a tabsa, n int, x string) int { //mencari kode saham dengan menggunakan binary search
	var left, right, idx, mid, idxasli int //syarat binary search adalah data harus terurut
	mengurutkanidsaham(&a, n)              //jadi biar dapat digunakan idnya diurutin dulu
	left = 0
	right = n - 1
	idx = -1
	for left <= right && idx == -1 {
		mid = (left + right) / 2
		if x < a[mid].id {
			right = mid - 1
		} else if x > a[mid].id {
			left = mid + 1
		} else {
			idx = mid
		}
	}
	if idx == -1 { //dilakkan pengecekan jika -1 
		urutDataAwal(&a, 10) 
		return -1
	}
	idxasli = a[idx].indeksAsli //karena sudah menyimpan data indeks yang asli / urutan awal jadi ketika di kembalikan pada urutan awal data dalam indeks tetap sama
	urutDataAwal(&a, 10)        //dikembalikan pada urutan awal
	return idxasli
}

func cariKodeSahamPengguna(dataSaham tabpe, searchKey string) int { // funsi cari di array pengguna
	var idx, i int
	idx = -1
	i = 0
	for idx == -1 && i < nData {
		if dataSaham[i].SahamDimiliki == searchKey {
			idx = i

		}
		i = i + 1
	}
	return idx

}

func beli() { //beli: nama saham, kode saham, beli berdasarkan lembar saham(bukan nilai dlm rupiah) dan pastikan saldo masih cukup,

	var pilih string
	var cariPe, cariSa, banyak int
	fmt.Printf("Saldo anda : Rp %.2f", saldo)
	fmt.Println("Masukkan kode saham yang ingin dibeli, 'buka' untuk melihat daftar: ")
	fmt.Scan(&pilih)
	switch pilih {
	case "buka": //pilihan ini akan membuka daftar saham
		tampilkanDaftarSaham()
	default: //pilihan ini akan melakukan transaksi pembelian
		fmt.Println("Anda ingin membeli berapa lembar saham: ")
		fmt.Scan(&banyak)                              // banyak lembar saham yang akan di beli
		cariSa = cariKodeSaham(daftarSaham, nData, pilih) //akan mencari kode saham di daftar saham
		if cariSa != -1 {                              //ketika kode saham ditemukan
			cariPe = cariKodeSahamPengguna(portofolio, pilih) // mencari kode saham di portofolio apakah pengguna sudah memiliki saham tersebut atau tidak

			if cariPe == -1 { // -1 berarti pengguna belum memiliki saham tersebut
				portofolio[sahamPunya].SahamDimiliki = pilih                     //menambahkan saham dimiliki pada portofolio
				if (float64(banyak) * daftarSaham[cariSa].hargaSaham) <= saldo { //jika pembelian masih di bawah atau sama dengan saldo
					portofolio[0].totalNilaiSahamDimiliki += daftarSaham[cariSa].hargaSaham * float64(banyak)    // total semua saham yang dimiliki
					portofolio[sahamPunya].nilaiSahamDimiliki = daftarSaham[cariSa].hargaSaham * float64(banyak) //hanya mencatat khusus
					saldo = saldo - daftarSaham[cariSa].hargaSaham*float64(banyak)                               //update sisa saldo
					portofolio[sahamPunya].lembarSahamDimiliki = banyak
					portofolio[sahamPunya].indeksAsli = indeksBenar
					fmt.Printf("Anda sudah membeli saham  %s \n", pilih)
					fmt.Printf("Saldo anda sekarang: Rp %.2f \n", saldo)
					sahamPunya += 1 // kepemilikan jenis saham bertambah 1
					indeksBenar += 1
				} else {
					fmt.Println("Saldo anda tidak mencukupi, pembelian gagal!")
				}
			} else { // pengguna sudah memiliki saham tersebut berarti hanya menambah nilai saham
				if float64(banyak)*daftarSaham[cariSa].hargaSaham <= saldo { //jika saldo masih cukup
					portofolio[0].totalNilaiSahamDimiliki += daftarSaham[cariSa].hargaSaham * float64(banyak) // total semua saham yang dimiliki
					portofolio[cariPe].nilaiSahamDimiliki += daftarSaham[cariSa].hargaSaham * float64(banyak) // hanya mencatat khusus
					portofolio[cariPe].lembarSahamDimiliki += banyak
					saldo = saldo - daftarSaham[cariSa].hargaSaham*float64(banyak)
					fmt.Printf("Anda sudah memperbarui jumlah saham %s \n", pilih)
					fmt.Printf("Saldo anda sekarang: Rp %.2f", saldo)
				} else {
					fmt.Println("Saldo anda tidak mencukupi, pembelian gagal!")
				}
			}
		} else { // jika kode saham yang diinputkan ternyata tidak ada
			fmt.Println("Kode saham yang Anda cari tidak ada!!")
		}

	}
}
func hapusSahamPengguna(index int) {
	var i int = index
	portofolio[index].totalNilaiSahamDimiliki = portofolio[index].totalNilaiSahamDimiliki - portofolio[index].nilaiSahamDimiliki

	for i < 10 {
		portofolio[i].nilaiSahamDimiliki = portofolio[i+1].nilaiSahamDimiliki
		portofolio[i].SahamDimiliki = portofolio[i+1].SahamDimiliki
		i++
	}
	portofolio[sahamPunya-1] = pengguna{} // Membersihkan slot terakhir
	sahamPunya--
}

// Untuk yang function jual masih dalam proses pengerjaan//
func jual() {
	var saham string
	var jumlah, cariPe, cariSa int
	var hargaPerLembar, totalJual float64

	fmt.Println("Ingin jual saham apa: ")
	fmt.Scan(&saham)
	fmt.Println("Berapa banyak jumlah saham yang ingin dijual: ")
	fmt.Scan(&jumlah)

	cariPe = cariKodeSahamPengguna(portofolio, saham)
	cariSa = cariKodeSaham(daftarSaham, nData, saham)
	if cariPe == -1 || cariSa == -1 {
		fmt.Println("Saham tidak ditemukan dalam portofolio Anda.")
		return
	}

	if jumlah > portofolio[cariPe].lembarSahamDimiliki {
		fmt.Println("Anda tidak memiliki cukup lembar saham untuk dijual.")
		return
	}

	hargaPerLembar = daftarSaham[cariSa].hargaSaham
	totalJual = float64(jumlah) * hargaPerLembar

	// update portofolio
	portofolio[cariPe].lembarSahamDimiliki -= jumlah
	portofolio[cariPe].nilaiSahamDimiliki -= totalJual
	portofolio[0].totalNilaiSahamDimiliki -= totalJual

	saldo += totalJual

	// hapus saham dari portofolio jika sudah 0 lembar
	if portofolio[cariPe].lembarSahamDimiliki == 0 {
		hapusSahamPengguna(cariPe)
		fmt.Println("Saham dijual habis dan dihapus dari portofolio.")
	} else {
		fmt.Printf("Saham %s telah berhasil dijual sebanyak %d lembar.\n", saham, jumlah)
	}
	fmt.Printf("Saldo anda sekarang: Rp %.2f\n", saldo)
}

//Untuk yang function jual masih dalam proses pengerjaan//

func portofolioPengguna() { //input nama pengguna lalu buat var saham dimiliki, nilai saham, keuntungan, kerugian
	fmt.Println("----------------------------------------------------------------------------------------------------------------------")
	fmt.Printf("| %-20s | %-10s | %-20s | %-25s | %-10s |\n", "NAMA", "KODE SAHAM", "NILAI SAHAM DIMILIKI", "TOTAL NILAI SAHAM DIMILIKI", "LEMBAR SAHAM DIMILIKI")
	fmt.Println("----------------------------------------------------------------------------------------------------------------------")
	for i := 0; i < sahamPunya; i++ {
		if portofolio[i].SahamDimiliki != "" {
			fmt.Printf("| %-20s | %-10s | Rp %-18.2f | Rp %-23.2f | %-10d |\n", portofolio[i].nama, portofolio[i].SahamDimiliki, portofolio[i].nilaiSahamDimiliki, portofolio[i].totalNilaiSahamDimiliki, portofolio[i].lembarSahamDimiliki)
		}
	}
	fmt.Println("----------------------------------------------------------------------------------------------------------------------")

}
func tampilkanStatistikUntungRugi() {
	fmt.Println("\n-------------------------STATISTIK KEUNTUNGAN / KERUGIAAN ----------------------------------------------------------")
	fmt.Printf("Total Keuntungan      : Rp %.2f\n", statstk[0].totalKeuntungan)
	fmt.Printf("Total Kerugian        : Rp %.2f\n", statstk[0].totalKerugian)
	fmt.Printf("Jumlah Harga Naik     : %.0f kali \n", statstk[0].naikHarga)
	fmt.Printf("Jumlah Harga Turun    : %.0f kali \n", statstk[0].turunHarga)
	fmt.Println("----------------------------------------------------------------------------------------------------------------------")

}

func simulasiTrading() { //n untuk berapa lama menahan
	/* idenya adalah asumsikan bahwa naik turun harga saham memiliki kemungkinan untung 3 : 1 rugi
	anggap kita mempunya angka random 1-5 dengan jika dapat 1 tidak ada perubahan jika dapat 2 Harga turun sekian dalam %
	jika dapat 3-5 harga naik sekian dalam %. anggap tiap kenaikkan adalah dalam rentang 1 bulan sehingga masukkan
	misal 4 jadi kita seperti perkembangan harga saham dalam 4 bulan atau 4 kali naik / turun */
	var n, random int
	var hargaSebelum, nilaiSebelum float64
	var selisihHarga, selisihNilai float64
	var adanaik, adaturun bool
	adanaik = false
	adaturun = false
	fmt.Println("Simulasi pergerakan harga saham sebanyak 1 perbulan, ingin menahan berapa lama: ")
	fmt.Scan(&n)
	for i := 1; i <= n; i++ {
		rand.Seed(int64(i + 4*n))
		random = rand.Intn(6) + 1
		for j := 0; j < nData; j++ { //berhenti ketika seluruh daftar saham terupdate
			if j <= sahamPunya { //kondisi update portofolio agar tidak melebihi saham yang dimiliki
				hargaSebelum = portofolio[j].nilaiSahamDimiliki
				nilaiSebelum = portofolio[j].totalNilaiSahamDimiliki
			}
			switch random {
			case 1:
				// tidak ada perubahan
			case 2:
				daftarSaham[j].hargaSaham *= 0.98 // turun 2%
				daftarSaham[j].nilaiSaham *= 0.98
				if j <= sahamPunya {
					portofolio[j].nilaiSahamDimiliki *= 0.98
					portofolio[0].totalNilaiSahamDimiliki *= 0.98
				}
			case 3:
				daftarSaham[j].hargaSaham *= 1.03 // naik 3%
				daftarSaham[j].nilaiSaham *= 1.03
				if j <= sahamPunya {
					portofolio[j].nilaiSahamDimiliki *= 1.03
					portofolio[0].totalNilaiSahamDimiliki *= 1.03
				}
			case 4:
				daftarSaham[j].hargaSaham *= 1.04 // naik 4%
				daftarSaham[j].nilaiSaham *= 1.04
				if j <= sahamPunya {
					portofolio[j].nilaiSahamDimiliki *= 1.04
					portofolio[0].totalNilaiSahamDimiliki *= 1.04
				}
			case 5:
				daftarSaham[j].hargaSaham *= 1.05 // naik 5%
				daftarSaham[j].nilaiSaham *= 1.05
				if j <= sahamPunya {
					portofolio[j].nilaiSahamDimiliki *= 1.05
					portofolio[0].totalNilaiSahamDimiliki *= 1.05
				}
			case 6:
				daftarSaham[j].hargaSaham *= 0.97 // turun 3%
				daftarSaham[j].nilaiSaham *= 0.97
				if j <= sahamPunya {
					portofolio[j].nilaiSahamDimiliki *= 0.97
					portofolio[0].totalNilaiSahamDimiliki *= 0.97
				}
			}
			if j <= sahamPunya {
				selisihHarga = portofolio[j].nilaiSahamDimiliki - hargaSebelum
				selisihNilai = portofolio[0].totalNilaiSahamDimiliki - nilaiSebelum

				if selisihHarga > 0 { //update nanti menyelesaikan permasalahan naik/turun
					statstk[0].totalKeuntungan += selisihNilai
					adanaik = true
				} else if selisihHarga < 0 {
					statstk[0].totalKerugian += -selisihNilai
					adaturun = true
				}
			}
		}
		if adanaik {
			statstk[0].naikHarga++
		} else if adaturun {
			statstk[0].turunHarga++
		}

	}
	fmt.Println("Harga telah diperbarui.")
}
func mengurutkanidsaham(a *tabsa, n int) {
	var i, pass int
	var temp saham
	pass = 1
	for pass <= n-1 {
		i = pass
		temp = (*a)[pass]
		for i > 0 && temp.id < (*a)[i-1].id {
			(*a)[i] = (*a)[i-1]
			i = i - 1
		}
		(*a)[i] = temp
		pass = pass + 1
	}
}

// function urutDataDescedingHarga masih dalam tahap pengerjaan//
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

func urutDataAwal(data *tabsa, ndata int) { //untuk mengurutkan data saham kembali ke setingan awal
	var i, idx, pass int
	var temp saham //penyimpanan data sementara
	pass = 1
	for pass < ndata {
		idx = pass - 1
		i = pass
		for i < ndata {
			if (*data)[i].indeksAsli < (*data)[idx].indeksAsli {
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

// function urutDataDescedingVolume masih dalam tahap pengerjaan//
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
func ascendingInsertionsortVolume(a *tabsa, n int) {
	var i, pass int
	var temp saham
	pass = 1
	for pass <= n-1 {
		i = pass
		temp = (*a)[pass]
		for i > 0 && temp.volume < (*a)[i-1].volume {
			(*a)[i] = (*a)[i-1]
			i = i - 1
		}
		(*a)[i] = temp
		pass = pass + 1
	}
}
func ascendingInsertionsortHargaSa(a *tabsa, n int) {
	var i, pass int
	var temp saham
	pass = 1
	for pass <= n-1 {
		i = pass
		temp = (*a)[pass]
		for i > 0 && temp.hargaSaham < (*a)[i-1].hargaSaham {
			(*a)[i] = (*a)[i-1]
			i = i - 1
		}
		(*a)[i] = temp
		pass = pass + 1
	}
}

func pilihan() { //pengguna memilih akan melakukan apa
	for {
		var pilih int
		fmt.Println("===================MENU UTAMA===================")
		fmt.Printf("  Pilih hal apa yang akan anda lakukan: \n  1. Tampilkan daftar saham. \n  2. Tampilkan portofolio anda. \n  3. Beli saham. \n  4. Jual saham. \n  5. Tampilkan statistik keuntungan/kerugian. \n  6. Melakukan simulasi saham. \n  7. Keluar.\n")
		fmt.Println("================================================")
		fmt.Scan(&pilih)
		switch pilih {
		case 1: //memanggil fungsi daftar saham
			tampilkanDaftarSaham()
		case 2: //fungsi porto
			portofolioPengguna()
		case 3:
			beli()
		case 4:
			jual()
		case 5:
			tampilkanStatistikUntungRugi()
		case 6:
			simulasiTrading()
		case 7:
			return
		default:
			fmt.Println("Pilihan tidak valid!")
		}
		fmt.Println()
	}
}

func main() {
	fmt.Print("Masukkan nama Anda: ")
	fmt.Scan((&portofolio[0].nama))
	inisialisasiDataSaham()
	pilihan()
	fmt.Println("Terima kasih telah menggunakan applikasi simulasi saham ini !!!")

}
