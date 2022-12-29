package eshort

import "eshort/pkg/base_model"

type Eshort struct {
	base_model.BaseModel
	//shortkey依然唯一，使用domain，减少索引长度
	Ext               string
	ShortKey          string
	FullData          string
	Exp               uint64
	Status, IsDeleted uint8
}
