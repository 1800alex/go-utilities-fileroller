package fileroller

import (
	"fmt"
	"github.com/spf13/afero"
	"os"
)

// NewFileRoller creates a new FileRoller object.
func NewFileRoller(file string, maxSize int64, logsToKeep int) (obj *FileRoller, err error) {
	err = nil

	if file == "" {
		err = ErrInvalidFile
		return
	}

	obj = &FileRoller{}
	obj.FileName = file
	obj.MaxSize = maxSize
	obj.LogsToKeep = logsToKeep

	return
}

// FileRoller defines a custom FileRoller object.
type FileRoller struct {
	FileName   string
	MaxSize    int64
	LogsToKeep int
}

// Roll flips all log files if the current size + bytesToAdd is > the configured MaxSize.
func (obj *FileRoller) Roll(bytesToAdd int) (rolled bool, err error) {

	rolled = false
	err = nil

	var f afero.File
	var fi os.FileInfo

	var exists bool

	exists, err = afero.Exists(DepFS, obj.FileName)

	if false == exists {
		return
	}

	f, err = DepFS.OpenFile(obj.FileName, os.O_WRONLY, 0666)

	if err != nil {
		return
	}

	defer f.Close()

	fi, err = f.Stat()

	if err != nil {
		return
	}

	if (fi.Size() + int64(bytesToAdd)) > obj.MaxSize {
		f.Close()

		var logFrom string
		var logTo string
		var exists bool

		if obj.LogsToKeep <= 1 {
			err = DepFS.Remove(obj.FileName)
		} else {
			for i := 0; i < obj.LogsToKeep-1; i++ {
				if i >= (obj.LogsToKeep - 2) {
					logFrom = fmt.Sprintf("%s", obj.FileName)
					logTo = fmt.Sprintf("%s%d", obj.FileName, obj.LogsToKeep-i-2)
				} else {
					logFrom = fmt.Sprintf("%s%d", obj.FileName, obj.LogsToKeep-i-3)
					logTo = fmt.Sprintf("%s%d", obj.FileName, obj.LogsToKeep-i-2)
				}

				exists, err = afero.Exists(DepFS, logFrom)

				if err != nil {
					return
				}

				if true == exists {
					err = DepFS.Rename(logFrom, logTo)
				}

				if err != nil {
					return
				}
			}
		}

		rolled = true

		/* Create our file after we moved the old one */
		f, err = DepFS.OpenFile(obj.FileName, os.O_CREATE, 0666)

		if err != nil {
			return
		}

		defer f.Close()
	}

	return
}
