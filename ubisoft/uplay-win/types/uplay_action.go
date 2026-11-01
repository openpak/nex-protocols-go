// Package types implements all the types used by the Uplay Win protocol
package types

import (
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go/v2/types"
)

// UplayAction is a type within the Uplay Win protocol
type UplayAction struct {
	types.Structure
	Code        types.String
	Name        types.String
	Description types.String
	Value       types.Int32
	GameCode    types.String
	Platforms   types.List[UplayActionPlatform]
}

// WriteTo writes the UplayAction to the given writable
func (s UplayAction) WriteTo(writable types.Writable) {
	contentWritable := writable.CopyNew()

	s.Code.WriteTo(contentWritable)
	s.Name.WriteTo(contentWritable)
	s.Description.WriteTo(contentWritable)
	s.Value.WriteTo(contentWritable)
	s.GameCode.WriteTo(contentWritable)
	s.Platforms.WriteTo(contentWritable)

	content := contentWritable.Bytes()

	s.WriteHeaderTo(writable, uint32(len(content)))

	writable.Write(content)
}

// ExtractFrom extracts the UplayAction from the given readable
func (s *UplayAction) ExtractFrom(readable types.Readable) error {
	if err := s.ExtractHeaderFrom(readable); err != nil {
		return fmt.Errorf("failed to extract UplayAction header. %s", err.Error())
	}

	if err := s.Code.ExtractFrom(readable); err != nil {
		return fmt.Errorf("failed to extract UplayAction.Code. %s", err.Error())
	}

	if err := s.Name.ExtractFrom(readable); err != nil {
		return fmt.Errorf("failed to extract UplayAction.Name. %s", err.Error())
	}

	if err := s.Description.ExtractFrom(readable); err != nil {
		return fmt.Errorf("failed to extract UplayAction.Description. %s", err.Error())
	}

	if err := s.Value.ExtractFrom(readable); err != nil {
		return fmt.Errorf("failed to extract UplayAction.Value. %s", err.Error())
	}

	if err := s.GameCode.ExtractFrom(readable); err != nil {
		return fmt.Errorf("failed to extract UplayAction.GameCode. %s", err.Error())
	}

	if err := s.Platforms.ExtractFrom(readable); err != nil {
		return fmt.Errorf("failed to extract UplayAction.Platforms. %s", err.Error())
	}

	return nil
}

// Copy returns a new copied instance of UplayAction
func (s UplayAction) Copy() types.RVType {
	copied := NewUplayAction()

	copied.StructureVersion = s.StructureVersion
	copied.Code = s.Code.Copy().(types.String)
	copied.Name = s.Name.Copy().(types.String)
	copied.Description = s.Description.Copy().(types.String)
	copied.Value = s.Value.Copy().(types.Int32)
	copied.GameCode = s.GameCode.Copy().(types.String)
	copied.Platforms = s.Platforms.Copy().(types.List[UplayActionPlatform])

	return copied
}

// Equals checks if the given UplayAction contains the same data as the current UplayAction
func (s UplayAction) Equals(o types.RVType) bool {
	other, ok := o.(UplayAction)
	if !ok {
		return false
	}

	if s.StructureVersion != other.StructureVersion {
		return false
	}

	if !s.Code.Equals(other.Code) {
		return false
	}

	if !s.Name.Equals(other.Name) {
		return false
	}

	if !s.Description.Equals(other.Description) {
		return false
	}

	if !s.Value.Equals(other.Value) {
		return false
	}

	if !s.GameCode.Equals(other.GameCode) {
		return false
	}

	if !s.Platforms.Equals(other.Platforms) {
		return false
	}

	return true
}

// CopyRef copies the current value of the UplayAction
// and returns a pointer to the new copy
func (s UplayAction) CopyRef() types.RVTypePtr {
	copied := s.Copy().(UplayAction)
	return &copied
}

// Deref takes a pointer to the UplayAction
// and dereferences it to the raw value.
// Only useful when working with an instance of RVTypePtr
func (s *UplayAction) Deref() types.RVType {
	return *s
}

// String returns the string representation of the UplayAction
func (s UplayAction) String() string {
	return s.FormatToString(0)
}

// FormatToString pretty-prints the UplayAction using the provided indentation level
func (s UplayAction) FormatToString(indentationLevel int) string {
	indentationValues := strings.Repeat("\t", indentationLevel+1)
	indentationEnd := strings.Repeat("\t", indentationLevel)

	var b strings.Builder

	b.WriteString("UplayAction{\n")
	fmt.Fprintf(&b, "%sCode: %s,\n", indentationValues, s.Code)
	fmt.Fprintf(&b, "%sName: %s,\n", indentationValues, s.Name)
	fmt.Fprintf(&b, "%sDescription: %s,\n", indentationValues, s.Description)
	fmt.Fprintf(&b, "%sValue: %s,\n", indentationValues, s.Value)
	fmt.Fprintf(&b, "%sGameCode: %s,\n", indentationValues, s.GameCode)
	fmt.Fprintf(&b, "%sPlatforms: %s,\n", indentationValues, s.Platforms)
	fmt.Fprintf(&b, "%s}", indentationEnd)

	return b.String()
}

// NewUplayAction returns a new UplayAction
func NewUplayAction() UplayAction {
	return UplayAction{
		Code:        types.NewString(""),
		Name:        types.NewString(""),
		Description: types.NewString(""),
		Value:       types.NewInt32(0),
		GameCode:    types.NewString(""),
		Platforms:   types.NewList[UplayActionPlatform](),
	}
}
