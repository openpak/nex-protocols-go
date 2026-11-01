package protocol

import "testing"

func TestProtocolID(t *testing.T) {
	if ProtocolID != 44 {
		t.Fatalf("ProtocolID = %d, want 44", ProtocolID)
	}
}
