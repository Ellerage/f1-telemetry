package hotkeyhandler

import (
	hook "github.com/robotn/gohook"
)

type SessionStorage interface {
	NewToken()
	RevokeToken()
}

type HotkeyHandlerParams struct {
	SessionStorage SessionStorage
}

type HotkeyHandler struct {
	sessionStorage SessionStorage
}

func NewHotkeyHandler(params HotkeyHandlerParams) *HotkeyHandler {
	return &HotkeyHandler{sessionStorage: params.SessionStorage}
}

func (h *HotkeyHandler) RegisterHotkeyListener() {
	// Create new token
	hook.Register(hook.KeyDown, []string{"space"}, func(e hook.Event) {
		h.sessionStorage.NewToken()
	})

	// Revoke token
	hook.Register(hook.KeyDown, []string{"delete"}, func(e hook.Event) {
		h.sessionStorage.RevokeToken()
	})

	hook.Start(3000)
}
