package service

import (
	"context"
	"database/sql"
	"github.com/go-playground/validator/v10"
	"go-simple-restful-api/helper"
	"go-simple-restful-api/model/domain"
	"go-simple-restful-api/model/web"
	"go-simple-restful-api/repository"
)

type CategoryServiceImpl struct {
	CategoryRepository repository.CategoryRepository // requires repository to manipulate data
	DB                 *sql.DB
	validate           *validator.Validate
}

// constructor function
func NewCategoryService(categoryRepository repository.CategoryRepository, DB *sql.DB, validate *validator.Validate) CategoryService {
	return &CategoryServiceImpl{
		CategoryRepository: categoryRepository,
		DB:                 DB,
		validate:           validate,
	}
}

func (service *CategoryServiceImpl) Create(ctx context.Context, request web.CategoryCreateRequest) web.CategoryResponse {
	// validation
	err := service.validate.Struct(request)
	helper.PanicIfError(err)

	tx, err := service.DB.Begin() // begin transaction
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	category := domain.Category{
		Name: request.Name,
	}

	category = service.CategoryRepository.Save(ctx, tx, category)

	return helper.ToCategoryResponse(category)
}

func (service *CategoryServiceImpl) Update(ctx context.Context, request web.CategoryUpdateRequest) web.CategoryResponse {
	err := service.validate.Struct(request)
	helper.PanicIfError(err)

	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	// find category before updating
	category, err := service.CategoryRepository.FindById(ctx, tx, request.Id)
	helper.PanicIfError(err)

	// change category name to request name
	category.Name = request.Name

	category = service.CategoryRepository.Update(ctx, tx, category)

	return helper.ToCategoryResponse(category)
}

func (service *CategoryServiceImpl) Delete(ctx context.Context, categoryId int) {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	// find category before deleting
	category, err := service.CategoryRepository.FindById(ctx, tx, categoryId)
	helper.PanicIfError(err)

	service.CategoryRepository.Delete(ctx, tx, category)
}

func (service *CategoryServiceImpl) FindById(ctx context.Context, categoryId int) web.CategoryResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	category, err := service.CategoryRepository.FindById(ctx, tx, categoryId)
	helper.PanicIfError(err)

	return helper.ToCategoryResponse(category)
}

func (service *CategoryServiceImpl) FindAll(ctx context.Context) []web.CategoryResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	categories := service.CategoryRepository.FindAll(ctx, tx)

	return helper.ToCategoryResponses(categories)
}
