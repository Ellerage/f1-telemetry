package filemanager

import (
	"encoding/csv"
	model "f1-telemetry/internal/model/csv"
	"os"
	"path/filepath"
)

type FileManager struct {
	filePath string
	file     *os.File
	writer   *csv.Writer
}

func NewFileManager() *FileManager {
	return &FileManager{}
}

func (fm *FileManager) OpenFile(path string) (bool, error) {
	fm.filePath = path

	if fm.file != nil {
		fm.CloseFile()
	}

	isFileExist := fm.IsFileExist()

	dir := filepath.Dir(path)
	if err := os.MkdirAll(dir, 0755); err != nil {
		return false, err
	}

	file, err := os.OpenFile(path, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		return isFileExist, err
	}
	fm.file = file

	writer := csv.NewWriter(file)
	fm.writer = writer

	// Write header - first row
	if !isFileExist {
		err = fm.WriteRow(model.TelemetryRowColumns)
		return isFileExist, err
	}

	return isFileExist, nil
}

func (fm *FileManager) CloseFile() {
	fm.writer.Flush()
	fm.file.Close()
}

func (fm *FileManager) WriteRows(values [][]string) error {
	defer fm.writer.Flush()
	return fm.writer.WriteAll(values)
}

func (fm *FileManager) IsFileExist() bool {
	if fm.filePath == "" {
		return false
	}
	_, err := os.Stat(fm.filePath)
	if err != nil {
		if os.IsNotExist(err) {
			return false
		}

		return false
	}
	return true
}

// Write one row - not effective for frequent writing
func (fm *FileManager) WriteRow(values []string) error {
	defer fm.writer.Flush()
	return fm.writer.Write(values)
}
