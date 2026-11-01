package types

import (
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go/v2/types"
)

// UserSlotCount is a type within the User Storage protocol
type UserSlotCount struct {
	types.Structure
	Pid        types.UInt32
	TypeID     types.UInt32
	UsedSlots  types.UInt32
	TotalSlots types.UInt32
}

// WriteTo writes the UserSlotCount to the given writable
func (usc UserSlotCount) WriteTo(writable types.Writable) {
	contentWritable := writable.CopyNew()

	usc.Pid.WriteTo(contentWritable)
	usc.TypeID.WriteTo(contentWritable)
	usc.UsedSlots.WriteTo(contentWritable)
	usc.TotalSlots.WriteTo(contentWritable)

	content := contentWritable.Bytes()

	usc.WriteHeaderTo(writable, uint32(len(content)))

	writable.Write(content)
}

// ExtractFrom extracts the UserSlotCount from the given readable
func (usc *UserSlotCount) ExtractFrom(readable types.Readable) error {
	if err := usc.ExtractHeaderFrom(readable); err != nil {
		return fmt.Errorf("failed to extract UserSlotCount header. %s", err.Error())
	}

	if err := usc.Pid.ExtractFrom(readable); err != nil {
		return fmt.Errorf("failed to extract UserSlotCount.Pid. %s", err.Error())
	}

	if err := usc.TypeID.ExtractFrom(readable); err != nil {
		return fmt.Errorf("failed to extract UserSlotCount.TypeID. %s", err.Error())
	}

	if err := usc.UsedSlots.ExtractFrom(readable); err != nil {
		return fmt.Errorf("failed to extract UserSlotCount.UsedSlots. %s", err.Error())
	}

	if err := usc.TotalSlots.ExtractFrom(readable); err != nil {
		return fmt.Errorf("failed to extract UserSlotCount.TotalSlots. %s", err.Error())
	}

	return nil
}

// Copy returns a new copied instance of UserSlotCount
func (usc UserSlotCount) Copy() types.RVType {
	copied := NewUserSlotCount()

	copied.StructureVersion = usc.StructureVersion
	copied.Pid = usc.Pid.Copy().(types.UInt32)
	copied.TypeID = usc.TypeID.Copy().(types.UInt32)
	copied.UsedSlots = usc.UsedSlots.Copy().(types.UInt32)
	copied.TotalSlots = usc.TotalSlots.Copy().(types.UInt32)

	return copied
}

// Equals checks if the given UserSlotCount contains the same data as the current UserSlotCount
func (usc UserSlotCount) Equals(o types.RVType) bool {
	other, ok := o.(UserSlotCount)
	if !ok {
		return false
	}

	if usc.StructureVersion != other.StructureVersion {
		return false
	}

	if !(usc.Pid.Equals(other.Pid)) {
		return false
	}

	if !(usc.TypeID.Equals(other.TypeID)) {
		return false
	}

	if !(usc.UsedSlots.Equals(other.UsedSlots)) {
		return false
	}

	if !(usc.TotalSlots.Equals(other.TotalSlots)) {
		return false
	}

	return true
}

// CopyRef copies the current value of the UserSlotCount
// and returns a pointer to the new copy
func (usc UserSlotCount) CopyRef() types.RVTypePtr {
	copied := usc.Copy().(UserSlotCount)
	return &copied
}

// Deref takes a pointer to the UserSlotCount
// and dereferences it to the raw value.
// Only useful when working with an instance of RVTypePtr
func (usc *UserSlotCount) Deref() types.RVType {
	return *usc
}

// String returns the string representation of the UserSlotCount
func (usc UserSlotCount) String() string {
	return usc.FormatToString(0)
}

// FormatToString pretty-prints the UserSlotCount using the provided indentation level
func (usc UserSlotCount) FormatToString(indentationLevel int) string {
	indentationValues := strings.Repeat("\t", indentationLevel+1)
	indentationEnd := strings.Repeat("\t", indentationLevel)

	var b strings.Builder

	b.WriteString("UserSlotCount{\n")
	fmt.Fprintf(&b, "%sPid: %s,\n", indentationValues, usc.Pid)
	fmt.Fprintf(&b, "%sTypeID: %s,\n", indentationValues, usc.TypeID)
	fmt.Fprintf(&b, "%sUsedSlots: %s,\n", indentationValues, usc.UsedSlots)
	fmt.Fprintf(&b, "%sTotalSlots: %s,\n", indentationValues, usc.TotalSlots)
	fmt.Fprintf(&b, "%s}", indentationEnd)

	return b.String()
}

// NewUserSlotCount returns a new UserSlotCount
func NewUserSlotCount() UserSlotCount {
	return UserSlotCount{
		Pid:        types.NewUInt32(0),
		TypeID:     types.NewUInt32(0),
		UsedSlots:  types.NewUInt32(0),
		TotalSlots: types.NewUInt32(0),
	}
}
