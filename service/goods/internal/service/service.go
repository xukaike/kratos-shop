package service

import (
	"github.com/google/wire"
	v1 "goods/api/goods/v1"
	"goods/internal/biz"
)

type GoodsService struct {
	v1.UnimplementedGoodsServer
	cac *biz.CategoryUsecase
	gt  *biz.GoodsTypeUsecase
	bc  *biz.BrandUsecase
	sc  *biz.SpecificationUsecase
	gac *biz.GoodsAttrUsecase
}

func NewGoodsService(
	cac *biz.CategoryUsecase,
	bc *biz.BrandUsecase,
	gt *biz.GoodsTypeUsecase,
	sc *biz.SpecificationUsecase,
	gac *biz.GoodsAttrUsecase,
) *GoodsService {
	return &GoodsService{
		cac: cac,
		bc:  bc,
		gt:  gt,
		sc:  sc,
		gac: gac,
	}
}

// ProviderSet is service providers.
var ProviderSet = wire.NewSet(NewGoodsService)
