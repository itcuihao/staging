package driver

import "testing"

func TestDrive(t *testing.T) {
	var idriver IDriver
	idriver = new(Driver)

	// 构造函数传递对象
	benz := new(Benz)
	// driver.drive(benz)
	idriver.Drive(benz)

	bmw := new(Bmw)
	idriver.Drive(bmw)
}

func TestDriveCar(t *testing.T) {
	var idriver IDriver
	idriver = new(Driver)

	// 构造方法传递对象
	benz := new(Benz)
	idriver.SetCar(benz)
	idriver.DriveCar()

	bmw := new(Bmw)
	idriver.SetCar(bmw)
	idriver.DriveCar()
}
