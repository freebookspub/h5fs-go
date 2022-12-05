package webapis

import (
	"h5fs/model"
)

type ListRequest struct {
	Username string `json:"username"`
	Offset   int    `json:"offset"`
	Limit    int    `json:"limit"`
}

type GetIndexListsRequest struct {
	Offset int `json:"offset"`
	Page   int `json:"page"`
}

type GetIndexListsRespones struct {
	TotalCount uint64        `json:"totalCount"`
	Lists      []*model.News `json:"lists"`
}

type GetIndexOneRespones struct {
	Lists *model.News `json:"lists"`
}

type GetNewsOneRespones struct {
	Lists *model.Details `json:"lists"`
}
type GetPostsOneRespones struct {
	Lists *model.NewsDeatis `json:"lists"`
}

type GetCategoryListsRespones struct {
	TotalCount uint64            `json:"totalCount"`
	Lists      []*model.Category `json:"lists"`
}

type GetReadListsRespones struct {
	Lists *model.Readinfo `json:"lists"`
}
