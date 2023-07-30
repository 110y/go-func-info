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
			pos:  95,
			expected: &analysis.FuncInfo{
				Name: "F1",
				Results: []*analysis.ResultInfo{
					{
						Name:      "",
						TypeName:  "int",
						ZeroValue: "0",
					},
				},
				StartPos: 70,
				EndPos:   109,
			},
		},
		"should return TestInfo includes expected TestFuncName and SubTestNames": {
			path: "package1/file1_test.go",
			pos:  124,
			expected: &analysis.FuncInfo{
				Name:     "F2",
				StartPos: 111,
				EndPos:   149,
			},
		},
		"should return a custom type receiver": {
			path: "package1/file1_test.go",
			pos:  204,
			expected: &analysis.FuncInfo{
				Name:     "F4",
				StartPos: 179,
				EndPos:   211,
				Receiver: &analysis.ReceiverInfo{
					Name:     "t",
					TypeName: "*github.com/110y/go-func-info/internal/analysis/testdata/package1.MyType",
				},
			},
		},
		"should return expected FunInfo and ResultInfo": {
			path: "package1/file1_test.go",
			pos:  233,
			expected: &analysis.FuncInfo{
				Name:     "F5",
				StartPos: 213,
				EndPos:   250,
				Results: []*analysis.ResultInfo{
					{
						TypeName:  "github.com/110y/go-func-info/internal/analysis/testdata/package1.MyType",
						ZeroValue: "MyType{}",
					},
				},
			},
		},
		"should return expected FunInfo and ResultInfo with error": {
			path: "package1/file1_test.go",
			pos:  278,
			expected: &analysis.FuncInfo{
				Name:     "F6",
				StartPos: 252,
				EndPos:   293,
				Results: []*analysis.ResultInfo{
					{
						TypeName:  "int",
						ZeroValue: "0",
					},
					{
						TypeName:  "error",
						ZeroValue: "nil",
					},
				},
			},
		},
		"should return expected ResultInfo with custom Interface": {
			path: "package1/file1_test.go",
			pos:  338,
			expected: &analysis.FuncInfo{
				Name:     "F7",
				StartPos: 295,
				EndPos:   346,
				Results: []*analysis.ResultInfo{
					{
						TypeName:  "github.com/110y/go-func-info/internal/analysis/testdata/package1.MyInterface",
						ZeroValue: "nil",
					},
					{
						TypeName:  "error",
						ZeroValue: "nil",
					},
				},
			},
		},
		"should return expected ResultInfo with a pointer to the custom Struct": {
			path: "package1/file1_test.go",
			pos:  378,
			expected: &analysis.FuncInfo{
				Name:     "F8",
				StartPos: 348,
				EndPos:   395,
				Results: []*analysis.ResultInfo{
					{
						TypeName:  "*github.com/110y/go-func-info/internal/analysis/testdata/package1.MyType",
						ZeroValue: "nil",
					},
					{
						TypeName:  "error",
						ZeroValue: "nil",
					},
				},
			},
		},
		"should return expected ResultInfo with custom Struct": {
			path: "package1/file1_test.go",
			pos:  426,
			expected: &analysis.FuncInfo{
				Name:     "F9",
				StartPos: 397,
				EndPos:   448,
				Results: []*analysis.ResultInfo{
					{
						TypeName:  "github.com/110y/go-func-info/internal/analysis/testdata/package1.MyType",
						ZeroValue: "MyType{}",
					},
					{
						TypeName:  "error",
						ZeroValue: "nil",
					},
				},
			},
		},
		"should return expected ResultInfo with custom generic struct receiver": {
			path: "package1/file1_test.go",
			pos:  482,
			expected: &analysis.FuncInfo{
				Name:     "F10",
				StartPos: 450,
				EndPos:   493,
				Receiver: &analysis.ReceiverInfo{
					Name:     "m",
					TypeName: "github.com/110y/go-func-info/internal/analysis/testdata/package1.MyType",
				},
				Results: []*analysis.ResultInfo{
					{
						TypeName:  "error",
						ZeroValue: "nil",
					},
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
