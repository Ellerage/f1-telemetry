package filemanager

import (
	"encoding/csv"
	"errors"
	"f1-telemetry/internal/model"
	"os"
)

type FileManager struct {
	filePath string
	file     *os.File
	writer   *csv.Writer
}

func NewFileManager() *FileManager {
	return &FileManager{}
}

func (fm *FileManager) OpenFile(filePath string) (bool, error) {
	if fm.file != nil {
		fm.CloseFile()
	}

	isFileExist := fm.IsFileExist()
	var errs []error

	file, err := os.OpenFile(fm.filePath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		errs = append(errs, err)
	}

	fm.file = file
	fm.filePath = filePath

	writer := csv.NewWriter(file)
	fm.writer = writer

	// Write header - first row
	if !isFileExist {
		err = fm.writeRow(model.TelemetryRowColumns)
		errs = append(errs, err)
	}

	return isFileExist, errors.Join(errs...)
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
	_, err := os.Stat(fm.filePath)
	return !os.IsNotExist(err)
}

// Write one row - use it for header
func (fm *FileManager) writeRow(values []string) error {
	defer fm.writer.Flush()
	return fm.writer.Write(values)
}
