package obgsstorage

import "time"

// Storage ...
type Storage interface {
	// User
	UserAdd(u User) error
	UserRemove(username string) error
	// Friendship
	UserRequestFriendship(usernameFrom, usernameTo string) error
	UserRemoveFriendship(usernameFrom, usernameTo string) error
	// Game
	// GameAdd(name string) error
}

// User is a user description
type User struct {
	ID          int       `json:"id"`
	Username    string    `json:"username"`
	Displayname string    `json:"displayname"`
	BirthDate   time.Time `json:"birthdate"`
	Gender      string    `json:"gender"`
}
