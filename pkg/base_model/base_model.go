package base_model

import (
	"eshort/pkg/easylogger"
	"eshort/pkg/model"
	"eshort/pkg/mtime"
	"eshort/pkg/struct_tool"
	"eshort/pkg/types"
)

type BaseModel struct {
	ID        uint64       `json:"id"`
	CreatedAt *mtime.Mtime `gorm:"column:created_at;index" json:"created_at"`
	UpdatedAt *mtime.Mtime `gorm:"column:updated_at;index" json:"updated_at"`
}

var ErrorMsg = map[string]string{
	"record not found": "参数有误",
}

func (model BaseModel) GetStringID() string {
	return types.Uint64ToString(model.ID)
}

func (model BaseModel) GetError(err error, msg string) string {
	return GetError(err, msg)
}

func GetError(err error, msg string) string {
	if s, ok := ErrorMsg[err.Error()]; ok {
		return s
	}
	return msg
}

func GetById[T any](data T, uint64id uint64) (T, error) {
	//id := types.StringToUint64(idstr)
	err := model.DB.First(&data, uint64id).Error
	if err != nil {
		return data, err
	}
	return data, nil
}

func Create[T any](data T) error {
	err := model.DB.Create(&data).Error
	if err != nil {
		easylogger.LogError(err, "添加失败")
		return err
	}
	return nil
}

func ToVo(vo any, src any) {
	struct_tool.Copy(vo, src)
}
