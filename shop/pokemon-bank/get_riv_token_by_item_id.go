// Package protocol implements the Pokemon Bank Shop protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go/v2"
	"github.com/PretendoNetwork/nex-go/v2/types"
	"github.com/PretendoNetwork/nex-protocols-go/v2/globals"
)

// extractGetRivTokenByItemIDParameters reads the GetRivTokenByItemID request parameters from the given stream
func extractGetRivTokenByItemIDParameters(stream *nex.ByteStreamIn) (itemID types.UInt32, err error) {
	if err = itemID.ExtractFrom(stream); err != nil {
		err = fmt.Errorf("failed to read itemID from parameters. %s", err.Error())
		return
	}

	return
}

func (protocol *Protocol) handleGetRivTokenByItemID(packet nex.PacketInterface) {
	if protocol.GetRivTokenByItemID == nil {
		err := nex.NewError(nex.ResultCodes.Core.NotImplemented, "ShopPokemonBank::GetRivTokenByItemID not implemented")

		globals.Logger.Warning(err.Message)
		globals.RespondError(packet, ProtocolID, err)

		return
	}

	request := packet.RMCMessage()
	callID := request.CallID
	parameters := request.Parameters
	endpoint := packet.Sender().Endpoint()
	parametersStream := nex.NewByteStreamIn(parameters, endpoint.LibraryVersions(), endpoint.ByteStreamSettings())

	itemID, err := extractGetRivTokenByItemIDParameters(parametersStream)
	if err != nil {
		_, rmcError := protocol.GetRivTokenByItemID(err, packet, callID, itemID)
		if rmcError != nil {
			globals.RespondError(packet, ProtocolID, rmcError)
		}

		return
	}

	rmcMessage, rmcError := protocol.GetRivTokenByItemID(nil, packet, callID, itemID)
	if rmcError != nil {
		globals.RespondError(packet, ProtocolID, rmcError)
		return
	}

	globals.Respond(packet, rmcMessage)
}
