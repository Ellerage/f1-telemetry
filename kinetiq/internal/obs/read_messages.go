package obs

import "log"

func (obs *OBSService) readMessages() {
	for {
		var message map[string]any
		if err := obs.conn.ReadJSON(&message); err != nil {
			log.Printf("error reading: %v", err)
			return
		}

		op, ok := message["op"].(float64)
		if !ok {
			continue
		}

		// op == 7 (RequestResponse)
		if op == 7 {
			if d, ok := message["d"].(map[string]any); ok {
				if requestID, ok := d["requestId"].(string); ok {
					obs.respMu.Lock()
					if ch, exists := obs.responses[requestID]; exists {
						ch <- d
						delete(obs.responses, requestID)
					}
					obs.respMu.Unlock()
				}
			}
		}
	}
}
