package words

import (
	"context"
	v1 "easyword/api/words/v1"
	"easyword/internal/dao"
	"easyword/internal/model/do"

	"github.com/gogf/gf/errors/gerror"
)

type Words struct {
}

func New() *Words {
	return &Words{}
}

type CreateInput struct {
	Uid                uint
	Word               string
	Definition         string
	ExampleSentence    string
	ChineseTranslation string
	Pronunciation      string
	ProficiencyLevel   v1.ProficiencyLevel
}

type UpdateInput struct {
	Uid                uint
	Word               string
	Definition         string
	ExampleSentence    string
	ChineseTranslation string
	Pronunciation      string
	ProficiencyLevel   v1.ProficiencyLevel
}

func (w *Words) Create(ctx context.Context, in CreateInput) error {
	var cls = dao.Words.Columns()

	count, err := dao.Words.Ctx(ctx).
		Where(cls.Uid, in.Uid).
		Where(cls.Word, in.Word).Count()
	if err != nil {
		return err
	}
	if count > 0 {
		return gerror.New("单词已存在")
	}

	_, err = dao.Words.Ctx(ctx).Data(do.Words{
		Uid:                in.Uid,
		Word:               in.Word,
		Definition:         in.Definition,
		ExampleSentence:    in.ExampleSentence,
		ChineseTranslation: in.ChineseTranslation,
		Pronunciation:      in.Pronunciation,
		ProficiencyLevel:   in.ProficiencyLevel,
	}).Insert()
	if err != nil {
		return err
	}
	return nil
}

func (w *Words) Update(ctx context.Context, id uint, in UpdateInput) error {
	var cls = dao.Words.Columns() // 获取数据库表列名常量

	// 检查更新后的单词是否已存在(排除当前单词)
	count, err := dao.Words.Ctx(ctx).
		Where(cls.Uid, in.Uid).   // 限定当前用户
		Where(cls.Word, in.Word). // 检查单词内容
		Where(cls.Id, id).        // 排除当前要更新的单词
		Count()                   // 统计匹配记录数
	if err != nil {
		return err // 返回查询错误
	}
	if count > 0 {
		return gerror.New("单词已存在") // 返回单词已存在错误
	}

	// 执行单词更新操作
	_, err = dao.Words.Ctx(ctx).Data(do.Words{
		Word:               in.Word,               // 更新单词内容
		Definition:         in.Definition,         // 更新单词定义
		ExampleSentence:    in.ExampleSentence,    // 更新例句
		ChineseTranslation: in.ChineseTranslation, // 更新中文翻译
		Pronunciation:      in.Pronunciation,      // 更新发音
		ProficiencyLevel:   in.ProficiencyLevel,   // 更新熟练度
	}).Where(cls.Id, id). // 限定更新指定ID的单词
				Where(cls.Uid, in.Uid). // 确保只能更新当前用户的单词
				Update()                // 执行更新
	if err != nil {
		return err // 返回更新错误
	}
	return nil // 更新成功
}
