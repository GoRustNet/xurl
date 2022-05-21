package errs

var typeMsgMp = map[Type]string{
	TypeNone:      "OK",
	TypeCommon:    "发生错误",
	TypeDb:        "数据库操作失败",
	TypeExists:    "记录已存在",
	TypeBcrypt:    "加密失败",
	TypeNotExists: "记录不存在",
}

func getDefinedMsg(types Type) string {
	msg, ok := typeMsgMp[types]
	if ok {
		return msg
	}
	return "发生错误"
}
