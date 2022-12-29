package controller

import (
	"eshort/app/biz"
	"eshort/app/http/controller/base"
	"eshort/app/models/eshort"
	"eshort/app/services/shortkey"
	"eshort/pkg/eredis"
	rsp "eshort/pkg/response"
	"github.com/gin-gonic/gin"
	"github.com/spf13/cast"
	"net/http"
	"time"
)

type Index struct {
	base.BaseController
}

// Generate 短链转换 组合 Passphrase
func (a *Index) Generate(c *gin.Context) {
	fulldata := c.PostForm("data")
	exp := cast.ToInt(c.PostForm("exp"))
	exps := map[int]int64{
		1:   86400,
		7:   7 * 86400,
		30:  30 * 86400,
		365: 365 * 86400,
	}
	if _, ok := exps[exp]; !ok {
		rsp.ErrMsg(c, "过期时间选择有误")
		return
	}
	vali, err := biz.BIZ.GenerateVali(fulldata)
	if !vali {
		//api.gen.vali.err
		rsp.ErrMsg(c, err.Error())
		return
	}
	skey, err := shortkey.Take()
	if err != nil {
		//api.gen.vali.err
		rsp.ErrMsg(c, "有误，暂不可用")
		return
	}
	data := biz.BIZ.GenerateResult(skey, gin.H{})
	//api.gen.vali.ok
	//下发消息，使用了这一条消息，标记该条消息已使用，进入布隆过滤器
	//获取已经注册过的策略接口，调用转换方法，返回
	update := eshort.Eshort{Status: 1, FullData: fulldata, Exp: uint64(exps[exp] + time.Now().Unix()), Ext: biz.BIZ.GetExt()}
	eredis.JoinBloom(skey)
	eshort.UpdateByKey(skey, update)
	rsp.RepData(c, data, "转换成功")
	return
	//hashKey(fulldata)
	//1.有效时间
	//2.长链接
	//3.防攻击，布隆过滤器，
	//4.同一个长链生成出来的短链，应当不一致。
	//5.依然采用发号器，号码生成为当前 murmur32(毫秒数+随机串+长链)。短链唯一索引。需满足重启恢复。
	//一，验证
	//同一个长链接，转换出来的短链接应当一致。取号器如何实现？
	//
	//没有取号服务api端从redis中取pop（利用通道，同时查询数量），有号码生成服务。
	//号码生成服务提供的接口：
	//生成单条key，供api服务在redis中取不到数据时使用。如果当时没有正在生成数据（redis设置一个状态），立即生成。
	//接收异步通知，批量生成key。当刚生成完后的60秒内，再次接收到单条key或者异步通知，则生成数量翻倍。
	//考虑批量插入，非唯一报错怎么办，怎么检测？
}

// Agent 接受访问
func (a *Index) Agent(c *gin.Context) {
	key := c.Param("key")
	vali, err := biz.BIZ.AgentVali(key)
	if !vali {
		//api.agent.vali.err
		rsp.ErrMsg(c, err.Error())
		return
	}
	key = biz.BIZ.ExtracKey(key)
	if !eredis.InBloom(key) {
		rsp.ErrMsg(c, "有误")
		return
	}
	shortData, err := eshort.GetShortByKey(key)
	if err != nil {
		rsp.ErrMsg(c, err.Error())
		return
	}
	ntime := uint64(time.Now().Unix())
	if shortData.Exp < ntime { //
		rsp.ErrMsg(c, "已过期")
		return
	}
	result, h, s := biz.BIZ.AgentResult(key, shortData.FullData, gin.H{})
	if s == "redirect" {
		c.Redirect(http.StatusFound, result)
		return
	}
	if s == "response" {
		rsp.RepData(c, h)
		return
	}
	//获取访问用户IP，生产日志进消息队列，是否使用go?
}
