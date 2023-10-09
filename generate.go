package generate

import (
	"fmt"

	"goa.design/goa/v3/codegen"
	"goa.design/goa/v3/eval"
	"goa.design/goa/v3/expr"
)

func addImports(t expr.UserType, imports map[string][]string) {
	fp := fmt.Sprintf("gen/%s/service.go", codegen.SnakeCase(t.Name()))
	for key, values := range t.Attribute().Meta {
		if key == "struct:pkg:path" {
			fp = fmt.Sprintf("gen/%s/%s.go", values[0], codegen.SnakeCase(t.Name()))
		}
		if key != "import" {
			continue
		}

		imports[fmt.Sprintf("gen/%s/views/view.go", codegen.SnakeCase(t.Name()))] = values
		imports[fp] = values
	}
}

// Generate adds the imports to the generated files.
func Generate(genpkg string, roots []eval.Root, files []*codegen.File) ([]*codegen.File, error) {
	imports := make(map[string][]string)
	for _, root := range roots {
		if r, ok := root.(*expr.RootExpr); ok {
			for _, t := range r.Types {
				addImports(t, imports)
			}
			for _, t := range r.ResultTypes {
				addImports(t, imports)
			}
		}
	}

	for _, f := range files {
		for key, importVals := range imports {
			if key == f.Path {
				for _, v := range importVals {
					codegen.AddImport(f.SectionTemplates[0], &codegen.ImportSpec{Path: v})
				}
			}
		}
	}

	return files, nil
}

func init() {
	codegen.RegisterPlugin("dynamic-imports", "gen", nil, Generate)
}
