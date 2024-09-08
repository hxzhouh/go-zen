package sqlite

import (
	"github.com/hxzhouh/go-zen.git/domain"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"testing"
)

type CategoryRepositoryTestSuite struct {
	suite.Suite
	DB                 *gorm.DB
	categoryRepository domain.CategoryRepository
}

func (suite *CategoryRepositoryTestSuite) SetupSuite() {
	suite.DB, _ = gorm.Open(sqlite.Open("file::memory:?cache=shared"), &gorm.Config{})
	suite.categoryRepository = NewCategoryRepository(suite.DB)
	err := suite.DB.AutoMigrate(&domain.Category{})
	assert.NoError(suite.T(), err)
}

func (suite *CategoryRepositoryTestSuite) SetupTest() {
	suite.DB.Exec("DELETE FROM categories")
}

func (suite *CategoryRepositoryTestSuite) TearDownSuite() {
	err := suite.DB.Migrator().DropTable(&domain.Category{})
	assert.NoError(suite.T(), err)
}

func TestCategoryRepositoryTestSuite(t *testing.T) {
	suite.Run(t, new(CategoryRepositoryTestSuite))
}

func (suite *CategoryRepositoryTestSuite) TestCreate() {
	category := &domain.Category{Name: "Test Category", Summary: "Test Summary"}
	err := suite.categoryRepository.Create(category)
	assert.NoError(suite.T(), err)
}

func (suite *CategoryRepositoryTestSuite) TestSearch() {
	category := &domain.Category{Name: "Test Category", Summary: "Test Summary"}
	suite.categoryRepository.Create(category)

	results, err := suite.categoryRepository.Search("Test")
	assert.NoError(suite.T(), err)
	assert.Len(suite.T(), results, 1)
}

func (suite *CategoryRepositoryTestSuite) TestUpdate() {
	category := &domain.Category{Name: "Test Category", Summary: "Test Summary"}
	suite.categoryRepository.Create(category)

	category.Name = "Updated Category"
	err := suite.categoryRepository.Update(category)
	assert.NoError(suite.T(), err)

	updatedCategory, err := suite.categoryRepository.GetByCategoryID(category.CategoryId)
	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), "Updated Category", updatedCategory.Name)
}

func (suite *CategoryRepositoryTestSuite) TestGetAll() {
	category1 := &domain.Category{Name: "Category 1", Summary: "Summary 1"}
	category2 := &domain.Category{Name: "Category 2", Summary: "Summary 2"}
	suite.categoryRepository.Create(category1)
	suite.categoryRepository.Create(category2)

	categories, err := suite.categoryRepository.GetAll()
	assert.NoError(suite.T(), err)
	assert.Len(suite.T(), categories, 2)
}

func (suite *CategoryRepositoryTestSuite) TestGetByID() {
	category := &domain.Category{Name: "Test Category", Summary: "Test Summary"}
	suite.categoryRepository.Create(category)

	fetchedCategory, err := suite.categoryRepository.GetByCategoryID(category.CategoryId)
	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), category.ID, fetchedCategory.CategoryId)
}

func (suite *CategoryRepositoryTestSuite) TestGetByIds() {
	category1 := &domain.Category{Name: "Category 1", Summary: "Summary 1"}
	category2 := &domain.Category{Name: "Category 2", Summary: "Summary 2"}
	suite.categoryRepository.Create(category1)
	suite.categoryRepository.Create(category2)

	categories, err := suite.categoryRepository.GetAll()
	assert.NoError(suite.T(), err)
	assert.Len(suite.T(), categories, 2)
}

func (suite *CategoryRepositoryTestSuite) TestDelete() {
	category := &domain.Category{Name: "Test Category", Summary: "Test Summary"}
	suite.categoryRepository.Create(category)

	err := suite.categoryRepository.Delete(category.CategoryId)
	assert.NoError(suite.T(), err)

	_, err = suite.categoryRepository.GetByCategoryID(category.CategoryId)
	assert.Error(suite.T(), err)
}
