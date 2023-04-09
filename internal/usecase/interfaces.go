package usecase

import "context"

type Rester interface {
	FindSubStr(context.Context, string) (string, error)
	CheckEmail(context.Context, string) ([]string, error)
	FindSelf(context.Context, string) ([]string, error)
}