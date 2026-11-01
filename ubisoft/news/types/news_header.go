package types

import (
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go/v2/types"
)

// NewsHeader is a type within the News protocol
type NewsHeader struct {
	types.Structure
	ID              types.UInt32
	RecipientID     types.UInt32
	RecipientType   types.UInt32
	PublisherPID    types.UInt32
	PublisherName   types.String
	PublicationTime types.DateTime
	DisplayTime     types.DateTime
	ExpirationTime  types.DateTime
	Title           types.String
	Link            types.String
}

// WriteTo writes the NewsHeader to the given writable
func (nh NewsHeader) WriteTo(writable types.Writable) {
	contentWritable := writable.CopyNew()

	nh.ID.WriteTo(contentWritable)
	nh.RecipientID.WriteTo(contentWritable)
	nh.RecipientType.WriteTo(contentWritable)
	nh.PublisherPID.WriteTo(contentWritable)
	nh.PublisherName.WriteTo(contentWritable)
	nh.PublicationTime.WriteTo(contentWritable)
	nh.DisplayTime.WriteTo(contentWritable)
	nh.ExpirationTime.WriteTo(contentWritable)
	nh.Title.WriteTo(contentWritable)
	nh.Link.WriteTo(contentWritable)

	content := contentWritable.Bytes()

	nh.WriteHeaderTo(writable, uint32(len(content)))

	writable.Write(content)
}

// ExtractFrom extracts the NewsHeader from the given readable
func (nh *NewsHeader) ExtractFrom(readable types.Readable) error {
	if err := nh.ExtractHeaderFrom(readable); err != nil {
		return fmt.Errorf("failed to extract NewsHeader header. %s", err.Error())
	}

	if err := nh.ID.ExtractFrom(readable); err != nil {
		return fmt.Errorf("failed to extract NewsHeader.ID. %s", err.Error())
	}

	if err := nh.RecipientID.ExtractFrom(readable); err != nil {
		return fmt.Errorf("failed to extract NewsHeader.RecipientID. %s", err.Error())
	}

	if err := nh.RecipientType.ExtractFrom(readable); err != nil {
		return fmt.Errorf("failed to extract NewsHeader.RecipientType. %s", err.Error())
	}

	if err := nh.PublisherPID.ExtractFrom(readable); err != nil {
		return fmt.Errorf("failed to extract NewsHeader.PublisherPID. %s", err.Error())
	}

	if err := nh.PublisherName.ExtractFrom(readable); err != nil {
		return fmt.Errorf("failed to extract NewsHeader.PublisherName. %s", err.Error())
	}

	if err := nh.PublicationTime.ExtractFrom(readable); err != nil {
		return fmt.Errorf("failed to extract NewsHeader.PublicationTime. %s", err.Error())
	}

	if err := nh.DisplayTime.ExtractFrom(readable); err != nil {
		return fmt.Errorf("failed to extract NewsHeader.DisplayTime. %s", err.Error())
	}

	if err := nh.ExpirationTime.ExtractFrom(readable); err != nil {
		return fmt.Errorf("failed to extract NewsHeader.ExpirationTime. %s", err.Error())
	}

	if err := nh.Title.ExtractFrom(readable); err != nil {
		return fmt.Errorf("failed to extract NewsHeader.Title. %s", err.Error())
	}

	if err := nh.Link.ExtractFrom(readable); err != nil {
		return fmt.Errorf("failed to extract NewsHeader.Link. %s", err.Error())
	}

	return nil
}

// Copy returns a new copied instance of NewsHeader
func (nh NewsHeader) Copy() types.RVType {
	copied := NewNewsHeader()

	copied.StructureVersion = nh.StructureVersion
	copied.ID = nh.ID.Copy().(types.UInt32)
	copied.RecipientID = nh.RecipientID.Copy().(types.UInt32)
	copied.RecipientType = nh.RecipientType.Copy().(types.UInt32)
	copied.PublisherPID = nh.PublisherPID.Copy().(types.UInt32)
	copied.PublisherName = nh.PublisherName.Copy().(types.String)
	copied.PublicationTime = nh.PublicationTime.Copy().(types.DateTime)
	copied.DisplayTime = nh.DisplayTime.Copy().(types.DateTime)
	copied.ExpirationTime = nh.ExpirationTime.Copy().(types.DateTime)
	copied.Title = nh.Title.Copy().(types.String)
	copied.Link = nh.Link.Copy().(types.String)

	return copied
}

// Equals checks if the given NewsHeader contains the same data as the current NewsHeader
func (nh NewsHeader) Equals(o types.RVType) bool {
	other, ok := o.(NewsHeader)
	if !ok {
		return false
	}

	if nh.StructureVersion != other.StructureVersion {
		return false
	}

	if !(nh.ID.Equals(other.ID)) {
		return false
	}

	if !(nh.RecipientID.Equals(other.RecipientID)) {
		return false
	}

	if !(nh.RecipientType.Equals(other.RecipientType)) {
		return false
	}

	if !(nh.PublisherPID.Equals(other.PublisherPID)) {
		return false
	}

	if !(nh.PublisherName.Equals(other.PublisherName)) {
		return false
	}

	if !(nh.PublicationTime.Equals(other.PublicationTime)) {
		return false
	}

	if !(nh.DisplayTime.Equals(other.DisplayTime)) {
		return false
	}

	if !(nh.ExpirationTime.Equals(other.ExpirationTime)) {
		return false
	}

	if !(nh.Title.Equals(other.Title)) {
		return false
	}

	if !(nh.Link.Equals(other.Link)) {
		return false
	}

	return true
}

// CopyRef copies the current value of the NewsHeader
// and returns a pointer to the new copy
func (nh NewsHeader) CopyRef() types.RVTypePtr {
	copied := nh.Copy().(NewsHeader)
	return &copied
}

// Deref takes a pointer to the NewsHeader
// and dereferences it to the raw value.
// Only useful when working with an instance of RVTypePtr
func (nh *NewsHeader) Deref() types.RVType {
	return *nh
}

// String returns the string representation of the NewsHeader
func (nh NewsHeader) String() string {
	return nh.FormatToString(0)
}

// FormatToString pretty-prints the NewsHeader using the provided indentation level
func (nh NewsHeader) FormatToString(indentationLevel int) string {
	indentationValues := strings.Repeat("\t", indentationLevel+1)
	indentationEnd := strings.Repeat("\t", indentationLevel)

	var b strings.Builder

	b.WriteString("NewsHeader{\n")
	fmt.Fprintf(&b, "%sID: %s,\n", indentationValues, nh.ID)
	fmt.Fprintf(&b, "%sRecipientID: %s,\n", indentationValues, nh.RecipientID)
	fmt.Fprintf(&b, "%sRecipientType: %s,\n", indentationValues, nh.RecipientType)
	fmt.Fprintf(&b, "%sPublisherPID: %s,\n", indentationValues, nh.PublisherPID)
	fmt.Fprintf(&b, "%sPublisherName: %s,\n", indentationValues, nh.PublisherName)
	fmt.Fprintf(&b, "%sPublicationTime: %s,\n", indentationValues, nh.PublicationTime.FormatToString(indentationLevel+1))
	fmt.Fprintf(&b, "%sDisplayTime: %s,\n", indentationValues, nh.DisplayTime.FormatToString(indentationLevel+1))
	fmt.Fprintf(&b, "%sExpirationTime: %s,\n", indentationValues, nh.ExpirationTime.FormatToString(indentationLevel+1))
	fmt.Fprintf(&b, "%sTitle: %s,\n", indentationValues, nh.Title)
	fmt.Fprintf(&b, "%sLink: %s,\n", indentationValues, nh.Link)
	fmt.Fprintf(&b, "%s}", indentationEnd)

	return b.String()
}

// NewNewsHeader returns a new NewsHeader
func NewNewsHeader() NewsHeader {
	return NewsHeader{
		ID:              types.NewUInt32(0),
		RecipientID:     types.NewUInt32(0),
		RecipientType:   types.NewUInt32(0),
		PublisherPID:    types.NewUInt32(0),
		PublisherName:   types.NewString(""),
		PublicationTime: types.NewDateTime(0),
		DisplayTime:     types.NewDateTime(0),
		ExpirationTime:  types.NewDateTime(0),
		Title:           types.NewString(""),
		Link:            types.NewString(""),
	}
}
