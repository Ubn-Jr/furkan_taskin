package main

import "fmt"

func main() {

	var weight float32 = 165.50
	fmt.Println("Your current weight : ", weight)
	var desiredWeight float32 = 105.0
	fmt.Println("The weight you want to reach :", desiredWeight)
	fmt.Println("*****************")
	var dailyFatPercentage float32 = 1.5

	//aşağıdaki fonksiyondan dönen değer ekrana
	//kaybedilen kilo ve şuan ulaşmış olunan kilo olarak yazılacaktır
	weightLost := weightLoss(&weight, desiredWeight, dailyFatPercentage)
	fmt.Println("Your new weight : ", weight)
	fmt.Println("Your lost weight : ", weightLost)
}

//kilo ölçümü yapılacak kişinin spora başlamadan önce ki kilosu ve
//bir günde yakacağı yağ oranı mevcut kilosundan çıkarılarak
//günlük bir döngüyle mevcut kilosundan düşecektir.

func weightLoss(weight *float32, desiredWeight float32, dailyFatPercentage float32) float32 {
	startingWeight := *weight
	for *weight > desiredWeight {
		fmt.Println(*weight)
		*weight -= dailyFatPercentage
	}
	return startingWeight - *weight
}
