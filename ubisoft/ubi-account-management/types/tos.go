// Package types implements all the types used by the Ubi Account Management protocol
package types

import (
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go/v2/types"
)

// TOS is a type within the Ubi Account Management protocol
type TOS struct {
	types.Structure
	LocaleCode          types.String
	Content             types.String
	StoringInfoQuestion types.String
}

// WriteTo writes the TOS to the given writable
func (s TOS) WriteTo(writable types.Writable) {
	contentWritable := writable.CopyNew()

	s.LocaleCode.WriteTo(contentWritable)
	s.Content.WriteTo(contentWritable)
	s.StoringInfoQuestion.WriteTo(contentWritable)

	content := contentWritable.Bytes()

	s.WriteHeaderTo(writable, uint32(len(content)))

	writable.Write(content)
}

// ExtractFrom extracts the TOS from the given readable
func (s *TOS) ExtractFrom(readable types.Readable) error {
	if err := s.ExtractHeaderFrom(readable); err != nil {
		return fmt.Errorf("failed to extract TOS header. %s", err.Error())
	}

	if err := s.LocaleCode.ExtractFrom(readable); err != nil {
		return fmt.Errorf("failed to extract TOS.LocaleCode. %s", err.Error())
	}

	if err := s.Content.ExtractFrom(readable); err != nil {
		return fmt.Errorf("failed to extract TOS.Content. %s", err.Error())
	}

	if err := s.StoringInfoQuestion.ExtractFrom(readable); err != nil {
		return fmt.Errorf("failed to extract TOS.StoringInfoQuestion. %s", err.Error())
	}

	return nil
}

// Copy returns a new copied instance of TOS
func (s TOS) Copy() types.RVType {
	copied := NewTOS()

	copied.StructureVersion = s.StructureVersion
	copied.LocaleCode = s.LocaleCode.Copy().(types.String)
	copied.Content = s.Content.Copy().(types.String)
	copied.StoringInfoQuestion = s.StoringInfoQuestion.Copy().(types.String)

	return copied
}

// Equals checks if the given TOS contains the same data as the current TOS
func (s TOS) Equals(o types.RVType) bool {
	other, ok := o.(TOS)
	if !ok {
		return false
	}

	if s.StructureVersion != other.StructureVersion {
		return false
	}

	if !s.LocaleCode.Equals(other.LocaleCode) {
		return false
	}

	if !s.Content.Equals(other.Content) {
		return false
	}

	if !s.StoringInfoQuestion.Equals(other.StoringInfoQuestion) {
		return false
	}

	return true
}

// CopyRef copies the current value of the TOS
// and returns a pointer to the new copy
func (s TOS) CopyRef() types.RVTypePtr {
	copied := s.Copy().(TOS)
	return &copied
}

// Deref takes a pointer to the TOS
// and dereferences it to the raw value.
// Only useful when working with an instance of RVTypePtr
func (s *TOS) Deref() types.RVType {
	return *s
}

// String returns the string representation of the TOS
func (s TOS) String() string {
	return s.FormatToString(0)
}

// FormatToString pretty-prints the TOS using the provided indentation level
func (s TOS) FormatToString(indentationLevel int) string {
	indentationValues := strings.Repeat("\t", indentationLevel+1)
	indentationEnd := strings.Repeat("\t", indentationLevel)

	var b strings.Builder

	b.WriteString("TOS{\n")
	fmt.Fprintf(&b, "%sLocaleCode: %s,\n", indentationValues, s.LocaleCode)
	fmt.Fprintf(&b, "%sContent: %s,\n", indentationValues, s.Content)
	fmt.Fprintf(&b, "%sStoringInfoQuestion: %s,\n", indentationValues, s.StoringInfoQuestion)
	fmt.Fprintf(&b, "%s}", indentationEnd)

	return b.String()
}

// NewTOS returns a new TOS
func NewTOS() TOS {
	return TOS{
		LocaleCode:          types.NewString(""),
		Content:             types.NewString(""),
		StoringInfoQuestion: types.NewString(""),
	}
}
