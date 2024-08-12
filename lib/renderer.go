package lib

import (
	"context"
	"net/http"

	"taskflow/web/templates"

	"github.com/a-h/templ"
	"github.com/gin-gonic/gin"

	"github.com/gin-gonic/gin/render"
)

var Default = &HTMLTemplRenderer{}

type HTMLTemplRenderer struct {
	FallbackHtmlRenderer render.HTMLRender
}

func (r *HTMLTemplRenderer) Instance(s string, d any) render.Render {
	templData, ok := d.(templ.Component)
	if !ok {
		if r.FallbackHtmlRenderer != nil {
			return r.FallbackHtmlRenderer.Instance(s, d)
		}
	}
	return &Renderer{
		Ctx:       context.Background(),
		Status:    -1,
		Component: templData,
	}
}

func New(ctx context.Context, status int, component templ.Component) *Renderer {
	return &Renderer{
		Ctx:       ctx,
		Status:    status,
		Component: component,
	}
}

type Renderer struct {
	Ctx       context.Context
	Status    int
	Component templ.Component
}

func (t Renderer) Render(w http.ResponseWriter) error {
	t.WriteContentType(w)
	if t.Status != -1 {
		w.WriteHeader(t.Status)
	}
	if t.Component != nil {
		return t.Component.Render(t.Ctx, w)
	}
	return nil
}

func (t Renderer) WriteContentType(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
}

// RenderWithLayout renders the given component with the layout and
// sets the response status code.
//
// Parameters:
//
// - ctx: gin.Context
//
// - t: templ.Component
func RenderWithLayout(ctx *gin.Context, t templ.Component) {
	r := New(ctx.Request.Context(), 200, templates.Layout(t))
	ctx.Render(200, r)
}

// Render renders the given component and sets the response status code.
//
// Parameters:
//
// - ctx: gin.Context
//
// - t: templ.Component
func Render(ctx *gin.Context, t templ.Component) {
	r := New(ctx.Request.Context(), 200, t)
	ctx.Render(200, r)
}
