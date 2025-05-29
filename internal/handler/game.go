package handler

import (
	"context"

	"github.com/danielgtaylor/huma/v2"
	_ "github.com/danielgtaylor/huma/v2/formats/cbor"
	types "github.com/thaynaCaixeta/lucky-admin/internal/api"
	"github.com/thaynaCaixeta/lucky-admin/internal/service"
)

type GameHandler interface {
	RegisterGameEndpoints(api huma.API)
}

type gameHandler struct {
	svc service.GameService
}

func NewGameHandler(svc service.GameService) GameHandler {
	return &gameHandler{
		svc: svc,
	}
}

func (h *gameHandler) RegisterGameEndpoints(api huma.API) {
	huma.Post(api, "/game", func(ctx context.Context, req *types.NewGameRequest) (*types.NewGameResponse, error) {
		stored, err := h.svc.StartNewGame(req.Body.NumRounds, req.Body.ClosesAt, req.Body.CreatedBy)
		if err != nil {
			return nil, err
		}

		res := &types.NewGameResponse{
			Body: types.NewGameResponseBody{
				Id:        string(stored.Id),
				NumRounds: stored.NumRounds,
				CreatedAt: stored.CreatedAt,
				ClosesAt:  stored.ClosesAt,
				Status:    stored.Status.String(),
				CreatedBy: stored.CreatedBy,
			},
		}

		return res, nil
	})
}
