package types

type Query map[string]interface{}

type Pagination struct {
	PageSize int
	Page     int
	Total    int64
}

type Response struct {
	Pagination
	Code    int
	Message string
	Data    interface{}
}

type Request struct {
	Query Query
	Pagination
}

type TaskStatus int

type ServiceStatus int
