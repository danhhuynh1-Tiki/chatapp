package usecase

import (
	"chat/domain"
)

type talkUsecase struct {
	talkRepo domain.TalkRepository
}

func NewTalkUsecase(talkRepo domain.TalkRepository) domain.TalkUsecase {
	return &talkUsecase{talkRepo}
}
func (t *talkUsecase) AddTalk(id string, id1 string) error {
	return t.talkRepo.AddTalk(id, id1)
}
func (t *talkUsecase) GetTalk(id string, id1 string) (domain.Talk, error) {
	return t.talkRepo.GetTalk(id, id1)
}
func (t *talkUsecase) GetMessage(talkId string) []domain.Message {
	// return t.talkRepo.GetMessage(talkId)
	return nil
}
func (t *talkUsecase) AddMessage(talkId string, m domain.Message) error {
	return t.talkRepo.AddMessage(talkId, m)
}
