package main

import "fmt"

func main() {

	//Laptop almak için elimizdeki mevcut bütçeyi tanımladık.
	//3 faklı markanın 3 farklı fiyatlı laptoplarını farklı iki array olarak tanımladık.
	//Mevcut bütçe diye bir bool değişken tanımlayıp true değerini atadık.

	var totalBudget int = 17250
	var adequateBudget bool = true

	var laptopBrands [3]string = [3]string{"asus", "msi", "monster"}
	var laptopPrices [3]int = [3]int{12500, 15500, 18100}

	for i := 0; i < len(laptopBrands); i++ {
		fmt.Println(laptopBrands[i], "Laptop", laptopPrices[i])
		fmt.Println() //bir satır atlayıp çıktının anlaşılır olması için yazıldı!
	}

	//KDV fiyatını, priceİncludingVat fonksiyonunda 100'e böleceğimiz için 118 olarak tanımladık (%18KDV).
	//Vergili fiyatını hesaplamak için taxedPrice adında int bir değişken tanımlayıp 0 değerini atadık.

	for j := 0; j < len(laptopPrices); j++ {
		var taxPrice int = 118
		var taxedPrice int = 0
		taxedPrice = priceİncludingVat(taxPrice, laptopPrices[j])
		fmt.Println("kdvliFiyat :", taxedPrice)

		//Laptop fiyatlarını baz alarak, mevcut bütçemizi KDV dahilinde if-else olarak hesapladık.
		//Sırasıyla ; bütçemiz, laptop fiyatlarından fazla mı veya laptop almaya yetiyor mu ?

		if laptopPrices[j] <= totalBudget {
			fmt.Println("Satin Alinabilir : ", adequateBudget)

			//Bütçemiz, tam tamına herhangi bir markanın laptop fiyatına eşit mi'nin kod dizini

		} else {
			if laptopPrices[j] == totalBudget {
				fmt.Println("Satin Alinabilir :", adequateBudget)

				//Bütçemiz, hiçbir markanın laptobunu almaya yetimiyorsa alttaki else kod dizini çalışacaktır.

			} else {
				adequateBudget = false
				fmt.Println("Satin Alinabilir :", adequateBudget)
			}
		}

	}
}

//KDV hespalamak için aşağıdaki fonksiyonu tanımladık.
//
//Laptop fiyatı ile yukarıda tanımladığımız KDV(118) yi çarpıp 100'e böldüğünde
//
//Laptobun Vergi fiyatı dahil fiyatının ne kadar olacağını artık öğrenebiliriz.
func priceİncludingVat(x int, y int) (priceİncludingTax int) {

	taxCalculation := (x * y) / 100

	return taxCalculation
}
