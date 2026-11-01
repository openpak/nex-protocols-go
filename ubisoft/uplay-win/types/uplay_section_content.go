// Package types implements all the types used by the Uplay Win protocol
package types

import (
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go/v2/types"
)

// UplaySectionContent is a type within the Uplay Win protocol
type UplaySectionContent struct {
	types.Structure
	Key           types.String
	Name          types.String
	Order         types.Int16
	TypeName      types.String
	LocalizedInfo UplaySectionContentLocalized
}

// WriteTo writes the UplaySectionContent to the given writable
func (s UplaySectionContent) WriteTo(writable types.Writable) {
	contentWritable := writable.CopyNew()

	s.Key.WriteTo(contentWritable)
	s.Name.WriteTo(contentWritable)
	s.Order.WriteTo(contentWritable)
	s.TypeName.WriteTo(contentWritable)
	s.LocalizedInfo.WriteTo(contentWritable)

	content := contentWritable.Bytes()

	s.WriteHeaderTo(writable, uint32(len(content)))

	writable.Write(content)
}

// ExtractFrom extracts the UplaySectionContent from the given readable
func (s *UplaySectionContent) ExtractFrom(readable types.Readable) error {
	if err := s.ExtractHeaderFrom(readable); err != nil {
		return fmt.Errorf("failed to extract UplaySectionContent header. %s", err.Error())
	}

	if err := s.Key.ExtractFrom(readable); err != nil {
		return fmt.Errorf("failed to extract UplaySectionContent.Key. %s", err.Error())
	}

	if err := s.Name.ExtractFrom(readable); err != nil {
		return fmt.Errorf("failed to extract UplaySectionContent.Name. %s", err.Error())
	}

	if err := s.Order.ExtractFrom(readable); err != nil {
		return fmt.Errorf("failed to extract UplaySectionContent.Order. %s", err.Error())
	}

	if err := s.TypeName.ExtractFrom(readable); err != nil {
		return fmt.Errorf("failed to extract UplaySectionContent.TypeName. %s", err.Error())
	}

	if err := s.LocalizedInfo.ExtractFrom(readable); err != nil {
		return fmt.Errorf("failed to extract UplaySectionContent.LocalizedInfo. %s", err.Error())
	}

	return nil
}

// Copy returns a new copied instance of UplaySectionContent
func (s UplaySectionContent) Copy() types.RVType {
	copied := NewUplaySectionContent()

	copied.StructureVersion = s.StructureVersion
	copied.Key = s.Key.Copy().(types.String)
	copied.Name = s.Name.Copy().(types.String)
	copied.Order = s.Order.Copy().(types.Int16)
	copied.TypeName = s.TypeName.Copy().(types.String)
	copied.LocalizedInfo = s.LocalizedInfo.Copy().(UplaySectionContentLocalized)

	return copied
}

// Equals checks if the given UplaySectionContent contains the same data as the current UplaySectionContent
func (s UplaySectionContent) Equals(o types.RVType) bool {
	other, ok := o.(UplaySectionContent)
	if !ok {
		return false
	}

	if s.StructureVersion != other.StructureVersion {
		return false
	}

	if !s.Key.Equals(other.Key) {
		return false
	}

	if !s.Name.Equals(other.Name) {
		return false
	}

	if !s.Order.Equals(other.Order) {
		return false
	}

	if !s.TypeName.Equals(other.TypeName) {
		return false
	}

	if !s.LocalizedInfo.Equals(other.LocalizedInfo) {
		return false
	}

	return true
}

// CopyRef copies the current value of the UplaySectionContent
// and returns a pointer to the new copy
func (s UplaySectionContent) CopyRef() types.RVTypePtr {
	copied := s.Copy().(UplaySectionContent)
	return &copied
}

// Deref takes a pointer to the UplaySectionContent
// and dereferences it to the raw value.
// Only useful when working with an instance of RVTypePtr
func (s *UplaySectionContent) Deref() types.RVType {
	return *s
}

// String returns the string representation of the UplaySectionContent
func (s UplaySectionContent) String() string {
	return s.FormatToString(0)
}

// FormatToString pretty-prints the UplaySectionContent using the provided indentation level
func (s UplaySectionContent) FormatToString(indentationLevel int) string {
	indentationValues := strings.Repeat("\t", indentationLevel+1)
	indentationEnd := strings.Repeat("\t", indentationLevel)

	var b strings.Builder

	b.WriteString("UplaySectionContent{\n")
	fmt.Fprintf(&b, "%sKey: %s,\n", indentationValues, s.Key)
	fmt.Fprintf(&b, "%sName: %s,\n", indentationValues, s.Name)
	fmt.Fprintf(&b, "%sOrder: %s,\n", indentationValues, s.Order)
	fmt.Fprintf(&b, "%sTypeName: %s,\n", indentationValues, s.TypeName)
	fmt.Fprintf(&b, "%sLocalizedInfo: %s,\n", indentationValues, s.LocalizedInfo)
	fmt.Fprintf(&b, "%s}", indentationEnd)

	return b.String()
}

// NewUplaySectionContent returns a new UplaySectionContent
func NewUplaySectionContent() UplaySectionContent {
	return UplaySectionContent{
		Key:           types.NewString(""),
		Name:          types.NewString(""),
		Order:         types.NewInt16(0),
		TypeName:      types.NewString(""),
		LocalizedInfo: NewUplaySectionContentLocalized(),
	}
}
