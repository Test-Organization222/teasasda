package pastebin

// KeyService provides the system with the unique keys.
type KeyService interface {
	GetKey() (key string, err error)
}

//Storage stores and retrieves pastes.
type Storage interface {
	GetPaste(key string) (Paste, error)
	StorePaste(paste Paste) (Paste, error)
}
