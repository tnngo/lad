package lad

import (
	"io"
	"os"
	"time"

	"github.com/tnngo/lad/ladcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

type LadOption struct {
	Level    ladcore.Level
	Filename string
}

func Build(opts ...*LadOption) {
	var opt *LadOption
	if len(opts) == 0 {
		opt = &LadOption{
			Level: DebugLevel,
		}
	} else {
		opt = opts[0]
	}

	var cores []ladcore.Core

	// 填充命令行配置
	cores = append(cores, (&Console{
		Level: opt.Level,
	}).Mode())

	if opt.Filename != "" {
		cores = append(cores, (&File{
			Filename:   opt.Filename,
			MaxSize:    64,
			MaxBackups: 10,
			MaxAge:     30,
			Compress:   true,
			LapLevel:   opt.Level,
		}).Mode())
		/** 定义日志文件输出核心。 */
		hook := &lumberjack.Logger{}

		fileWrite := ladcore.AddSync(io.MultiWriter(hook))
		fileConfig := NewProductionEncoderConfig()
		fileConfig.EncodeTime = timeFormat
		fileCore := ladcore.NewCore(
			ladcore.NewConsoleEncoder(fileConfig),
			fileWrite,
			opt.Level,
		)
		cores = append(cores, fileCore)

	}
	core := ladcore.NewTee(cores...)

	ReplaceGlobals(New(core, AddCaller()))
}

// 日志时间格式。
func timeFormat(t time.Time, enc ladcore.PrimitiveArrayEncoder) {
	enc.AppendString(t.Format("2006-01-02 15:04:05.000"))
}

type Console struct {
	Level ladcore.Level
}

func (c *Console) Mode() ladcore.Core {
	write := ladcore.AddSync(io.MultiWriter(os.Stdout))
	config := NewProductionEncoderConfig()
	config.EncodeTime = timeFormat
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

func (c *Console) Build() {
	ReplaceGlobals(New(c.Mode(), AddCaller()))
}

type File struct {
	// Level 日志级别，默认为info。
	LapLevel ladcore.Level
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

func (f *File) Mode() ladcore.Core {
	hook := &lumberjack.Logger{
		Filename:   f.Filename,
		MaxSize:    f.MaxSize,
		MaxBackups: f.MaxBackups,
		MaxAge:     f.MaxAge,
		Compress:   f.Compress,
	}
	write := ladcore.AddSync(io.MultiWriter(hook))
	config := NewProductionEncoderConfig()
	config.EncodeTime = timeFormat
	return ladcore.NewCore(
		ladcore.NewConsoleEncoder(config),
		write,
		f.LapLevel,
	)
}

func (f *File) Build() {
	ReplaceGlobals(New(f.Mode(), AddCaller()))
}
