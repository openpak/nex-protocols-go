// Package types implements all the types used by the Uplay Win protocol
package types

import (
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go/v2/types"
)

// UplaySection is a type within the Uplay Win protocol
type UplaySection struct {
	types.Structure
	Key          types.String
	Name         types.String
	TypeName     types.String
	MenuTypeName types.String
	ContentList  types.List[UplaySectionContent]
	GameCode     types.String
	PlatformCode types.String
}

// WriteTo writes the UplaySection to the given writable
func (s UplaySection) WriteTo(writable types.Writable) {
	contentWritable := writable.CopyNew()

	s.Key.WriteTo(contentWritable)
	s.Name.WriteTo(contentWritable)
	s.TypeName.WriteTo(contentWritable)
	s.MenuTypeName.WriteTo(contentWritable)
	s.ContentList.WriteTo(contentWritable)
	s.GameCode.WriteTo(contentWritable)
	s.PlatformCode.WriteTo(contentWritable)

	content := contentWritable.Bytes()

	s.WriteHeaderTo(writable, uint32(len(content)))

	writable.Write(content)
}

// ExtractFrom extracts the UplaySection from the given readable
func (s *UplaySection) ExtractFrom(readable types.Readable) error {
	if err := s.ExtractHeaderFrom(readable); err != nil {
		return fmt.Errorf("failed to extract UplaySection header. %s", err.Error())
	}

	if err := s.Key.ExtractFrom(readable); err != nil {
		return fmt.Errorf("failed to extract UplaySection.Key. %s", err.Error())
	}

	if err := s.Name.ExtractFrom(readable); err != nil {
		return fmt.Errorf("failed to extract UplaySection.Name. %s", err.Error())
	}

	if err := s.TypeName.ExtractFrom(readable); err != nil {
		return fmt.Errorf("failed to extract UplaySection.TypeName. %s", err.Error())
	}

	if err := s.MenuTypeName.ExtractFrom(readable); err != nil {
		return fmt.Errorf("failed to extract UplaySection.MenuTypeName. %s", err.Error())
	}

	if err := s.ContentList.ExtractFrom(readable); err != nil {
		return fmt.Errorf("failed to extract UplaySection.ContentList. %s", err.Error())
	}

	if err := s.GameCode.ExtractFrom(readable); err != nil {
		return fmt.Errorf("failed to extract UplaySection.GameCode. %s", err.Error())
	}

	if err := s.PlatformCode.ExtractFrom(readable); err != nil {
		return fmt.Errorf("failed to extract UplaySection.PlatformCode. %s", err.Error())
	}

	return nil
}

// Copy returns a new copied instance of UplaySection
func (s UplaySection) Copy() types.RVType {
	copied := NewUplaySection()

	copied.StructureVersion = s.StructureVersion
	copied.Key = s.Key.Copy().(types.String)
	copied.Name = s.Name.Copy().(types.String)
	copied.TypeName = s.TypeName.Copy().(types.String)
	copied.MenuTypeName = s.MenuTypeName.Copy().(types.String)
	copied.ContentList = s.ContentList.Copy().(types.List[UplaySectionContent])
	copied.GameCode = s.GameCode.Copy().(types.String)
	copied.PlatformCode = s.PlatformCode.Copy().(types.String)

	return copied
}

// Equals checks if the given UplaySection contains the same data as the current UplaySection
func (s UplaySection) Equals(o types.RVType) bool {
	other, ok := o.(UplaySection)
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

	if !s.TypeName.Equals(other.TypeName) {
		return false
	}

	if !s.MenuTypeName.Equals(other.MenuTypeName) {
		return false
	}

	if !s.ContentList.Equals(other.ContentList) {
		return false
	}

	if !s.GameCode.Equals(other.GameCode) {
		return false
	}

	if !s.PlatformCode.Equals(other.PlatformCode) {
		return false
	}

	return true
}

// CopyRef copies the current value of the UplaySection
// and returns a pointer to the new copy
func (s UplaySection) CopyRef() types.RVTypePtr {
	copied := s.Copy().(UplaySection)
	return &copied
}

// Deref takes a pointer to the UplaySection
// and dereferences it to the raw value.
// Only useful when working with an instance of RVTypePtr
func (s *UplaySection) Deref() types.RVType {
	return *s
}

// String returns the string representation of the UplaySection
func (s UplaySection) String() string {
	return s.FormatToString(0)
}

// FormatToString pretty-prints the UplaySection using the provided indentation level
func (s UplaySection) FormatToString(indentationLevel int) string {
	indentationValues := strings.Repeat("\t", indentationLevel+1)
	indentationEnd := strings.Repeat("\t", indentationLevel)

	var b strings.Builder

	b.WriteString("UplaySection{\n")
	fmt.Fprintf(&b, "%sKey: %s,\n", indentationValues, s.Key)
	fmt.Fprintf(&b, "%sName: %s,\n", indentationValues, s.Name)
	fmt.Fprintf(&b, "%sTypeName: %s,\n", indentationValues, s.TypeName)
	fmt.Fprintf(&b, "%sMenuTypeName: %s,\n", indentationValues, s.MenuTypeName)
	fmt.Fprintf(&b, "%sContentList: %s,\n", indentationValues, s.ContentList)
	fmt.Fprintf(&b, "%sGameCode: %s,\n", indentationValues, s.GameCode)
	fmt.Fprintf(&b, "%sPlatformCode: %s,\n", indentationValues, s.PlatformCode)
	fmt.Fprintf(&b, "%s}", indentationEnd)

	return b.String()
}

// NewUplaySection returns a new UplaySection
func NewUplaySection() UplaySection {
	return UplaySection{
		Key:          types.NewString(""),
		Name:         types.NewString(""),
		TypeName:     types.NewString(""),
		MenuTypeName: types.NewString(""),
		ContentList:  types.NewList[UplaySectionContent](),
		GameCode:     types.NewString(""),
		PlatformCode: types.NewString(""),
	}
}
