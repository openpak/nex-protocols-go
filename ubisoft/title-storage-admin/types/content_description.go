package types

import (
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go/v2/types"
)

// ContentDescription is a type within the Title Storage Admin protocol
type ContentDescription struct {
	types.Structure
	ContentDisplayName types.String
	ContentDescription types.String
}

// WriteTo writes the ContentDescription to the given writable
func (cd ContentDescription) WriteTo(writable types.Writable) {
	contentWritable := writable.CopyNew()

	cd.ContentDisplayName.WriteTo(contentWritable)
	cd.ContentDescription.WriteTo(contentWritable)

	content := contentWritable.Bytes()

	cd.WriteHeaderTo(writable, uint32(len(content)))

	writable.Write(content)
}

// ExtractFrom extracts the ContentDescription from the given readable
func (cd *ContentDescription) ExtractFrom(readable types.Readable) error {
	if err := cd.ExtractHeaderFrom(readable); err != nil {
		return fmt.Errorf("failed to extract ContentDescription header. %s", err.Error())
	}

	if err := cd.ContentDisplayName.ExtractFrom(readable); err != nil {
		return fmt.Errorf("failed to extract ContentDescription.ContentDisplayName. %s", err.Error())
	}

	if err := cd.ContentDescription.ExtractFrom(readable); err != nil {
		return fmt.Errorf("failed to extract ContentDescription.ContentDescription. %s", err.Error())
	}

	return nil
}

// Copy returns a new copied instance of ContentDescription
func (cd ContentDescription) Copy() types.RVType {
	copied := NewContentDescription()

	copied.StructureVersion = cd.StructureVersion
	copied.ContentDisplayName = cd.ContentDisplayName.Copy().(types.String)
	copied.ContentDescription = cd.ContentDescription.Copy().(types.String)

	return copied
}

// Equals checks if the given ContentDescription contains the same data as the current ContentDescription
func (cd ContentDescription) Equals(o types.RVType) bool {
	other, ok := o.(ContentDescription)
	if !ok {
		return false
	}

	if cd.StructureVersion != other.StructureVersion {
		return false
	}

	if !(cd.ContentDisplayName.Equals(other.ContentDisplayName)) {
		return false
	}

	if !(cd.ContentDescription.Equals(other.ContentDescription)) {
		return false
	}

	return true
}

// CopyRef copies the current value of the ContentDescription
// and returns a pointer to the new copy
func (cd ContentDescription) CopyRef() types.RVTypePtr {
	copied := cd.Copy().(ContentDescription)
	return &copied
}

// Deref takes a pointer to the ContentDescription
// and dereferences it to the raw value.
// Only useful when working with an instance of RVTypePtr
func (cd *ContentDescription) Deref() types.RVType {
	return *cd
}

// String returns the string representation of the ContentDescription
func (cd ContentDescription) String() string {
	return cd.FormatToString(0)
}

// FormatToString pretty-prints the ContentDescription using the provided indentation level
func (cd ContentDescription) FormatToString(indentationLevel int) string {
	indentationValues := strings.Repeat("\t", indentationLevel+1)
	indentationEnd := strings.Repeat("\t", indentationLevel)

	var b strings.Builder

	b.WriteString("ContentDescription{\n")
	fmt.Fprintf(&b, "%sContentDisplayName: %s,\n", indentationValues, cd.ContentDisplayName)
	fmt.Fprintf(&b, "%sContentDescription: %s,\n", indentationValues, cd.ContentDescription)
	fmt.Fprintf(&b, "%s}", indentationEnd)

	return b.String()
}

// NewContentDescription returns a new ContentDescription
func NewContentDescription() ContentDescription {
	return ContentDescription{
		ContentDisplayName: types.NewString(""),
		ContentDescription: types.NewString(""),
	}
}
