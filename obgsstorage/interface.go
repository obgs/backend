package obgsstorage

// Storage ...
type Storage interface {
	// User
	UserAdd(name, birth, gender string) error
	UserRemove(name string) error
	// Friendship
	UserRequestFriendship(nameFrom, nameTo string) error
	UserRemoveFriendship(nameFrom, nameTo string) error
	// Game
	// GameAdd(name string) error
}
