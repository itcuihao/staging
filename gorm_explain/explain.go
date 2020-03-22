package gorm_explain

import (
	"database/sql"
	"encoding/csv"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/bndr/gotabulate"
	"github.com/jinzhu/gorm"
)

type ExplainType int

const (
	ExplainOutDefault ExplainType = iota
	ExplainOutCsv
	ExplainOutWeb
)

type Explain struct {
	out ExplainType
}

type explainStore struct {
	out    ExplainType
	detail []*explainSql
}

type explainSql struct {
	sql       string
	resultStr string
	columns   []string
	results   [][]interface{}
}

var (
	store *explainStore
)

func SqlExplain(db *gorm.DB, num int, out ExplainType) {
	initStore(num*2, out)
	e := &Explain{out: out}

	db.Callback().Query().After("gorm:query").Register("my_plugin:explain", e.CallbackSelect)
}

func (e *Explain) CallbackSelect(scope *gorm.Scope) {

	if !strings.HasPrefix(strings.ToUpper(scope.SQL), "SELECT") {
		return
	}

	rows, err := scope.SQLDB().Query("EXPLAIN "+scope.SQL, scope.SQLVars...)
	if scope.Err(err) != nil {
		return
	}
	defer rows.Close()

	columns, _ := rows.Columns()
	results, _ := e.convertToResult(rows)

	detail := &explainSql{}
	switch e.out {
	case ExplainOutCsv:
		detail.sql = scope.SQL
		detail.columns = columns
		detail.results = results
	case ExplainOutWeb:
		detail.sql = e.outTableSql(scope.SQL, scope.SQLVars)
		detail.resultStr = e.outTableExplain(columns, results)
	default:
		detail.sql = e.outTableSql(scope.SQL, scope.SQLVars)
		detail.resultStr = e.outTableExplain(columns, results)
	}

	store.detail = append(store.detail, detail)
}

func (e *Explain) outTableSql(sql string, values []interface{}) string {
	sqlTitle := []string{"sql", "value"}
	sqlValue := [][]interface{}{[]interface{}{sql, fmt.Sprintf("%+v", values)}}
	sqlTable := gotabulate.Create(sqlValue)
	sqlTable.SetHeaders(sqlTitle)
	return sqlTable.Render("simple")
}

func (e *Explain) outTableExplain(columns []string, results [][]interface{}) string {
	table := gotabulate.Create(results)
	table.SetHeaders(columns)
	table.SetEmptyString("None")
	table.SetAlign("right")
	return table.Render("grid")
}

func (e *Explain) convertToResult(rows *sql.Rows) (res [][]interface{}, err error) {
	columns, err := rows.Columns()
	if err != nil {
		return nil, err
	}

	values := make([]sql.RawBytes, len(columns))
	args := make([]interface{}, len(values))

	for i := range values {
		args[i] = &values[i]
	}

	for rows.Next() {
		if err := rows.Scan(args...); err != nil {
			return nil, err
		}
		row := []interface{}{}
		for _, col := range values {
			row = append(row, string(col))
		}
		res = append(res, row)
	}

	return res, nil
}

func initStore(size int, o ExplainType) {
	store = &explainStore{out: o, detail: make([]*explainSql, 0, size)}
}

func OutExplain() {
	fmt.Println("all sql number:", len(store.detail))
	switch store.out {
	case ExplainOutCsv:

		file, err := os.OpenFile(
			"explain.csv",
			os.O_WRONLY|os.O_TRUNC|os.O_CREATE,
			0666,
		)
		if err != nil {
			log.Printf("error create file:%v\n", err)
			return
		}
		defer file.Close()

		w := csv.NewWriter(file)
		defer w.Flush()

		if err := w.Write([]string{fmt.Sprintf("all sql number:%d", len(store.detail))}); err != nil {
			log.Printf("error writing record to csv:%v\n", err)
		}
		for _, d := range store.detail {

			if err := w.Write([]string{d.sql}); err != nil {
				log.Printf("error writing record to csv:%v\n", err)
				continue
			}
			if err := w.Write(d.columns); err != nil {
				log.Printf("error writing record to csv:%v\n", err)
				continue
			}
			for _, result := range d.results {
				record := make([]string, 0, len(result))
				for _, r := range result {
					record = append(record, fmt.Sprintf("%v", r))
				}
				if err := w.Write(record); err != nil {
					log.Printf("error writing record to csv:%v\n", err)
					continue
				}
			}

			if err := w.Write([]string{""}); err != nil {
				log.Printf("error writing record to csv:%v\n", err)
				continue
			}
		}
	case ExplainOutWeb:
		var sb strings.Builder
		for _, exp := range store.detail {
			if exp == nil {
				sb.WriteString("")
				continue
			}
			sb.WriteString(exp.sql)
			sb.WriteString(exp.resultStr)
		}
		httpExplain(sb.String())
	default:
		fmt.Println("all sql number:", len(store.detail))
		for _, exp := range store.detail {
			if exp == nil {
				fmt.Println("")
				continue
			}
			fmt.Printf("%+v\n", exp.sql)
			fmt.Printf("%+v\n", exp.resultStr)
		}
	}
}

func httpExplain(str string) {
	fmt.Println("please click: http://localhost:1030/explain")
	mux := http.NewServeMux()
	mux.HandleFunc("/explain", func(w http.ResponseWriter, req *http.Request) {
		w.Write([]byte(str))
	})
	http.ListenAndServe(":1030", mux)
}
