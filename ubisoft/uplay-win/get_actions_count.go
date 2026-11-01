package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go/v2"
	"github.com/PretendoNetwork/nex-go/v2/types"
	"github.com/PretendoNetwork/nex-protocols-go/v2/globals"
)

func (protocol *Protocol) handleGetActionsCount(packet nex.PacketInterface) {
	if protocol.GetActionsCount == nil {
		err := nex.NewError(nex.ResultCodes.Core.NotImplemented, "UplayWin::GetActionsCount not implemented")

		globals.Logger.Warning(err.Message)
		globals.RespondError(packet, ProtocolID, err)

		return
	}

	request := packet.RMCMessage()
	callID := request.CallID
	parameters := request.Parameters
	endpoint := packet.Sender().Endpoint()
	parametersStream := nex.NewByteStreamIn(parameters, endpoint.LibraryVersions(), endpoint.ByteStreamSettings())

	platformCode, gameCode, err := parseGetActionsCountRequest(parametersStream)
	if err != nil {
		_, rmcError := protocol.GetActionsCount(err, packet, callID, platformCode, gameCode)
		if rmcError != nil {
			globals.RespondError(packet, ProtocolID, rmcError)
		}

		return
	}

	rmcMessage, rmcError := protocol.GetActionsCount(nil, packet, callID, platformCode, gameCode)
	if rmcError != nil {
		globals.RespondError(packet, ProtocolID, rmcError)
		return
	}

	globals.Respond(packet, rmcMessage)
}

// parseGetActionsCountRequest decodes the parameters of Uplay Win::GetActionsCount from the given stream
func parseGetActionsCountRequest(stream *nex.ByteStreamIn) (types.String, types.String, error) {
	var platformCode types.String
	var gameCode types.String

	if err := platformCode.ExtractFrom(stream); err != nil {
		return platformCode, gameCode, fmt.Errorf("failed to read platformCode from parameters. %s", err.Error())
	}

	if err := gameCode.ExtractFrom(stream); err != nil {
		return platformCode, gameCode, fmt.Errorf("failed to read gameCode from parameters. %s", err.Error())
	}

	return platformCode, gameCode, nil
}
