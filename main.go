package main

import (
	"fmt"

	"github.com/pingcap/tidb/parser"
	_ "github.com/pingcap/tidb/types/parser_driver"
)

func main() {
	p := parser.New()
	fmt.Print(p)
	src := fmt.Sprintf("SELECT * FROM demo;")
	stmtNode, _ := p.ParseOneStmt(src, "", "")
	fmt.Print(stmtNode)

}
