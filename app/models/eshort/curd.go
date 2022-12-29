package eshort

import (
	"errors"
	"eshort/pkg/eviction/elru"
	"eshort/pkg/model"
	"fmt"
)

func UpdateByKey(key string, update Eshort) {
	model.DB.Model(&update).Where("short_key=?", key).Updates(update)
}

func GetShortByKey(key string) (Eshort, error) {
	lru := elru.ELRU
	v, in := lru.Get(key)
	fmt.Println(v, in)
	shortUrl := Eshort{}
	if in {
		t, _ := v.(Eshort)
		shortUrl = t
	} else {
		where := Eshort{
			ShortKey:  key,
			IsDeleted: 0,
			Status:    1,
		}
		take := model.DB.Model(shortUrl).Where(where).Take(&shortUrl)
		if take.RowsAffected == 0 {
			return shortUrl, errors.New("找不到对应数据")
		}
		if take.Error != nil {
			return shortUrl, errors.New("查询有误，应当入库")
		}
		_, _ = lru.Put(elru.NewNode(key, shortUrl))
	}
	return shortUrl, nil
}
