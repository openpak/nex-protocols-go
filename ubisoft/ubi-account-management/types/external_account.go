// Package types implements all the types used by the Ubi Account Management protocol
package types

import (
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go/v2/types"
)

// ExternalAccount is a type within the Ubi Account Management protocol
type ExternalAccount struct {
	types.Structure
	AccountType types.UInt32
	ID          types.String
	Username    types.String
}

// WriteTo writes the ExternalAccount to the given writable
func (s ExternalAccount) WriteTo(writable types.Writable) {
	contentWritable := writable.CopyNew()

	s.AccountType.WriteTo(contentWritable)
	s.ID.WriteTo(contentWritable)
	s.Username.WriteTo(contentWritable)

	content := contentWritable.Bytes()

	s.WriteHeaderTo(writable, uint32(len(content)))

	writable.Write(content)
}

// ExtractFrom extracts the ExternalAccount from the given readable
func (s *ExternalAccount) ExtractFrom(readable types.Readable) error {
	if err := s.ExtractHeaderFrom(readable); err != nil {
		return fmt.Errorf("failed to extract ExternalAccount header. %s", err.Error())
	}

	if err := s.AccountType.ExtractFrom(readable); err != nil {
		return fmt.Errorf("failed to extract ExternalAccount.AccountType. %s", err.Error())
	}

	if err := s.ID.ExtractFrom(readable); err != nil {
		return fmt.Errorf("failed to extract ExternalAccount.ID. %s", err.Error())
	}

	if err := s.Username.ExtractFrom(readable); err != nil {
		return fmt.Errorf("failed to extract ExternalAccount.Username. %s", err.Error())
	}

	return nil
}

// Copy returns a new copied instance of ExternalAccount
func (s ExternalAccount) Copy() types.RVType {
	copied := NewExternalAccount()

	copied.StructureVersion = s.StructureVersion
	copied.AccountType = s.AccountType.Copy().(types.UInt32)
	copied.ID = s.ID.Copy().(types.String)
	copied.Username = s.Username.Copy().(types.String)

	return copied
}

// Equals checks if the given ExternalAccount contains the same data as the current ExternalAccount
func (s ExternalAccount) Equals(o types.RVType) bool {
	other, ok := o.(ExternalAccount)
	if !ok {
		return false
	}

	if s.StructureVersion != other.StructureVersion {
		return false
	}

	if !s.AccountType.Equals(other.AccountType) {
		return false
	}

	if !s.ID.Equals(other.ID) {
		return false
	}

	if !s.Username.Equals(other.Username) {
		return false
	}

	return true
}

// CopyRef copies the current value of the ExternalAccount
// and returns a pointer to the new copy
func (s ExternalAccount) CopyRef() types.RVTypePtr {
	copied := s.Copy().(ExternalAccount)
	return &copied
}

// Deref takes a pointer to the ExternalAccount
// and dereferences it to the raw value.
// Only useful when working with an instance of RVTypePtr
func (s *ExternalAccount) Deref() types.RVType {
	return *s
}

// String returns the string representation of the ExternalAccount
func (s ExternalAccount) String() string {
	return s.FormatToString(0)
}

// FormatToString pretty-prints the ExternalAccount using the provided indentation level
func (s ExternalAccount) FormatToString(indentationLevel int) string {
	indentationValues := strings.Repeat("\t", indentationLevel+1)
	indentationEnd := strings.Repeat("\t", indentationLevel)

	var b strings.Builder

	b.WriteString("ExternalAccount{\n")
	fmt.Fprintf(&b, "%sAccountType: %s,\n", indentationValues, s.AccountType)
	fmt.Fprintf(&b, "%sID: %s,\n", indentationValues, s.ID)
	fmt.Fprintf(&b, "%sUsername: %s,\n", indentationValues, s.Username)
	fmt.Fprintf(&b, "%s}", indentationEnd)

	return b.String()
}

// NewExternalAccount returns a new ExternalAccount
func NewExternalAccount() ExternalAccount {
	return ExternalAccount{
		AccountType: types.NewUInt32(0),
		ID:          types.NewString(""),
		Username:    types.NewString(""),
	}
}
