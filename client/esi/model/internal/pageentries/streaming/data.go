package streaming

import (
	"time"

	"github.com/WiggidyW/weve-esi/cache"
	"github.com/WiggidyW/weve-esi/util"
)

type HeadRepWithChan[E any] struct {
	NumPages int
	Expires  time.Time // initially head expires
	ChanRecv util.ChanRecvResult[cache.ExpirableData[[]E]]
}

func (hrwc *HeadRepWithChan[E]) RecvUpdateExpires() ([]E, error) {
	rep, err := hrwc.ChanRecv.Recv()
	if err != nil {
		return nil, err
	} else {
		pageExpires := rep.Expires()
		if pageExpires.After(hrwc.Expires) {
			hrwc.Expires = pageExpires
		}
	}
	return rep.Data(), nil
}