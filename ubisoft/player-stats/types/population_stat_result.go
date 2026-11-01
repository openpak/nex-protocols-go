package types

import (
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go/v2/types"
)

// PopulationStatResult is a type within the Player Stats protocol
type PopulationStatResult struct {
	types.Structure
	BoardID           types.UInt32
	ContextID         types.UInt32
	StatID            types.UInt32
	Sum               types.Double
	Average           types.Double
	StandardDeviation types.Double
}

// WriteTo writes the PopulationStatResult to the given writable
func (psr PopulationStatResult) WriteTo(writable types.Writable) {
	contentWritable := writable.CopyNew()

	psr.BoardID.WriteTo(contentWritable)
	psr.ContextID.WriteTo(contentWritable)
	psr.StatID.WriteTo(contentWritable)
	psr.Sum.WriteTo(contentWritable)
	psr.Average.WriteTo(contentWritable)
	psr.StandardDeviation.WriteTo(contentWritable)

	content := contentWritable.Bytes()

	psr.WriteHeaderTo(writable, uint32(len(content)))

	writable.Write(content)
}

// ExtractFrom extracts the PopulationStatResult from the given readable
func (psr *PopulationStatResult) ExtractFrom(readable types.Readable) error {
	if err := psr.ExtractHeaderFrom(readable); err != nil {
		return fmt.Errorf("failed to extract PopulationStatResult header. %s", err.Error())
	}

	if err := psr.BoardID.ExtractFrom(readable); err != nil {
		return fmt.Errorf("failed to extract PopulationStatResult.BoardID. %s", err.Error())
	}

	if err := psr.ContextID.ExtractFrom(readable); err != nil {
		return fmt.Errorf("failed to extract PopulationStatResult.ContextID. %s", err.Error())
	}

	if err := psr.StatID.ExtractFrom(readable); err != nil {
		return fmt.Errorf("failed to extract PopulationStatResult.StatID. %s", err.Error())
	}

	if err := psr.Sum.ExtractFrom(readable); err != nil {
		return fmt.Errorf("failed to extract PopulationStatResult.Sum. %s", err.Error())
	}

	if err := psr.Average.ExtractFrom(readable); err != nil {
		return fmt.Errorf("failed to extract PopulationStatResult.Average. %s", err.Error())
	}

	if err := psr.StandardDeviation.ExtractFrom(readable); err != nil {
		return fmt.Errorf("failed to extract PopulationStatResult.StandardDeviation. %s", err.Error())
	}

	return nil
}

// Copy returns a new copied instance of PopulationStatResult
func (psr PopulationStatResult) Copy() types.RVType {
	copied := NewPopulationStatResult()

	copied.StructureVersion = psr.StructureVersion
	copied.BoardID = psr.BoardID.Copy().(types.UInt32)
	copied.ContextID = psr.ContextID.Copy().(types.UInt32)
	copied.StatID = psr.StatID.Copy().(types.UInt32)
	copied.Sum = psr.Sum.Copy().(types.Double)
	copied.Average = psr.Average.Copy().(types.Double)
	copied.StandardDeviation = psr.StandardDeviation.Copy().(types.Double)

	return copied
}

// Equals checks if the given PopulationStatResult contains the same data as the current PopulationStatResult
func (psr PopulationStatResult) Equals(o types.RVType) bool {
	other, ok := o.(PopulationStatResult)
	if !ok {
		return false
	}

	if psr.StructureVersion != other.StructureVersion {
		return false
	}

	if !(psr.BoardID.Equals(other.BoardID)) {
		return false
	}

	if !(psr.ContextID.Equals(other.ContextID)) {
		return false
	}

	if !(psr.StatID.Equals(other.StatID)) {
		return false
	}

	if !(psr.Sum.Equals(other.Sum)) {
		return false
	}

	if !(psr.Average.Equals(other.Average)) {
		return false
	}

	if !(psr.StandardDeviation.Equals(other.StandardDeviation)) {
		return false
	}

	return true
}

// CopyRef copies the current value of the PopulationStatResult
// and returns a pointer to the new copy
func (psr PopulationStatResult) CopyRef() types.RVTypePtr {
	copied := psr.Copy().(PopulationStatResult)
	return &copied
}

// Deref takes a pointer to the PopulationStatResult
// and dereferences it to the raw value.
// Only useful when working with an instance of RVTypePtr
func (psr *PopulationStatResult) Deref() types.RVType {
	return *psr
}

// String returns the string representation of the PopulationStatResult
func (psr PopulationStatResult) String() string {
	return psr.FormatToString(0)
}

// FormatToString pretty-prints the PopulationStatResult using the provided indentation level
func (psr PopulationStatResult) FormatToString(indentationLevel int) string {
	indentationValues := strings.Repeat("\t", indentationLevel+1)
	indentationEnd := strings.Repeat("\t", indentationLevel)

	var b strings.Builder

	b.WriteString("PopulationStatResult{\n")
	fmt.Fprintf(&b, "%sBoardID: %s,\n", indentationValues, psr.BoardID)
	fmt.Fprintf(&b, "%sContextID: %s,\n", indentationValues, psr.ContextID)
	fmt.Fprintf(&b, "%sStatID: %s,\n", indentationValues, psr.StatID)
	fmt.Fprintf(&b, "%sSum: %s,\n", indentationValues, psr.Sum)
	fmt.Fprintf(&b, "%sAverage: %s,\n", indentationValues, psr.Average)
	fmt.Fprintf(&b, "%sStandardDeviation: %s,\n", indentationValues, psr.StandardDeviation)
	fmt.Fprintf(&b, "%s}", indentationEnd)

	return b.String()
}

// NewPopulationStatResult returns a new PopulationStatResult
func NewPopulationStatResult() PopulationStatResult {
	return PopulationStatResult{
		BoardID:           types.NewUInt32(0),
		ContextID:         types.NewUInt32(0),
		StatID:            types.NewUInt32(0),
		Sum:               types.NewDouble(0),
		Average:           types.NewDouble(0),
		StandardDeviation: types.NewDouble(0),
	}
}
