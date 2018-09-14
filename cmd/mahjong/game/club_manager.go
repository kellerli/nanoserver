package game

import (
	"github.com/lonnng/nanoserver/internal/protocol"

	"github.com/lonnng/nanoserver/db"
	"github.com/lonnng/nanoserver/internal/async"

	"github.com/lonnng/nano/component"
	"github.com/lonnng/nano/session"
)

type ClubManager struct {
	component.Base
}

func (c *ClubManager) ApplyClub(s *session.Session, payload *protocol.ApplyClubRequest) error {
	mid := s.MID()
	logger.Debugf("玩家申请加入俱乐部，UID=%d，俱乐部ID=%d", s.UID(), payload.ClubId)
	async.Run(func() {
		if err := db.ApplyClub(s.UID(), payload.ClubId); err != nil {
			s.ResponseMID(mid, &protocol.ErrorResponse{
				Code:  -1,
				Error: err.Error(),
			})
		} else {
			s.ResponseMID(mid, &protocol.SuccessResponse)
		}
	})
	return nil
}
