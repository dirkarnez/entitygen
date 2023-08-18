package main

import (
	"bytes"
	"fmt"
	"log"
	"os"

	"github.com/pingcap/tidb/parser"
	"github.com/pingcap/tidb/parser/ast"
	_ "github.com/pingcap/tidb/types/parser_driver"
)

func main() {
	file, err := os.Open("input.sql")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	buf := new(bytes.Buffer)
	_, err = buf.ReadFrom(file)
	if err != nil {
		log.Fatal(err)
	}

	src := buf.String()

	p := parser.New()
	stmtNode, _ := p.ParseOneStmt(src, "", "")

	if createTableStmt, ok := stmtNode.(*ast.CreateTableStmt); ok {
		// Handle CREATE TABLE statement
		fmt.Println("Table name:", createTableStmt.Table.Name.String())
		fmt.Println("Columns:")
		for _, colDef := range createTableStmt.Cols {
			fmt.Println(colDef.Name.Name.L, colDef.Tp.InfoSchemaStr())
		}
	}
}
