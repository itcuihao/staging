package gorm_explain

import (
	"testing"

	"github.com/itcuihao/staging/db"
)

func TestExplain(t *testing.T) {
	dao := db.InitDebug(&db.MysqlConfig{})

	// ExplainOutDefault 在控制台输出
	//SqlExplain(dao.db, 3, ExplainOutDefault)

	// ExplainOutCsv 输出csv文件
	//SqlExplain(dao.db, 3, ExplainOutCsv)

	// ExplainOutWeb 输出到浏览器
	SqlExplain(dao.DB, 3, ExplainOutWeb)
	defer OutExplain()
	// 下面为需要查看explain的方法
	//dao.GetWorkXx(0)
	//dao.GetWorkXxx(0, "a")
	//dao.GetWorkXxx("a")
}
