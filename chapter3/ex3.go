package main

type Creator interface {
	Create() error
}

type Updater interface {
	Update() error
}

type Deleter interface {
	Delete() error
}

type Finder interface {
	Find() (string, error)
}
