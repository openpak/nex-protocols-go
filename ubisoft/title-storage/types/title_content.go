// Package types implements the structures used by the Title Storage protocol

package types

import (
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go/v2/types"
)

// TitleContent is a type within the Title Storage protocol
// and inherits from Data
type TitleContent struct {
	types.Structure
	types.Data
	Id            types.UInt32
	Name          types.String
	Description   types.String
	Size          types.Int64
	ContentType   types.String
	Flags         types.UInt32
	ExtendedFlags types.UInt32
}

// WriteTo writes the TitleContent to the given writable
func (tc TitleContent) WriteTo(writable types.Writable) {
	tc.Data.WriteTo(writable)

	contentWritable := writable.CopyNew()

	tc.Id.WriteTo(contentWritable)
	tc.Name.WriteTo(contentWritable)
	tc.Description.WriteTo(contentWritable)
	tc.Size.WriteTo(contentWritable)
	tc.ContentType.WriteTo(contentWritable)
	tc.Flags.WriteTo(contentWritable)
	tc.ExtendedFlags.WriteTo(contentWritable)

	content := contentWritable.Bytes()

	tc.WriteHeaderTo(writable, uint32(len(content)))

	writable.Write(content)
}

// ExtractFrom extracts the TitleContent from the given readable
func (tc *TitleContent) ExtractFrom(readable types.Readable) error {
	if err := tc.Data.ExtractFrom(readable); err != nil {
		return fmt.Errorf("failed to extract TitleContent.Data. %s", err.Error())
	}

	if err := tc.ExtractHeaderFrom(readable); err != nil {
		return fmt.Errorf("failed to extract TitleContent header. %s", err.Error())
	}

	if err := tc.Id.ExtractFrom(readable); err != nil {
		return fmt.Errorf("failed to extract TitleContent.Id. %s", err.Error())
	}

	if err := tc.Name.ExtractFrom(readable); err != nil {
		return fmt.Errorf("failed to extract TitleContent.Name. %s", err.Error())
	}

	if err := tc.Description.ExtractFrom(readable); err != nil {
		return fmt.Errorf("failed to extract TitleContent.Description. %s", err.Error())
	}

	if err := tc.Size.ExtractFrom(readable); err != nil {
		return fmt.Errorf("failed to extract TitleContent.Size. %s", err.Error())
	}

	if err := tc.ContentType.ExtractFrom(readable); err != nil {
		return fmt.Errorf("failed to extract TitleContent.ContentType. %s", err.Error())
	}

	if err := tc.Flags.ExtractFrom(readable); err != nil {
		return fmt.Errorf("failed to extract TitleContent.Flags. %s", err.Error())
	}

	if err := tc.ExtendedFlags.ExtractFrom(readable); err != nil {
		return fmt.Errorf("failed to extract TitleContent.ExtendedFlags. %s", err.Error())
	}

	return nil
}

// Copy returns a new copied instance of TitleContent
func (tc TitleContent) Copy() types.RVType {
	copied := NewTitleContent()

	copied.StructureVersion = tc.StructureVersion
	copied.Data = tc.Data.Copy().(types.Data)
	copied.Id = tc.Id.Copy().(types.UInt32)
	copied.Name = tc.Name.Copy().(types.String)
	copied.Description = tc.Description.Copy().(types.String)
	copied.Size = tc.Size.Copy().(types.Int64)
	copied.ContentType = tc.ContentType.Copy().(types.String)
	copied.Flags = tc.Flags.Copy().(types.UInt32)
	copied.ExtendedFlags = tc.ExtendedFlags.Copy().(types.UInt32)

	return copied
}

// Equals checks if the given TitleContent contains the same data as the current TitleContent
func (tc TitleContent) Equals(o types.RVType) bool {
	other, ok := o.(TitleContent)
	if !ok {
		return false
	}

	if tc.StructureVersion != other.StructureVersion {
		return false
	}

	if !tc.Data.Equals(other.Data) {
		return false
	}

	if !(tc.Id.Equals(other.Id)) {
		return false
	}

	if !(tc.Name.Equals(other.Name)) {
		return false
	}

	if !(tc.Description.Equals(other.Description)) {
		return false
	}

	if !(tc.Size.Equals(other.Size)) {
		return false
	}

	if !(tc.ContentType.Equals(other.ContentType)) {
		return false
	}

	if !(tc.Flags.Equals(other.Flags)) {
		return false
	}

	if !(tc.ExtendedFlags.Equals(other.ExtendedFlags)) {
		return false
	}

	return true
}

// CopyRef copies the current value of the TitleContent
// and returns a pointer to the new copy
func (tc TitleContent) CopyRef() types.RVTypePtr {
	copied := tc.Copy().(TitleContent)
	return &copied
}

// Deref takes a pointer to the TitleContent
// and dereferences it to the raw value.
// Only useful when working with an instance of RVTypePtr
func (tc *TitleContent) Deref() types.RVType {
	return *tc
}

// String returns the string representation of the TitleContent
func (tc TitleContent) String() string {
	return tc.FormatToString(0)
}

// FormatToString pretty-prints the TitleContent using the provided indentation level
func (tc TitleContent) FormatToString(indentationLevel int) string {
	indentationValues := strings.Repeat("\t", indentationLevel+1)
	indentationEnd := strings.Repeat("\t", indentationLevel)

	var b strings.Builder

	b.WriteString("TitleContent{\n")
	fmt.Fprintf(&b, "%sData (parent): %s,\n", indentationValues, tc.Data.FormatToString(indentationLevel+1))
	fmt.Fprintf(&b, "%sId: %s,\n", indentationValues, tc.Id)
	fmt.Fprintf(&b, "%sName: %s,\n", indentationValues, tc.Name)
	fmt.Fprintf(&b, "%sDescription: %s,\n", indentationValues, tc.Description)
	fmt.Fprintf(&b, "%sSize: %s,\n", indentationValues, tc.Size)
	fmt.Fprintf(&b, "%sContentType: %s,\n", indentationValues, tc.ContentType)
	fmt.Fprintf(&b, "%sFlags: %s,\n", indentationValues, tc.Flags)
	fmt.Fprintf(&b, "%sExtendedFlags: %s,\n", indentationValues, tc.ExtendedFlags)
	fmt.Fprintf(&b, "%s}", indentationEnd)

	return b.String()
}

// NewTitleContent returns a new TitleContent
func NewTitleContent() TitleContent {
	return TitleContent{
		Data:          types.NewData(),
		Id:            types.NewUInt32(0),
		Name:          types.NewString(""),
		Description:   types.NewString(""),
		Size:          types.NewInt64(0),
		ContentType:   types.NewString(""),
		Flags:         types.NewUInt32(0),
		ExtendedFlags: types.NewUInt32(0),
	}
}
