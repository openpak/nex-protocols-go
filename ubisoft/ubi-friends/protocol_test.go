package protocol

import "testing"

var _ Interface = (*Protocol)(nil)

func TestProtocolID(t *testing.T) {
	if ProtocolID != 47 {
		t.Fatalf("ProtocolID = %d, want 47", ProtocolID)
	}
}
