// Package types implements all the types used by the Privileges protocol
package types

import (
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go/v2/types"
)

// Privilege is a type within the Privileges protocol
type Privilege struct {
	types.Structure
	ID          types.UInt32
	Description types.String
}

// WriteTo writes the Privilege to the given writable
func (s Privilege) WriteTo(writable types.Writable) {
	contentWritable := writable.CopyNew()

	s.ID.WriteTo(contentWritable)
	s.Description.WriteTo(contentWritable)

	content := contentWritable.Bytes()

	s.WriteHeaderTo(writable, uint32(len(content)))

	writable.Write(content)
}

// ExtractFrom extracts the Privilege from the given readable
func (s *Privilege) ExtractFrom(readable types.Readable) error {
	if err := s.ExtractHeaderFrom(readable); err != nil {
		return fmt.Errorf("failed to extract Privilege header. %s", err.Error())
	}

	if err := s.ID.ExtractFrom(readable); err != nil {
		return fmt.Errorf("failed to extract Privilege.ID. %s", err.Error())
	}

	if err := s.Description.ExtractFrom(readable); err != nil {
		return fmt.Errorf("failed to extract Privilege.Description. %s", err.Error())
	}

	return nil
}

// Copy returns a new copied instance of Privilege
func (s Privilege) Copy() types.RVType {
	copied := NewPrivilege()

	copied.StructureVersion = s.StructureVersion
	copied.ID = s.ID.Copy().(types.UInt32)
	copied.Description = s.Description.Copy().(types.String)

	return copied
}

// Equals checks if the given Privilege contains the same data as the current Privilege
func (s Privilege) Equals(o types.RVType) bool {
	other, ok := o.(Privilege)
	if !ok {
		return false
	}

	if s.StructureVersion != other.StructureVersion {
		return false
	}

	if !s.ID.Equals(other.ID) {
		return false
	}

	if !s.Description.Equals(other.Description) {
		return false
	}

	return true
}

// CopyRef copies the current value of the Privilege
// and returns a pointer to the new copy
func (s Privilege) CopyRef() types.RVTypePtr {
	copied := s.Copy().(Privilege)
	return &copied
}

// Deref takes a pointer to the Privilege
// and dereferences it to the raw value.
// Only useful when working with an instance of RVTypePtr
func (s *Privilege) Deref() types.RVType {
	return *s
}

// String returns the string representation of the Privilege
func (s Privilege) String() string {
	return s.FormatToString(0)
}

// FormatToString pretty-prints the Privilege using the provided indentation level
func (s Privilege) FormatToString(indentationLevel int) string {
	indentationValues := strings.Repeat("\t", indentationLevel+1)
	indentationEnd := strings.Repeat("\t", indentationLevel)

	var b strings.Builder

	b.WriteString("Privilege{\n")
	fmt.Fprintf(&b, "%sID: %s,\n", indentationValues, s.ID)
	fmt.Fprintf(&b, "%sDescription: %s,\n", indentationValues, s.Description)
	fmt.Fprintf(&b, "%s}", indentationEnd)

	return b.String()
}

// NewPrivilege returns a new Privilege
func NewPrivilege() Privilege {
	return Privilege{
		ID:          types.NewUInt32(0),
		Description: types.NewString(""),
	}
}
