package pastebin

//App represents the business logic
type App struct {
}

//Store stores a Paste object.
func (a App) Store(p Paste) (Paste, error) {
	return Paste{}, nil
}

//Get retrieves the paste by given key.
func (a App) Get(k string) (Paste, error) {
	return Paste{}, nil
}
