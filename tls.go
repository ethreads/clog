package clog

import (
	"sync"
	"time"
)

type LogCtx struct {
	ExchangeId int
	DspIds     map[int]struct{}
}

type logCtxWrapper struct {
	ctx        *LogCtx
	lastAccess time.Time
	mu         sync.RWMutex
}

var (
	tlsPool = sync.Map{}
)

func getCtx(sid int64) *logCtxWrapper {
	w, _ := tlsPool.LoadOrStore(sid, &logCtxWrapper{
		ctx: &LogCtx{DspIds: make(map[int]struct{})},
	})
	wrapper := w.(*logCtxWrapper)
	wrapper.lastAccess = time.Now()
	return wrapper
}

func SetCtx(sid int64, exchangeId, dspId int) {
	wrapper := getCtx(sid)
	wrapper.mu.Lock()
	wrapper.ctx.ExchangeId = exchangeId
	wrapper.ctx.DspIds[dspId] = struct{}{}
	wrapper.mu.Unlock()
}

func SetCtxEid(sid int64, exchangeId int) {
	wrapper := getCtx(sid)
	wrapper.mu.Lock()
	wrapper.ctx.ExchangeId = exchangeId
	wrapper.mu.Unlock()
}

func SetCtxDid(sid int64, dspId int) {
	wrapper := getCtx(sid)
	wrapper.mu.Lock()
	wrapper.ctx.DspIds[dspId] = struct{}{}
	wrapper.mu.Unlock()
}

func ClearCtx(sid int64) {
	tlsPool.Delete(sid)
}

const (
	cleanInterval = 5 * time.Second
	maxIdle       = 5 * time.Second
)

func init() {
	go cleaner()
}

func cleaner() {
	ticker := time.NewTicker(cleanInterval)
	defer ticker.Stop()

	for range ticker.C {
		now := time.Now()
		tlsPool.Range(func(key, value interface{}) bool {
			w := value.(*logCtxWrapper)
			if now.Sub(w.lastAccess) > maxIdle {
				tlsPool.Delete(key)
			}
			return true
		})
	}
}
