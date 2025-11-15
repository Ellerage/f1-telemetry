package filemanager

type BufferFileMangerParams struct {
	FileManager *FileManager
	BufferSize  int
}

type BufferFileManager struct {
	fileManager *FileManager
	buffer      [][]string
	bufferSize  int
}

func NewBufferFileManger(params BufferFileMangerParams) *BufferFileManager {
	return &BufferFileManager{fileManager: params.FileManager, bufferSize: params.BufferSize}
}

func (fm *BufferFileManager) AddRow(values []string) error {
	fm.buffer = append(fm.buffer, values)

	if len(fm.buffer) >= fm.bufferSize {
		return fm.commitWrite()
	}

	return nil
}

func (fm *BufferFileManager) clearBuffer() {
	fm.buffer = fm.buffer[:0]
}

func (fm *BufferFileManager) commitWrite() error {
	if err := fm.fileManager.WriteRows(fm.buffer); err != nil {
		return err
	}

	fm.clearBuffer()

	return nil
}
