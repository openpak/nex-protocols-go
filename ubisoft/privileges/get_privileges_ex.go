package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go/v2"
	"github.com/PretendoNetwork/nex-go/v2/types"
	"github.com/PretendoNetwork/nex-protocols-go/v2/globals"
)

func (protocol *Protocol) handleGetPrivilegesEx(packet nex.PacketInterface) {
	if protocol.GetPrivilegesEx == nil {
		err := nex.NewError(nex.ResultCodes.Core.NotImplemented, "Privileges::GetPrivilegesEx not implemented")

		globals.Logger.Warning(err.Message)
		globals.RespondError(packet, ProtocolID, err)

		return
	}

	request := packet.RMCMessage()
	callID := request.CallID
	parameters := request.Parameters
	endpoint := packet.Sender().Endpoint()
	parametersStream := nex.NewByteStreamIn(parameters, endpoint.LibraryVersions(), endpoint.ByteStreamSettings())

	localeCode, err := parseGetPrivilegesExRequest(parametersStream)
	if err != nil {
		_, rmcError := protocol.GetPrivilegesEx(err, packet, callID, localeCode)
		if rmcError != nil {
			globals.RespondError(packet, ProtocolID, rmcError)
		}

		return
	}

	rmcMessage, rmcError := protocol.GetPrivilegesEx(nil, packet, callID, localeCode)
	if rmcError != nil {
		globals.RespondError(packet, ProtocolID, rmcError)
		return
	}

	globals.Respond(packet, rmcMessage)
}

// parseGetPrivilegesExRequest decodes the parameters of Privileges::GetPrivilegesEx from the given stream
func parseGetPrivilegesExRequest(stream *nex.ByteStreamIn) (types.String, error) {
	var localeCode types.String

	if err := localeCode.ExtractFrom(stream); err != nil {
		return localeCode, fmt.Errorf("failed to read localeCode from parameters. %s", err.Error())
	}

	return localeCode, nil
}
