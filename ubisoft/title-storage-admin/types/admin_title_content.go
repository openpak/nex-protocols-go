package types

import (
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go/v2/types"
)

// AdminTitleContent is a type within the Title Storage Admin protocol
type AdminTitleContent struct {
	types.Structure
	Id                 types.UInt32
	Name               types.String
	ContentDescription ContentDescription
	Flags              types.UInt32
	ExtendedFlags      types.UInt32
	DateAvailable      types.DateTime
	DateExpired        types.DateTime
	ContentType        types.String
	DownloadURL        types.String
}

// WriteTo writes the AdminTitleContent to the given writable
func (atc AdminTitleContent) WriteTo(writable types.Writable) {
	contentWritable := writable.CopyNew()

	atc.Id.WriteTo(contentWritable)
	atc.Name.WriteTo(contentWritable)
	atc.ContentDescription.WriteTo(contentWritable)
	atc.Flags.WriteTo(contentWritable)
	atc.ExtendedFlags.WriteTo(contentWritable)
	atc.DateAvailable.WriteTo(contentWritable)
	atc.DateExpired.WriteTo(contentWritable)
	atc.ContentType.WriteTo(contentWritable)
	atc.DownloadURL.WriteTo(contentWritable)

	content := contentWritable.Bytes()

	atc.WriteHeaderTo(writable, uint32(len(content)))

	writable.Write(content)
}

// ExtractFrom extracts the AdminTitleContent from the given readable
func (atc *AdminTitleContent) ExtractFrom(readable types.Readable) error {
	if err := atc.ExtractHeaderFrom(readable); err != nil {
		return fmt.Errorf("failed to extract AdminTitleContent header. %s", err.Error())
	}

	if err := atc.Id.ExtractFrom(readable); err != nil {
		return fmt.Errorf("failed to extract AdminTitleContent.Id. %s", err.Error())
	}

	if err := atc.Name.ExtractFrom(readable); err != nil {
		return fmt.Errorf("failed to extract AdminTitleContent.Name. %s", err.Error())
	}

	if err := atc.ContentDescription.ExtractFrom(readable); err != nil {
		return fmt.Errorf("failed to extract AdminTitleContent.ContentDescription. %s", err.Error())
	}

	if err := atc.Flags.ExtractFrom(readable); err != nil {
		return fmt.Errorf("failed to extract AdminTitleContent.Flags. %s", err.Error())
	}

	if err := atc.ExtendedFlags.ExtractFrom(readable); err != nil {
		return fmt.Errorf("failed to extract AdminTitleContent.ExtendedFlags. %s", err.Error())
	}

	if err := atc.DateAvailable.ExtractFrom(readable); err != nil {
		return fmt.Errorf("failed to extract AdminTitleContent.DateAvailable. %s", err.Error())
	}

	if err := atc.DateExpired.ExtractFrom(readable); err != nil {
		return fmt.Errorf("failed to extract AdminTitleContent.DateExpired. %s", err.Error())
	}

	if err := atc.ContentType.ExtractFrom(readable); err != nil {
		return fmt.Errorf("failed to extract AdminTitleContent.ContentType. %s", err.Error())
	}

	if err := atc.DownloadURL.ExtractFrom(readable); err != nil {
		return fmt.Errorf("failed to extract AdminTitleContent.DownloadURL. %s", err.Error())
	}

	return nil
}

// Copy returns a new copied instance of AdminTitleContent
func (atc AdminTitleContent) Copy() types.RVType {
	copied := NewAdminTitleContent()

	copied.StructureVersion = atc.StructureVersion
	copied.Id = atc.Id.Copy().(types.UInt32)
	copied.Name = atc.Name.Copy().(types.String)
	copied.ContentDescription = atc.ContentDescription.Copy().(ContentDescription)
	copied.Flags = atc.Flags.Copy().(types.UInt32)
	copied.ExtendedFlags = atc.ExtendedFlags.Copy().(types.UInt32)
	copied.DateAvailable = atc.DateAvailable.Copy().(types.DateTime)
	copied.DateExpired = atc.DateExpired.Copy().(types.DateTime)
	copied.ContentType = atc.ContentType.Copy().(types.String)
	copied.DownloadURL = atc.DownloadURL.Copy().(types.String)

	return copied
}

// Equals checks if the given AdminTitleContent contains the same data as the current AdminTitleContent
func (atc AdminTitleContent) Equals(o types.RVType) bool {
	other, ok := o.(AdminTitleContent)
	if !ok {
		return false
	}

	if atc.StructureVersion != other.StructureVersion {
		return false
	}

	if !(atc.Id.Equals(other.Id)) {
		return false
	}

	if !(atc.Name.Equals(other.Name)) {
		return false
	}

	if !(atc.ContentDescription.Equals(other.ContentDescription)) {
		return false
	}

	if !(atc.Flags.Equals(other.Flags)) {
		return false
	}

	if !(atc.ExtendedFlags.Equals(other.ExtendedFlags)) {
		return false
	}

	if !(atc.DateAvailable.Equals(other.DateAvailable)) {
		return false
	}

	if !(atc.DateExpired.Equals(other.DateExpired)) {
		return false
	}

	if !(atc.ContentType.Equals(other.ContentType)) {
		return false
	}

	if !(atc.DownloadURL.Equals(other.DownloadURL)) {
		return false
	}

	return true
}

// CopyRef copies the current value of the AdminTitleContent
// and returns a pointer to the new copy
func (atc AdminTitleContent) CopyRef() types.RVTypePtr {
	copied := atc.Copy().(AdminTitleContent)
	return &copied
}

// Deref takes a pointer to the AdminTitleContent
// and dereferences it to the raw value.
// Only useful when working with an instance of RVTypePtr
func (atc *AdminTitleContent) Deref() types.RVType {
	return *atc
}

// String returns the string representation of the AdminTitleContent
func (atc AdminTitleContent) String() string {
	return atc.FormatToString(0)
}

// FormatToString pretty-prints the AdminTitleContent using the provided indentation level
func (atc AdminTitleContent) FormatToString(indentationLevel int) string {
	indentationValues := strings.Repeat("\t", indentationLevel+1)
	indentationEnd := strings.Repeat("\t", indentationLevel)

	var b strings.Builder

	b.WriteString("AdminTitleContent{\n")
	fmt.Fprintf(&b, "%sId: %s,\n", indentationValues, atc.Id)
	fmt.Fprintf(&b, "%sName: %s,\n", indentationValues, atc.Name)
	fmt.Fprintf(&b, "%sContentDescription: %s,\n", indentationValues, atc.ContentDescription.FormatToString(indentationLevel+1))
	fmt.Fprintf(&b, "%sFlags: %s,\n", indentationValues, atc.Flags)
	fmt.Fprintf(&b, "%sExtendedFlags: %s,\n", indentationValues, atc.ExtendedFlags)
	fmt.Fprintf(&b, "%sDateAvailable: %s,\n", indentationValues, atc.DateAvailable.FormatToString(indentationLevel+1))
	fmt.Fprintf(&b, "%sDateExpired: %s,\n", indentationValues, atc.DateExpired.FormatToString(indentationLevel+1))
	fmt.Fprintf(&b, "%sContentType: %s,\n", indentationValues, atc.ContentType)
	fmt.Fprintf(&b, "%sDownloadURL: %s,\n", indentationValues, atc.DownloadURL)
	fmt.Fprintf(&b, "%s}", indentationEnd)

	return b.String()
}

// NewAdminTitleContent returns a new AdminTitleContent
func NewAdminTitleContent() AdminTitleContent {
	return AdminTitleContent{
		Id:                 types.NewUInt32(0),
		Name:               types.NewString(""),
		ContentDescription: NewContentDescription(),
		Flags:              types.NewUInt32(0),
		ExtendedFlags:      types.NewUInt32(0),
		DateAvailable:      types.NewDateTime(0),
		DateExpired:        types.NewDateTime(0),
		ContentType:        types.NewString(""),
		DownloadURL:        types.NewString(""),
	}
}
