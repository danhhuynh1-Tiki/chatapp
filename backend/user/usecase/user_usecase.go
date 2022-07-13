package usecase

import "chat/domain"

type userUseCase struct {
	userRepo domain.UserRepository
}

func NewUserUseCase(userRepo domain.UserRepository) domain.UserUsecase {
	return &userUseCase{userRepo}
}

func (u *userUseCase) GetUser(user domain.User) {

}
func (u *userUseCase) Create(user domain.User) (*string, error) {
	return u.userRepo.Create(user)
}

// func (u *userUseCase) IsExisted(email string) (*domain.User, error) {
// 	return u.userRepo.IsExisted(email)
// }
