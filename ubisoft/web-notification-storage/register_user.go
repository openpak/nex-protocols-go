package protocol

import (
	nex "github.com/PretendoNetwork/nex-go/v2"
	"github.com/PretendoNetwork/nex-go/v2/types"
	"github.com/PretendoNetwork/nex-protocols-go/v2/globals"
)

// RegisterUserRequest holds the request parameters of Web Notifications Storage::RegisterUser
type RegisterUserRequest struct{}

// WriteTo writes the RegisterUserRequest to the given writable
func (m RegisterUserRequest) WriteTo(writable types.Writable) {
	_ = writable
}

// ExtractFrom extracts the RegisterUserRequest from the given readable
func (m *RegisterUserRequest) ExtractFrom(readable types.Readable) error {
	_ = readable
	return nil
}

func (protocol *Protocol) handleRegisterUser(packet nex.PacketInterface) {
	if protocol.RegisterUser == nil {
		err := nex.NewError(nex.ResultCodes.Core.NotImplemented, "Web Notifications Storage::RegisterUser not implemented")

		globals.Logger.Warning(err.Message)
		globals.RespondError(packet, ProtocolID, err)

		return
	}

	request := packet.RMCMessage()
	callID := request.CallID

	rmcMessage, rmcError := protocol.RegisterUser(nil, packet, callID)
	if rmcError != nil {
		globals.RespondError(packet, ProtocolID, rmcError)
		return
	}

	globals.Respond(packet, rmcMessage)
}
