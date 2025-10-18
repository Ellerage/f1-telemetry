package obs

import (
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"log/slog"
	"net/url"

	"github.com/gorilla/websocket"
)

func (obs *OBSService) Connect() error {
	u := url.URL{Scheme: "ws", Host: fmt.Sprintf("%s:%d", obs.host, obs.port), Path: "/"}

	conn, _, err := websocket.DefaultDialer.Dial(u.String(), nil)
	if err != nil {
		return fmt.Errorf("err connection to OBS: %w", err)
	}

	obs.conn = conn

	var hello HelloMessage
	if err := conn.ReadJSON(&hello); err != nil {
		return fmt.Errorf("reding Hello: %w", err)
	}

	if obs.password != "" {
		if err := obs.authenticate(hello.D.Authentication.Challenge, hello.D.Authentication.Salt); err != nil {
			return fmt.Errorf("auth Error: %w", err)
		}
	} else {
		identifyMsg := map[string]any{
			"op": 1,
			"d": map[string]any{
				"rpcVersion": 1,
			},
		}
		if err := conn.WriteJSON(identifyMsg); err != nil {
			return fmt.Errorf("identify: %w", err)
		}
	}

	var identified map[string]any
	if err := conn.ReadJSON(&identified); err != nil {
		return fmt.Errorf("identified: %w", err)
	}

	go obs.readMessages()

	slog.Info("success connect")
	return nil
}

func (obs *OBSService) authenticate(challenge, salt string) error {
	hash1 := sha256.Sum256([]byte(obs.password + salt))
	secret := base64.StdEncoding.EncodeToString(hash1[:])

	hash2 := sha256.Sum256([]byte(secret + challenge))
	auth := base64.StdEncoding.EncodeToString(hash2[:])

	identifyMsg := map[string]any{
		"op": 1,
		"d": map[string]any{
			"rpcVersion":     1,
			"authentication": auth,
		},
	}

	return obs.conn.WriteJSON(identifyMsg)
}
