// Package types implements the structures used by the Tracking 3 protocol

package types

import (
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go/v2/types"
)

// TrackingInformation is a type within the Tracking 3 protocol
type TrackingInformation struct {
	types.Structure
	Ipn        types.UInt32
	UserID     types.String
	MachineID  types.String
	VisitorID  types.String
	UtsVersion types.String
}

// WriteTo writes the TrackingInformation to the given writable
func (ti TrackingInformation) WriteTo(writable types.Writable) {
	contentWritable := writable.CopyNew()

	ti.Ipn.WriteTo(contentWritable)
	ti.UserID.WriteTo(contentWritable)
	ti.MachineID.WriteTo(contentWritable)
	ti.VisitorID.WriteTo(contentWritable)
	ti.UtsVersion.WriteTo(contentWritable)

	content := contentWritable.Bytes()

	ti.WriteHeaderTo(writable, uint32(len(content)))

	writable.Write(content)
}

// ExtractFrom extracts the TrackingInformation from the given readable
func (ti *TrackingInformation) ExtractFrom(readable types.Readable) error {
	if err := ti.ExtractHeaderFrom(readable); err != nil {
		return fmt.Errorf("failed to extract TrackingInformation header. %s", err.Error())
	}

	if err := ti.Ipn.ExtractFrom(readable); err != nil {
		return fmt.Errorf("failed to extract TrackingInformation.Ipn. %s", err.Error())
	}

	if err := ti.UserID.ExtractFrom(readable); err != nil {
		return fmt.Errorf("failed to extract TrackingInformation.UserID. %s", err.Error())
	}

	if err := ti.MachineID.ExtractFrom(readable); err != nil {
		return fmt.Errorf("failed to extract TrackingInformation.MachineID. %s", err.Error())
	}

	if err := ti.VisitorID.ExtractFrom(readable); err != nil {
		return fmt.Errorf("failed to extract TrackingInformation.VisitorID. %s", err.Error())
	}

	if err := ti.UtsVersion.ExtractFrom(readable); err != nil {
		return fmt.Errorf("failed to extract TrackingInformation.UtsVersion. %s", err.Error())
	}

	return nil
}

// Copy returns a new copied instance of TrackingInformation
func (ti TrackingInformation) Copy() types.RVType {
	copied := NewTrackingInformation()

	copied.StructureVersion = ti.StructureVersion
	copied.Ipn = ti.Ipn.Copy().(types.UInt32)
	copied.UserID = ti.UserID.Copy().(types.String)
	copied.MachineID = ti.MachineID.Copy().(types.String)
	copied.VisitorID = ti.VisitorID.Copy().(types.String)
	copied.UtsVersion = ti.UtsVersion.Copy().(types.String)

	return copied
}

// Equals checks if the given TrackingInformation contains the same data as the current TrackingInformation
func (ti TrackingInformation) Equals(o types.RVType) bool {
	other, ok := o.(TrackingInformation)
	if !ok {
		return false
	}

	if ti.StructureVersion != other.StructureVersion {
		return false
	}

	if !(ti.Ipn.Equals(other.Ipn)) {
		return false
	}

	if !(ti.UserID.Equals(other.UserID)) {
		return false
	}

	if !(ti.MachineID.Equals(other.MachineID)) {
		return false
	}

	if !(ti.VisitorID.Equals(other.VisitorID)) {
		return false
	}

	if !(ti.UtsVersion.Equals(other.UtsVersion)) {
		return false
	}

	return true
}

// CopyRef copies the current value of the TrackingInformation
// and returns a pointer to the new copy
func (ti TrackingInformation) CopyRef() types.RVTypePtr {
	copied := ti.Copy().(TrackingInformation)
	return &copied
}

// Deref takes a pointer to the TrackingInformation
// and dereferences it to the raw value.
// Only useful when working with an instance of RVTypePtr
func (ti *TrackingInformation) Deref() types.RVType {
	return *ti
}

// String returns the string representation of the TrackingInformation
func (ti TrackingInformation) String() string {
	return ti.FormatToString(0)
}

// FormatToString pretty-prints the TrackingInformation using the provided indentation level
func (ti TrackingInformation) FormatToString(indentationLevel int) string {
	indentationValues := strings.Repeat("\t", indentationLevel+1)
	indentationEnd := strings.Repeat("\t", indentationLevel)

	var b strings.Builder

	b.WriteString("TrackingInformation{\n")
	fmt.Fprintf(&b, "%sIpn: %s,\n", indentationValues, ti.Ipn)
	fmt.Fprintf(&b, "%sUserID: %s,\n", indentationValues, ti.UserID)
	fmt.Fprintf(&b, "%sMachineID: %s,\n", indentationValues, ti.MachineID)
	fmt.Fprintf(&b, "%sVisitorID: %s,\n", indentationValues, ti.VisitorID)
	fmt.Fprintf(&b, "%sUtsVersion: %s,\n", indentationValues, ti.UtsVersion)
	fmt.Fprintf(&b, "%s}", indentationEnd)

	return b.String()
}

// NewTrackingInformation returns a new TrackingInformation
func NewTrackingInformation() TrackingInformation {
	return TrackingInformation{
		Ipn:        types.NewUInt32(0),
		UserID:     types.NewString(""),
		MachineID:  types.NewString(""),
		VisitorID:  types.NewString(""),
		UtsVersion: types.NewString(""),
	}
}
