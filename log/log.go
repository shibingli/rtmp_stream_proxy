package log

import (
	"os"
	"path/filepath"
	"rtmp_stream_proxy/utils"
	"strings"

	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var (
	DefaultLogFile = "/tmp/hsp/app.log"
)

type Config struct {
	Filename   string `json:"filename" yaml:"filename"`
	MaxSize    int    `json:"max_size" yaml:"max_size"`
	MaxBackups int    `json:"max_backups" yaml:"max_backups"`
	MaxAge     int    `json:"max_age" yaml:"max_age"`
	Compress   bool   `json:"compress" yaml:"compress"`
}

type Logger struct {
	ZapLogger *zap.Logger
}

func NewLogger(level zapcore.Level, config ...Config) (*Logger, error) {

	encoderConfig := zapcore.EncoderConfig{
		TimeKey:        "T",
		LevelKey:       "L",
		NameKey:        "LOG",
		CallerKey:      "F",
		MessageKey:     "M",
		StacktraceKey:  "S",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.LowercaseLevelEncoder,
		EncodeTime:     zapcore.ISO8601TimeEncoder,
		EncodeDuration: zapcore.SecondsDurationEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder,
		EncodeName:     zapcore.FullNameEncoder,
	}

	var zws zapcore.WriteSyncer

	var conf Config
	if len(config) > 0 && len(strings.TrimSpace(config[0].Filename)) > 0 {
		conf = config[0]

		filename := strings.TrimSpace(conf.Filename)

		logDir := filepath.Dir(filename)

		if !utils.IsDir(logDir) {
			if err := os.Mkdir(logDir, os.ModePerm); nil != err {
				return nil, err
			}
		}

		zws = zapcore.NewMultiWriteSyncer(zapcore.AddSync(os.Stdout), zapcore.AddSync(&lumberjack.Logger{
			Filename:   conf.Filename,
			MaxSize:    conf.MaxSize,
			MaxBackups: conf.MaxBackups,
			MaxAge:     conf.MaxAge,
			Compress:   conf.Compress,
		}))
	} else {
		zws = zapcore.NewMultiWriteSyncer(zapcore.AddSync(os.Stdout))
	}

	l := zap.New(zapcore.NewCore(
		zapcore.NewJSONEncoder(encoderConfig),
		zws,
		zap.NewAtomicLevelAt(level),
	), zap.AddCaller(), zap.Development())

	defer func() {
		_ = l.Sync()
	}()

	return &Logger{
		ZapLogger: l,
	}, nil
}

func GetLogLevel(lv string) (level zapcore.Level) {
	lv = strings.TrimSpace(lv)
	lv = strings.ToLower(lv)

	switch lv {
	case "debug":
		level = zap.DebugLevel
	case "warn":
		level = zap.WarnLevel
	case "error":
		level = zap.ErrorLevel
	case "info":
		level = zap.InfoLevel
	case "fatal":
		level = zap.FatalLevel
	default:
		level = zap.InfoLevel
	}

	return
}

func (l *Logger) IsZapField(arg interface{}) (boo bool) {
	switch arg.(type) {
	case zap.Field:
		boo = true
	}
	return
}

func (l *Logger) Printf(format string, args ...interface{}) {
	l.ZapLogger.Sugar().Infof(format, args...)
}

func (l *Logger) Infof(format string, args ...interface{}) {
	l.ZapLogger.Sugar().Infof(format, args...)
}

func (l *Logger) Debugf(format string, args ...interface{}) {
	l.ZapLogger.Sugar().Debugf(format, args...)
}

func (l *Logger) Warnf(format string, args ...interface{}) {
	l.ZapLogger.Sugar().Warnf(format, args...)
}

func (l *Logger) Errorf(format string, args ...interface{}) {
	l.ZapLogger.Sugar().Errorf(format, args...)
}

func (l *Logger) Fatalf(format string, args ...interface{}) {
	l.ZapLogger.Sugar().Fatalf(format, args...)
}
