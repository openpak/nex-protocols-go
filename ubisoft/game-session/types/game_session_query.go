package types

import (
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go/v2/types"
)

// GameSessionQuery is a type within the Game Session protocol
type GameSessionQuery struct {
	types.Structure
	TypeID     types.UInt32
	QueryID    types.UInt32
	Parameters types.List[Property]
}

// WriteTo writes the GameSessionQuery to the given writable
func (gsq GameSessionQuery) WriteTo(writable types.Writable) {
	contentWritable := writable.CopyNew()

	gsq.TypeID.WriteTo(contentWritable)
	gsq.QueryID.WriteTo(contentWritable)
	gsq.Parameters.WriteTo(contentWritable)

	content := contentWritable.Bytes()

	gsq.WriteHeaderTo(writable, uint32(len(content)))

	writable.Write(content)
}

// ExtractFrom extracts the GameSessionQuery from the given readable
func (gsq *GameSessionQuery) ExtractFrom(readable types.Readable) error {
	if err := gsq.ExtractHeaderFrom(readable); err != nil {
		return fmt.Errorf("failed to extract GameSessionQuery header. %s", err.Error())
	}

	if err := gsq.TypeID.ExtractFrom(readable); err != nil {
		return fmt.Errorf("failed to extract GameSessionQuery.TypeID. %s", err.Error())
	}

	if err := gsq.QueryID.ExtractFrom(readable); err != nil {
		return fmt.Errorf("failed to extract GameSessionQuery.QueryID. %s", err.Error())
	}

	if err := gsq.Parameters.ExtractFrom(readable); err != nil {
		return fmt.Errorf("failed to extract GameSessionQuery.Parameters. %s", err.Error())
	}

	return nil
}

// Copy returns a new copied instance of GameSessionQuery
func (gsq GameSessionQuery) Copy() types.RVType {
	copied := NewGameSessionQuery()

	copied.StructureVersion = gsq.StructureVersion
	copied.TypeID = gsq.TypeID.Copy().(types.UInt32)
	copied.QueryID = gsq.QueryID.Copy().(types.UInt32)
	copied.Parameters = gsq.Parameters.Copy().(types.List[Property])

	return copied
}

// Equals checks if the given GameSessionQuery contains the same data as the current GameSessionQuery
func (gsq GameSessionQuery) Equals(o types.RVType) bool {
	other, ok := o.(GameSessionQuery)
	if !ok {
		return false
	}

	if gsq.StructureVersion != other.StructureVersion {
		return false
	}

	if !(gsq.TypeID.Equals(other.TypeID)) {
		return false
	}

	if !(gsq.QueryID.Equals(other.QueryID)) {
		return false
	}

	if !(gsq.Parameters.Equals(other.Parameters)) {
		return false
	}

	return true
}

// CopyRef copies the current value of the GameSessionQuery
// and returns a pointer to the new copy
func (gsq GameSessionQuery) CopyRef() types.RVTypePtr {
	copied := gsq.Copy().(GameSessionQuery)
	return &copied
}

// Deref takes a pointer to the GameSessionQuery
// and dereferences it to the raw value.
// Only useful when working with an instance of RVTypePtr
func (gsq *GameSessionQuery) Deref() types.RVType {
	return *gsq
}

// String returns the string representation of the GameSessionQuery
func (gsq GameSessionQuery) String() string {
	return gsq.FormatToString(0)
}

// FormatToString pretty-prints the GameSessionQuery using the provided indentation level
func (gsq GameSessionQuery) FormatToString(indentationLevel int) string {
	indentationValues := strings.Repeat("\t", indentationLevel+1)
	indentationEnd := strings.Repeat("\t", indentationLevel)

	var b strings.Builder

	b.WriteString("GameSessionQuery{\n")
	fmt.Fprintf(&b, "%sTypeID: %s,\n", indentationValues, gsq.TypeID)
	fmt.Fprintf(&b, "%sQueryID: %s,\n", indentationValues, gsq.QueryID)
	fmt.Fprintf(&b, "%sParameters: %s,\n", indentationValues, gsq.Parameters)
	fmt.Fprintf(&b, "%s}", indentationEnd)

	return b.String()
}

// NewGameSessionQuery returns a new GameSessionQuery
func NewGameSessionQuery() GameSessionQuery {
	return GameSessionQuery{
		TypeID:     types.NewUInt32(0),
		QueryID:    types.NewUInt32(0),
		Parameters: types.NewList[Property](),
	}
}
