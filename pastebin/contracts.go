package pastebin

type KeyGenerator interface {
	// GetKey() (key string, err error)
}

type ObjectStorage interface {
	// GetPaste(key string) (key Paste, err error)
	// StorePaste(paste Paste) (key string, err error)
}

type MetadataStorage interface {
	// GetPasteBody(objId string) (pasteBody string, err error)
	// StorePasteBody(pasteBody string) (objId string, err error)
}
