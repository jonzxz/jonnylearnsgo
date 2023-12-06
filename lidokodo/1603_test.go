package main

import "testing"

func TestParkingSystem(t *testing.T) {
	system := ParkingSystemConstructor(1, 1, 0)

	want := []bool{true, true, false, false}
	got := parkCars(system, []int{1, 2, 3, 1})

	for i := range want {
		if want[i] != got[i] {
			t.Errorf("Want %v, Got %v\n", want, got)
			return
		}
	}
}

func parkCars(system *ParkingSystem, cars []int) []bool {
	var results []bool
	for _, v := range cars {
		results = append(results, system.AddCar(v))
	}
	return results
}
