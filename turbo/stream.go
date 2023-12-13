package turbo

import (
	"net/http"

	"github.com/a-h/templ"
	"github.com/bilus/microwire/turbo/templates"
)

type ActionType string

const (
	ActionAppend  ActionType = "append"
	ActionPrepend ActionType = "prepend"
	ActionReplace ActionType = "replace"
	ActionUpdate  ActionType = "update"
	ActionRemove  ActionType = "remove"
)

type Action struct {
	Type     ActionType
	Target   string
	Template templ.Component
}

func newAction(at ActionType, target string, template templ.Component) Action {
	return Action{Type: at, Target: target, Template: template}
}

func Append(target string, template templ.Component) Action {
	return newAction(ActionAppend, target, template)
}

func Prepend(target string, template templ.Component) Action {
	return newAction(ActionPrepend, target, template)
}

func Replace(target string, template templ.Component) Action {
	return newAction(ActionReplace, target, template)
}

func Update(target string, template templ.Component) Action {
	return newAction(ActionUpdate, target, template)
}

func Remove(target string, template templ.Component) Action {
	return newAction(ActionRemove, target, template)
}

func Stream(actions ...Action) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/vnd.turbo-stream.html")
		for _, action := range actions {
			action := action
			_ = templates.Action(string(action.Type), action.Target, action.Template).Render(r.Context(), w)
		}
	})
}
