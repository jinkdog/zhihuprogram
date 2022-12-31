package boot

import (
	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	g "main/app/global"
	"os"
	"time"
)

func LoggerSetup() {
	dynamicLevel := zap.NewAtomicLevel() //定义日志等级
	//dynamicLevel动态等级//AtomicLevel原子等级

	switch g.Config.Logger.LogLevel {
	case "debug": //设置debug等级
		dynamicLevel.SetLevel(zap.DebugLevel)
	case "info": //设置info等级
		dynamicLevel.SetLevel(zap.InfoLevel)
	case "warn": //设置warn等级
		dynamicLevel.SetLevel(zap.WarnLevel)
	case "error": //设置error等级
		dynamicLevel.SetLevel(zap.ErrorLevel)
	}

	encoder := zapcore.NewConsoleEncoder(zapcore.EncoderConfig{ //设置编码器
		MessageKey:     "message",                        // 日志消息键
		LevelKey:       "level",                          // 日志等级键
		TimeKey:        "time",                           // 时间键
		NameKey:        "logger",                         // 日志记录器名
		CallerKey:      "caller",                         // 日志文件信息键
		StacktraceKey:  "stacktrace",                     // 堆栈键
		LineEnding:     zapcore.DefaultLineEnding,        // 友好日志换行符
		EncodeLevel:    zapcore.CapitalColorLevelEncoder, // 友好日志等级名大小写（info INFO）
		EncodeTime:     CustomTimeEncoder,                // 友好日志时日期格式化
		EncodeDuration: zapcore.StringDurationEncoder,    // 时间序列化
		EncodeCaller:   zapcore.FullCallerEncoder,        // 日志文件信息 short（包/文件.go:行号） full (文件位置.go:行号)
	})

	cores := [...]zapcore.Core{ //此处...意义为自行推断数组长度
		zapcore.NewCore(encoder, os.Stdout, dynamicLevel), //输出到控制台
		zapcore.NewCore(
			encoder,
			zapcore.AddSync(&lumberjack.Logger{
				Filename:   g.Config.Logger.SavePath,
				MaxSize:    g.Config.Logger.MaxSize,
				MaxAge:     g.Config.Logger.MaxAge,
				MaxBackups: g.Config.Logger.MaxBackups,
				LocalTime:  true,
				Compress:   g.Config.Logger.IsCompress,
			}), //将日志输出到文件当中
			dynamicLevel,
		),
	}

	g.Logger = zap.New(zapcore.NewTee(cores[:]...), zap.AddCaller()) //将数组打散成一个个个体
	defer func(Logger *zap.Logger) {
		_ = Logger.Sync()
	}(g.Logger)

	g.Logger.Info("initialize logger successfully!")
}

// CustomTimeEncoder 格式化时间//自定义输出日志的格式
func CustomTimeEncoder(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendString(t.Format("[2006-01-02 15:04:05.000]"))
}

func getWriteSyncer(file string) zapcore.WriteSyncer { //为了使日志能保存到文件中而使用的一个插件
	lumberJackLogger := &lumberjack.Logger{ //可以自动将日志文件切割，使其从大文件变为小文件//可以使用os包替代其使用
		Filename:   file,                       // 日志保存位置
		MaxSize:    g.Config.Logger.MaxSize,    // 日志文件最大大小 (MB)
		MaxBackups: g.Config.Logger.MaxBackups, // 日志文件备份数量
		MaxAge:     g.Config.Logger.MaxAge,     // 日志保存时间
		Compress:   true,
	}
	return zapcore.AddSync(lumberJackLogger)
}
