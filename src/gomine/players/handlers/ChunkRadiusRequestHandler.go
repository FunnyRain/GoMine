package handlers

import (
	"gomine/players"
	"gomine/net/info"
	"gomine/interfaces"
	"goraklib/server"
	"gomine/net/packets"
	"fmt"
)

type ChunkRadiusRequestHandler struct {
	*players.PacketHandler
}

func NewChunkRadiusRequestHandler() ChunkRadiusRequestHandler {
	return ChunkRadiusRequestHandler{players.NewPacketHandler(info.RequestChunkRadiusPacket)}
}

func (handler ChunkRadiusRequestHandler) Handle(packet interfaces.IPacket, player interfaces.IPlayer, session *server.Session, server interfaces.IServer) bool {
	if chunkRadiusPacket, ok := packet.(*packets.ChunkRadiusRequestPacket); ok {
		player.SetViewDistance(uint(chunkRadiusPacket.Radius))

		var pk = packets.NewChunkRadiusUpdatedPacket()
		pk.Radius = uint32(player.GetViewDistance())
		server.GetRakLibAdapter().SendPacket(pk, session)

		pk2 := packets.NewPlayStatusPacket()
		pk2.Status = 3
		server.GetRakLibAdapter().SendPacket(pk2, session)

		fmt.Println("Sent chunk radius updated packet with: ", uint32(player.GetViewDistance()))
	}

	return true
}