// Package types implements all the types used by the Ubi Account Management protocol
package types

import (
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go/v2/types"
)

// Country is a type within the Ubi Account Management protocol
type Country struct {
	types.Structure
	Code types.String
	Name types.String
}

// WriteTo writes the Country to the given writable
func (s Country) WriteTo(writable types.Writable) {
	contentWritable := writable.CopyNew()

	s.Code.WriteTo(contentWritable)
	s.Name.WriteTo(contentWritable)

	content := contentWritable.Bytes()

	s.WriteHeaderTo(writable, uint32(len(content)))

	writable.Write(content)
}

// ExtractFrom extracts the Country from the given readable
func (s *Country) ExtractFrom(readable types.Readable) error {
	if err := s.ExtractHeaderFrom(readable); err != nil {
		return fmt.Errorf("failed to extract Country header. %s", err.Error())
	}

	if err := s.Code.ExtractFrom(readable); err != nil {
		return fmt.Errorf("failed to extract Country.Code. %s", err.Error())
	}

	if err := s.Name.ExtractFrom(readable); err != nil {
		return fmt.Errorf("failed to extract Country.Name. %s", err.Error())
	}

	return nil
}

// Copy returns a new copied instance of Country
func (s Country) Copy() types.RVType {
	copied := NewCountry()

	copied.StructureVersion = s.StructureVersion
	copied.Code = s.Code.Copy().(types.String)
	copied.Name = s.Name.Copy().(types.String)

	return copied
}

// Equals checks if the given Country contains the same data as the current Country
func (s Country) Equals(o types.RVType) bool {
	other, ok := o.(Country)
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

	return true
}

// CopyRef copies the current value of the Country
// and returns a pointer to the new copy
func (s Country) CopyRef() types.RVTypePtr {
	copied := s.Copy().(Country)
	return &copied
}

// Deref takes a pointer to the Country
// and dereferences it to the raw value.
// Only useful when working with an instance of RVTypePtr
func (s *Country) Deref() types.RVType {
	return *s
}

// String returns the string representation of the Country
func (s Country) String() string {
	return s.FormatToString(0)
}

// FormatToString pretty-prints the Country using the provided indentation level
func (s Country) FormatToString(indentationLevel int) string {
	indentationValues := strings.Repeat("\t", indentationLevel+1)
	indentationEnd := strings.Repeat("\t", indentationLevel)

	var b strings.Builder

	b.WriteString("Country{\n")
	fmt.Fprintf(&b, "%sCode: %s,\n", indentationValues, s.Code)
	fmt.Fprintf(&b, "%sName: %s,\n", indentationValues, s.Name)
	fmt.Fprintf(&b, "%s}", indentationEnd)

	return b.String()
}

// NewCountry returns a new Country
func NewCountry() Country {
	return Country{
		Code: types.NewString(""),
		Name: types.NewString(""),
	}
}
