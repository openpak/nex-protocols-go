// Package protocol implements the Pokemon Bank Shop protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go/v2"
	"github.com/PretendoNetwork/nex-go/v2/types"
	"github.com/PretendoNetwork/nex-protocols-go/v2/globals"
)

// extractDebugSetExpireTimeParameters reads the DebugSetExpireTime request parameters from the given stream
func extractDebugSetExpireTimeParameters(stream *nex.ByteStreamIn) (expireTime types.DateTime, err error) {
	if err = expireTime.ExtractFrom(stream); err != nil {
		err = fmt.Errorf("failed to read expireTime from parameters. %s", err.Error())
		return
	}

	return
}

func (protocol *Protocol) handleDebugSetExpireTime(packet nex.PacketInterface) {
	if protocol.DebugSetExpireTime == nil {
		err := nex.NewError(nex.ResultCodes.Core.NotImplemented, "ShopPokemonBank::DebugSetExpireTime not implemented")

		globals.Logger.Warning(err.Message)
		globals.RespondError(packet, ProtocolID, err)

		return
	}

	request := packet.RMCMessage()
	callID := request.CallID
	parameters := request.Parameters
	endpoint := packet.Sender().Endpoint()
	parametersStream := nex.NewByteStreamIn(parameters, endpoint.LibraryVersions(), endpoint.ByteStreamSettings())

	expireTime, err := extractDebugSetExpireTimeParameters(parametersStream)
	if err != nil {
		_, rmcError := protocol.DebugSetExpireTime(err, packet, callID, expireTime)
		if rmcError != nil {
			globals.RespondError(packet, ProtocolID, rmcError)
		}

		return
	}

	rmcMessage, rmcError := protocol.DebugSetExpireTime(nil, packet, callID, expireTime)
	if rmcError != nil {
		globals.RespondError(packet, ProtocolID, rmcError)
		return
	}

	globals.Respond(packet, rmcMessage)
}
