package main

import (
	"github.com/robfig/cron"
	"go-gin-example/models"
	"go-gin-example/pkg/logging"
	"time"
)

func main() {
	logging.Info("Starting...")

	// 根据本地时间创建一个新（空白）的 Cron job runner
	c := cron.New()

	//AddFunc 会向 Cron job runner 添加一个 func ，以按给定的时间表运行
	c.AddFunc("* * * * * *", func() {
		logging.Info("RUN models.CleanAllTag...")
		models.CleanAllTag()
	})

	c.AddFunc("* * * * * *", func() {
		logging.Info("RUN models.ClanAllArticle...")
		models.CleanAllArticle()
	})

	// 在当前执行的程序中启动 Cron 调度程序。其实这里的主体是 goroutine + for + select + timer 的调度控制哦
	c.Start()

	t1 := time.NewTimer(time.Second * 10)
	for {
		select {
		case <- t1.C:
			t1.Reset(time.Second * 10)
		}
	}
}

/*
**（1）time.NewTimer **

会创建一个新的定时器，持续你设定的时间 d 后发送一个 channel 消息

（2）for + select

阻塞 select 等待 channel

（3）t1.Reset

会重置定时器，让它重新开始计时

注：本文适用于 “t.C 已经取走，可直接使用 Reset”。
 */