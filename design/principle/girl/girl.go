package girl

import "fmt"

// 漂亮的女孩
type IPettyGirl interface {
	goodLooking()
	niceFigure()
	greatTemperament()
}

// 实现星探
type Searcher struct {
	pettyGirl IPettyGirl
}

func (s *Searcher) AbstractSearcher(i IPettyGirl) {
	s.pettyGirl = i
}

func (s Searcher) show() {
	fmt.Println("信息如下：")

	s.pettyGirl.goodLooking()
	s.pettyGirl.niceFigure()
	s.pettyGirl.greatTemperament()
}

// 实现女孩
type PettyGirl struct {
	name string
}

func (p *PettyGirl) SetName(n string) {
	p.name = n
}

func (p PettyGirl) goodLooking() {
	fmt.Println(p.name + "漂亮")
}

func (p PettyGirl) greatTemperament() {
	fmt.Println(p.name + "气质好")
}

func (p PettyGirl) niceFigure() {
	fmt.Println(p.name + "身材好")
}

func run() {
	girl := &PettyGirl{}
	girl.SetName("春")

	var iGirl IPettyGirl
	iGirl = girl

	searcher := &Searcher{}
	searcher.AbstractSearcher(iGirl)
	searcher.show()
}
