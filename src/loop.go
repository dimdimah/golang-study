package main

import "fmt"

func main() {
	// points := 5
	// nilai := 8849

	// if points == 10 {
	// 	fmt.Println("bagus")
	// } else if points > 5 {
	// 	fmt.Println("Lumayan")
	// } else if points == 4 {
	// 	fmt.Println("Kurang")
	// } else {
	// 	fmt.Println("Sangat Kurang")
	// }

	// if percents := nilai / 100; percents >= 75 {
	// 	fmt.Println("Nilai kamu", percents, "Sempurna")
	// } else if percents >= 50 {
	// 	fmt.Println("Nilai kamu", percents, "Baik")
	// } else {
	// 	fmt.Println("Nilai kamu", percents, "Kurang")
	// }

	// switch {
	// case points == 9:
	// 	fmt.Println("Sempurna")
	// case points >= 8:
	// 	fmt.Println("Hampir sempurna")
	// default:
	// 	fmt.Println("Kurang")
	// }

	// if points > 8 {
	// 	switch points {
	// 	case 10:
	// 		fmt.Println("Perfect")
	// 	default:
	// 		fmt.Println("Good Job")
	// 	}
	// } else {
	// 	if points > 6 {
	// 		fmt.Println("better")
	// 	} else if points >= 3 {
	// 		fmt.Println("nice")
	// 	} else {
	// 		fmt.Println("try again")
	// 	}
	// }

	// var i = 0

	// for i < points {
	// 	fmt.Println("Angka ", i)
	// 	i++
	// }

	// var xs = [5]int{10, 20, 30, 40, 50}
	// for _, v := range xs {
	// 	fmt.Println("value= ", v)
	// }

	// var zx = xs[0:2]
	// for _, v := range zx {
	// 	fmt.Println("value= ", v)
	// }
	// var fruits = [5]string{"jeruk", "ceri", "apel", "alpukat", "lemon"}
	// fmt.Println("jumlah element ", len(fruits))

	ayam := map[string]int{}

	ayam["ayam boiler"] = 50
	ayam["ayam bakar"] = 40

	fmt.Println(ayam["ayam bakar"])
}
