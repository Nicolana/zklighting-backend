package user

import (
	"strconv"

	. "zklighting-backend/handler"
	"zklighting-backend/model"
	"zklighting-backend/pkg/errno"

	"github.com/gin-gonic/gin"
)

// Delete delete an user by the user identifier.
func Delete(c *gin.Context) {
	userId, _ := strconv.Atoi(c.Param("id"))
	if err := model.DeleteUser(int64(userId)); err != nil {
		SendResponse(c, errno.ErrDatabase, nil)
		return
	}

	SendResponse(c, nil, nil)
}
