package template

import (
	_ "embed"
	"strings"
)

//go:embed route/model_common.go
var ModelCommon string

//go:embed route/router_common.go
var RouterCommon string

//go:embed dao/common.go
var DaoCommon string

//go:embed route/router_getter.go
var __RouterGetter string

func init() {
	ModelCommon = clean(ModelCommon)
	DaoCommon = clean(DaoCommon)
	sb := strings.Builder{}
	sb.WriteString(clean(RouterCommon))
	sb.WriteString(clean(__RouterGetter))
	RouterCommon = sb.String()
}

func clean(o string) string {
	n := strings.Builder{}
	temp := strings.Builder{}

	importsStarted := false
	for _, c := range o {
		if c == '\n' {
			n.WriteRune(c)

			ts := strings.TrimSpace(temp.String())
			if importsStarted {
				if strings.HasPrefix(ts, ")") {
					importsStarted = false
				}
			} else if strings.HasPrefix(ts, "import") {
				content := strings.TrimSpace(strings.TrimPrefix(ts, "import"))
				if strings.HasPrefix(content, "(") {
					importsStarted = true
				}
			} else if strings.HasPrefix(ts, "package") {
			} else {
				n.WriteString(temp.String())
			}

			temp.Reset()
		} else {
			temp.WriteRune(c)
		}
	}

	return n.String()
}
