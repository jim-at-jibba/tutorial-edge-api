package comment

import (
	"context"
	"errors"
	"fmt"
)

// By defining specific error messages for each method
// we avoid exposing sensitive info from the errors passed by
// the repositiry layer
var (
	ErrFetchingComment = errors.New("failed to fetch comment by id")
	ErrNotImplemented  = errors.New("not implemented")
)

// Comment - a representation of the Comment
// structure for our service
type Comment struct {
	ID     string
	Slug   string
	Body   string
	Author string
}

// Store - this interface defines all of the methods
// that our service needs in order to operate
// With this we keep things decoupled, we dont need to know
// the implimentation details of the data access and only that
// we need the fillowing methods to work
type Store interface {
	GetComment(context.Context, string) (Comment, error)
	PostComment(context.Context, Comment) (Comment, error)
}

// Service - is the struct that all our
// logic will be built on top of
// TIP: Accept interface and return struct
type Service struct {
	Store Store
}

// NewService - returns a pointer to a
// new service
func NewService(store Store) *Service {
	return &Service{
		Store: store,
	}
}

func (s *Service) GetComment(ctx context.Context, id string) (Comment, error) {
	fmt.Println("retreiving a comment")
	cmt, err := s.Store.GetComment(ctx, id)

	if err != nil {
		// TIP: Important to log the original error to
		// allow easy debugging
		fmt.Println(err)
		return Comment{}, ErrFetchingComment
	}
	return cmt, nil
}

func (s *Service) UpdateComment(ctx context.Context, cmt Comment) error {
	return ErrNotImplemented
}

func (s *Service) DeleteComment(ctx context.Context, id string) error {
	return ErrNotImplemented
}

// context.Content - allows us to pass important information through our services
// request and trace ids -  corrolation ids
func (s *Service) PostComment(ctx context.Context, cmt Comment) (Comment, error) {
	insertedCmt, err := s.Store.PostComment(ctx, cmt)
	if err != nil {
		return Comment{}, err
	}
	return insertedCmt, nil
}
