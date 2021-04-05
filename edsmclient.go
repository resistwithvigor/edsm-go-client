package edsm_go_client

import (
	"encoding/json"
	"github.com/resistwithvigor/edsm-go-client/logging"
)

const (
	EDSMEndpoint         = "https://www.edsm.net/"
	EDSMServerStatus     = EDSMEndpoint + "api-status-v1/elite-server"
	EDSMCmdrRanks        = EDSMEndpoint + "api-commander-v1/get-ranks"
	EDSMCmdrCredits      = EDSMEndpoint + "api-commander-v1/get-credits"
	EDSMCmdrMats         = EDSMEndpoint + "api-commander-v1/get-materials"
	EDSMLogsGetLogs      = EDSMEndpoint + "api-logs-v1/get-logs"
	EDSMLogsGetPosition  = EDSMEndpoint + "api-logs-v1/get-position"
	EDSMLogsSetComment   = EDSMEndpoint + "api-logs-v1/set-comment"
	EDSMLogsGetComment   = EDSMEndpoint + "api-logs-v1/get-comment"
	EDSMLogsGetCommetns  = EDSMEndpoint + "api-logs-v1/get-comments"
	EDSMSystemBodies     = EDSMEndpoint + "api-system-v1/bodies"
	EDSMSystemEstValue   = EDSMEndpoint + "api-system-v1/estimated-value"
	EDSMSystemStations   = EDSMEndpoint + "api-system-v1/stations"
	EDSMSystemMarket     = EDSMEndpoint + "api-system-v1/stations/market"
	EDSMSystemShipyard   = EDSMEndpoint + "api-system-v1/stations/shipyard"
	EDSMSystemOutfitting = EDSMEndpoint + "api-system-v1/stations/outfitting"
	EDSMSystemFactions   = EDSMEndpoint + "api-system-v1/factions"
	EDSMSystemTraffic    = EDSMEndpoint + "api-system-v1/traffic"
	EDSMSystemDeaths     = EDSMEndpoint + "api-system-v1/deaths"
	EDSMSystemsSingle    = EDSMEndpoint + "api-v1/system"
	EDSMSystemsSystems   = EDSMEndpoint + "api-v1/systems"
	EDSMSystemsSphere    = EDSMEndpoint + "api-v1/sphere-systems"
	EDSMSystemsCube      = EDSMEndpoint + "api-v1/cube-systems"
)

type EDSMClient struct {
	Debug     bool
	UserAgent string
}

func GetEDSMServerStatus() (ServerStatus, error) {
	client := NewEDSMClient(false, "edsmgoclient/0.0.1")
	b := new([]byte)
	response, err := client.request("GET", EDSMServerStatus, *b)
	if err != nil {
		logging.Error.Println("get server status error", err)
		return ServerStatus{}, err
	}
	var m ServerStatus
	errM := json.Unmarshal(response, &m)
	if errM != nil {
		logging.Error.Println("get server status error", errM)
	}
	return m, err

}
