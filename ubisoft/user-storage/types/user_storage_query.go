package types

import (
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go/v2/types"
)

// UserStorageQuery is a type within the User Storage protocol
type UserStorageQuery struct {
	types.Structure
	TypeID      types.UInt32
	QueryID     types.UInt32
	ResultRange types.ResultRange
	Parameters  types.List[ContentProperty]
}

// WriteTo writes the UserStorageQuery to the given writable
func (usq UserStorageQuery) WriteTo(writable types.Writable) {
	contentWritable := writable.CopyNew()

	usq.TypeID.WriteTo(contentWritable)
	usq.QueryID.WriteTo(contentWritable)
	usq.ResultRange.WriteTo(contentWritable)
	usq.Parameters.WriteTo(contentWritable)

	content := contentWritable.Bytes()

	usq.WriteHeaderTo(writable, uint32(len(content)))

	writable.Write(content)
}

// ExtractFrom extracts the UserStorageQuery from the given readable
func (usq *UserStorageQuery) ExtractFrom(readable types.Readable) error {
	if err := usq.ExtractHeaderFrom(readable); err != nil {
		return fmt.Errorf("failed to extract UserStorageQuery header. %s", err.Error())
	}

	if err := usq.TypeID.ExtractFrom(readable); err != nil {
		return fmt.Errorf("failed to extract UserStorageQuery.TypeID. %s", err.Error())
	}

	if err := usq.QueryID.ExtractFrom(readable); err != nil {
		return fmt.Errorf("failed to extract UserStorageQuery.QueryID. %s", err.Error())
	}

	if err := usq.ResultRange.ExtractFrom(readable); err != nil {
		return fmt.Errorf("failed to extract UserStorageQuery.ResultRange. %s", err.Error())
	}

	if err := usq.Parameters.ExtractFrom(readable); err != nil {
		return fmt.Errorf("failed to extract UserStorageQuery.Parameters. %s", err.Error())
	}

	return nil
}

// Copy returns a new copied instance of UserStorageQuery
func (usq UserStorageQuery) Copy() types.RVType {
	copied := NewUserStorageQuery()

	copied.StructureVersion = usq.StructureVersion
	copied.TypeID = usq.TypeID.Copy().(types.UInt32)
	copied.QueryID = usq.QueryID.Copy().(types.UInt32)
	copied.ResultRange = usq.ResultRange.Copy().(types.ResultRange)
	copied.Parameters = usq.Parameters.Copy().(types.List[ContentProperty])

	return copied
}

// Equals checks if the given UserStorageQuery contains the same data as the current UserStorageQuery
func (usq UserStorageQuery) Equals(o types.RVType) bool {
	other, ok := o.(UserStorageQuery)
	if !ok {
		return false
	}

	if usq.StructureVersion != other.StructureVersion {
		return false
	}

	if !(usq.TypeID.Equals(other.TypeID)) {
		return false
	}

	if !(usq.QueryID.Equals(other.QueryID)) {
		return false
	}

	if !(usq.ResultRange == other.ResultRange) {
		return false
	}

	if !(usq.Parameters.Equals(other.Parameters)) {
		return false
	}

	return true
}

// CopyRef copies the current value of the UserStorageQuery
// and returns a pointer to the new copy
func (usq UserStorageQuery) CopyRef() types.RVTypePtr {
	copied := usq.Copy().(UserStorageQuery)
	return &copied
}

// Deref takes a pointer to the UserStorageQuery
// and dereferences it to the raw value.
// Only useful when working with an instance of RVTypePtr
func (usq *UserStorageQuery) Deref() types.RVType {
	return *usq
}

// String returns the string representation of the UserStorageQuery
func (usq UserStorageQuery) String() string {
	return usq.FormatToString(0)
}

// FormatToString pretty-prints the UserStorageQuery using the provided indentation level
func (usq UserStorageQuery) FormatToString(indentationLevel int) string {
	indentationValues := strings.Repeat("\t", indentationLevel+1)
	indentationEnd := strings.Repeat("\t", indentationLevel)

	var b strings.Builder

	b.WriteString("UserStorageQuery{\n")
	fmt.Fprintf(&b, "%sTypeID: %s,\n", indentationValues, usq.TypeID)
	fmt.Fprintf(&b, "%sQueryID: %s,\n", indentationValues, usq.QueryID)
	fmt.Fprintf(&b, "%sResultRange: %s,\n", indentationValues, usq.ResultRange.FormatToString(indentationLevel+1))
	fmt.Fprintf(&b, "%sParameters: %s,\n", indentationValues, usq.Parameters)
	fmt.Fprintf(&b, "%s}", indentationEnd)

	return b.String()
}

// NewUserStorageQuery returns a new UserStorageQuery
func NewUserStorageQuery() UserStorageQuery {
	return UserStorageQuery{
		TypeID:      types.NewUInt32(0),
		QueryID:     types.NewUInt32(0),
		ResultRange: types.NewResultRange(),
		Parameters:  types.NewList[ContentProperty](),
	}
}
