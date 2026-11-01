// Package types implements all the types used by the Ubi Account Management protocol
package types

import (
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go/v2/types"
)

// TOSEx is a type within the Ubi Account Management protocol
type TOSEx struct {
	types.Structure
	LocaleCode          types.String
	Contents            types.List[types.String]
	StoringInfoQuestion types.String
}

// WriteTo writes the TOSEx to the given writable
func (s TOSEx) WriteTo(writable types.Writable) {
	contentWritable := writable.CopyNew()

	s.LocaleCode.WriteTo(contentWritable)
	s.Contents.WriteTo(contentWritable)
	s.StoringInfoQuestion.WriteTo(contentWritable)

	content := contentWritable.Bytes()

	s.WriteHeaderTo(writable, uint32(len(content)))

	writable.Write(content)
}

// ExtractFrom extracts the TOSEx from the given readable
func (s *TOSEx) ExtractFrom(readable types.Readable) error {
	if err := s.ExtractHeaderFrom(readable); err != nil {
		return fmt.Errorf("failed to extract TOSEx header. %s", err.Error())
	}

	if err := s.LocaleCode.ExtractFrom(readable); err != nil {
		return fmt.Errorf("failed to extract TOSEx.LocaleCode. %s", err.Error())
	}

	if err := s.Contents.ExtractFrom(readable); err != nil {
		return fmt.Errorf("failed to extract TOSEx.Contents. %s", err.Error())
	}

	if err := s.StoringInfoQuestion.ExtractFrom(readable); err != nil {
		return fmt.Errorf("failed to extract TOSEx.StoringInfoQuestion. %s", err.Error())
	}

	return nil
}

// Copy returns a new copied instance of TOSEx
func (s TOSEx) Copy() types.RVType {
	copied := NewTOSEx()

	copied.StructureVersion = s.StructureVersion
	copied.LocaleCode = s.LocaleCode.Copy().(types.String)
	copied.Contents = s.Contents.Copy().(types.List[types.String])
	copied.StoringInfoQuestion = s.StoringInfoQuestion.Copy().(types.String)

	return copied
}

// Equals checks if the given TOSEx contains the same data as the current TOSEx
func (s TOSEx) Equals(o types.RVType) bool {
	other, ok := o.(TOSEx)
	if !ok {
		return false
	}

	if s.StructureVersion != other.StructureVersion {
		return false
	}

	if !s.LocaleCode.Equals(other.LocaleCode) {
		return false
	}

	if !s.Contents.Equals(other.Contents) {
		return false
	}

	if !s.StoringInfoQuestion.Equals(other.StoringInfoQuestion) {
		return false
	}

	return true
}

// CopyRef copies the current value of the TOSEx
// and returns a pointer to the new copy
func (s TOSEx) CopyRef() types.RVTypePtr {
	copied := s.Copy().(TOSEx)
	return &copied
}

// Deref takes a pointer to the TOSEx
// and dereferences it to the raw value.
// Only useful when working with an instance of RVTypePtr
func (s *TOSEx) Deref() types.RVType {
	return *s
}

// String returns the string representation of the TOSEx
func (s TOSEx) String() string {
	return s.FormatToString(0)
}

// FormatToString pretty-prints the TOSEx using the provided indentation level
func (s TOSEx) FormatToString(indentationLevel int) string {
	indentationValues := strings.Repeat("\t", indentationLevel+1)
	indentationEnd := strings.Repeat("\t", indentationLevel)

	var b strings.Builder

	b.WriteString("TOSEx{\n")
	fmt.Fprintf(&b, "%sLocaleCode: %s,\n", indentationValues, s.LocaleCode)
	fmt.Fprintf(&b, "%sContents: %s,\n", indentationValues, s.Contents)
	fmt.Fprintf(&b, "%sStoringInfoQuestion: %s,\n", indentationValues, s.StoringInfoQuestion)
	fmt.Fprintf(&b, "%s}", indentationEnd)

	return b.String()
}

// NewTOSEx returns a new TOSEx
func NewTOSEx() TOSEx {
	return TOSEx{
		LocaleCode:          types.NewString(""),
		Contents:            types.NewList[types.String](),
		StoringInfoQuestion: types.NewString(""),
	}
}
