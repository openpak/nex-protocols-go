// Package protocol implements the User Account Management protocol (Rendez-Vous/Quazal, protocol ID 72).
//
// Method IDs and structures follow the NintendoClients wiki page "User Account Management Protocol".
package protocol

import (
	"fmt"
	"slices"

	nex "github.com/PretendoNetwork/nex-go/v2"
	"github.com/PretendoNetwork/nex-go/v2/types"
	"github.com/PretendoNetwork/nex-protocols-go/v2/globals"
)

const (
	// ProtocolID is the protocol ID for the User Account Management protocol
	ProtocolID = 0x48

	// MethodLookupSceNpIds is the method ID for the method LookupSceNpIds
	MethodLookupSceNpIds = 0x1

	// MethodLookupPrincipalIDs is the method ID for the method LookupPrincipalIDs
	MethodLookupPrincipalIDs = 0x2

	// MethodLookupFirstPartyIds is the method ID for the method LookupFirstPartyIds
	MethodLookupFirstPartyIds = 0x3

	// MethodUserHasPlayed is the method ID for the method UserHasPlayed
	MethodUserHasPlayed = 0x4

	// MethodIsUserPlaying is the method ID for the method IsUserPlaying
	MethodIsUserPlaying = 0x5

	// MethodUpdateSonyAccountInfo is the method ID for the method UpdateSonyAccountInfo
	MethodUpdateSonyAccountInfo = 0x6

	// MethodLookupUsernames is the method ID for the method LookupUsernames
	MethodLookupUsernames = 0x7
)

// Protocol stores all the RMC method handlers for the User Account Management protocol and listens for requests
type Protocol struct {
	endpoint              nex.EndpointInterface
	LookupSceNpIds        func(err error, packet nex.PacketInterface, callID uint32, pids types.List[types.UInt32]) (*nex.RMCMessage, *nex.Error)
	LookupPrincipalIDs    func(err error, packet nex.PacketInterface, callID uint32, firstPartyIds types.List[types.String], platformId types.UInt32) (*nex.RMCMessage, *nex.Error)
	LookupFirstPartyIds   func(err error, packet nex.PacketInterface, callID uint32, pids types.List[types.UInt32], platformId types.UInt32) (*nex.RMCMessage, *nex.Error)
	UserHasPlayed         func(err error, packet nex.PacketInterface, callID uint32, firstPartyIds types.List[types.String], platformId types.UInt32) (*nex.RMCMessage, *nex.Error)
	IsUserPlaying         func(err error, packet nex.PacketInterface, callID uint32, firstPartyIds types.List[types.String], platformId types.UInt32) (*nex.RMCMessage, *nex.Error)
	UpdateSonyAccountInfo func(err error, packet nex.PacketInterface, callID uint32, ticketData types.QBuffer, ticketSize types.UInt32) (*nex.RMCMessage, *nex.Error)
	LookupUsernames       func(err error, packet nex.PacketInterface, callID uint32, pids types.List[types.UInt32], platformId types.UInt32) (*nex.RMCMessage, *nex.Error)
	Patches               nex.ServiceProtocol
	PatchedMethods        []uint32
}

// Interface implements the methods present on the User Account Management protocol struct
type Interface interface {
	Endpoint() nex.EndpointInterface
	SetEndpoint(endpoint nex.EndpointInterface)
	SetHandlerLookupSceNpIds(handler func(err error, packet nex.PacketInterface, callID uint32, pids types.List[types.UInt32]) (*nex.RMCMessage, *nex.Error))
	SetHandlerLookupPrincipalIDs(handler func(err error, packet nex.PacketInterface, callID uint32, firstPartyIds types.List[types.String], platformId types.UInt32) (*nex.RMCMessage, *nex.Error))
	SetHandlerLookupFirstPartyIds(handler func(err error, packet nex.PacketInterface, callID uint32, pids types.List[types.UInt32], platformId types.UInt32) (*nex.RMCMessage, *nex.Error))
	SetHandlerUserHasPlayed(handler func(err error, packet nex.PacketInterface, callID uint32, firstPartyIds types.List[types.String], platformId types.UInt32) (*nex.RMCMessage, *nex.Error))
	SetHandlerIsUserPlaying(handler func(err error, packet nex.PacketInterface, callID uint32, firstPartyIds types.List[types.String], platformId types.UInt32) (*nex.RMCMessage, *nex.Error))
	SetHandlerUpdateSonyAccountInfo(handler func(err error, packet nex.PacketInterface, callID uint32, ticketData types.QBuffer, ticketSize types.UInt32) (*nex.RMCMessage, *nex.Error))
	SetHandlerLookupUsernames(handler func(err error, packet nex.PacketInterface, callID uint32, pids types.List[types.UInt32], platformId types.UInt32) (*nex.RMCMessage, *nex.Error))
}

// Endpoint returns the endpoint implementing the protocol
func (protocol *Protocol) Endpoint() nex.EndpointInterface {
	return protocol.endpoint
}

// SetEndpoint sets the endpoint implementing the protocol
func (protocol *Protocol) SetEndpoint(endpoint nex.EndpointInterface) {
	protocol.endpoint = endpoint
}

// SetHandlerLookupSceNpIds sets the handler for the LookupSceNpIds method
func (protocol *Protocol) SetHandlerLookupSceNpIds(handler func(err error, packet nex.PacketInterface, callID uint32, pids types.List[types.UInt32]) (*nex.RMCMessage, *nex.Error)) {
	protocol.LookupSceNpIds = handler
}

// SetHandlerLookupPrincipalIDs sets the handler for the LookupPrincipalIDs method
func (protocol *Protocol) SetHandlerLookupPrincipalIDs(handler func(err error, packet nex.PacketInterface, callID uint32, firstPartyIds types.List[types.String], platformId types.UInt32) (*nex.RMCMessage, *nex.Error)) {
	protocol.LookupPrincipalIDs = handler
}

// SetHandlerLookupFirstPartyIds sets the handler for the LookupFirstPartyIds method
func (protocol *Protocol) SetHandlerLookupFirstPartyIds(handler func(err error, packet nex.PacketInterface, callID uint32, pids types.List[types.UInt32], platformId types.UInt32) (*nex.RMCMessage, *nex.Error)) {
	protocol.LookupFirstPartyIds = handler
}

// SetHandlerUserHasPlayed sets the handler for the UserHasPlayed method
func (protocol *Protocol) SetHandlerUserHasPlayed(handler func(err error, packet nex.PacketInterface, callID uint32, firstPartyIds types.List[types.String], platformId types.UInt32) (*nex.RMCMessage, *nex.Error)) {
	protocol.UserHasPlayed = handler
}

// SetHandlerIsUserPlaying sets the handler for the IsUserPlaying method
func (protocol *Protocol) SetHandlerIsUserPlaying(handler func(err error, packet nex.PacketInterface, callID uint32, firstPartyIds types.List[types.String], platformId types.UInt32) (*nex.RMCMessage, *nex.Error)) {
	protocol.IsUserPlaying = handler
}

// SetHandlerUpdateSonyAccountInfo sets the handler for the UpdateSonyAccountInfo method
func (protocol *Protocol) SetHandlerUpdateSonyAccountInfo(handler func(err error, packet nex.PacketInterface, callID uint32, ticketData types.QBuffer, ticketSize types.UInt32) (*nex.RMCMessage, *nex.Error)) {
	protocol.UpdateSonyAccountInfo = handler
}

// SetHandlerLookupUsernames sets the handler for the LookupUsernames method
func (protocol *Protocol) SetHandlerLookupUsernames(handler func(err error, packet nex.PacketInterface, callID uint32, pids types.List[types.UInt32], platformId types.UInt32) (*nex.RMCMessage, *nex.Error)) {
	protocol.LookupUsernames = handler
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
	case MethodLookupSceNpIds:
		protocol.handleLookupSceNpIds(packet)
	case MethodLookupPrincipalIDs:
		protocol.handleLookupPrincipalIDs(packet)
	case MethodLookupFirstPartyIds:
		protocol.handleLookupFirstPartyIds(packet)
	case MethodUserHasPlayed:
		protocol.handleUserHasPlayed(packet)
	case MethodIsUserPlaying:
		protocol.handleIsUserPlaying(packet)
	case MethodUpdateSonyAccountInfo:
		protocol.handleUpdateSonyAccountInfo(packet)
	case MethodLookupUsernames:
		protocol.handleLookupUsernames(packet)
	default:
		errMessage := fmt.Sprintf("Unsupported User Account Management method ID: %#v\n", message.MethodID)
		err := nex.NewError(nex.ResultCodes.Core.NotImplemented, errMessage)

		globals.RespondError(packet, ProtocolID, err)
		globals.Logger.Warning(err.Message)
	}
}

// NewProtocol returns a new User Account Management protocol
func NewProtocol() *Protocol {
	return &Protocol{}
}
