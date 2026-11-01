package protocol

import "testing"

var _ Interface = (*Protocol)(nil)

func TestProtocolID(t *testing.T) {
	if ProtocolID != 26 {
		t.Fatalf("ProtocolID = %d, want 26", ProtocolID)
	}
}
