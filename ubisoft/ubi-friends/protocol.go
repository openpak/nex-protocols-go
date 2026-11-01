// Package protocol declares the Ubi friends protocol of the Quazal Rendez-Vous (Ubisoft) family.
//
// The NintendoClients wiki lists the protocol ID (47) on its "NEX Protocols" page
// but documents no methods, so none are defined here. Every request is answered
// with Core::NotImplemented.
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go/v2"
	"github.com/PretendoNetwork/nex-protocols-go/v2/globals"
)

// ProtocolID is the protocol ID for the Ubi friends protocol
const ProtocolID = 0x2f

// Protocol handles the Ubi friends protocol
type Protocol struct {
	endpoint nex.EndpointInterface
}

// Interface implements the methods present on the Ubi friends protocol struct
type Interface interface {
	Endpoint() nex.EndpointInterface
	SetEndpoint(endpoint nex.EndpointInterface)
}

// Endpoint returns the endpoint implementing the protocol
func (protocol *Protocol) Endpoint() nex.EndpointInterface {
	return protocol.endpoint
}

// SetEndpoint sets the endpoint implementing the protocol
func (protocol *Protocol) SetEndpoint(endpoint nex.EndpointInterface) {
	protocol.endpoint = endpoint
}

// HandlePacket answers every request with Core::NotImplemented
func (protocol *Protocol) HandlePacket(packet nex.PacketInterface) {
	message := packet.RMCMessage()

	if !message.IsRequest || message.ProtocolID != ProtocolID {
		return
	}

	errMessage := fmt.Sprintf("Ubi friends protocol methods are undocumented, unsupported method ID: %#v", message.MethodID)
	err := nex.NewError(nex.ResultCodes.Core.NotImplemented, errMessage)

	globals.Logger.Warning(err.Message)
	globals.RespondError(packet, ProtocolID, err)
}

// NewProtocol returns a new Ubi friends protocol
func NewProtocol() *Protocol {
	return &Protocol{}
}
