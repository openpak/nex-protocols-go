package types

import (
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go/v2/types"
)

// NewsMessage is a type within the News protocol
type NewsMessage struct {
	types.Structure
	Body types.String
}

// WriteTo writes the NewsMessage to the given writable
func (nm NewsMessage) WriteTo(writable types.Writable) {
	contentWritable := writable.CopyNew()

	nm.Body.WriteTo(contentWritable)

	content := contentWritable.Bytes()

	nm.WriteHeaderTo(writable, uint32(len(content)))

	writable.Write(content)
}

// ExtractFrom extracts the NewsMessage from the given readable
func (nm *NewsMessage) ExtractFrom(readable types.Readable) error {
	if err := nm.ExtractHeaderFrom(readable); err != nil {
		return fmt.Errorf("failed to extract NewsMessage header. %s", err.Error())
	}

	if err := nm.Body.ExtractFrom(readable); err != nil {
		return fmt.Errorf("failed to extract NewsMessage.Body. %s", err.Error())
	}

	return nil
}

// Copy returns a new copied instance of NewsMessage
func (nm NewsMessage) Copy() types.RVType {
	copied := NewNewsMessage()

	copied.StructureVersion = nm.StructureVersion
	copied.Body = nm.Body.Copy().(types.String)

	return copied
}

// Equals checks if the given NewsMessage contains the same data as the current NewsMessage
func (nm NewsMessage) Equals(o types.RVType) bool {
	other, ok := o.(NewsMessage)
	if !ok {
		return false
	}

	if nm.StructureVersion != other.StructureVersion {
		return false
	}

	if !(nm.Body.Equals(other.Body)) {
		return false
	}

	return true
}

// CopyRef copies the current value of the NewsMessage
// and returns a pointer to the new copy
func (nm NewsMessage) CopyRef() types.RVTypePtr {
	copied := nm.Copy().(NewsMessage)
	return &copied
}

// Deref takes a pointer to the NewsMessage
// and dereferences it to the raw value.
// Only useful when working with an instance of RVTypePtr
func (nm *NewsMessage) Deref() types.RVType {
	return *nm
}

// String returns the string representation of the NewsMessage
func (nm NewsMessage) String() string {
	return nm.FormatToString(0)
}

// FormatToString pretty-prints the NewsMessage using the provided indentation level
func (nm NewsMessage) FormatToString(indentationLevel int) string {
	indentationValues := strings.Repeat("\t", indentationLevel+1)
	indentationEnd := strings.Repeat("\t", indentationLevel)

	var b strings.Builder

	b.WriteString("NewsMessage{\n")
	fmt.Fprintf(&b, "%sBody: %s,\n", indentationValues, nm.Body)
	fmt.Fprintf(&b, "%s}", indentationEnd)

	return b.String()
}

// NewNewsMessage returns a new NewsMessage
func NewNewsMessage() NewsMessage {
	return NewsMessage{
		Body: types.NewString(""),
	}
}
