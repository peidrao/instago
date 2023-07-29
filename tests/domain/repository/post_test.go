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
