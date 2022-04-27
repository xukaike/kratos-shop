package data

import (
	"context"
	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/log"
	"goods/internal/biz"
	"goods/internal/domain"
	"gorm.io/gorm"
	"time"
)

// SpecificationsAttr 规格参数信息表
type SpecificationsAttr struct {
	ID        int64          `gorm:"primarykey;type:int" json:"id"`
	TypeID    int64          `gorm:"index:type_id;type:int;comment:商品类型ID;not null"`
	Name      string         `gorm:"type:varchar(250);not null;comment:规格参数名称" json:"name"`
	Sort      int32          `gorm:"comment:规格排序;default:99;not null;type:int" json:"sort"`
	Status    bool           `gorm:"comment:参数状态;default:false" json:"status"`
	IsSKU     bool           `gorm:"comment:是否通用的SKU持有;default:false" json:"is_sku"`
	IsSelect  bool           `gorm:"comment:是否可查询;default:false" json:"is_select"`
	CreatedAt time.Time      `gorm:"column:add_time" json:"created_at"`
	UpdatedAt time.Time      `gorm:"column:update_time" json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at"`
}

// SpecificationsAttrValue 规格参数信息选项表
type SpecificationsAttrValue struct {
	ID        int64          `gorm:"primarykey;type:int" json:"id"`
	AttrId    int64          `gorm:"index:attr_id;type:int;comment:规格ID;not null"`
	Value     string         `gorm:"type:varchar(250);not null;comment:规格参数信息值" json:"value"`
	Sort      int32          `gorm:"comment:规格参数值排序;default:99;not null;type:int" json:"sort"`
	CreatedAt time.Time      `gorm:"column:add_time" json:"created_at"`
	UpdatedAt time.Time      `gorm:"column:update_time" json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at"`
}

type specificationRepo struct {
	data *Data
	log  *log.Helper
}

// NewSpecificationRepo .
func NewSpecificationRepo(data *Data, logger log.Logger) biz.SpecificationRepo {
	return &specificationRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (r *specificationRepo) CreateSpecification(ctx context.Context, req *domain.Specification) (int64, error) {
	s := &SpecificationsAttr{
		TypeID:    req.TypeID,
		Name:      req.Name,
		Sort:      req.Sort,
		Status:    req.Status,
		IsSKU:     req.IsSKU,
		IsSelect:  req.IsSelect,
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
	}
	if err := r.data.DB(ctx).Save(s).Error; err != nil {
		return 0, errors.InternalServer("SPECIFICATION_SAVED_ERROR", err.Error())
	}
	return s.ID, nil
}

func (r *specificationRepo) CreateSpecificationValue(ctx context.Context, attrId int64, req []*domain.SpecificationValue) error {
	var value []*SpecificationsAttrValue
	for _, v := range req {
		res := &SpecificationsAttrValue{
			AttrId:    attrId,
			Value:     v.Value,
			Sort:      v.Sort,
			CreatedAt: time.Time{},
			UpdatedAt: time.Time{},
		}
		value = append(value, res)
	}
	if err := r.data.DB(ctx).Create(&value).Error; err != nil {
		return errors.InternalServer("SPECIFICATION_VALUE_SAVED_ERROR", err.Error())
	}

	return nil
}
