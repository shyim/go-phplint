package tester

import "github.com/google/go-cmp/cmp"

// Golden parser trees were recorded before columns were tracked. Comparisons
// still check offsets and lines.
var ignorePositionColumns = cmp.FilterPath(func(path cmp.Path) bool {
	field, ok := path.Index(-1).(cmp.StructField)
	return ok && (field.Name() == "StartCol" || field.Name() == "EndCol")
}, cmp.Ignore())
