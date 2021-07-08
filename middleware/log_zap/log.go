package log_zap

import (
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var sugarLogger *zap.SugaredLogger

func NewEncoderConfig() zapcore.EncoderConfig {
	return zapcore.EncoderConfig{
		// Keys can be anything except the empty string.
		TimeKey:        "T",
		LevelKey:       "L",
		NameKey:        "N",
		CallerKey:      "C",
		MessageKey:     "M",
		StacktraceKey:  "S",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.CapitalLevelEncoder,
		EncodeTime:     TimeEncoder,
		EncodeDuration: zapcore.StringDurationEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder,
	}
}

func TimeEncoder(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendString(t.Format("2006-01-02 15:04:05.000"))
}

func init() {
	var err error

	//logger, err := zap.NewProduction(zap.AddCaller())
	logger, err := zap.NewDevelopment(
		zap.AddCaller(),
		zap.AddCallerSkip(1),

		//zap.AddStacktrace(zapcore.ErrorLevel),
	) // NewExample or NewProduction, or NewDevelopment

	if err != nil {
		panic(err)
	}

	//w := zapcore.AddSync(&lumberjack.Logger{
	//	Filename:   "foo.log",
	//	MaxSize:    500, // megabytes
	//	MaxBackups: 3,
	//	MaxAge:     28, // days
	//})
	//
	//core := zapcore.NewCore(
	//	zapcore.NewConsoleEncoder(NewEncoderConfig()),
	//	zapcore.NewMultiWriteSyncer(zapcore.AddSync(os.Stdout), w),
	//	zap.DebugLevel,
	//)

	//logger = zap.New(core, zap.AddCaller())

	defer func() {
		err = logger.Sync()

		//if err != nil {
		//	panic(err)
		//}
	}()

	sugarLogger = logger.Sugar()
	defer func() {
		err = sugarLogger.Sync()
	}()

	return
}

func Info(template string, args ...interface{}) {
	sugarLogger.Infof(template, args)
}

func Debug(template string, args ...interface{}) {
	sugarLogger.Debugf(template, args)
}

func Warn(template string, args ...interface{}) {
	sugarLogger.Warnf(template, args)
}

func Error(template string, args ...interface{}) {
	sugarLogger.Errorf(template, args)
}
