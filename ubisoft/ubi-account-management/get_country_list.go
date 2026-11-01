package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go/v2"
	"github.com/PretendoNetwork/nex-go/v2/types"
	"github.com/PretendoNetwork/nex-protocols-go/v2/globals"
)

func (protocol *Protocol) handleGetCountryList(packet nex.PacketInterface) {
	if protocol.GetCountryList == nil {
		err := nex.NewError(nex.ResultCodes.Core.NotImplemented, "UbiAccountManagement::GetCountryList not implemented")

		globals.Logger.Warning(err.Message)
		globals.RespondError(packet, ProtocolID, err)

		return
	}

	request := packet.RMCMessage()
	callID := request.CallID
	parameters := request.Parameters
	endpoint := packet.Sender().Endpoint()
	parametersStream := nex.NewByteStreamIn(parameters, endpoint.LibraryVersions(), endpoint.ByteStreamSettings())

	languageCode, err := parseGetCountryListRequest(parametersStream)
	if err != nil {
		_, rmcError := protocol.GetCountryList(err, packet, callID, languageCode)
		if rmcError != nil {
			globals.RespondError(packet, ProtocolID, rmcError)
		}

		return
	}

	rmcMessage, rmcError := protocol.GetCountryList(nil, packet, callID, languageCode)
	if rmcError != nil {
		globals.RespondError(packet, ProtocolID, rmcError)
		return
	}

	globals.Respond(packet, rmcMessage)
}

// parseGetCountryListRequest decodes the parameters of Ubi Account Management::GetCountryList from the given stream
func parseGetCountryListRequest(stream *nex.ByteStreamIn) (types.String, error) {
	var languageCode types.String

	if err := languageCode.ExtractFrom(stream); err != nil {
		return languageCode, fmt.Errorf("failed to read languageCode from parameters. %s", err.Error())
	}

	return languageCode, nil
}
