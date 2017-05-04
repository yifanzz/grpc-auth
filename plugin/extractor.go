package plugin

import (
  "regexp"
  "strings"
  "fmt"
)

var (
  // regex used to extract substitution variables
  substitutionRegex = regexp.MustCompile("\\{(.+?)}")
)

// takes a string in the form of "prj/{deployment_id.project_id.name}/dpl/{deployment_id.name}"
// where field references are strings that are wrapped by `{...}`. The field references extracted from
// the exmaple would be `deployment_id.project_id.name` and `deployment_id.name` and the template
// string `prj/%v/dpl/%v`
func ExtractFieldPath(path string) *SubstitutablePath {
  matches := substitutionRegex.FindAllStringSubmatch(path, -1)

  var fieldPaths []*ProtoFieldPath
  for _, match := range matches {
    fieldPaths = append(fieldPaths, &ProtoFieldPath{
      protoFieldNames: strings.Split(match[1], "."),
    })
  }

  return &SubstitutablePath{
    protoFieldPaths: fieldPaths,
    template:        substitutionRegex.ReplaceAllString(path, "%v"),
  }
}

type SubstitutablePath struct {

  protoFieldPaths []*ProtoFieldPath

  // template should be a substitutable string in the form of "prj/%v/dpl/%v".
  // the number of substitution vars should match the length of `protoFieldNames`
  template string
}

func (p *SubstitutablePath) Generate(varName string) string {
  var fullPaths []string
  for _, protoFieldPath := range p.protoFieldPaths {
    fullPaths = append(fullPaths, fmt.Sprintf("%s.%s", varName, protoFieldPath.GoProtoField()))
  }
  return fmt.Sprintf(`fmt.Sprintf("%s", %v)`, p.template, strings.Join(fullPaths, ", "))
}

type ProtoFieldPath struct {
  // protoFieldNames contains the list of field references to be substituted in
  protoFieldNames []string
}

func (p *ProtoFieldPath) GoProtoField() string {
  var convertedField []string
  for _, elem := range p.protoFieldNames {
    convertedField = append(convertedField, underscoreToCamel(elem))
  }
  return strings.Join(convertedField, ".")
}

func underscoreToCamel(underscored string) string {
  parts := strings.Split(underscored, "_")
  var newParts []string
  for _, part := range parts {
    newParts = append(newParts, strings.Title(part))
  }
  return strings.Join(newParts, "")
}
