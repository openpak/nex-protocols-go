package types

import (
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go/v2/types"
)

// PlayerStatSortCriteria is a type within the Player Stats protocol
type PlayerStatSortCriteria struct {
	types.Structure
	SortStatID types.UInt32
	SortOrder  types.UInt32
}

// WriteTo writes the PlayerStatSortCriteria to the given writable
func (pssc PlayerStatSortCriteria) WriteTo(writable types.Writable) {
	contentWritable := writable.CopyNew()

	pssc.SortStatID.WriteTo(contentWritable)
	pssc.SortOrder.WriteTo(contentWritable)

	content := contentWritable.Bytes()

	pssc.WriteHeaderTo(writable, uint32(len(content)))

	writable.Write(content)
}

// ExtractFrom extracts the PlayerStatSortCriteria from the given readable
func (pssc *PlayerStatSortCriteria) ExtractFrom(readable types.Readable) error {
	if err := pssc.ExtractHeaderFrom(readable); err != nil {
		return fmt.Errorf("failed to extract PlayerStatSortCriteria header. %s", err.Error())
	}

	if err := pssc.SortStatID.ExtractFrom(readable); err != nil {
		return fmt.Errorf("failed to extract PlayerStatSortCriteria.SortStatID. %s", err.Error())
	}

	if err := pssc.SortOrder.ExtractFrom(readable); err != nil {
		return fmt.Errorf("failed to extract PlayerStatSortCriteria.SortOrder. %s", err.Error())
	}

	return nil
}

// Copy returns a new copied instance of PlayerStatSortCriteria
func (pssc PlayerStatSortCriteria) Copy() types.RVType {
	copied := NewPlayerStatSortCriteria()

	copied.StructureVersion = pssc.StructureVersion
	copied.SortStatID = pssc.SortStatID.Copy().(types.UInt32)
	copied.SortOrder = pssc.SortOrder.Copy().(types.UInt32)

	return copied
}

// Equals checks if the given PlayerStatSortCriteria contains the same data as the current PlayerStatSortCriteria
func (pssc PlayerStatSortCriteria) Equals(o types.RVType) bool {
	other, ok := o.(PlayerStatSortCriteria)
	if !ok {
		return false
	}

	if pssc.StructureVersion != other.StructureVersion {
		return false
	}

	if !(pssc.SortStatID.Equals(other.SortStatID)) {
		return false
	}

	if !(pssc.SortOrder.Equals(other.SortOrder)) {
		return false
	}

	return true
}

// CopyRef copies the current value of the PlayerStatSortCriteria
// and returns a pointer to the new copy
func (pssc PlayerStatSortCriteria) CopyRef() types.RVTypePtr {
	copied := pssc.Copy().(PlayerStatSortCriteria)
	return &copied
}

// Deref takes a pointer to the PlayerStatSortCriteria
// and dereferences it to the raw value.
// Only useful when working with an instance of RVTypePtr
func (pssc *PlayerStatSortCriteria) Deref() types.RVType {
	return *pssc
}

// String returns the string representation of the PlayerStatSortCriteria
func (pssc PlayerStatSortCriteria) String() string {
	return pssc.FormatToString(0)
}

// FormatToString pretty-prints the PlayerStatSortCriteria using the provided indentation level
func (pssc PlayerStatSortCriteria) FormatToString(indentationLevel int) string {
	indentationValues := strings.Repeat("\t", indentationLevel+1)
	indentationEnd := strings.Repeat("\t", indentationLevel)

	var b strings.Builder

	b.WriteString("PlayerStatSortCriteria{\n")
	fmt.Fprintf(&b, "%sSortStatID: %s,\n", indentationValues, pssc.SortStatID)
	fmt.Fprintf(&b, "%sSortOrder: %s,\n", indentationValues, pssc.SortOrder)
	fmt.Fprintf(&b, "%s}", indentationEnd)

	return b.String()
}

// NewPlayerStatSortCriteria returns a new PlayerStatSortCriteria
func NewPlayerStatSortCriteria() PlayerStatSortCriteria {
	return PlayerStatSortCriteria{
		SortStatID: types.NewUInt32(0),
		SortOrder:  types.NewUInt32(0),
	}
}
