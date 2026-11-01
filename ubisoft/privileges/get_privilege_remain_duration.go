package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go/v2"
	"github.com/PretendoNetwork/nex-go/v2/types"
	"github.com/PretendoNetwork/nex-protocols-go/v2/globals"
)

func (protocol *Protocol) handleGetPrivilegeRemainDuration(packet nex.PacketInterface) {
	if protocol.GetPrivilegeRemainDuration == nil {
		err := nex.NewError(nex.ResultCodes.Core.NotImplemented, "Privileges::GetPrivilegeRemainDuration not implemented")

		globals.Logger.Warning(err.Message)
		globals.RespondError(packet, ProtocolID, err)

		return
	}

	request := packet.RMCMessage()
	callID := request.CallID
	parameters := request.Parameters
	endpoint := packet.Sender().Endpoint()
	parametersStream := nex.NewByteStreamIn(parameters, endpoint.LibraryVersions(), endpoint.ByteStreamSettings())

	privilegeID, err := parseGetPrivilegeRemainDurationRequest(parametersStream)
	if err != nil {
		_, rmcError := protocol.GetPrivilegeRemainDuration(err, packet, callID, privilegeID)
		if rmcError != nil {
			globals.RespondError(packet, ProtocolID, rmcError)
		}

		return
	}

	rmcMessage, rmcError := protocol.GetPrivilegeRemainDuration(nil, packet, callID, privilegeID)
	if rmcError != nil {
		globals.RespondError(packet, ProtocolID, rmcError)
		return
	}

	globals.Respond(packet, rmcMessage)
}

// parseGetPrivilegeRemainDurationRequest decodes the parameters of Privileges::GetPrivilegeRemainDuration from the given stream
func parseGetPrivilegeRemainDurationRequest(stream *nex.ByteStreamIn) (types.UInt32, error) {
	var privilegeID types.UInt32

	if err := privilegeID.ExtractFrom(stream); err != nil {
		return privilegeID, fmt.Errorf("failed to read privilegeID from parameters. %s", err.Error())
	}

	return privilegeID, nil
}
