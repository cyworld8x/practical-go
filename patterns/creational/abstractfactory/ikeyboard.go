package abstractfactory

// IKeyboard interface
type IKeyboard interface {
	typing() string
	press(key string)
}

// Keyboard struct
type Keyboard struct {
	keys string
}

func (k *Keyboard) typing() string {
	return "Typing  '" + k.keys + "' on keyboard"
}

func (k *Keyboard) press(key string) {
	k.keys = key
}
