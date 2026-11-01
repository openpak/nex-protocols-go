package protocol

import "testing"

var _ Interface = (*Protocol)(nil)

func TestProtocolID(t *testing.T) {
	if ProtocolID != 30 {
		t.Fatalf("ProtocolID = %d, want 30", ProtocolID)
	}
}
