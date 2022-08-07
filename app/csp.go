package app

import (
	"fmt"
	"strings"
)

type CSPHeader struct {
	DefaultSRC     []string
	StyleSRC       []string
	ScriptSRC      []string
	FontSRC        []string
	ImgSRC         []string
	BaseURI        []string
	FormAction     []string
	FrameAncestors []string
	ConnectSRC     []string
}

func (c *CSPHeader) String() string {
	return fmt.Sprintf("default-src %s; style-src %s; script-src %s; font-src %s; img-src %s; base-uri %s; form-action %s; frame-ancestors %s; connect-src %s;",
		strings.Join(c.DefaultSRC, " "), strings.Join(c.StyleSRC, " "), strings.Join(c.ScriptSRC, " "), strings.Join(c.FontSRC, " "), strings.Join(c.ImgSRC, " "), strings.Join(c.BaseURI, " "), strings.Join(c.FormAction, " "), strings.Join(c.FrameAncestors, " "), strings.Join(c.ConnectSRC, " "))
}

func NewCSPHeader() *CSPHeader {
	// 個人用なので、汎用的に使えなくてもいい。
	return &CSPHeader{
		DefaultSRC:     []string{"'none'"},
		StyleSRC:       []string{"'self'"},
		ScriptSRC:      []string{"'strict-dynamic'"},
		FontSRC:        []string{"'self'"},
		ImgSRC:         []string{"'self'"},
		BaseURI:        []string{"'none'"},
		FormAction:     []string{"'none'"},
		FrameAncestors: []string{"'none'"},
		ConnectSRC:     []string{"'self'", "*.google-analytics.com"},
	}
}
