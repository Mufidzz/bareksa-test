package response

type ErrorResponse struct {
	Success bool        `json:"success"`
	Message string      `json:"message"`
	Type    int         `json:"type"` // 0 = Error, 1 = Warning
	Data    interface{} `json:"data"`
}
