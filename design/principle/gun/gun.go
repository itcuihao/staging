package gun

import "fmt"

// 枪的接口
type IGun interface {
	Shoot()
}

// 普通枪
type Gun struct {
}

// 射击
func (g Gun) Shoot() {
}

// 手枪
type Handgun struct {
	Gun
}

// 射击
func (h Handgun) Shoot() {
	fmt.Println("Handgun shoot.")
}

// 步枪
type Rifle struct {
	Gun
}

func (r Rifle) Shoot() {
	fmt.Println("Rifle shoot.")
}

// 机枪
type MachineGun struct {
	Gun
}

func (m MachineGun) Shoot() {
	fmt.Println("MachineGun shoot.")
}

// 士兵
type Soldier struct {
	gun IGun
}

// 分配枪支
func (s *Soldier) SetGun(gun IGun) {
	s.gun = gun
}

// 杀敌
func (s Soldier) KillEnemy() {
	fmt.Println("kill enemy...")
	s.gun.Shoot()
}

// 狙击步枪
type Aug struct {
	Rifle
}

// 望远镜
func (a Aug) ZoomOut() {
	fmt.Println("zoom out...")
}

// 步枪接口
type IRifle interface {
	IGun
	ZoomOut()
}

// 狙击手
type Snipper struct {
	gun IRifle
}

func (s *Snipper) SetRifle(gun IRifle) {
	s.gun = gun
}

func (s Snipper) KillEnemy() {
	fmt.Println("kill enemy...")
	s.gun.ZoomOut()
	// 这里可以使用 IGun 的 shoot 方法
	s.gun.Shoot()
}

type ToyGun struct {

}
