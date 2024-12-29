package ladglobal

import (
	"io"
	"os"
	"time"

	"github.com/tnngo/lad"
	"github.com/tnngo/lad/ladcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

type Builder interface {
	mode() ladcore.Core
}

type Console struct {
	Level      ladcore.Level
	TimeFormat string
}

func (c *Console) mode() ladcore.Core {
	write := ladcore.AddSync(io.MultiWriter(os.Stdout))
	config := lad.NewProductionEncoderConfig()
	config.EncodeTime = func(t time.Time, pae ladcore.PrimitiveArrayEncoder) {
		if c.TimeFormat == "" {
			pae.AppendString(t.Format("2006-01-02 15:04:05.000"))
		} else {
			pae.AppendString(t.Format(c.TimeFormat))
		}
	}

	// 控制台输出颜色
	config.EncodeLevel = ladcore.CapitalColorLevelEncoder
	// 定义日志核心
	return ladcore.NewCore(
		// 控制台
		ladcore.NewConsoleEncoder(config),
		write,
		c.Level,
	)
}

func BuildDefaultConsole() {
	c := &Console{
		Level: lad.DebugLevel,
	}
	lad.ReplaceGlobals(lad.New(c.mode(), lad.AddCaller()))
}

type File struct {
	// Level 日志级别，默认为info。
	LadLevel ladcore.Level
	// TimeFormat 日期格式
	TimeFormat string
	// Filename 日志文件名称。
	Filename string `json:"filename"`
	// MaxSize 日志最大尺寸，默认为100MB。
	MaxSize int `json:"max_size" yaml:"max_size"`
	// MaxBackups 最大备份数量。
	MaxBackups int `json:"max_backups" yaml:"max_backups"`
	// MaxAge 最大保存时间。
	MaxAge int `json:"max_age" yaml:"max_age"`
	// Compress 是否压缩打包。
	Compress bool `json:"compress"`
}

func (f *File) mode() ladcore.Core {
	hook := &lumberjack.Logger{
		Filename:   f.Filename,
		MaxSize:    f.MaxSize,
		MaxBackups: f.MaxBackups,
		MaxAge:     f.MaxAge,
		Compress:   f.Compress,
	}
	write := ladcore.AddSync(io.MultiWriter(hook))
	config := lad.NewProductionEncoderConfig()
	config.EncodeTime = func(t time.Time, pae ladcore.PrimitiveArrayEncoder) {
		if f.TimeFormat == "" {
			pae.AppendString(t.Format("2006-01-02 15:04:05.000"))
		} else {
			pae.AppendString(t.Format(f.TimeFormat))
		}

	}
	return ladcore.NewCore(
		ladcore.NewConsoleEncoder(config),
		write,
		f.LadLevel,
	)
}

func BuildDefaultFile() {
	var cores []ladcore.Core
	cores = append(cores, (&Console{
		Level: lad.DebugLevel,
	}).mode())

	f := &File{
		LadLevel:   lad.DebugLevel,
		MaxSize:    64,
		MaxBackups: 10,
		MaxAge:     30,
		Compress:   true,
	}
	cores = append(cores, f.mode())
	core := ladcore.NewTee(cores...)

	lad.ReplaceGlobals(lad.New(core, lad.AddCaller()))
}

func Build(builders ...Builder) {
	var cores []ladcore.Core
	if len(builders) == 0 {
		cores = append(cores, (&Console{
			Level: lad.DebugLevel,
		}).mode())
	} else {
		for _, v := range builders {
			cores = append(cores, v.mode())
		}
	}

	core := ladcore.NewTee(cores...)

	lad.ReplaceGlobals(lad.New(core, lad.AddCaller()))
}
