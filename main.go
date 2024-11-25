package main

import (
	"day03/pustaka"
	"fmt"
)
func makan(){
	fmt.Println("Saya Makan Nasi")
}


func main(){
	// makan()
	// pustaka.BilangHalo()
	//pustaka.JalanAntrian()
	//pustaka.JalanAntrianWG()
	//pustaka.JalanAntrianChannel()
	pustaka.JalanAntrianGabungan()
	dMhs := struct {
		nama string
		umur int
	}{
		"budi",
		25,
	}
	dMhs.nama = "Anto"
	dMhs.umur = 50

	//fmt.Println(dMhs)
	type dKaryawan struct {
		nama string
		umur int
		alamat string
	}

	var arrData = []dKaryawan{
		{
			nama: "Budi",
			umur: 21,
			alamat: "Jakarta",
		},{
			nama: "Anto",
			umur: 50,
			alamat: "Surabaya",
		},{
			nama: "Dina",
			umur: 21,
			alamat: "Medan",
		},
	}
	fmt.Println(arrData)
	// for index, isi := range arrData {
	// 	fmt.Println(index,isi.nama, isi.umur,isi.alamat)
	// }
}