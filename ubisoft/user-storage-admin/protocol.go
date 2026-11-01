// Package protocol implements the User Storage Admin protocol of the Quazal Rendez-Vous (Ubisoft) family.
//
// Method IDs and layouts follow the NintendoClients wiki page "User Storage Admin Protocol" (MIT).
//
// Layout provisional:
//   - The wiki gives no protocol ID for this protocol; ProtocolID is 0 until it is known.
//   - UserContent, UserContentKey and ContentProperty are documented on the (unavailable) User Content page; the User Storage page layouts are reused.
package protocol

import (
	"fmt"
	"slices"

	nex "github.com/PretendoNetwork/nex-go/v2"
	"github.com/PretendoNetwork/nex-go/v2/types"
	"github.com/PretendoNetwork/nex-protocols-go/v2/globals"
	user_storage_types "github.com/PretendoNetwork/nex-protocols-go/v2/ubisoft/user-storage/types"
)

const (
	// ProtocolID is the protocol ID for the User Storage Admin protocol.
	// Layout provisional: the wiki does not specify the ID, 0 is a placeholder.
	ProtocolID = 0x0

	// MethodGetContentsToModerate is the method ID for method GetContentsToModerate
	MethodGetContentsToModerate = 0x1

	// MethodFlagContentAsVerified is the method ID for method FlagContentAsVerified
	MethodFlagContentAsVerified = 0x2

	// MethodBanContent is the method ID for method BanContent
	MethodBanContent = 0x3

	// MethodBanUser is the method ID for method BanUser
	MethodBanUser = 0x4

	// MethodBanUserFromContentType is the method ID for method BanUserFromContentType
	MethodBanUserFromContentType = 0x5

	// MethodUnbanUser is the method ID for method UnbanUser
	MethodUnbanUser = 0x6

	// MethodUnbanUserFromContentType is the method ID for method UnbanUserFromContentType
	MethodUnbanUserFromContentType = 0x7

	// MethodGetContentsToModerateWithThreshold is the method ID for method GetContentsToModerateWithThreshold
	MethodGetContentsToModerateWithThreshold = 0x8

	// MethodUpdateMetaData is the method ID for method UpdateMetaData
	MethodUpdateMetaData = 0x9
)

// Protocol handles the User Storage Admin protocol
type Protocol struct {
	endpoint                           nex.EndpointInterface
	GetContentsToModerate              func(err error, packet nex.PacketInterface, callID uint32, typeID types.UInt32, offset types.UInt32, size types.UInt32) (*nex.RMCMessage, *nex.Error)
	FlagContentAsVerified              func(err error, packet nex.PacketInterface, callID uint32, contentKey user_storage_types.UserContentKey) (*nex.RMCMessage, *nex.Error)
	BanContent                         func(err error, packet nex.PacketInterface, callID uint32, contentKey user_storage_types.UserContentKey) (*nex.RMCMessage, *nex.Error)
	BanUser                            func(err error, packet nex.PacketInterface, callID uint32, pid types.UInt32, reason types.String, banContents types.Bool, expireDate types.DateTime) (*nex.RMCMessage, *nex.Error)
	BanUserFromContentType             func(err error, packet nex.PacketInterface, callID uint32, typeID types.UInt32, pid types.UInt32, reason types.String, banContents types.Bool, expireDate types.DateTime) (*nex.RMCMessage, *nex.Error)
	UnbanUser                          func(err error, packet nex.PacketInterface, callID uint32, pid types.UInt32) (*nex.RMCMessage, *nex.Error)
	UnbanUserFromContentType           func(err error, packet nex.PacketInterface, callID uint32, typeID types.UInt32, pid types.UInt32) (*nex.RMCMessage, *nex.Error)
	GetContentsToModerateWithThreshold func(err error, packet nex.PacketInterface, callID uint32, typeID types.UInt32, threshold types.UInt32, offset types.UInt32, size types.UInt32) (*nex.RMCMessage, *nex.Error)
	UpdateMetaData                     func(err error, packet nex.PacketInterface, callID uint32, contentKey user_storage_types.UserContentKey, properties types.List[user_storage_types.ContentProperty]) (*nex.RMCMessage, *nex.Error)
	Patches                            nex.ServiceProtocol
	PatchedMethods                     []uint32
}

// Interface implements the methods present on the User Storage Admin protocol struct
type Interface interface {
	Endpoint() nex.EndpointInterface
	SetEndpoint(endpoint nex.EndpointInterface)
	SetHandlerGetContentsToModerate(handler func(err error, packet nex.PacketInterface, callID uint32, typeID types.UInt32, offset types.UInt32, size types.UInt32) (*nex.RMCMessage, *nex.Error))
	SetHandlerFlagContentAsVerified(handler func(err error, packet nex.PacketInterface, callID uint32, contentKey user_storage_types.UserContentKey) (*nex.RMCMessage, *nex.Error))
	SetHandlerBanContent(handler func(err error, packet nex.PacketInterface, callID uint32, contentKey user_storage_types.UserContentKey) (*nex.RMCMessage, *nex.Error))
	SetHandlerBanUser(handler func(err error, packet nex.PacketInterface, callID uint32, pid types.UInt32, reason types.String, banContents types.Bool, expireDate types.DateTime) (*nex.RMCMessage, *nex.Error))
	SetHandlerBanUserFromContentType(handler func(err error, packet nex.PacketInterface, callID uint32, typeID types.UInt32, pid types.UInt32, reason types.String, banContents types.Bool, expireDate types.DateTime) (*nex.RMCMessage, *nex.Error))
	SetHandlerUnbanUser(handler func(err error, packet nex.PacketInterface, callID uint32, pid types.UInt32) (*nex.RMCMessage, *nex.Error))
	SetHandlerUnbanUserFromContentType(handler func(err error, packet nex.PacketInterface, callID uint32, typeID types.UInt32, pid types.UInt32) (*nex.RMCMessage, *nex.Error))
	SetHandlerGetContentsToModerateWithThreshold(handler func(err error, packet nex.PacketInterface, callID uint32, typeID types.UInt32, threshold types.UInt32, offset types.UInt32, size types.UInt32) (*nex.RMCMessage, *nex.Error))
	SetHandlerUpdateMetaData(handler func(err error, packet nex.PacketInterface, callID uint32, contentKey user_storage_types.UserContentKey, properties types.List[user_storage_types.ContentProperty]) (*nex.RMCMessage, *nex.Error))
}

// Endpoint returns the endpoint implementing the protocol
func (protocol *Protocol) Endpoint() nex.EndpointInterface {
	return protocol.endpoint
}

// SetEndpoint sets the endpoint implementing the protocol
func (protocol *Protocol) SetEndpoint(endpoint nex.EndpointInterface) {
	protocol.endpoint = endpoint
}

// SetHandlerGetContentsToModerate sets the handler for the GetContentsToModerate method
func (protocol *Protocol) SetHandlerGetContentsToModerate(handler func(err error, packet nex.PacketInterface, callID uint32, typeID types.UInt32, offset types.UInt32, size types.UInt32) (*nex.RMCMessage, *nex.Error)) {
	protocol.GetContentsToModerate = handler
}

// SetHandlerFlagContentAsVerified sets the handler for the FlagContentAsVerified method
func (protocol *Protocol) SetHandlerFlagContentAsVerified(handler func(err error, packet nex.PacketInterface, callID uint32, contentKey user_storage_types.UserContentKey) (*nex.RMCMessage, *nex.Error)) {
	protocol.FlagContentAsVerified = handler
}

// SetHandlerBanContent sets the handler for the BanContent method
func (protocol *Protocol) SetHandlerBanContent(handler func(err error, packet nex.PacketInterface, callID uint32, contentKey user_storage_types.UserContentKey) (*nex.RMCMessage, *nex.Error)) {
	protocol.BanContent = handler
}

// SetHandlerBanUser sets the handler for the BanUser method
func (protocol *Protocol) SetHandlerBanUser(handler func(err error, packet nex.PacketInterface, callID uint32, pid types.UInt32, reason types.String, banContents types.Bool, expireDate types.DateTime) (*nex.RMCMessage, *nex.Error)) {
	protocol.BanUser = handler
}

// SetHandlerBanUserFromContentType sets the handler for the BanUserFromContentType method
func (protocol *Protocol) SetHandlerBanUserFromContentType(handler func(err error, packet nex.PacketInterface, callID uint32, typeID types.UInt32, pid types.UInt32, reason types.String, banContents types.Bool, expireDate types.DateTime) (*nex.RMCMessage, *nex.Error)) {
	protocol.BanUserFromContentType = handler
}

// SetHandlerUnbanUser sets the handler for the UnbanUser method
func (protocol *Protocol) SetHandlerUnbanUser(handler func(err error, packet nex.PacketInterface, callID uint32, pid types.UInt32) (*nex.RMCMessage, *nex.Error)) {
	protocol.UnbanUser = handler
}

// SetHandlerUnbanUserFromContentType sets the handler for the UnbanUserFromContentType method
func (protocol *Protocol) SetHandlerUnbanUserFromContentType(handler func(err error, packet nex.PacketInterface, callID uint32, typeID types.UInt32, pid types.UInt32) (*nex.RMCMessage, *nex.Error)) {
	protocol.UnbanUserFromContentType = handler
}

// SetHandlerGetContentsToModerateWithThreshold sets the handler for the GetContentsToModerateWithThreshold method
func (protocol *Protocol) SetHandlerGetContentsToModerateWithThreshold(handler func(err error, packet nex.PacketInterface, callID uint32, typeID types.UInt32, threshold types.UInt32, offset types.UInt32, size types.UInt32) (*nex.RMCMessage, *nex.Error)) {
	protocol.GetContentsToModerateWithThreshold = handler
}

// SetHandlerUpdateMetaData sets the handler for the UpdateMetaData method
func (protocol *Protocol) SetHandlerUpdateMetaData(handler func(err error, packet nex.PacketInterface, callID uint32, contentKey user_storage_types.UserContentKey, properties types.List[user_storage_types.ContentProperty]) (*nex.RMCMessage, *nex.Error)) {
	protocol.UpdateMetaData = handler
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
	case MethodGetContentsToModerate:
		protocol.handleGetContentsToModerate(packet)
	case MethodFlagContentAsVerified:
		protocol.handleFlagContentAsVerified(packet)
	case MethodBanContent:
		protocol.handleBanContent(packet)
	case MethodBanUser:
		protocol.handleBanUser(packet)
	case MethodBanUserFromContentType:
		protocol.handleBanUserFromContentType(packet)
	case MethodUnbanUser:
		protocol.handleUnbanUser(packet)
	case MethodUnbanUserFromContentType:
		protocol.handleUnbanUserFromContentType(packet)
	case MethodGetContentsToModerateWithThreshold:
		protocol.handleGetContentsToModerateWithThreshold(packet)
	case MethodUpdateMetaData:
		protocol.handleUpdateMetaData(packet)
	default:
		errMessage := fmt.Sprintf("Unsupported User Storage Admin method ID: %#v\n", message.MethodID)
		err := nex.NewError(nex.ResultCodes.Core.NotImplemented, errMessage)

		globals.RespondError(packet, ProtocolID, err)
		globals.Logger.Warning(err.Message)
	}
}

// NewProtocol returns a new User Storage Admin protocol
func NewProtocol() *Protocol {
	return &Protocol{}
}
