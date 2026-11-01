package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go/v2"
	"github.com/PretendoNetwork/nex-go/v2/types"
	"github.com/PretendoNetwork/nex-protocols-go/v2/globals"
)

func (protocol *Protocol) handleActivateKey(packet nex.PacketInterface) {
	if protocol.ActivateKey == nil {
		err := nex.NewError(nex.ResultCodes.Core.NotImplemented, "Privileges::ActivateKey not implemented")

		globals.Logger.Warning(err.Message)
		globals.RespondError(packet, ProtocolID, err)

		return
	}

	request := packet.RMCMessage()
	callID := request.CallID
	parameters := request.Parameters
	endpoint := packet.Sender().Endpoint()
	parametersStream := nex.NewByteStreamIn(parameters, endpoint.LibraryVersions(), endpoint.ByteStreamSettings())

	uniqueKey, languageCode, err := parseActivateKeyRequest(parametersStream)
	if err != nil {
		_, rmcError := protocol.ActivateKey(err, packet, callID, uniqueKey, languageCode)
		if rmcError != nil {
			globals.RespondError(packet, ProtocolID, rmcError)
		}

		return
	}

	rmcMessage, rmcError := protocol.ActivateKey(nil, packet, callID, uniqueKey, languageCode)
	if rmcError != nil {
		globals.RespondError(packet, ProtocolID, rmcError)
		return
	}

	globals.Respond(packet, rmcMessage)
}

// parseActivateKeyRequest decodes the parameters of Privileges::ActivateKey from the given stream
func parseActivateKeyRequest(stream *nex.ByteStreamIn) (types.String, types.String, error) {
	var uniqueKey types.String
	var languageCode types.String

	if err := uniqueKey.ExtractFrom(stream); err != nil {
		return uniqueKey, languageCode, fmt.Errorf("failed to read uniqueKey from parameters. %s", err.Error())
	}

	if err := languageCode.ExtractFrom(stream); err != nil {
		return uniqueKey, languageCode, fmt.Errorf("failed to read languageCode from parameters. %s", err.Error())
	}

	return uniqueKey, languageCode, nil
}
