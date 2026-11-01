package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go/v2"
	"github.com/PretendoNetwork/nex-go/v2/types"
	"github.com/PretendoNetwork/nex-protocols-go/v2/globals"
)

func (protocol *Protocol) handleUplayWelcome(packet nex.PacketInterface) {
	if protocol.UplayWelcome == nil {
		err := nex.NewError(nex.ResultCodes.Core.NotImplemented, "UplayWin::UplayWelcome not implemented")

		globals.Logger.Warning(err.Message)
		globals.RespondError(packet, ProtocolID, err)

		return
	}

	request := packet.RMCMessage()
	callID := request.CallID
	parameters := request.Parameters
	endpoint := packet.Sender().Endpoint()
	parametersStream := nex.NewByteStreamIn(parameters, endpoint.LibraryVersions(), endpoint.ByteStreamSettings())

	cultureName, platformCode, gameCode, err := parseUplayWelcomeRequest(parametersStream)
	if err != nil {
		_, rmcError := protocol.UplayWelcome(err, packet, callID, cultureName, platformCode, gameCode)
		if rmcError != nil {
			globals.RespondError(packet, ProtocolID, rmcError)
		}

		return
	}

	rmcMessage, rmcError := protocol.UplayWelcome(nil, packet, callID, cultureName, platformCode, gameCode)
	if rmcError != nil {
		globals.RespondError(packet, ProtocolID, rmcError)
		return
	}

	globals.Respond(packet, rmcMessage)
}

// parseUplayWelcomeRequest decodes the parameters of Uplay Win::UplayWelcome from the given stream
func parseUplayWelcomeRequest(stream *nex.ByteStreamIn) (types.String, types.String, types.String, error) {
	var cultureName types.String
	var platformCode types.String
	var gameCode types.String

	if err := cultureName.ExtractFrom(stream); err != nil {
		return cultureName, platformCode, gameCode, fmt.Errorf("failed to read cultureName from parameters. %s", err.Error())
	}

	if err := platformCode.ExtractFrom(stream); err != nil {
		return cultureName, platformCode, gameCode, fmt.Errorf("failed to read platformCode from parameters. %s", err.Error())
	}

	if err := gameCode.ExtractFrom(stream); err != nil {
		return cultureName, platformCode, gameCode, fmt.Errorf("failed to read gameCode from parameters. %s", err.Error())
	}

	return cultureName, platformCode, gameCode, nil
}
