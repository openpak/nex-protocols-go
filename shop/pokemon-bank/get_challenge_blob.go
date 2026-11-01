// Package protocol implements the Pokemon Bank Shop protocol
package protocol

import (
	nex "github.com/PretendoNetwork/nex-go/v2"
	"github.com/PretendoNetwork/nex-protocols-go/v2/globals"
)

func (protocol *Protocol) handleGetChallengeBlob(packet nex.PacketInterface) {
	if protocol.GetChallengeBlob == nil {
		err := nex.NewError(nex.ResultCodes.Core.NotImplemented, "ShopPokemonBank::GetChallengeBlob not implemented")

		globals.Logger.Warning(err.Message)
		globals.RespondError(packet, ProtocolID, err)

		return
	}

	request := packet.RMCMessage()
	callID := request.CallID

	rmcMessage, rmcError := protocol.GetChallengeBlob(nil, packet, callID)
	if rmcError != nil {
		globals.RespondError(packet, ProtocolID, rmcError)
		return
	}

	globals.Respond(packet, rmcMessage)
}
