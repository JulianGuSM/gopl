package main

import (
	"go.uber.org/zap"
	"time"
)

/**
 * @Author: sm.gu
 * @Email: sm.gu@aftership.com
 * @Date: 2021/7/22 1:43 下午
 * @Desc:
 */

func main() {
	logger, _ := zap.NewProduction()
	defer logger.Sync()

	const url = "https://example.com"

	sugar := logger.Sugar()
	sugar.Infow("Failed to fetch URL. err:%s duration:%s 10",
		"url", url,
		"a", 3,
		"backoff", time.Second,
	)
	sugar.Infof("Failed to fetch URL:%s", url)

	logger.Info("failed to fetch URL",
		zap.String("url", url),
		zap.Int("attempt", 3),
		zap.Duration("backoff", time.Second),
	)

}
