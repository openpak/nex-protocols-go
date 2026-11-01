package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go/v2"
	"github.com/PretendoNetwork/nex-go/v2/types"
	"github.com/PretendoNetwork/nex-protocols-go/v2/globals"
)

func (protocol *Protocol) handleGetTOSEx(packet nex.PacketInterface) {
	if protocol.GetTOSEx == nil {
		err := nex.NewError(nex.ResultCodes.Core.NotImplemented, "UbiAccountManagement::GetTOSEx not implemented")

		globals.Logger.Warning(err.Message)
		globals.RespondError(packet, ProtocolID, err)

		return
	}

	request := packet.RMCMessage()
	callID := request.CallID
	parameters := request.Parameters
	endpoint := packet.Sender().Endpoint()
	parametersStream := nex.NewByteStreamIn(parameters, endpoint.LibraryVersions(), endpoint.ByteStreamSettings())

	countryCode, languageCode, htmlVersion, err := parseGetTOSExRequest(parametersStream)
	if err != nil {
		_, rmcError := protocol.GetTOSEx(err, packet, callID, countryCode, languageCode, htmlVersion)
		if rmcError != nil {
			globals.RespondError(packet, ProtocolID, rmcError)
		}

		return
	}

	rmcMessage, rmcError := protocol.GetTOSEx(nil, packet, callID, countryCode, languageCode, htmlVersion)
	if rmcError != nil {
		globals.RespondError(packet, ProtocolID, rmcError)
		return
	}

	globals.Respond(packet, rmcMessage)
}

// parseGetTOSExRequest decodes the parameters of Ubi Account Management::GetTOSEx from the given stream
func parseGetTOSExRequest(stream *nex.ByteStreamIn) (types.String, types.String, types.Bool, error) {
	var countryCode types.String
	var languageCode types.String
	var htmlVersion types.Bool

	if err := countryCode.ExtractFrom(stream); err != nil {
		return countryCode, languageCode, htmlVersion, fmt.Errorf("failed to read countryCode from parameters. %s", err.Error())
	}

	if err := languageCode.ExtractFrom(stream); err != nil {
		return countryCode, languageCode, htmlVersion, fmt.Errorf("failed to read languageCode from parameters. %s", err.Error())
	}

	if err := htmlVersion.ExtractFrom(stream); err != nil {
		return countryCode, languageCode, htmlVersion, fmt.Errorf("failed to read htmlVersion from parameters. %s", err.Error())
	}

	return countryCode, languageCode, htmlVersion, nil
}
