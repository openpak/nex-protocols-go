// Package protocol implements the Title Storage Admin protocol of the Quazal Rendez-Vous (Ubisoft) family.
//
// Method IDs and layouts follow the NintendoClients wiki page "Title Storage Admin Protocol" (MIT).
//
// Layout provisional:
//   - The wiki gives no protocol ID for this protocol; ProtocolID is 0 until it is known.
package protocol

import (
	"fmt"
	"slices"

	nex "github.com/PretendoNetwork/nex-go/v2"
	"github.com/PretendoNetwork/nex-go/v2/types"
	"github.com/PretendoNetwork/nex-protocols-go/v2/globals"
)

const (
	// ProtocolID is the protocol ID for the Title Storage Admin protocol.
	// Layout provisional: the wiki does not specify the ID, 0 is a placeholder.
	ProtocolID = 0x0

	// MethodGetContentTypesDescription is the method ID for method GetContentTypesDescription
	MethodGetContentTypesDescription = 0x1

	// MethodGetContentTypeDescription is the method ID for method GetContentTypeDescription
	MethodGetContentTypeDescription = 0x2

	// MethodBrowseContents is the method ID for method BrowseContents
	MethodBrowseContents = 0x3

	// MethodGetContentDetails is the method ID for method GetContentDetails
	MethodGetContentDetails = 0x4

	// MethodCreateContentAndGetUploadInformation is the method ID for method CreateContentAndGetUploadInformation
	MethodCreateContentAndGetUploadInformation = 0x5

	// MethodRemoveContent is the method ID for method RemoveContent
	MethodRemoveContent = 0x6

	// MethodUpdateContentMetaData is the method ID for method UpdateContentMetaData
	MethodUpdateContentMetaData = 0x7

	// MethodContentInsertionCompleted is the method ID for method ContentInsertionCompleted
	MethodContentInsertionCompleted = 0x8
)

// Protocol handles the Title Storage Admin protocol
type Protocol struct {
	endpoint                             nex.EndpointInterface
	GetContentTypesDescription           func(err error, packet nex.PacketInterface, callID uint32) (*nex.RMCMessage, *nex.Error)
	GetContentTypeDescription            func(err error, packet nex.PacketInterface, callID uint32, contentType types.String) (*nex.RMCMessage, *nex.Error)
	BrowseContents                       func(err error, packet nex.PacketInterface, callID uint32) (*nex.RMCMessage, *nex.Error)
	GetContentDetails                    func(err error, packet nex.PacketInterface, callID uint32, contentId types.UInt32) (*nex.RMCMessage, *nex.Error)
	CreateContentAndGetUploadInformation func(err error, packet nex.PacketInterface, callID uint32, name types.String, description types.String, key types.String, dateAvailable types.DateTime, dateExpired types.DateTime, flags types.UInt32, extendedFlags types.UInt32, contentType types.String, extension types.DataHolder, mime types.String) (*nex.RMCMessage, *nex.Error)
	RemoveContent                        func(err error, packet nex.PacketInterface, callID uint32, contentId types.UInt32) (*nex.RMCMessage, *nex.Error)
	UpdateContentMetaData                func(err error, packet nex.PacketInterface, callID uint32, contentId types.UInt32, name types.String, description types.String, dateAvailable types.DateTime, dateExpired types.DateTime, flags types.UInt32, extendedFlags types.UInt32, extension types.DataHolder) (*nex.RMCMessage, *nex.Error)
	ContentInsertionCompleted            func(err error, packet nex.PacketInterface, callID uint32, contentId types.UInt32, contentSize types.UInt32) (*nex.RMCMessage, *nex.Error)
	Patches                              nex.ServiceProtocol
	PatchedMethods                       []uint32
}

// Interface implements the methods present on the Title Storage Admin protocol struct
type Interface interface {
	Endpoint() nex.EndpointInterface
	SetEndpoint(endpoint nex.EndpointInterface)
	SetHandlerGetContentTypesDescription(handler func(err error, packet nex.PacketInterface, callID uint32) (*nex.RMCMessage, *nex.Error))
	SetHandlerGetContentTypeDescription(handler func(err error, packet nex.PacketInterface, callID uint32, contentType types.String) (*nex.RMCMessage, *nex.Error))
	SetHandlerBrowseContents(handler func(err error, packet nex.PacketInterface, callID uint32) (*nex.RMCMessage, *nex.Error))
	SetHandlerGetContentDetails(handler func(err error, packet nex.PacketInterface, callID uint32, contentId types.UInt32) (*nex.RMCMessage, *nex.Error))
	SetHandlerCreateContentAndGetUploadInformation(handler func(err error, packet nex.PacketInterface, callID uint32, name types.String, description types.String, key types.String, dateAvailable types.DateTime, dateExpired types.DateTime, flags types.UInt32, extendedFlags types.UInt32, contentType types.String, extension types.DataHolder, mime types.String) (*nex.RMCMessage, *nex.Error))
	SetHandlerRemoveContent(handler func(err error, packet nex.PacketInterface, callID uint32, contentId types.UInt32) (*nex.RMCMessage, *nex.Error))
	SetHandlerUpdateContentMetaData(handler func(err error, packet nex.PacketInterface, callID uint32, contentId types.UInt32, name types.String, description types.String, dateAvailable types.DateTime, dateExpired types.DateTime, flags types.UInt32, extendedFlags types.UInt32, extension types.DataHolder) (*nex.RMCMessage, *nex.Error))
	SetHandlerContentInsertionCompleted(handler func(err error, packet nex.PacketInterface, callID uint32, contentId types.UInt32, contentSize types.UInt32) (*nex.RMCMessage, *nex.Error))
}

// Endpoint returns the endpoint implementing the protocol
func (protocol *Protocol) Endpoint() nex.EndpointInterface {
	return protocol.endpoint
}

// SetEndpoint sets the endpoint implementing the protocol
func (protocol *Protocol) SetEndpoint(endpoint nex.EndpointInterface) {
	protocol.endpoint = endpoint
}

// SetHandlerGetContentTypesDescription sets the handler for the GetContentTypesDescription method
func (protocol *Protocol) SetHandlerGetContentTypesDescription(handler func(err error, packet nex.PacketInterface, callID uint32) (*nex.RMCMessage, *nex.Error)) {
	protocol.GetContentTypesDescription = handler
}

// SetHandlerGetContentTypeDescription sets the handler for the GetContentTypeDescription method
func (protocol *Protocol) SetHandlerGetContentTypeDescription(handler func(err error, packet nex.PacketInterface, callID uint32, contentType types.String) (*nex.RMCMessage, *nex.Error)) {
	protocol.GetContentTypeDescription = handler
}

// SetHandlerBrowseContents sets the handler for the BrowseContents method
func (protocol *Protocol) SetHandlerBrowseContents(handler func(err error, packet nex.PacketInterface, callID uint32) (*nex.RMCMessage, *nex.Error)) {
	protocol.BrowseContents = handler
}

// SetHandlerGetContentDetails sets the handler for the GetContentDetails method
func (protocol *Protocol) SetHandlerGetContentDetails(handler func(err error, packet nex.PacketInterface, callID uint32, contentId types.UInt32) (*nex.RMCMessage, *nex.Error)) {
	protocol.GetContentDetails = handler
}

// SetHandlerCreateContentAndGetUploadInformation sets the handler for the CreateContentAndGetUploadInformation method
func (protocol *Protocol) SetHandlerCreateContentAndGetUploadInformation(handler func(err error, packet nex.PacketInterface, callID uint32, name types.String, description types.String, key types.String, dateAvailable types.DateTime, dateExpired types.DateTime, flags types.UInt32, extendedFlags types.UInt32, contentType types.String, extension types.DataHolder, mime types.String) (*nex.RMCMessage, *nex.Error)) {
	protocol.CreateContentAndGetUploadInformation = handler
}

// SetHandlerRemoveContent sets the handler for the RemoveContent method
func (protocol *Protocol) SetHandlerRemoveContent(handler func(err error, packet nex.PacketInterface, callID uint32, contentId types.UInt32) (*nex.RMCMessage, *nex.Error)) {
	protocol.RemoveContent = handler
}

// SetHandlerUpdateContentMetaData sets the handler for the UpdateContentMetaData method
func (protocol *Protocol) SetHandlerUpdateContentMetaData(handler func(err error, packet nex.PacketInterface, callID uint32, contentId types.UInt32, name types.String, description types.String, dateAvailable types.DateTime, dateExpired types.DateTime, flags types.UInt32, extendedFlags types.UInt32, extension types.DataHolder) (*nex.RMCMessage, *nex.Error)) {
	protocol.UpdateContentMetaData = handler
}

// SetHandlerContentInsertionCompleted sets the handler for the ContentInsertionCompleted method
func (protocol *Protocol) SetHandlerContentInsertionCompleted(handler func(err error, packet nex.PacketInterface, callID uint32, contentId types.UInt32, contentSize types.UInt32) (*nex.RMCMessage, *nex.Error)) {
	protocol.ContentInsertionCompleted = handler
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
	case MethodGetContentTypesDescription:
		protocol.handleGetContentTypesDescription(packet)
	case MethodGetContentTypeDescription:
		protocol.handleGetContentTypeDescription(packet)
	case MethodBrowseContents:
		protocol.handleBrowseContents(packet)
	case MethodGetContentDetails:
		protocol.handleGetContentDetails(packet)
	case MethodCreateContentAndGetUploadInformation:
		protocol.handleCreateContentAndGetUploadInformation(packet)
	case MethodRemoveContent:
		protocol.handleRemoveContent(packet)
	case MethodUpdateContentMetaData:
		protocol.handleUpdateContentMetaData(packet)
	case MethodContentInsertionCompleted:
		protocol.handleContentInsertionCompleted(packet)
	default:
		errMessage := fmt.Sprintf("Unsupported Title Storage Admin method ID: %#v\n", message.MethodID)
		err := nex.NewError(nex.ResultCodes.Core.NotImplemented, errMessage)

		globals.RespondError(packet, ProtocolID, err)
		globals.Logger.Warning(err.Message)
	}
}

// NewProtocol returns a new Title Storage Admin protocol
func NewProtocol() *Protocol {
	return &Protocol{}
}
