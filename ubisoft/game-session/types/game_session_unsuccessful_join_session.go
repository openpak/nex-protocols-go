package types

import (
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go/v2/types"
)

// GameSessionUnsuccessfulJoinSession is a type within the Game Session protocol
type GameSessionUnsuccessfulJoinSession struct {
	types.Structure
	SessionKey    GameSessionKey
	ErrorCategory types.UInt32
	ErrorCode     types.Int32
}

// WriteTo writes the GameSessionUnsuccessfulJoinSession to the given writable
func (gsujs GameSessionUnsuccessfulJoinSession) WriteTo(writable types.Writable) {
	contentWritable := writable.CopyNew()

	gsujs.SessionKey.WriteTo(contentWritable)
	gsujs.ErrorCategory.WriteTo(contentWritable)
	gsujs.ErrorCode.WriteTo(contentWritable)

	content := contentWritable.Bytes()

	gsujs.WriteHeaderTo(writable, uint32(len(content)))

	writable.Write(content)
}

// ExtractFrom extracts the GameSessionUnsuccessfulJoinSession from the given readable
func (gsujs *GameSessionUnsuccessfulJoinSession) ExtractFrom(readable types.Readable) error {
	if err := gsujs.ExtractHeaderFrom(readable); err != nil {
		return fmt.Errorf("failed to extract GameSessionUnsuccessfulJoinSession header. %s", err.Error())
	}

	if err := gsujs.SessionKey.ExtractFrom(readable); err != nil {
		return fmt.Errorf("failed to extract GameSessionUnsuccessfulJoinSession.SessionKey. %s", err.Error())
	}

	if err := gsujs.ErrorCategory.ExtractFrom(readable); err != nil {
		return fmt.Errorf("failed to extract GameSessionUnsuccessfulJoinSession.ErrorCategory. %s", err.Error())
	}

	if err := gsujs.ErrorCode.ExtractFrom(readable); err != nil {
		return fmt.Errorf("failed to extract GameSessionUnsuccessfulJoinSession.ErrorCode. %s", err.Error())
	}

	return nil
}

// Copy returns a new copied instance of GameSessionUnsuccessfulJoinSession
func (gsujs GameSessionUnsuccessfulJoinSession) Copy() types.RVType {
	copied := NewGameSessionUnsuccessfulJoinSession()

	copied.StructureVersion = gsujs.StructureVersion
	copied.SessionKey = gsujs.SessionKey.Copy().(GameSessionKey)
	copied.ErrorCategory = gsujs.ErrorCategory.Copy().(types.UInt32)
	copied.ErrorCode = gsujs.ErrorCode.Copy().(types.Int32)

	return copied
}

// Equals checks if the given GameSessionUnsuccessfulJoinSession contains the same data as the current GameSessionUnsuccessfulJoinSession
func (gsujs GameSessionUnsuccessfulJoinSession) Equals(o types.RVType) bool {
	other, ok := o.(GameSessionUnsuccessfulJoinSession)
	if !ok {
		return false
	}

	if gsujs.StructureVersion != other.StructureVersion {
		return false
	}

	if !(gsujs.SessionKey.Equals(other.SessionKey)) {
		return false
	}

	if !(gsujs.ErrorCategory.Equals(other.ErrorCategory)) {
		return false
	}

	if !(gsujs.ErrorCode.Equals(other.ErrorCode)) {
		return false
	}

	return true
}

// CopyRef copies the current value of the GameSessionUnsuccessfulJoinSession
// and returns a pointer to the new copy
func (gsujs GameSessionUnsuccessfulJoinSession) CopyRef() types.RVTypePtr {
	copied := gsujs.Copy().(GameSessionUnsuccessfulJoinSession)
	return &copied
}

// Deref takes a pointer to the GameSessionUnsuccessfulJoinSession
// and dereferences it to the raw value.
// Only useful when working with an instance of RVTypePtr
func (gsujs *GameSessionUnsuccessfulJoinSession) Deref() types.RVType {
	return *gsujs
}

// String returns the string representation of the GameSessionUnsuccessfulJoinSession
func (gsujs GameSessionUnsuccessfulJoinSession) String() string {
	return gsujs.FormatToString(0)
}

// FormatToString pretty-prints the GameSessionUnsuccessfulJoinSession using the provided indentation level
func (gsujs GameSessionUnsuccessfulJoinSession) FormatToString(indentationLevel int) string {
	indentationValues := strings.Repeat("\t", indentationLevel+1)
	indentationEnd := strings.Repeat("\t", indentationLevel)

	var b strings.Builder

	b.WriteString("GameSessionUnsuccessfulJoinSession{\n")
	fmt.Fprintf(&b, "%sSessionKey: %s,\n", indentationValues, gsujs.SessionKey.FormatToString(indentationLevel+1))
	fmt.Fprintf(&b, "%sErrorCategory: %s,\n", indentationValues, gsujs.ErrorCategory)
	fmt.Fprintf(&b, "%sErrorCode: %s,\n", indentationValues, gsujs.ErrorCode)
	fmt.Fprintf(&b, "%s}", indentationEnd)

	return b.String()
}

// NewGameSessionUnsuccessfulJoinSession returns a new GameSessionUnsuccessfulJoinSession
func NewGameSessionUnsuccessfulJoinSession() GameSessionUnsuccessfulJoinSession {
	return GameSessionUnsuccessfulJoinSession{
		SessionKey:    NewGameSessionKey(),
		ErrorCategory: types.NewUInt32(0),
		ErrorCode:     types.NewInt32(0),
	}
}
