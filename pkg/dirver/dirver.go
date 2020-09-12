package dirver

import (
	"errors"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
	"github.com/goftp/server"
	"github.com/lunny/log"
)

type PutFileOnlyDriver struct {
	RootPath string
	server.Perm
}

type FileInfo struct {
	os.FileInfo

	mode  os.FileMode
	owner string
	group string
}

func (f *FileInfo) Mode() os.FileMode {
	return f.mode
}

func (f *FileInfo) Owner() string {
	return f.owner
}

func (f *FileInfo) Group() string {
	return f.group
}

func (driver *PutFileOnlyDriver) realPath(path string) string {
	paths := strings.Split(path, "/")
	destPath := filepath.Join(append([]string{driver.RootPath}, paths[:len(paths) - 1]...)...)
	//os.MkdirAll(filepath.Join(append([]string{driver.RootPath}, paths2...)...), os.ModePerm)
	//fmt.Println(destPath)
	_, err := os.Lstat(destPath)
	if os.IsNotExist(err) {
		log.Info("############### MkdirPath : ", destPath)
		os.MkdirAll(destPath, os.ModePerm)
	}
	return filepath.Join(append([]string{driver.RootPath}, paths...)...)
}

func (driver *PutFileOnlyDriver) Init(conn *server.Conn) {
	//driver.conn = conn
}

func (driver *PutFileOnlyDriver) ChangeDir(path string) error {
	rPath := driver.realPath(path)
	f, err := os.Lstat(rPath)
	if err != nil {
		return err
	}
	if f.IsDir() {
		return nil
	}
	return errors.New("Not a directory")
}

func (driver *PutFileOnlyDriver) Stat(path string) (server.FileInfo, error) {
	basepath := driver.realPath(path)
	rPath, err := filepath.Abs(basepath)
	if err != nil {
		return nil, err
	}
	f, err := os.Lstat(rPath)
	if err != nil {
		return nil, err
	}
	mode, err := driver.Perm.GetMode(path)
	if err != nil {
		return nil, err
	}
	if f.IsDir() {
		mode |= os.ModeDir
	}
	owner, err := driver.Perm.GetOwner(path)
	if err != nil {
		return nil, err
	}
	group, err := driver.Perm.GetGroup(path)
	if err != nil {
		return nil, err
	}
	return &FileInfo{f, mode, owner, group}, nil
}

func (driver *PutFileOnlyDriver) ListDir(path string, callback func(server.FileInfo) error) error {
	return nil
}

func (driver *PutFileOnlyDriver) DeleteDir(path string) error {
	return errors.New("Not a directory")
}

func (driver *PutFileOnlyDriver) DeleteFile(path string) error {
	return errors.New("Not a file")
}

func (driver *PutFileOnlyDriver) Rename(fromPath string, toPath string) error {
	return errors.New("Not a file")
}

func (driver *PutFileOnlyDriver) MakeDir(path string) error {
	return errors.New("Not a directory")
}

func (driver *PutFileOnlyDriver) GetFile(path string, offset int64) (int64, io.ReadCloser, error) {
	return 0, nil, errors.New("Not a file")
}

func (driver *PutFileOnlyDriver) PutFile(destPath string, data io.Reader, appendData bool) (int64, error) {
	log.Info("############### PutFile destPath : ", destPath)
	rPath := driver.realPath(destPath)
	var isExist bool
	//fmt.Println(rPath)
	f, err := os.Lstat(rPath)
	if err == nil {
		isExist = true
		if f.IsDir() {
			return 0, errors.New("A dir has the same name")
		}
	} else {
		fmt.Println(err)
		if os.IsNotExist(err) {
			isExist = false
		} else {
			return 0, errors.New(fmt.Sprintln("Put File error:", err))
		}
	}

	if appendData && !isExist {
		appendData = false
	}

	if !appendData {
		if isExist {
			err = os.Remove(rPath)
			if err != nil {
				return 0, err
			}
		}
		f, err := os.Create(rPath)
		if err != nil {
			return 0, err
		}
		defer f.Close()
		bytes, err := io.Copy(f, data)
		if err != nil {
			return 0, err
		}
		return bytes, nil
	}

	of, err := os.OpenFile(rPath, os.O_APPEND|os.O_RDWR, 0660)
	if err != nil {
		return 0, err
	}
	defer of.Close()

	_, err = of.Seek(0, os.SEEK_END)
	if err != nil {
		return 0, err
	}

	bytes, err := io.Copy(of, data)
	if err != nil {
		return 0, err
	}

	return bytes, nil
}

type PutFileOnlyDriverFactory struct {
	RootPath string
	server.Perm
}

func (factory *PutFileOnlyDriverFactory) NewDriver() (server.Driver, error) {
	return &PutFileOnlyDriver{factory.RootPath, factory.Perm}, nil
}