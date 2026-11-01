package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go/v2"
	"github.com/PretendoNetwork/nex-go/v2/types"
	"github.com/PretendoNetwork/nex-protocols-go/v2/globals"
	user_storage_types "github.com/PretendoNetwork/nex-protocols-go/v2/ubisoft/user-storage/types"
)

// GetFavouritesRequest holds the request parameters of User Storage::GetFavourites
type GetFavouritesRequest struct {
	ContentTypes types.List[types.UInt32]
}

// WriteTo writes the GetFavouritesRequest to the given writable
func (m GetFavouritesRequest) WriteTo(writable types.Writable) {
	m.ContentTypes.WriteTo(writable)
}

// ExtractFrom extracts the GetFavouritesRequest from the given readable
func (m *GetFavouritesRequest) ExtractFrom(readable types.Readable) error {
	if err := m.ContentTypes.ExtractFrom(readable); err != nil {
		return fmt.Errorf("failed to read GetFavourites.contentTypes. %s", err.Error())
	}

	return nil
}

// GetFavouritesResponse holds the response parameters of User Storage::GetFavourites
type GetFavouritesResponse struct {
	Favourites types.List[user_storage_types.UserContent]
}

// WriteTo writes the GetFavouritesResponse to the given writable
func (m GetFavouritesResponse) WriteTo(writable types.Writable) {
	m.Favourites.WriteTo(writable)
}

// ExtractFrom extracts the GetFavouritesResponse from the given readable
func (m *GetFavouritesResponse) ExtractFrom(readable types.Readable) error {
	if err := m.Favourites.ExtractFrom(readable); err != nil {
		return fmt.Errorf("failed to read GetFavourites.favourites. %s", err.Error())
	}

	return nil
}

func (protocol *Protocol) handleGetFavourites(packet nex.PacketInterface) {
	if protocol.GetFavourites == nil {
		err := nex.NewError(nex.ResultCodes.Core.NotImplemented, "User Storage::GetFavourites not implemented")

		globals.Logger.Warning(err.Message)
		globals.RespondError(packet, ProtocolID, err)

		return
	}

	request := packet.RMCMessage()
	callID := request.CallID
	parameters := request.Parameters
	endpoint := packet.Sender().Endpoint()
	parametersStream := nex.NewByteStreamIn(parameters, endpoint.LibraryVersions(), endpoint.ByteStreamSettings())

	var params GetFavouritesRequest

	if err := params.ExtractFrom(parametersStream); err != nil {
		_, rmcError := protocol.GetFavourites(err, packet, callID, params.ContentTypes)
		if rmcError != nil {
			globals.RespondError(packet, ProtocolID, rmcError)
		}

		return
	}

	rmcMessage, rmcError := protocol.GetFavourites(nil, packet, callID, params.ContentTypes)
	if rmcError != nil {
		globals.RespondError(packet, ProtocolID, rmcError)
		return
	}

	globals.Respond(packet, rmcMessage)
}
