package types

import (
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go/v2/types"
)

// PopulationStatQuery is a type within the Player Stats protocol
type PopulationStatQuery struct {
	types.Structure
	BoardID   types.UInt32
	ContextID types.UInt32
	StatIDs   types.List[types.UInt32]
}

// WriteTo writes the PopulationStatQuery to the given writable
func (psq PopulationStatQuery) WriteTo(writable types.Writable) {
	contentWritable := writable.CopyNew()

	psq.BoardID.WriteTo(contentWritable)
	psq.ContextID.WriteTo(contentWritable)
	psq.StatIDs.WriteTo(contentWritable)

	content := contentWritable.Bytes()

	psq.WriteHeaderTo(writable, uint32(len(content)))

	writable.Write(content)
}

// ExtractFrom extracts the PopulationStatQuery from the given readable
func (psq *PopulationStatQuery) ExtractFrom(readable types.Readable) error {
	if err := psq.ExtractHeaderFrom(readable); err != nil {
		return fmt.Errorf("failed to extract PopulationStatQuery header. %s", err.Error())
	}

	if err := psq.BoardID.ExtractFrom(readable); err != nil {
		return fmt.Errorf("failed to extract PopulationStatQuery.BoardID. %s", err.Error())
	}

	if err := psq.ContextID.ExtractFrom(readable); err != nil {
		return fmt.Errorf("failed to extract PopulationStatQuery.ContextID. %s", err.Error())
	}

	if err := psq.StatIDs.ExtractFrom(readable); err != nil {
		return fmt.Errorf("failed to extract PopulationStatQuery.StatIDs. %s", err.Error())
	}

	return nil
}

// Copy returns a new copied instance of PopulationStatQuery
func (psq PopulationStatQuery) Copy() types.RVType {
	copied := NewPopulationStatQuery()

	copied.StructureVersion = psq.StructureVersion
	copied.BoardID = psq.BoardID.Copy().(types.UInt32)
	copied.ContextID = psq.ContextID.Copy().(types.UInt32)
	copied.StatIDs = psq.StatIDs.Copy().(types.List[types.UInt32])

	return copied
}

// Equals checks if the given PopulationStatQuery contains the same data as the current PopulationStatQuery
func (psq PopulationStatQuery) Equals(o types.RVType) bool {
	other, ok := o.(PopulationStatQuery)
	if !ok {
		return false
	}

	if psq.StructureVersion != other.StructureVersion {
		return false
	}

	if !(psq.BoardID.Equals(other.BoardID)) {
		return false
	}

	if !(psq.ContextID.Equals(other.ContextID)) {
		return false
	}

	if !(psq.StatIDs.Equals(other.StatIDs)) {
		return false
	}

	return true
}

// CopyRef copies the current value of the PopulationStatQuery
// and returns a pointer to the new copy
func (psq PopulationStatQuery) CopyRef() types.RVTypePtr {
	copied := psq.Copy().(PopulationStatQuery)
	return &copied
}

// Deref takes a pointer to the PopulationStatQuery
// and dereferences it to the raw value.
// Only useful when working with an instance of RVTypePtr
func (psq *PopulationStatQuery) Deref() types.RVType {
	return *psq
}

// String returns the string representation of the PopulationStatQuery
func (psq PopulationStatQuery) String() string {
	return psq.FormatToString(0)
}

// FormatToString pretty-prints the PopulationStatQuery using the provided indentation level
func (psq PopulationStatQuery) FormatToString(indentationLevel int) string {
	indentationValues := strings.Repeat("\t", indentationLevel+1)
	indentationEnd := strings.Repeat("\t", indentationLevel)

	var b strings.Builder

	b.WriteString("PopulationStatQuery{\n")
	fmt.Fprintf(&b, "%sBoardID: %s,\n", indentationValues, psq.BoardID)
	fmt.Fprintf(&b, "%sContextID: %s,\n", indentationValues, psq.ContextID)
	fmt.Fprintf(&b, "%sStatIDs: %s,\n", indentationValues, psq.StatIDs)
	fmt.Fprintf(&b, "%s}", indentationEnd)

	return b.String()
}

// NewPopulationStatQuery returns a new PopulationStatQuery
func NewPopulationStatQuery() PopulationStatQuery {
	return PopulationStatQuery{
		BoardID:   types.NewUInt32(0),
		ContextID: types.NewUInt32(0),
		StatIDs:   types.NewList[types.UInt32](),
	}
}
