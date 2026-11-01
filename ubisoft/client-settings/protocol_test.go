package protocol

import "testing"

var _ Interface = (*Protocol)(nil)

func TestProtocolID(t *testing.T) {
	if ProtocolID != 28 {
		t.Fatalf("ProtocolID = %d, want 28", ProtocolID)
	}
}
