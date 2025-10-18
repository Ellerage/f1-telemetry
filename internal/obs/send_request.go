package obs

import (
	"fmt"
	"time"
)

func (obs *OBSService) sendRequest(requestType string, requestData map[string]any) (map[string]any, error) {
	obs.mu.Lock()
	obs.msgID++
	msgID := fmt.Sprintf("req-%d", obs.msgID)
	obs.mu.Unlock()

	respChan := make(chan map[string]any, 1)
	obs.respMu.Lock()
	obs.responses[msgID] = respChan
	obs.respMu.Unlock()

	request := map[string]any{
		"op": 6,
		"d": map[string]any{
			"requestType": requestType,
			"requestId":   msgID,
		},
	}

	if requestData != nil {
		request["d"].(map[string]any)["requestData"] = requestData
	}

	if err := obs.conn.WriteJSON(request); err != nil {
		obs.respMu.Lock()
		delete(obs.responses, msgID)
		obs.respMu.Unlock()
		return nil, fmt.Errorf("err seding request: %w", err)
	}

	select {
	case d := <-respChan:
		var responseData map[string]any
		if data, ok := d["responseData"].(map[string]any); ok {
			responseData = data
		}

		if status, ok := d["requestStatus"].(map[string]any); ok {
			if !status["result"].(bool) {
				code := status["code"].(float64)
				return nil, fmt.Errorf("code: %.0f", code)
			}
		}
		return responseData, nil
	case <-time.After(5 * time.Second):
		obs.respMu.Lock()
		delete(obs.responses, msgID)
		obs.respMu.Unlock()
		return nil, fmt.Errorf("timeout waiting for response")
	}
}
