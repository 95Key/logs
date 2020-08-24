package logs

import (
	"os"
	"os/user"
	"path"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

var log *zap.Logger

func init() {


	hook := lumberjack.Logger{
		Filename:   logName(), // 日志文件路径
		MaxSize:    1,                    // 每个日志文件保存的最大尺寸 单位：M
		MaxBackups: 1,                    // 日志文件最多保存多少个备份
		MaxAge:     2,                    // 文件最多保存多少天
		Compress:   true,                 // 是否压缩
	}

	encoderConfig := zapcore.EncoderConfig{
		TimeKey:        "ts",
		LevelKey:       "level",
		NameKey:        "name",
		CallerKey:      "caller",
		MessageKey:     "msg",
		StacktraceKey:  "stacktrace",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.LowercaseLevelEncoder,  // 小写编码器
		EncodeTime:     zapcore.ISO8601TimeEncoder,     // ISO8601 UTC 时间格式
		EncodeDuration: zapcore.SecondsDurationEncoder, //
		EncodeCaller:   zapcore.ShortCallerEncoder,     // FullCallerEncoder 全路径编码器  ShortCallerEncoder
		EncodeName:     zapcore.FullNameEncoder,
	}

	// 设置日志级别
	atomicLevel := zap.NewAtomicLevel()
	atomicLevel.SetLevel(zap.DebugLevel)

	core := zapcore.NewCore(
		zapcore.NewJSONEncoder(encoderConfig),                                           // 编码器配置
		zapcore.NewMultiWriteSyncer(zapcore.AddSync(os.Stdout), zapcore.AddSync(&hook)), // 打印到控制台和文件
		atomicLevel, // 日志级别
	)

	// 开启开发模式，堆栈跟踪
	caller := zap.AddCaller()
	AddCallerSkip := zap.AddCallerSkip(1)
	// 开启文件及行号
	development := zap.Development()
	// 设置初始化字段
	//filed := zap.Fields(zap.String("serviceName", "serviceName"))
	// 构造日志
	log = zap.New(core, caller, AddCallerSkip, development, zap.Fields())
	log.Info("系统启动")
}


func logName()string{
	u,err:= user.Current()
	if err!= nil {
		println("获取不到 user.Current",err.Error())
		panic(err)
	}
	return path.Join("./","data","log",u.Username+".log")
}

func Warn(msg string, fields ...zap.Field) {
	log.Warn(msg, fields...)
}

func Debug(msg string, fields ...zap.Field) {
	log.Debug(msg, fields...)
}

func Info(msg string, fields ...zap.Field) {
	log.Info(msg, fields...)
}

func Panic(msg string, fields ...zap.Field) {
	log.Panic(msg, fields...)
}

func Error(msg string, fields ...zap.Field) {
	log.Error(msg, fields...)
}
