package types

import (
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go/v2/types"
)

// StatboardQuery is a type within the Player Stats protocol
type StatboardQuery struct {
	types.Structure
	BoardID        types.UInt32
	ContextIDs     types.List[types.UInt32]
	ResetFrequency types.UInt32
	StatIDs        types.List[types.UInt32]
	SortCriterias  types.List[PlayerStatSortCriteria]
}

// WriteTo writes the StatboardQuery to the given writable
func (sq StatboardQuery) WriteTo(writable types.Writable) {
	contentWritable := writable.CopyNew()

	sq.BoardID.WriteTo(contentWritable)
	sq.ContextIDs.WriteTo(contentWritable)
	sq.ResetFrequency.WriteTo(contentWritable)
	sq.StatIDs.WriteTo(contentWritable)
	sq.SortCriterias.WriteTo(contentWritable)

	content := contentWritable.Bytes()

	sq.WriteHeaderTo(writable, uint32(len(content)))

	writable.Write(content)
}

// ExtractFrom extracts the StatboardQuery from the given readable
func (sq *StatboardQuery) ExtractFrom(readable types.Readable) error {
	if err := sq.ExtractHeaderFrom(readable); err != nil {
		return fmt.Errorf("failed to extract StatboardQuery header. %s", err.Error())
	}

	if err := sq.BoardID.ExtractFrom(readable); err != nil {
		return fmt.Errorf("failed to extract StatboardQuery.BoardID. %s", err.Error())
	}

	if err := sq.ContextIDs.ExtractFrom(readable); err != nil {
		return fmt.Errorf("failed to extract StatboardQuery.ContextIDs. %s", err.Error())
	}

	if err := sq.ResetFrequency.ExtractFrom(readable); err != nil {
		return fmt.Errorf("failed to extract StatboardQuery.ResetFrequency. %s", err.Error())
	}

	if err := sq.StatIDs.ExtractFrom(readable); err != nil {
		return fmt.Errorf("failed to extract StatboardQuery.StatIDs. %s", err.Error())
	}

	if err := sq.SortCriterias.ExtractFrom(readable); err != nil {
		return fmt.Errorf("failed to extract StatboardQuery.SortCriterias. %s", err.Error())
	}

	return nil
}

// Copy returns a new copied instance of StatboardQuery
func (sq StatboardQuery) Copy() types.RVType {
	copied := NewStatboardQuery()

	copied.StructureVersion = sq.StructureVersion
	copied.BoardID = sq.BoardID.Copy().(types.UInt32)
	copied.ContextIDs = sq.ContextIDs.Copy().(types.List[types.UInt32])
	copied.ResetFrequency = sq.ResetFrequency.Copy().(types.UInt32)
	copied.StatIDs = sq.StatIDs.Copy().(types.List[types.UInt32])
	copied.SortCriterias = sq.SortCriterias.Copy().(types.List[PlayerStatSortCriteria])

	return copied
}

// Equals checks if the given StatboardQuery contains the same data as the current StatboardQuery
func (sq StatboardQuery) Equals(o types.RVType) bool {
	other, ok := o.(StatboardQuery)
	if !ok {
		return false
	}

	if sq.StructureVersion != other.StructureVersion {
		return false
	}

	if !(sq.BoardID.Equals(other.BoardID)) {
		return false
	}

	if !(sq.ContextIDs.Equals(other.ContextIDs)) {
		return false
	}

	if !(sq.ResetFrequency.Equals(other.ResetFrequency)) {
		return false
	}

	if !(sq.StatIDs.Equals(other.StatIDs)) {
		return false
	}

	if !(sq.SortCriterias.Equals(other.SortCriterias)) {
		return false
	}

	return true
}

// CopyRef copies the current value of the StatboardQuery
// and returns a pointer to the new copy
func (sq StatboardQuery) CopyRef() types.RVTypePtr {
	copied := sq.Copy().(StatboardQuery)
	return &copied
}

// Deref takes a pointer to the StatboardQuery
// and dereferences it to the raw value.
// Only useful when working with an instance of RVTypePtr
func (sq *StatboardQuery) Deref() types.RVType {
	return *sq
}

// String returns the string representation of the StatboardQuery
func (sq StatboardQuery) String() string {
	return sq.FormatToString(0)
}

// FormatToString pretty-prints the StatboardQuery using the provided indentation level
func (sq StatboardQuery) FormatToString(indentationLevel int) string {
	indentationValues := strings.Repeat("\t", indentationLevel+1)
	indentationEnd := strings.Repeat("\t", indentationLevel)

	var b strings.Builder

	b.WriteString("StatboardQuery{\n")
	fmt.Fprintf(&b, "%sBoardID: %s,\n", indentationValues, sq.BoardID)
	fmt.Fprintf(&b, "%sContextIDs: %s,\n", indentationValues, sq.ContextIDs)
	fmt.Fprintf(&b, "%sResetFrequency: %s,\n", indentationValues, sq.ResetFrequency)
	fmt.Fprintf(&b, "%sStatIDs: %s,\n", indentationValues, sq.StatIDs)
	fmt.Fprintf(&b, "%sSortCriterias: %s,\n", indentationValues, sq.SortCriterias)
	fmt.Fprintf(&b, "%s}", indentationEnd)

	return b.String()
}

// NewStatboardQuery returns a new StatboardQuery
func NewStatboardQuery() StatboardQuery {
	return StatboardQuery{
		BoardID:        types.NewUInt32(0),
		ContextIDs:     types.NewList[types.UInt32](),
		ResetFrequency: types.NewUInt32(0),
		StatIDs:        types.NewList[types.UInt32](),
		SortCriterias:  types.NewList[PlayerStatSortCriteria](),
	}
}
