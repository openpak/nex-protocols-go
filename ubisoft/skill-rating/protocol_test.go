package protocol

import "testing"

var _ Interface = (*Protocol)(nil)

func TestProtocolID(t *testing.T) {
	if ProtocolID != 48 {
		t.Fatalf("ProtocolID = %d, want 48", ProtocolID)
	}
}
