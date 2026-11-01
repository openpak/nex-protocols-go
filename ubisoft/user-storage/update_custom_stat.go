package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go/v2"
	"github.com/PretendoNetwork/nex-go/v2/types"
	"github.com/PretendoNetwork/nex-protocols-go/v2/globals"
	user_storage_types "github.com/PretendoNetwork/nex-protocols-go/v2/ubisoft/user-storage/types"
)

// UpdateCustomStatRequest holds the request parameters of User Storage::UpdateCustomStat
type UpdateCustomStatRequest struct {
	ContentKey user_storage_types.UserContentKey
	StatID     types.UInt16
	IncValue   types.Int64
}

// WriteTo writes the UpdateCustomStatRequest to the given writable
func (m UpdateCustomStatRequest) WriteTo(writable types.Writable) {
	m.ContentKey.WriteTo(writable)
	m.StatID.WriteTo(writable)
	m.IncValue.WriteTo(writable)
}

// ExtractFrom extracts the UpdateCustomStatRequest from the given readable
func (m *UpdateCustomStatRequest) ExtractFrom(readable types.Readable) error {
	if err := m.ContentKey.ExtractFrom(readable); err != nil {
		return fmt.Errorf("failed to read UpdateCustomStat.contentKey. %s", err.Error())
	}

	if err := m.StatID.ExtractFrom(readable); err != nil {
		return fmt.Errorf("failed to read UpdateCustomStat.statID. %s", err.Error())
	}

	if err := m.IncValue.ExtractFrom(readable); err != nil {
		return fmt.Errorf("failed to read UpdateCustomStat.incValue. %s", err.Error())
	}

	return nil
}

func (protocol *Protocol) handleUpdateCustomStat(packet nex.PacketInterface) {
	if protocol.UpdateCustomStat == nil {
		err := nex.NewError(nex.ResultCodes.Core.NotImplemented, "User Storage::UpdateCustomStat not implemented")

		globals.Logger.Warning(err.Message)
		globals.RespondError(packet, ProtocolID, err)

		return
	}

	request := packet.RMCMessage()
	callID := request.CallID
	parameters := request.Parameters
	endpoint := packet.Sender().Endpoint()
	parametersStream := nex.NewByteStreamIn(parameters, endpoint.LibraryVersions(), endpoint.ByteStreamSettings())

	var params UpdateCustomStatRequest

	if err := params.ExtractFrom(parametersStream); err != nil {
		_, rmcError := protocol.UpdateCustomStat(err, packet, callID, params.ContentKey, params.StatID, params.IncValue)
		if rmcError != nil {
			globals.RespondError(packet, ProtocolID, rmcError)
		}

		return
	}

	rmcMessage, rmcError := protocol.UpdateCustomStat(nil, packet, callID, params.ContentKey, params.StatID, params.IncValue)
	if rmcError != nil {
		globals.RespondError(packet, ProtocolID, rmcError)
		return
	}

	globals.Respond(packet, rmcMessage)
}
