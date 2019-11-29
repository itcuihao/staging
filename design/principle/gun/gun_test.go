package gun

import "testing"

func TestSoldier(t *testing.T) {
	soldier := new(Soldier)

	handgun := new(Handgun)
	soldier.SetGun(handgun)
	soldier.KillEnemy()

	machineGun := new(MachineGun)
	soldier.SetGun(machineGun)
	soldier.KillEnemy()

	rifle := new(Rifle)
	soldier.SetGun(rifle)
	soldier.KillEnemy()
}

func TestSnipper(t *testing.T) {
	snipper := new(Snipper)

	// aug := new(Aug)
	// snipper.SetRifle(aug)

	rifle:=new(Rifle)
	snipper.SetRifle(rifle)
	snipper.KillEnemy()
}
