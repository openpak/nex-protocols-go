// Package protocol implements the Web Notifications Storage protocol of the Quazal Rendez-Vous (Ubisoft) family.
//
// Method IDs and layouts follow the NintendoClients wiki page "Web Notification Storage Protocol" (MIT).
//
// Layout provisional:
//   - The wiki lists the protocol ID as unknown ("?"); ProtocolID is 0 until it is known.
package protocol

import (
	"fmt"
	"slices"

	nex "github.com/PretendoNetwork/nex-go/v2"
	"github.com/PretendoNetwork/nex-protocols-go/v2/globals"
)

const (
	// ProtocolID is the protocol ID for the Web Notifications Storage protocol.
	// Layout provisional: the wiki does not specify the ID, 0 is a placeholder.
	ProtocolID = 0x0

	// MethodRegisterUser is the method ID for method RegisterUser
	MethodRegisterUser = 0x1

	// MethodPollNotifications is the method ID for method PollNotifications
	MethodPollNotifications = 0x2

	// MethodUnregisterUser is the method ID for method UnregisterUser
	MethodUnregisterUser = 0x3
)

// Protocol handles the Web Notifications Storage protocol
type Protocol struct {
	endpoint          nex.EndpointInterface
	RegisterUser      func(err error, packet nex.PacketInterface, callID uint32) (*nex.RMCMessage, *nex.Error)
	PollNotifications func(err error, packet nex.PacketInterface, callID uint32) (*nex.RMCMessage, *nex.Error)
	UnregisterUser    func(err error, packet nex.PacketInterface, callID uint32) (*nex.RMCMessage, *nex.Error)
	Patches           nex.ServiceProtocol
	PatchedMethods    []uint32
}

// Interface implements the methods present on the Web Notifications Storage protocol struct
type Interface interface {
	Endpoint() nex.EndpointInterface
	SetEndpoint(endpoint nex.EndpointInterface)
	SetHandlerRegisterUser(handler func(err error, packet nex.PacketInterface, callID uint32) (*nex.RMCMessage, *nex.Error))
	SetHandlerPollNotifications(handler func(err error, packet nex.PacketInterface, callID uint32) (*nex.RMCMessage, *nex.Error))
	SetHandlerUnregisterUser(handler func(err error, packet nex.PacketInterface, callID uint32) (*nex.RMCMessage, *nex.Error))
}

// Endpoint returns the endpoint implementing the protocol
func (protocol *Protocol) Endpoint() nex.EndpointInterface {
	return protocol.endpoint
}

// SetEndpoint sets the endpoint implementing the protocol
func (protocol *Protocol) SetEndpoint(endpoint nex.EndpointInterface) {
	protocol.endpoint = endpoint
}

// SetHandlerRegisterUser sets the handler for the RegisterUser method
func (protocol *Protocol) SetHandlerRegisterUser(handler func(err error, packet nex.PacketInterface, callID uint32) (*nex.RMCMessage, *nex.Error)) {
	protocol.RegisterUser = handler
}

// SetHandlerPollNotifications sets the handler for the PollNotifications method
func (protocol *Protocol) SetHandlerPollNotifications(handler func(err error, packet nex.PacketInterface, callID uint32) (*nex.RMCMessage, *nex.Error)) {
	protocol.PollNotifications = handler
}

// SetHandlerUnregisterUser sets the handler for the UnregisterUser method
func (protocol *Protocol) SetHandlerUnregisterUser(handler func(err error, packet nex.PacketInterface, callID uint32) (*nex.RMCMessage, *nex.Error)) {
	protocol.UnregisterUser = handler
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
	case MethodRegisterUser:
		protocol.handleRegisterUser(packet)
	case MethodPollNotifications:
		protocol.handlePollNotifications(packet)
	case MethodUnregisterUser:
		protocol.handleUnregisterUser(packet)
	default:
		errMessage := fmt.Sprintf("Unsupported Web Notifications Storage method ID: %#v\n", message.MethodID)
		err := nex.NewError(nex.ResultCodes.Core.NotImplemented, errMessage)

		globals.RespondError(packet, ProtocolID, err)
		globals.Logger.Warning(err.Message)
	}
}

// NewProtocol returns a new Web Notifications Storage protocol
func NewProtocol() *Protocol {
	return &Protocol{}
}
