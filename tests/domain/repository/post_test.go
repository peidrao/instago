package tests

import (
	"testing"

	"github.com/peidrao/instago/internal/domain/entity"
	"github.com/stretchr/testify/assert"
)

func TestPostCreateRepository(t *testing.T) {
	setupTest()
	defer tearDownTest()

	user := createUser(t, "test", "test@test.com", "Test")

	post := createPost(t, user.ID)

	var firsPost entity.Post
	err := db.First(&firsPost, post.ID).Error
	assert.NoError(t, err, "Erro ao buscar postagem do banco de dados")

	assert.Equal(t, post.Caption, firsPost.Caption, "Legenda não coincidem")
	assert.Equal(t, post.Location, firsPost.Location, "Localização não cincidem")
	assert.Equal(t, post.ImageURL, firsPost.ImageURL, "Imagem não coincidem")
	assert.Equal(t, post.UserID, firsPost.UserID, "Usuários não coincidem")
}

func TestFindPostsByUserRepository(t *testing.T) {
	setupTest()
	defer tearDownTest()

	user := createUser(t, "test", "test@test.com", "Test")

	for i := 0; i < 5; i++ {
		createPost(t, user.ID)
	}

	postsByUser, err := postRepo.FindPostsByUser(user.ID)
	assert.NoError(t, err, "Erro ao buscar postagens no banco de dados")
	assert.Len(t, postsByUser, 5, "Quantidade de postagens não coincidem")
}

func TestFindPostByIDRepository(t *testing.T) {
	setupTest()
	defer tearDownTest()

	user := createUser(t, "test", "test@test.com", "Test")

	post := createPost(t, user.ID)

	findPost, err := postRepo.FindPostByID(post.ID)
	assert.NoError(t, err, "Erro ao buscar postagem pelo ID")

	assert.Equal(t, post.Caption, findPost.Caption, "Legenda não coincidem")
	assert.Equal(t, post.Location, findPost.Location, "Localização não cincidem")
	assert.Equal(t, post.ImageURL, findPost.ImageURL, "Imagem não coincidem")
	assert.Equal(t, post.UserID, findPost.UserID, "Usuários não coincidem")
}

func TestRemovePostRepository(t *testing.T) {
	setupTest()
	defer tearDownTest()

	user := createUser(t, "test", "test@test.com", "Test")

	post := createPost(t, user.ID)

	err := postRepo.RemovePost(post.ID)
	assert.NoError(t, err, "Erro ao remover postagem")

	var postRemoved entity.Post
	err = db.First(&postRemoved, post.ID).Error
	assert.Contains(t, err.Error(), "record not found", "Existe uma postagem no banco de dados")
}
