package gen

import (
	"github.com/wuhc1010/go-zero/tools/goctl/model/sql/template"
	"github.com/wuhc1010/go-zero/tools/goctl/util"
	"github.com/wuhc1010/go-zero/tools/goctl/util/pathx"
)

func genTableName(table Table) (string, error) {
	text, err := pathx.LoadTemplate(category, tableNameTemplateFile, template.TableName)
	if err != nil {
		return "", err
	}

	output, err := util.With("tableName").
		Parse(text).
		Execute(map[string]any{
			"tableName":             table.Name.Source(),
			"upperStartCamelObject": table.Name.ToCamel(),
		})
	if err != nil {
		return "", nil
	}

	return output.String(), nil
}
