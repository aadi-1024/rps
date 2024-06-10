package server

import (
	"context"

	"github.com/aadi-1024/rps/protobuf"
)

type Player struct {
	move chan protobuf.Moves
	res  chan protobuf.Result
}

type GameServer struct {
	p1, p2 Player
	protobuf.UnimplementedGameServer
}

func NewGameServer() (GameServer, func()) {
	g := GameServer{}
	f := func() {
		wins_against := map[protobuf.Moves]protobuf.Moves{
			protobuf.Moves_Rock:     protobuf.Moves_Scissors,
			protobuf.Moves_Paper:    protobuf.Moves_Rock,
			protobuf.Moves_Scissors: protobuf.Moves_Rock,
		}

		m1 := <-g.p1.move
		m2 := <-g.p2.move

		if m1 == m2 || m1 == protobuf.Moves_Quit || m2 == protobuf.Moves_Quit {
			g.p1.res <- protobuf.Result_Draw
			g.p2.res <- protobuf.Result_Draw
		} else if m2 == wins_against[m1] {
			g.p1.res <- protobuf.Result_Win
			g.p2.res <- protobuf.Result_Lose
		} else {
			g.p1.res <- protobuf.Result_Lose
			g.p2.res <- protobuf.Result_Win
		}
	}
	return g, f
}

func (s GameServer) PlayMove(ctx context.Context, in *protobuf.Action) (*protobuf.Response, error) {
	player := in.GetPlayerId()
	resp := &protobuf.Response{}

	if player == 1 {
		s.p1.move <- in.GetMove()
		res := <-s.p1.res
		resp.Res = res
	} else {
		s.p2.move <- in.GetMove()
		res := <-s.p2.res
		resp.Res = res
	}

	switch resp.Res {
	case protobuf.Result_Win:
		resp.Msg = "You Win!"
	case protobuf.Result_Lose:
		resp.Msg = "You Lose!"
	case protobuf.Result_Draw:
		resp.Msg = "Draw!"
	}

	return resp, nil
}
