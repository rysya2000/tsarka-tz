package usecase

import "context"

type UseCase struct {
	rest Rester
}

func New() *UseCase {
	return &UseCase{}
}

func (u *UseCase) FindSubStr(ctx context.Context, s string) (string, error) {
	return Find_SubStr(s), nil
}

func (u *UseCase) CheckEmail(ctx context.Context, s string) ([]string, error) {
	return Check_Email(s), nil
}

func (u *UseCase) FindSelf(ctx context.Context, s string) ([]string, error) {
	return Find_Self(s)
}
