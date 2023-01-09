package file

//
//import (
//	"fmt"
//	"github.com/fsnotify/fsnotify"
//)
//
//type Notify struct {
//	dir string
//	notifyOption
//}
//
//type notifyOption struct {
//	create Create
//	write  Write
//	rename Rename
//	remove Remove
//}
//
//func WithCreate(create Create) Option {
//	return func(option *notifyOption) {
//		if create == nil {
//			return
//		}
//		option.create = create
//	}
//}
//
//func WithWrite(write Write) Option {
//	return func(option *notifyOption) {
//		if write == nil {
//			return
//		}
//		option.write = write
//	}
//}
//
//func WithRename(rename Rename) Option {
//	return func(option *notifyOption) {
//		if rename == nil {
//			return
//		}
//		option.rename = rename
//	}
//}
//
//func WithRemove(remove Remove) Option {
//	return func(option *notifyOption) {
//		if remove == nil {
//			return
//		}
//		option.remove = remove
//	}
//}
//
//type Option func(*notifyOption)
//
//func NewListener(dir string, opts ...Option) Listener {
//	defaultOption := newDefaultOption()
//	for _, apply := range opts {
//		apply(&defaultOption)
//	}
//	n := Notify{
//		dir:          dir,
//		notifyOption: defaultOption,
//	}
//	return &n
//}
//
//func newDefaultOption() notifyOption {
//
//	return notifyOption{
//		create: &filenameFmt{},
//		write:  &filenameFmt{},
//		rename: &filenameFmt{},
//		remove: &filenameFmt{},
//	}
//}
//
//type filenameFmt struct {
//}
//
//func (ff *filenameFmt) Create(filename string) {
//	fmt.Printf("Create: %s\n", filename)
//}
//
//func (ff *filenameFmt) Write(filename string) {
//	fmt.Printf("Write: %s\n", filename)
//}
//
//func (ff *filenameFmt) Rename(filename string) {
//	fmt.Printf("Rename: %s\n", filename)
//}
//
//func (ff *filenameFmt) Remove(filename string) {
//	fmt.Printf("Remove: %s\n", filename)
//}
//
//func (n *Notify) Listen(done chan struct{}) {
//	watcher, err := fsnotify.NewWatcher()
//	if err != nil {
//		fmt.Println(err)
//		return
//	}
//	defer watcher.Close()
//	go func() {
//		for {
//			select {
//			case event, ok := <-watcher.Events:
//				if !ok {
//					return
//				}
//				if event.Op&fsnotify.Create == fsnotify.Create && n.create != nil {
//					n.create.Create(event.Name)
//				}
//				if event.Op&fsnotify.Write == fsnotify.Write && n.write != nil {
//					n.write.Write(event.Name)
//				}
//				if event.Op&fsnotify.Rename == fsnotify.Rename && n.rename != nil {
//					n.rename.Rename(event.Name)
//				}
//				if event.Op&fsnotify.Remove == fsnotify.Remove && n.remove != nil {
//					n.remove.Remove(event.Name)
//				}
//			case err, ok := <-watcher.Errors:
//				if !ok {
//					return
//				}
//				fmt.Println(err)
//			}
//		}
//	}()
//	err = watcher.Add(n.dir)
//	if err != nil {
//		fmt.Println(err)
//		return
//	}
//	<-done
//}
