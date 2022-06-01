package handlers

import (
	"net/http"
	"postman/internal/httptools"
	"postman/internal/model"
	"postman/internal/responsewriter"
	"postman/internal/service"

	"github.com/gorilla/mux"
)

const (
	contentType     = "Content-Type"
	applicationJson = "application/json"
)

type PostmanHandler struct {
	postmanService *service.PostmanService
}

func NewPostmanHandler(PostmanService *service.PostmanService) *PostmanHandler {
	return &PostmanHandler{
		postmanService: PostmanService,
	}
}

func (p *PostmanHandler) AddRoutes(router *mux.Router) {
	router.HandleFunc("/helloworld", p.handler).Methods(http.MethodGet)
	router.HandleFunc("/newuser", p.handleNewUser).Methods(http.MethodPost)
	router.HandleFunc("/getuser/{id}", p.handleGetUser).Methods(http.MethodGet)
}

func (p *PostmanHandler) handler(resp http.ResponseWriter, _ *http.Request) {
	responsewriter.WriteSuccessResponse(resp, "0000", "SUCCESS", "1", "HelloWorld")
}

func (p *PostmanHandler) handleNewUser(resp http.ResponseWriter, request *http.Request) {
	resp.Header().Set(contentType, applicationJson)

	ctx := request.Context()

	requestModel := &model.PostmanRequest{}

	err := httptools.ParseRequest(request, requestModel)
	if err != nil {
		responsewriter.WriteFailResponse(resp, "0002", "Invalid Request", "1", 400)
		return
	}

	serviceResponse, err := p.postmanService.SimpleDBCreate(ctx, requestModel.GetServiceRequest())
	if err != nil {
		responsewriter.WriteFailResponse(resp, "0001", "Internal Server Error", "1", 500)
		return
	}

	responsewriter.WriteSuccessResponse(resp, "0000", "SUCCESS", "1", serviceResponse)
}

func (p *PostmanHandler) handleGetUser(resp http.ResponseWriter, request *http.Request) {
	resp.Header().Set(contentType, applicationJson)

	ctx := request.Context()

	params := httptools.GetMuxValue(ctx)
	id := params["id"]
	if id == "" {
		responsewriter.WriteFailResponse(resp, "0002", "Invalid Request", "1", 400)
		return
	}

	serviceResponse, err := p.postmanService.SimpleDBSearch(ctx, &model.ServiceRequest{ID: id})
	if err != nil {
		responsewriter.WriteFailResponse(resp, "0001", "Internal Server Error", "1", 500)
		return
	}

	responsewriter.WriteSuccessResponse(resp, "0000", "SUCCESS", "1", serviceResponse)
}