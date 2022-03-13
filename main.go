package main

import (
	"fmt"

	modul "github.com/damian3197/tugasgomodules"
)

func main(){
	I := 5
	V := 120
	modul.Welcome()
	modul.Verifikasi("Damian", "Sitanggang")
	modul.AnonymousDataBuku()
	var power, resistor = modul.PowerCal(I, V)
	fmt.Println("Arus =", I, " Tegangan = ", V, " Maka Daya = ", power, " Tahanan = ", resistor)
}