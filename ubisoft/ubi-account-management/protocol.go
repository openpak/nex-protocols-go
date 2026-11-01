// Package protocol implements the Ubi Account Management protocol (Rendez-Vous/Quazal, protocol ID 29).
//
// Method IDs and structures follow the NintendoClients wiki page "Ubi Account Management Protocol".
package protocol

import (
	"fmt"
	"slices"

	nex "github.com/PretendoNetwork/nex-go/v2"
	"github.com/PretendoNetwork/nex-go/v2/types"
	"github.com/PretendoNetwork/nex-protocols-go/v2/globals"
	ubi_account_management_types "github.com/PretendoNetwork/nex-protocols-go/v2/ubisoft/ubi-account-management/types"
)

const (
	// ProtocolID is the protocol ID for the Ubi Account Management protocol
	ProtocolID = 0x1d

	// MethodCreateAccount is the method ID for the method CreateAccount
	MethodCreateAccount = 0x1

	// MethodUpdateAccount is the method ID for the method UpdateAccount
	MethodUpdateAccount = 0x2

	// MethodGetAccount is the method ID for the method GetAccount
	MethodGetAccount = 0x3

	// MethodLinkAccount is the method ID for the method LinkAccount
	MethodLinkAccount = 0x4

	// MethodGetTOS is the method ID for the method GetTOS
	MethodGetTOS = 0x5

	// MethodValidateUsername is the method ID for the method ValidateUsername
	MethodValidateUsername = 0x6

	// MethodValidatePassword is the method ID for the method ValidatePassword
	MethodValidatePassword = 0x7

	// MethodValidateEmail is the method ID for the method ValidateEmail
	MethodValidateEmail = 0x8

	// MethodGetCountryList is the method ID for the method GetCountryList
	MethodGetCountryList = 0x9

	// MethodForgetPassword is the method ID for the method ForgetPassword
	MethodForgetPassword = 0xa

	// MethodLookupPrincipalIds is the method ID for the method LookupPrincipalIds
	MethodLookupPrincipalIds = 0xb

	// MethodLookupUbiAccountIDsByPids is the method ID for the method LookupUbiAccountIDsByPids
	MethodLookupUbiAccountIDsByPids = 0xc

	// MethodLookupUbiAccountIDsByUsernames is the method ID for the method LookupUbiAccountIDsByUsernames
	MethodLookupUbiAccountIDsByUsernames = 0xd

	// MethodLookupUsernamesByUbiAccountIDs is the method ID for the method LookupUsernamesByUbiAccountIDs
	MethodLookupUsernamesByUbiAccountIDs = 0xe

	// MethodLookupUbiAccountIDsByUsernameSubString is the method ID for the method LookupUbiAccountIDsByUsernameSubString
	MethodLookupUbiAccountIDsByUsernameSubString = 0xf

	// MethodUserHasPlayed is the method ID for the method UserHasPlayed
	MethodUserHasPlayed = 0x10

	// MethodIsUserPlaying is the method ID for the method IsUserPlaying
	MethodIsUserPlaying = 0x11

	// MethodLookupUbiAccountIDsByUsernamesGlobal is the method ID for the method LookupUbiAccountIDsByUsernamesGlobal
	MethodLookupUbiAccountIDsByUsernamesGlobal = 0x12

	// MethodLookupUbiAccountIDsByEmailsGlobal is the method ID for the method LookupUbiAccountIDsByEmailsGlobal
	MethodLookupUbiAccountIDsByEmailsGlobal = 0x13

	// MethodLookupUsernamesByUbiAccountIDsGlobal is the method ID for the method LookupUsernamesByUbiAccountIDsGlobal
	MethodLookupUsernamesByUbiAccountIDsGlobal = 0x14

	// MethodGetTOSEx is the method ID for the method GetTOSEx
	MethodGetTOSEx = 0x15

	// MethodHasAcceptedLatestTOS is the method ID for the method HasAcceptedLatestTOS
	MethodHasAcceptedLatestTOS = 0x16

	// MethodAcceptLatestTOS is the method ID for the method AcceptLatestTOS
	MethodAcceptLatestTOS = 0x17
)

// Protocol stores all the RMC method handlers for the Ubi Account Management protocol and listens for requests
type Protocol struct {
	endpoint                               nex.EndpointInterface
	CreateAccount                          func(err error, packet nex.PacketInterface, callID uint32, ubiAccount ubi_account_management_types.UbiAccount) (*nex.RMCMessage, *nex.Error)
	UpdateAccount                          func(err error, packet nex.PacketInterface, callID uint32, ubiAccount ubi_account_management_types.UbiAccount) (*nex.RMCMessage, *nex.Error)
	GetAccount                             func(err error, packet nex.PacketInterface, callID uint32) (*nex.RMCMessage, *nex.Error)
	LinkAccount                            func(err error, packet nex.PacketInterface, callID uint32, ubiAccountUsername types.String, ubiAccountPassword types.String) (*nex.RMCMessage, *nex.Error)
	GetTOS                                 func(err error, packet nex.PacketInterface, callID uint32, countryCode types.String, languageCode types.String, htmlVersion types.Bool) (*nex.RMCMessage, *nex.Error)
	ValidateUsername                       func(err error, packet nex.PacketInterface, callID uint32, username types.String) (*nex.RMCMessage, *nex.Error)
	ValidatePassword                       func(err error, packet nex.PacketInterface, callID uint32, password types.String, username types.String) (*nex.RMCMessage, *nex.Error)
	ValidateEmail                          func(err error, packet nex.PacketInterface, callID uint32, email types.String) (*nex.RMCMessage, *nex.Error)
	GetCountryList                         func(err error, packet nex.PacketInterface, callID uint32, languageCode types.String) (*nex.RMCMessage, *nex.Error)
	ForgetPassword                         func(err error, packet nex.PacketInterface, callID uint32, usernameOrEmail types.String) (*nex.RMCMessage, *nex.Error)
	LookupPrincipalIds                     func(err error, packet nex.PacketInterface, callID uint32, ubiAccountIds types.List[types.String]) (*nex.RMCMessage, *nex.Error)
	LookupUbiAccountIDsByPids              func(err error, packet nex.PacketInterface, callID uint32, pids types.List[types.UInt32]) (*nex.RMCMessage, *nex.Error)
	LookupUbiAccountIDsByUsernames         func(err error, packet nex.PacketInterface, callID uint32, usernames types.List[types.String]) (*nex.RMCMessage, *nex.Error)
	LookupUsernamesByUbiAccountIDs         func(err error, packet nex.PacketInterface, callID uint32, ubiAccountIds types.List[types.String]) (*nex.RMCMessage, *nex.Error)
	LookupUbiAccountIDsByUsernameSubString func(err error, packet nex.PacketInterface, callID uint32, usernameSubString types.String) (*nex.RMCMessage, *nex.Error)
	UserHasPlayed                          func(err error, packet nex.PacketInterface, callID uint32, ubiAccountIDs types.List[types.String]) (*nex.RMCMessage, *nex.Error)
	IsUserPlaying                          func(err error, packet nex.PacketInterface, callID uint32, ubiAccountIDs types.List[types.String]) (*nex.RMCMessage, *nex.Error)
	LookupUbiAccountIDsByUsernamesGlobal   func(err error, packet nex.PacketInterface, callID uint32, usernames types.List[types.String]) (*nex.RMCMessage, *nex.Error)
	LookupUbiAccountIDsByEmailsGlobal      func(err error, packet nex.PacketInterface, callID uint32, emails types.List[types.String]) (*nex.RMCMessage, *nex.Error)
	LookupUsernamesByUbiAccountIDsGlobal   func(err error, packet nex.PacketInterface, callID uint32, ubiAccountIds types.List[types.String]) (*nex.RMCMessage, *nex.Error)
	GetTOSEx                               func(err error, packet nex.PacketInterface, callID uint32, countryCode types.String, languageCode types.String, htmlVersion types.Bool) (*nex.RMCMessage, *nex.Error)
	HasAcceptedLatestTOS                   func(err error, packet nex.PacketInterface, callID uint32) (*nex.RMCMessage, *nex.Error)
	AcceptLatestTOS                        func(err error, packet nex.PacketInterface, callID uint32) (*nex.RMCMessage, *nex.Error)
	Patches                                nex.ServiceProtocol
	PatchedMethods                         []uint32
}

// Interface implements the methods present on the Ubi Account Management protocol struct
type Interface interface {
	Endpoint() nex.EndpointInterface
	SetEndpoint(endpoint nex.EndpointInterface)
	SetHandlerCreateAccount(handler func(err error, packet nex.PacketInterface, callID uint32, ubiAccount ubi_account_management_types.UbiAccount) (*nex.RMCMessage, *nex.Error))
	SetHandlerUpdateAccount(handler func(err error, packet nex.PacketInterface, callID uint32, ubiAccount ubi_account_management_types.UbiAccount) (*nex.RMCMessage, *nex.Error))
	SetHandlerGetAccount(handler func(err error, packet nex.PacketInterface, callID uint32) (*nex.RMCMessage, *nex.Error))
	SetHandlerLinkAccount(handler func(err error, packet nex.PacketInterface, callID uint32, ubiAccountUsername types.String, ubiAccountPassword types.String) (*nex.RMCMessage, *nex.Error))
	SetHandlerGetTOS(handler func(err error, packet nex.PacketInterface, callID uint32, countryCode types.String, languageCode types.String, htmlVersion types.Bool) (*nex.RMCMessage, *nex.Error))
	SetHandlerValidateUsername(handler func(err error, packet nex.PacketInterface, callID uint32, username types.String) (*nex.RMCMessage, *nex.Error))
	SetHandlerValidatePassword(handler func(err error, packet nex.PacketInterface, callID uint32, password types.String, username types.String) (*nex.RMCMessage, *nex.Error))
	SetHandlerValidateEmail(handler func(err error, packet nex.PacketInterface, callID uint32, email types.String) (*nex.RMCMessage, *nex.Error))
	SetHandlerGetCountryList(handler func(err error, packet nex.PacketInterface, callID uint32, languageCode types.String) (*nex.RMCMessage, *nex.Error))
	SetHandlerForgetPassword(handler func(err error, packet nex.PacketInterface, callID uint32, usernameOrEmail types.String) (*nex.RMCMessage, *nex.Error))
	SetHandlerLookupPrincipalIds(handler func(err error, packet nex.PacketInterface, callID uint32, ubiAccountIds types.List[types.String]) (*nex.RMCMessage, *nex.Error))
	SetHandlerLookupUbiAccountIDsByPids(handler func(err error, packet nex.PacketInterface, callID uint32, pids types.List[types.UInt32]) (*nex.RMCMessage, *nex.Error))
	SetHandlerLookupUbiAccountIDsByUsernames(handler func(err error, packet nex.PacketInterface, callID uint32, usernames types.List[types.String]) (*nex.RMCMessage, *nex.Error))
	SetHandlerLookupUsernamesByUbiAccountIDs(handler func(err error, packet nex.PacketInterface, callID uint32, ubiAccountIds types.List[types.String]) (*nex.RMCMessage, *nex.Error))
	SetHandlerLookupUbiAccountIDsByUsernameSubString(handler func(err error, packet nex.PacketInterface, callID uint32, usernameSubString types.String) (*nex.RMCMessage, *nex.Error))
	SetHandlerUserHasPlayed(handler func(err error, packet nex.PacketInterface, callID uint32, ubiAccountIDs types.List[types.String]) (*nex.RMCMessage, *nex.Error))
	SetHandlerIsUserPlaying(handler func(err error, packet nex.PacketInterface, callID uint32, ubiAccountIDs types.List[types.String]) (*nex.RMCMessage, *nex.Error))
	SetHandlerLookupUbiAccountIDsByUsernamesGlobal(handler func(err error, packet nex.PacketInterface, callID uint32, usernames types.List[types.String]) (*nex.RMCMessage, *nex.Error))
	SetHandlerLookupUbiAccountIDsByEmailsGlobal(handler func(err error, packet nex.PacketInterface, callID uint32, emails types.List[types.String]) (*nex.RMCMessage, *nex.Error))
	SetHandlerLookupUsernamesByUbiAccountIDsGlobal(handler func(err error, packet nex.PacketInterface, callID uint32, ubiAccountIds types.List[types.String]) (*nex.RMCMessage, *nex.Error))
	SetHandlerGetTOSEx(handler func(err error, packet nex.PacketInterface, callID uint32, countryCode types.String, languageCode types.String, htmlVersion types.Bool) (*nex.RMCMessage, *nex.Error))
	SetHandlerHasAcceptedLatestTOS(handler func(err error, packet nex.PacketInterface, callID uint32) (*nex.RMCMessage, *nex.Error))
	SetHandlerAcceptLatestTOS(handler func(err error, packet nex.PacketInterface, callID uint32) (*nex.RMCMessage, *nex.Error))
}

// Endpoint returns the endpoint implementing the protocol
func (protocol *Protocol) Endpoint() nex.EndpointInterface {
	return protocol.endpoint
}

// SetEndpoint sets the endpoint implementing the protocol
func (protocol *Protocol) SetEndpoint(endpoint nex.EndpointInterface) {
	protocol.endpoint = endpoint
}

// SetHandlerCreateAccount sets the handler for the CreateAccount method
func (protocol *Protocol) SetHandlerCreateAccount(handler func(err error, packet nex.PacketInterface, callID uint32, ubiAccount ubi_account_management_types.UbiAccount) (*nex.RMCMessage, *nex.Error)) {
	protocol.CreateAccount = handler
}

// SetHandlerUpdateAccount sets the handler for the UpdateAccount method
func (protocol *Protocol) SetHandlerUpdateAccount(handler func(err error, packet nex.PacketInterface, callID uint32, ubiAccount ubi_account_management_types.UbiAccount) (*nex.RMCMessage, *nex.Error)) {
	protocol.UpdateAccount = handler
}

// SetHandlerGetAccount sets the handler for the GetAccount method
func (protocol *Protocol) SetHandlerGetAccount(handler func(err error, packet nex.PacketInterface, callID uint32) (*nex.RMCMessage, *nex.Error)) {
	protocol.GetAccount = handler
}

// SetHandlerLinkAccount sets the handler for the LinkAccount method
func (protocol *Protocol) SetHandlerLinkAccount(handler func(err error, packet nex.PacketInterface, callID uint32, ubiAccountUsername types.String, ubiAccountPassword types.String) (*nex.RMCMessage, *nex.Error)) {
	protocol.LinkAccount = handler
}

// SetHandlerGetTOS sets the handler for the GetTOS method
func (protocol *Protocol) SetHandlerGetTOS(handler func(err error, packet nex.PacketInterface, callID uint32, countryCode types.String, languageCode types.String, htmlVersion types.Bool) (*nex.RMCMessage, *nex.Error)) {
	protocol.GetTOS = handler
}

// SetHandlerValidateUsername sets the handler for the ValidateUsername method
func (protocol *Protocol) SetHandlerValidateUsername(handler func(err error, packet nex.PacketInterface, callID uint32, username types.String) (*nex.RMCMessage, *nex.Error)) {
	protocol.ValidateUsername = handler
}

// SetHandlerValidatePassword sets the handler for the ValidatePassword method
func (protocol *Protocol) SetHandlerValidatePassword(handler func(err error, packet nex.PacketInterface, callID uint32, password types.String, username types.String) (*nex.RMCMessage, *nex.Error)) {
	protocol.ValidatePassword = handler
}

// SetHandlerValidateEmail sets the handler for the ValidateEmail method
func (protocol *Protocol) SetHandlerValidateEmail(handler func(err error, packet nex.PacketInterface, callID uint32, email types.String) (*nex.RMCMessage, *nex.Error)) {
	protocol.ValidateEmail = handler
}

// SetHandlerGetCountryList sets the handler for the GetCountryList method
func (protocol *Protocol) SetHandlerGetCountryList(handler func(err error, packet nex.PacketInterface, callID uint32, languageCode types.String) (*nex.RMCMessage, *nex.Error)) {
	protocol.GetCountryList = handler
}

// SetHandlerForgetPassword sets the handler for the ForgetPassword method
func (protocol *Protocol) SetHandlerForgetPassword(handler func(err error, packet nex.PacketInterface, callID uint32, usernameOrEmail types.String) (*nex.RMCMessage, *nex.Error)) {
	protocol.ForgetPassword = handler
}

// SetHandlerLookupPrincipalIds sets the handler for the LookupPrincipalIds method
func (protocol *Protocol) SetHandlerLookupPrincipalIds(handler func(err error, packet nex.PacketInterface, callID uint32, ubiAccountIds types.List[types.String]) (*nex.RMCMessage, *nex.Error)) {
	protocol.LookupPrincipalIds = handler
}

// SetHandlerLookupUbiAccountIDsByPids sets the handler for the LookupUbiAccountIDsByPids method
func (protocol *Protocol) SetHandlerLookupUbiAccountIDsByPids(handler func(err error, packet nex.PacketInterface, callID uint32, pids types.List[types.UInt32]) (*nex.RMCMessage, *nex.Error)) {
	protocol.LookupUbiAccountIDsByPids = handler
}

// SetHandlerLookupUbiAccountIDsByUsernames sets the handler for the LookupUbiAccountIDsByUsernames method
func (protocol *Protocol) SetHandlerLookupUbiAccountIDsByUsernames(handler func(err error, packet nex.PacketInterface, callID uint32, usernames types.List[types.String]) (*nex.RMCMessage, *nex.Error)) {
	protocol.LookupUbiAccountIDsByUsernames = handler
}

// SetHandlerLookupUsernamesByUbiAccountIDs sets the handler for the LookupUsernamesByUbiAccountIDs method
func (protocol *Protocol) SetHandlerLookupUsernamesByUbiAccountIDs(handler func(err error, packet nex.PacketInterface, callID uint32, ubiAccountIds types.List[types.String]) (*nex.RMCMessage, *nex.Error)) {
	protocol.LookupUsernamesByUbiAccountIDs = handler
}

// SetHandlerLookupUbiAccountIDsByUsernameSubString sets the handler for the LookupUbiAccountIDsByUsernameSubString method
func (protocol *Protocol) SetHandlerLookupUbiAccountIDsByUsernameSubString(handler func(err error, packet nex.PacketInterface, callID uint32, usernameSubString types.String) (*nex.RMCMessage, *nex.Error)) {
	protocol.LookupUbiAccountIDsByUsernameSubString = handler
}

// SetHandlerUserHasPlayed sets the handler for the UserHasPlayed method
func (protocol *Protocol) SetHandlerUserHasPlayed(handler func(err error, packet nex.PacketInterface, callID uint32, ubiAccountIDs types.List[types.String]) (*nex.RMCMessage, *nex.Error)) {
	protocol.UserHasPlayed = handler
}

// SetHandlerIsUserPlaying sets the handler for the IsUserPlaying method
func (protocol *Protocol) SetHandlerIsUserPlaying(handler func(err error, packet nex.PacketInterface, callID uint32, ubiAccountIDs types.List[types.String]) (*nex.RMCMessage, *nex.Error)) {
	protocol.IsUserPlaying = handler
}

// SetHandlerLookupUbiAccountIDsByUsernamesGlobal sets the handler for the LookupUbiAccountIDsByUsernamesGlobal method
func (protocol *Protocol) SetHandlerLookupUbiAccountIDsByUsernamesGlobal(handler func(err error, packet nex.PacketInterface, callID uint32, usernames types.List[types.String]) (*nex.RMCMessage, *nex.Error)) {
	protocol.LookupUbiAccountIDsByUsernamesGlobal = handler
}

// SetHandlerLookupUbiAccountIDsByEmailsGlobal sets the handler for the LookupUbiAccountIDsByEmailsGlobal method
func (protocol *Protocol) SetHandlerLookupUbiAccountIDsByEmailsGlobal(handler func(err error, packet nex.PacketInterface, callID uint32, emails types.List[types.String]) (*nex.RMCMessage, *nex.Error)) {
	protocol.LookupUbiAccountIDsByEmailsGlobal = handler
}

// SetHandlerLookupUsernamesByUbiAccountIDsGlobal sets the handler for the LookupUsernamesByUbiAccountIDsGlobal method
func (protocol *Protocol) SetHandlerLookupUsernamesByUbiAccountIDsGlobal(handler func(err error, packet nex.PacketInterface, callID uint32, ubiAccountIds types.List[types.String]) (*nex.RMCMessage, *nex.Error)) {
	protocol.LookupUsernamesByUbiAccountIDsGlobal = handler
}

// SetHandlerGetTOSEx sets the handler for the GetTOSEx method
func (protocol *Protocol) SetHandlerGetTOSEx(handler func(err error, packet nex.PacketInterface, callID uint32, countryCode types.String, languageCode types.String, htmlVersion types.Bool) (*nex.RMCMessage, *nex.Error)) {
	protocol.GetTOSEx = handler
}

// SetHandlerHasAcceptedLatestTOS sets the handler for the HasAcceptedLatestTOS method
func (protocol *Protocol) SetHandlerHasAcceptedLatestTOS(handler func(err error, packet nex.PacketInterface, callID uint32) (*nex.RMCMessage, *nex.Error)) {
	protocol.HasAcceptedLatestTOS = handler
}

// SetHandlerAcceptLatestTOS sets the handler for the AcceptLatestTOS method
func (protocol *Protocol) SetHandlerAcceptLatestTOS(handler func(err error, packet nex.PacketInterface, callID uint32) (*nex.RMCMessage, *nex.Error)) {
	protocol.AcceptLatestTOS = handler
}

// HandlePacket sends the packet to the correct RMC method handler
func (protocol *Protocol) HandlePacket(packet nex.PacketInterface) {
	message := packet.RMCMessage()

	if !message.IsRequest || message.ProtocolID != ProtocolID {
		return
	}

	if protocol.Patches != nil && slices.Contains(protocol.PatchedMethods, message.MethodID) {
		protocol.Patches.HandlePacket(packet)
		return
	}

	switch message.MethodID {
	case MethodCreateAccount:
		protocol.handleCreateAccount(packet)
	case MethodUpdateAccount:
		protocol.handleUpdateAccount(packet)
	case MethodGetAccount:
		protocol.handleGetAccount(packet)
	case MethodLinkAccount:
		protocol.handleLinkAccount(packet)
	case MethodGetTOS:
		protocol.handleGetTOS(packet)
	case MethodValidateUsername:
		protocol.handleValidateUsername(packet)
	case MethodValidatePassword:
		protocol.handleValidatePassword(packet)
	case MethodValidateEmail:
		protocol.handleValidateEmail(packet)
	case MethodGetCountryList:
		protocol.handleGetCountryList(packet)
	case MethodForgetPassword:
		protocol.handleForgetPassword(packet)
	case MethodLookupPrincipalIds:
		protocol.handleLookupPrincipalIds(packet)
	case MethodLookupUbiAccountIDsByPids:
		protocol.handleLookupUbiAccountIDsByPids(packet)
	case MethodLookupUbiAccountIDsByUsernames:
		protocol.handleLookupUbiAccountIDsByUsernames(packet)
	case MethodLookupUsernamesByUbiAccountIDs:
		protocol.handleLookupUsernamesByUbiAccountIDs(packet)
	case MethodLookupUbiAccountIDsByUsernameSubString:
		protocol.handleLookupUbiAccountIDsByUsernameSubString(packet)
	case MethodUserHasPlayed:
		protocol.handleUserHasPlayed(packet)
	case MethodIsUserPlaying:
		protocol.handleIsUserPlaying(packet)
	case MethodLookupUbiAccountIDsByUsernamesGlobal:
		protocol.handleLookupUbiAccountIDsByUsernamesGlobal(packet)
	case MethodLookupUbiAccountIDsByEmailsGlobal:
		protocol.handleLookupUbiAccountIDsByEmailsGlobal(packet)
	case MethodLookupUsernamesByUbiAccountIDsGlobal:
		protocol.handleLookupUsernamesByUbiAccountIDsGlobal(packet)
	case MethodGetTOSEx:
		protocol.handleGetTOSEx(packet)
	case MethodHasAcceptedLatestTOS:
		protocol.handleHasAcceptedLatestTOS(packet)
	case MethodAcceptLatestTOS:
		protocol.handleAcceptLatestTOS(packet)
	default:
		errMessage := fmt.Sprintf("Unsupported Ubi Account Management method ID: %#v\n", message.MethodID)
		err := nex.NewError(nex.ResultCodes.Core.NotImplemented, errMessage)

		globals.RespondError(packet, ProtocolID, err)
		globals.Logger.Warning(err.Message)
	}
}

// NewProtocol returns a new Ubi Account Management protocol
func NewProtocol() *Protocol {
	return &Protocol{}
}
