package plugin

import "testing"

func TestExtractFieldPath(t *testing.T) {
	for _, run := range []struct{
		name string
		path string
		expectedOutput string
	}{
		{
			name: "path without templating",
			path: "prj/foo",
			expectedOutput: `fmt.Sprintf("prj/foo", )`,
		},
		{
			name: "single template var",
			path: "prj/{some_path}",
			expectedOutput: `fmt.Sprintf("prj/%v", msg.SomePath)`,
		},
		{
			name: "multiple template var",
			path: "prj/{some_path}/yoyo/{other_path}",
			expectedOutput: `fmt.Sprintf("prj/%v/yoyo/%v", msg.SomePath, msg.OtherPath)`,
		},
		{
			name: "nested field",
			path: "prj/{some_path.haha}",
			expectedOutput: `fmt.Sprintf("prj/%v", msg.SomePath.Haha)`,
		},
		{
			name: "multiple mixed template var",
			path: "prj/{some_path.foo}/yoyo/{other_path}",
			expectedOutput: `fmt.Sprintf("prj/%v/yoyo/%v", msg.SomePath.Foo, msg.OtherPath)`,
		},

	}{
		path:= run.path
		expectedPath := run.expectedOutput
		t.Run(run.name, func(subT *testing.T) {
			substitutablePath := ExtractFieldPath(path)
			subed := substitutablePath.Generate("msg")
			if subed !=  expectedPath {
				subT.Errorf("Expected('%s') Actual('%s')", expectedPath, subed)
			}
		})
	}
}
