package generator

import (
	"bytes"
	"fmt"
	"os"
	"reflect"
	"strings"

	"golang.org/x/tools/imports"
)

func Add(pkgPath string, filename string, modelTargets ...Target) {
	for _, target := range modelTargets {
		modelType := reflect.TypeOf(target.Model)
		for modelType.Kind() == reflect.Pointer {
			modelType = modelType.Elem()
		}
		pkgParts := strings.Split(modelType.PkgPath(), "/")
		source := bytes.NewBuffer([]byte{})
		source.WriteString(fmt.Sprintf("package %s\n", pkgParts[len(pkgParts)-1]))

		source.WriteString(fmt.Sprintf("func (m *%s) Create(db *sql.DB) error {\n", modelType.Name()))
		source.WriteString("columnString := \"\"\n")
		source.WriteString("columnValueString := \"\"\n\n")

		fieldNames := []string{}
		fieldJsonNames := []string{}
		for i := 0; i < modelType.NumField(); i++ {
			field := modelType.Field(i)
			if field.Tag.Get("json") == "-" {
				continue
			}
			fieldName := field.Name
			fieldNames = append(fieldNames, fieldName)
			fieldJsonName := field.Name
			fieldJsonNames = append(fieldJsonNames, fieldJsonName)

			fieldType := field.Type
			for fieldType.Kind() == reflect.Pointer {
				fieldType = fieldType.Elem()
			}

			switch fieldType.Kind() {
			case reflect.String:
				source.WriteString(fmt.Sprintf("if m.%s != \"\" {\n", fieldName))
				source.WriteString(fmt.Sprintf("columnString += \"%s\"\n", fieldJsonName))
				source.WriteString("columnString += \",\"\n")
				source.WriteString("columnValueString += \"\\\"\"\n")
				source.WriteString(fmt.Sprintf("columnValueString += m.%s\n", fieldName))
				source.WriteString("columnValueString += \"\\\",\"\n")
				source.WriteString("}\n\n")
			case reflect.Int:
				source.WriteString(fmt.Sprintf("columnString += \"%s\"\n", fieldJsonName))
				source.WriteString("columnString += \",\"\n")
				source.WriteString(fmt.Sprintf("columnValueString += strconv.Itoa(m.%s)\n", fieldName))
				source.WriteString("columnValueString += \",\"\n\n")
			default:
				if fieldType.Implements(typeOfStringer) {
					source.WriteString(fmt.Sprintf("if m.%s.String() != \"\" {\n", fieldName))
					source.WriteString(fmt.Sprintf("columnString += \"%s\"\n", fieldJsonName))
					source.WriteString("columnString += \",\"\n")
					source.WriteString("columnValueString += \"\\\"\"\n")
					source.WriteString(fmt.Sprintf("columnValueString += m.%s.String()\n", fieldName))
					source.WriteString("columnValueString += \"\\\",\"\n")
					source.WriteString("}\n\n")
				} else {
					source.WriteString(fmt.Sprintf("columnString += \"%s\"\n", fieldJsonName))
					source.WriteString("columnString += \",\"\n")
					source.WriteString(fmt.Sprintf("columnValueString += fmt.Sprintf(\"%%v\", m.%s)\n", fieldName))
					source.WriteString("columnValueString += \",\"\n\n")
				}
			}
		}
		source.WriteString(fmt.Sprintf("query := \"insert into %s (\"\n"+
			"query += columnString\n"+
			"query += \") values (\"\n"+
			"query += columnValueString\n"+
			"query += \")\"\n", target.TableName))

		source.WriteString("_, err := db.Exec(query)\n")
		source.WriteString("if err != nil {\n")
		source.WriteString("return err\n")
		source.WriteString("}\n")
		source.WriteString("return nil\n")
		source.WriteString("}\n")

		data, err := imports.Process("", source.Bytes(), &imports.Options{
			TabIndent: true,
		})
		if err != nil {
			fmt.Println("do imports failed: ", err.Error())
			return
		}

		os.WriteFile(pkgPath+"/"+filename, data, 0777)
	}
}
