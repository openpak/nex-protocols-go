package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go/v2"
	"github.com/PretendoNetwork/nex-go/v2/types"
	"github.com/PretendoNetwork/nex-protocols-go/v2/globals"
	user_storage_types "github.com/PretendoNetwork/nex-protocols-go/v2/ubisoft/user-storage/types"
)

// FlagContentAsVerifiedRequest holds the request parameters of User Storage Admin::FlagContentAsVerified
type FlagContentAsVerifiedRequest struct {
	ContentKey user_storage_types.UserContentKey
}

// WriteTo writes the FlagContentAsVerifiedRequest to the given writable
func (m FlagContentAsVerifiedRequest) WriteTo(writable types.Writable) {
	m.ContentKey.WriteTo(writable)
}

// ExtractFrom extracts the FlagContentAsVerifiedRequest from the given readable
func (m *FlagContentAsVerifiedRequest) ExtractFrom(readable types.Readable) error {
	if err := m.ContentKey.ExtractFrom(readable); err != nil {
		return fmt.Errorf("failed to read FlagContentAsVerified.contentKey. %s", err.Error())
	}

	return nil
}

func (protocol *Protocol) handleFlagContentAsVerified(packet nex.PacketInterface) {
	if protocol.FlagContentAsVerified == nil {
		err := nex.NewError(nex.ResultCodes.Core.NotImplemented, "User Storage Admin::FlagContentAsVerified not implemented")

		globals.Logger.Warning(err.Message)
		globals.RespondError(packet, ProtocolID, err)

		return
	}

	request := packet.RMCMessage()
	callID := request.CallID
	parameters := request.Parameters
	endpoint := packet.Sender().Endpoint()
	parametersStream := nex.NewByteStreamIn(parameters, endpoint.LibraryVersions(), endpoint.ByteStreamSettings())

	var params FlagContentAsVerifiedRequest

	if err := params.ExtractFrom(parametersStream); err != nil {
		_, rmcError := protocol.FlagContentAsVerified(err, packet, callID, params.ContentKey)
		if rmcError != nil {
			globals.RespondError(packet, ProtocolID, rmcError)
		}

		return
	}

	rmcMessage, rmcError := protocol.FlagContentAsVerified(nil, packet, callID, params.ContentKey)
	if rmcError != nil {
		globals.RespondError(packet, ProtocolID, rmcError)
		return
	}

	globals.Respond(packet, rmcMessage)
}
