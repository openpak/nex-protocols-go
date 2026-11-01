// Package types implements all the types used by the Uplay Win protocol
package types

import (
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go/v2/types"
)

// UplaySectionContentLocalized is a type within the Uplay Win protocol
type UplaySectionContentLocalized struct {
	types.Structure
	Key      types.String
	Culture  types.String
	Text     types.String
	Url      types.String
	Duration types.Int32
	Size     types.String
	Width    types.String
	Height   types.String
}

// WriteTo writes the UplaySectionContentLocalized to the given writable
func (s UplaySectionContentLocalized) WriteTo(writable types.Writable) {
	contentWritable := writable.CopyNew()

	s.Key.WriteTo(contentWritable)
	s.Culture.WriteTo(contentWritable)
	s.Text.WriteTo(contentWritable)
	s.Url.WriteTo(contentWritable)
	s.Duration.WriteTo(contentWritable)
	s.Size.WriteTo(contentWritable)
	s.Width.WriteTo(contentWritable)
	s.Height.WriteTo(contentWritable)

	content := contentWritable.Bytes()

	s.WriteHeaderTo(writable, uint32(len(content)))

	writable.Write(content)
}

// ExtractFrom extracts the UplaySectionContentLocalized from the given readable
func (s *UplaySectionContentLocalized) ExtractFrom(readable types.Readable) error {
	if err := s.ExtractHeaderFrom(readable); err != nil {
		return fmt.Errorf("failed to extract UplaySectionContentLocalized header. %s", err.Error())
	}

	if err := s.Key.ExtractFrom(readable); err != nil {
		return fmt.Errorf("failed to extract UplaySectionContentLocalized.Key. %s", err.Error())
	}

	if err := s.Culture.ExtractFrom(readable); err != nil {
		return fmt.Errorf("failed to extract UplaySectionContentLocalized.Culture. %s", err.Error())
	}

	if err := s.Text.ExtractFrom(readable); err != nil {
		return fmt.Errorf("failed to extract UplaySectionContentLocalized.Text. %s", err.Error())
	}

	if err := s.Url.ExtractFrom(readable); err != nil {
		return fmt.Errorf("failed to extract UplaySectionContentLocalized.Url. %s", err.Error())
	}

	if err := s.Duration.ExtractFrom(readable); err != nil {
		return fmt.Errorf("failed to extract UplaySectionContentLocalized.Duration. %s", err.Error())
	}

	if err := s.Size.ExtractFrom(readable); err != nil {
		return fmt.Errorf("failed to extract UplaySectionContentLocalized.Size. %s", err.Error())
	}

	if err := s.Width.ExtractFrom(readable); err != nil {
		return fmt.Errorf("failed to extract UplaySectionContentLocalized.Width. %s", err.Error())
	}

	if err := s.Height.ExtractFrom(readable); err != nil {
		return fmt.Errorf("failed to extract UplaySectionContentLocalized.Height. %s", err.Error())
	}

	return nil
}

// Copy returns a new copied instance of UplaySectionContentLocalized
func (s UplaySectionContentLocalized) Copy() types.RVType {
	copied := NewUplaySectionContentLocalized()

	copied.StructureVersion = s.StructureVersion
	copied.Key = s.Key.Copy().(types.String)
	copied.Culture = s.Culture.Copy().(types.String)
	copied.Text = s.Text.Copy().(types.String)
	copied.Url = s.Url.Copy().(types.String)
	copied.Duration = s.Duration.Copy().(types.Int32)
	copied.Size = s.Size.Copy().(types.String)
	copied.Width = s.Width.Copy().(types.String)
	copied.Height = s.Height.Copy().(types.String)

	return copied
}

// Equals checks if the given UplaySectionContentLocalized contains the same data as the current UplaySectionContentLocalized
func (s UplaySectionContentLocalized) Equals(o types.RVType) bool {
	other, ok := o.(UplaySectionContentLocalized)
	if !ok {
		return false
	}

	if s.StructureVersion != other.StructureVersion {
		return false
	}

	if !s.Key.Equals(other.Key) {
		return false
	}

	if !s.Culture.Equals(other.Culture) {
		return false
	}

	if !s.Text.Equals(other.Text) {
		return false
	}

	if !s.Url.Equals(other.Url) {
		return false
	}

	if !s.Duration.Equals(other.Duration) {
		return false
	}

	if !s.Size.Equals(other.Size) {
		return false
	}

	if !s.Width.Equals(other.Width) {
		return false
	}

	if !s.Height.Equals(other.Height) {
		return false
	}

	return true
}

// CopyRef copies the current value of the UplaySectionContentLocalized
// and returns a pointer to the new copy
func (s UplaySectionContentLocalized) CopyRef() types.RVTypePtr {
	copied := s.Copy().(UplaySectionContentLocalized)
	return &copied
}

// Deref takes a pointer to the UplaySectionContentLocalized
// and dereferences it to the raw value.
// Only useful when working with an instance of RVTypePtr
func (s *UplaySectionContentLocalized) Deref() types.RVType {
	return *s
}

// String returns the string representation of the UplaySectionContentLocalized
func (s UplaySectionContentLocalized) String() string {
	return s.FormatToString(0)
}

// FormatToString pretty-prints the UplaySectionContentLocalized using the provided indentation level
func (s UplaySectionContentLocalized) FormatToString(indentationLevel int) string {
	indentationValues := strings.Repeat("\t", indentationLevel+1)
	indentationEnd := strings.Repeat("\t", indentationLevel)

	var b strings.Builder

	b.WriteString("UplaySectionContentLocalized{\n")
	fmt.Fprintf(&b, "%sKey: %s,\n", indentationValues, s.Key)
	fmt.Fprintf(&b, "%sCulture: %s,\n", indentationValues, s.Culture)
	fmt.Fprintf(&b, "%sText: %s,\n", indentationValues, s.Text)
	fmt.Fprintf(&b, "%sUrl: %s,\n", indentationValues, s.Url)
	fmt.Fprintf(&b, "%sDuration: %s,\n", indentationValues, s.Duration)
	fmt.Fprintf(&b, "%sSize: %s,\n", indentationValues, s.Size)
	fmt.Fprintf(&b, "%sWidth: %s,\n", indentationValues, s.Width)
	fmt.Fprintf(&b, "%sHeight: %s,\n", indentationValues, s.Height)
	fmt.Fprintf(&b, "%s}", indentationEnd)

	return b.String()
}

// NewUplaySectionContentLocalized returns a new UplaySectionContentLocalized
func NewUplaySectionContentLocalized() UplaySectionContentLocalized {
	return UplaySectionContentLocalized{
		Key:      types.NewString(""),
		Culture:  types.NewString(""),
		Text:     types.NewString(""),
		Url:      types.NewString(""),
		Duration: types.NewInt32(0),
		Size:     types.NewString(""),
		Width:    types.NewString(""),
		Height:   types.NewString(""),
	}
}
