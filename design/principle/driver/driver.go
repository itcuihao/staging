package driver

import "fmt"

type Benz struct {
}

func (b Benz) run() {
	fmt.Println("benz run. ")
}

type Bmw struct {
}

func (b Bmw) run() {
	fmt.Println("bmw run. ")
}

type ICar interface {
	run()
}

type IDriver interface {
	// 构造函数传递
	Drive(car ICar)

	// 方法传递
	SetCar(car ICar)
	DriveCar()
}

type Driver struct {
	Car ICar
}

// 司机只能开奔驰，假如有宝马则没法开
// func (d Driver) drive(b *Benz) {
// 	b.run()
// }

func (d *Driver) Drive(car ICar) {
	car.run()
}

func (d *Driver) SetCar(car ICar) {
	d.Car = car
}

func (d *Driver) DriveCar() {
	d.Car.run()
}
