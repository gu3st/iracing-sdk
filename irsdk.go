package irsdk

import (
	"fmt"
	"github.com/gu3st/iracing-sdk/lib/winevents"
	"github.com/gu3st/yaml"
	"github.com/hidez8891/shm"
	"io/ioutil"
	"log"
	"time"
)

// IRSDK is the main SDK object clients must use
type IRSDK struct {
	r             reader
	h             *header
	tVars         *TelemetryVars
	sessionData *SessionData
	lastValidData int64
}

func (sdk *IRSDK) WaitForData(timeout time.Duration) bool {
	if !sdk.IsConnected() {
		initIRSDK(sdk)
	}
	if winevents.WaitForSingleObject(timeout) {
		readSessionData(sdk)
		return readVariableValues(sdk)
	}
	return false
}

func (sdk *IRSDK) GetVar(name string) (Variable, error) {
	if !sessionStatusOK(sdk.h.status) {
		return Variable{}, fmt.Errorf("Session is not active")
	}
	sdk.tVars.mux.Lock()
	if v, ok := sdk.tVars.vars[name]; ok {
		sdk.tVars.mux.Unlock()
		return v, nil
	}
	sdk.tVars.mux.Unlock()
	return Variable{}, fmt.Errorf("Telemetry Variable %q not found", name)
}

func (sdk *IRSDK) GetLastVersion() int {
	if !sessionStatusOK(sdk.h.status) {
		return -1
	}
	sdk.tVars.mux.Lock()
	last := sdk.tVars.lastVersion
	sdk.tVars.mux.Unlock()
	return last
}

func (sdk *IRSDK) GetSessionData() (*SessionData, error) {
	if !sessionStatusOK(sdk.h.status) {
		return nil, fmt.Errorf("Session not connected")
	}
	return sdk.sessionData, nil
}

func (sdk *IRSDK) IsConnected() bool {
	if sdk.h != nil {
		if sessionStatusOK(sdk.h.status) && (sdk.lastValidData+connTimeout > time.Now().Unix()) {
			return true
		}
	}

	return false
}

// ExportTo exports current memory data to a file
func (sdk *IRSDK) ExportIbtTo(fileName string) {
	rbuf := make([]byte, fileMapSize)
	_, err := sdk.r.ReadAt(rbuf, 0)
	if err != nil {
		log.Fatal(err)
	}
	ioutil.WriteFile(fileName, rbuf, 0644)
}

// ExportTo exports current session yaml data to a file
func (sdk *IRSDK) ExportSessionTo(fileName string) error {
	sd, err := yaml.Marshal(sdk.sessionData)
	if err != nil {
		return fmt.Errorf("Failed to convert session data to YAML string")
	}
	err = ioutil.WriteFile(fileName, sd, 0644)
	if err != nil {
		return fmt.Errorf("Failed to write session data to file")
	}
	return nil
}

func (sdk *IRSDK) BroadcastMsg(msg Msg) {
	if msg.P2 == nil {
		msg.P2 = 0
	}
	winevents.BroadcastMsg(broadcastMsgName, msg.Cmd, msg.P1, msg.P2, msg.P3)
}

// Close clean up sdk resources
func (sdk *IRSDK) Close() {
	sdk.r.Close()
}

// Init creates a SDK instance to operate with
func Init(r reader) IRSDK {
	if r == nil {
		var err error
		r, err = shm.Open(fileMapName, fileMapSize)
		if err != nil {
			log.Fatal(err)
		}
	}

	sdk := IRSDK{r: r, lastValidData: 0}
	winevents.OpenEvent(dataValidEventName)
	initIRSDK(&sdk)
	return sdk
}

func initIRSDK(sdk *IRSDK) {
	h := readHeader(sdk.r)
	sdk.h = &h
	if sdk.tVars != nil {
		sdk.tVars.vars = nil
	}
	if sessionStatusOK(h.status) {
		sdk.tVars = readVariableHeaders(sdk.r, &h)
		sdk.sessionData = &SessionData{}
		readSessionData(sdk)
		readVariableValues(sdk)
	}
}

func sessionStatusOK(status int) bool {
	return (status & stConnected) > 0
}
