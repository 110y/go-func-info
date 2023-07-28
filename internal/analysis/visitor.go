package analysis

import (
	"fmt"
	"go/ast"
	"go/token"
	"go/types"
)

var _ ast.Visitor = (*visitor)(nil)

var builtinTypeZeroValues = map[string]string{
	"bool":       "false",
	"uint8":      "0",
	"uint16":     "0",
	"uint32":     "0",
	"uint64":     "0",
	"int8":       "0",
	"int16":      "0",
	"int32":      "0",
	"int64":      "0",
	"float32":    "0",
	"float64":    "0",
	"complex64":  "0",
	"complex128": "0",
	"string":     `""`,
	"int":        "0",
	"uint":       "0",
	"uintptr":    "0",
	"byte":       "0",
	"rune":       "0",
	"any":        "nil",
	"comparable": "nil",
}

func newVisitor(pos int, fs *token.FileSet, info *types.Info) *visitor {
	return &visitor{
		fileset:   fs,
		cursorPos: pos,
		info:      info,
	}
}

type visitor struct {
	cursorPos int
	fileset   *token.FileSet
	info      *types.Info

	funcInfo *FuncInfo
}

func (v *visitor) Visit(node ast.Node) ast.Visitor {
	if node == nil {
		return nil
	}

	startPos := v.getPositionOffset(node.Pos())
	endPos := v.getPositionOffset(node.End())

	if v.cursorPos < startPos || v.cursorPos > endPos {
		return nil
	}

	if fd, ok := node.(*ast.FuncDecl); ok {
		if fd.Name == nil {
			return nil
		}

		v.funcInfo = &FuncInfo{
			Name: fd.Name.Name,
		}

		if fd.Recv != nil {
			recv := fd.Recv.List[0]
			v.funcInfo.Receiver = &ReceiverInfo{
				Name:     recv.Names[0].Name,
				TypeName: recv.Type.(*ast.StarExpr).X.(*ast.Ident).Name,
			}
		}

		if fd.Type.Results != nil {
			for _, f := range fd.Type.Results.List {
				if t, ok := v.info.Types[f.Type]; ok {
					ri := &ResultInfo{
						TypeName:  t.Type.String(),
						ZeroValue: getZeroValueStringRepresentationFromType(t.Type),
					}

					if t.Value != nil {
						ri.Name = t.Value.String()
					}

					v.funcInfo.Results = append(v.funcInfo.Results, ri)
				}
			}
		}

		v.funcInfo.StartPos = startPos
		v.funcInfo.EndPos = endPos

		return nil
	}

	return v
}

func (v *visitor) GetFuncInfo() *FuncInfo {
	return v.funcInfo
}

func (v *visitor) getPositionOffset(pos token.Pos) int {
	return v.fileset.Position(pos).Offset
}

func getZeroValueStringRepresentationFromType(typ types.Type) string {
	if _, ok := typ.(*types.Interface); ok {
		return "nil"
	}

	u := typ.Underlying()
	if _, ok := u.(*types.Interface); ok {
		return "nil"
	}

	t := typ.String()

	typeStr, ok := builtinTypeZeroValues[t]
	if ok {
		return typeStr
	}

	if n, ok := typ.(*types.Named); ok {
		return fmt.Sprintf("%s{}", n.Obj().Name())
	}

	// NOTE: This line should never be reached
	return "nil"
}
