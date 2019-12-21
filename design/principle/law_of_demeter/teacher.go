package law_of_demeter

import "fmt"

type Teacher struct {
}

func (t Teacher) commond(groupLeader GroupLeader) {
	// 此处 Girl 不属于 Teacher 的朋友类，出现在此处，产生了依赖关系，违背了迪米特法则
	// listGirls := make([]*Girl, 20)
	//groupLeader.countGirl(listGirls)

	groupLeader.countGirl()
}

type Girl struct {
}

type GroupLeader struct {
	listGirls []*Girl
}

func (g *GroupLeader) newListGirls(girls []*Girl) {
	g.listGirls = girls
}

//func (g GroupLeader) countGirl(girls []*Girl) {
	//fmt.Println("女生数量是：", len(girls))
//}

func (g GroupLeader) countGirl() {
	fmt.Println("女生数量是：", len(g.listGirls))
}
