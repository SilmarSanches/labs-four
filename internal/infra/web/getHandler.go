package web

import (
	"labs-four/config"
	"labs-four/internal/usecases"
	"net/http"
)

type GetHelloHandler struct {
	config       *config.AppSettings
	HelloUseCase usecases.HelloUseCaseInterface
}

func NewGetHelloHandler(helloUC usecases.HelloUseCaseInterface, cfg *config.AppSettings) *GetHelloHandler {
	return &GetHelloHandler{
		config:       cfg,
		HelloUseCase: helloUC,
	}
}

// HandleLabsFour godoc
// @Summary Endpoint simples que retorna uma string
// @Description Endpoint que retorna uma mensagem de "Hello, World!"
// @Tags Labs-Four
// @Accept json
// @Produce plain
// @Success 200 {string} string "OK"
// @Router /hello [get]
func (h *GetHelloHandler) HandleHello(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "text/plain")
    w.WriteHeader(http.StatusOK)
    w.Write([]byte(h.HelloUseCase.Hello()))
}

