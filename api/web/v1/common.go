package v1

type Common struct {
	Lang string `v:"required|in:zh-CN,zh-TW,en#语言必填|语言传送错误" dc:"语言" json:"lang" form:"lang" example:"zh_CN"`
}
