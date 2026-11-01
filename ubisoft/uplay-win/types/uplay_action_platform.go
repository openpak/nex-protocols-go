// Package types implements all the types used by the Uplay Win protocol
package types

import (
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go/v2/types"
)

// UplayActionPlatform is a type within the Uplay Win protocol
type UplayActionPlatform struct {
	types.Structure
	PlatformCode types.String
	Completed    types.Bool
	SpecificKey  types.String
}

// WriteTo writes the UplayActionPlatform to the given writable
func (s UplayActionPlatform) WriteTo(writable types.Writable) {
	contentWritable := writable.CopyNew()

	s.PlatformCode.WriteTo(contentWritable)
	s.Completed.WriteTo(contentWritable)
	s.SpecificKey.WriteTo(contentWritable)

	content := contentWritable.Bytes()

	s.WriteHeaderTo(writable, uint32(len(content)))

	writable.Write(content)
}

// ExtractFrom extracts the UplayActionPlatform from the given readable
func (s *UplayActionPlatform) ExtractFrom(readable types.Readable) error {
	if err := s.ExtractHeaderFrom(readable); err != nil {
		return fmt.Errorf("failed to extract UplayActionPlatform header. %s", err.Error())
	}

	if err := s.PlatformCode.ExtractFrom(readable); err != nil {
		return fmt.Errorf("failed to extract UplayActionPlatform.PlatformCode. %s", err.Error())
	}

	if err := s.Completed.ExtractFrom(readable); err != nil {
		return fmt.Errorf("failed to extract UplayActionPlatform.Completed. %s", err.Error())
	}

	if err := s.SpecificKey.ExtractFrom(readable); err != nil {
		return fmt.Errorf("failed to extract UplayActionPlatform.SpecificKey. %s", err.Error())
	}

	return nil
}

// Copy returns a new copied instance of UplayActionPlatform
func (s UplayActionPlatform) Copy() types.RVType {
	copied := NewUplayActionPlatform()

	copied.StructureVersion = s.StructureVersion
	copied.PlatformCode = s.PlatformCode.Copy().(types.String)
	copied.Completed = s.Completed.Copy().(types.Bool)
	copied.SpecificKey = s.SpecificKey.Copy().(types.String)

	return copied
}

// Equals checks if the given UplayActionPlatform contains the same data as the current UplayActionPlatform
func (s UplayActionPlatform) Equals(o types.RVType) bool {
	other, ok := o.(UplayActionPlatform)
	if !ok {
		return false
	}

	if s.StructureVersion != other.StructureVersion {
		return false
	}

	if !s.PlatformCode.Equals(other.PlatformCode) {
		return false
	}

	if !s.Completed.Equals(other.Completed) {
		return false
	}

	if !s.SpecificKey.Equals(other.SpecificKey) {
		return false
	}

	return true
}

// CopyRef copies the current value of the UplayActionPlatform
// and returns a pointer to the new copy
func (s UplayActionPlatform) CopyRef() types.RVTypePtr {
	copied := s.Copy().(UplayActionPlatform)
	return &copied
}

// Deref takes a pointer to the UplayActionPlatform
// and dereferences it to the raw value.
// Only useful when working with an instance of RVTypePtr
func (s *UplayActionPlatform) Deref() types.RVType {
	return *s
}

// String returns the string representation of the UplayActionPlatform
func (s UplayActionPlatform) String() string {
	return s.FormatToString(0)
}

// FormatToString pretty-prints the UplayActionPlatform using the provided indentation level
func (s UplayActionPlatform) FormatToString(indentationLevel int) string {
	indentationValues := strings.Repeat("\t", indentationLevel+1)
	indentationEnd := strings.Repeat("\t", indentationLevel)

	var b strings.Builder

	b.WriteString("UplayActionPlatform{\n")
	fmt.Fprintf(&b, "%sPlatformCode: %s,\n", indentationValues, s.PlatformCode)
	fmt.Fprintf(&b, "%sCompleted: %s,\n", indentationValues, s.Completed)
	fmt.Fprintf(&b, "%sSpecificKey: %s,\n", indentationValues, s.SpecificKey)
	fmt.Fprintf(&b, "%s}", indentationEnd)

	return b.String()
}

// NewUplayActionPlatform returns a new UplayActionPlatform
func NewUplayActionPlatform() UplayActionPlatform {
	return UplayActionPlatform{
		PlatformCode: types.NewString(""),
		Completed:    types.NewBool(false),
		SpecificKey:  types.NewString(""),
	}
}
