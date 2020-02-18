package internl

import (
	"strconv"
)

func (pm *PushMessage) GetCodeStr() string {
	return strconv.FormatInt(int64(pm.Code), 10)
}

func (mr *MessageRequest) GetCodeStr() string {
	return strconv.FormatInt(int64(mr.Code), 10)
}

