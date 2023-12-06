package main

type ParkingSystem struct {
	lots []int
}

func ParkingSystemConstructor(big int, medium int, small int) *ParkingSystem {
	ps := ParkingSystem{[]int{big, medium, small}}
	return &ps
}

func (this *ParkingSystem) AddCar(carType int) bool {
	if this.lots[carType-1] > 0 {
		this.lots[carType-1]--
		return true
	}
	return false
}

// Switch case
// type ParkingSystem struct {
// 	big    int
// 	medium int
// 	small  int
// }

// func ParkingSystemConstructor(big int, medium int, small int) *ParkingSystem {
// 	ps := ParkingSystem{big, medium, small}
// 	return &ps
// }

// func (this *ParkingSystem) AddCar(carType int) bool {
// 	switch carType {
// 	case 1:
// 		if this.big > 0 {
// 			this.big--
// 			return true
// 		}
// 	case 2:
// 		if this.medium > 0 {
// 			this.medium--
// 			return true
// 		}
// 	case 3:
// 		if this.small > 0 {
// 			this.small--
// 			return true
// 		}
// 	}
// 	return false
// }

// func main() {
// 	var re []bool
// 	ps := Constructor(1, 1, 0)
// 	re = append(re, ps.AddCar(1))
// 	re = append(re, ps.AddCar(2))
// 	re = append(re, ps.AddCar(3))
// 	re = append(re, ps.AddCar(1))
// 	fmt.Println(re)
// }
