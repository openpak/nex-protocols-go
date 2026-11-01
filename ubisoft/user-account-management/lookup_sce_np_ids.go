package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go/v2"
	"github.com/PretendoNetwork/nex-go/v2/types"
	"github.com/PretendoNetwork/nex-protocols-go/v2/globals"
)

func (protocol *Protocol) handleLookupSceNpIds(packet nex.PacketInterface) {
	if protocol.LookupSceNpIds == nil {
		err := nex.NewError(nex.ResultCodes.Core.NotImplemented, "UserAccountManagement::LookupSceNpIds not implemented")

		globals.Logger.Warning(err.Message)
		globals.RespondError(packet, ProtocolID, err)

		return
	}

	request := packet.RMCMessage()
	callID := request.CallID
	parameters := request.Parameters
	endpoint := packet.Sender().Endpoint()
	parametersStream := nex.NewByteStreamIn(parameters, endpoint.LibraryVersions(), endpoint.ByteStreamSettings())

	pids, err := parseLookupSceNpIdsRequest(parametersStream)
	if err != nil {
		_, rmcError := protocol.LookupSceNpIds(err, packet, callID, pids)
		if rmcError != nil {
			globals.RespondError(packet, ProtocolID, rmcError)
		}

		return
	}

	rmcMessage, rmcError := protocol.LookupSceNpIds(nil, packet, callID, pids)
	if rmcError != nil {
		globals.RespondError(packet, ProtocolID, rmcError)
		return
	}

	globals.Respond(packet, rmcMessage)
}

// parseLookupSceNpIdsRequest decodes the parameters of User Account Management::LookupSceNpIds from the given stream
func parseLookupSceNpIdsRequest(stream *nex.ByteStreamIn) (types.List[types.UInt32], error) {
	var pids types.List[types.UInt32]

	if err := pids.ExtractFrom(stream); err != nil {
		return pids, fmt.Errorf("failed to read pids from parameters. %s", err.Error())
	}

	return pids, nil
}
