package file

type Create interface {
	Create(filename string)
}

type Write interface {
	Write(filename string)
}

type Rename interface {
	Rename(filename string)
}

type Remove interface {
	Remove(filename string)
}
