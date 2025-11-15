package obs

type OBSRequest struct {
	RequestType string         `json:"requestType"`
	RequestID   string         `json:"requestId"`
	RequestData map[string]any `json:"requestData,omitempty"`
}

type OBSResponse struct {
	Op          int            `json:"op"`
	RequestID   string         `json:"requestId,omitempty"`
	RequestType string         `json:"requestType,omitempty"`
	Success     bool           `json:"requestStatus.result,omitempty"`
	Code        int            `json:"requestStatus.code,omitempty"`
	Data        map[string]any `json:"responseData,omitempty"`
}

type HelloMessage struct {
	Op int `json:"op"`
	D  struct {
		OBSWebSocketVersion string `json:"obsWebSocketVersion"`
		Authentication      struct {
			Challenge string `json:"challenge"`
			Salt      string `json:"salt"`
		} `json:"authentication"`
	} `json:"d"`
}
