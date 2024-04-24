package main

import (
	"fmt"
	"math/rand"
)

func main() {
	ngc2c1()
	ngc2c2()
}

func ngc2c1() {

	var name string

	fmt.Println("Masukkan Nama: ")
	fmt.Scanln(&name)

	// rand.Seed(time.Now().UnixNamo())
	number := rand.Intn(100) + 1

	switch {
	case number > 80:
		fmt.Printf("Random angka: %v ", number)
		fmt.Printf("Selamat %s, anda sangat beruntung\n", name)
	case number <= 80 && number > 60:
		fmt.Printf("Random angka: %v ", number)
		fmt.Printf("Selamat %s, anda beruntung\n", name)
	case number <= 60 && number > 40:
		fmt.Printf("random angka: %v ", number)
		fmt.Printf("Mohon maaf %s, anda kurang beruntung\n", name)
	case number <= 40:
		fmt.Printf("Random angka: %v ", number)
		fmt.Printf("Mohon maaf %s, anda sial\n", name)
	}

}

func ngc2c2() {
	var nama string
	fmt.Println("Masukkan nama Anda: ")
	fmt.Scan(&nama)

	// reader := buFio.NewReader(os.Stdin)
	// fmt.Print("Masukkan Nama")
	// nama, _ := reader.ReadString('\n')

	// nama = strings.TrimSpace(nama)

	var umur int
	fmt.Println("Masukkan umur Anda: ")
	fmt.Scan(&umur)

	if umur < 0 {
		fmt.Println("Error: Umur tidak boleh kurang dari 0")
	} else if umur > 100 {
		fmt.Println("Error: Umur tidak boleh lebih dari 100")
	} else {
		if umur > 18 {
			fmt.Println("Silahkan masuk")
		} else {
			fmt.Println("Dilarang masuk, maksimal umur 19")
		}
	}
}
