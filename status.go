package statusAPI

import (
	"github.com/labstack/echo"
	"io"
	"io/ioutil"
	"json"
	"net/http"
	"os"
)

func StatusPage(c echo.Context) error {
	return c.String(http.StatusOK, "こんにちはdesune！")
}

type Status struct {
	Name              string `json:"Name"`
	ID                string `json:"ID"`
	Session           bool   `json:"Session"`
	Login             bool   `json:"Login"`
	LastRequestStatus bool   `json:"LastRequestStatus"`
}

type StatusJSON []Status

func NewStatusJSON() StatusJSON {
	return StatusJSON{}
}

func (s *StatusJSON) Read(r io.Reader) error {
	b, err := ioutil.ReadAll(r)
	if err != nil {
		return err
	}
	return json.Unmarshal(b, &s)
}

func (s *StatusJSON) ReadFile(path string) error {
	file, err := os.Open(path)
	if err != nil {
		return err
	}
	defer file.Close()

	return s.Read(file)
}

type StatusFactory struct {
	path string
	json StatusJSON
}

func NewStatusFactory(path string) StatusFactory {
	return StatusFactory{
		path: path,
	}
}

func (s StatusFactory) Handler(c echo.Context) error {
	if err := s.json.ReadFile(s.path); err != nil {
		return err
	}
	return c.JSON(http.StatusOK, s.json)
}
