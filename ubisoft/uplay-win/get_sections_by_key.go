package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go/v2"
	"github.com/PretendoNetwork/nex-go/v2/types"
	"github.com/PretendoNetwork/nex-protocols-go/v2/globals"
)

func (protocol *Protocol) handleGetSectionsByKey(packet nex.PacketInterface) {
	if protocol.GetSectionsByKey == nil {
		err := nex.NewError(nex.ResultCodes.Core.NotImplemented, "UplayWin::GetSectionsByKey not implemented")

		globals.Logger.Warning(err.Message)
		globals.RespondError(packet, ProtocolID, err)

		return
	}

	request := packet.RMCMessage()
	callID := request.CallID
	parameters := request.Parameters
	endpoint := packet.Sender().Endpoint()
	parametersStream := nex.NewByteStreamIn(parameters, endpoint.LibraryVersions(), endpoint.ByteStreamSettings())

	cultureName, sectionKey, platformCode, gameCode, err := parseGetSectionsByKeyRequest(parametersStream)
	if err != nil {
		_, rmcError := protocol.GetSectionsByKey(err, packet, callID, cultureName, sectionKey, platformCode, gameCode)
		if rmcError != nil {
			globals.RespondError(packet, ProtocolID, rmcError)
		}

		return
	}

	rmcMessage, rmcError := protocol.GetSectionsByKey(nil, packet, callID, cultureName, sectionKey, platformCode, gameCode)
	if rmcError != nil {
		globals.RespondError(packet, ProtocolID, rmcError)
		return
	}

	globals.Respond(packet, rmcMessage)
}

// parseGetSectionsByKeyRequest decodes the parameters of Uplay Win::GetSectionsByKey from the given stream
func parseGetSectionsByKeyRequest(stream *nex.ByteStreamIn) (types.String, types.String, types.String, types.String, error) {
	var cultureName types.String
	var sectionKey types.String
	var platformCode types.String
	var gameCode types.String

	if err := cultureName.ExtractFrom(stream); err != nil {
		return cultureName, sectionKey, platformCode, gameCode, fmt.Errorf("failed to read cultureName from parameters. %s", err.Error())
	}

	if err := sectionKey.ExtractFrom(stream); err != nil {
		return cultureName, sectionKey, platformCode, gameCode, fmt.Errorf("failed to read sectionKey from parameters. %s", err.Error())
	}

	if err := platformCode.ExtractFrom(stream); err != nil {
		return cultureName, sectionKey, platformCode, gameCode, fmt.Errorf("failed to read platformCode from parameters. %s", err.Error())
	}

	if err := gameCode.ExtractFrom(stream); err != nil {
		return cultureName, sectionKey, platformCode, gameCode, fmt.Errorf("failed to read gameCode from parameters. %s", err.Error())
	}

	return cultureName, sectionKey, platformCode, gameCode, nil
}
