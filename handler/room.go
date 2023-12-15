package handler

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/mustafawidiarto/go-boilerplate/handler/middleware"
	"github.com/mustafawidiarto/go-boilerplate/model"
	"github.com/mustafawidiarto/go-boilerplate/model/entity"

	"github.com/go-chi/chi/v5"
)

type room struct {
	roomUC RoomUsecase
}

// NewRoom creates a new instance of the room struct with the provided room usecase.
func NewRoom(roomUC RoomUsecase) *room {
	return &room{
		roomUC: roomUC,
	}
}

type newSessionWebhookReq struct {
	IsNewSession bool `json:"is_new_session"`
	Payload      struct {
		Room struct {
			ID              string `json:"id"`
			IDStr           string `json:"id_str"`
			IsPublicChannel bool   `json:"is_public_channel"`
			Name            string `json:"name"`
			Options         string `json:"options"`
			Participants    []struct {
				Email string `json:"email"`
			} `json:"participants"`
			RoomAvatar string `json:"room_avatar"`
			TopicID    string `json:"topic_id"`
			TopicIDStr string `json:"topic_id_str"`
			Type       string `json:"type"`
		} `json:"room"`
	} `json:"payload"`
	WebhookType string `json:"webhook_type"`
}

func (h *room) HandleNewSessionWebhook() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		var body newSessionWebhookReq
		if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
			respJSON(w, http.StatusInternalServerError, ResponseError{Message: err.Error()})
			return
		}

		room := &entity.Room{
			MultichannelRoomID: body.Payload.Room.IDStr,
		}

		if err := h.roomUC.CreateRoom(ctx, room); err != nil {
			respJSON(w, getErrStatusCode(err), ResponseError{Message: err.Error()})
			return
		}

		respJSON(w, http.StatusOK, "ok")
	}
}

func (h *room) GetRoomByID() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		id, err := strconv.Atoi(chi.URLParam(r, "id"))
		if err != nil {
			respJSON(w, http.StatusBadRequest, ResponseError{Message: model.ErrBadParamInput.Error()})
			return
		}

		room, err := h.roomUC.GetRoomByID(ctx, int64(id))
		if err != nil {
			respJSON(w, getErrStatusCode(err), ResponseError{Message: err.Error()})
			return
		}

		respJSON(w, http.StatusOK, room)
	}
}

func (h *room) HandleRoute(r chi.Router) {
	r.Post("/wh/multichannel/new_session", h.HandleNewSessionWebhook())

	r.Group(func(r chi.Router) {
		r.Use(middleware.ApiKey)
		r.Route("/api/v1", func(r chi.Router) {
			r.Get("/rooms/{id}", h.GetRoomByID())
		})
	})
}
