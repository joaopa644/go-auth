package helloworldrequestmodel

type HelloWorldRequestModel struct {
	Message       string `form:"helloWorldRequestModel" json:"message" binding:"required"`
	NameRequester string `form:"helloWorldRequestModel" json:"nameRequester" binding:"required"`
}
