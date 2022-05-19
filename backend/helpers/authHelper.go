package helpers

import (
	"errors"

	"github.com/gin-gonic/gin"
)

//CheckUserType renews the user tokens when they log in
func CheckUserType(c *gin.Context, role string) (err error) {
	userType := c.GetString("usertype")
	if userType == "CHIEF" && role == "TEACHER" {
		return nil
	}
	err = nil
	if userType != role {
		err = errors.New("unauthorized to access this resource")
		return err
	}

	return err
}

//MatchUserTypeToUid only allows the user to access their data and no other data. Only the admin can access all user data
func MatchUserTypeToUid(c *gin.Context, username_ string) (err error) {
	userType := c.GetString("usertype")
	username := c.GetString("username")
	err = nil

	if userType == "STUDENT" && username != username_ {
		err = errors.New("unauthorized to access this resource")
		return err
	}
	err = CheckUserType(c, userType)
	return err
}

func MatchUserToUsername(c *gin.Context, username_ string) (err error) {
	username := c.GetString("username")
	err = nil
	if username != username_ {
		err = errors.New("not authorized to access this resource")
		return err
	}

	return err
}
