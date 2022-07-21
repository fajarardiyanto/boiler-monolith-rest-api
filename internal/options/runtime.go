package options

import (
	"encoding/json"
	"path/filepath"
	"runtime"
	"strings"
)

type Caller struct {
	File       string `json:"file"`
	Line       int    `json:"line"`
	Fname      string `json:"fname"`
	FnameShort string `json:"-"`
}

func (c Caller) String() string {
	bs, _ := json.Marshal(c)
	return string(bs)
}

func GetCaller(skip int) (cs Caller) {
	if skip == 0 {
		skip = 1
	}

	pc, file, line, ok := runtime.Caller(skip)
	if ok {
		cs.File = file
		cs.Line = line
		fc := runtime.FuncForPC(pc)
		cs.Fname = fc.Name()
		drname := filepath.Dir(fc.Name())
		spsd := strings.Split(drname, "/")
		if len(spsd) >= 2 {
			spsd = spsd[len(spsd)-2:]
		}
		spsd = append(spsd, filepath.Base(fc.Name()))
		cs.FnameShort = filepath.Join(spsd...)
	}

	return cs
}
