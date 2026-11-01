package types

import (
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go/v2/types"
)

// TrackingTag is a type within the Tracking 3 protocol
type TrackingTag struct {
	types.Structure
	TrackingID types.UInt32
	Tag        types.String
	Attributes types.String
	DeltaTime  types.UInt32
	NewUserId  types.String
}

// WriteTo writes the TrackingTag to the given writable
func (tt TrackingTag) WriteTo(writable types.Writable) {
	contentWritable := writable.CopyNew()

	tt.TrackingID.WriteTo(contentWritable)
	tt.Tag.WriteTo(contentWritable)
	tt.Attributes.WriteTo(contentWritable)
	tt.DeltaTime.WriteTo(contentWritable)
	tt.NewUserId.WriteTo(contentWritable)

	content := contentWritable.Bytes()

	tt.WriteHeaderTo(writable, uint32(len(content)))

	writable.Write(content)
}

// ExtractFrom extracts the TrackingTag from the given readable
func (tt *TrackingTag) ExtractFrom(readable types.Readable) error {
	if err := tt.ExtractHeaderFrom(readable); err != nil {
		return fmt.Errorf("failed to extract TrackingTag header. %s", err.Error())
	}

	if err := tt.TrackingID.ExtractFrom(readable); err != nil {
		return fmt.Errorf("failed to extract TrackingTag.TrackingID. %s", err.Error())
	}

	if err := tt.Tag.ExtractFrom(readable); err != nil {
		return fmt.Errorf("failed to extract TrackingTag.Tag. %s", err.Error())
	}

	if err := tt.Attributes.ExtractFrom(readable); err != nil {
		return fmt.Errorf("failed to extract TrackingTag.Attributes. %s", err.Error())
	}

	if err := tt.DeltaTime.ExtractFrom(readable); err != nil {
		return fmt.Errorf("failed to extract TrackingTag.DeltaTime. %s", err.Error())
	}

	if err := tt.NewUserId.ExtractFrom(readable); err != nil {
		return fmt.Errorf("failed to extract TrackingTag.NewUserId. %s", err.Error())
	}

	return nil
}

// Copy returns a new copied instance of TrackingTag
func (tt TrackingTag) Copy() types.RVType {
	copied := NewTrackingTag()

	copied.StructureVersion = tt.StructureVersion
	copied.TrackingID = tt.TrackingID.Copy().(types.UInt32)
	copied.Tag = tt.Tag.Copy().(types.String)
	copied.Attributes = tt.Attributes.Copy().(types.String)
	copied.DeltaTime = tt.DeltaTime.Copy().(types.UInt32)
	copied.NewUserId = tt.NewUserId.Copy().(types.String)

	return copied
}

// Equals checks if the given TrackingTag contains the same data as the current TrackingTag
func (tt TrackingTag) Equals(o types.RVType) bool {
	other, ok := o.(TrackingTag)
	if !ok {
		return false
	}

	if tt.StructureVersion != other.StructureVersion {
		return false
	}

	if !(tt.TrackingID.Equals(other.TrackingID)) {
		return false
	}

	if !(tt.Tag.Equals(other.Tag)) {
		return false
	}

	if !(tt.Attributes.Equals(other.Attributes)) {
		return false
	}

	if !(tt.DeltaTime.Equals(other.DeltaTime)) {
		return false
	}

	if !(tt.NewUserId.Equals(other.NewUserId)) {
		return false
	}

	return true
}

// CopyRef copies the current value of the TrackingTag
// and returns a pointer to the new copy
func (tt TrackingTag) CopyRef() types.RVTypePtr {
	copied := tt.Copy().(TrackingTag)
	return &copied
}

// Deref takes a pointer to the TrackingTag
// and dereferences it to the raw value.
// Only useful when working with an instance of RVTypePtr
func (tt *TrackingTag) Deref() types.RVType {
	return *tt
}

// String returns the string representation of the TrackingTag
func (tt TrackingTag) String() string {
	return tt.FormatToString(0)
}

// FormatToString pretty-prints the TrackingTag using the provided indentation level
func (tt TrackingTag) FormatToString(indentationLevel int) string {
	indentationValues := strings.Repeat("\t", indentationLevel+1)
	indentationEnd := strings.Repeat("\t", indentationLevel)

	var b strings.Builder

	b.WriteString("TrackingTag{\n")
	fmt.Fprintf(&b, "%sTrackingID: %s,\n", indentationValues, tt.TrackingID)
	fmt.Fprintf(&b, "%sTag: %s,\n", indentationValues, tt.Tag)
	fmt.Fprintf(&b, "%sAttributes: %s,\n", indentationValues, tt.Attributes)
	fmt.Fprintf(&b, "%sDeltaTime: %s,\n", indentationValues, tt.DeltaTime)
	fmt.Fprintf(&b, "%sNewUserId: %s,\n", indentationValues, tt.NewUserId)
	fmt.Fprintf(&b, "%s}", indentationEnd)

	return b.String()
}

// NewTrackingTag returns a new TrackingTag
func NewTrackingTag() TrackingTag {
	return TrackingTag{
		TrackingID: types.NewUInt32(0),
		Tag:        types.NewString(""),
		Attributes: types.NewString(""),
		DeltaTime:  types.NewUInt32(0),
		NewUserId:  types.NewString(""),
	}
}
