// Package protocol implements the Pokemon Bank Shop protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go/v2"
	"github.com/PretendoNetwork/nex-go/v2/types"
	"github.com/PretendoNetwork/nex-protocols-go/v2/globals"
	shop "github.com/PretendoNetwork/nex-protocols-go/v2/shop"
	"golang.org/x/exp/slices"
)

const (
	// ProtocolID is the Protocol ID for the Shop (Pokemon Bank) protocol
	ProtocolID = 0xC8

	// MethodGetItems is the method ID for GetItems
	MethodGetItems = 0x1

	// MethodGetChallengeBlob is the method ID for GetChallengeBlob
	MethodGetChallengeBlob = 0x2

	// MethodGetRivToken is the method ID for GetRivToken
	MethodGetRivToken = 0x3

	// MethodGetRivTokenByItemID is the method ID for GetRivTokenByItemID
	MethodGetRivTokenByItemID = 0x4

	// MethodGetItemRights is the method ID for GetItemRights
	MethodGetItemRights = 0x5

	// MethodVerifyAndRegisterTicket is the method ID for VerifyAndRegisterTicket
	MethodVerifyAndRegisterTicket = 0x6

	// MethodDebugSetExpireTime is the method ID for DebugSetExpireTime
	MethodDebugSetExpireTime = 0x7

	// MethodPrincipalIDToSupportNumber is the method ID for PrincipalIDToSupportNumber
	MethodPrincipalIDToSupportNumber = 0x8

	// MethodSupportNumberToPrincipalID is the method ID for SupportNumberToPrincipalID
	MethodSupportNumberToPrincipalID = 0x9

	// MethodGetGameServerTime is the method ID for GetGameServerTime
	MethodGetGameServerTime = 0xA
)

var patchedMethods = []uint32{
	MethodGetItems,
	MethodGetChallengeBlob,
	MethodGetRivToken,
	MethodGetRivTokenByItemID,
	MethodGetItemRights,
	MethodVerifyAndRegisterTicket,
	MethodDebugSetExpireTime,
	MethodPrincipalIDToSupportNumber,
	MethodSupportNumberToPrincipalID,
	MethodGetGameServerTime,
}

type shopProtocol = shop.Protocol

// Protocol stores all the RMC method handlers for the Shop (Pokemon Bank) protocol and listens for requests
// Embeds the Shop protocol
type Protocol struct {
	endpoint nex.EndpointInterface
	shopProtocol
	GetItems                   func(err error, packet nex.PacketInterface, callID uint32) (*nex.RMCMessage, *nex.Error)
	GetChallengeBlob           func(err error, packet nex.PacketInterface, callID uint32) (*nex.RMCMessage, *nex.Error)
	GetRivToken                func(err error, packet nex.PacketInterface, callID uint32, itemCode types.String, referenceID types.QBuffer) (*nex.RMCMessage, *nex.Error)
	GetRivTokenByItemID        func(err error, packet nex.PacketInterface, callID uint32, itemID types.UInt32) (*nex.RMCMessage, *nex.Error)
	GetItemRights              func(err error, packet nex.PacketInterface, callID uint32, ticketEnvelope types.QBuffer) (*nex.RMCMessage, *nex.Error)
	VerifyAndRegisterTicket    func(err error, packet nex.PacketInterface, callID uint32, ticketEnvelope types.QBuffer, purchasedTime types.DateTime) (*nex.RMCMessage, *nex.Error)
	DebugSetExpireTime         func(err error, packet nex.PacketInterface, callID uint32, expireTime types.DateTime) (*nex.RMCMessage, *nex.Error)
	PrincipalIDToSupportNumber func(err error, packet nex.PacketInterface, callID uint32, pid types.UInt32) (*nex.RMCMessage, *nex.Error)
	SupportNumberToPrincipalID func(err error, packet nex.PacketInterface, callID uint32, supportNumber types.String) (*nex.RMCMessage, *nex.Error)
	GetGameServerTime          func(err error, packet nex.PacketInterface, callID uint32) (*nex.RMCMessage, *nex.Error)
}

// Interface implements the methods present on the Shop (Pokemon Bank) protocol struct
type Interface interface {
	Endpoint() nex.EndpointInterface
	SetEndpoint(endpoint nex.EndpointInterface)
	SetHandlerGetItems(handler func(err error, packet nex.PacketInterface, callID uint32) (*nex.RMCMessage, *nex.Error))
	SetHandlerGetChallengeBlob(handler func(err error, packet nex.PacketInterface, callID uint32) (*nex.RMCMessage, *nex.Error))
	SetHandlerGetRivToken(handler func(err error, packet nex.PacketInterface, callID uint32, itemCode types.String, referenceID types.QBuffer) (*nex.RMCMessage, *nex.Error))
	SetHandlerGetRivTokenByItemID(handler func(err error, packet nex.PacketInterface, callID uint32, itemID types.UInt32) (*nex.RMCMessage, *nex.Error))
	SetHandlerGetItemRights(handler func(err error, packet nex.PacketInterface, callID uint32, ticketEnvelope types.QBuffer) (*nex.RMCMessage, *nex.Error))
	SetHandlerVerifyAndRegisterTicket(handler func(err error, packet nex.PacketInterface, callID uint32, ticketEnvelope types.QBuffer, purchasedTime types.DateTime) (*nex.RMCMessage, *nex.Error))
	SetHandlerDebugSetExpireTime(handler func(err error, packet nex.PacketInterface, callID uint32, expireTime types.DateTime) (*nex.RMCMessage, *nex.Error))
	SetHandlerPrincipalIDToSupportNumber(handler func(err error, packet nex.PacketInterface, callID uint32, pid types.UInt32) (*nex.RMCMessage, *nex.Error))
	SetHandlerSupportNumberToPrincipalID(handler func(err error, packet nex.PacketInterface, callID uint32, supportNumber types.String) (*nex.RMCMessage, *nex.Error))
	SetHandlerGetGameServerTime(handler func(err error, packet nex.PacketInterface, callID uint32) (*nex.RMCMessage, *nex.Error))
}

// SetHandlerGetItems sets the handler for the GetItems method
func (protocol *Protocol) SetHandlerGetItems(handler func(err error, packet nex.PacketInterface, callID uint32) (*nex.RMCMessage, *nex.Error)) {
	protocol.GetItems = handler
}

// SetHandlerGetChallengeBlob sets the handler for the GetChallengeBlob method
func (protocol *Protocol) SetHandlerGetChallengeBlob(handler func(err error, packet nex.PacketInterface, callID uint32) (*nex.RMCMessage, *nex.Error)) {
	protocol.GetChallengeBlob = handler
}

// SetHandlerGetRivToken sets the handler for the GetRivToken method
func (protocol *Protocol) SetHandlerGetRivToken(handler func(err error, packet nex.PacketInterface, callID uint32, itemCode types.String, referenceID types.QBuffer) (*nex.RMCMessage, *nex.Error)) {
	protocol.GetRivToken = handler
}

// SetHandlerGetRivTokenByItemID sets the handler for the GetRivTokenByItemID method
func (protocol *Protocol) SetHandlerGetRivTokenByItemID(handler func(err error, packet nex.PacketInterface, callID uint32, itemID types.UInt32) (*nex.RMCMessage, *nex.Error)) {
	protocol.GetRivTokenByItemID = handler
}

// SetHandlerGetItemRights sets the handler for the GetItemRights method
func (protocol *Protocol) SetHandlerGetItemRights(handler func(err error, packet nex.PacketInterface, callID uint32, ticketEnvelope types.QBuffer) (*nex.RMCMessage, *nex.Error)) {
	protocol.GetItemRights = handler
}

// SetHandlerVerifyAndRegisterTicket sets the handler for the VerifyAndRegisterTicket method
func (protocol *Protocol) SetHandlerVerifyAndRegisterTicket(handler func(err error, packet nex.PacketInterface, callID uint32, ticketEnvelope types.QBuffer, purchasedTime types.DateTime) (*nex.RMCMessage, *nex.Error)) {
	protocol.VerifyAndRegisterTicket = handler
}

// SetHandlerDebugSetExpireTime sets the handler for the DebugSetExpireTime method
func (protocol *Protocol) SetHandlerDebugSetExpireTime(handler func(err error, packet nex.PacketInterface, callID uint32, expireTime types.DateTime) (*nex.RMCMessage, *nex.Error)) {
	protocol.DebugSetExpireTime = handler
}

// SetHandlerPrincipalIDToSupportNumber sets the handler for the PrincipalIDToSupportNumber method
func (protocol *Protocol) SetHandlerPrincipalIDToSupportNumber(handler func(err error, packet nex.PacketInterface, callID uint32, pid types.UInt32) (*nex.RMCMessage, *nex.Error)) {
	protocol.PrincipalIDToSupportNumber = handler
}

// SetHandlerSupportNumberToPrincipalID sets the handler for the SupportNumberToPrincipalID method
func (protocol *Protocol) SetHandlerSupportNumberToPrincipalID(handler func(err error, packet nex.PacketInterface, callID uint32, supportNumber types.String) (*nex.RMCMessage, *nex.Error)) {
	protocol.SupportNumberToPrincipalID = handler
}

// SetHandlerGetGameServerTime sets the handler for the GetGameServerTime method
func (protocol *Protocol) SetHandlerGetGameServerTime(handler func(err error, packet nex.PacketInterface, callID uint32) (*nex.RMCMessage, *nex.Error)) {
	protocol.GetGameServerTime = handler
}

// HandlePacket sends the packet to the correct RMC method handler
func (protocol *Protocol) HandlePacket(packet nex.PacketInterface) {
	message := packet.RMCMessage()

	if !message.IsRequest || message.ProtocolID != ProtocolID {
		return
	}

	if !slices.Contains(patchedMethods, message.MethodID) {
		protocol.shopProtocol.HandlePacket(packet)
		return
	}

	switch message.MethodID {
	case MethodGetItems:
		protocol.handleGetItems(packet)
	case MethodGetChallengeBlob:
		protocol.handleGetChallengeBlob(packet)
	case MethodGetRivToken:
		protocol.handleGetRivToken(packet)
	case MethodGetRivTokenByItemID:
		protocol.handleGetRivTokenByItemID(packet)
	case MethodGetItemRights:
		protocol.handleGetItemRights(packet)
	case MethodVerifyAndRegisterTicket:
		protocol.handleVerifyAndRegisterTicket(packet)
	case MethodDebugSetExpireTime:
		protocol.handleDebugSetExpireTime(packet)
	case MethodPrincipalIDToSupportNumber:
		protocol.handlePrincipalIDToSupportNumber(packet)
	case MethodSupportNumberToPrincipalID:
		protocol.handleSupportNumberToPrincipalID(packet)
	case MethodGetGameServerTime:
		protocol.handleGetGameServerTime(packet)
	default:
		errMessage := fmt.Sprintf("Unsupported ShopPokemonBank method ID: %#v\n", message.MethodID)
		err := nex.NewError(nex.ResultCodes.Core.NotImplemented, errMessage)

		globals.RespondError(packet, ProtocolID, err)
		globals.Logger.Warning(err.Message)
	}
}

// NewProtocol returns a new Shop (Pokemon Bank)
func NewProtocol(endpoint nex.EndpointInterface) *Protocol {
	protocol := &Protocol{endpoint: endpoint}
	protocol.SetEndpoint(endpoint)

	return protocol
}
