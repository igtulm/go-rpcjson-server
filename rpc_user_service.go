package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type UserService struct{
	Repo IUserRepository
}

type UserArgs struct {
	Login string
}

type UserUpdateArgs struct {
	UserArgs
	NewLogin string
}

type Response struct {
	Result string
}

// TODO make json preparation smarter
func jsonErrMsg(err error) string {
	return fmt.Sprintf("{\"status\": \"fail\", \"message\": \"%v\"}", err)
}

func jsonStatusOk() string {
	return fmt.Sprintf("{\"status\": \"ok\"}")
}

func jsonFieldWithValue(field string, value string) string {
	return fmt.Sprintf("{\"%s\": \"%s\"}", field, value)
}

func NewUserService(repo IUserRepository) *UserService{
	service := new(UserService)
	service.Repo = repo
	return service
}

func (s *UserService) GetByLogin(r *http.Request, args *UserArgs, result *Response) error {
	user, err := s.Repo.GetByLogin(args.Login)
	if err != nil {
		result.Result = jsonErrMsg(err)
		return err
	}
	item, err := json.Marshal(user)
	if err != nil {
		result.Result = jsonErrMsg(err)
		return err
	}
	result.Result = string(item)
	return nil
}

func (s *UserService) Create(r *http.Request, args *UserArgs, result *Response) error {
	userID, err := s.Repo.Create(args.Login)
	if err != nil {
		result.Result = jsonErrMsg(err)
		return err
	}
	result.Result = jsonFieldWithValue("id", userID)
	return nil
}

func (s *UserService) Update(r *http.Request, args *UserUpdateArgs, result *Response) error {
	err := s.Repo.Update(args.Login, args.NewLogin)
	if err != nil {
		result.Result = jsonErrMsg(err)
		return err
	}
	result.Result = jsonStatusOk()
	return nil
}
