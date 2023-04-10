package services

import (
	"restful_api_testing/repositories"
	"restful_api_testing/models"
)

type BlogService struct {
	repository repositories.BlogRepository
}

func InitBlogService() BlogService {
	return BlogService{
		repository: &repositories.BlogRepositoryImp{},
	}
}

func (bs *BlogService) GetAll() ([]models.Blog, error){
	return bs.repository.GetAll()
}

func (bs *BlogService) GetById(id string) (models.Blog, error){
	return bs.repository.GetById(id)
}

func (bs *BlogService) Create(blogInput models.BlogInput) (models.Blog, error){
	return bs.repository.Create(blogInput)
}

func (bs *BlogService) Update(blogInput models.BlogInput, id string) (models.Blog, error){
	return bs.repository.Update(blogInput, id)
}

func (bs *BlogService) Delete(id string) (error){
	return bs.repository.Delete(id)
}



