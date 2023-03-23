package service

import (
	"github.com/rokinanda/golang-unit-test.git/entity"
	"github.com/rokinanda/golang-unit-test.git/repository"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

var categoryRepository = &repository.CategoryRepositoryMock{Mock: mock.Mock{}}
var categoryService = CategoryService{Repository: categoryRepository}

func TestCategoryService_GetNotFound(t *testing.T) {
	//program mock
	categoryRepository.Mock.On("FindById", "1").Return(nil)

	category, err := categoryService.Get("1")
	assert.Nil(t, category)
	assert.NotNil(t, err)
}

func TestCategoryService_GetSuccess(t *testing.T) {
	category := entity.Category{
		Id:   "1000",
		Name: "Laptop",
	}
	categoryRepository.Mock.On("FindById", "1").Return(category)
	result, err := categoryService.Get("1")
	assert.Nil(t, err)
	assert.NotNil(t, result)
	//fmt.Println(category.Id)
	//fmt.Println(result)
	assert.Equal(t, category.Id, result.Id)
	assert.Equal(t, category.Name, result.Name)

}
