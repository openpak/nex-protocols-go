package types

import (
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go/v2/types"
)

// NewsRecipient is a type within the News protocol
type NewsRecipient struct {
	types.Structure
	RecipientID   types.UInt32
	RecipientType types.UInt32
}

// WriteTo writes the NewsRecipient to the given writable
func (nr NewsRecipient) WriteTo(writable types.Writable) {
	contentWritable := writable.CopyNew()

	nr.RecipientID.WriteTo(contentWritable)
	nr.RecipientType.WriteTo(contentWritable)

	content := contentWritable.Bytes()

	nr.WriteHeaderTo(writable, uint32(len(content)))

	writable.Write(content)
}

// ExtractFrom extracts the NewsRecipient from the given readable
func (nr *NewsRecipient) ExtractFrom(readable types.Readable) error {
	if err := nr.ExtractHeaderFrom(readable); err != nil {
		return fmt.Errorf("failed to extract NewsRecipient header. %s", err.Error())
	}

	if err := nr.RecipientID.ExtractFrom(readable); err != nil {
		return fmt.Errorf("failed to extract NewsRecipient.RecipientID. %s", err.Error())
	}

	if err := nr.RecipientType.ExtractFrom(readable); err != nil {
		return fmt.Errorf("failed to extract NewsRecipient.RecipientType. %s", err.Error())
	}

	return nil
}

// Copy returns a new copied instance of NewsRecipient
func (nr NewsRecipient) Copy() types.RVType {
	copied := NewNewsRecipient()

	copied.StructureVersion = nr.StructureVersion
	copied.RecipientID = nr.RecipientID.Copy().(types.UInt32)
	copied.RecipientType = nr.RecipientType.Copy().(types.UInt32)

	return copied
}

// Equals checks if the given NewsRecipient contains the same data as the current NewsRecipient
func (nr NewsRecipient) Equals(o types.RVType) bool {
	other, ok := o.(NewsRecipient)
	if !ok {
		return false
	}

	if nr.StructureVersion != other.StructureVersion {
		return false
	}

	if !(nr.RecipientID.Equals(other.RecipientID)) {
		return false
	}

	if !(nr.RecipientType.Equals(other.RecipientType)) {
		return false
	}

	return true
}

// CopyRef copies the current value of the NewsRecipient
// and returns a pointer to the new copy
func (nr NewsRecipient) CopyRef() types.RVTypePtr {
	copied := nr.Copy().(NewsRecipient)
	return &copied
}

// Deref takes a pointer to the NewsRecipient
// and dereferences it to the raw value.
// Only useful when working with an instance of RVTypePtr
func (nr *NewsRecipient) Deref() types.RVType {
	return *nr
}

// String returns the string representation of the NewsRecipient
func (nr NewsRecipient) String() string {
	return nr.FormatToString(0)
}

// FormatToString pretty-prints the NewsRecipient using the provided indentation level
func (nr NewsRecipient) FormatToString(indentationLevel int) string {
	indentationValues := strings.Repeat("\t", indentationLevel+1)
	indentationEnd := strings.Repeat("\t", indentationLevel)

	var b strings.Builder

	b.WriteString("NewsRecipient{\n")
	fmt.Fprintf(&b, "%sRecipientID: %s,\n", indentationValues, nr.RecipientID)
	fmt.Fprintf(&b, "%sRecipientType: %s,\n", indentationValues, nr.RecipientType)
	fmt.Fprintf(&b, "%s}", indentationEnd)

	return b.String()
}

// NewNewsRecipient returns a new NewsRecipient
func NewNewsRecipient() NewsRecipient {
	return NewsRecipient{
		RecipientID:   types.NewUInt32(0),
		RecipientType: types.NewUInt32(0),
	}
}
