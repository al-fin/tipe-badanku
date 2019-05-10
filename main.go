package main

import (
	"fmt"
	"bufio"
	"os"
	"strings"
	"strconv"
)


func main() {
	var banner string = `
 _____ ___ ____  _____   ____    _    ____    _    _   _ _  ___   _ 
|_   _|_ _|  _ \| ____| | __ )  / \  |  _ \  / \  | \ | | |/ / | | |
  | |  | || |_) |  _|   |  _ \ / _ \ | | | |/ _ \ |  \| | ' /| | | |
  | |  | ||  __/| |___  | |_) / ___ \| |_| / ___ \| |\  | . \| |_| |
  |_| |___|_|   |_____| |____/_/   \_\____/_/   \_\_| \_|_|\_\\___/ 
   
  						Created by : Alfin
  `
	fmt.Println(banner)
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("[?] Tinggi badan kamu (cm) : ")
	tinggi,_ := reader.ReadString('\n')
	tinggi = strings.TrimSuffix(tinggi,"\n")

	fmt.Print("[?] Berat badan kamu (kg) : ")
	berat,_ := reader.ReadString('\n')
	berat = strings.TrimSuffix(berat,"\n")

	t,_ := strconv.ParseFloat(tinggi, 64)
	b,_ := strconv.ParseFloat(berat, 64)
	hasil,ideal := hitung(t,b)

	fmt.Println("====================================\n")
	fmt.Printf("[1] Tipe badan kamu adalah %s.\n", hasil)
	if hasil != "normal" {
		fmt.Printf("[i] Berat badan ideal kamu adalah %s kg.\n\n", ideal)
	}
}

func hitung(t float64, b float64) (hasil string,ideal string){

	bmi := b/((t/100)*(t/100))
	i := 20*((t/100)*(t/100))
	ideal = fmt.Sprintf("%.2f", i)

	fmt.Println()
	switch {
		case bmi<=16.0:
			hasil = "kurus"
		case bmi<=18.4:
			hasil = "agak kurus"
		case bmi<=25.0:
			hasil = "normal"
		case bmi<=27.0:
			hasil = "agak gemuk"
		case bmi>27.0:
			hasil = "gemuk"
	}
	return hasil, ideal
}
