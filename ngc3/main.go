package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

func main() {
	menu()
}

func looping1() {
	//Looping 1
	//Membuat slice yang berisi data map orang-orang
	people := []map[string]interface{}{
		{"name": "Hank", "age": 50, "job": "Polisi"},
		{"name": "Heisenberg", "age": 52, "job": "Ilmuwan"},
		{"name": "Skler", "age": 48, "job": "Akuntan"},
	}

	//looping untuk menampilkan data setiap orang
	for _, person := range people {
		name := person["name"].(string)
		age := person["age"].(int)
		job := person["job"].(string)

		fmt.Printf("Hi Perkenalkan, Nama saya %s, umur saya %d, dan saya bekerja sebagai %s\n", name, age, job)
	}
	//End looping 1
}

func looping2() {
	//Looping 2
	// Slice1
	slice1 := []float64{1, 5, 7, 8, 10, 24, 33}
	// Slice2
	slice2 := []float64{1.1, 5.4, 6.7, 9.2, 11.3, 25.2, 33.1}

	//Menghitung Rata-rata, jumlah, dan median dari slice1
	avg1, sum1, median1 := calculate(slice1)
	fmt.Println("Slice1: ")
	fmt.Printf("Rata-rata: %.2f\n", avg1)
	fmt.Printf("jumlah: %.2f\n", sum1)
	fmt.Printf("Median: %.2f\n", median1)

	//Menghitung Rata-rata, jumlah, dan median dari slice2
	avg2, sum2, median2 := calculate(slice2)
	fmt.Println("Slice2: ")
	fmt.Printf("Rata-rata: %.2f\n", avg2)
	fmt.Printf("Jumlah: %.2f\n", sum2)
	fmt.Printf("Median: %.2f\n", median2)
}

func calculate(slice []float64) (float64, float64, float64) {
	//Menghitung jumlah
	sum := 0.0
	for _, value := range slice {
		sum += value
	}

	//Menghitung rata-rata
	avg := sum / float64(len(slice))

	//Menghitung median
	sort.Float64s(slice)
	median := 0.0
	if len(slice)%2 == 0 {
		median = (slice[len(slice)/2-1] + slice[len(slice)/2]) / 2.0
	} else {
		median = slice[len(slice)/2]
	}

	return avg, sum, median
	//End Looping2
}

// Looping 3
func looping3() {
	// Variabel kata yang akan diperiksa
	kata := "katak"

	// Memeriksa apakah kata merupakan palindrome
	isPalindrome := checkPalindrome(kata)

	// Menampilkan hasil
	fmt.Println(isPalindrome)
}

func checkPalindrome(kata string) bool {
	// Mengubah kata menjadi lowercase untuk memastikan kecocokan case
	kata = strings.ToLower(kata)

	// Membalik kata
	reversed := ""
	for i := len(kata) - 1; i >= 0; i-- {
		reversed += string(kata[i])
	}

	// Membandingkan kata asli dengan kata yang dibalik
	return kata == reversed
	//End looping 3
}

// Looping 4
func looping4() {
	// Variabel kata yang akan diperiksa
	kata := "xoxo"

	// Menghitung jumlah 'x' dan 'o'
	jumlahX, jumlahO := countXO(kata)

	// Menentukan apakah jumlah 'x' dan 'o' sama
	isSama := jumlahX == jumlahO

	// Menampilkan hasil
	fmt.Println(isSama)
}

func countXO(kata string) (int, int) {
	jumlahX := 0
	jumlahO := 0

	// Menghitung jumlah 'x' dan 'o'
	for _, char := range kata {
		if char == 'x' || char == 'X' {
			jumlahX++
		} else if char == 'o' || char == 'O' {
			jumlahO++
		}
	}

	return jumlahX, jumlahO
	//End Looping 4
}

// Looping 5 bubble sort
func looping5() {
	// Variabel kata yang akan diurutkan
	kata := "56432891"

	// Mengubah string menjadi slice dari int
	slice := make([]int, len(kata))
	for i, char := range kata {
		slice[i] = int(char - '0')
	}

	// Bubble Sort
	for i := 0; i < len(slice)-1; i++ {
		for j := 0; j < len(slice)-i-1; j++ {
			if slice[j] < slice[j+1] {
				// Menukar posisi jika slice[j] lebih kecil dari slice[j+1]
				slice[j], slice[j+1] = slice[j+1], slice[j]
			}
		}
	}

	// Menampilkan hasil
	fmt.Println(slice)
	//End Looping 5
}

// looping 6
func looping6() {
	//jumlah baris atau (row)
	rows1 := 5

	//looping untuk menampilkan barisan *
	for i := 0; i < rows1; i++ {
		fmt.Println("*")
	}
	//end looping 6
}

// looping 7
func looping7() {
	//jumlah baris atau (row)
	rows2 := 5

	// Nested looping untuk menampilkan *
	for i := 0; i < rows2; i++ {
		for j := 0; j < rows2; j++ {
			fmt.Print("*")
		}
		fmt.Println()
	}
	//End Looping 7
}

// looping 8
func looping8() {
	// Jumlah baris (rows)
	rows3 := 5

	// Nested looping untuk menampilkan barisan tangga bintang
	for i := 0; i < rows3; i++ {
		for j := 0; j <= i; j++ {
			fmt.Print("*")
		}
		fmt.Println()
	}
	//end looping 8
}

func looping9() {
	// Jumlah baris (rows)
	rows4 := 5

	// Nested looping untuk menampilkan barisan tangga bintang tterbalik
	for i := rows4; i > 0; i-- {
		for j := 0; j < i; j++ {
			fmt.Print("*")
		}
		fmt.Println()
	}
	//End looping 9
}

func menu() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("List")
	fmt.Println("1. Looping 1")
	fmt.Println("2. Looping 2")
	fmt.Println("3. Looping 2")
	fmt.Println("4. Looping 2")
	fmt.Println("5. Looping 2")
	fmt.Println("6. Looping 2")
	fmt.Println("7. Looping 2")
	fmt.Println("8. Looping 2")
	fmt.Println("9. Looping 2")
	fmt.Println("Masukkan Nomor Soal Yang Anda Inginkan: ")

	scanner.Scan()
	i := scanner.Text()
	switch i {
	case "1":
		looping1()
	case "2":
		looping2()
	case "3":
		looping3()
	case "4":
		looping4()
	case "5":
		looping5()
	case "6":
		looping6()
	case "7":
		looping7()
	case "8":
		looping8()
	case "9":
		looping9()
	}
}
