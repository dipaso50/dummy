package fileGenerator

import (
	"io/ioutil"
	"log"
	"math/rand"
	"path/filepath"
	"sync"
	"time"

	"github.com/google/uuid"
)

var (
	bigBuff []byte
	ext     = [...]string{"pdf", "jpg", "doc", "txt", "xls"}
	r       = rand.New(rand.NewSource(time.Now().UnixNano()))
)

func Generate(fileNumber int, destinationFolder string, mgNumber int, fileExtension string) {

	var wg sync.WaitGroup

	bigBuff = make([]byte, mgNumber*1000000)

	for i := 0; i < fileNumber; i++ {
		wg.Add(1)

		go func(folder string, fileExt string) {
			defer wg.Done()

			aFile, err := createFileName(destinationFolder, fileExt)

			if err != nil {
				log.Printf("Error creating filename %v \n", err)
			}

			if err = generateFile(aFile); err != nil {
				log.Printf("Error creating file %v \n", err)
			}
		}(destinationFolder, fileExtension)
	}

	wg.Wait()
}

func createFileName(destinationFolder, fileExtension string) (string, error) {
	var fileName uuid.UUID
	var err error

	if fileName, err = uuid.NewUUID(); err != nil {
		return "", err
	}

	extension := ext[r.Intn(len(ext))]

	if len(fileExtension) > 0 {
		extension = fileExtension
	}

	finalFile := filepath.Join(destinationFolder, fileName.String()+"."+extension)

	return finalFile, nil
}

func generateFile(dummy string) error {

	err := ioutil.WriteFile(dummy, bigBuff, 0666)

	if err != nil {
		return err
	}

	log.Printf("File %s created!! \n", dummy)
	return nil
}
