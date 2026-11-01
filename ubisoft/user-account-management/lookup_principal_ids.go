package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go/v2"
	"github.com/PretendoNetwork/nex-go/v2/types"
	"github.com/PretendoNetwork/nex-protocols-go/v2/globals"
)

func (protocol *Protocol) handleLookupPrincipalIDs(packet nex.PacketInterface) {
	if protocol.LookupPrincipalIDs == nil {
		err := nex.NewError(nex.ResultCodes.Core.NotImplemented, "UserAccountManagement::LookupPrincipalIDs not implemented")

		globals.Logger.Warning(err.Message)
		globals.RespondError(packet, ProtocolID, err)

		return
	}

	request := packet.RMCMessage()
	callID := request.CallID
	parameters := request.Parameters
	endpoint := packet.Sender().Endpoint()
	parametersStream := nex.NewByteStreamIn(parameters, endpoint.LibraryVersions(), endpoint.ByteStreamSettings())

	firstPartyIds, platformId, err := parseLookupPrincipalIDsRequest(parametersStream)
	if err != nil {
		_, rmcError := protocol.LookupPrincipalIDs(err, packet, callID, firstPartyIds, platformId)
		if rmcError != nil {
			globals.RespondError(packet, ProtocolID, rmcError)
		}

		return
	}

	rmcMessage, rmcError := protocol.LookupPrincipalIDs(nil, packet, callID, firstPartyIds, platformId)
	if rmcError != nil {
		globals.RespondError(packet, ProtocolID, rmcError)
		return
	}

	globals.Respond(packet, rmcMessage)
}

// parseLookupPrincipalIDsRequest decodes the parameters of User Account Management::LookupPrincipalIDs from the given stream
func parseLookupPrincipalIDsRequest(stream *nex.ByteStreamIn) (types.List[types.String], types.UInt32, error) {
	var firstPartyIds types.List[types.String]
	var platformId types.UInt32

	if err := firstPartyIds.ExtractFrom(stream); err != nil {
		return firstPartyIds, platformId, fmt.Errorf("failed to read firstPartyIds from parameters. %s", err.Error())
	}

	if err := platformId.ExtractFrom(stream); err != nil {
		return firstPartyIds, platformId, fmt.Errorf("failed to read platformId from parameters. %s", err.Error())
	}

	return firstPartyIds, platformId, nil
}
