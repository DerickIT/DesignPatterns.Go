package main

import (
	"fmt"
	"time"
)

type UserService interface {
	GetUserData(userID int) string
}

type RealUserService struct{}

func (s *RealUserService) GetUserData(userID int) string {
	time.Sleep(2 * time.Second)
	fmt.Println("Fetching user data from database")
	return fmt.Sprintf("User data for user ID %d", userID)
}

type CachedUserService struct {
	realService UserService
	cache       map[int]string
}

func NewCachedUserService(realService UserService) *CachedUserService {
	return &CachedUserService{
		realService: realService,
		cache:       make(map[int]string),
	}
}

func (s *CachedUserService) GetUserData(userID int) string {
	if data, found := s.cache[userID]; found {
		fmt.Println("Fetching user data from cache")
		return data
	}

	data := s.realService.GetUserData(userID)
	s.cache[userID] = data
	return data
}

func main() {
	realService := &RealUserService{}
	cachedService := NewCachedUserService(realService)

	fmt.Println(cachedService.GetUserData(1))
	fmt.Println(cachedService.GetUserData(1))
}
