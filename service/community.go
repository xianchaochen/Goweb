package service

import (
	"bluebell/repository"
	"bluebell/entity"
)

type ICommunityService interface {
	SelectCommunityList() (list []*entity.Community, err error)
}

var _ ICommunityService = (*CommunityService)(nil)

type CommunityService struct {
	communityRepository repository.ICommunityRepository
}

func NewCommunityService(repository repository.ICommunityRepository) ICommunityService {
	return &CommunityService{communityRepository: repository}
}

func (c *CommunityService) SelectCommunityList() (list []*entity.Community, err error) {
	list, err = c.communityRepository.SelectCommunityList()
	if err != nil {
		return nil, err
	}
	return list, nil
}

