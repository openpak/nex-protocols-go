package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go/v2"
	"github.com/PretendoNetwork/nex-go/v2/types"
	"github.com/PretendoNetwork/nex-protocols-go/v2/globals"
)

func (protocol *Protocol) handleActivateKeyWithExpectedPrivileges(packet nex.PacketInterface) {
	if protocol.ActivateKeyWithExpectedPrivileges == nil {
		err := nex.NewError(nex.ResultCodes.Core.NotImplemented, "Privileges::ActivateKeyWithExpectedPrivileges not implemented")

		globals.Logger.Warning(err.Message)
		globals.RespondError(packet, ProtocolID, err)

		return
	}

	request := packet.RMCMessage()
	callID := request.CallID
	parameters := request.Parameters
	endpoint := packet.Sender().Endpoint()
	parametersStream := nex.NewByteStreamIn(parameters, endpoint.LibraryVersions(), endpoint.ByteStreamSettings())

	uniqueKey, languageCode, expectedPrivilegeIDs, err := parseActivateKeyWithExpectedPrivilegesRequest(parametersStream)
	if err != nil {
		_, rmcError := protocol.ActivateKeyWithExpectedPrivileges(err, packet, callID, uniqueKey, languageCode, expectedPrivilegeIDs)
		if rmcError != nil {
			globals.RespondError(packet, ProtocolID, rmcError)
		}

		return
	}

	rmcMessage, rmcError := protocol.ActivateKeyWithExpectedPrivileges(nil, packet, callID, uniqueKey, languageCode, expectedPrivilegeIDs)
	if rmcError != nil {
		globals.RespondError(packet, ProtocolID, rmcError)
		return
	}

	globals.Respond(packet, rmcMessage)
}

// parseActivateKeyWithExpectedPrivilegesRequest decodes the parameters of Privileges::ActivateKeyWithExpectedPrivileges from the given stream
func parseActivateKeyWithExpectedPrivilegesRequest(stream *nex.ByteStreamIn) (types.String, types.String, types.List[types.UInt32], error) {
	var uniqueKey types.String
	var languageCode types.String
	var expectedPrivilegeIDs types.List[types.UInt32]

	if err := uniqueKey.ExtractFrom(stream); err != nil {
		return uniqueKey, languageCode, expectedPrivilegeIDs, fmt.Errorf("failed to read uniqueKey from parameters. %s", err.Error())
	}

	if err := languageCode.ExtractFrom(stream); err != nil {
		return uniqueKey, languageCode, expectedPrivilegeIDs, fmt.Errorf("failed to read languageCode from parameters. %s", err.Error())
	}

	if err := expectedPrivilegeIDs.ExtractFrom(stream); err != nil {
		return uniqueKey, languageCode, expectedPrivilegeIDs, fmt.Errorf("failed to read expectedPrivilegeIDs from parameters. %s", err.Error())
	}

	return uniqueKey, languageCode, expectedPrivilegeIDs, nil
}
