// Package types implements the structures used by the Game Session protocol

package types

import (
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go/v2/types"
)

// GameSessionKey is a type within the Game Session protocol
type GameSessionKey struct {
	types.Structure
	TypeID    types.UInt32
	SessionID types.UInt32
}

// WriteTo writes the GameSessionKey to the given writable
func (gsk GameSessionKey) WriteTo(writable types.Writable) {
	contentWritable := writable.CopyNew()

	gsk.TypeID.WriteTo(contentWritable)
	gsk.SessionID.WriteTo(contentWritable)

	content := contentWritable.Bytes()

	gsk.WriteHeaderTo(writable, uint32(len(content)))

	writable.Write(content)
}

// ExtractFrom extracts the GameSessionKey from the given readable
func (gsk *GameSessionKey) ExtractFrom(readable types.Readable) error {
	if err := gsk.ExtractHeaderFrom(readable); err != nil {
		return fmt.Errorf("failed to extract GameSessionKey header. %s", err.Error())
	}

	if err := gsk.TypeID.ExtractFrom(readable); err != nil {
		return fmt.Errorf("failed to extract GameSessionKey.TypeID. %s", err.Error())
	}

	if err := gsk.SessionID.ExtractFrom(readable); err != nil {
		return fmt.Errorf("failed to extract GameSessionKey.SessionID. %s", err.Error())
	}

	return nil
}

// Copy returns a new copied instance of GameSessionKey
func (gsk GameSessionKey) Copy() types.RVType {
	copied := NewGameSessionKey()

	copied.StructureVersion = gsk.StructureVersion
	copied.TypeID = gsk.TypeID.Copy().(types.UInt32)
	copied.SessionID = gsk.SessionID.Copy().(types.UInt32)

	return copied
}

// Equals checks if the given GameSessionKey contains the same data as the current GameSessionKey
func (gsk GameSessionKey) Equals(o types.RVType) bool {
	other, ok := o.(GameSessionKey)
	if !ok {
		return false
	}

	if gsk.StructureVersion != other.StructureVersion {
		return false
	}

	if !(gsk.TypeID.Equals(other.TypeID)) {
		return false
	}

	if !(gsk.SessionID.Equals(other.SessionID)) {
		return false
	}

	return true
}

// CopyRef copies the current value of the GameSessionKey
// and returns a pointer to the new copy
func (gsk GameSessionKey) CopyRef() types.RVTypePtr {
	copied := gsk.Copy().(GameSessionKey)
	return &copied
}

// Deref takes a pointer to the GameSessionKey
// and dereferences it to the raw value.
// Only useful when working with an instance of RVTypePtr
func (gsk *GameSessionKey) Deref() types.RVType {
	return *gsk
}

// String returns the string representation of the GameSessionKey
func (gsk GameSessionKey) String() string {
	return gsk.FormatToString(0)
}

// FormatToString pretty-prints the GameSessionKey using the provided indentation level
func (gsk GameSessionKey) FormatToString(indentationLevel int) string {
	indentationValues := strings.Repeat("\t", indentationLevel+1)
	indentationEnd := strings.Repeat("\t", indentationLevel)

	var b strings.Builder

	b.WriteString("GameSessionKey{\n")
	fmt.Fprintf(&b, "%sTypeID: %s,\n", indentationValues, gsk.TypeID)
	fmt.Fprintf(&b, "%sSessionID: %s,\n", indentationValues, gsk.SessionID)
	fmt.Fprintf(&b, "%s}", indentationEnd)

	return b.String()
}

// NewGameSessionKey returns a new GameSessionKey
func NewGameSessionKey() GameSessionKey {
	return GameSessionKey{
		TypeID:    types.NewUInt32(0),
		SessionID: types.NewUInt32(0),
	}
}
