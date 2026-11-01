package types

import (
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go/v2/types"
)

// StatboardResult is a type within the Player Stats protocol
type StatboardResult struct {
	types.Structure
	BoardID           types.UInt32
	ContextID         types.UInt32
	ResetFrequency    types.UInt32
	PlayerStatSets    types.List[PlayerStatSet]
	DefaultStatValues types.List[PropertyVariant]
}

// WriteTo writes the StatboardResult to the given writable
func (sr StatboardResult) WriteTo(writable types.Writable) {
	contentWritable := writable.CopyNew()

	sr.BoardID.WriteTo(contentWritable)
	sr.ContextID.WriteTo(contentWritable)
	sr.ResetFrequency.WriteTo(contentWritable)
	sr.PlayerStatSets.WriteTo(contentWritable)
	sr.DefaultStatValues.WriteTo(contentWritable)

	content := contentWritable.Bytes()

	sr.WriteHeaderTo(writable, uint32(len(content)))

	writable.Write(content)
}

// ExtractFrom extracts the StatboardResult from the given readable
func (sr *StatboardResult) ExtractFrom(readable types.Readable) error {
	if err := sr.ExtractHeaderFrom(readable); err != nil {
		return fmt.Errorf("failed to extract StatboardResult header. %s", err.Error())
	}

	if err := sr.BoardID.ExtractFrom(readable); err != nil {
		return fmt.Errorf("failed to extract StatboardResult.BoardID. %s", err.Error())
	}

	if err := sr.ContextID.ExtractFrom(readable); err != nil {
		return fmt.Errorf("failed to extract StatboardResult.ContextID. %s", err.Error())
	}

	if err := sr.ResetFrequency.ExtractFrom(readable); err != nil {
		return fmt.Errorf("failed to extract StatboardResult.ResetFrequency. %s", err.Error())
	}

	if err := sr.PlayerStatSets.ExtractFrom(readable); err != nil {
		return fmt.Errorf("failed to extract StatboardResult.PlayerStatSets. %s", err.Error())
	}

	if err := sr.DefaultStatValues.ExtractFrom(readable); err != nil {
		return fmt.Errorf("failed to extract StatboardResult.DefaultStatValues. %s", err.Error())
	}

	return nil
}

// Copy returns a new copied instance of StatboardResult
func (sr StatboardResult) Copy() types.RVType {
	copied := NewStatboardResult()

	copied.StructureVersion = sr.StructureVersion
	copied.BoardID = sr.BoardID.Copy().(types.UInt32)
	copied.ContextID = sr.ContextID.Copy().(types.UInt32)
	copied.ResetFrequency = sr.ResetFrequency.Copy().(types.UInt32)
	copied.PlayerStatSets = sr.PlayerStatSets.Copy().(types.List[PlayerStatSet])
	copied.DefaultStatValues = sr.DefaultStatValues.Copy().(types.List[PropertyVariant])

	return copied
}

// Equals checks if the given StatboardResult contains the same data as the current StatboardResult
func (sr StatboardResult) Equals(o types.RVType) bool {
	other, ok := o.(StatboardResult)
	if !ok {
		return false
	}

	if sr.StructureVersion != other.StructureVersion {
		return false
	}

	if !(sr.BoardID.Equals(other.BoardID)) {
		return false
	}

	if !(sr.ContextID.Equals(other.ContextID)) {
		return false
	}

	if !(sr.ResetFrequency.Equals(other.ResetFrequency)) {
		return false
	}

	if !(sr.PlayerStatSets.Equals(other.PlayerStatSets)) {
		return false
	}

	if !(sr.DefaultStatValues.Equals(other.DefaultStatValues)) {
		return false
	}

	return true
}

// CopyRef copies the current value of the StatboardResult
// and returns a pointer to the new copy
func (sr StatboardResult) CopyRef() types.RVTypePtr {
	copied := sr.Copy().(StatboardResult)
	return &copied
}

// Deref takes a pointer to the StatboardResult
// and dereferences it to the raw value.
// Only useful when working with an instance of RVTypePtr
func (sr *StatboardResult) Deref() types.RVType {
	return *sr
}

// String returns the string representation of the StatboardResult
func (sr StatboardResult) String() string {
	return sr.FormatToString(0)
}

// FormatToString pretty-prints the StatboardResult using the provided indentation level
func (sr StatboardResult) FormatToString(indentationLevel int) string {
	indentationValues := strings.Repeat("\t", indentationLevel+1)
	indentationEnd := strings.Repeat("\t", indentationLevel)

	var b strings.Builder

	b.WriteString("StatboardResult{\n")
	fmt.Fprintf(&b, "%sBoardID: %s,\n", indentationValues, sr.BoardID)
	fmt.Fprintf(&b, "%sContextID: %s,\n", indentationValues, sr.ContextID)
	fmt.Fprintf(&b, "%sResetFrequency: %s,\n", indentationValues, sr.ResetFrequency)
	fmt.Fprintf(&b, "%sPlayerStatSets: %s,\n", indentationValues, sr.PlayerStatSets)
	fmt.Fprintf(&b, "%sDefaultStatValues: %s,\n", indentationValues, sr.DefaultStatValues)
	fmt.Fprintf(&b, "%s}", indentationEnd)

	return b.String()
}

// NewStatboardResult returns a new StatboardResult
func NewStatboardResult() StatboardResult {
	return StatboardResult{
		BoardID:           types.NewUInt32(0),
		ContextID:         types.NewUInt32(0),
		ResetFrequency:    types.NewUInt32(0),
		PlayerStatSets:    types.NewList[PlayerStatSet](),
		DefaultStatValues: types.NewList[PropertyVariant](),
	}
}
