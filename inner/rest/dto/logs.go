package dto

type GetLogs struct {
	ID     string `uri:"id" binding:"required"`
    Limit  int    `form:"limit"`
}
