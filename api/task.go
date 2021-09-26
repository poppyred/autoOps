package api

import (
	"github.com/gorilla/mux"
	"net/http"
)

type TaskHandle struct {

}

// @Summary getTaskList
// @Description 获取JSON
// @Tags 任务
// @Success 200 {string} string "{"code": 200, "data": [...]}"
// @Success 200 {string} string "{"code": -1, "message": "抱歉未找到相关信息"}"
// @Router /api/v1/tasks [get]
// @Security Bearer
func (t *TaskHandle) getTaskList(w http.ResponseWriter, r *http.Request) {

}

func (t *TaskHandle) getTask(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	_ = vars["id"]
}
