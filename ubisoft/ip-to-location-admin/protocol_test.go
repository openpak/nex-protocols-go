package protocol

import "testing"

var _ Interface = (*Protocol)(nil)

func TestProtocolID(t *testing.T) {
	if ProtocolID != 46 {
		t.Fatalf("ProtocolID = %d, want 46", ProtocolID)
	}
}
