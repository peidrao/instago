package repository

// func (u *UserRepository) FollowUser(userID, followerID uint) error {

// 	follower, err := u.FindUserByID(userID)

// 	if err != nil {
// 		return err
// 	}

// 	followed, err := u.FindUserByID(followerID)
// 	if err != nil {
// 		return err
// 	}

// 	follower.Following = append(follower.Following, followed)

// 	_, err = u.UpdateUser(follower)

// 	return err
// }

// func (u *UserRepository) UnFollowUser(userID, unfollowerID uint) error {
// 	result := u.db.Exec("DELETE FROM user_followers WHERE follower_id = ? and followed_id = ?", unfollowerID, userID)

// 	if result.Error != nil {
// 		return result.Error
// 	}

// 	return nil
// }

// func (u *UserRepository) FindFollowers(username string) ([]*entity.User, error) {
// 	var user entity.User
// 	result := u.db.Preload("Followers").Preload("Following").Where("username = ?", username).First(&user)

// 	if result.Error != nil {
// 		return nil, result.Error
// 	}

// 	return user.Followers, nil
// }

// func (u *UserRepository) FindFollowing(username string) ([]*entity.User, error) {
// 	var user entity.User
// 	result := u.db.Preload("Followers").Preload("Following").Where("username = ?", username).First(&user)

// 	if result.Error != nil {
// 		return nil, result.Error
// 	}

// 	return user.Following, nil
// }
