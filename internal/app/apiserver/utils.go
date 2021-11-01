package apiserver

import (
	"fmt"
	"io/ioutil"
	"net/http"
)



func (s *server) ParseFiles(w http.ResponseWriter, r *http.Request) ([][]byte, error) {
	if err := r.ParseMultipartForm(16384); err != nil {
		return nil, err
	}

	res := make([][]byte, FilesCount)

	for i := 1; i <= FilesCount ; i++ {
		file, _, err := r.FormFile(fmt.Sprintf("file%v", i))
		if err != nil {
			return nil, makeMissingOrIncorrectFileErr(i)
		}
		defer file.Close()

		fileBytes, err := ioutil.ReadAll(file)
		if err != nil {
			return nil, err
		}

		res[i-1] = fileBytes
	}


	s.logger.Info("Files successfully parsed!")
	return res, nil
}
