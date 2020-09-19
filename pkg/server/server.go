package server

import (
	"crypto/tls"
	"fmt"
	"github.com/labstack/echo/v4"
	"hydro_monitor/web_api/pkg/configs"
	"io/ioutil"
	"net/http"
	"time"
)

type Server interface {
	Start() error
}

type server struct {
	e          *echo.Echo
	httpServer *http.Server
}

func (s *server) Start() error {
	return s.e.StartServer(s.httpServer)
}
func loadCertFiles(server *http.Server, config *configs.Configuration) error {
	var err error
	var cert []byte
	if cert, err = ioutil.ReadFile(config.TLSCert); err != nil {
		return err
	}
	var key []byte
	if key, err = ioutil.ReadFile(config.TLSKey); err != nil {
		return err
	}
	if server.TLSConfig.Certificates[0], err = tls.X509KeyPair(cert, key); err != nil {
		return err
	}
	return nil
}

func NewServer(e *echo.Echo, config *configs.Configuration) Server {
	s := &http.Server{
		Addr:        fmt.Sprintf(":%s", config.Port),
		ReadTimeout: 25 * time.Minute,
	}
	if config.HTTPSEnabled {
		s.TLSConfig = new(tls.Config)
		s.TLSConfig.Certificates = make([]tls.Certificate, 1)
		if err := loadCertFiles(s, config); err != nil {
			e.Logger.Fatal("Failed to load TLS cert files")
		}
	}
	return &server{e: e, httpServer: s}
}
