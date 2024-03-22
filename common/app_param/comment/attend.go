package comment

import (
	"fmt"
	"github.com/juetun/base-wrapper/lib/base"
	"strconv"
	"strings"
)

type (
	ArgHasAttendStatus struct {
		CurrentUid int64   `json:"current_uid" form:"current_uid"`
		TargetUid  string  `json:"target_uid" form:"target_uid"` //多个用逗号隔开
		TargetUids []int64 `json:"-" form:"-"`
	}
	ResultHasAttendStatus map[int64]bool
)

func (r *ArgHasAttendStatus) Default(ctx *base.Context) (err error) {
	r.TargetUids = make([]int64, 0)
	var (
		uid int64
	)
	if r.TargetUid != "" {
		taUid := strings.Split(r.TargetUid, ",")
		r.TargetUids = make([]int64, 0, len(taUid))
		for _, item := range taUid {
			if uid, err = strconv.ParseInt(item, 10, 64); err != nil {
				err = fmt.Errorf("target_uid参数格式错误")
				return
			}
			if uid == 0 {
				continue
			}
			r.TargetUids = append(r.TargetUids, uid)
		}
	}

	return
}
