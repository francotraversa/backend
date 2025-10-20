package utils

import "github.com/francotraversa/siriusbackend/internal/types"

func FindUserByUserId(userId string) *types.User {
	db := DatabaseInstance{}.Instance()

	user := types.User{}
	result := db.Limit(1).Find(&user, "user_id = ?", userId)
	if result.Error == nil && result.RowsAffected > 0 {
		return &user
	}
	return nil
}

func FindUserByUsername(username string) *types.User {
	db := DatabaseInstance{}.Instance()

	user := types.User{}
	result := db.Limit(1).Find(&user, "username = ?", username)
	if result.Error == nil && result.RowsAffected > 0 {
		return &user
	}
	return nil
}
func FindUserByEmail(email string) *types.User {
	db := DatabaseInstance{}.Instance()

	user := types.User{}
	result := db.Limit(1).Find(&user, "email = ?", email)
	if result.Error == nil && result.RowsAffected > 0 {
		return &user
	}
	return nil
}
