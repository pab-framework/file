package file

type Listener interface {
	Listen(done chan struct{})
}
