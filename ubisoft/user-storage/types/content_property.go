// Package types implements the structures used by the User Storage protocol

package types

import (
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go/v2/types"
)

// ContentProperty is a type within the User Storage protocol
type ContentProperty struct {
	types.Structure
	ID    types.UInt32
	Value types.Variant
}

// WriteTo writes the ContentProperty to the given writable
func (cp ContentProperty) WriteTo(writable types.Writable) {
	contentWritable := writable.CopyNew()

	cp.ID.WriteTo(contentWritable)
	cp.Value.WriteTo(contentWritable)

	content := contentWritable.Bytes()

	cp.WriteHeaderTo(writable, uint32(len(content)))

	writable.Write(content)
}

// ExtractFrom extracts the ContentProperty from the given readable
func (cp *ContentProperty) ExtractFrom(readable types.Readable) error {
	if err := cp.ExtractHeaderFrom(readable); err != nil {
		return fmt.Errorf("failed to extract ContentProperty header. %s", err.Error())
	}

	if err := cp.ID.ExtractFrom(readable); err != nil {
		return fmt.Errorf("failed to extract ContentProperty.ID. %s", err.Error())
	}

	if err := cp.Value.ExtractFrom(readable); err != nil {
		return fmt.Errorf("failed to extract ContentProperty.Value. %s", err.Error())
	}

	return nil
}

// Copy returns a new copied instance of ContentProperty
func (cp ContentProperty) Copy() types.RVType {
	copied := NewContentProperty()

	copied.StructureVersion = cp.StructureVersion
	copied.ID = cp.ID.Copy().(types.UInt32)
	copied.Value = cp.Value.Copy().(types.Variant)

	return copied
}

// Equals checks if the given ContentProperty contains the same data as the current ContentProperty
func (cp ContentProperty) Equals(o types.RVType) bool {
	other, ok := o.(ContentProperty)
	if !ok {
		return false
	}

	if cp.StructureVersion != other.StructureVersion {
		return false
	}

	if !(cp.ID.Equals(other.ID)) {
		return false
	}

	if !(cp.Value.Equals(other.Value)) {
		return false
	}

	return true
}

// CopyRef copies the current value of the ContentProperty
// and returns a pointer to the new copy
func (cp ContentProperty) CopyRef() types.RVTypePtr {
	copied := cp.Copy().(ContentProperty)
	return &copied
}

// Deref takes a pointer to the ContentProperty
// and dereferences it to the raw value.
// Only useful when working with an instance of RVTypePtr
func (cp *ContentProperty) Deref() types.RVType {
	return *cp
}

// String returns the string representation of the ContentProperty
func (cp ContentProperty) String() string {
	return cp.FormatToString(0)
}

// FormatToString pretty-prints the ContentProperty using the provided indentation level
func (cp ContentProperty) FormatToString(indentationLevel int) string {
	indentationValues := strings.Repeat("\t", indentationLevel+1)
	indentationEnd := strings.Repeat("\t", indentationLevel)

	var b strings.Builder

	b.WriteString("ContentProperty{\n")
	fmt.Fprintf(&b, "%sID: %s,\n", indentationValues, cp.ID)
	fmt.Fprintf(&b, "%sValue: %s,\n", indentationValues, cp.Value)
	fmt.Fprintf(&b, "%s}", indentationEnd)

	return b.String()
}

// NewContentProperty returns a new ContentProperty
func NewContentProperty() ContentProperty {
	return ContentProperty{
		ID:    types.NewUInt32(0),
		Value: types.NewVariant(),
	}
}
