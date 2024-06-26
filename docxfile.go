package docx

import (
	"archive/zip"
	"errors"
	"io"
	"os"
)

// DocxFile is the structure that allow to access the internal represntation
// in memory of the doc (either read or about to be written)
type DocxFile struct {
	Document    Document
	DocRelation Relationships

	rId int
}

func (d *DocxFile) GetStyleById(styleId string) *DefinedStyle {
	return d.Document.Styles.GetStyleById(styleId)
}

// New generates a new empty docx file that we can manipulate and
// later on, save
func New() *DocxFile {
	return emptyFile()
}

// Parse generates a new docx file in memory from a reader
// You can it invoke from a file
//
//	readFile, err := os.Open(FILE_PATH)
//	if err != nil {
//		panic(err)
//	}
//	fileinfo, err := readFile.Stat()
//	if err != nil {
//		panic(err)
//	}
//	size := fileinfo.Size()
//	doc, err := docx.Parse(readFile, int64(size))
//
// but also you can invoke from a webform (BEWARE of trusting users data!!!)
//
//	func uploadFile(w http.ResponseWriter, r *http.Request) {
//		r.ParseMultipartForm(10 << 20)
//
//		file, handler, err := r.FormFile("file")
//		if err != nil {
//			fmt.Println("Error Retrieving the File")
//			fmt.Println(err)
//			http.Error(w, err.Error(), http.StatusBadRequest)
//			return
//		}
//		defer file.Close()
//		docx.Parse(file, handler.Size)
//	}
func Parse(reader io.ReaderAt, size int64) (doc *DocxFile, err error) {
	zipReader, err := zip.NewReader(reader, size)
	if err != nil {
		return nil, err
	}
	doc, err = unpack(zipReader)
	return
}

// Write allows to save a docx to a writer
func (f *DocxFile) Write(writer io.Writer) (err error) {
	zipWriter := zip.NewWriter(writer)
	defer zipWriter.Close()

	return f.pack(zipWriter)
}

// References gets the url for a reference
func (f *DocxFile) References(id string) (href string, err error) {
	for _, a := range f.DocRelation.Relationships {
		if a.ID == id {
			href = a.Target
			return
		}
	}
	err = errors.New("id not found")
	return
}

func Open(fileName string) (doc *DocxFile, err error) {
	readFile, err := os.Open(fileName)
	if err != nil {
		return nil, err
	}
	fileinfo, err := readFile.Stat()
	if err != nil {
		return nil, err
	}
	size := fileinfo.Size()
	return Parse(readFile, int64(size))
}
