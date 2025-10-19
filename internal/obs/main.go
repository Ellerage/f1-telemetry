package obs

import (
	"fmt"
	"log/slog"
	"os"
	"path/filepath"
	"sync"
	"time"

	"github.com/gorilla/websocket"
)

type OBSService struct {
	conn          *websocket.Conn
	host          string
	port          int
	password      string
	mu            sync.Mutex
	msgID         int
	responses     map[string]chan map[string]any
	respMu        sync.Mutex
	bufferSeconds int
	enabled       bool
}

func NewOBSService(host string, port int, password string, bufferSeconds int, enabled bool) *OBSService {
	return &OBSService{
		host:          host,
		port:          port,
		password:      password,
		responses:     make(map[string]chan map[string]any),
		bufferSeconds: bufferSeconds,
		enabled:       enabled,
	}
}

func (obs *OBSService) GetReplayBufferStatus() (bool, error) {
	if !obs.enabled {
		return false, fmt.Errorf("obs is disabled")
	}

	data, err := obs.sendRequest("GetReplayBufferStatus", nil)
	if err != nil {
		return false, err
	}

	if outputActive, ok := data["outputActive"].(bool); ok {
		return outputActive, nil
	}

	return false, fmt.Errorf("unable to get status")
}

func (obs *OBSService) StartReplayBuffer() error {
	if !obs.enabled {
		return fmt.Errorf("obs is disabled")
	}

	slog.Info("Starting Replay Buffer...")
	_, err := obs.sendRequest("StartReplayBuffer", nil)
	if err != nil {
		return err
	}

	for i := 0; i < obs.bufferSeconds; i++ {
		time.Sleep(1 * time.Second)
		active, err := obs.GetReplayBufferStatus()
		if err == nil && active {
			slog.Info("Replay Buffer is now active")
			return nil
		}
	}

	return fmt.Errorf("replay buffer failed to start")
}

func (obs *OBSService) StopBuffer() error {
	if !obs.enabled {
		return fmt.Errorf("obs is disabled")
	}

	slog.Info("Stop Replay Buffer...")
	_, err := obs.sendRequest("StopReplayBuffer", nil)
	return err
}

func (obs *OBSService) SaveReplay(filename string) error {
	if !obs.enabled {
		return fmt.Errorf("obs is disabled")
	}

	active, err := obs.GetReplayBufferStatus()
	if err != nil {
		return fmt.Errorf("failed to check status: %w", err)
	}

	if !active {
		return fmt.Errorf("replay buffer is not active, start it first")
	}

	slog.Info("Save Replay...")

	exePath, err := os.Executable()
	if err != nil {
		return fmt.Errorf("cannot get executable path: %w", err)
	}

	exeDir := filepath.Dir(exePath)

	savePath := filepath.Join(exeDir, filename+".mp4")

	_, err = obs.sendRequest("SaveReplayBuffer", map[string]any{
		"savedReplayPath": savePath,
	})
	return err
}

func (obs *OBSService) Close() error {
	if obs.conn != nil {
		return obs.conn.Close()
	}
	return nil
}
