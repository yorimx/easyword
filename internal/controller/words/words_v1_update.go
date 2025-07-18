package words

import (
	"context"

	v1 "easyword/api/words/v1"
	"easyword/internal/logic/words"
)

func (c *ControllerV1) Update(ctx context.Context, req *v1.UpdateReq) (res *v1.UpdateRes, err error) {
	uid, err := c.users.GetUid(ctx) // 从上下文中获取当前用户ID
	if err != nil {
		return nil, err // 返回获取用户ID失败的错误
	}
	// 调用业务逻辑层的Update方法更新单词
	err = c.words.Update(ctx, req.Id, words.UpdateInput{
		Uid:                uid,                    // 设置用户ID
		Word:               req.Word,               // 设置单词内容
		Definition:         req.Definition,         // 设置单词定义
		ExampleSentence:    req.ExampleSentence,    // 设置例句
		ChineseTranslation: req.ChineseTranslation, // 设置中文翻译
		Pronunciation:      req.Pronunciation,      // 设置发音
		ProficiencyLevel:   req.ProficiencyLevel,   // 设置熟练度
	})
	return nil, err // 返回操作结果(成功为nil，失败为错误)
}
