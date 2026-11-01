package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go/v2"
	"github.com/PretendoNetwork/nex-go/v2/types"
	"github.com/PretendoNetwork/nex-protocols-go/v2/globals"
)

func (protocol *Protocol) handleLookupUbiAccountIDsByPids(packet nex.PacketInterface) {
	if protocol.LookupUbiAccountIDsByPids == nil {
		err := nex.NewError(nex.ResultCodes.Core.NotImplemented, "UbiAccountManagement::LookupUbiAccountIDsByPids not implemented")

		globals.Logger.Warning(err.Message)
		globals.RespondError(packet, ProtocolID, err)

		return
	}

	request := packet.RMCMessage()
	callID := request.CallID
	parameters := request.Parameters
	endpoint := packet.Sender().Endpoint()
	parametersStream := nex.NewByteStreamIn(parameters, endpoint.LibraryVersions(), endpoint.ByteStreamSettings())

	pids, err := parseLookupUbiAccountIDsByPidsRequest(parametersStream)
	if err != nil {
		_, rmcError := protocol.LookupUbiAccountIDsByPids(err, packet, callID, pids)
		if rmcError != nil {
			globals.RespondError(packet, ProtocolID, rmcError)
		}

		return
	}

	rmcMessage, rmcError := protocol.LookupUbiAccountIDsByPids(nil, packet, callID, pids)
	if rmcError != nil {
		globals.RespondError(packet, ProtocolID, rmcError)
		return
	}

	globals.Respond(packet, rmcMessage)
}

// parseLookupUbiAccountIDsByPidsRequest decodes the parameters of Ubi Account Management::LookupUbiAccountIDsByPids from the given stream
func parseLookupUbiAccountIDsByPidsRequest(stream *nex.ByteStreamIn) (types.List[types.UInt32], error) {
	var pids types.List[types.UInt32]

	if err := pids.ExtractFrom(stream); err != nil {
		return pids, fmt.Errorf("failed to read pids from parameters. %s", err.Error())
	}

	return pids, nil
}
