package router

import (
	_ "autoOps/docs"
	"github.com/gorilla/mux"
	"github.com/swaggo/http-swagger"
	"net/http"
)

// @title Swagger Example API
// @version 1.0
// @description This is a sample server Petstore server.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host petstore.swagger.io
// @BasePath /v2
func InitRouter() http.Handler {
	r := mux.NewRouter()
	//r := chi.NewRouter()

	r.PathPrefix("/swagger/").HandlerFunc(httpSwagger.Handler(
		httpSwagger.URL("http://127.0.0.1:8002/swagger/doc.json"),
	)).Methods(http.MethodGet)
	taskHandle := &FakeHandle{}
	//任务的增删改查
	r.HandleFunc("/api/v1/tasks", taskHandle.fake).Methods(http.MethodGet)
	r.HandleFunc("/api/v1/task", taskHandle.fake).Methods(http.MethodPost)
	r.HandleFunc("/api/v1/task/{id}", taskHandle.fake).Methods(http.MethodDelete)
	r.HandleFunc("/api/v1/task/{id}", taskHandle.fake).Methods(http.MethodPut)
	r.HandleFunc("/api/v1/task/{id}", taskHandle.fake).Methods(http.MethodGet)

	//agent的增删改查
	r.HandleFunc("/api/v1/agents", taskHandle.fake).Methods(http.MethodGet)
	r.HandleFunc("/api/v1/agent", taskHandle.fake).Methods(http.MethodPost)
	r.HandleFunc("/api/v1/agent/{id}", taskHandle.fake).Methods(http.MethodDelete)
	r.HandleFunc("/api/v1/agent/{id}", taskHandle.fake).Methods(http.MethodPut)
	r.HandleFunc("/api/v1/agent/{id}", taskHandle.fake).Methods(http.MethodGet)

	//script的增删改查
	r.HandleFunc("/api/v1/scripts", taskHandle.fake).Methods(http.MethodGet)
	r.HandleFunc("/api/v1/script", taskHandle.fake).Methods(http.MethodPost)
	r.HandleFunc("/api/v1/script/{id}", taskHandle.fake).Methods(http.MethodDelete)
	r.HandleFunc("/api/v1/script/{id}", taskHandle.fake).Methods(http.MethodPut)
	r.HandleFunc("/api/v1/script/{id}", taskHandle.fake).Methods(http.MethodGet)

	//任务结果的增删改查
	r.HandleFunc("/api/v1/task_results", taskHandle.fake).Methods(http.MethodGet)
	r.HandleFunc("/api/v1/task_result", taskHandle.fake).Methods(http.MethodPost)
	r.HandleFunc("/api/v1/task_result/{id}", taskHandle.fake).Methods(http.MethodDelete)
	r.HandleFunc("/api/v1/task_result/{id}", taskHandle.fake).Methods(http.MethodPut)
	r.HandleFunc("/api/v1/task_result/{id}", taskHandle.fake).Methods(http.MethodGet)

	return r
}

type FakeHandle struct {
}

func (receiver *FakeHandle) fake(w http.ResponseWriter, r *http.Request) {
	panic("我是假的")
}
