package file

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"os"
	"path/filepath"
)

// Copy 文件拷贝
func Copy(source, target string) error {
	if !IsExist(source) || IsExist(target) {
		return errors.New(fmt.Sprintf("source: %s is not exist or target: %s is exist", source, target))
	}
	tDir := filepath.Dir(target)
	if !IsExist(tDir) {
		err := os.MkdirAll(tDir, os.ModePerm)
		if err != nil {
			return err
		}
	}
	sf, err := os.Open(source)
	if err != nil {
		return err
	}
	defer sf.Close()
	tf, err := os.Create(target)
	if err != nil {
		return err
	}
	defer tf.Close()
	_, err = io.Copy(bufio.NewWriter(tf), bufio.NewReader(sf))
	if err != nil {
		return err
	}
	return nil
}

// ByBytesCreate 通过字节数组创建文件
func ByBytesCreate(data []byte, target string) error {
	if IsExist(target) {
		return errors.New(fmt.Sprintf("target: %s is exist", target))
	}
	tf, err := os.Create(target)
	if err != nil {
		return err
	}
	defer tf.Close()
	wf := bufio.NewWriter(tf)
	_, err = wf.Write(data)
	if err != nil {
		return err
	}
	err = wf.Flush()
	if err != nil {
		return err
	}
	return nil
}

// IsExist 判断文件是否存在
func IsExist(path string) bool {
	_, err := os.Stat(path)
	if err == nil || os.IsExist(err) {
		return true
	}
	return false
}
