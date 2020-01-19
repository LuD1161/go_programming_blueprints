package trace

import (
	"bytes"
	"testing"
)

func TestNew(t *testing.T) {
	var buf bytes.Buffer
	tracer := New(&buf)
	if tracer == nil {
		t.Error("Return from New shouldn't be nil")
	} else {
		tracer.Trace("Hello trace package.")
		if buf.String() != "Hello trace package.\n" {
			t.Errorf("Trace shouldn't write %s", buf.String())
		}
	}
}

func TraceOff(t *testing.T) {
	var silentTracer Tracer = Off()
	silentTracer.Trace("something")
}
