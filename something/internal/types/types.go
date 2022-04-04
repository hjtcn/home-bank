// Code generated by goctl. DO NOT EDIT.
package types

type Request struct {
	Name string `path:"name,options=you|me|BenBird|Annie"`
}

type Response struct {
	Message string `json:"message"`
}

type CreateRequest struct {
	Name        string `json:"name"`
	CategoryId  int64  `json:"category_id"`
	ContainerId int64  `json:"container_id"`
}

type CreateResponse struct {
	Code    int64  `json:"code"`
	Message string `json:"message"`
}
