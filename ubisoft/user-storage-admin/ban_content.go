package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go/v2"
	"github.com/PretendoNetwork/nex-go/v2/types"
	"github.com/PretendoNetwork/nex-protocols-go/v2/globals"
	user_storage_types "github.com/PretendoNetwork/nex-protocols-go/v2/ubisoft/user-storage/types"
)

// BanContentRequest holds the request parameters of User Storage Admin::BanContent
type BanContentRequest struct {
	ContentKey user_storage_types.UserContentKey
}

// WriteTo writes the BanContentRequest to the given writable
func (m BanContentRequest) WriteTo(writable types.Writable) {
	m.ContentKey.WriteTo(writable)
}

// ExtractFrom extracts the BanContentRequest from the given readable
func (m *BanContentRequest) ExtractFrom(readable types.Readable) error {
	if err := m.ContentKey.ExtractFrom(readable); err != nil {
		return fmt.Errorf("failed to read BanContent.contentKey. %s", err.Error())
	}

	return nil
}

func (protocol *Protocol) handleBanContent(packet nex.PacketInterface) {
	if protocol.BanContent == nil {
		err := nex.NewError(nex.ResultCodes.Core.NotImplemented, "User Storage Admin::BanContent not implemented")

		globals.Logger.Warning(err.Message)
		globals.RespondError(packet, ProtocolID, err)

		return
	}

	request := packet.RMCMessage()
	callID := request.CallID
	parameters := request.Parameters
	endpoint := packet.Sender().Endpoint()
	parametersStream := nex.NewByteStreamIn(parameters, endpoint.LibraryVersions(), endpoint.ByteStreamSettings())

	var params BanContentRequest

	if err := params.ExtractFrom(parametersStream); err != nil {
		_, rmcError := protocol.BanContent(err, packet, callID, params.ContentKey)
		if rmcError != nil {
			globals.RespondError(packet, ProtocolID, rmcError)
		}

		return
	}

	rmcMessage, rmcError := protocol.BanContent(nil, packet, callID, params.ContentKey)
	if rmcError != nil {
		globals.RespondError(packet, ProtocolID, rmcError)
		return
	}

	globals.Respond(packet, rmcMessage)
}
