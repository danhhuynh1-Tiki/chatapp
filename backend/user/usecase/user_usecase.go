package usecase

import "chat/domain"

type userUseCase struct {
	userRepo domain.UserRepository
}

func NewUserUsecase(userRepo domain.UserRepository) domain.UserUsecase {
	return &userUseCase{userRepo}
}

func (u *userUseCase) GetUser(user domain.User) {

}
func (u *userUseCase) Create(user domain.User) (*string, error) {
	return u.userRepo.Create(user)
}
func (u *userUseCase) GetAllUser() []domain.User {
	return u.userRepo.GetAllUser()
}

// func (u *userUseCase) IsExisted(email string) (*domain.User, error) {
// 	return u.userRepo.IsExisted(email)
// }
