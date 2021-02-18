package pastebin

import "errors"

// KeyService provides the system with the unique keys.
type KeyService interface {
	GetKey() (key string, err error)
}

//Storage stores and retrieves pastes.
type Storage interface {
	GetPaste(key string) (Paste, error)
	StorePaste(paste Paste) error
}

// ErrKeyExists happens when #StorePaste is triggered with the already-existing Key
var ErrKeyExists error = errors.New("Key already exists")
