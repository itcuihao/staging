package law_of_demeter

import "testing"

func TestCountGirl(t *testing.T) {
	teacher := new(Teacher)
	var groupLeader  GroupLeader
	teacher.commond(groupLeader)
}

func TestCountGirl2(t *testing.T) {
	listGirls:=make([]*Girl,21)
	teacher := new(Teacher)
	var groupLeader  GroupLeader
	groupLeader.newListGirls(listGirls)
	teacher.commond(groupLeader)
}
