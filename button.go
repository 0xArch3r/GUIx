package GUIx

import (
	"strings"

	"github.com/a-h/templ"
)

const (
	LinkType       = "link"
	HTMXType       = "htmx"
	FormSubmitType = "submit"
)

type ButtonOpts struct {
	key   string
	value string
}

func LinkOpt() ButtonOpts {
	return ButtonOpts{
		"type",
		LinkType,
	}
}

func HrefOpt(href string) ButtonOpts {
	return ButtonOpts{
		"href",
		href,
	}
}

func WithClass(className string) ButtonOpts {
	return ButtonOpts{
		"class",
		className,
	}
}

type buttonParams struct {
	buttonType string
	id         string
	label      string
	classes    []string
	href       string
}

func BasicButton(id, label string, opts ...ButtonOpts) templ.Component {
	params := buttonParams{
		id:      id,
		label:   label,
		classes: []string{},
	}
	for _, opt := range opts {
		switch opt.key {
		case "type":
			params.buttonType = opt.value
		case "href":
			params.href = opt.value
		case "class":
			params.classes = append(params.classes, opt.value)
		}
	}

	switch params.buttonType {
	case LinkType:
		url := templ.SafeURL(params.href)
		return linkButton(params.id, params.label, strings.Join(params.classes, " "), url)
	default:
		panic("invalid button type")
	}
}
