// Package protocol implements the User Storage protocol of the Quazal Rendez-Vous (Ubisoft) family.
//
// Method IDs and layouts follow the NintendoClients wiki page "User Storage Protocol" (MIT).
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
	// ProtocolID is the protocol ID for the User Storage protocol
	ProtocolID = 0x35

	// MethodSearchContents is the method ID for method SearchContents
	MethodSearchContents = 0x1

	// MethodSearchContentsWithTotal is the method ID for method SearchContentsWithTotal
	MethodSearchContentsWithTotal = 0x2

	// MethodDeleteContent is the method ID for method DeleteContent
	MethodDeleteContent = 0x3

	// MethodSaveMetaData is the method ID for method SaveMetaData
	MethodSaveMetaData = 0x4

	// MethodSaveContentDB is the method ID for method SaveContentDB
	MethodSaveContentDB = 0x5

	// MethodSaveContentAndGetUploadInfo is the method ID for method SaveContentAndGetUploadInfo
	MethodSaveContentAndGetUploadInfo = 0x6

	// MethodUploadEnd is the method ID for method UploadEnd
	MethodUploadEnd = 0x7

	// MethodGetContentDB is the method ID for method GetContentDB
	MethodGetContentDB = 0x8

	// MethodGetContentURL is the method ID for method GetContentURL
	MethodGetContentURL = 0x9

	// MethodGetSlotCount is the method ID for method GetSlotCount
	MethodGetSlotCount = 0xa

	// MethodGetMetaData is the method ID for method GetMetaData
	MethodGetMetaData = 0xb

	// MethodLike is the method ID for method Like
	MethodLike = 0xc

	// MethodUnlike is the method ID for method Unlike
	MethodUnlike = 0xd

	// MethodIsLiked is the method ID for method IsLiked
	MethodIsLiked = 0xe

	// MethodGetFavourites is the method ID for method GetFavourites
	MethodGetFavourites = 0xf

	// MethodMakeFavourite is the method ID for method MakeFavourite
	MethodMakeFavourite = 0x10

	// MethodRemoveFromFavourites is the method ID for method RemoveFromFavourites
	MethodRemoveFromFavourites = 0x11

	// MethodReportInappropriate is the method ID for method ReportInappropriate
	MethodReportInappropriate = 0x12

	// MethodIncrementPlayCount is the method ID for method IncrementPlayCount
	MethodIncrementPlayCount = 0x13

	// MethodUpdateCustomStat is the method ID for method UpdateCustomStat
	MethodUpdateCustomStat = 0x14

	// MethodGetOwnContents is the method ID for method GetOwnContents
	MethodGetOwnContents = 0x15

	// MethodGetMostPopularTags is the method ID for method GetMostPopularTags
	MethodGetMostPopularTags = 0x16

	// MethodGetTags is the method ID for method GetTags
	MethodGetTags = 0x17

	// MethodTagContent is the method ID for method TagContent
	MethodTagContent = 0x18

	// MethodSearchContentsByPlayers is the method ID for method SearchContentsByPlayers
	MethodSearchContentsByPlayers = 0x19

	// MethodSearchContentsByPlayersWithTotal is the method ID for method SearchContentsByPlayersWithTotal
	MethodSearchContentsByPlayersWithTotal = 0x1a
)

// Protocol handles the User Storage protocol
type Protocol struct {
	endpoint                         nex.EndpointInterface
	SearchContents                   func(err error, packet nex.PacketInterface, callID uint32, query user_storage_types.UserStorageQuery) (*nex.RMCMessage, *nex.Error)
	SearchContentsWithTotal          func(err error, packet nex.PacketInterface, callID uint32, query user_storage_types.UserStorageQuery) (*nex.RMCMessage, *nex.Error)
	DeleteContent                    func(err error, packet nex.PacketInterface, callID uint32, contentKey user_storage_types.UserContentKey) (*nex.RMCMessage, *nex.Error)
	SaveMetaData                     func(err error, packet nex.PacketInterface, callID uint32, properties types.List[user_storage_types.ContentProperty], contentKey user_storage_types.UserContentKey) (*nex.RMCMessage, *nex.Error)
	SaveContentDB                    func(err error, packet nex.PacketInterface, callID uint32, properties types.List[user_storage_types.ContentProperty], data types.Buffer, contentKey user_storage_types.UserContentKey) (*nex.RMCMessage, *nex.Error)
	SaveContentAndGetUploadInfo      func(err error, packet nex.PacketInterface, callID uint32, properties types.List[user_storage_types.ContentProperty], size types.UInt32, contentKey user_storage_types.UserContentKey) (*nex.RMCMessage, *nex.Error)
	UploadEnd                        func(err error, packet nex.PacketInterface, callID uint32, pendingID types.UInt64, result types.Bool, contentKey user_storage_types.UserContentKey) (*nex.RMCMessage, *nex.Error)
	GetContentDB                     func(err error, packet nex.PacketInterface, callID uint32, contentKey user_storage_types.UserContentKey) (*nex.RMCMessage, *nex.Error)
	GetContentURL                    func(err error, packet nex.PacketInterface, callID uint32, contentKey user_storage_types.UserContentKey) (*nex.RMCMessage, *nex.Error)
	GetSlotCount                     func(err error, packet nex.PacketInterface, callID uint32, typeID types.UInt32) (*nex.RMCMessage, *nex.Error)
	GetMetaData                      func(err error, packet nex.PacketInterface, callID uint32, contentKey user_storage_types.UserContentKey) (*nex.RMCMessage, *nex.Error)
	Like                             func(err error, packet nex.PacketInterface, callID uint32, contentKey user_storage_types.UserContentKey) (*nex.RMCMessage, *nex.Error)
	Unlike                           func(err error, packet nex.PacketInterface, callID uint32, contentKey user_storage_types.UserContentKey) (*nex.RMCMessage, *nex.Error)
	IsLiked                          func(err error, packet nex.PacketInterface, callID uint32, contentKey user_storage_types.UserContentKey) (*nex.RMCMessage, *nex.Error)
	GetFavourites                    func(err error, packet nex.PacketInterface, callID uint32, contentTypes types.List[types.UInt32]) (*nex.RMCMessage, *nex.Error)
	MakeFavourite                    func(err error, packet nex.PacketInterface, callID uint32, contentKey user_storage_types.UserContentKey) (*nex.RMCMessage, *nex.Error)
	RemoveFromFavourites             func(err error, packet nex.PacketInterface, callID uint32, contentKey user_storage_types.UserContentKey) (*nex.RMCMessage, *nex.Error)
	ReportInappropriate              func(err error, packet nex.PacketInterface, callID uint32, contentKey user_storage_types.UserContentKey, reason types.String) (*nex.RMCMessage, *nex.Error)
	IncrementPlayCount               func(err error, packet nex.PacketInterface, callID uint32, contentKey user_storage_types.UserContentKey) (*nex.RMCMessage, *nex.Error)
	UpdateCustomStat                 func(err error, packet nex.PacketInterface, callID uint32, contentKey user_storage_types.UserContentKey, statID types.UInt16, incValue types.Int64) (*nex.RMCMessage, *nex.Error)
	GetOwnContents                   func(err error, packet nex.PacketInterface, callID uint32, typeID types.UInt32) (*nex.RMCMessage, *nex.Error)
	GetMostPopularTags               func(err error, packet nex.PacketInterface, callID uint32, contentKey user_storage_types.UserContentKey) (*nex.RMCMessage, *nex.Error)
	GetTags                          func(err error, packet nex.PacketInterface, callID uint32, contentKey user_storage_types.UserContentKey) (*nex.RMCMessage, *nex.Error)
	TagContent                       func(err error, packet nex.PacketInterface, callID uint32, contentKey user_storage_types.UserContentKey, newTagIds types.List[types.UInt32]) (*nex.RMCMessage, *nex.Error)
	SearchContentsByPlayers          func(err error, packet nex.PacketInterface, callID uint32, pids types.List[types.UInt32], query user_storage_types.UserStorageQuery) (*nex.RMCMessage, *nex.Error)
	SearchContentsByPlayersWithTotal func(err error, packet nex.PacketInterface, callID uint32, pids types.List[types.UInt32], query user_storage_types.UserStorageQuery) (*nex.RMCMessage, *nex.Error)
	Patches                          nex.ServiceProtocol
	PatchedMethods                   []uint32
}

// Interface implements the methods present on the User Storage protocol struct
type Interface interface {
	Endpoint() nex.EndpointInterface
	SetEndpoint(endpoint nex.EndpointInterface)
	SetHandlerSearchContents(handler func(err error, packet nex.PacketInterface, callID uint32, query user_storage_types.UserStorageQuery) (*nex.RMCMessage, *nex.Error))
	SetHandlerSearchContentsWithTotal(handler func(err error, packet nex.PacketInterface, callID uint32, query user_storage_types.UserStorageQuery) (*nex.RMCMessage, *nex.Error))
	SetHandlerDeleteContent(handler func(err error, packet nex.PacketInterface, callID uint32, contentKey user_storage_types.UserContentKey) (*nex.RMCMessage, *nex.Error))
	SetHandlerSaveMetaData(handler func(err error, packet nex.PacketInterface, callID uint32, properties types.List[user_storage_types.ContentProperty], contentKey user_storage_types.UserContentKey) (*nex.RMCMessage, *nex.Error))
	SetHandlerSaveContentDB(handler func(err error, packet nex.PacketInterface, callID uint32, properties types.List[user_storage_types.ContentProperty], data types.Buffer, contentKey user_storage_types.UserContentKey) (*nex.RMCMessage, *nex.Error))
	SetHandlerSaveContentAndGetUploadInfo(handler func(err error, packet nex.PacketInterface, callID uint32, properties types.List[user_storage_types.ContentProperty], size types.UInt32, contentKey user_storage_types.UserContentKey) (*nex.RMCMessage, *nex.Error))
	SetHandlerUploadEnd(handler func(err error, packet nex.PacketInterface, callID uint32, pendingID types.UInt64, result types.Bool, contentKey user_storage_types.UserContentKey) (*nex.RMCMessage, *nex.Error))
	SetHandlerGetContentDB(handler func(err error, packet nex.PacketInterface, callID uint32, contentKey user_storage_types.UserContentKey) (*nex.RMCMessage, *nex.Error))
	SetHandlerGetContentURL(handler func(err error, packet nex.PacketInterface, callID uint32, contentKey user_storage_types.UserContentKey) (*nex.RMCMessage, *nex.Error))
	SetHandlerGetSlotCount(handler func(err error, packet nex.PacketInterface, callID uint32, typeID types.UInt32) (*nex.RMCMessage, *nex.Error))
	SetHandlerGetMetaData(handler func(err error, packet nex.PacketInterface, callID uint32, contentKey user_storage_types.UserContentKey) (*nex.RMCMessage, *nex.Error))
	SetHandlerLike(handler func(err error, packet nex.PacketInterface, callID uint32, contentKey user_storage_types.UserContentKey) (*nex.RMCMessage, *nex.Error))
	SetHandlerUnlike(handler func(err error, packet nex.PacketInterface, callID uint32, contentKey user_storage_types.UserContentKey) (*nex.RMCMessage, *nex.Error))
	SetHandlerIsLiked(handler func(err error, packet nex.PacketInterface, callID uint32, contentKey user_storage_types.UserContentKey) (*nex.RMCMessage, *nex.Error))
	SetHandlerGetFavourites(handler func(err error, packet nex.PacketInterface, callID uint32, contentTypes types.List[types.UInt32]) (*nex.RMCMessage, *nex.Error))
	SetHandlerMakeFavourite(handler func(err error, packet nex.PacketInterface, callID uint32, contentKey user_storage_types.UserContentKey) (*nex.RMCMessage, *nex.Error))
	SetHandlerRemoveFromFavourites(handler func(err error, packet nex.PacketInterface, callID uint32, contentKey user_storage_types.UserContentKey) (*nex.RMCMessage, *nex.Error))
	SetHandlerReportInappropriate(handler func(err error, packet nex.PacketInterface, callID uint32, contentKey user_storage_types.UserContentKey, reason types.String) (*nex.RMCMessage, *nex.Error))
	SetHandlerIncrementPlayCount(handler func(err error, packet nex.PacketInterface, callID uint32, contentKey user_storage_types.UserContentKey) (*nex.RMCMessage, *nex.Error))
	SetHandlerUpdateCustomStat(handler func(err error, packet nex.PacketInterface, callID uint32, contentKey user_storage_types.UserContentKey, statID types.UInt16, incValue types.Int64) (*nex.RMCMessage, *nex.Error))
	SetHandlerGetOwnContents(handler func(err error, packet nex.PacketInterface, callID uint32, typeID types.UInt32) (*nex.RMCMessage, *nex.Error))
	SetHandlerGetMostPopularTags(handler func(err error, packet nex.PacketInterface, callID uint32, contentKey user_storage_types.UserContentKey) (*nex.RMCMessage, *nex.Error))
	SetHandlerGetTags(handler func(err error, packet nex.PacketInterface, callID uint32, contentKey user_storage_types.UserContentKey) (*nex.RMCMessage, *nex.Error))
	SetHandlerTagContent(handler func(err error, packet nex.PacketInterface, callID uint32, contentKey user_storage_types.UserContentKey, newTagIds types.List[types.UInt32]) (*nex.RMCMessage, *nex.Error))
	SetHandlerSearchContentsByPlayers(handler func(err error, packet nex.PacketInterface, callID uint32, pids types.List[types.UInt32], query user_storage_types.UserStorageQuery) (*nex.RMCMessage, *nex.Error))
	SetHandlerSearchContentsByPlayersWithTotal(handler func(err error, packet nex.PacketInterface, callID uint32, pids types.List[types.UInt32], query user_storage_types.UserStorageQuery) (*nex.RMCMessage, *nex.Error))
}

// Endpoint returns the endpoint implementing the protocol
func (protocol *Protocol) Endpoint() nex.EndpointInterface {
	return protocol.endpoint
}

// SetEndpoint sets the endpoint implementing the protocol
func (protocol *Protocol) SetEndpoint(endpoint nex.EndpointInterface) {
	protocol.endpoint = endpoint
}

// SetHandlerSearchContents sets the handler for the SearchContents method
func (protocol *Protocol) SetHandlerSearchContents(handler func(err error, packet nex.PacketInterface, callID uint32, query user_storage_types.UserStorageQuery) (*nex.RMCMessage, *nex.Error)) {
	protocol.SearchContents = handler
}

// SetHandlerSearchContentsWithTotal sets the handler for the SearchContentsWithTotal method
func (protocol *Protocol) SetHandlerSearchContentsWithTotal(handler func(err error, packet nex.PacketInterface, callID uint32, query user_storage_types.UserStorageQuery) (*nex.RMCMessage, *nex.Error)) {
	protocol.SearchContentsWithTotal = handler
}

// SetHandlerDeleteContent sets the handler for the DeleteContent method
func (protocol *Protocol) SetHandlerDeleteContent(handler func(err error, packet nex.PacketInterface, callID uint32, contentKey user_storage_types.UserContentKey) (*nex.RMCMessage, *nex.Error)) {
	protocol.DeleteContent = handler
}

// SetHandlerSaveMetaData sets the handler for the SaveMetaData method
func (protocol *Protocol) SetHandlerSaveMetaData(handler func(err error, packet nex.PacketInterface, callID uint32, properties types.List[user_storage_types.ContentProperty], contentKey user_storage_types.UserContentKey) (*nex.RMCMessage, *nex.Error)) {
	protocol.SaveMetaData = handler
}

// SetHandlerSaveContentDB sets the handler for the SaveContentDB method
func (protocol *Protocol) SetHandlerSaveContentDB(handler func(err error, packet nex.PacketInterface, callID uint32, properties types.List[user_storage_types.ContentProperty], data types.Buffer, contentKey user_storage_types.UserContentKey) (*nex.RMCMessage, *nex.Error)) {
	protocol.SaveContentDB = handler
}

// SetHandlerSaveContentAndGetUploadInfo sets the handler for the SaveContentAndGetUploadInfo method
func (protocol *Protocol) SetHandlerSaveContentAndGetUploadInfo(handler func(err error, packet nex.PacketInterface, callID uint32, properties types.List[user_storage_types.ContentProperty], size types.UInt32, contentKey user_storage_types.UserContentKey) (*nex.RMCMessage, *nex.Error)) {
	protocol.SaveContentAndGetUploadInfo = handler
}

// SetHandlerUploadEnd sets the handler for the UploadEnd method
func (protocol *Protocol) SetHandlerUploadEnd(handler func(err error, packet nex.PacketInterface, callID uint32, pendingID types.UInt64, result types.Bool, contentKey user_storage_types.UserContentKey) (*nex.RMCMessage, *nex.Error)) {
	protocol.UploadEnd = handler
}

// SetHandlerGetContentDB sets the handler for the GetContentDB method
func (protocol *Protocol) SetHandlerGetContentDB(handler func(err error, packet nex.PacketInterface, callID uint32, contentKey user_storage_types.UserContentKey) (*nex.RMCMessage, *nex.Error)) {
	protocol.GetContentDB = handler
}

// SetHandlerGetContentURL sets the handler for the GetContentURL method
func (protocol *Protocol) SetHandlerGetContentURL(handler func(err error, packet nex.PacketInterface, callID uint32, contentKey user_storage_types.UserContentKey) (*nex.RMCMessage, *nex.Error)) {
	protocol.GetContentURL = handler
}

// SetHandlerGetSlotCount sets the handler for the GetSlotCount method
func (protocol *Protocol) SetHandlerGetSlotCount(handler func(err error, packet nex.PacketInterface, callID uint32, typeID types.UInt32) (*nex.RMCMessage, *nex.Error)) {
	protocol.GetSlotCount = handler
}

// SetHandlerGetMetaData sets the handler for the GetMetaData method
func (protocol *Protocol) SetHandlerGetMetaData(handler func(err error, packet nex.PacketInterface, callID uint32, contentKey user_storage_types.UserContentKey) (*nex.RMCMessage, *nex.Error)) {
	protocol.GetMetaData = handler
}

// SetHandlerLike sets the handler for the Like method
func (protocol *Protocol) SetHandlerLike(handler func(err error, packet nex.PacketInterface, callID uint32, contentKey user_storage_types.UserContentKey) (*nex.RMCMessage, *nex.Error)) {
	protocol.Like = handler
}

// SetHandlerUnlike sets the handler for the Unlike method
func (protocol *Protocol) SetHandlerUnlike(handler func(err error, packet nex.PacketInterface, callID uint32, contentKey user_storage_types.UserContentKey) (*nex.RMCMessage, *nex.Error)) {
	protocol.Unlike = handler
}

// SetHandlerIsLiked sets the handler for the IsLiked method
func (protocol *Protocol) SetHandlerIsLiked(handler func(err error, packet nex.PacketInterface, callID uint32, contentKey user_storage_types.UserContentKey) (*nex.RMCMessage, *nex.Error)) {
	protocol.IsLiked = handler
}

// SetHandlerGetFavourites sets the handler for the GetFavourites method
func (protocol *Protocol) SetHandlerGetFavourites(handler func(err error, packet nex.PacketInterface, callID uint32, contentTypes types.List[types.UInt32]) (*nex.RMCMessage, *nex.Error)) {
	protocol.GetFavourites = handler
}

// SetHandlerMakeFavourite sets the handler for the MakeFavourite method
func (protocol *Protocol) SetHandlerMakeFavourite(handler func(err error, packet nex.PacketInterface, callID uint32, contentKey user_storage_types.UserContentKey) (*nex.RMCMessage, *nex.Error)) {
	protocol.MakeFavourite = handler
}

// SetHandlerRemoveFromFavourites sets the handler for the RemoveFromFavourites method
func (protocol *Protocol) SetHandlerRemoveFromFavourites(handler func(err error, packet nex.PacketInterface, callID uint32, contentKey user_storage_types.UserContentKey) (*nex.RMCMessage, *nex.Error)) {
	protocol.RemoveFromFavourites = handler
}

// SetHandlerReportInappropriate sets the handler for the ReportInappropriate method
func (protocol *Protocol) SetHandlerReportInappropriate(handler func(err error, packet nex.PacketInterface, callID uint32, contentKey user_storage_types.UserContentKey, reason types.String) (*nex.RMCMessage, *nex.Error)) {
	protocol.ReportInappropriate = handler
}

// SetHandlerIncrementPlayCount sets the handler for the IncrementPlayCount method
func (protocol *Protocol) SetHandlerIncrementPlayCount(handler func(err error, packet nex.PacketInterface, callID uint32, contentKey user_storage_types.UserContentKey) (*nex.RMCMessage, *nex.Error)) {
	protocol.IncrementPlayCount = handler
}

// SetHandlerUpdateCustomStat sets the handler for the UpdateCustomStat method
func (protocol *Protocol) SetHandlerUpdateCustomStat(handler func(err error, packet nex.PacketInterface, callID uint32, contentKey user_storage_types.UserContentKey, statID types.UInt16, incValue types.Int64) (*nex.RMCMessage, *nex.Error)) {
	protocol.UpdateCustomStat = handler
}

// SetHandlerGetOwnContents sets the handler for the GetOwnContents method
func (protocol *Protocol) SetHandlerGetOwnContents(handler func(err error, packet nex.PacketInterface, callID uint32, typeID types.UInt32) (*nex.RMCMessage, *nex.Error)) {
	protocol.GetOwnContents = handler
}

// SetHandlerGetMostPopularTags sets the handler for the GetMostPopularTags method
func (protocol *Protocol) SetHandlerGetMostPopularTags(handler func(err error, packet nex.PacketInterface, callID uint32, contentKey user_storage_types.UserContentKey) (*nex.RMCMessage, *nex.Error)) {
	protocol.GetMostPopularTags = handler
}

// SetHandlerGetTags sets the handler for the GetTags method
func (protocol *Protocol) SetHandlerGetTags(handler func(err error, packet nex.PacketInterface, callID uint32, contentKey user_storage_types.UserContentKey) (*nex.RMCMessage, *nex.Error)) {
	protocol.GetTags = handler
}

// SetHandlerTagContent sets the handler for the TagContent method
func (protocol *Protocol) SetHandlerTagContent(handler func(err error, packet nex.PacketInterface, callID uint32, contentKey user_storage_types.UserContentKey, newTagIds types.List[types.UInt32]) (*nex.RMCMessage, *nex.Error)) {
	protocol.TagContent = handler
}

// SetHandlerSearchContentsByPlayers sets the handler for the SearchContentsByPlayers method
func (protocol *Protocol) SetHandlerSearchContentsByPlayers(handler func(err error, packet nex.PacketInterface, callID uint32, pids types.List[types.UInt32], query user_storage_types.UserStorageQuery) (*nex.RMCMessage, *nex.Error)) {
	protocol.SearchContentsByPlayers = handler
}

// SetHandlerSearchContentsByPlayersWithTotal sets the handler for the SearchContentsByPlayersWithTotal method
func (protocol *Protocol) SetHandlerSearchContentsByPlayersWithTotal(handler func(err error, packet nex.PacketInterface, callID uint32, pids types.List[types.UInt32], query user_storage_types.UserStorageQuery) (*nex.RMCMessage, *nex.Error)) {
	protocol.SearchContentsByPlayersWithTotal = handler
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
	case MethodSearchContents:
		protocol.handleSearchContents(packet)
	case MethodSearchContentsWithTotal:
		protocol.handleSearchContentsWithTotal(packet)
	case MethodDeleteContent:
		protocol.handleDeleteContent(packet)
	case MethodSaveMetaData:
		protocol.handleSaveMetaData(packet)
	case MethodSaveContentDB:
		protocol.handleSaveContentDB(packet)
	case MethodSaveContentAndGetUploadInfo:
		protocol.handleSaveContentAndGetUploadInfo(packet)
	case MethodUploadEnd:
		protocol.handleUploadEnd(packet)
	case MethodGetContentDB:
		protocol.handleGetContentDB(packet)
	case MethodGetContentURL:
		protocol.handleGetContentURL(packet)
	case MethodGetSlotCount:
		protocol.handleGetSlotCount(packet)
	case MethodGetMetaData:
		protocol.handleGetMetaData(packet)
	case MethodLike:
		protocol.handleLike(packet)
	case MethodUnlike:
		protocol.handleUnlike(packet)
	case MethodIsLiked:
		protocol.handleIsLiked(packet)
	case MethodGetFavourites:
		protocol.handleGetFavourites(packet)
	case MethodMakeFavourite:
		protocol.handleMakeFavourite(packet)
	case MethodRemoveFromFavourites:
		protocol.handleRemoveFromFavourites(packet)
	case MethodReportInappropriate:
		protocol.handleReportInappropriate(packet)
	case MethodIncrementPlayCount:
		protocol.handleIncrementPlayCount(packet)
	case MethodUpdateCustomStat:
		protocol.handleUpdateCustomStat(packet)
	case MethodGetOwnContents:
		protocol.handleGetOwnContents(packet)
	case MethodGetMostPopularTags:
		protocol.handleGetMostPopularTags(packet)
	case MethodGetTags:
		protocol.handleGetTags(packet)
	case MethodTagContent:
		protocol.handleTagContent(packet)
	case MethodSearchContentsByPlayers:
		protocol.handleSearchContentsByPlayers(packet)
	case MethodSearchContentsByPlayersWithTotal:
		protocol.handleSearchContentsByPlayersWithTotal(packet)
	default:
		errMessage := fmt.Sprintf("Unsupported User Storage method ID: %#v\n", message.MethodID)
		err := nex.NewError(nex.ResultCodes.Core.NotImplemented, errMessage)

		globals.RespondError(packet, ProtocolID, err)
		globals.Logger.Warning(err.Message)
	}
}

// NewProtocol returns a new User Storage protocol
func NewProtocol() *Protocol {
	return &Protocol{}
}
