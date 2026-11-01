// Package types implements the structures used by the Title Storage Admin protocol

package types

import (
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go/v2/types"
)

// ContentTypeExtensionField is a type within the Title Storage Admin protocol
type ContentTypeExtensionField struct {
	types.Structure
	FieldName types.String
	FieldType types.String
}

// WriteTo writes the ContentTypeExtensionField to the given writable
func (ctef ContentTypeExtensionField) WriteTo(writable types.Writable) {
	contentWritable := writable.CopyNew()

	ctef.FieldName.WriteTo(contentWritable)
	ctef.FieldType.WriteTo(contentWritable)

	content := contentWritable.Bytes()

	ctef.WriteHeaderTo(writable, uint32(len(content)))

	writable.Write(content)
}

// ExtractFrom extracts the ContentTypeExtensionField from the given readable
func (ctef *ContentTypeExtensionField) ExtractFrom(readable types.Readable) error {
	if err := ctef.ExtractHeaderFrom(readable); err != nil {
		return fmt.Errorf("failed to extract ContentTypeExtensionField header. %s", err.Error())
	}

	if err := ctef.FieldName.ExtractFrom(readable); err != nil {
		return fmt.Errorf("failed to extract ContentTypeExtensionField.FieldName. %s", err.Error())
	}

	if err := ctef.FieldType.ExtractFrom(readable); err != nil {
		return fmt.Errorf("failed to extract ContentTypeExtensionField.FieldType. %s", err.Error())
	}

	return nil
}

// Copy returns a new copied instance of ContentTypeExtensionField
func (ctef ContentTypeExtensionField) Copy() types.RVType {
	copied := NewContentTypeExtensionField()

	copied.StructureVersion = ctef.StructureVersion
	copied.FieldName = ctef.FieldName.Copy().(types.String)
	copied.FieldType = ctef.FieldType.Copy().(types.String)

	return copied
}

// Equals checks if the given ContentTypeExtensionField contains the same data as the current ContentTypeExtensionField
func (ctef ContentTypeExtensionField) Equals(o types.RVType) bool {
	other, ok := o.(ContentTypeExtensionField)
	if !ok {
		return false
	}

	if ctef.StructureVersion != other.StructureVersion {
		return false
	}

	if !(ctef.FieldName.Equals(other.FieldName)) {
		return false
	}

	if !(ctef.FieldType.Equals(other.FieldType)) {
		return false
	}

	return true
}

// CopyRef copies the current value of the ContentTypeExtensionField
// and returns a pointer to the new copy
func (ctef ContentTypeExtensionField) CopyRef() types.RVTypePtr {
	copied := ctef.Copy().(ContentTypeExtensionField)
	return &copied
}

// Deref takes a pointer to the ContentTypeExtensionField
// and dereferences it to the raw value.
// Only useful when working with an instance of RVTypePtr
func (ctef *ContentTypeExtensionField) Deref() types.RVType {
	return *ctef
}

// String returns the string representation of the ContentTypeExtensionField
func (ctef ContentTypeExtensionField) String() string {
	return ctef.FormatToString(0)
}

// FormatToString pretty-prints the ContentTypeExtensionField using the provided indentation level
func (ctef ContentTypeExtensionField) FormatToString(indentationLevel int) string {
	indentationValues := strings.Repeat("\t", indentationLevel+1)
	indentationEnd := strings.Repeat("\t", indentationLevel)

	var b strings.Builder

	b.WriteString("ContentTypeExtensionField{\n")
	fmt.Fprintf(&b, "%sFieldName: %s,\n", indentationValues, ctef.FieldName)
	fmt.Fprintf(&b, "%sFieldType: %s,\n", indentationValues, ctef.FieldType)
	fmt.Fprintf(&b, "%s}", indentationEnd)

	return b.String()
}

// NewContentTypeExtensionField returns a new ContentTypeExtensionField
func NewContentTypeExtensionField() ContentTypeExtensionField {
	return ContentTypeExtensionField{
		FieldName: types.NewString(""),
		FieldType: types.NewString(""),
	}
}
