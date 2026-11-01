package types

import (
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go/v2/types"
)

// ContentTypeDescription is a type within the Title Storage Admin protocol
type ContentTypeDescription struct {
	types.Structure
	ContentType     types.String
	ExtensionFields types.List[ContentTypeExtensionField]
}

// WriteTo writes the ContentTypeDescription to the given writable
func (ctd ContentTypeDescription) WriteTo(writable types.Writable) {
	contentWritable := writable.CopyNew()

	ctd.ContentType.WriteTo(contentWritable)
	ctd.ExtensionFields.WriteTo(contentWritable)

	content := contentWritable.Bytes()

	ctd.WriteHeaderTo(writable, uint32(len(content)))

	writable.Write(content)
}

// ExtractFrom extracts the ContentTypeDescription from the given readable
func (ctd *ContentTypeDescription) ExtractFrom(readable types.Readable) error {
	if err := ctd.ExtractHeaderFrom(readable); err != nil {
		return fmt.Errorf("failed to extract ContentTypeDescription header. %s", err.Error())
	}

	if err := ctd.ContentType.ExtractFrom(readable); err != nil {
		return fmt.Errorf("failed to extract ContentTypeDescription.ContentType. %s", err.Error())
	}

	if err := ctd.ExtensionFields.ExtractFrom(readable); err != nil {
		return fmt.Errorf("failed to extract ContentTypeDescription.ExtensionFields. %s", err.Error())
	}

	return nil
}

// Copy returns a new copied instance of ContentTypeDescription
func (ctd ContentTypeDescription) Copy() types.RVType {
	copied := NewContentTypeDescription()

	copied.StructureVersion = ctd.StructureVersion
	copied.ContentType = ctd.ContentType.Copy().(types.String)
	copied.ExtensionFields = ctd.ExtensionFields.Copy().(types.List[ContentTypeExtensionField])

	return copied
}

// Equals checks if the given ContentTypeDescription contains the same data as the current ContentTypeDescription
func (ctd ContentTypeDescription) Equals(o types.RVType) bool {
	other, ok := o.(ContentTypeDescription)
	if !ok {
		return false
	}

	if ctd.StructureVersion != other.StructureVersion {
		return false
	}

	if !(ctd.ContentType.Equals(other.ContentType)) {
		return false
	}

	if !(ctd.ExtensionFields.Equals(other.ExtensionFields)) {
		return false
	}

	return true
}

// CopyRef copies the current value of the ContentTypeDescription
// and returns a pointer to the new copy
func (ctd ContentTypeDescription) CopyRef() types.RVTypePtr {
	copied := ctd.Copy().(ContentTypeDescription)
	return &copied
}

// Deref takes a pointer to the ContentTypeDescription
// and dereferences it to the raw value.
// Only useful when working with an instance of RVTypePtr
func (ctd *ContentTypeDescription) Deref() types.RVType {
	return *ctd
}

// String returns the string representation of the ContentTypeDescription
func (ctd ContentTypeDescription) String() string {
	return ctd.FormatToString(0)
}

// FormatToString pretty-prints the ContentTypeDescription using the provided indentation level
func (ctd ContentTypeDescription) FormatToString(indentationLevel int) string {
	indentationValues := strings.Repeat("\t", indentationLevel+1)
	indentationEnd := strings.Repeat("\t", indentationLevel)

	var b strings.Builder

	b.WriteString("ContentTypeDescription{\n")
	fmt.Fprintf(&b, "%sContentType: %s,\n", indentationValues, ctd.ContentType)
	fmt.Fprintf(&b, "%sExtensionFields: %s,\n", indentationValues, ctd.ExtensionFields)
	fmt.Fprintf(&b, "%s}", indentationEnd)

	return b.String()
}

// NewContentTypeDescription returns a new ContentTypeDescription
func NewContentTypeDescription() ContentTypeDescription {
	return ContentTypeDescription{
		ContentType:     types.NewString(""),
		ExtensionFields: types.NewList[ContentTypeExtensionField](),
	}
}
