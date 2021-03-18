package main

type ParkingSystem struct {
	big, medium, small int
}

func (p *ParkingSystem) Init(a, b, c int) ParkingSystem {
	p.small, p.medium, p.big = c, b, a
	return *p
}

func Constructor(big int, medium int, small int) ParkingSystem {
	return new(ParkingSystem).Init(big, medium, small)
}

func (p *ParkingSystem) AddCar(carType int) bool {
	switch carType {
	case 1:
		if p.big > 0 {
			p.big--
			return true
		}
	case 2:
		if p.medium > 0 {
			p.medium--
			return true
		}
	case 3:
		if p.small > 0 {
			p.small--
			return true
		}
	}
	return false
}

/**
 * Your ParkingSystem object will be instantiated and called as such:
 * obj := Constructor(big, medium, small);
 * param_1 := obj.AddCar(carType);
 */
