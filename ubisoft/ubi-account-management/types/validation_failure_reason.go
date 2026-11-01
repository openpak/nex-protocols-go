// Package types implements all the types used by the Ubi Account Management protocol
package types

import (
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go/v2/types"
)

// ValidationFailureReason is a type within the Ubi Account Management protocol
type ValidationFailureReason struct {
	types.Structure
	ValidationID types.UInt32
	Description  types.String
}

// WriteTo writes the ValidationFailureReason to the given writable
func (s ValidationFailureReason) WriteTo(writable types.Writable) {
	contentWritable := writable.CopyNew()

	s.ValidationID.WriteTo(contentWritable)
	s.Description.WriteTo(contentWritable)

	content := contentWritable.Bytes()

	s.WriteHeaderTo(writable, uint32(len(content)))

	writable.Write(content)
}

// ExtractFrom extracts the ValidationFailureReason from the given readable
func (s *ValidationFailureReason) ExtractFrom(readable types.Readable) error {
	if err := s.ExtractHeaderFrom(readable); err != nil {
		return fmt.Errorf("failed to extract ValidationFailureReason header. %s", err.Error())
	}

	if err := s.ValidationID.ExtractFrom(readable); err != nil {
		return fmt.Errorf("failed to extract ValidationFailureReason.ValidationID. %s", err.Error())
	}

	if err := s.Description.ExtractFrom(readable); err != nil {
		return fmt.Errorf("failed to extract ValidationFailureReason.Description. %s", err.Error())
	}

	return nil
}

// Copy returns a new copied instance of ValidationFailureReason
func (s ValidationFailureReason) Copy() types.RVType {
	copied := NewValidationFailureReason()

	copied.StructureVersion = s.StructureVersion
	copied.ValidationID = s.ValidationID.Copy().(types.UInt32)
	copied.Description = s.Description.Copy().(types.String)

	return copied
}

// Equals checks if the given ValidationFailureReason contains the same data as the current ValidationFailureReason
func (s ValidationFailureReason) Equals(o types.RVType) bool {
	other, ok := o.(ValidationFailureReason)
	if !ok {
		return false
	}

	if s.StructureVersion != other.StructureVersion {
		return false
	}

	if !s.ValidationID.Equals(other.ValidationID) {
		return false
	}

	if !s.Description.Equals(other.Description) {
		return false
	}

	return true
}

// CopyRef copies the current value of the ValidationFailureReason
// and returns a pointer to the new copy
func (s ValidationFailureReason) CopyRef() types.RVTypePtr {
	copied := s.Copy().(ValidationFailureReason)
	return &copied
}

// Deref takes a pointer to the ValidationFailureReason
// and dereferences it to the raw value.
// Only useful when working with an instance of RVTypePtr
func (s *ValidationFailureReason) Deref() types.RVType {
	return *s
}

// String returns the string representation of the ValidationFailureReason
func (s ValidationFailureReason) String() string {
	return s.FormatToString(0)
}

// FormatToString pretty-prints the ValidationFailureReason using the provided indentation level
func (s ValidationFailureReason) FormatToString(indentationLevel int) string {
	indentationValues := strings.Repeat("\t", indentationLevel+1)
	indentationEnd := strings.Repeat("\t", indentationLevel)

	var b strings.Builder

	b.WriteString("ValidationFailureReason{\n")
	fmt.Fprintf(&b, "%sValidationID: %s,\n", indentationValues, s.ValidationID)
	fmt.Fprintf(&b, "%sDescription: %s,\n", indentationValues, s.Description)
	fmt.Fprintf(&b, "%s}", indentationEnd)

	return b.String()
}

// NewValidationFailureReason returns a new ValidationFailureReason
func NewValidationFailureReason() ValidationFailureReason {
	return ValidationFailureReason{
		ValidationID: types.NewUInt32(0),
		Description:  types.NewString(""),
	}
}
