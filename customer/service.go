package customer

import (
	"context"
	"fmt"
	"math/rand"
	"sync"
	"time"

	"github.com/micro/go-micro/v2/logger"
	"gitlab.lrz.devss/semester/ob-20ws/blatt2/blatt2-grp03/api"
)

type users struct {
	name string
}
type CustomerHandlerService struct {
	user         map[int32]*users

}

const (
	maxuserid int32 = 987654321
)

func (u *CustomerHandlerService) appendANewUser(id int32, user *users) bool {
	if id != 0 && user != nil {
		(*u.getUserMap())[id] = user
		return true
	}
	return false
}
// getUserMap will return a pointer to the current user map in order to work in that. //
func (u *CustomerHandlerService) getUserMap() *map[int32]*users {
	return &u.user
}
func (u *CustomerHandlerService) containsID(id int32) bool {
	_, inMap := (*u.getUserMap())[id]
	return inMap
}

func (u *CustomerHandlerService) getRandomUserID(length int32) int32 {
	rand.Seed(time.Now().UnixNano())
	for {
		potantialID := rand.Int31n(length)
		if !u.containsID(potantialID) {
			return potantialID
		}
	}
}



func (u *CustomerHandlerService) CreateUser(context context.Context, request *api.CreateUserRequest, response *api.CreatedUserResponse) error {
	if request.GetName() != "" {
		uid := u.getRandomUserID(maxuserid)
		if u.appendANewUser(uid, &users{name: request.GetName()}) {
			response.UserId =uid
			return nil
		}
	}
	return fmt.Errorf("cannot create user with empty name")
}
