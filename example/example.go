package main

import (
	"github.com/ethreads/clog"
	"github.com/ethreads/clog/cloghttp"
	"github.com/ethreads/clog/storage"
	"net/http"
	"os"
	"sync"
	"time"
)

func main() {
	//changeLogLevel()
	//newSearchExample()
	writeLogFile()
}

func writeLogFile() {

	path := "./log/api.log"
	//s := storage.NewSizeSplitFile(path).Backups(10).MaxSize(50).SaveTime(4).Compress(3).Finish()
	s := storage.NewTimeSplitFile(path, time.Minute).Backups(3).SaveTime(3).Compress(2).Finish()
	defer s.Close()
	clog.NewOption().WithTimestamp().WithPrefix("host", "local").WithWriter(s).WithRandom(1000).Default()
	clog.Set.TimeFormat(time.RFC3339Nano)
	clog.Set.FiledName().TimestampFieldName("tm")
	clog.SetGlobalLevel(clog.Level(2))
	clog.Info().Msg("receive req data")
	log := clog.CopyDefault()
	log.ResetStrPrefix("host", "default")
	log.Info().Msg("receive req data")
	log.Random(-159403579887936999).Msg("receive req data")
	wg := &sync.WaitGroup{}
	for i := 0; i < 100; i++ {
		wg.Add(1)
		gid := i
		go func() {
			for j := -1000; j < 1000; j++ {
				log.Random(int64(j)).Int("goroutine id", gid).Int("idx", j).Str("msg", "foo").Msg("bar")
				log.Random(int64(j)).Int("goroutine id", gid).Int("idx-1", j).Str("msg", "foo").Msg("bar")
			}
			wg.Done()
		}()
	}
	clog.Info().Msg("receive req data")
	wg.Wait()
}

func changeLogLevel() {
	var mux = http.NewServeMux()
	mux.Handle("/changelog", cloghttp.Handler())
	srv := http.Server{
		Addr:    ":9999",
		Handler: mux,
	}
	go func() {
		if err := srv.ListenAndServe(); err != nil {
			if err == http.ErrServerClosed {
				clog.Info().Msg("service exit...")
				return
			}
			panic(err)
		}
	}()
	clog.NewOption().WithLogLevel(clog.TraceLevel).WithTimestamp().WithWriter(os.Stdout).Default()
	clog.Set.BaseTimeDurationInteger()
	for i := 0; i < 1000; i++ {
		clog.Trace().Msg("trace")
		clog.Debug().Msg("debug")
		clog.Info().Msg("info")
		clog.Warn().Msg("warn")
		clog.Error().Msg("error")
		clog.Log().Msg("log")
		time.Sleep(time.Millisecond * 300)
	}
}

type SearchExample struct {
	count int
	log   *clog.Logger
}

func newSearchExample() {
	clog.NewOption().WithLogLevel(clog.InfoLevel).WithWriter(os.Stdout).Default()
	example := SearchExample{
		log: clog.CopyDefault(),
	}
	example.log.ResetStrPrefix("searchId", time.Now().UnixNano())
	example.log.Info().Interface("data", 1).Msg("")
	example.log.Error().Interface("data", 1).Msg("")
	example.log.Info().Interface("data", 1).Msg("")
	example.log.AppendStrPrefix("sec", time.Now().UnixNano())
	example.log.Info().Interface("data", 1).Msg("")
	example.log.Info().Interface("data", 1).Msg("")
	example.log.Info().Interface("data", 1).Msg("")
}
