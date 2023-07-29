package tests

import (
	"testing"

	"github.com/peidrao/instago/internal/domain/entity"
	"github.com/stretchr/testify/assert"
)

func TestFollowCreateRepository(t *testing.T) {
	setupTest()
	defer tearDownTest()

	user := createUser(t, "test", "test@test.com", "Test")
	user2 := createUser(t, "test1", "test1@test1.com", "Test One")

	follow := createFollow(t, user.ID, user2.ID)

	var firstFollow entity.Follow
	err := db.First(&firstFollow, follow.ID).Error
	assert.NoError(t, err, "Erro ao buscar postagem do banco de dados")

	assert.Equal(t, follow.FollowerID, firstFollow.FollowerID, "Legenda não coincidem")
	assert.Equal(t, follow.FollowingID, firstFollow.FollowingID, "Localização não cincidem")
}

func TestFollowDeleteRepository(t *testing.T) {
	setupTest()
	defer tearDownTest()

	user := createUser(t, "test", "test@test.com", "Test")
	user2 := createUser(t, "test1", "test1@test1.com", "Test One")

	follow := createFollow(t, user.ID, user2.ID)

	err := followRepo.DeleteFollow(follow, follow.ID)
	assert.NoError(t, err, "Erro ao remover seguidor")

}
