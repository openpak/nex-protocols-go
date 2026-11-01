package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go/v2"
	"github.com/PretendoNetwork/nex-go/v2/types"
	"github.com/PretendoNetwork/nex-protocols-go/v2/globals"
)

func (protocol *Protocol) handleLookupFirstPartyIds(packet nex.PacketInterface) {
	if protocol.LookupFirstPartyIds == nil {
		err := nex.NewError(nex.ResultCodes.Core.NotImplemented, "UserAccountManagement::LookupFirstPartyIds not implemented")

		globals.Logger.Warning(err.Message)
		globals.RespondError(packet, ProtocolID, err)

		return
	}

	request := packet.RMCMessage()
	callID := request.CallID
	parameters := request.Parameters
	endpoint := packet.Sender().Endpoint()
	parametersStream := nex.NewByteStreamIn(parameters, endpoint.LibraryVersions(), endpoint.ByteStreamSettings())

	pids, platformId, err := parseLookupFirstPartyIdsRequest(parametersStream)
	if err != nil {
		_, rmcError := protocol.LookupFirstPartyIds(err, packet, callID, pids, platformId)
		if rmcError != nil {
			globals.RespondError(packet, ProtocolID, rmcError)
		}

		return
	}

	rmcMessage, rmcError := protocol.LookupFirstPartyIds(nil, packet, callID, pids, platformId)
	if rmcError != nil {
		globals.RespondError(packet, ProtocolID, rmcError)
		return
	}

	globals.Respond(packet, rmcMessage)
}

// parseLookupFirstPartyIdsRequest decodes the parameters of User Account Management::LookupFirstPartyIds from the given stream
func parseLookupFirstPartyIdsRequest(stream *nex.ByteStreamIn) (types.List[types.UInt32], types.UInt32, error) {
	var pids types.List[types.UInt32]
	var platformId types.UInt32

	if err := pids.ExtractFrom(stream); err != nil {
		return pids, platformId, fmt.Errorf("failed to read pids from parameters. %s", err.Error())
	}

	if err := platformId.ExtractFrom(stream); err != nil {
		return pids, platformId, fmt.Errorf("failed to read platformId from parameters. %s", err.Error())
	}

	return pids, platformId, nil
}
