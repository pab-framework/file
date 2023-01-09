package file

import "testing"

func TestName(t *testing.T) {

}

//
//import (
//	"os"
//	"testing"
//	"time"
//)
//
//func TestNotify_Listen(t *testing.T) {
//	type fields struct {
//		dir          string
//		notifyOption notifyOption
//	}
//	type args struct {
//		done chan struct{}
//	}
//	tests := []struct {
//		name   string
//		fields fields
//		args   args
//	}{
//		{
//			name: "test_Notify_Listen_1",
//			args: args{done: make(chan struct{})},
//			fields: fields{
//				dir:          "C:\\Users\\leig\\Developer\\github\\file\\20230107\\file",
//				notifyOption: newDefaultOption(),
//			},
//		},
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			n := &Notify{
//				dir:          tt.fields.dir,
//				notifyOption: tt.fields.notifyOption,
//			}
//			go func() {
//				time.Sleep(10 * time.Second)
//				tt.args.done <- struct{}{}
//			}()
//			go func() {
//				time.Sleep(3 * time.Second)
//				f, err := os.Create("C:\\Users\\leig\\Developer\\github\\file\\20230107\\file\\test.txt")
//				if err != nil {
//					t.Error(err)
//					tt.args.done <- struct{}{}
//				}
//				defer f.Close()
//			}()
//			go func() {
//				time.Sleep(6 * time.Second)
//				err := os.Remove("C:\\Users\\leig\\Developer\\github\\file\\20230107\\file\\test.txt")
//				if err != nil {
//					t.Error(err)
//					tt.args.done <- struct{}{}
//				}
//			}()
//			n.Listen(tt.args.done)
//		})
//	}
//}
