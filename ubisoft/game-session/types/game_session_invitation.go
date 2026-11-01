package types

import (
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go/v2/types"
)

// GameSessionInvitation is a type within the Game Session protocol
type GameSessionInvitation struct {
	types.Structure
	SessionKey    GameSessionKey
	RecipientPIDs types.List[types.UInt32]
	Message       types.String
}

// WriteTo writes the GameSessionInvitation to the given writable
func (gsi GameSessionInvitation) WriteTo(writable types.Writable) {
	contentWritable := writable.CopyNew()

	gsi.SessionKey.WriteTo(contentWritable)
	gsi.RecipientPIDs.WriteTo(contentWritable)
	gsi.Message.WriteTo(contentWritable)

	content := contentWritable.Bytes()

	gsi.WriteHeaderTo(writable, uint32(len(content)))

	writable.Write(content)
}

// ExtractFrom extracts the GameSessionInvitation from the given readable
func (gsi *GameSessionInvitation) ExtractFrom(readable types.Readable) error {
	if err := gsi.ExtractHeaderFrom(readable); err != nil {
		return fmt.Errorf("failed to extract GameSessionInvitation header. %s", err.Error())
	}

	if err := gsi.SessionKey.ExtractFrom(readable); err != nil {
		return fmt.Errorf("failed to extract GameSessionInvitation.SessionKey. %s", err.Error())
	}

	if err := gsi.RecipientPIDs.ExtractFrom(readable); err != nil {
		return fmt.Errorf("failed to extract GameSessionInvitation.RecipientPIDs. %s", err.Error())
	}

	if err := gsi.Message.ExtractFrom(readable); err != nil {
		return fmt.Errorf("failed to extract GameSessionInvitation.Message. %s", err.Error())
	}

	return nil
}

// Copy returns a new copied instance of GameSessionInvitation
func (gsi GameSessionInvitation) Copy() types.RVType {
	copied := NewGameSessionInvitation()

	copied.StructureVersion = gsi.StructureVersion
	copied.SessionKey = gsi.SessionKey.Copy().(GameSessionKey)
	copied.RecipientPIDs = gsi.RecipientPIDs.Copy().(types.List[types.UInt32])
	copied.Message = gsi.Message.Copy().(types.String)

	return copied
}

// Equals checks if the given GameSessionInvitation contains the same data as the current GameSessionInvitation
func (gsi GameSessionInvitation) Equals(o types.RVType) bool {
	other, ok := o.(GameSessionInvitation)
	if !ok {
		return false
	}

	if gsi.StructureVersion != other.StructureVersion {
		return false
	}

	if !(gsi.SessionKey.Equals(other.SessionKey)) {
		return false
	}

	if !(gsi.RecipientPIDs.Equals(other.RecipientPIDs)) {
		return false
	}

	if !(gsi.Message.Equals(other.Message)) {
		return false
	}

	return true
}

// CopyRef copies the current value of the GameSessionInvitation
// and returns a pointer to the new copy
func (gsi GameSessionInvitation) CopyRef() types.RVTypePtr {
	copied := gsi.Copy().(GameSessionInvitation)
	return &copied
}

// Deref takes a pointer to the GameSessionInvitation
// and dereferences it to the raw value.
// Only useful when working with an instance of RVTypePtr
func (gsi *GameSessionInvitation) Deref() types.RVType {
	return *gsi
}

// String returns the string representation of the GameSessionInvitation
func (gsi GameSessionInvitation) String() string {
	return gsi.FormatToString(0)
}

// FormatToString pretty-prints the GameSessionInvitation using the provided indentation level
func (gsi GameSessionInvitation) FormatToString(indentationLevel int) string {
	indentationValues := strings.Repeat("\t", indentationLevel+1)
	indentationEnd := strings.Repeat("\t", indentationLevel)

	var b strings.Builder

	b.WriteString("GameSessionInvitation{\n")
	fmt.Fprintf(&b, "%sSessionKey: %s,\n", indentationValues, gsi.SessionKey.FormatToString(indentationLevel+1))
	fmt.Fprintf(&b, "%sRecipientPIDs: %s,\n", indentationValues, gsi.RecipientPIDs)
	fmt.Fprintf(&b, "%sMessage: %s,\n", indentationValues, gsi.Message)
	fmt.Fprintf(&b, "%s}", indentationEnd)

	return b.String()
}

// NewGameSessionInvitation returns a new GameSessionInvitation
func NewGameSessionInvitation() GameSessionInvitation {
	return GameSessionInvitation{
		SessionKey:    NewGameSessionKey(),
		RecipientPIDs: types.NewList[types.UInt32](),
		Message:       types.NewString(""),
	}
}
