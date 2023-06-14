package util

import (
	"fmt"
	"os"
	"text/template"
)

type Renderable struct {
	Values              map[string]interface{}
	ComposeFileTemplate string
}

func (renderable Renderable) Render() {
	tpl := template.New("gotemplate")

	templateContent, err := os.ReadFile(renderable.ComposeFileTemplate)
	tpl, _ = tpl.Parse(string(templateContent))

	if err != nil {
		fmt.Println(err)
	}

	tpl.Execute(os.Stdout, renderable)
}
