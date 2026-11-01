package protocol

import "testing"

var _ Interface = (*Protocol)(nil)

func TestProtocolID(t *testing.T) {
	if ProtocolID != 60 {
		t.Fatalf("ProtocolID = %d, want 60", ProtocolID)
	}
}
