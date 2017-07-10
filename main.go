package main

import (
	"flag"
	"fmt"
	"net/http"
	"os"
	ep "resuelve/api/endpoint"
	s "resuelve/api/service"

	"github.com/bmizerany/pat"
	"github.com/go-kit/kit/log"
	httpTransport "github.com/go-kit/kit/transport/http"
	"github.com/lnx1337/go/api"
	"github.com/lnx1337/go/config"
)

const (
	configFile = `service.invoice.toml`
)

type settings struct {
	Server map[string]string
}

var (
	cfg          settings
	listen       = flag.String("listen", ":8080", "Address to listen on")
	invSvc       = s.InvoiceSvc{}
	mux          *pat.PatternServeMux
	errorEncoder = httpTransport.ServerErrorEncoder(api.EncodeError)
)

func main() {
	flag.Parse()

	// Configuración inicial
	if err := config.Load(&cfg, configFile); err != nil {
		fmt.Println(`CONFIG_LOAD_ERR: `, err)
	}
	if addr, ok := cfg.Server[`port`]; ok {
		*listen = addr
	}

	// loggers
	logger := log.NewLogfmtLogger(os.Stderr)
	logger = log.With(logger, "listen", *listen, "caller", log.DefaultCaller)

	// Enpoints
	ViewInvoiceHandler := httpTransport.NewServer(
		ep.MakeReadInvoiceEndpoint(invSvc),
		ep.DecodeReadInoviceRequest,
		api.EncodeResponse,
		errorEncoder,
	)

	// Rutas
	mux = pat.New()
	mux.Get(`/v1/invoice/:id/:startDate/:endDate`, ViewInvoiceHandler)
	http.Handle(`/`, mux)

	logger.Log(`msg`, `SERVING`, `address`, *listen)
	logger.Log(`err`, http.ListenAndServe(*listen, nil))
}
