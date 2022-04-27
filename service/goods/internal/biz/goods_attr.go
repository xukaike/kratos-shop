package biz

import (
	"context"
	"errors"
	"github.com/go-kratos/kratos/v2/log"
	"goods/internal/domain"
)

type GoodsAttrRepo interface {
	CreateGoodsGroupAttr(context.Context, *domain.AttrGroup) (*domain.AttrGroup, error)
}

type GoodsAttrUsecase struct {
	repo     GoodsAttrRepo
	typeRepo GoodsTypeRepo // 引入goods type 的 repo
	tx       Transaction   // 引入事务
	log      *log.Helper
}

func NewGoodsAttrUsecase(repo GoodsAttrRepo, tx Transaction, gRepo GoodsTypeRepo, logger log.Logger) *GoodsAttrUsecase {
	return &GoodsAttrUsecase{
		repo:     repo,
		tx:       tx,
		typeRepo: gRepo,
		log:      log.NewHelper(logger),
	}
}

func (ga *GoodsAttrUsecase) CreateAttrGroup(ctx context.Context, r *domain.AttrGroup) (*domain.AttrGroup, error) {
	if r.IsTypeIDEmpty() {
		return nil, errors.New("请选择商品类型进行绑定")
	}

	_, err := ga.typeRepo.IsExistsByID(ctx, r.TypeID)
	if err != nil {
		return nil, err
	}

	attr, err := ga.repo.CreateGoodsGroupAttr(ctx, r)
	if err != nil {
		return nil, err
	}
	return attr, nil
}
