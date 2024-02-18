package main

import (
	"fmt"

	"github.com/cdipaolo/sentiment"
)

func main() {
	// 创建情感分析器
	model, err := sentiment.Restore()
	if err != nil {
		fmt.Printf("加载情感分析模型失败: %v\n", err)
		return
	}

	// 待分析的文本
	text := "这是一个很棒的电影！"

	// 进行情感分析
	analysis := model.SentimentAnalysis(text, sentiment.ChineseSimplified)

	// 输出情感分析结果
	fmt.Printf("文本: %s\n", text)
	fmt.Println(analysis.Score)
	fmt.Printf("情感类型: %s\n", analysis.Language)
}
