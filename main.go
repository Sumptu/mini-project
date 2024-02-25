package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Buku struct {
	KodeBuku string
	JudulBuku string
	Pengarang string
	Penerbit string
	JumlahHalaman int
	TahunTerbit int
}

var ListBuku []Buku
	// di os windows pada strings/replace terdapat perbedaan sytax pada kode "\r\n" yang awalnya "\n

func TambahBuku() {
	inputanUser := bufio.NewReader(os.Stdin)
	jumlahHalaman := 0
	tahunTerbit := 0

	fmt.Println("===========================================")
	fmt.Println("Tambah Buku Baru")
	fmt.Println("===========================================")
	fmt.Print("Masukkan Kode Buku : ")

	kodeBuku, err := inputanUser.ReadString('\n')
	if err != nil {
		fmt.Println("Terjadi Error", err)
		return
	}

	kodeBuku = strings.Replace(
		kodeBuku,
		"\r\n",
		"",
		1,
	)

	fmt.Print("Masukkan Judul Buku : ")
	judulBuku, err := inputanUser.ReadString('\n')
	if err != nil {
		fmt.Println("Terjadi Error", err)
		return
	}

	judulBuku = strings.Replace(
		judulBuku,
		"\r\n",
		"",
		1,
	)

	fmt.Print("Masukkan Pengarang : ")
	pengarang, err := inputanUser.ReadString('\n')
	if err != nil {
		fmt.Println("Terjadi Error", err)
		return
	}

	pengarang = strings.Replace(
		pengarang,
		"\r\n",
		"",
		1,
	)

	fmt.Print("Masukkan Penerbit : ")
	penerbit, err := inputanUser.ReadString('\n')
	if err != nil {
		fmt.Println("Terjadi Error", err)
		return
	}

	penerbit = strings.Replace(
		penerbit,
		"\r\n",
		"",
		1,
	)

	fmt.Print("Masukkan Jumlah Halaman : ")
	_, err = fmt.Scanln(&jumlahHalaman)
	if err != nil {
		fmt.Println("Terjadi Error:", err)
		return
	}

	fmt.Print("Masukkan Tahun terbit : ")
	_, err = fmt.Scanln(&tahunTerbit)
	if err != nil {
		fmt.Println("Terjadi Error:", err)
		return
	}

	ListBuku = append(ListBuku, Buku{
		KodeBuku: kodeBuku,
		JudulBuku: judulBuku,
		Pengarang: pengarang,
		Penerbit: penerbit,
		JumlahHalaman: jumlahHalaman,
		TahunTerbit: tahunTerbit,
	})

	fmt.Println("Berhasil Menambah Buku!")
}

func LiatBuku() {
	fmt.Println("===========================================")
	fmt.Println("Lihat Buku")
	fmt.Println("===========================================")

	for urutan, buku := range ListBuku {
		fmt.Printf("%d, Kode: %s, Judul : %s, Pengarang : %s, Penerbit : %s, Halaman : %d, Terbit : %d\n",
			urutan+1,
			buku.KodeBuku,
			buku.JudulBuku,
			buku.Pengarang,
			buku.Penerbit,
			buku.JumlahHalaman,
			buku.TahunTerbit,
		)
	}
}

func EditBuku() {
	fmt.Println("===========================================")
	fmt.Println("Edit Buku")
	fmt.Println("===========================================")
	
	LiatBuku()
	
	fmt.Println("===========================================")
	var kodeBuku string
	fmt.Print("Masukkan Kode Buku yang akan diedit : ")

	inputanUser := bufio.NewReader(os.Stdin)
	kodeBuku, err := inputanUser.ReadString('\n')
	if err != nil {
		fmt.Println("Terjadi error:", err)
		return
	}

	kodeBuku = strings.Replace(
		kodeBuku,
		"\r\n",
		"",
		1,
	)
	
	bukuDitemukan := false
	for i, buku := range ListBuku {
		if buku.KodeBuku == kodeBuku {
			bukuDitemukan = true
	
				fmt.Printf("Buku yang akan diubah:\nKode Buku: %s, Judul: %s, Pengarang: %s, Penerbit: %s, Halaman: %d, Terbit: %d\n",
				buku.KodeBuku, buku.JudulBuku, buku.Pengarang, buku.Penerbit, buku.JumlahHalaman, buku.TahunTerbit)
	
				inputanUser := bufio.NewReader(os.Stdin)
				judulBuku := ""
				pengarang := ""
				penerbit := ""
				jumlahHalaman := 0
				tahunTerbit := 0

				fmt.Println("===========================================")
				fmt.Print("Silahkan Masukkan Judul Buku baru: ")
				judulBuku, err = inputanUser.ReadString('\n')
				if err != nil {
					fmt.Println("Terjadi Error:", err)
					return
				}
				judulBuku = strings.Replace(
					judulBuku,
					"\r\n",
					"",
					1,
				)
	
				fmt.Print("Silahkan Masukkan Pengarang baru: ")
				pengarang, err = inputanUser.ReadString('\n')
				if err != nil {
					fmt.Println("Terjadi Error:", err)
					return
				}
				pengarang = strings.Replace(
					pengarang, 
					"\r\n", 
					"",
					1,
				)
	
				fmt.Print("Silahkan Masukkan Penerbit baru: ")
				penerbit, err = inputanUser.ReadString('\n')
				if err != nil {
					fmt.Println("Terjadi Error:", err)
					return
				}
				penerbit = strings.Replace(
					penerbit, 
					"\r\n", 
					"", 
					1,
				)
	
				fmt.Print("Silahkan Masukkan Jumlah Halaman baru: ")
				_, err = fmt.Scanln(&jumlahHalaman)
				if err != nil {
					fmt.Println("Terjadi Error:", err)
					return
				}
	
				fmt.Print("Silahkan Masukkan Tahun Terbit baru: ")
				_, err = fmt.Scanln(&tahunTerbit)
				if err != nil {
					fmt.Println("Terjadi Error:", err)
					return
				}
	
				ListBuku[i] = Buku{
					KodeBuku:     kodeBuku,
					JudulBuku:    judulBuku,
					Pengarang:    pengarang,
					Penerbit:     penerbit,
					JumlahHalaman: jumlahHalaman,
					TahunTerbit:  tahunTerbit,
				}
	
				fmt.Println("Buku Berhasil Diubah!")
				break
			}
		}
	
		if !bukuDitemukan {
			fmt.Println("Buku dengan Kode Buku tersebut tidak ditemukan.")
		}
}
	


func HapusBuku() {
	fmt.Println("=================================")
	fmt.Println("Hapus Pesanan")
	fmt.Println("=================================")
	LiatBuku()
	fmt.Println("=================================")

	var urutanBuku int
	fmt.Print("Masukkan Urutan Buku : ")
	_, err := fmt.Scanln(&urutanBuku)
	if err != nil {
		fmt.Println("Terjadi error:", err)
	}

	if (urutanBuku-1) < 0 || 
		(urutanBuku-1) > len (ListBuku) {
			fmt.Println("Urutan Buku Tidak Sesuai")
			HapusBuku()
			return
		}
	
		ListBuku = append(
			ListBuku[:urutanBuku-1],
			ListBuku[urutanBuku:]...,
		)
	fmt.Println("Buku berhasil dihapus!")
}

func main() {
	var pilihan int
	fmt.Println("===========================================")
	fmt.Println("Aplikasi Manajemen Daftar Buku Perpustakaan")
	fmt.Println("===========================================")
	fmt.Println("Silahkan Pilih : ")
	fmt.Println("1. Tambah Buku Baru")
	fmt.Println("2. Liat Daftar Buku")
	fmt.Println("3. Edit Buku")
	fmt.Println("4. Hapus Buku")
	fmt.Println("5. Keluar")
	fmt.Println("===========================================")
	fmt.Print("Masukan Pilihan : ")
	_, err := fmt.Scanln(&pilihan)
	if err != nil {
		fmt.Println("Terjadi error:", err)
	}	

	switch (pilihan) {
	case 1:
		TambahBuku()
	case 2: 
		LiatBuku()
	case 3:
		EditBuku()
	case 4:
		HapusBuku()
	case 5:
		os.Exit(0)
	}
	main()
}

