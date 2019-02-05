package main

import (
	"encoding/json"
	"net/http"
)

type RPCStatus string

const (
	RPCStatusSuccess = RPCStatus("success")
	RPCStatusFail    = RPCStatus("fail")
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
	Status     RPCStatus
	Body       string
	ErrMessage string
}

func prepareErrMsg(err error) *Response {
	return &Response{
		Status: RPCStatusFail,
		ErrMessage: err.Error(),
	}
}

func NewUserService(repo IUserRepository) *UserService{
	service := new(UserService)
	service.Repo = repo
	return service
}

func (s *UserService) GetByLogin(r *http.Request, args *UserArgs, result *Response) error {
	user, err := s.Repo.GetByLogin(args.Login)
	if err != nil {
		result = prepareErrMsg(err)
		return err
	}
	item, err := json.Marshal(user)
	if err != nil {
		result = prepareErrMsg(err)
		return err
	}
	result.Status = RPCStatusSuccess
	result.Body = string(item)
	return nil
}

func (s *UserService) Create(r *http.Request, args *UserArgs, result *Response) error {
	userID, err := s.Repo.Create(args.Login)
	if err != nil {
		result = prepareErrMsg(err)
		return err
	}
	result.Status = RPCStatusSuccess
	result.Body = userID
	return nil
}

func (s *UserService) Update(r *http.Request, args *UserUpdateArgs, result *Response) error {
	err := s.Repo.Update(args.Login, args.NewLogin)
	if err != nil {
		result = prepareErrMsg(err)
		return err
	}
	result.Status = RPCStatusSuccess
	return nil
}
