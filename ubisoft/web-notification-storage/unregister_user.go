package protocol

import (
	nex "github.com/PretendoNetwork/nex-go/v2"
	"github.com/PretendoNetwork/nex-go/v2/types"
	"github.com/PretendoNetwork/nex-protocols-go/v2/globals"
)

// UnregisterUserRequest holds the request parameters of Web Notifications Storage::UnregisterUser
type UnregisterUserRequest struct{}

// WriteTo writes the UnregisterUserRequest to the given writable
func (m UnregisterUserRequest) WriteTo(writable types.Writable) {
	_ = writable
}

// ExtractFrom extracts the UnregisterUserRequest from the given readable
func (m *UnregisterUserRequest) ExtractFrom(readable types.Readable) error {
	_ = readable
	return nil
}

func (protocol *Protocol) handleUnregisterUser(packet nex.PacketInterface) {
	if protocol.UnregisterUser == nil {
		err := nex.NewError(nex.ResultCodes.Core.NotImplemented, "Web Notifications Storage::UnregisterUser not implemented")

		globals.Logger.Warning(err.Message)
		globals.RespondError(packet, ProtocolID, err)

		return
	}

	request := packet.RMCMessage()
	callID := request.CallID

	rmcMessage, rmcError := protocol.UnregisterUser(nil, packet, callID)
	if rmcError != nil {
		globals.RespondError(packet, ProtocolID, rmcError)
		return
	}

	globals.Respond(packet, rmcMessage)
}
