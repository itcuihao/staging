package gun

import "fmt"

type IGun interface {
	Shoot()
}

type Gun struct {
}

func (g Gun) Shoot() {

}

type Handgun struct {
}

func (h Handgun) Shoot() {
	fmt.Println("Handgun shoot.")
}

type Rifle struct {
}

func (r Rifle) Shoot() {
	fmt.Println("Rifle shoot.")
}

type MachineGun struct {
}

func (m MachineGun) Shoot() {
	fmt.Println("MachineGun shoot.")
}

type Soldier struct {
	gun IGun
}

func (s *Soldier) SetGun(gun IGun) {
	s.gun = gun
}

func (s Soldier) KillEnemy() {
	fmt.Println("kill enemy...")
	s.gun.Shoot()
}

// 另类枪
type Aug struct {
	Rifle
}

func (a Aug) ZoomOut() {
	fmt.Println("zoom out...")
}

type IRifle interface{
	IGun
	ZoomOut()
}

type Snipper struct{
	gun IRifle
}

func (s *Snipper) SetRifle(gun IRifle) {
	s.gun = gun
}

func (s Snipper) KillEnemy() {
	fmt.Println("kill enemy...")
	s.gun.ZoomOut()
	s.gun.Shoot()
}