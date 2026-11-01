// Package protocol implements the Title Storage protocol of the Quazal Rendez-Vous (Ubisoft) family.
//
// Method IDs and layouts follow the NintendoClients wiki page "Title Storage Protocol" (MIT).
//
// Layout provisional:
//   - BrowseContentsByType returns List<any> on the wiki; modelled as a list of AnyDataHolder (DataHolder).
package protocol

import (
	"fmt"
	"slices"

	nex "github.com/PretendoNetwork/nex-go/v2"
	"github.com/PretendoNetwork/nex-go/v2/types"
	"github.com/PretendoNetwork/nex-protocols-go/v2/globals"
)

const (
	// ProtocolID is the protocol ID for the Title Storage protocol
	ProtocolID = 0x33

	// MethodBrowseContents is the method ID for method BrowseContents
	MethodBrowseContents = 0x1

	// MethodBrowseContentsByType is the method ID for method BrowseContentsByType
	MethodBrowseContentsByType = 0x2

	// MethodGetContentDetails is the method ID for method GetContentDetails
	MethodGetContentDetails = 0x3

	// MethodDownloadContent is the method ID for method DownloadContent
	MethodDownloadContent = 0x4
)

// Protocol handles the Title Storage protocol
type Protocol struct {
	endpoint             nex.EndpointInterface
	BrowseContents       func(err error, packet nex.PacketInterface, callID uint32, flags types.UInt32, flagsMaskType types.UInt8, extendedFlags types.UInt32, extendedFlagsMaskType types.UInt8) (*nex.RMCMessage, *nex.Error)
	BrowseContentsByType func(err error, packet nex.PacketInterface, callID uint32, contentType types.String, flags types.UInt32, flagsMaskType types.UInt8, extendedFlags types.UInt32, extendedFlagsMaskType types.UInt8) (*nex.RMCMessage, *nex.Error)
	GetContentDetails    func(err error, packet nex.PacketInterface, callID uint32, contentID types.UInt32) (*nex.RMCMessage, *nex.Error)
	DownloadContent      func(err error, packet nex.PacketInterface, callID uint32, contentID types.UInt32) (*nex.RMCMessage, *nex.Error)
	Patches              nex.ServiceProtocol
	PatchedMethods       []uint32
}

// Interface implements the methods present on the Title Storage protocol struct
type Interface interface {
	Endpoint() nex.EndpointInterface
	SetEndpoint(endpoint nex.EndpointInterface)
	SetHandlerBrowseContents(handler func(err error, packet nex.PacketInterface, callID uint32, flags types.UInt32, flagsMaskType types.UInt8, extendedFlags types.UInt32, extendedFlagsMaskType types.UInt8) (*nex.RMCMessage, *nex.Error))
	SetHandlerBrowseContentsByType(handler func(err error, packet nex.PacketInterface, callID uint32, contentType types.String, flags types.UInt32, flagsMaskType types.UInt8, extendedFlags types.UInt32, extendedFlagsMaskType types.UInt8) (*nex.RMCMessage, *nex.Error))
	SetHandlerGetContentDetails(handler func(err error, packet nex.PacketInterface, callID uint32, contentID types.UInt32) (*nex.RMCMessage, *nex.Error))
	SetHandlerDownloadContent(handler func(err error, packet nex.PacketInterface, callID uint32, contentID types.UInt32) (*nex.RMCMessage, *nex.Error))
}

// Endpoint returns the endpoint implementing the protocol
func (protocol *Protocol) Endpoint() nex.EndpointInterface {
	return protocol.endpoint
}

// SetEndpoint sets the endpoint implementing the protocol
func (protocol *Protocol) SetEndpoint(endpoint nex.EndpointInterface) {
	protocol.endpoint = endpoint
}

// SetHandlerBrowseContents sets the handler for the BrowseContents method
func (protocol *Protocol) SetHandlerBrowseContents(handler func(err error, packet nex.PacketInterface, callID uint32, flags types.UInt32, flagsMaskType types.UInt8, extendedFlags types.UInt32, extendedFlagsMaskType types.UInt8) (*nex.RMCMessage, *nex.Error)) {
	protocol.BrowseContents = handler
}

// SetHandlerBrowseContentsByType sets the handler for the BrowseContentsByType method
func (protocol *Protocol) SetHandlerBrowseContentsByType(handler func(err error, packet nex.PacketInterface, callID uint32, contentType types.String, flags types.UInt32, flagsMaskType types.UInt8, extendedFlags types.UInt32, extendedFlagsMaskType types.UInt8) (*nex.RMCMessage, *nex.Error)) {
	protocol.BrowseContentsByType = handler
}

// SetHandlerGetContentDetails sets the handler for the GetContentDetails method
func (protocol *Protocol) SetHandlerGetContentDetails(handler func(err error, packet nex.PacketInterface, callID uint32, contentID types.UInt32) (*nex.RMCMessage, *nex.Error)) {
	protocol.GetContentDetails = handler
}

// SetHandlerDownloadContent sets the handler for the DownloadContent method
func (protocol *Protocol) SetHandlerDownloadContent(handler func(err error, packet nex.PacketInterface, callID uint32, contentID types.UInt32) (*nex.RMCMessage, *nex.Error)) {
	protocol.DownloadContent = handler
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
	case MethodBrowseContents:
		protocol.handleBrowseContents(packet)
	case MethodBrowseContentsByType:
		protocol.handleBrowseContentsByType(packet)
	case MethodGetContentDetails:
		protocol.handleGetContentDetails(packet)
	case MethodDownloadContent:
		protocol.handleDownloadContent(packet)
	default:
		errMessage := fmt.Sprintf("Unsupported Title Storage method ID: %#v\n", message.MethodID)
		err := nex.NewError(nex.ResultCodes.Core.NotImplemented, errMessage)

		globals.RespondError(packet, ProtocolID, err)
		globals.Logger.Warning(err.Message)
	}
}

// NewProtocol returns a new Title Storage protocol
func NewProtocol() *Protocol {
	return &Protocol{}
}
