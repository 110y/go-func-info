package analysis_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/google/go-cmp/cmp"

	"github.com/110y/go-func-info/internal/analysis"
)

func TestGetFuncInfo(t *testing.T) {
	t.Parallel()

	tests := map[string]struct {
		path     string
		pos      int
		expected *analysis.FuncInfo
	}{
		"should return expected FunInfo": {
			path: "package1/file1_test.go",
			pos:  62,
			expected: &analysis.FuncInfo{
				Name:     "F1",
				StartPos: 40,
				EndPos:   79,
			},
		},
		"should return TestInfo includes expected TestFuncName and SubTestNames": {
			path: "package1/file1_test.go",
			pos:  95,
			expected: &analysis.FuncInfo{
				Name:     "F2",
				StartPos: 81,
				EndPos:   119,
			},
		},
		"xx": {
			path: "package1/file1_test.go",
			pos:  175,
			expected: &analysis.FuncInfo{
				Name:     "F4",
				StartPos: 149,
				EndPos:   181,
				Receiver: &analysis.ReceiverInfo{
					Name:     "t",
					TypeName: "MyType",
				},
			},
		},
	}

	for name, test := range tests {
		test := test
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			path := fmt.Sprintf("testdata/%s", test.path)
			actual, err := analysis.GetFuncInfo(context.Background(), path, test.pos)
			if err != nil {
				t.Fatalf("error: %s\n", err.Error())
			}

			if diff := cmp.Diff(test.expected, actual); diff != "" {
				t.Errorf("\n(-expected, +actual)\n%s", diff)
			}
		})
	}
}
