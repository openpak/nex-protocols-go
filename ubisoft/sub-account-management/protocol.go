// Package protocol declares the Sub Account Management protocol (Rendez-Vous/Quazal, protocol ID 44).
//
// The NintendoClients wiki lists this protocol ID under "Not provided by NEX"
// (page "NEX Protocols") but has no page for it, so its methods and structures
// are undocumented. Only the protocol ID is defined here; every request is
// answered with Core::NotImplemented.
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go/v2"
	"github.com/PretendoNetwork/nex-protocols-go/v2/globals"
)

const (
	// ProtocolID is the protocol ID for the Sub Account Management protocol
	ProtocolID = 0x2C
)

// Protocol listens for Sub Account Management requests and rejects all of them
type Protocol struct {
	endpoint nex.EndpointInterface
}

// Interface implements the methods present on the Sub Account Management protocol struct
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

	errMessage := fmt.Sprintf("Sub Account Management method ID %#v is undocumented and not implemented\n", message.MethodID)
	err := nex.NewError(nex.ResultCodes.Core.NotImplemented, errMessage)

	globals.RespondError(packet, ProtocolID, err)
	globals.Logger.Warning(err.Message)
}

// NewProtocol returns a new Sub Account Management protocol
func NewProtocol() *Protocol {
	return &Protocol{}
}
