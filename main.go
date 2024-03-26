package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"encoding/json"
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

// disini menggunakan map untuk menyimpan kode buku yang sudah digunakan atau belum
var IsKodeBukuUsed map[string]bool

func init() {
	IsKodeBukuUsed = make(map[string]bool)
}

func TambahBuku() {
	inputanUser := bufio.NewReader(os.Stdin)
	jumlahHalaman := 0
	tahunTerbit := 0

	for {
		fmt.Println("===========================================")
		fmt.Println("Tambah Buku Baru")
		fmt.Println("===========================================")

		var kodeBuku string
		for {
			fmt.Print("Masukkan Kode Buku : ")
			kodeBuku, _ = inputanUser.ReadString('\n')
			kodeBuku = strings.TrimSpace(kodeBuku)

			if _, found := IsKodeBukuUsed[kodeBuku]; found {
				fmt.Println("Kode buku sudah digunakan. Silakan masukkan kode buku yang berbeda.")
			} else {
				break
			}
		}

		IsKodeBukuUsed[kodeBuku] = true
		fileName := fmt.Sprintf("books/book-%s.json", kodeBuku)
		if _, err := os.Stat(fileName); err == nil {
			fmt.Println("File JSON sudah ada untuk kode buku ini.")
			continue
		}

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

		err = os.MkdirAll("books", 0755)
		if err != nil {
			fmt.Println("Gagal membuat direktori:", err)
			return
		}

		file, err := os.Create(fileName)
		if err != nil {
			fmt.Println("Gagal membuat file:", err)
			return
		}
		defer file.Close()

		encoder := json.NewEncoder(file)
		if err := encoder.Encode(ListBuku); err != nil {
			fmt.Println("Gagal menyimpan buku ke dalam file:", err)
		}

		fmt.Println("Berhasil Menambah Buku!")

		fmt.Print("Tambahkan buku lagi? (y/n): ")
		var jawaban string
		_, err = fmt.Scanln(&jawaban)
		if err != nil {
			fmt.Println("Terjadi error:", err)
			return
		}

		if strings.ToLower(strings.TrimSpace(jawaban)) != "y" {
			break 
		}
	}
}

func LiatBuku() {
	fmt.Println("===========================================")
	fmt.Println("Lihat Buku")
	fmt.Println("===========================================")

	files, err := os.ReadDir("books")
	if err != nil {
		fmt.Println("Gagal membaca direktori:", err)
		return
	}

	ListBuku = nil
	foundFiles := false

	for _, file := range files {
		if !file.IsDir() && strings.HasSuffix(file.Name(), ".json") {
			foundFiles = true
			fileName := fmt.Sprintf("books/%s", file.Name())

			jsonFile, err := os.Open(fileName)
			if err != nil {
				fmt.Println("Gagal membuka file:", err)
				continue
			}
			defer jsonFile.Close()

			var buku []Buku
			decoder := json.NewDecoder(jsonFile)
			if err := decoder.Decode(&buku); err != nil {
				fmt.Println("Gagal membaca isi file JSON:", err)
				continue
			}

			for _, bukuBaru := range buku {
				bukuDitemukan := false
				for _, bukuSudahAda := range ListBuku {
					if bukuBaru.KodeBuku == bukuSudahAda.KodeBuku {
						bukuDitemukan = true
						break
					}
				}
				if !bukuDitemukan {
					ListBuku = append(ListBuku, bukuBaru)
				}
			}
		}
	}

	if !foundFiles {
		fmt.Println("Tidak ada file buku di direktori 'books'.")
		return
	}

	for urutan, buku := range ListBuku {
		fmt.Printf("%d. Kode: %s, Judul : %s, Pengarang : %s, Penerbit : %s, Halaman : %d, Terbit : %d\n",
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
                KodeBuku:      kodeBuku,
                JudulBuku:     judulBuku,
                Pengarang:     pengarang,
                Penerbit:      penerbit,
                JumlahHalaman: jumlahHalaman,
                TahunTerbit:   tahunTerbit,
            }

            fileName := fmt.Sprintf("books/book-%s.json", kodeBuku)
            file, err := os.Create(fileName)
            if err != nil {
                fmt.Println("Gagal membuat file:", err)
                return
            }
            defer file.Close()

            encoder := json.NewEncoder(file)
            if err := encoder.Encode(ListBuku); err != nil {
                fmt.Println("Gagal menyimpan buku ke dalam file:", err)
                return
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

	if (urutanBuku-1) < 0 || (urutanBuku-1) >= len(ListBuku) {
		fmt.Println("Urutan Buku Tidak Sesuai")
		return
	}

	buku := ListBuku[urutanBuku-1]

	ListBuku = append(ListBuku[:urutanBuku-1], ListBuku[urutanBuku:]...)

	fileName := fmt.Sprintf("books/book-%s.json", buku.KodeBuku)
	err = os.Remove(fileName)
	if err != nil {
		fmt.Println("Gagal menghapus file JSON:", err)
		return
	}

	fmt.Println("Buku berhasil dihapus!")
}
	
func PrintBuku() {
	fmt.Println("=================================")
	fmt.Println("Print Buku")
	fmt.Println("=================================")
	LiatBuku()
	fmt.Println("=================================")

	err := os.MkdirAll("pdf", os.ModePerm)
    if err != nil {
        fmt.Println("Gagal membuat direktori 'pdf':", err)
        return
    }

	var pilihan int
	fmt.Println("Pilih opsi yang anda inginkan")
	fmt.Println("1. Print Satu Buku")
	fmt.Println("2. Print Semua Buku")
	fmt.Print("Pilih opsi: ")
	_, err = fmt.Scanln(&pilihan)
	if err != nil {
		fmt.Println("Terjadi error:", err)
		return
	}

	switch pilihan {
	case 1:
		fmt.Print("Masukkan Urutan Buku yang akan di-print: ")
		var urutanBuku int
		_, err := fmt.Scanln(&urutanBuku)
		if err != nil {
			fmt.Println("Terjadi error:", err)
			return
		}
		if urutanBuku < 1 || urutanBuku > len(ListBuku) {
			fmt.Println("Urutan Buku tidak valid.")
			return
		}

		err = generateSatuBukuPDF(ListBuku[urutanBuku-1])
		if err != nil {
			fmt.Println("Gagal membuat PDF:", err)
			return
		}
		fmt.Println("Buku berhasil di-print sebagai PDF.")
	case 2:
		err := generateSemuaBukuPDF()
		if err != nil {
			fmt.Println("Gagal membuat PDF:", err)
			return
		}
		fmt.Println("Semua buku berhasil di-print sebagai PDF.")
	default:
		fmt.Println("Pilihan tidak valid.")
	}
}

func generateSatuBukuPDF(buku Buku) error {
	fileName := fmt.Sprintf("pdf/%s.pdf", buku.KodeBuku)
	pdf := fpdf.New("P", "mm", "A4", "")
	pdf.AddPage()
	pdf.SetFont("Arial", "B", 16)
	pdf.Cell(40, 10, "Kode Buku: "+buku.KodeBuku)
	pdf.Ln(-1)
	pdf.Cell(40, 10, "Judul Buku: "+buku.JudulBuku)
	pdf.Ln(-1)
	pdf.Cell(40, 10, "Pengarang: "+buku.Pengarang)
	pdf.Ln(-1)
	pdf.Cell(40, 10, "Penerbit: "+buku.Penerbit)
	pdf.Ln(-1)
	pdf.Cell(40, 10, fmt.Sprintf("Jumlah Halaman: %d", buku.JumlahHalaman))
	pdf.Ln(-1)
	pdf.Cell(40, 10, fmt.Sprintf("Tahun Terbit: %d", buku.TahunTerbit))
	return pdf.OutputFileAndClose(fileName)
}

func generateSemuaBukuPDF() error {
	fileName := "pdf/semua_buku.pdf"
	pdf := fpdf.New("P", "mm", "A4", "")
	pdf.AddPage()
	pdf.SetFont("Arial", "B", 16)
	for _, buku := range ListBuku {
		pdf.Cell(40, 10, "Kode Buku: "+buku.KodeBuku)
		pdf.Ln(-1)
		pdf.Cell(40, 10, "Judul Buku: "+buku.JudulBuku)
		pdf.Ln(-1)
		pdf.Cell(40, 10, "Pengarang: "+buku.Pengarang)
		pdf.Ln(-1)
		pdf.Cell(40, 10, "Penerbit: "+buku.Penerbit)
		pdf.Ln(-1)
		pdf.Cell(40, 10, fmt.Sprintf("Jumlah Halaman: %d", buku.JumlahHalaman))
		pdf.Ln(-1)
		pdf.Cell(40, 10, fmt.Sprintf("Tahun Terbit: %d", buku.TahunTerbit))
		pdf.Ln(-1)
		pdf.Ln(-1)
	}
	return pdf.OutputFileAndClose(fileName)
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
	fmt.Println("5. Print Buku")
	fmt.Println("6. Keluar")
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
		PrintBuku()
	case 6:
		os.Exit(0)
	}
	main()
}

