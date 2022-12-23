package d2

import (
	"bytes"
	"context"

	"github.com/yuin/goldmark/ast"
	"github.com/yuin/goldmark/renderer"
	"github.com/yuin/goldmark/util"
	"oss.terrastruct.com/d2/d2layouts/d2dagrelayout"
	"oss.terrastruct.com/d2/d2lib"
	"oss.terrastruct.com/d2/d2renderers/d2svg"
	"oss.terrastruct.com/d2/d2themes/d2themescatalog"
	"oss.terrastruct.com/d2/lib/textmeasure"
)

type HTMLRenderer struct {
}

func (r *HTMLRenderer) RegisterFuncs(reg renderer.NodeRendererFuncRegisterer) {
	reg.Register(KindBlock, r.Render)
}

func (r *HTMLRenderer) Render(w util.BufWriter, src []byte, node ast.Node, entering bool) (ast.WalkStatus, error) {
	n := node.(*Block)
	if !entering {
		w.WriteString("</p>")
		return ast.WalkContinue, nil
	}
	w.WriteString(`<p class="d2">`)

	b := bytes.Buffer{}
	lines := n.Lines()
	for i := 0; i < lines.Len(); i++ {
		line := lines.At(i)
		b.Write(line.Value(src))
	}

	if b.Len() == 0 {
		return ast.WalkContinue, nil
	}

	ruler, err := textmeasure.NewRuler()
	if err != nil {
		return ast.WalkStop, err
	}
	diagram, _, err := d2lib.Compile(context.Background(), b.String(), &d2lib.CompileOptions{
		Layout:  d2dagrelayout.Layout,
		Ruler:   ruler,
		ThemeID: d2themescatalog.CoolClassics.ID,
	})
	if err != nil {
		_, err = w.Write(b.Bytes())
		return ast.WalkContinue, err
	}
	out, err := d2svg.Render(diagram, d2svg.DEFAULT_PADDING)
	if err != nil {
		_, err = w.Write(b.Bytes())
		return ast.WalkContinue, err
	}

	_, err = w.Write(out)
	return ast.WalkContinue, err
}
