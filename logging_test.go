package appwrap

import (
	"bytes"
	. "gopkg.in/check.v1"
)

func (s *AppengineInterfacesTest) TestLogging(c *C) {
	w := &bytes.Buffer{}

	log := NewWriterLogger(w)
	log.Debugf("msg %d", 0)
	log.Infof("msg %d", 1)
	log.Warningf("msg %d", 2)
	log.Errorf("msg %d", 3)
	log.Criticalf("msg %d", 4)

	c.Check(w.String(), Equals,
		"debug: msg 0\n"+
			"info: msg 1\n"+
			"Warning: msg 2\n"+
			"Error: msg 3\n"+
			"CRITICAL: msg 4\n")
}

func (s *AppengineInterfacesTest) TestLevelLogger(c *C) {
	w := &bytes.Buffer{}

	jay := NewWriterLogger(w)
	silentBob := NewLevelLogger(LogLevelSilence, jay)

	silentBob.Debugf("WASSUP")
	silentBob.Infof("YO")
	silentBob.Warningf("WATCH OUT")
	silentBob.Errorf("BUSTED")
	silentBob.Criticalf("OUTTA HERE")

	c.Check(w.String(), Equals, "")
}
