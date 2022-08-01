package main

import "fmt"

/*CANİK, GLOCK VE  SİG MARKALARININ ŞARJÖRÜN İÇİNE
ALDIKLARI MÜHİMMAT KAPASİTELERİNİN TOPLAMININ İNCELENDİĞİ BÖLÜM..*/

func pistols(canik int, glock int, sig int) int {
	magCapacity := canik + glock + sig
	return magCapacity
	/*MARKALARIN ŞARJÖR KAPASİTELERİNİN AVRUPA STANDARTLARINA
	  UYGUN OLUP OLMADIĞININ İNCELENDİĞİ BÖLÜM..*/
}
func euStandCap(canik int, glock int, sig int) int {

	if canik > 15 {
		fmt.Println("Canik :Avrupa standartlarına uygun")
	} else {
		fmt.Println("Canik : sınıf dışı")
	}
	if glock > 15 {
		fmt.Println("Glock :Avrupa standartlarına uygun")
	} else {
		fmt.Println("Glock : sınıf dışı", glock)
	}
	if sig > 15 {
		fmt.Println("Sig :Avrupa standartlarına uygun")
	} else {
		fmt.Println("Sig : Sınıf dışı")
	}

	return pistols(canik, glock, sig)
	/*MARKALARIN AĞIRLIKLARININ AVRUPA STANDARTLARINA
	  UYGUN OLUP OLMADIĞININ İNCELENDİĞİ BÖLÜM..*/
}

func weigth(glockw float32, canikw float32, sigw float32) {

	if 200 < glockw && glockw < 700 {
		fmt.Println("Glock : Ağırlık standartlarına uygundur")
	} else {
		fmt.Println("Glock : Olması gereken ağırlıta değil, yan sanayi ürünü")
	}
	if 450 < canikw && canikw < 900 {
		fmt.Println("Canik :Ağırlık standartlarına uygundur")
	} else {
		if canikw > 1000 {
			var error float32 = canikw - 1000
			fmt.Println("Canik :Max ağırlığın üstünde, fark :", error)
		}
	}
	if 600 < sigw && sigw < 1250 {
		fmt.Println("Ağırlık standartlarına uygundur: sig ")
	} else {
		if sigw > 1250 {
			var error1 float32 = sigw - 1250
			fmt.Println(" Sig Max ağırlığın üstünde, fark :", error1)
		}
		if sigw < 600 {
			var error2 float32 = 600 - sigw
			fmt.Println("Sig minimum ağarlığın altında, fark :", error2)
		}

	}
	/*MARKALARIN NAMLU UZUNLUKLARININ AVRUPA STANDARTLARINA
	  UYGUN OLUP OLMADIĞIN İNCELENDİĞİ BÖLÜM..*/
}

func euBarrel(canikBar int, glockBar int, sigBar int) {

	var barrelLenght int = 102

	if canikBar > barrelLenght {
		fmt.Println("Canik Namlu uzunluğu = geçerli")
	} else {
		fmt.Println("Canik namlu Fazlalığı mm =", barrelLenght-canikBar)
	}
	if glockBar > barrelLenght {
		fmt.Println("Glock Namlu uzunluğu = geçerli")
	} else {
		fmt.Println("Glock namlu Fazlalığı mm =", barrelLenght-glockBar)
	}
	if sigBar > barrelLenght {
		fmt.Println("Sig Namlu uzunluğu = geçerli")
	} else {
		if sigBar == 102 {
			fmt.Println("Sig Namlu uzunluğu = kritik değer")
		} else {
			fmt.Println("Sig Namlu Fazlalığı mm =", barrelLenght-sigBar)
		}
	}
	return
	//MARKALARIN FİRMA YAŞLARININ İNCELENDİĞİ BÖLÜM..

}
func firmAge(canikComp, glockComp, sigComp int) int {

	var year int = 2022

	fmt.Printf("Canik Firma Yaşı : %v\n", year-canikComp)
	fmt.Printf("Glock Firma Yaşı : %v\n", year-glockComp)
	fmt.Printf("Sig Firma Yaşı : %v\n", year-sigComp)

	return year

}

func main() {
	canik := 16
	glock := 19
	sig := 11

	var glockw float32 = 452.25
	var canikw float32 = 1003.78
	var sigw float32 = 525.22

	pistols(canik, glock, sig)

	pistols := euStandCap(canik, glock, sig)
	fmt.Println("Şarjör Kapasiteleri Toplam :", pistols)
	weigth(glockw, canikw, sigw)

	canikBar := 110
	glockBar := 97
	sigBar := 102
	euBarrel(canikBar, glockBar, sigBar)

	canikComp := 2013
	glockComp := 1999
	sigComp := 1974

	firmAge(canikComp, glockComp, sigComp)

}
