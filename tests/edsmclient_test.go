package tests

import (
	"github.com/resistwithvigor/edsm-go-client"
	"testing"
)

func TestNewEDSMClientUserAgent(t *testing.T) {
	var (
		client = edsm_go_client.NewEDSMClient(false, "edsmclient/0.0.1")
	)
	if client.UserAgent != "edsmclient/0.0.1" {
		t.Errorf("UserAgentString Incorrect. got: %s, wanted: \"%s\"", client.UserAgent, "edsmclient/0.0.1")
	} else {
		t.Logf("Got client Debug: %s", client.UserAgent)
	}
}

func TestNewEDSMClientDebug(t *testing.T) {
	var (
		client = edsm_go_client.NewEDSMClient(false, "edsmclient/0.0.1")
	)
	if client.Debug != false {
		t.Errorf("Debug Incorrect. got: %t, wanted: \"%t\"", client.Debug, false)
	} else {
		t.Logf("Got client Debug: %t", client.Debug)
	}
}

func TestGetEliteServerStatus(t *testing.T) {
	status, err := edsm_go_client.GetEDSMServerStatus()

	if len(status.Message) == 0 || err != nil {
		t.Errorf("Server Status incorrect. got: %s, wanted %s", status.Message, "OK")
	} else {
		t.Logf("Server status: %s", status.Message)
	}
}
