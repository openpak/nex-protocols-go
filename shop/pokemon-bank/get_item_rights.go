// Package protocol implements the Pokemon Bank Shop protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go/v2"
	"github.com/PretendoNetwork/nex-go/v2/types"
	"github.com/PretendoNetwork/nex-protocols-go/v2/globals"
)

// extractGetItemRightsParameters reads the GetItemRights request parameters from the given stream
func extractGetItemRightsParameters(stream *nex.ByteStreamIn) (ticketEnvelope types.QBuffer, err error) {
	if err = ticketEnvelope.ExtractFrom(stream); err != nil {
		err = fmt.Errorf("failed to read ticketEnvelope from parameters. %s", err.Error())
		return
	}

	return
}

func (protocol *Protocol) handleGetItemRights(packet nex.PacketInterface) {
	if protocol.GetItemRights == nil {
		err := nex.NewError(nex.ResultCodes.Core.NotImplemented, "ShopPokemonBank::GetItemRights not implemented")

		globals.Logger.Warning(err.Message)
		globals.RespondError(packet, ProtocolID, err)

		return
	}

	request := packet.RMCMessage()
	callID := request.CallID
	parameters := request.Parameters
	endpoint := packet.Sender().Endpoint()
	parametersStream := nex.NewByteStreamIn(parameters, endpoint.LibraryVersions(), endpoint.ByteStreamSettings())

	ticketEnvelope, err := extractGetItemRightsParameters(parametersStream)
	if err != nil {
		_, rmcError := protocol.GetItemRights(err, packet, callID, ticketEnvelope)
		if rmcError != nil {
			globals.RespondError(packet, ProtocolID, rmcError)
		}

		return
	}

	rmcMessage, rmcError := protocol.GetItemRights(nil, packet, callID, ticketEnvelope)
	if rmcError != nil {
		globals.RespondError(packet, ProtocolID, rmcError)
		return
	}

	globals.Respond(packet, rmcMessage)
}
