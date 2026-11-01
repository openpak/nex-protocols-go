package protocol

import "testing"

func TestProtocolID(t *testing.T) {
	if ProtocolID != 12 {
		t.Fatalf("ProtocolID = %d, want 12", ProtocolID)
	}
}
