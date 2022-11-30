package main

import "fmt"

type creditCard struct {
	customerName string
	cardID       int
	password     int
}

func main() {

	var (
		moneyInAccount float32 = 15430.32
		withdrawMoney  float32
	)
	hesapP := &moneyInAccount

	var user creditCard = creditCard{
		customerName: "FURKAN TAŞKIN",
		cardID:       3210,
		password:     1905,
	}

	fmt.Printf("Lütfen Kartınızın Şifresini Giriniz : ")
	fmt.Scanln(&user.password)

	if user.cardID == 3210 && user.password == 1905 {

		fmt.Printf("HOŞ GELDİN %s\n", user.customerName)
		fmt.Println("Hesapta Olan Paranız : ", moneyInAccount)
		fmt.Printf("Çekmek İstediğiniz Para Miktarını Giriniz :")
		fmt.Scanln(&withdrawMoney)
		moneyOutOfAccount(&moneyInAccount, withdrawMoney)

		for i := 0; i < 1; i++ {
			if withdrawMoney >= 3000 || withdrawMoney <= 0 {

				fmt.Println("Üzgünüz Günlük Limiti Aştınız..")

			} else {
				fmt.Println("Makbuz Yazırıldı Kalan Para : ", *hesapP)
			}
		}

	} else {
		fmt.Println("Yanlış Şifre Girdiniz Kart İade Ediliyor..")

	}

}
func moneyOutOfAccount(moneyInAccount *float32, withdrawMoney float32) {

	*moneyInAccount = *moneyInAccount - withdrawMoney

}
