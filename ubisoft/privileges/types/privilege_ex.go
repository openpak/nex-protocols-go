// Package types implements all the types used by the Privileges protocol
package types

import (
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go/v2/types"
)

// PrivilegeEx is a type within the Privileges protocol
type PrivilegeEx struct {
	types.Structure
	ID          types.UInt32
	Description types.String
	Duration    types.Int32
}

// WriteTo writes the PrivilegeEx to the given writable
func (s PrivilegeEx) WriteTo(writable types.Writable) {
	contentWritable := writable.CopyNew()

	s.ID.WriteTo(contentWritable)
	s.Description.WriteTo(contentWritable)
	s.Duration.WriteTo(contentWritable)

	content := contentWritable.Bytes()

	s.WriteHeaderTo(writable, uint32(len(content)))

	writable.Write(content)
}

// ExtractFrom extracts the PrivilegeEx from the given readable
func (s *PrivilegeEx) ExtractFrom(readable types.Readable) error {
	if err := s.ExtractHeaderFrom(readable); err != nil {
		return fmt.Errorf("failed to extract PrivilegeEx header. %s", err.Error())
	}

	if err := s.ID.ExtractFrom(readable); err != nil {
		return fmt.Errorf("failed to extract PrivilegeEx.ID. %s", err.Error())
	}

	if err := s.Description.ExtractFrom(readable); err != nil {
		return fmt.Errorf("failed to extract PrivilegeEx.Description. %s", err.Error())
	}

	if err := s.Duration.ExtractFrom(readable); err != nil {
		return fmt.Errorf("failed to extract PrivilegeEx.Duration. %s", err.Error())
	}

	return nil
}

// Copy returns a new copied instance of PrivilegeEx
func (s PrivilegeEx) Copy() types.RVType {
	copied := NewPrivilegeEx()

	copied.StructureVersion = s.StructureVersion
	copied.ID = s.ID.Copy().(types.UInt32)
	copied.Description = s.Description.Copy().(types.String)
	copied.Duration = s.Duration.Copy().(types.Int32)

	return copied
}

// Equals checks if the given PrivilegeEx contains the same data as the current PrivilegeEx
func (s PrivilegeEx) Equals(o types.RVType) bool {
	other, ok := o.(PrivilegeEx)
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

	if !s.Duration.Equals(other.Duration) {
		return false
	}

	return true
}

// CopyRef copies the current value of the PrivilegeEx
// and returns a pointer to the new copy
func (s PrivilegeEx) CopyRef() types.RVTypePtr {
	copied := s.Copy().(PrivilegeEx)
	return &copied
}

// Deref takes a pointer to the PrivilegeEx
// and dereferences it to the raw value.
// Only useful when working with an instance of RVTypePtr
func (s *PrivilegeEx) Deref() types.RVType {
	return *s
}

// String returns the string representation of the PrivilegeEx
func (s PrivilegeEx) String() string {
	return s.FormatToString(0)
}

// FormatToString pretty-prints the PrivilegeEx using the provided indentation level
func (s PrivilegeEx) FormatToString(indentationLevel int) string {
	indentationValues := strings.Repeat("\t", indentationLevel+1)
	indentationEnd := strings.Repeat("\t", indentationLevel)

	var b strings.Builder

	b.WriteString("PrivilegeEx{\n")
	fmt.Fprintf(&b, "%sID: %s,\n", indentationValues, s.ID)
	fmt.Fprintf(&b, "%sDescription: %s,\n", indentationValues, s.Description)
	fmt.Fprintf(&b, "%sDuration: %s,\n", indentationValues, s.Duration)
	fmt.Fprintf(&b, "%s}", indentationEnd)

	return b.String()
}

// NewPrivilegeEx returns a new PrivilegeEx
func NewPrivilegeEx() PrivilegeEx {
	return PrivilegeEx{
		ID:          types.NewUInt32(0),
		Description: types.NewString(""),
		Duration:    types.NewInt32(0),
	}
}
