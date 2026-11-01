// Package protocol implements the News protocol of the Quazal Rendez-Vous (Ubisoft) family.
//
// Method IDs and layouts follow the NintendoClients wiki page "News Protocol" (MIT).
package protocol

import (
	"fmt"
	"slices"

	nex "github.com/PretendoNetwork/nex-go/v2"
	"github.com/PretendoNetwork/nex-go/v2/types"
	"github.com/PretendoNetwork/nex-protocols-go/v2/globals"
	news_types "github.com/PretendoNetwork/nex-protocols-go/v2/ubisoft/news/types"
)

const (
	// ProtocolID is the protocol ID for the News protocol
	ProtocolID = 0x1f

	// MethodGetChannels is the method ID for method GetChannels
	MethodGetChannels = 0x1

	// MethodGetChannelsByTypes is the method ID for method GetChannelsByTypes
	MethodGetChannelsByTypes = 0x2

	// MethodGetSubscribableChannels is the method ID for method GetSubscribableChannels
	MethodGetSubscribableChannels = 0x3

	// MethodGetChannelsByIDs is the method ID for method GetChannelsByIDs
	MethodGetChannelsByIDs = 0x4

	// MethodGetSubscribedChannels is the method ID for method GetSubscribedChannels
	MethodGetSubscribedChannels = 0x5

	// MethodSubscribeChannel is the method ID for method SubscribeChannel
	MethodSubscribeChannel = 0x6

	// MethodUnsubscribeChannel is the method ID for method UnsubscribeChannel
	MethodUnsubscribeChannel = 0x7

	// MethodGetNewsHeaders is the method ID for method GetNewsHeaders
	MethodGetNewsHeaders = 0x8

	// MethodGetNewsMessages is the method ID for method GetNewsMessages
	MethodGetNewsMessages = 0x9

	// MethodGetNumberOfNews is the method ID for method GetNumberOfNews
	MethodGetNumberOfNews = 0xa

	// MethodGetChannelByType is the method ID for method GetChannelByType
	MethodGetChannelByType = 0xb

	// MethodGetNewsHeadersByType is the method ID for method GetNewsHeadersByType
	MethodGetNewsHeadersByType = 0xc

	// MethodGetNewsMessagesByType is the method ID for method GetNewsMessagesByType
	MethodGetNewsMessagesByType = 0xd
)

// Protocol handles the News protocol
type Protocol struct {
	endpoint                nex.EndpointInterface
	GetChannels             func(err error, packet nex.PacketInterface, callID uint32, resultRange types.ResultRange) (*nex.RMCMessage, *nex.Error)
	GetChannelsByTypes      func(err error, packet nex.PacketInterface, callID uint32, newsChannelTypes types.List[types.String], resultRange types.ResultRange) (*nex.RMCMessage, *nex.Error)
	GetSubscribableChannels func(err error, packet nex.PacketInterface, callID uint32, resultRange types.ResultRange) (*nex.RMCMessage, *nex.Error)
	GetChannelsByIDs        func(err error, packet nex.PacketInterface, callID uint32, newsChannelIDs types.List[types.UInt32]) (*nex.RMCMessage, *nex.Error)
	GetSubscribedChannels   func(err error, packet nex.PacketInterface, callID uint32, resultRange types.ResultRange) (*nex.RMCMessage, *nex.Error)
	SubscribeChannel        func(err error, packet nex.PacketInterface, callID uint32, newsChannelID types.UInt32) (*nex.RMCMessage, *nex.Error)
	UnsubscribeChannel      func(err error, packet nex.PacketInterface, callID uint32, newsChannelID types.UInt32) (*nex.RMCMessage, *nex.Error)
	GetNewsHeaders          func(err error, packet nex.PacketInterface, callID uint32, recipient news_types.NewsRecipient, rangeParam types.ResultRange) (*nex.RMCMessage, *nex.Error)
	GetNewsMessages         func(err error, packet nex.PacketInterface, callID uint32, newsMessageIDs types.List[types.UInt32]) (*nex.RMCMessage, *nex.Error)
	GetNumberOfNews         func(err error, packet nex.PacketInterface, callID uint32, recipient news_types.NewsRecipient) (*nex.RMCMessage, *nex.Error)
	GetChannelByType        func(err error, packet nex.PacketInterface, callID uint32, newsChannelType types.String) (*nex.RMCMessage, *nex.Error)
	GetNewsHeadersByType    func(err error, packet nex.PacketInterface, callID uint32, newsChannelType types.String, rangeParam types.ResultRange) (*nex.RMCMessage, *nex.Error)
	GetNewsMessagesByType   func(err error, packet nex.PacketInterface, callID uint32, newsChannelType types.String, rangeParam types.ResultRange) (*nex.RMCMessage, *nex.Error)
	Patches                 nex.ServiceProtocol
	PatchedMethods          []uint32
}

// Interface implements the methods present on the News protocol struct
type Interface interface {
	Endpoint() nex.EndpointInterface
	SetEndpoint(endpoint nex.EndpointInterface)
	SetHandlerGetChannels(handler func(err error, packet nex.PacketInterface, callID uint32, resultRange types.ResultRange) (*nex.RMCMessage, *nex.Error))
	SetHandlerGetChannelsByTypes(handler func(err error, packet nex.PacketInterface, callID uint32, newsChannelTypes types.List[types.String], resultRange types.ResultRange) (*nex.RMCMessage, *nex.Error))
	SetHandlerGetSubscribableChannels(handler func(err error, packet nex.PacketInterface, callID uint32, resultRange types.ResultRange) (*nex.RMCMessage, *nex.Error))
	SetHandlerGetChannelsByIDs(handler func(err error, packet nex.PacketInterface, callID uint32, newsChannelIDs types.List[types.UInt32]) (*nex.RMCMessage, *nex.Error))
	SetHandlerGetSubscribedChannels(handler func(err error, packet nex.PacketInterface, callID uint32, resultRange types.ResultRange) (*nex.RMCMessage, *nex.Error))
	SetHandlerSubscribeChannel(handler func(err error, packet nex.PacketInterface, callID uint32, newsChannelID types.UInt32) (*nex.RMCMessage, *nex.Error))
	SetHandlerUnsubscribeChannel(handler func(err error, packet nex.PacketInterface, callID uint32, newsChannelID types.UInt32) (*nex.RMCMessage, *nex.Error))
	SetHandlerGetNewsHeaders(handler func(err error, packet nex.PacketInterface, callID uint32, recipient news_types.NewsRecipient, rangeParam types.ResultRange) (*nex.RMCMessage, *nex.Error))
	SetHandlerGetNewsMessages(handler func(err error, packet nex.PacketInterface, callID uint32, newsMessageIDs types.List[types.UInt32]) (*nex.RMCMessage, *nex.Error))
	SetHandlerGetNumberOfNews(handler func(err error, packet nex.PacketInterface, callID uint32, recipient news_types.NewsRecipient) (*nex.RMCMessage, *nex.Error))
	SetHandlerGetChannelByType(handler func(err error, packet nex.PacketInterface, callID uint32, newsChannelType types.String) (*nex.RMCMessage, *nex.Error))
	SetHandlerGetNewsHeadersByType(handler func(err error, packet nex.PacketInterface, callID uint32, newsChannelType types.String, rangeParam types.ResultRange) (*nex.RMCMessage, *nex.Error))
	SetHandlerGetNewsMessagesByType(handler func(err error, packet nex.PacketInterface, callID uint32, newsChannelType types.String, rangeParam types.ResultRange) (*nex.RMCMessage, *nex.Error))
}

// Endpoint returns the endpoint implementing the protocol
func (protocol *Protocol) Endpoint() nex.EndpointInterface {
	return protocol.endpoint
}

// SetEndpoint sets the endpoint implementing the protocol
func (protocol *Protocol) SetEndpoint(endpoint nex.EndpointInterface) {
	protocol.endpoint = endpoint
}

// SetHandlerGetChannels sets the handler for the GetChannels method
func (protocol *Protocol) SetHandlerGetChannels(handler func(err error, packet nex.PacketInterface, callID uint32, resultRange types.ResultRange) (*nex.RMCMessage, *nex.Error)) {
	protocol.GetChannels = handler
}

// SetHandlerGetChannelsByTypes sets the handler for the GetChannelsByTypes method
func (protocol *Protocol) SetHandlerGetChannelsByTypes(handler func(err error, packet nex.PacketInterface, callID uint32, newsChannelTypes types.List[types.String], resultRange types.ResultRange) (*nex.RMCMessage, *nex.Error)) {
	protocol.GetChannelsByTypes = handler
}

// SetHandlerGetSubscribableChannels sets the handler for the GetSubscribableChannels method
func (protocol *Protocol) SetHandlerGetSubscribableChannels(handler func(err error, packet nex.PacketInterface, callID uint32, resultRange types.ResultRange) (*nex.RMCMessage, *nex.Error)) {
	protocol.GetSubscribableChannels = handler
}

// SetHandlerGetChannelsByIDs sets the handler for the GetChannelsByIDs method
func (protocol *Protocol) SetHandlerGetChannelsByIDs(handler func(err error, packet nex.PacketInterface, callID uint32, newsChannelIDs types.List[types.UInt32]) (*nex.RMCMessage, *nex.Error)) {
	protocol.GetChannelsByIDs = handler
}

// SetHandlerGetSubscribedChannels sets the handler for the GetSubscribedChannels method
func (protocol *Protocol) SetHandlerGetSubscribedChannels(handler func(err error, packet nex.PacketInterface, callID uint32, resultRange types.ResultRange) (*nex.RMCMessage, *nex.Error)) {
	protocol.GetSubscribedChannels = handler
}

// SetHandlerSubscribeChannel sets the handler for the SubscribeChannel method
func (protocol *Protocol) SetHandlerSubscribeChannel(handler func(err error, packet nex.PacketInterface, callID uint32, newsChannelID types.UInt32) (*nex.RMCMessage, *nex.Error)) {
	protocol.SubscribeChannel = handler
}

// SetHandlerUnsubscribeChannel sets the handler for the UnsubscribeChannel method
func (protocol *Protocol) SetHandlerUnsubscribeChannel(handler func(err error, packet nex.PacketInterface, callID uint32, newsChannelID types.UInt32) (*nex.RMCMessage, *nex.Error)) {
	protocol.UnsubscribeChannel = handler
}

// SetHandlerGetNewsHeaders sets the handler for the GetNewsHeaders method
func (protocol *Protocol) SetHandlerGetNewsHeaders(handler func(err error, packet nex.PacketInterface, callID uint32, recipient news_types.NewsRecipient, rangeParam types.ResultRange) (*nex.RMCMessage, *nex.Error)) {
	protocol.GetNewsHeaders = handler
}

// SetHandlerGetNewsMessages sets the handler for the GetNewsMessages method
func (protocol *Protocol) SetHandlerGetNewsMessages(handler func(err error, packet nex.PacketInterface, callID uint32, newsMessageIDs types.List[types.UInt32]) (*nex.RMCMessage, *nex.Error)) {
	protocol.GetNewsMessages = handler
}

// SetHandlerGetNumberOfNews sets the handler for the GetNumberOfNews method
func (protocol *Protocol) SetHandlerGetNumberOfNews(handler func(err error, packet nex.PacketInterface, callID uint32, recipient news_types.NewsRecipient) (*nex.RMCMessage, *nex.Error)) {
	protocol.GetNumberOfNews = handler
}

// SetHandlerGetChannelByType sets the handler for the GetChannelByType method
func (protocol *Protocol) SetHandlerGetChannelByType(handler func(err error, packet nex.PacketInterface, callID uint32, newsChannelType types.String) (*nex.RMCMessage, *nex.Error)) {
	protocol.GetChannelByType = handler
}

// SetHandlerGetNewsHeadersByType sets the handler for the GetNewsHeadersByType method
func (protocol *Protocol) SetHandlerGetNewsHeadersByType(handler func(err error, packet nex.PacketInterface, callID uint32, newsChannelType types.String, rangeParam types.ResultRange) (*nex.RMCMessage, *nex.Error)) {
	protocol.GetNewsHeadersByType = handler
}

// SetHandlerGetNewsMessagesByType sets the handler for the GetNewsMessagesByType method
func (protocol *Protocol) SetHandlerGetNewsMessagesByType(handler func(err error, packet nex.PacketInterface, callID uint32, newsChannelType types.String, rangeParam types.ResultRange) (*nex.RMCMessage, *nex.Error)) {
	protocol.GetNewsMessagesByType = handler
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
	case MethodGetChannels:
		protocol.handleGetChannels(packet)
	case MethodGetChannelsByTypes:
		protocol.handleGetChannelsByTypes(packet)
	case MethodGetSubscribableChannels:
		protocol.handleGetSubscribableChannels(packet)
	case MethodGetChannelsByIDs:
		protocol.handleGetChannelsByIDs(packet)
	case MethodGetSubscribedChannels:
		protocol.handleGetSubscribedChannels(packet)
	case MethodSubscribeChannel:
		protocol.handleSubscribeChannel(packet)
	case MethodUnsubscribeChannel:
		protocol.handleUnsubscribeChannel(packet)
	case MethodGetNewsHeaders:
		protocol.handleGetNewsHeaders(packet)
	case MethodGetNewsMessages:
		protocol.handleGetNewsMessages(packet)
	case MethodGetNumberOfNews:
		protocol.handleGetNumberOfNews(packet)
	case MethodGetChannelByType:
		protocol.handleGetChannelByType(packet)
	case MethodGetNewsHeadersByType:
		protocol.handleGetNewsHeadersByType(packet)
	case MethodGetNewsMessagesByType:
		protocol.handleGetNewsMessagesByType(packet)
	default:
		errMessage := fmt.Sprintf("Unsupported News method ID: %#v\n", message.MethodID)
		err := nex.NewError(nex.ResultCodes.Core.NotImplemented, errMessage)

		globals.RespondError(packet, ProtocolID, err)
		globals.Logger.Warning(err.Message)
	}
}

// NewProtocol returns a new News protocol
func NewProtocol() *Protocol {
	return &Protocol{}
}
