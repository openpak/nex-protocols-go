// Package protocol implements the Offline Game Notifications protocol of the Quazal Rendez-Vous (Ubisoft) family.
//
// Method IDs and layouts follow the NintendoClients wiki page "Offline Game Notification Protocol" (MIT).
//
// Layout provisional:
//   - NotificationEvent is the nex-protocols-go notifications type; its field widths depend on the endpoint library version.
package protocol

import (
	"fmt"
	"slices"

	nex "github.com/PretendoNetwork/nex-go/v2"
	"github.com/PretendoNetwork/nex-go/v2/types"
	"github.com/PretendoNetwork/nex-protocols-go/v2/globals"
)

const (
	// ProtocolID is the protocol ID for the Offline Game Notifications protocol
	ProtocolID = 0x47

	// MethodPollNotifications is the method ID for method PollNotifications
	MethodPollNotifications = 0x1

	// MethodPollSpecificOfflineNotifications is the method ID for method PollSpecificOfflineNotifications
	MethodPollSpecificOfflineNotifications = 0x2

	// MethodPollAnyOfflineNotifications is the method ID for method PollAnyOfflineNotifications
	MethodPollAnyOfflineNotifications = 0x3
)

// Protocol handles the Offline Game Notifications protocol
type Protocol struct {
	endpoint                         nex.EndpointInterface
	PollNotifications                func(err error, packet nex.PacketInterface, callID uint32) (*nex.RMCMessage, *nex.Error)
	PollSpecificOfflineNotifications func(err error, packet nex.PacketInterface, callID uint32, majortype types.List[types.UInt32]) (*nex.RMCMessage, *nex.Error)
	PollAnyOfflineNotifications      func(err error, packet nex.PacketInterface, callID uint32) (*nex.RMCMessage, *nex.Error)
	Patches                          nex.ServiceProtocol
	PatchedMethods                   []uint32
}

// Interface implements the methods present on the Offline Game Notifications protocol struct
type Interface interface {
	Endpoint() nex.EndpointInterface
	SetEndpoint(endpoint nex.EndpointInterface)
	SetHandlerPollNotifications(handler func(err error, packet nex.PacketInterface, callID uint32) (*nex.RMCMessage, *nex.Error))
	SetHandlerPollSpecificOfflineNotifications(handler func(err error, packet nex.PacketInterface, callID uint32, majortype types.List[types.UInt32]) (*nex.RMCMessage, *nex.Error))
	SetHandlerPollAnyOfflineNotifications(handler func(err error, packet nex.PacketInterface, callID uint32) (*nex.RMCMessage, *nex.Error))
}

// Endpoint returns the endpoint implementing the protocol
func (protocol *Protocol) Endpoint() nex.EndpointInterface {
	return protocol.endpoint
}

// SetEndpoint sets the endpoint implementing the protocol
func (protocol *Protocol) SetEndpoint(endpoint nex.EndpointInterface) {
	protocol.endpoint = endpoint
}

// SetHandlerPollNotifications sets the handler for the PollNotifications method
func (protocol *Protocol) SetHandlerPollNotifications(handler func(err error, packet nex.PacketInterface, callID uint32) (*nex.RMCMessage, *nex.Error)) {
	protocol.PollNotifications = handler
}

// SetHandlerPollSpecificOfflineNotifications sets the handler for the PollSpecificOfflineNotifications method
func (protocol *Protocol) SetHandlerPollSpecificOfflineNotifications(handler func(err error, packet nex.PacketInterface, callID uint32, majortype types.List[types.UInt32]) (*nex.RMCMessage, *nex.Error)) {
	protocol.PollSpecificOfflineNotifications = handler
}

// SetHandlerPollAnyOfflineNotifications sets the handler for the PollAnyOfflineNotifications method
func (protocol *Protocol) SetHandlerPollAnyOfflineNotifications(handler func(err error, packet nex.PacketInterface, callID uint32) (*nex.RMCMessage, *nex.Error)) {
	protocol.PollAnyOfflineNotifications = handler
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
	case MethodPollNotifications:
		protocol.handlePollNotifications(packet)
	case MethodPollSpecificOfflineNotifications:
		protocol.handlePollSpecificOfflineNotifications(packet)
	case MethodPollAnyOfflineNotifications:
		protocol.handlePollAnyOfflineNotifications(packet)
	default:
		errMessage := fmt.Sprintf("Unsupported Offline Game Notifications method ID: %#v\n", message.MethodID)
		err := nex.NewError(nex.ResultCodes.Core.NotImplemented, errMessage)

		globals.RespondError(packet, ProtocolID, err)
		globals.Logger.Warning(err.Message)
	}
}

// NewProtocol returns a new Offline Game Notifications protocol
func NewProtocol() *Protocol {
	return &Protocol{}
}
