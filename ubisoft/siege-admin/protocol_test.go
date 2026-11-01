package protocol

import "testing"

var _ Interface = (*Protocol)(nil)

func TestProtocolID(t *testing.T) {
	if ProtocolID != 84 {
		t.Fatalf("ProtocolID = %d, want 84", ProtocolID)
	}
}
