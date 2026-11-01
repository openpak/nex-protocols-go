// Package protocol implements the Pokemon Bank Shop protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go/v2"
	"github.com/PretendoNetwork/nex-go/v2/types"
	"github.com/PretendoNetwork/nex-protocols-go/v2/globals"
)

// extractGetRivTokenParameters reads the GetRivToken request parameters from the given stream
func extractGetRivTokenParameters(stream *nex.ByteStreamIn) (itemCode types.String, referenceID types.QBuffer, err error) {
	if err = itemCode.ExtractFrom(stream); err != nil {
		err = fmt.Errorf("failed to read itemCode from parameters. %s", err.Error())
		return
	}

	if err = referenceID.ExtractFrom(stream); err != nil {
		err = fmt.Errorf("failed to read referenceID from parameters. %s", err.Error())
		return
	}

	return
}

func (protocol *Protocol) handleGetRivToken(packet nex.PacketInterface) {
	if protocol.GetRivToken == nil {
		err := nex.NewError(nex.ResultCodes.Core.NotImplemented, "ShopPokemonBank::GetRivToken not implemented")

		globals.Logger.Warning(err.Message)
		globals.RespondError(packet, ProtocolID, err)

		return
	}

	request := packet.RMCMessage()
	callID := request.CallID
	parameters := request.Parameters
	endpoint := packet.Sender().Endpoint()
	parametersStream := nex.NewByteStreamIn(parameters, endpoint.LibraryVersions(), endpoint.ByteStreamSettings())

	itemCode, referenceID, err := extractGetRivTokenParameters(parametersStream)
	if err != nil {
		_, rmcError := protocol.GetRivToken(err, packet, callID, itemCode, referenceID)
		if rmcError != nil {
			globals.RespondError(packet, ProtocolID, rmcError)
		}

		return
	}

	rmcMessage, rmcError := protocol.GetRivToken(nil, packet, callID, itemCode, referenceID)
	if rmcError != nil {
		globals.RespondError(packet, ProtocolID, rmcError)
		return
	}

	globals.Respond(packet, rmcMessage)
}
