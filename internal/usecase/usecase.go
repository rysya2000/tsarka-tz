package usecase

import "context"

type UseCase struct {
}

func New() *UseCase {
	return &UseCase{}
}

func (u *UseCase) findSubStr(context.Context, string) (string, error) {
	return "", nil
}

func (u *UseCase) checkEmail(context.Context, string) ([]string, error) {
	return []string{}, nil
}

func (u *UseCase) findSelf(context.Context) ([]string, error) {
	return []string{}, nil
}
