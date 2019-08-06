package irsdk

import (
	"errors"
	"fmt"
	"log"
	"sync"
	"time"
)

type varBuffer struct {
	tickCount int // used to detect changes in data
	bufOffset int // offset from header
}

type Variable struct {
	varType     varType // irsdk_VarType
	offset      int     // offset fron start of buffer row
	count       int     // number of entrys (array) so length in bytes would be irsdk_VarTypeBytes[type] * count
	countAsTime bool
	Name        string
	Desc        string
	Unit        string
	Value       interface{}
	rawBytes    []byte
}

type varType int

const (
	irVarTypeChar     varType = 0
	irVarTypeBool             = 1
	irVarTypeInt              = 2
	irVarTypeBitField         = 3
	irVarTypeFloat            = 4
	irVarTypeDouble           = 5
)

func (v Variable) String() string {
	var ret string
	switch v.varType {
	case irVarTypeChar:
		ret = fmt.Sprintf("%c", v.Value)
	case irVarTypeBool:
		ret = fmt.Sprintf("%v", v.Value)
	case irVarTypeInt:
		ret = fmt.Sprintf("%d", v.Value)
	case irVarTypeBitField:
		ret = fmt.Sprintf("%s", v.Value)
	case irVarTypeFloat:
		ret = fmt.Sprintf("%f", v.Value)
	case irVarTypeDouble:
		ret = fmt.Sprintf("%f", v.Value)
	default:
		ret = fmt.Sprintf("Unknown (%d)", v.varType)
	}
	return ret
}

func (v Variable) getSize() int {
	switch v.varType {
	case irVarTypeChar, irVarTypeBool:
		return 1
	case irVarTypeInt, irVarTypeBitField, irVarTypeFloat:
		return 4
	case irVarTypeDouble:
		return 8
	}
	log.Fatalf("Attempted to get size on unknown variable type %d", v.varType)
	return -1
}

func (v Variable) getVal(bytes []byte) (interface{}, error) {
	switch v.varType {
	case irVarTypeChar:
		return string(bytes[0]), nil
	case irVarTypeBool:
		return bytes[0] > 0, nil
	case irVarTypeInt:
		return byte4ToInt32(bytes), nil
	case irVarTypeBitField:
		return byte4toBitField(bytes), nil
	case irVarTypeFloat:
		return byte4ToFloat(bytes), nil
	case irVarTypeDouble:
		return byte8ToFloat(bytes), nil
	}
	return nil, errors.New(fmt.Sprintf("Unable to convert type %d to a value", v.varType))
}

// TelemetryVars holds all variables we can read from telemetry live
type TelemetryVars struct {
	lastVersion int
	vars        map[string]Variable
	mux         sync.Mutex
}

func findLatestBuffer(r reader, h *header) varBuffer {
	var vb varBuffer
	foundTickCount := 0
	for i := 0; i < h.numBuf; i++ {
		rbuf := make([]byte, 16)
		_, err := r.ReadAt(rbuf, int64(48+i*16))
		if err != nil {
			log.Fatal(err)
		}
		currentVb := varBuffer{
			byte4ToInt(rbuf[0:4]),
			byte4ToInt(rbuf[4:8]),
		}
		//fmt.Printf("BUFF?: %+v\n", currentVb)
		if foundTickCount < currentVb.tickCount {
			foundTickCount = currentVb.tickCount
			vb = currentVb
		}
	}
	//fmt.Printf("BUFF: %+v\n", vb)
	return vb
}

func readVariableHeaders(r reader, h *header) *TelemetryVars {
	vars := TelemetryVars{vars: make(map[string]Variable, h.numVars)}
	for i := 0; i < h.numVars; i++ {
		rbuf := make([]byte, 144)
		_, err := r.ReadAt(rbuf, int64(h.headerOffset+i*144))
		if err != nil {
			log.Fatal(err)
		}
		v := Variable{
			byte4ToVarType(rbuf[0:4]),
			byte4ToInt(rbuf[4:8]),
			byte4ToInt(rbuf[8:12]),
			int(rbuf[12]) > 0,
			bytesToString(rbuf[16:48]),
			bytesToString(rbuf[48:112]),
			bytesToString(rbuf[112:144]),
			nil,
			nil,
		}
		vars.vars[v.Name] = v
	}
	return &vars
}

func readVariableValues(sdk *IRSDK) bool {
	newData := false
	if sessionStatusOK(sdk.h.status) {
		// find latest buffer for variables
		vb := findLatestBuffer(sdk.r, sdk.h)
		sdk.tVars.mux.Lock()
		if sdk.tVars.lastVersion < vb.tickCount {
			newData = true
			sdk.tVars.lastVersion = vb.tickCount
			sdk.lastValidData = time.Now().Unix()
			for varName, v := range sdk.tVars.vars {
				var rbuf []byte
				var bufIdx int
				rbuf = make([]byte, v.count*v.getSize())
				_, err := sdk.r.ReadAt(rbuf, int64(vb.bufOffset+v.offset))
				if err != nil {
					log.Fatal(err)
				}
				if v.count > 1 {
					vals := make([]interface{}, v.count)
					for bufIdx = 0; bufIdx < v.count; bufIdx++ {
						startOff := bufIdx * v.getSize()
						endOff := startOff + v.getSize()
						val, _ := v.getVal(rbuf[startOff:endOff])
						vals[bufIdx] = val
					}
					v.Value = vals
				} else {
					val, _ := v.getVal(rbuf)
					v.Value = val
				}
				v.rawBytes = rbuf
				sdk.tVars.vars[varName] = v
			}
		}
		sdk.tVars.mux.Unlock()
	}

	return newData
}
