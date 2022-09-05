package file_subscriber

import (
	"io/fs"
	"log"
	"os"
	"syscall"
)

type FileSubsciber struct {
	dbFile string

}

func NewFileSubsciber() *FileSubsciber {
	return &FileSubsciber{
		dbFile: dbFile(),
	}

}

func dbFile() string {
	return os.TempDir() + "/fileWatcher.txt";
}

func (subsciber *FileSubsciber) FileSubscribe(filePath string) {
	file, err := os.OpenFile(filePath, syscall.O_RDONLY, fs.ModeSticky)
	if err != nil {
		log.Fatalln(err)
	}

	stat, err := file.Stat()
	if err != nil {
		return
	}
	stat.ModTime()
}

func (subsciber *FileSubsciber) readFile(filePath string) {

}
