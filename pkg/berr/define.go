package berr

type BizError struct {
	ErrCode int64  `json:"err_code"`
	Msg     string `json:"msg"`
	Detail  string `json:"detail"`
}

func (b *BizError) Error() string {
	return b.Msg
}

func NewBizError(code int64, msg string) *BizError {
	return &BizError{
		ErrCode: code,
		Msg:     msg,
	}
}

func (b *BizError) CloneWithError(err error) *BizError {
	return &BizError{
		ErrCode: b.ErrCode,
		Msg:     b.Msg,
		Detail:  err.Error(),
	}
}

func (b *BizError) AppendMsg(s string) *BizError {
	return &BizError{
		ErrCode: b.ErrCode,
		Msg:     b.Msg + s,
		Detail:  b.Detail,
	}
}

func (b *BizError) SetDetail(s string) *BizError {
	return &BizError{
		ErrCode: b.ErrCode,
		Msg:     b.Msg,
		Detail:  s,
	}
}
