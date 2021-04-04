package user

import (
	"fmt"

	. "zklighting-backend/handler"
	"zklighting-backend/pkg/errno"

	"github.com/gin-gonic/gin"
	"github.com/lexkong/log"
)

// Create creates a new user account
func Create(c *gin.Context) {
	var r CreateRequest
	var err error
	if err := c.Bind(&r); err != nil {
		// c.JSON(http.StatusOK, gin.H{"error": errno.ErrBind})
		SendResponse(c, errno.ErrBind, nil)
		return
	}

	admin2 := c.Param("username")
	log.Infof("URL username: %s", admin2)

	desc := c.Query("desc")
	log.Infof("URL key param desc: %s", desc)

	contentType := c.GetHeader("Content-Type")
	log.Infof("Header Content-Type: %s", contentType)

	log.Debugf("username is: [%s], password is [%s]", r.Username, r.Password)
	if r.Username == "" {
		SendResponse(c, errno.New(errno.ErrUserNotFound, fmt.Errorf("username can not found in db: xx.xx.xx.xx")), nil)
		// err = errno.New(errno.ErrUserNotFound, fmt.Errorf("username can not found in db: xx.xx.xx.xx")).Add("This is add message.")
		log.Errorf(err, "Get an error")
	}

	// if errno.IsErrUserNotFound(err) {
	// log.Debug("err type is ErrUserNotFound")
	// }

	if r.Password == "" {
		SendResponse(c, fmt.Errorf("password is empty"), nil)
		// err = fmt.Error("password is empty")
	}

	rsp := CreateResponse{
		Username: r.Username,
	}

	// Show the user infomation
	SendResponse(c, nil, rsp)

	// code, message := errno.DecodeErr(err)
	// c.JSON(http.StatusOK, gin.H{"code": code, "message": message})
}
