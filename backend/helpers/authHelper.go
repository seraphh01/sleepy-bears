package helpers

import (
	"errors"

	"github.com/gin-gonic/gin"
)

//CheckUserType renews the user tokens when they login
func CheckUserType(c *gin.Context, role string) (err error) {
	userType := c.GetString("UserType")
	err = nil
	if userType != role {
		err = errors.New("Unauthorized to access this resource")
		return err
	}

	return err
}

//MatchUserTypeToUid only allows the user to access their data and no other data. Only the admin can access all user data
func MatchUserTypeToUid(c *gin.Context, userId string) (err error) {
	userType := c.GetString("UserType")
	uid := c.GetString("uid")
	err = nil

	if userType == "STUDENT" && uid != userId {
		err = errors.New("Unauthorized to access this resource")
		return err
	}
	err = CheckUserType(c, userType)
	return err
}
