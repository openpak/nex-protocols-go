package types

import (
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go/v2/types"
)

// WeightedTag is a type within the User Storage protocol
type WeightedTag struct {
	types.Structure
	Id                 types.UInt32
	NumberOfOccurences types.UInt32
}

// WriteTo writes the WeightedTag to the given writable
func (wt WeightedTag) WriteTo(writable types.Writable) {
	contentWritable := writable.CopyNew()

	wt.Id.WriteTo(contentWritable)
	wt.NumberOfOccurences.WriteTo(contentWritable)

	content := contentWritable.Bytes()

	wt.WriteHeaderTo(writable, uint32(len(content)))

	writable.Write(content)
}

// ExtractFrom extracts the WeightedTag from the given readable
func (wt *WeightedTag) ExtractFrom(readable types.Readable) error {
	if err := wt.ExtractHeaderFrom(readable); err != nil {
		return fmt.Errorf("failed to extract WeightedTag header. %s", err.Error())
	}

	if err := wt.Id.ExtractFrom(readable); err != nil {
		return fmt.Errorf("failed to extract WeightedTag.Id. %s", err.Error())
	}

	if err := wt.NumberOfOccurences.ExtractFrom(readable); err != nil {
		return fmt.Errorf("failed to extract WeightedTag.NumberOfOccurences. %s", err.Error())
	}

	return nil
}

// Copy returns a new copied instance of WeightedTag
func (wt WeightedTag) Copy() types.RVType {
	copied := NewWeightedTag()

	copied.StructureVersion = wt.StructureVersion
	copied.Id = wt.Id.Copy().(types.UInt32)
	copied.NumberOfOccurences = wt.NumberOfOccurences.Copy().(types.UInt32)

	return copied
}

// Equals checks if the given WeightedTag contains the same data as the current WeightedTag
func (wt WeightedTag) Equals(o types.RVType) bool {
	other, ok := o.(WeightedTag)
	if !ok {
		return false
	}

	if wt.StructureVersion != other.StructureVersion {
		return false
	}

	if !(wt.Id.Equals(other.Id)) {
		return false
	}

	if !(wt.NumberOfOccurences.Equals(other.NumberOfOccurences)) {
		return false
	}

	return true
}

// CopyRef copies the current value of the WeightedTag
// and returns a pointer to the new copy
func (wt WeightedTag) CopyRef() types.RVTypePtr {
	copied := wt.Copy().(WeightedTag)
	return &copied
}

// Deref takes a pointer to the WeightedTag
// and dereferences it to the raw value.
// Only useful when working with an instance of RVTypePtr
func (wt *WeightedTag) Deref() types.RVType {
	return *wt
}

// String returns the string representation of the WeightedTag
func (wt WeightedTag) String() string {
	return wt.FormatToString(0)
}

// FormatToString pretty-prints the WeightedTag using the provided indentation level
func (wt WeightedTag) FormatToString(indentationLevel int) string {
	indentationValues := strings.Repeat("\t", indentationLevel+1)
	indentationEnd := strings.Repeat("\t", indentationLevel)

	var b strings.Builder

	b.WriteString("WeightedTag{\n")
	fmt.Fprintf(&b, "%sId: %s,\n", indentationValues, wt.Id)
	fmt.Fprintf(&b, "%sNumberOfOccurences: %s,\n", indentationValues, wt.NumberOfOccurences)
	fmt.Fprintf(&b, "%s}", indentationEnd)

	return b.String()
}

// NewWeightedTag returns a new WeightedTag
func NewWeightedTag() WeightedTag {
	return WeightedTag{
		Id:                 types.NewUInt32(0),
		NumberOfOccurences: types.NewUInt32(0),
	}
}
