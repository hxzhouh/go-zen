package sqlite

import (
	"context"
	"github.com/hxzhouh/go-zen.git/domain"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"testing"
)

type UserRepositoryTestSuite struct {
	suite.Suite
	DefaultStorage *gorm.DB
	userRepository domain.UserRepository
}

func (suite *UserRepositoryTestSuite) SetupSuite() {
	suite.DefaultStorage, _ = gorm.Open(sqlite.Open("file::memory:?cache=shared"), &gorm.Config{})

	suite.userRepository = NewUserRepository(suite.DefaultStorage)
	err := suite.DefaultStorage.AutoMigrate(&domain.User{})
	assert.NoError(suite.T(), err)
}

func (suite *UserRepositoryTestSuite) SetupTest() {
	allUsers, err := suite.userRepository.Fetch(context.Background())
	assert.NoError(suite.T(), err)
	for _, user := range allUsers {
		err = suite.userRepository.DeleteByID(context.Background(), user.ID)
		assert.NoError(suite.T(), err)
	}
}
func (suite *UserRepositoryTestSuite) TestCreate() {
	user := domain.User{
		Name:     "huizhou92",
		Email:    "huizhou92",
		Password: "huizhou92",
	}
	err := suite.userRepository.Create(context.Background(), &user)
	assert.NoError(suite.T(), err)
}

func (suite *UserRepositoryTestSuite) TestFetch() {

	users, err := suite.userRepository.Fetch(context.Background())
	assert.NoError(suite.T(), err)
	assert.Len(suite.T(), users, 0)
	suite.TestCreate()
	users, err = suite.userRepository.Fetch(context.Background())
	assert.NoError(suite.T(), err)
	assert.Len(suite.T(), users, 1)
	// get By ID
	user, err := suite.userRepository.GetByID(context.Background(), users[0].ID)
	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), user.ID, users[0].ID)
	// get By Email
	user, err = suite.userRepository.GetByEmail(context.Background(), users[0].Email)
	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), user.Email, users[0].Email)
	// delete
	err = suite.userRepository.DeleteByID(context.Background(), users[0].ID)
	assert.NoError(suite.T(), err)
	users, err = suite.userRepository.Fetch(context.Background())
	assert.NoError(suite.T(), err)
	assert.Len(suite.T(), users, 0)
}

func (suite *UserRepositoryTestSuite) TearDownSuite() {
	err := suite.DefaultStorage.Migrator().DropTable(&domain.User{})
	assert.NoError(suite.T(), err)
}

func TestBackendNodeSormTestSuite(t *testing.T) {
	suite.Run(t, new(UserRepositoryTestSuite))
}
