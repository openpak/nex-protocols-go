package types

import (
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go/v2/types"
)

// GameSessionInvitationReceived is a type within the Game Session protocol
type GameSessionInvitationReceived struct {
	types.Structure
	SessionKey   GameSessionKey
	SenderPID    types.UInt32
	Message      types.String
	CreationTime types.DateTime
}

// WriteTo writes the GameSessionInvitationReceived to the given writable
func (gsir GameSessionInvitationReceived) WriteTo(writable types.Writable) {
	contentWritable := writable.CopyNew()

	gsir.SessionKey.WriteTo(contentWritable)
	gsir.SenderPID.WriteTo(contentWritable)
	gsir.Message.WriteTo(contentWritable)
	gsir.CreationTime.WriteTo(contentWritable)

	content := contentWritable.Bytes()

	gsir.WriteHeaderTo(writable, uint32(len(content)))

	writable.Write(content)
}

// ExtractFrom extracts the GameSessionInvitationReceived from the given readable
func (gsir *GameSessionInvitationReceived) ExtractFrom(readable types.Readable) error {
	if err := gsir.ExtractHeaderFrom(readable); err != nil {
		return fmt.Errorf("failed to extract GameSessionInvitationReceived header. %s", err.Error())
	}

	if err := gsir.SessionKey.ExtractFrom(readable); err != nil {
		return fmt.Errorf("failed to extract GameSessionInvitationReceived.SessionKey. %s", err.Error())
	}

	if err := gsir.SenderPID.ExtractFrom(readable); err != nil {
		return fmt.Errorf("failed to extract GameSessionInvitationReceived.SenderPID. %s", err.Error())
	}

	if err := gsir.Message.ExtractFrom(readable); err != nil {
		return fmt.Errorf("failed to extract GameSessionInvitationReceived.Message. %s", err.Error())
	}

	if err := gsir.CreationTime.ExtractFrom(readable); err != nil {
		return fmt.Errorf("failed to extract GameSessionInvitationReceived.CreationTime. %s", err.Error())
	}

	return nil
}

// Copy returns a new copied instance of GameSessionInvitationReceived
func (gsir GameSessionInvitationReceived) Copy() types.RVType {
	copied := NewGameSessionInvitationReceived()

	copied.StructureVersion = gsir.StructureVersion
	copied.SessionKey = gsir.SessionKey.Copy().(GameSessionKey)
	copied.SenderPID = gsir.SenderPID.Copy().(types.UInt32)
	copied.Message = gsir.Message.Copy().(types.String)
	copied.CreationTime = gsir.CreationTime.Copy().(types.DateTime)

	return copied
}

// Equals checks if the given GameSessionInvitationReceived contains the same data as the current GameSessionInvitationReceived
func (gsir GameSessionInvitationReceived) Equals(o types.RVType) bool {
	other, ok := o.(GameSessionInvitationReceived)
	if !ok {
		return false
	}

	if gsir.StructureVersion != other.StructureVersion {
		return false
	}

	if !(gsir.SessionKey.Equals(other.SessionKey)) {
		return false
	}

	if !(gsir.SenderPID.Equals(other.SenderPID)) {
		return false
	}

	if !(gsir.Message.Equals(other.Message)) {
		return false
	}

	if !(gsir.CreationTime.Equals(other.CreationTime)) {
		return false
	}

	return true
}

// CopyRef copies the current value of the GameSessionInvitationReceived
// and returns a pointer to the new copy
func (gsir GameSessionInvitationReceived) CopyRef() types.RVTypePtr {
	copied := gsir.Copy().(GameSessionInvitationReceived)
	return &copied
}

// Deref takes a pointer to the GameSessionInvitationReceived
// and dereferences it to the raw value.
// Only useful when working with an instance of RVTypePtr
func (gsir *GameSessionInvitationReceived) Deref() types.RVType {
	return *gsir
}

// String returns the string representation of the GameSessionInvitationReceived
func (gsir GameSessionInvitationReceived) String() string {
	return gsir.FormatToString(0)
}

// FormatToString pretty-prints the GameSessionInvitationReceived using the provided indentation level
func (gsir GameSessionInvitationReceived) FormatToString(indentationLevel int) string {
	indentationValues := strings.Repeat("\t", indentationLevel+1)
	indentationEnd := strings.Repeat("\t", indentationLevel)

	var b strings.Builder

	b.WriteString("GameSessionInvitationReceived{\n")
	fmt.Fprintf(&b, "%sSessionKey: %s,\n", indentationValues, gsir.SessionKey.FormatToString(indentationLevel+1))
	fmt.Fprintf(&b, "%sSenderPID: %s,\n", indentationValues, gsir.SenderPID)
	fmt.Fprintf(&b, "%sMessage: %s,\n", indentationValues, gsir.Message)
	fmt.Fprintf(&b, "%sCreationTime: %s,\n", indentationValues, gsir.CreationTime.FormatToString(indentationLevel+1))
	fmt.Fprintf(&b, "%s}", indentationEnd)

	return b.String()
}

// NewGameSessionInvitationReceived returns a new GameSessionInvitationReceived
func NewGameSessionInvitationReceived() GameSessionInvitationReceived {
	return GameSessionInvitationReceived{
		SessionKey:   NewGameSessionKey(),
		SenderPID:    types.NewUInt32(0),
		Message:      types.NewString(""),
		CreationTime: types.NewDateTime(0),
	}
}
