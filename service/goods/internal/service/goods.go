package service

import (
	"context"
	"goods/internal/biz"

	v1 "goods/api/goods/v1"
)

func (g *GoodsService) CreateCategory(ctx context.Context, r *v1.CategoryInfoRequest) (*v1.CategoryInfoResponse, error) {
	result, err := g.cac.CreateGreeter(ctx, &biz.CategoryInfo{
		Name:           r.Name,
		ParentCategory: r.ParentCategory,
		Level:          r.Level,
		IsTab:          r.IsTab,
		Sort:           r.Sort,
	})
	if err != nil {
		return nil, err
	}

	return &v1.CategoryInfoResponse{
		Id:             result.ID,
		Name:           result.Name,
		ParentCategory: result.ParentCategory,
		Level:          result.Level,
		IsTab:          result.IsTab,
		Sort:           result.Sort,
	}, nil
}
