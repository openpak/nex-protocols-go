package protocol

import "testing"

var _ Interface = (*Protocol)(nil)

func TestProtocolID(t *testing.T) {
	if ProtocolID != 17 {
		t.Fatalf("ProtocolID = %d, want 17", ProtocolID)
	}
}
