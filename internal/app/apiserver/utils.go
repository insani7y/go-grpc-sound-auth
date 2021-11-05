package apiserver

import (
	"fmt"
	"mime/multipart"
	"net/http"
)



func (s *server) ParseFiles(w http.ResponseWriter, r *http.Request, numFiles int) ([]multipart.File, error) {
	if err := r.ParseMultipartForm(16384); err != nil {
		return nil, err
	}

	var res []multipart.File

	for i := 1; i <= numFiles ; i++ {
		file, _, err := r.FormFile(fmt.Sprintf("file%v", i))
		if err != nil {
			return nil, makeMissingOrIncorrectFileErr(i)
		}

		res = append(res, file)
	}


	s.logger.Info("Files successfully parsed!")
	return res, nil
}
