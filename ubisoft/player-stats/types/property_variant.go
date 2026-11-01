package types

import (
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go/v2/types"
)

// PropertyVariant is a type within the Player Stats protocol
type PropertyVariant struct {
	types.Structure
	ID    types.UInt32
	Value types.Variant
}

// WriteTo writes the PropertyVariant to the given writable
func (pv PropertyVariant) WriteTo(writable types.Writable) {
	contentWritable := writable.CopyNew()

	pv.ID.WriteTo(contentWritable)
	pv.Value.WriteTo(contentWritable)

	content := contentWritable.Bytes()

	pv.WriteHeaderTo(writable, uint32(len(content)))

	writable.Write(content)
}

// ExtractFrom extracts the PropertyVariant from the given readable
func (pv *PropertyVariant) ExtractFrom(readable types.Readable) error {
	if err := pv.ExtractHeaderFrom(readable); err != nil {
		return fmt.Errorf("failed to extract PropertyVariant header. %s", err.Error())
	}

	if err := pv.ID.ExtractFrom(readable); err != nil {
		return fmt.Errorf("failed to extract PropertyVariant.ID. %s", err.Error())
	}

	if err := pv.Value.ExtractFrom(readable); err != nil {
		return fmt.Errorf("failed to extract PropertyVariant.Value. %s", err.Error())
	}

	return nil
}

// Copy returns a new copied instance of PropertyVariant
func (pv PropertyVariant) Copy() types.RVType {
	copied := NewPropertyVariant()

	copied.StructureVersion = pv.StructureVersion
	copied.ID = pv.ID.Copy().(types.UInt32)
	copied.Value = pv.Value.Copy().(types.Variant)

	return copied
}

// Equals checks if the given PropertyVariant contains the same data as the current PropertyVariant
func (pv PropertyVariant) Equals(o types.RVType) bool {
	other, ok := o.(PropertyVariant)
	if !ok {
		return false
	}

	if pv.StructureVersion != other.StructureVersion {
		return false
	}

	if !(pv.ID.Equals(other.ID)) {
		return false
	}

	if !(pv.Value.Equals(other.Value)) {
		return false
	}

	return true
}

// CopyRef copies the current value of the PropertyVariant
// and returns a pointer to the new copy
func (pv PropertyVariant) CopyRef() types.RVTypePtr {
	copied := pv.Copy().(PropertyVariant)
	return &copied
}

// Deref takes a pointer to the PropertyVariant
// and dereferences it to the raw value.
// Only useful when working with an instance of RVTypePtr
func (pv *PropertyVariant) Deref() types.RVType {
	return *pv
}

// String returns the string representation of the PropertyVariant
func (pv PropertyVariant) String() string {
	return pv.FormatToString(0)
}

// FormatToString pretty-prints the PropertyVariant using the provided indentation level
func (pv PropertyVariant) FormatToString(indentationLevel int) string {
	indentationValues := strings.Repeat("\t", indentationLevel+1)
	indentationEnd := strings.Repeat("\t", indentationLevel)

	var b strings.Builder

	b.WriteString("PropertyVariant{\n")
	fmt.Fprintf(&b, "%sID: %s,\n", indentationValues, pv.ID)
	fmt.Fprintf(&b, "%sValue: %s,\n", indentationValues, pv.Value)
	fmt.Fprintf(&b, "%s}", indentationEnd)

	return b.String()
}

// NewPropertyVariant returns a new PropertyVariant
func NewPropertyVariant() PropertyVariant {
	return PropertyVariant{
		ID:    types.NewUInt32(0),
		Value: types.NewVariant(),
	}
}
