package plugin

import (
  "regexp"
)

var (
  substitutionRegex = regexp.MustCompile("(.*?\\{(.+?)\\})+?")
)

// takes a string in the form of "prj/{deployment_id.project_id.name}/dpl/{deployment_id.name}"
// where field references are strings that are wrapped by `{...}`. The field references extracted from
// the exmaple would be `deployment_id.project_id.name` and `deployment_id.name` and the template
// string `prj/%v/dpl/%v`
func ExtractFieldPath(path string) []*FieldPath {
  return nil
}

type SubstitutablePath struct {

  fieldPaths []*FieldPath

  // template should be a substutable string in the form of "prj/%v/dpl/%v".
  // the number of substitution vars should match the length of `fields`
  template string
}

type FieldPath struct {
  // fields contains the list of field references to be substituted in
  fields []string
}

func (p *FieldPath) GoProtoField() string {
  return ""
}
