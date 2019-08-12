package irsdk

import (
	"bytes"
	"github.com/gu3st/yaml"
	"golang.org/x/text/encoding/charmap"
	"log"
)

func readSessionData(sdk *IRSDK) bool {
	dec := charmap.Windows1252.NewDecoder()
	rbuf := make([]byte, sdk.h.sessionInfoLen)
	_, err := sdk.r.ReadAt(rbuf, int64(sdk.h.sessionInfoOffset))
	if err != nil {
		log.Fatal(err)
	}

	rbuf, err = dec.Bytes(rbuf)
	if err != nil {
		log.Fatal(err)
	}

	nulTrimmed := bytes.TrimRight(rbuf, "\x00")
	err = yaml.Unmarshal(nulTrimmed, &sdk.sessionData)
	if err != nil {
		log.Fatal(err)
	}
	return true
}

