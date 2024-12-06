package file

import "strings"

func UnifySlashes(path string) string {
	return strings.ReplaceAll(strings.ReplaceAll(path, `\\`, `\`), `\`, `/`)
}
