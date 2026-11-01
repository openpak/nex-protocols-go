package types

import (
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go/v2/types"
)

// GameSessionInvitationSent is a type within the Game Session protocol
type GameSessionInvitationSent struct {
	types.Structure
	SessionKey   GameSessionKey
	RecipientPID types.UInt32
	Message      types.String
	CreationTime types.DateTime
}

// WriteTo writes the GameSessionInvitationSent to the given writable
func (gsis GameSessionInvitationSent) WriteTo(writable types.Writable) {
	contentWritable := writable.CopyNew()

	gsis.SessionKey.WriteTo(contentWritable)
	gsis.RecipientPID.WriteTo(contentWritable)
	gsis.Message.WriteTo(contentWritable)
	gsis.CreationTime.WriteTo(contentWritable)

	content := contentWritable.Bytes()

	gsis.WriteHeaderTo(writable, uint32(len(content)))

	writable.Write(content)
}

// ExtractFrom extracts the GameSessionInvitationSent from the given readable
func (gsis *GameSessionInvitationSent) ExtractFrom(readable types.Readable) error {
	if err := gsis.ExtractHeaderFrom(readable); err != nil {
		return fmt.Errorf("failed to extract GameSessionInvitationSent header. %s", err.Error())
	}

	if err := gsis.SessionKey.ExtractFrom(readable); err != nil {
		return fmt.Errorf("failed to extract GameSessionInvitationSent.SessionKey. %s", err.Error())
	}

	if err := gsis.RecipientPID.ExtractFrom(readable); err != nil {
		return fmt.Errorf("failed to extract GameSessionInvitationSent.RecipientPID. %s", err.Error())
	}

	if err := gsis.Message.ExtractFrom(readable); err != nil {
		return fmt.Errorf("failed to extract GameSessionInvitationSent.Message. %s", err.Error())
	}

	if err := gsis.CreationTime.ExtractFrom(readable); err != nil {
		return fmt.Errorf("failed to extract GameSessionInvitationSent.CreationTime. %s", err.Error())
	}

	return nil
}

// Copy returns a new copied instance of GameSessionInvitationSent
func (gsis GameSessionInvitationSent) Copy() types.RVType {
	copied := NewGameSessionInvitationSent()

	copied.StructureVersion = gsis.StructureVersion
	copied.SessionKey = gsis.SessionKey.Copy().(GameSessionKey)
	copied.RecipientPID = gsis.RecipientPID.Copy().(types.UInt32)
	copied.Message = gsis.Message.Copy().(types.String)
	copied.CreationTime = gsis.CreationTime.Copy().(types.DateTime)

	return copied
}

// Equals checks if the given GameSessionInvitationSent contains the same data as the current GameSessionInvitationSent
func (gsis GameSessionInvitationSent) Equals(o types.RVType) bool {
	other, ok := o.(GameSessionInvitationSent)
	if !ok {
		return false
	}

	if gsis.StructureVersion != other.StructureVersion {
		return false
	}

	if !(gsis.SessionKey.Equals(other.SessionKey)) {
		return false
	}

	if !(gsis.RecipientPID.Equals(other.RecipientPID)) {
		return false
	}

	if !(gsis.Message.Equals(other.Message)) {
		return false
	}

	if !(gsis.CreationTime.Equals(other.CreationTime)) {
		return false
	}

	return true
}

// CopyRef copies the current value of the GameSessionInvitationSent
// and returns a pointer to the new copy
func (gsis GameSessionInvitationSent) CopyRef() types.RVTypePtr {
	copied := gsis.Copy().(GameSessionInvitationSent)
	return &copied
}

// Deref takes a pointer to the GameSessionInvitationSent
// and dereferences it to the raw value.
// Only useful when working with an instance of RVTypePtr
func (gsis *GameSessionInvitationSent) Deref() types.RVType {
	return *gsis
}

// String returns the string representation of the GameSessionInvitationSent
func (gsis GameSessionInvitationSent) String() string {
	return gsis.FormatToString(0)
}

// FormatToString pretty-prints the GameSessionInvitationSent using the provided indentation level
func (gsis GameSessionInvitationSent) FormatToString(indentationLevel int) string {
	indentationValues := strings.Repeat("\t", indentationLevel+1)
	indentationEnd := strings.Repeat("\t", indentationLevel)

	var b strings.Builder

	b.WriteString("GameSessionInvitationSent{\n")
	fmt.Fprintf(&b, "%sSessionKey: %s,\n", indentationValues, gsis.SessionKey.FormatToString(indentationLevel+1))
	fmt.Fprintf(&b, "%sRecipientPID: %s,\n", indentationValues, gsis.RecipientPID)
	fmt.Fprintf(&b, "%sMessage: %s,\n", indentationValues, gsis.Message)
	fmt.Fprintf(&b, "%sCreationTime: %s,\n", indentationValues, gsis.CreationTime.FormatToString(indentationLevel+1))
	fmt.Fprintf(&b, "%s}", indentationEnd)

	return b.String()
}

// NewGameSessionInvitationSent returns a new GameSessionInvitationSent
func NewGameSessionInvitationSent() GameSessionInvitationSent {
	return GameSessionInvitationSent{
		SessionKey:   NewGameSessionKey(),
		RecipientPID: types.NewUInt32(0),
		Message:      types.NewString(""),
		CreationTime: types.NewDateTime(0),
	}
}
