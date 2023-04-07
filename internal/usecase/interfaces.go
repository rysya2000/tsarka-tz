package usecase

import "context"

type Rester interface {
	findSubStr(context.Context, string) (string, error)
	checkEmail(context.Context, string) ([]string, error)
	findSelf(context.Context) ([]string, error)
}