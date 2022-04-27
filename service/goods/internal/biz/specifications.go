package biz

import (
	"context"
	"errors"
	"github.com/go-kratos/kratos/v2/log"
	"goods/internal/domain"
)

type SpecificationRepo interface {
	CreateSpecification(context.Context, *domain.Specification) (int64, error)
	CreateSpecificationValue(context.Context, int64, []*domain.SpecificationValue) error
}

type SpecificationUsecase struct {
	repo  SpecificationRepo
	gRepo GoodsTypeRepo
	tx    Transaction
	log   *log.Helper
}

func NewSpecificationUsecase(repo SpecificationRepo, gRepo GoodsTypeRepo, tx Transaction, logger log.Logger) *SpecificationUsecase {
	return &SpecificationUsecase{
		repo:  repo,
		gRepo: gRepo,
		tx:    tx,
		log:   log.NewHelper(logger),
	}
}

func (s *SpecificationUsecase) CreateSpecification(ctx context.Context, r *domain.Specification) (int64, error) {
	var (
		id  int64
		err error
	)
	if r.IsTypeIDEmpty() {
		return id, errors.New("请选择商品类型进行绑定")
	}

	if r.IsValueEmpty() {
		return id, errors.New("请填写商品规格下的参数")
	}

	_, err = s.gRepo.IsExistsByID(ctx, r.TypeID)
	if err != nil {
		return id, err
	}

	err = s.tx.ExecTx(ctx, func(ctx context.Context) error {
		id, err = s.repo.CreateSpecification(ctx, r)
		if err != nil {
			return err
		}

		err = s.repo.CreateSpecificationValue(ctx, id, r.SpecificationValue)
		if err != nil {
			return err
		}
		return nil
	})
	return id, err
}
