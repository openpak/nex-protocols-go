// Package types implements all the types used by the Ubi Account Management protocol
package types

import (
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go/v2/types"
)

// UbiAccount is a type within the Ubi Account Management protocol
type UbiAccount struct {
	types.Structure
	UbiAccountID      types.String
	Username          types.String
	Password          types.String
	Status            UbiAccountStatus
	Email             types.String
	DateOfBirth       types.DateTime
	Gender            types.UInt32
	CountryCode       types.String
	OptIn             types.Bool
	ThirdPartyOptIn   types.Bool
	FirstName         types.String
	LastName          types.String
	PreferredLanguage types.String
	ExternalAccounts  types.List[ExternalAccount]
}

// WriteTo writes the UbiAccount to the given writable
func (s UbiAccount) WriteTo(writable types.Writable) {
	contentWritable := writable.CopyNew()

	s.UbiAccountID.WriteTo(contentWritable)
	s.Username.WriteTo(contentWritable)
	s.Password.WriteTo(contentWritable)
	s.Status.WriteTo(contentWritable)
	s.Email.WriteTo(contentWritable)
	s.DateOfBirth.WriteTo(contentWritable)
	s.Gender.WriteTo(contentWritable)
	s.CountryCode.WriteTo(contentWritable)
	s.OptIn.WriteTo(contentWritable)
	s.ThirdPartyOptIn.WriteTo(contentWritable)
	s.FirstName.WriteTo(contentWritable)
	s.LastName.WriteTo(contentWritable)
	s.PreferredLanguage.WriteTo(contentWritable)
	s.ExternalAccounts.WriteTo(contentWritable)

	content := contentWritable.Bytes()

	s.WriteHeaderTo(writable, uint32(len(content)))

	writable.Write(content)
}

// ExtractFrom extracts the UbiAccount from the given readable
func (s *UbiAccount) ExtractFrom(readable types.Readable) error {
	if err := s.ExtractHeaderFrom(readable); err != nil {
		return fmt.Errorf("failed to extract UbiAccount header. %s", err.Error())
	}

	if err := s.UbiAccountID.ExtractFrom(readable); err != nil {
		return fmt.Errorf("failed to extract UbiAccount.UbiAccountID. %s", err.Error())
	}

	if err := s.Username.ExtractFrom(readable); err != nil {
		return fmt.Errorf("failed to extract UbiAccount.Username. %s", err.Error())
	}

	if err := s.Password.ExtractFrom(readable); err != nil {
		return fmt.Errorf("failed to extract UbiAccount.Password. %s", err.Error())
	}

	if err := s.Status.ExtractFrom(readable); err != nil {
		return fmt.Errorf("failed to extract UbiAccount.Status. %s", err.Error())
	}

	if err := s.Email.ExtractFrom(readable); err != nil {
		return fmt.Errorf("failed to extract UbiAccount.Email. %s", err.Error())
	}

	if err := s.DateOfBirth.ExtractFrom(readable); err != nil {
		return fmt.Errorf("failed to extract UbiAccount.DateOfBirth. %s", err.Error())
	}

	if err := s.Gender.ExtractFrom(readable); err != nil {
		return fmt.Errorf("failed to extract UbiAccount.Gender. %s", err.Error())
	}

	if err := s.CountryCode.ExtractFrom(readable); err != nil {
		return fmt.Errorf("failed to extract UbiAccount.CountryCode. %s", err.Error())
	}

	if err := s.OptIn.ExtractFrom(readable); err != nil {
		return fmt.Errorf("failed to extract UbiAccount.OptIn. %s", err.Error())
	}

	if err := s.ThirdPartyOptIn.ExtractFrom(readable); err != nil {
		return fmt.Errorf("failed to extract UbiAccount.ThirdPartyOptIn. %s", err.Error())
	}

	if err := s.FirstName.ExtractFrom(readable); err != nil {
		return fmt.Errorf("failed to extract UbiAccount.FirstName. %s", err.Error())
	}

	if err := s.LastName.ExtractFrom(readable); err != nil {
		return fmt.Errorf("failed to extract UbiAccount.LastName. %s", err.Error())
	}

	if err := s.PreferredLanguage.ExtractFrom(readable); err != nil {
		return fmt.Errorf("failed to extract UbiAccount.PreferredLanguage. %s", err.Error())
	}

	if err := s.ExternalAccounts.ExtractFrom(readable); err != nil {
		return fmt.Errorf("failed to extract UbiAccount.ExternalAccounts. %s", err.Error())
	}

	return nil
}

// Copy returns a new copied instance of UbiAccount
func (s UbiAccount) Copy() types.RVType {
	copied := NewUbiAccount()

	copied.StructureVersion = s.StructureVersion
	copied.UbiAccountID = s.UbiAccountID.Copy().(types.String)
	copied.Username = s.Username.Copy().(types.String)
	copied.Password = s.Password.Copy().(types.String)
	copied.Status = s.Status.Copy().(UbiAccountStatus)
	copied.Email = s.Email.Copy().(types.String)
	copied.DateOfBirth = s.DateOfBirth.Copy().(types.DateTime)
	copied.Gender = s.Gender.Copy().(types.UInt32)
	copied.CountryCode = s.CountryCode.Copy().(types.String)
	copied.OptIn = s.OptIn.Copy().(types.Bool)
	copied.ThirdPartyOptIn = s.ThirdPartyOptIn.Copy().(types.Bool)
	copied.FirstName = s.FirstName.Copy().(types.String)
	copied.LastName = s.LastName.Copy().(types.String)
	copied.PreferredLanguage = s.PreferredLanguage.Copy().(types.String)
	copied.ExternalAccounts = s.ExternalAccounts.Copy().(types.List[ExternalAccount])

	return copied
}

// Equals checks if the given UbiAccount contains the same data as the current UbiAccount
func (s UbiAccount) Equals(o types.RVType) bool {
	other, ok := o.(UbiAccount)
	if !ok {
		return false
	}

	if s.StructureVersion != other.StructureVersion {
		return false
	}

	if !s.UbiAccountID.Equals(other.UbiAccountID) {
		return false
	}

	if !s.Username.Equals(other.Username) {
		return false
	}

	if !s.Password.Equals(other.Password) {
		return false
	}

	if !s.Status.Equals(other.Status) {
		return false
	}

	if !s.Email.Equals(other.Email) {
		return false
	}

	if !s.DateOfBirth.Equals(other.DateOfBirth) {
		return false
	}

	if !s.Gender.Equals(other.Gender) {
		return false
	}

	if !s.CountryCode.Equals(other.CountryCode) {
		return false
	}

	if !s.OptIn.Equals(other.OptIn) {
		return false
	}

	if !s.ThirdPartyOptIn.Equals(other.ThirdPartyOptIn) {
		return false
	}

	if !s.FirstName.Equals(other.FirstName) {
		return false
	}

	if !s.LastName.Equals(other.LastName) {
		return false
	}

	if !s.PreferredLanguage.Equals(other.PreferredLanguage) {
		return false
	}

	if !s.ExternalAccounts.Equals(other.ExternalAccounts) {
		return false
	}

	return true
}

// CopyRef copies the current value of the UbiAccount
// and returns a pointer to the new copy
func (s UbiAccount) CopyRef() types.RVTypePtr {
	copied := s.Copy().(UbiAccount)
	return &copied
}

// Deref takes a pointer to the UbiAccount
// and dereferences it to the raw value.
// Only useful when working with an instance of RVTypePtr
func (s *UbiAccount) Deref() types.RVType {
	return *s
}

// String returns the string representation of the UbiAccount
func (s UbiAccount) String() string {
	return s.FormatToString(0)
}

// FormatToString pretty-prints the UbiAccount using the provided indentation level
func (s UbiAccount) FormatToString(indentationLevel int) string {
	indentationValues := strings.Repeat("\t", indentationLevel+1)
	indentationEnd := strings.Repeat("\t", indentationLevel)

	var b strings.Builder

	b.WriteString("UbiAccount{\n")
	fmt.Fprintf(&b, "%sUbiAccountID: %s,\n", indentationValues, s.UbiAccountID)
	fmt.Fprintf(&b, "%sUsername: %s,\n", indentationValues, s.Username)
	fmt.Fprintf(&b, "%sPassword: %s,\n", indentationValues, s.Password)
	fmt.Fprintf(&b, "%sStatus: %s,\n", indentationValues, s.Status)
	fmt.Fprintf(&b, "%sEmail: %s,\n", indentationValues, s.Email)
	fmt.Fprintf(&b, "%sDateOfBirth: %s,\n", indentationValues, s.DateOfBirth)
	fmt.Fprintf(&b, "%sGender: %s,\n", indentationValues, s.Gender)
	fmt.Fprintf(&b, "%sCountryCode: %s,\n", indentationValues, s.CountryCode)
	fmt.Fprintf(&b, "%sOptIn: %s,\n", indentationValues, s.OptIn)
	fmt.Fprintf(&b, "%sThirdPartyOptIn: %s,\n", indentationValues, s.ThirdPartyOptIn)
	fmt.Fprintf(&b, "%sFirstName: %s,\n", indentationValues, s.FirstName)
	fmt.Fprintf(&b, "%sLastName: %s,\n", indentationValues, s.LastName)
	fmt.Fprintf(&b, "%sPreferredLanguage: %s,\n", indentationValues, s.PreferredLanguage)
	fmt.Fprintf(&b, "%sExternalAccounts: %s,\n", indentationValues, s.ExternalAccounts)
	fmt.Fprintf(&b, "%s}", indentationEnd)

	return b.String()
}

// NewUbiAccount returns a new UbiAccount
func NewUbiAccount() UbiAccount {
	return UbiAccount{
		UbiAccountID:      types.NewString(""),
		Username:          types.NewString(""),
		Password:          types.NewString(""),
		Status:            NewUbiAccountStatus(),
		Email:             types.NewString(""),
		DateOfBirth:       types.NewDateTime(0),
		Gender:            types.NewUInt32(0),
		CountryCode:       types.NewString(""),
		OptIn:             types.NewBool(false),
		ThirdPartyOptIn:   types.NewBool(false),
		FirstName:         types.NewString(""),
		LastName:          types.NewString(""),
		PreferredLanguage: types.NewString(""),
		ExternalAccounts:  types.NewList[ExternalAccount](),
	}
}
