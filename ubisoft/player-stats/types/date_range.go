package types

import (
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go/v2/types"
)

// DateRange is a type within the Player Stats protocol
type DateRange struct {
	types.Structure
	StartingDatetime types.DateTime
	EndingDatetime   types.DateTime
}

// WriteTo writes the DateRange to the given writable
func (dr DateRange) WriteTo(writable types.Writable) {
	contentWritable := writable.CopyNew()

	dr.StartingDatetime.WriteTo(contentWritable)
	dr.EndingDatetime.WriteTo(contentWritable)

	content := contentWritable.Bytes()

	dr.WriteHeaderTo(writable, uint32(len(content)))

	writable.Write(content)
}

// ExtractFrom extracts the DateRange from the given readable
func (dr *DateRange) ExtractFrom(readable types.Readable) error {
	if err := dr.ExtractHeaderFrom(readable); err != nil {
		return fmt.Errorf("failed to extract DateRange header. %s", err.Error())
	}

	if err := dr.StartingDatetime.ExtractFrom(readable); err != nil {
		return fmt.Errorf("failed to extract DateRange.StartingDatetime. %s", err.Error())
	}

	if err := dr.EndingDatetime.ExtractFrom(readable); err != nil {
		return fmt.Errorf("failed to extract DateRange.EndingDatetime. %s", err.Error())
	}

	return nil
}

// Copy returns a new copied instance of DateRange
func (dr DateRange) Copy() types.RVType {
	copied := NewDateRange()

	copied.StructureVersion = dr.StructureVersion
	copied.StartingDatetime = dr.StartingDatetime.Copy().(types.DateTime)
	copied.EndingDatetime = dr.EndingDatetime.Copy().(types.DateTime)

	return copied
}

// Equals checks if the given DateRange contains the same data as the current DateRange
func (dr DateRange) Equals(o types.RVType) bool {
	other, ok := o.(DateRange)
	if !ok {
		return false
	}

	if dr.StructureVersion != other.StructureVersion {
		return false
	}

	if !(dr.StartingDatetime.Equals(other.StartingDatetime)) {
		return false
	}

	if !(dr.EndingDatetime.Equals(other.EndingDatetime)) {
		return false
	}

	return true
}

// CopyRef copies the current value of the DateRange
// and returns a pointer to the new copy
func (dr DateRange) CopyRef() types.RVTypePtr {
	copied := dr.Copy().(DateRange)
	return &copied
}

// Deref takes a pointer to the DateRange
// and dereferences it to the raw value.
// Only useful when working with an instance of RVTypePtr
func (dr *DateRange) Deref() types.RVType {
	return *dr
}

// String returns the string representation of the DateRange
func (dr DateRange) String() string {
	return dr.FormatToString(0)
}

// FormatToString pretty-prints the DateRange using the provided indentation level
func (dr DateRange) FormatToString(indentationLevel int) string {
	indentationValues := strings.Repeat("\t", indentationLevel+1)
	indentationEnd := strings.Repeat("\t", indentationLevel)

	var b strings.Builder

	b.WriteString("DateRange{\n")
	fmt.Fprintf(&b, "%sStartingDatetime: %s,\n", indentationValues, dr.StartingDatetime.FormatToString(indentationLevel+1))
	fmt.Fprintf(&b, "%sEndingDatetime: %s,\n", indentationValues, dr.EndingDatetime.FormatToString(indentationLevel+1))
	fmt.Fprintf(&b, "%s}", indentationEnd)

	return b.String()
}

// NewDateRange returns a new DateRange
func NewDateRange() DateRange {
	return DateRange{
		StartingDatetime: types.NewDateTime(0),
		EndingDatetime:   types.NewDateTime(0),
	}
}
