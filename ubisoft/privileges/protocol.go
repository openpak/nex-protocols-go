// Package protocol implements the Privileges protocol (Rendez-Vous/Quazal, protocol ID 35).
//
// Method IDs and structures follow the NintendoClients wiki page "Privileges Protocol".
package protocol

import (
	"fmt"
	"slices"

	nex "github.com/PretendoNetwork/nex-go/v2"
	"github.com/PretendoNetwork/nex-go/v2/types"
	"github.com/PretendoNetwork/nex-protocols-go/v2/globals"
)

const (
	// ProtocolID is the protocol ID for the Privileges protocol
	ProtocolID = 0x23

	// MethodGetPrivileges is the method ID for the method GetPrivileges
	MethodGetPrivileges = 0x1

	// MethodActivateKey is the method ID for the method ActivateKey
	MethodActivateKey = 0x2

	// MethodActivateKeyWithExpectedPrivileges is the method ID for the method ActivateKeyWithExpectedPrivileges
	MethodActivateKeyWithExpectedPrivileges = 0x3

	// MethodGetPrivilegeRemainDuration is the method ID for the method GetPrivilegeRemainDuration
	MethodGetPrivilegeRemainDuration = 0x4

	// MethodGetExpiredPrivileges is the method ID for the method GetExpiredPrivileges
	MethodGetExpiredPrivileges = 0x5

	// MethodGetPrivilegesEx is the method ID for the method GetPrivilegesEx
	MethodGetPrivilegesEx = 0x6
)

// Protocol stores all the RMC method handlers for the Privileges protocol and listens for requests
type Protocol struct {
	endpoint                          nex.EndpointInterface
	GetPrivileges                     func(err error, packet nex.PacketInterface, callID uint32, localeCode types.String) (*nex.RMCMessage, *nex.Error)
	ActivateKey                       func(err error, packet nex.PacketInterface, callID uint32, uniqueKey types.String, languageCode types.String) (*nex.RMCMessage, *nex.Error)
	ActivateKeyWithExpectedPrivileges func(err error, packet nex.PacketInterface, callID uint32, uniqueKey types.String, languageCode types.String, expectedPrivilegeIDs types.List[types.UInt32]) (*nex.RMCMessage, *nex.Error)
	GetPrivilegeRemainDuration        func(err error, packet nex.PacketInterface, callID uint32, privilegeID types.UInt32) (*nex.RMCMessage, *nex.Error)
	GetExpiredPrivileges              func(err error, packet nex.PacketInterface, callID uint32) (*nex.RMCMessage, *nex.Error)
	GetPrivilegesEx                   func(err error, packet nex.PacketInterface, callID uint32, localeCode types.String) (*nex.RMCMessage, *nex.Error)
	Patches                           nex.ServiceProtocol
	PatchedMethods                    []uint32
}

// Interface implements the methods present on the Privileges protocol struct
type Interface interface {
	Endpoint() nex.EndpointInterface
	SetEndpoint(endpoint nex.EndpointInterface)
	SetHandlerGetPrivileges(handler func(err error, packet nex.PacketInterface, callID uint32, localeCode types.String) (*nex.RMCMessage, *nex.Error))
	SetHandlerActivateKey(handler func(err error, packet nex.PacketInterface, callID uint32, uniqueKey types.String, languageCode types.String) (*nex.RMCMessage, *nex.Error))
	SetHandlerActivateKeyWithExpectedPrivileges(handler func(err error, packet nex.PacketInterface, callID uint32, uniqueKey types.String, languageCode types.String, expectedPrivilegeIDs types.List[types.UInt32]) (*nex.RMCMessage, *nex.Error))
	SetHandlerGetPrivilegeRemainDuration(handler func(err error, packet nex.PacketInterface, callID uint32, privilegeID types.UInt32) (*nex.RMCMessage, *nex.Error))
	SetHandlerGetExpiredPrivileges(handler func(err error, packet nex.PacketInterface, callID uint32) (*nex.RMCMessage, *nex.Error))
	SetHandlerGetPrivilegesEx(handler func(err error, packet nex.PacketInterface, callID uint32, localeCode types.String) (*nex.RMCMessage, *nex.Error))
}

// Endpoint returns the endpoint implementing the protocol
func (protocol *Protocol) Endpoint() nex.EndpointInterface {
	return protocol.endpoint
}

// SetEndpoint sets the endpoint implementing the protocol
func (protocol *Protocol) SetEndpoint(endpoint nex.EndpointInterface) {
	protocol.endpoint = endpoint
}

// SetHandlerGetPrivileges sets the handler for the GetPrivileges method
func (protocol *Protocol) SetHandlerGetPrivileges(handler func(err error, packet nex.PacketInterface, callID uint32, localeCode types.String) (*nex.RMCMessage, *nex.Error)) {
	protocol.GetPrivileges = handler
}

// SetHandlerActivateKey sets the handler for the ActivateKey method
func (protocol *Protocol) SetHandlerActivateKey(handler func(err error, packet nex.PacketInterface, callID uint32, uniqueKey types.String, languageCode types.String) (*nex.RMCMessage, *nex.Error)) {
	protocol.ActivateKey = handler
}

// SetHandlerActivateKeyWithExpectedPrivileges sets the handler for the ActivateKeyWithExpectedPrivileges method
func (protocol *Protocol) SetHandlerActivateKeyWithExpectedPrivileges(handler func(err error, packet nex.PacketInterface, callID uint32, uniqueKey types.String, languageCode types.String, expectedPrivilegeIDs types.List[types.UInt32]) (*nex.RMCMessage, *nex.Error)) {
	protocol.ActivateKeyWithExpectedPrivileges = handler
}

// SetHandlerGetPrivilegeRemainDuration sets the handler for the GetPrivilegeRemainDuration method
func (protocol *Protocol) SetHandlerGetPrivilegeRemainDuration(handler func(err error, packet nex.PacketInterface, callID uint32, privilegeID types.UInt32) (*nex.RMCMessage, *nex.Error)) {
	protocol.GetPrivilegeRemainDuration = handler
}

// SetHandlerGetExpiredPrivileges sets the handler for the GetExpiredPrivileges method
func (protocol *Protocol) SetHandlerGetExpiredPrivileges(handler func(err error, packet nex.PacketInterface, callID uint32) (*nex.RMCMessage, *nex.Error)) {
	protocol.GetExpiredPrivileges = handler
}

// SetHandlerGetPrivilegesEx sets the handler for the GetPrivilegesEx method
func (protocol *Protocol) SetHandlerGetPrivilegesEx(handler func(err error, packet nex.PacketInterface, callID uint32, localeCode types.String) (*nex.RMCMessage, *nex.Error)) {
	protocol.GetPrivilegesEx = handler
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
	case MethodGetPrivileges:
		protocol.handleGetPrivileges(packet)
	case MethodActivateKey:
		protocol.handleActivateKey(packet)
	case MethodActivateKeyWithExpectedPrivileges:
		protocol.handleActivateKeyWithExpectedPrivileges(packet)
	case MethodGetPrivilegeRemainDuration:
		protocol.handleGetPrivilegeRemainDuration(packet)
	case MethodGetExpiredPrivileges:
		protocol.handleGetExpiredPrivileges(packet)
	case MethodGetPrivilegesEx:
		protocol.handleGetPrivilegesEx(packet)
	default:
		errMessage := fmt.Sprintf("Unsupported Privileges method ID: %#v\n", message.MethodID)
		err := nex.NewError(nex.ResultCodes.Core.NotImplemented, errMessage)

		globals.RespondError(packet, ProtocolID, err)
		globals.Logger.Warning(err.Message)
	}
}

// NewProtocol returns a new Privileges protocol
func NewProtocol() *Protocol {
	return &Protocol{}
}
