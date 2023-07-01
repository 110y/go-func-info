package analysis

import (
	"go/ast"
	"go/token"
	"go/types"
)

var _ ast.Visitor = (*visitor)(nil)

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
