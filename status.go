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
	Name              string `Json:"Name"`
	ID                string `Json:"ID"`
	Session           bool   `Json:"Session"`
	Login             bool   `Json:"Login"`
	LastRequestStatus bool   `Json:"LastRequestStatus"`
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
	Path string
	Json StatusJSON
}

func NewStatusFactory(path string) StatusFactory {
	return StatusFactory{
		Path: path,
	}
}

func (s StatusFactory) Handler(c echo.Context) error {
	//if err := s.Json.ReadFile(s.Path); err != nil {
	//	return err
	//}
	return c.JSON(http.StatusOK, s.Json)
}
