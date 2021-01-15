package customer

import (
	"context"
	"fmt"
	"math/rand"
	// "sync"
	"time"
	// "github.com/micro/go-micro/v2/logger"
	"blatt2-grp03/api"
)


type CustomerHandlerService struct {
	customer         map[int32]string

}

const (
	maxuserid int32 = 987654321
)

func (u *CustomerHandlerService) appendANewUser(id int32, user string) bool {
	if id != 0 && user != "" {
		(*u.getUserMap())[id] = user
		return true
	}
	return false
}
// getUserMap will return a pointer to the current user map in order to work in that. //
func (u *CustomerHandlerService) getUserMap() *map[int32]string {
	return &u.customer
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

func CreateNewCustomerHandleInstance() *CustomerHandlerService {
	return &CustomerHandlerService{
		customer:  make(map[int32]string),
	
	}
}
func (u *CustomerHandlerService) CreateCustomer(context context.Context, request *api.CreateCustomerRequest, response *api.CreateCustomerResponse) error {
	if request.GetName() != "" {
		uid := u.getRandomUserID(maxuserid)
		if u.appendANewUser(uid, request.GetName()) {
			response.Userid =uid
			return nil
		}
	}
	return fmt.Errorf("cannot create user with empty name")
}
