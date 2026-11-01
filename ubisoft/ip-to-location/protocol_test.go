package protocol

import "testing"

var _ Interface = (*Protocol)(nil)

func TestProtocolID(t *testing.T) {
	if ProtocolID != 45 {
		t.Fatalf("ProtocolID = %d, want 45", ProtocolID)
	}
}
