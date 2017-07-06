package plugin

import "text/template"

var simpleAuthTemplate = template.Must(template.New("simple_auth").Parse(`
package grpc_auth_test
`))