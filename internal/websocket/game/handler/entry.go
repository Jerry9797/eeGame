package handler

import (
	"context"
	"eeGame/internal/dao"
	"eeGame/internal/websocket/framework/net"
	"eeGame/internal/websocket/game/models/request"
	"encoding/json"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
)

type EntryHandler struct {
}

func NewEntryHandler() *EntryHandler {
	return &EntryHandler{}
}

func (h *EntryHandler) Entry(session *net.Session, body []byte) (any, error) {
	logrus.Info("Entry request: %v\n", string(body))
	var entryR request.EntryReq
	err := json.Unmarshal(body, &entryR)
	if err != nil {
		return nil, err
	}
	// 校验token合法性
	username := entryR.User.Username
	user, err := dao.User.GetUserByUsername(context.Background(), username)
	if err != nil && user == nil {
		return nil, errors.New("非法用户")
	}
	return user, nil
}
