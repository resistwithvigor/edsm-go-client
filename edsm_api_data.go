package edsm_go_client

//type APIStatus struct {
//	LastUpdate *time.Time `json:"lastUpdate,omitempty"`
//	Type       string     `json:"type,omitempty"`
//	Message    string     `json:"message,omitempty"`
//	Status     int32      `json:"status"`
//}

type CmdrPosition struct {
	MSGNum         int64  `json:"msgnum"`
	Message        string `json:"msg"`
	System         string `json:"system"`
	SystemId       int64  `json:"systemId"`
	FirstDiscovery bool   `json:"firstDiscvoer"`
	Date           string `json:"date,omitempty"`
}

type ServerStatus struct {
	LastUpdate *EDSMServerStatusDate `json:"lastUpdate,omitempty"`
	Type       string                `json:"type,omitempty"`
	Message    string                `json:"message,omitempty"`
	status     int                   `json:"status,omitempty"`
}
