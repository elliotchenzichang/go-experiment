package table_gen

import (
	"fmt"
	"os"
	"os/exec"
	"testing"
	"text/template"
)

func TestTemplateParse(t *testing.T) {
	//var structTemplate = `
	//
	//`
	type Column struct {
		ColumnName string
		DataType   string
	}

	var data = struct {
		PackageName string
		StructName  string
		Columns     []*Column
	}{
		PackageName: "table_gen",
		StructName:  "TestStruct",
		Columns: []*Column{
			{"Age", "int"},
			{"Name", "string"},
		},
	}
	//f, err := os.Create("test_gen.go")
	//if err != nil {
	//	fmt.Println(err)
	//	return
	//}
	//teml, _ := template.New("test").Parse(structTemplate)
	teml, _ := template.ParseFiles("struct_gen_template")
	teml.Execute(os.Stdout, data)
	fmt.Println("")
}

func TestGenerator_Build(t *testing.T) {
	config := &Config{
		Host:     "127.0.0.1",
		Port:     3306,
		Username: "root",
		Password: "12345678",
	}
	g, err := NewGenerator(config)
	if err != nil {
		t.Errorf("have err during NewGenerator, err is %s", err)
		return
	}
	genInfo := &GenInfo{
		Schema:       "elliot_test",
		Table:        "test_table",
		ExportFolder: "",
		TemplatePath: "struct_gen_test_template",
		FileName:     "test_gen.go",
		PackageName:  "table_gen",
		StructName:   "TestGenStruct",
	}
	isSuccess, err := g.Gen(genInfo)
	if err != nil {
		t.Errorf("have err during Gen file, err is %s", err)
		return
	}
	t.Logf("does it gen file successully?  %v", isSuccess)
}

func TestCommand(t *testing.T) {
	if _, e := exec.Command("go", "fmt", "test_gen.go").Output(); e != nil {
		t.Error(e)
	}
}
