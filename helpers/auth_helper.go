package helpers

import (
	"errors"

	"github.com/gin-gonic/gin"
)

func CheckUserType(c *gin.Context, role string) (err error) {
	userType := c.GetString("user_type")
	err = nil

	if userType != role {
		err = errors.New("unauthorized access")
		return err
	}
	return err

}

func MatchUserTypeToUId(c *gin.Context, userId string) (err error) {
	userType := c.GetString("user_type")
	uId := c.GetString("uId")
	err = nil

	if userType == "USER" && uId != userId {
		err = errors.New("unauthorized access")
		return err
	}

	err = CheckUserType(c, userType)
	return err
}
