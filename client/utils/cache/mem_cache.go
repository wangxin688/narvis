package mem_cache

import (
	"time"

	"github.com/patrickmn/go-cache"
	"github.com/wangxin688/narvis/client/utils/logger"
)

var MemCache cache.Cache

func InitCache() {
	c := cache.New(24*time.Hour, 48*time.Hour)
	MemCache = *c
	logger.Logger.Info("[memCacheInit]: Init cache successfully")
}
