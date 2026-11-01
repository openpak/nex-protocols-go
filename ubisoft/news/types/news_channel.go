// Package types implements the structures used by the News protocol

package types

import (
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go/v2/types"
)

// NewsChannel is a type within the News protocol
type NewsChannel struct {
	types.Structure
	ID             types.UInt32
	OwnerPID       types.UInt32
	Name           types.String
	Description    types.String
	CreationTime   types.DateTime
	ExpirationTime types.DateTime
	Type           types.String
	Locale         types.String
	Subscribable   types.Bool
}

// WriteTo writes the NewsChannel to the given writable
func (nc NewsChannel) WriteTo(writable types.Writable) {
	contentWritable := writable.CopyNew()

	nc.ID.WriteTo(contentWritable)
	nc.OwnerPID.WriteTo(contentWritable)
	nc.Name.WriteTo(contentWritable)
	nc.Description.WriteTo(contentWritable)
	nc.CreationTime.WriteTo(contentWritable)
	nc.ExpirationTime.WriteTo(contentWritable)
	nc.Type.WriteTo(contentWritable)
	nc.Locale.WriteTo(contentWritable)
	nc.Subscribable.WriteTo(contentWritable)

	content := contentWritable.Bytes()

	nc.WriteHeaderTo(writable, uint32(len(content)))

	writable.Write(content)
}

// ExtractFrom extracts the NewsChannel from the given readable
func (nc *NewsChannel) ExtractFrom(readable types.Readable) error {
	if err := nc.ExtractHeaderFrom(readable); err != nil {
		return fmt.Errorf("failed to extract NewsChannel header. %s", err.Error())
	}

	if err := nc.ID.ExtractFrom(readable); err != nil {
		return fmt.Errorf("failed to extract NewsChannel.ID. %s", err.Error())
	}

	if err := nc.OwnerPID.ExtractFrom(readable); err != nil {
		return fmt.Errorf("failed to extract NewsChannel.OwnerPID. %s", err.Error())
	}

	if err := nc.Name.ExtractFrom(readable); err != nil {
		return fmt.Errorf("failed to extract NewsChannel.Name. %s", err.Error())
	}

	if err := nc.Description.ExtractFrom(readable); err != nil {
		return fmt.Errorf("failed to extract NewsChannel.Description. %s", err.Error())
	}

	if err := nc.CreationTime.ExtractFrom(readable); err != nil {
		return fmt.Errorf("failed to extract NewsChannel.CreationTime. %s", err.Error())
	}

	if err := nc.ExpirationTime.ExtractFrom(readable); err != nil {
		return fmt.Errorf("failed to extract NewsChannel.ExpirationTime. %s", err.Error())
	}

	if err := nc.Type.ExtractFrom(readable); err != nil {
		return fmt.Errorf("failed to extract NewsChannel.Type. %s", err.Error())
	}

	if err := nc.Locale.ExtractFrom(readable); err != nil {
		return fmt.Errorf("failed to extract NewsChannel.Locale. %s", err.Error())
	}

	if err := nc.Subscribable.ExtractFrom(readable); err != nil {
		return fmt.Errorf("failed to extract NewsChannel.Subscribable. %s", err.Error())
	}

	return nil
}

// Copy returns a new copied instance of NewsChannel
func (nc NewsChannel) Copy() types.RVType {
	copied := NewNewsChannel()

	copied.StructureVersion = nc.StructureVersion
	copied.ID = nc.ID.Copy().(types.UInt32)
	copied.OwnerPID = nc.OwnerPID.Copy().(types.UInt32)
	copied.Name = nc.Name.Copy().(types.String)
	copied.Description = nc.Description.Copy().(types.String)
	copied.CreationTime = nc.CreationTime.Copy().(types.DateTime)
	copied.ExpirationTime = nc.ExpirationTime.Copy().(types.DateTime)
	copied.Type = nc.Type.Copy().(types.String)
	copied.Locale = nc.Locale.Copy().(types.String)
	copied.Subscribable = nc.Subscribable.Copy().(types.Bool)

	return copied
}

// Equals checks if the given NewsChannel contains the same data as the current NewsChannel
func (nc NewsChannel) Equals(o types.RVType) bool {
	other, ok := o.(NewsChannel)
	if !ok {
		return false
	}

	if nc.StructureVersion != other.StructureVersion {
		return false
	}

	if !(nc.ID.Equals(other.ID)) {
		return false
	}

	if !(nc.OwnerPID.Equals(other.OwnerPID)) {
		return false
	}

	if !(nc.Name.Equals(other.Name)) {
		return false
	}

	if !(nc.Description.Equals(other.Description)) {
		return false
	}

	if !(nc.CreationTime.Equals(other.CreationTime)) {
		return false
	}

	if !(nc.ExpirationTime.Equals(other.ExpirationTime)) {
		return false
	}

	if !(nc.Type.Equals(other.Type)) {
		return false
	}

	if !(nc.Locale.Equals(other.Locale)) {
		return false
	}

	if !(nc.Subscribable.Equals(other.Subscribable)) {
		return false
	}

	return true
}

// CopyRef copies the current value of the NewsChannel
// and returns a pointer to the new copy
func (nc NewsChannel) CopyRef() types.RVTypePtr {
	copied := nc.Copy().(NewsChannel)
	return &copied
}

// Deref takes a pointer to the NewsChannel
// and dereferences it to the raw value.
// Only useful when working with an instance of RVTypePtr
func (nc *NewsChannel) Deref() types.RVType {
	return *nc
}

// String returns the string representation of the NewsChannel
func (nc NewsChannel) String() string {
	return nc.FormatToString(0)
}

// FormatToString pretty-prints the NewsChannel using the provided indentation level
func (nc NewsChannel) FormatToString(indentationLevel int) string {
	indentationValues := strings.Repeat("\t", indentationLevel+1)
	indentationEnd := strings.Repeat("\t", indentationLevel)

	var b strings.Builder

	b.WriteString("NewsChannel{\n")
	fmt.Fprintf(&b, "%sID: %s,\n", indentationValues, nc.ID)
	fmt.Fprintf(&b, "%sOwnerPID: %s,\n", indentationValues, nc.OwnerPID)
	fmt.Fprintf(&b, "%sName: %s,\n", indentationValues, nc.Name)
	fmt.Fprintf(&b, "%sDescription: %s,\n", indentationValues, nc.Description)
	fmt.Fprintf(&b, "%sCreationTime: %s,\n", indentationValues, nc.CreationTime.FormatToString(indentationLevel+1))
	fmt.Fprintf(&b, "%sExpirationTime: %s,\n", indentationValues, nc.ExpirationTime.FormatToString(indentationLevel+1))
	fmt.Fprintf(&b, "%sType: %s,\n", indentationValues, nc.Type)
	fmt.Fprintf(&b, "%sLocale: %s,\n", indentationValues, nc.Locale)
	fmt.Fprintf(&b, "%sSubscribable: %s,\n", indentationValues, nc.Subscribable)
	fmt.Fprintf(&b, "%s}", indentationEnd)

	return b.String()
}

// NewNewsChannel returns a new NewsChannel
func NewNewsChannel() NewsChannel {
	return NewsChannel{
		ID:             types.NewUInt32(0),
		OwnerPID:       types.NewUInt32(0),
		Name:           types.NewString(""),
		Description:    types.NewString(""),
		CreationTime:   types.NewDateTime(0),
		ExpirationTime: types.NewDateTime(0),
		Type:           types.NewString(""),
		Locale:         types.NewString(""),
		Subscribable:   types.NewBool(false),
	}
}
