// Package types implements all the types used by the Ubi Account Management protocol
package types

import (
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go/v2/types"
)

// UsernameValidation is a type within the Ubi Account Management protocol
type UsernameValidation struct {
	types.Structure
	Suggestions types.List[types.String]
	Reasons     types.List[ValidationFailureReason]
}

// WriteTo writes the UsernameValidation to the given writable
func (s UsernameValidation) WriteTo(writable types.Writable) {
	contentWritable := writable.CopyNew()

	s.Suggestions.WriteTo(contentWritable)
	s.Reasons.WriteTo(contentWritable)

	content := contentWritable.Bytes()

	s.WriteHeaderTo(writable, uint32(len(content)))

	writable.Write(content)
}

// ExtractFrom extracts the UsernameValidation from the given readable
func (s *UsernameValidation) ExtractFrom(readable types.Readable) error {
	if err := s.ExtractHeaderFrom(readable); err != nil {
		return fmt.Errorf("failed to extract UsernameValidation header. %s", err.Error())
	}

	if err := s.Suggestions.ExtractFrom(readable); err != nil {
		return fmt.Errorf("failed to extract UsernameValidation.Suggestions. %s", err.Error())
	}

	if err := s.Reasons.ExtractFrom(readable); err != nil {
		return fmt.Errorf("failed to extract UsernameValidation.Reasons. %s", err.Error())
	}

	return nil
}

// Copy returns a new copied instance of UsernameValidation
func (s UsernameValidation) Copy() types.RVType {
	copied := NewUsernameValidation()

	copied.StructureVersion = s.StructureVersion
	copied.Suggestions = s.Suggestions.Copy().(types.List[types.String])
	copied.Reasons = s.Reasons.Copy().(types.List[ValidationFailureReason])

	return copied
}

// Equals checks if the given UsernameValidation contains the same data as the current UsernameValidation
func (s UsernameValidation) Equals(o types.RVType) bool {
	other, ok := o.(UsernameValidation)
	if !ok {
		return false
	}

	if s.StructureVersion != other.StructureVersion {
		return false
	}

	if !s.Suggestions.Equals(other.Suggestions) {
		return false
	}

	if !s.Reasons.Equals(other.Reasons) {
		return false
	}

	return true
}

// CopyRef copies the current value of the UsernameValidation
// and returns a pointer to the new copy
func (s UsernameValidation) CopyRef() types.RVTypePtr {
	copied := s.Copy().(UsernameValidation)
	return &copied
}

// Deref takes a pointer to the UsernameValidation
// and dereferences it to the raw value.
// Only useful when working with an instance of RVTypePtr
func (s *UsernameValidation) Deref() types.RVType {
	return *s
}

// String returns the string representation of the UsernameValidation
func (s UsernameValidation) String() string {
	return s.FormatToString(0)
}

// FormatToString pretty-prints the UsernameValidation using the provided indentation level
func (s UsernameValidation) FormatToString(indentationLevel int) string {
	indentationValues := strings.Repeat("\t", indentationLevel+1)
	indentationEnd := strings.Repeat("\t", indentationLevel)

	var b strings.Builder

	b.WriteString("UsernameValidation{\n")
	fmt.Fprintf(&b, "%sSuggestions: %s,\n", indentationValues, s.Suggestions)
	fmt.Fprintf(&b, "%sReasons: %s,\n", indentationValues, s.Reasons)
	fmt.Fprintf(&b, "%s}", indentationEnd)

	return b.String()
}

// NewUsernameValidation returns a new UsernameValidation
func NewUsernameValidation() UsernameValidation {
	return UsernameValidation{
		Suggestions: types.NewList[types.String](),
		Reasons:     types.NewList[ValidationFailureReason](),
	}
}
