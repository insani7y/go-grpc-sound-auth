package apiserver

import (
	"fmt"
	"github.com/reqww/go-rest-api/internal/app/model"
	"io/ioutil"
	"net/http"
)

func (s *server) ParseFiles(w http.ResponseWriter, r *http.Request, userId int) error {
	if err := r.ParseMultipartForm(1024); err != nil {
		return err
	}

	file, handler, err := r.FormFile("file")
	if err != nil {
		return err
	}
	defer file.Close()


	fileBytes, err := ioutil.ReadAll(file)
	if err != nil {
		return err
	}

	if err := ioutil.WriteFile(fmt.Sprintf("%v/%v/%v", model.SoundDir, userId, handler.Filename), fileBytes, 0644); err != nil {
		return err
	}

	s.logger.Info("Files successfully parsed!")
	return nil
}
