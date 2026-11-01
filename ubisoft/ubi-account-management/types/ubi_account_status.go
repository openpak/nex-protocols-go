// Package types implements all the types used by the Ubi Account Management protocol
package types

import (
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go/v2/types"
)

// UbiAccountStatus is a type within the Ubi Account Management protocol
type UbiAccountStatus struct {
	types.Structure
	BasicStatus                 types.UInt32
	MissingRequiredInformations types.Bool
	RecoveringPassword          types.Bool
	PendingDeactivation         types.Bool
}

// WriteTo writes the UbiAccountStatus to the given writable
func (s UbiAccountStatus) WriteTo(writable types.Writable) {
	contentWritable := writable.CopyNew()

	s.BasicStatus.WriteTo(contentWritable)
	s.MissingRequiredInformations.WriteTo(contentWritable)
	s.RecoveringPassword.WriteTo(contentWritable)
	s.PendingDeactivation.WriteTo(contentWritable)

	content := contentWritable.Bytes()

	s.WriteHeaderTo(writable, uint32(len(content)))

	writable.Write(content)
}

// ExtractFrom extracts the UbiAccountStatus from the given readable
func (s *UbiAccountStatus) ExtractFrom(readable types.Readable) error {
	if err := s.ExtractHeaderFrom(readable); err != nil {
		return fmt.Errorf("failed to extract UbiAccountStatus header. %s", err.Error())
	}

	if err := s.BasicStatus.ExtractFrom(readable); err != nil {
		return fmt.Errorf("failed to extract UbiAccountStatus.BasicStatus. %s", err.Error())
	}

	if err := s.MissingRequiredInformations.ExtractFrom(readable); err != nil {
		return fmt.Errorf("failed to extract UbiAccountStatus.MissingRequiredInformations. %s", err.Error())
	}

	if err := s.RecoveringPassword.ExtractFrom(readable); err != nil {
		return fmt.Errorf("failed to extract UbiAccountStatus.RecoveringPassword. %s", err.Error())
	}

	if err := s.PendingDeactivation.ExtractFrom(readable); err != nil {
		return fmt.Errorf("failed to extract UbiAccountStatus.PendingDeactivation. %s", err.Error())
	}

	return nil
}

// Copy returns a new copied instance of UbiAccountStatus
func (s UbiAccountStatus) Copy() types.RVType {
	copied := NewUbiAccountStatus()

	copied.StructureVersion = s.StructureVersion
	copied.BasicStatus = s.BasicStatus.Copy().(types.UInt32)
	copied.MissingRequiredInformations = s.MissingRequiredInformations.Copy().(types.Bool)
	copied.RecoveringPassword = s.RecoveringPassword.Copy().(types.Bool)
	copied.PendingDeactivation = s.PendingDeactivation.Copy().(types.Bool)

	return copied
}

// Equals checks if the given UbiAccountStatus contains the same data as the current UbiAccountStatus
func (s UbiAccountStatus) Equals(o types.RVType) bool {
	other, ok := o.(UbiAccountStatus)
	if !ok {
		return false
	}

	if s.StructureVersion != other.StructureVersion {
		return false
	}

	if !s.BasicStatus.Equals(other.BasicStatus) {
		return false
	}

	if !s.MissingRequiredInformations.Equals(other.MissingRequiredInformations) {
		return false
	}

	if !s.RecoveringPassword.Equals(other.RecoveringPassword) {
		return false
	}

	if !s.PendingDeactivation.Equals(other.PendingDeactivation) {
		return false
	}

	return true
}

// CopyRef copies the current value of the UbiAccountStatus
// and returns a pointer to the new copy
func (s UbiAccountStatus) CopyRef() types.RVTypePtr {
	copied := s.Copy().(UbiAccountStatus)
	return &copied
}

// Deref takes a pointer to the UbiAccountStatus
// and dereferences it to the raw value.
// Only useful when working with an instance of RVTypePtr
func (s *UbiAccountStatus) Deref() types.RVType {
	return *s
}

// String returns the string representation of the UbiAccountStatus
func (s UbiAccountStatus) String() string {
	return s.FormatToString(0)
}

// FormatToString pretty-prints the UbiAccountStatus using the provided indentation level
func (s UbiAccountStatus) FormatToString(indentationLevel int) string {
	indentationValues := strings.Repeat("\t", indentationLevel+1)
	indentationEnd := strings.Repeat("\t", indentationLevel)

	var b strings.Builder

	b.WriteString("UbiAccountStatus{\n")
	fmt.Fprintf(&b, "%sBasicStatus: %s,\n", indentationValues, s.BasicStatus)
	fmt.Fprintf(&b, "%sMissingRequiredInformations: %s,\n", indentationValues, s.MissingRequiredInformations)
	fmt.Fprintf(&b, "%sRecoveringPassword: %s,\n", indentationValues, s.RecoveringPassword)
	fmt.Fprintf(&b, "%sPendingDeactivation: %s,\n", indentationValues, s.PendingDeactivation)
	fmt.Fprintf(&b, "%s}", indentationEnd)

	return b.String()
}

// NewUbiAccountStatus returns a new UbiAccountStatus
func NewUbiAccountStatus() UbiAccountStatus {
	return UbiAccountStatus{
		BasicStatus:                 types.NewUInt32(0),
		MissingRequiredInformations: types.NewBool(false),
		RecoveringPassword:          types.NewBool(false),
		PendingDeactivation:         types.NewBool(false),
	}
}
