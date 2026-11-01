// Package protocol implements the Tracking 3 protocol of the Quazal Rendez-Vous (Ubisoft) family.
//
// Method IDs and layouts follow the NintendoClients wiki page "Tracking Protocol 3" (MIT).
package protocol

import (
	"fmt"
	"slices"

	nex "github.com/PretendoNetwork/nex-go/v2"
	"github.com/PretendoNetwork/nex-go/v2/types"
	"github.com/PretendoNetwork/nex-protocols-go/v2/globals"
	tracking_3_types "github.com/PretendoNetwork/nex-protocols-go/v2/ubisoft/tracking-3/types"
)

const (
	// ProtocolID is the protocol ID for the Tracking 3 protocol
	ProtocolID = 0x24

	// MethodSendTag is the method ID for method SendTag
	MethodSendTag = 0x1

	// MethodSendTagAndUpdateUserInfo is the method ID for method SendTagAndUpdateUserInfo
	MethodSendTagAndUpdateUserInfo = 0x2

	// MethodSendUserInfo is the method ID for method SendUserInfo
	MethodSendUserInfo = 0x3

	// MethodGetConfiguration is the method ID for method GetConfiguration
	MethodGetConfiguration = 0x4

	// MethodSendTags is the method ID for method SendTags
	MethodSendTags = 0x5
)

// Protocol handles the Tracking 3 protocol
type Protocol struct {
	endpoint                 nex.EndpointInterface
	SendTag                  func(err error, packet nex.PacketInterface, callID uint32, trackingID types.UInt32, tag types.String, attributes types.String, deltaTime types.UInt32) (*nex.RMCMessage, *nex.Error)
	SendTagAndUpdateUserInfo func(err error, packet nex.PacketInterface, callID uint32, trackingID types.UInt32, tag types.String, attributes types.String, deltaTime types.UInt32, userID types.String) (*nex.RMCMessage, *nex.Error)
	SendUserInfo             func(err error, packet nex.PacketInterface, callID uint32, userInfo tracking_3_types.TrackingInformation, deltaTime types.UInt32) (*nex.RMCMessage, *nex.Error)
	GetConfiguration         func(err error, packet nex.PacketInterface, callID uint32) (*nex.RMCMessage, *nex.Error)
	SendTags                 func(err error, packet nex.PacketInterface, callID uint32, tagData types.List[tracking_3_types.TrackingTag]) (*nex.RMCMessage, *nex.Error)
	Patches                  nex.ServiceProtocol
	PatchedMethods           []uint32
}

// Interface implements the methods present on the Tracking 3 protocol struct
type Interface interface {
	Endpoint() nex.EndpointInterface
	SetEndpoint(endpoint nex.EndpointInterface)
	SetHandlerSendTag(handler func(err error, packet nex.PacketInterface, callID uint32, trackingID types.UInt32, tag types.String, attributes types.String, deltaTime types.UInt32) (*nex.RMCMessage, *nex.Error))
	SetHandlerSendTagAndUpdateUserInfo(handler func(err error, packet nex.PacketInterface, callID uint32, trackingID types.UInt32, tag types.String, attributes types.String, deltaTime types.UInt32, userID types.String) (*nex.RMCMessage, *nex.Error))
	SetHandlerSendUserInfo(handler func(err error, packet nex.PacketInterface, callID uint32, userInfo tracking_3_types.TrackingInformation, deltaTime types.UInt32) (*nex.RMCMessage, *nex.Error))
	SetHandlerGetConfiguration(handler func(err error, packet nex.PacketInterface, callID uint32) (*nex.RMCMessage, *nex.Error))
	SetHandlerSendTags(handler func(err error, packet nex.PacketInterface, callID uint32, tagData types.List[tracking_3_types.TrackingTag]) (*nex.RMCMessage, *nex.Error))
}

// Endpoint returns the endpoint implementing the protocol
func (protocol *Protocol) Endpoint() nex.EndpointInterface {
	return protocol.endpoint
}

// SetEndpoint sets the endpoint implementing the protocol
func (protocol *Protocol) SetEndpoint(endpoint nex.EndpointInterface) {
	protocol.endpoint = endpoint
}

// SetHandlerSendTag sets the handler for the SendTag method
func (protocol *Protocol) SetHandlerSendTag(handler func(err error, packet nex.PacketInterface, callID uint32, trackingID types.UInt32, tag types.String, attributes types.String, deltaTime types.UInt32) (*nex.RMCMessage, *nex.Error)) {
	protocol.SendTag = handler
}

// SetHandlerSendTagAndUpdateUserInfo sets the handler for the SendTagAndUpdateUserInfo method
func (protocol *Protocol) SetHandlerSendTagAndUpdateUserInfo(handler func(err error, packet nex.PacketInterface, callID uint32, trackingID types.UInt32, tag types.String, attributes types.String, deltaTime types.UInt32, userID types.String) (*nex.RMCMessage, *nex.Error)) {
	protocol.SendTagAndUpdateUserInfo = handler
}

// SetHandlerSendUserInfo sets the handler for the SendUserInfo method
func (protocol *Protocol) SetHandlerSendUserInfo(handler func(err error, packet nex.PacketInterface, callID uint32, userInfo tracking_3_types.TrackingInformation, deltaTime types.UInt32) (*nex.RMCMessage, *nex.Error)) {
	protocol.SendUserInfo = handler
}

// SetHandlerGetConfiguration sets the handler for the GetConfiguration method
func (protocol *Protocol) SetHandlerGetConfiguration(handler func(err error, packet nex.PacketInterface, callID uint32) (*nex.RMCMessage, *nex.Error)) {
	protocol.GetConfiguration = handler
}

// SetHandlerSendTags sets the handler for the SendTags method
func (protocol *Protocol) SetHandlerSendTags(handler func(err error, packet nex.PacketInterface, callID uint32, tagData types.List[tracking_3_types.TrackingTag]) (*nex.RMCMessage, *nex.Error)) {
	protocol.SendTags = handler
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
	case MethodSendTag:
		protocol.handleSendTag(packet)
	case MethodSendTagAndUpdateUserInfo:
		protocol.handleSendTagAndUpdateUserInfo(packet)
	case MethodSendUserInfo:
		protocol.handleSendUserInfo(packet)
	case MethodGetConfiguration:
		protocol.handleGetConfiguration(packet)
	case MethodSendTags:
		protocol.handleSendTags(packet)
	default:
		errMessage := fmt.Sprintf("Unsupported Tracking 3 method ID: %#v\n", message.MethodID)
		err := nex.NewError(nex.ResultCodes.Core.NotImplemented, errMessage)

		globals.RespondError(packet, ProtocolID, err)
		globals.Logger.Warning(err.Message)
	}
}

// NewProtocol returns a new Tracking 3 protocol
func NewProtocol() *Protocol {
	return &Protocol{}
}
