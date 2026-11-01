// Package types implements all the types used by the Privileges protocol
package types

import (
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go/v2/types"
)

// PrivilegeGroup is a type within the Privileges protocol
type PrivilegeGroup struct {
	types.Structure
	Description types.String
	Privileges  types.List[Privilege]
}

// WriteTo writes the PrivilegeGroup to the given writable
func (s PrivilegeGroup) WriteTo(writable types.Writable) {
	contentWritable := writable.CopyNew()

	s.Description.WriteTo(contentWritable)
	s.Privileges.WriteTo(contentWritable)

	content := contentWritable.Bytes()

	s.WriteHeaderTo(writable, uint32(len(content)))

	writable.Write(content)
}

// ExtractFrom extracts the PrivilegeGroup from the given readable
func (s *PrivilegeGroup) ExtractFrom(readable types.Readable) error {
	if err := s.ExtractHeaderFrom(readable); err != nil {
		return fmt.Errorf("failed to extract PrivilegeGroup header. %s", err.Error())
	}

	if err := s.Description.ExtractFrom(readable); err != nil {
		return fmt.Errorf("failed to extract PrivilegeGroup.Description. %s", err.Error())
	}

	if err := s.Privileges.ExtractFrom(readable); err != nil {
		return fmt.Errorf("failed to extract PrivilegeGroup.Privileges. %s", err.Error())
	}

	return nil
}

// Copy returns a new copied instance of PrivilegeGroup
func (s PrivilegeGroup) Copy() types.RVType {
	copied := NewPrivilegeGroup()

	copied.StructureVersion = s.StructureVersion
	copied.Description = s.Description.Copy().(types.String)
	copied.Privileges = s.Privileges.Copy().(types.List[Privilege])

	return copied
}

// Equals checks if the given PrivilegeGroup contains the same data as the current PrivilegeGroup
func (s PrivilegeGroup) Equals(o types.RVType) bool {
	other, ok := o.(PrivilegeGroup)
	if !ok {
		return false
	}

	if s.StructureVersion != other.StructureVersion {
		return false
	}

	if !s.Description.Equals(other.Description) {
		return false
	}

	if !s.Privileges.Equals(other.Privileges) {
		return false
	}

	return true
}

// CopyRef copies the current value of the PrivilegeGroup
// and returns a pointer to the new copy
func (s PrivilegeGroup) CopyRef() types.RVTypePtr {
	copied := s.Copy().(PrivilegeGroup)
	return &copied
}

// Deref takes a pointer to the PrivilegeGroup
// and dereferences it to the raw value.
// Only useful when working with an instance of RVTypePtr
func (s *PrivilegeGroup) Deref() types.RVType {
	return *s
}

// String returns the string representation of the PrivilegeGroup
func (s PrivilegeGroup) String() string {
	return s.FormatToString(0)
}

// FormatToString pretty-prints the PrivilegeGroup using the provided indentation level
func (s PrivilegeGroup) FormatToString(indentationLevel int) string {
	indentationValues := strings.Repeat("\t", indentationLevel+1)
	indentationEnd := strings.Repeat("\t", indentationLevel)

	var b strings.Builder

	b.WriteString("PrivilegeGroup{\n")
	fmt.Fprintf(&b, "%sDescription: %s,\n", indentationValues, s.Description)
	fmt.Fprintf(&b, "%sPrivileges: %s,\n", indentationValues, s.Privileges)
	fmt.Fprintf(&b, "%s}", indentationEnd)

	return b.String()
}

// NewPrivilegeGroup returns a new PrivilegeGroup
func NewPrivilegeGroup() PrivilegeGroup {
	return PrivilegeGroup{
		Description: types.NewString(""),
		Privileges:  types.NewList[Privilege](),
	}
}
