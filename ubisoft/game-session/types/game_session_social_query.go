package types

import (
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go/v2/types"
)

// GameSessionSocialQuery is a type within the Game Session protocol
type GameSessionSocialQuery struct {
	types.Structure
	TypeID         types.UInt32
	QueryID        types.UInt32
	Parameters     types.List[Property]
	ParticipantIDs types.List[types.UInt32]
}

// WriteTo writes the GameSessionSocialQuery to the given writable
func (gssq GameSessionSocialQuery) WriteTo(writable types.Writable) {
	contentWritable := writable.CopyNew()

	gssq.TypeID.WriteTo(contentWritable)
	gssq.QueryID.WriteTo(contentWritable)
	gssq.Parameters.WriteTo(contentWritable)
	gssq.ParticipantIDs.WriteTo(contentWritable)

	content := contentWritable.Bytes()

	gssq.WriteHeaderTo(writable, uint32(len(content)))

	writable.Write(content)
}

// ExtractFrom extracts the GameSessionSocialQuery from the given readable
func (gssq *GameSessionSocialQuery) ExtractFrom(readable types.Readable) error {
	if err := gssq.ExtractHeaderFrom(readable); err != nil {
		return fmt.Errorf("failed to extract GameSessionSocialQuery header. %s", err.Error())
	}

	if err := gssq.TypeID.ExtractFrom(readable); err != nil {
		return fmt.Errorf("failed to extract GameSessionSocialQuery.TypeID. %s", err.Error())
	}

	if err := gssq.QueryID.ExtractFrom(readable); err != nil {
		return fmt.Errorf("failed to extract GameSessionSocialQuery.QueryID. %s", err.Error())
	}

	if err := gssq.Parameters.ExtractFrom(readable); err != nil {
		return fmt.Errorf("failed to extract GameSessionSocialQuery.Parameters. %s", err.Error())
	}

	if err := gssq.ParticipantIDs.ExtractFrom(readable); err != nil {
		return fmt.Errorf("failed to extract GameSessionSocialQuery.ParticipantIDs. %s", err.Error())
	}

	return nil
}

// Copy returns a new copied instance of GameSessionSocialQuery
func (gssq GameSessionSocialQuery) Copy() types.RVType {
	copied := NewGameSessionSocialQuery()

	copied.StructureVersion = gssq.StructureVersion
	copied.TypeID = gssq.TypeID.Copy().(types.UInt32)
	copied.QueryID = gssq.QueryID.Copy().(types.UInt32)
	copied.Parameters = gssq.Parameters.Copy().(types.List[Property])
	copied.ParticipantIDs = gssq.ParticipantIDs.Copy().(types.List[types.UInt32])

	return copied
}

// Equals checks if the given GameSessionSocialQuery contains the same data as the current GameSessionSocialQuery
func (gssq GameSessionSocialQuery) Equals(o types.RVType) bool {
	other, ok := o.(GameSessionSocialQuery)
	if !ok {
		return false
	}

	if gssq.StructureVersion != other.StructureVersion {
		return false
	}

	if !(gssq.TypeID.Equals(other.TypeID)) {
		return false
	}

	if !(gssq.QueryID.Equals(other.QueryID)) {
		return false
	}

	if !(gssq.Parameters.Equals(other.Parameters)) {
		return false
	}

	if !(gssq.ParticipantIDs.Equals(other.ParticipantIDs)) {
		return false
	}

	return true
}

// CopyRef copies the current value of the GameSessionSocialQuery
// and returns a pointer to the new copy
func (gssq GameSessionSocialQuery) CopyRef() types.RVTypePtr {
	copied := gssq.Copy().(GameSessionSocialQuery)
	return &copied
}

// Deref takes a pointer to the GameSessionSocialQuery
// and dereferences it to the raw value.
// Only useful when working with an instance of RVTypePtr
func (gssq *GameSessionSocialQuery) Deref() types.RVType {
	return *gssq
}

// String returns the string representation of the GameSessionSocialQuery
func (gssq GameSessionSocialQuery) String() string {
	return gssq.FormatToString(0)
}

// FormatToString pretty-prints the GameSessionSocialQuery using the provided indentation level
func (gssq GameSessionSocialQuery) FormatToString(indentationLevel int) string {
	indentationValues := strings.Repeat("\t", indentationLevel+1)
	indentationEnd := strings.Repeat("\t", indentationLevel)

	var b strings.Builder

	b.WriteString("GameSessionSocialQuery{\n")
	fmt.Fprintf(&b, "%sTypeID: %s,\n", indentationValues, gssq.TypeID)
	fmt.Fprintf(&b, "%sQueryID: %s,\n", indentationValues, gssq.QueryID)
	fmt.Fprintf(&b, "%sParameters: %s,\n", indentationValues, gssq.Parameters)
	fmt.Fprintf(&b, "%sParticipantIDs: %s,\n", indentationValues, gssq.ParticipantIDs)
	fmt.Fprintf(&b, "%s}", indentationEnd)

	return b.String()
}

// NewGameSessionSocialQuery returns a new GameSessionSocialQuery
func NewGameSessionSocialQuery() GameSessionSocialQuery {
	return GameSessionSocialQuery{
		TypeID:         types.NewUInt32(0),
		QueryID:        types.NewUInt32(0),
		Parameters:     types.NewList[Property](),
		ParticipantIDs: types.NewList[types.UInt32](),
	}
}
