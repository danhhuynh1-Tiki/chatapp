package domain

type SingupUsecase interface {
	Create(user User) (*string, error)
}
