package clog

var clog *Logger

// Err starts a new message with error level with err as a field if not nil or
// with info level if err is nil.
//
// You must call Msg on the returned event in order to send the event.
func Err(err error) *Event {
	return clog.Err(err)
}

// Trace starts a new message with trace level.
//
// You must call Msg on the returned event in order to send the event.
func Trace() *Event {
	return clog.Trace()
}

// Debug starts a new message with debug level.
//
// You must call Msg on the returned event in order to send the event.
func Debug() *Event {
	return clog.Debug()
}

// Random starts a new message with random level.
//
// You must call Msg on the returned event in order to send the event.
func Random(seed int64) *Event {
	return clog.Random(seed)
}

// Info starts a new message with info level.
//
// You must call Msg on the returned event in order to send the event.
func Info() *Event {
	return clog.Info()
}

// Warn starts a new message with warn level.
//
// You must call Msg on the returned event in order to send the event.
func Warn() *Event {
	return clog.Warn()
}

// Error starts a new message with error level.
//
// You must call Msg on the returned event in order to send the event.
func Error() *Event {
	return clog.Error()
}

// Fatal starts a new message with fatal level. The os.Exit(1) function
// is called by the Msg method.
//
// You must call Msg on the returned event in order to send the event.
func Fatal() *Event {
	return clog.Fatal()
}

// Panic starts a new message with panic level. The message is also sent
// to the panic function.
//
// You must call Msg on the returned event in order to send the event.
func Panic() *Event {
	return clog.Panic()
}

// Log starts a new message with no level. Setting clog.GlobalLevel to
// clog.Disabled will still disable events produced by this method.
//
// You must call Msg on the returned event in order to send the event.
func Log() *Event {
	return clog.Log()
}

// Print sends a log event using debug level and no extra field.
// Arguments are handled in the manner of fmt.Print.
func Print(v ...interface{}) {
	clog.Print(v...)
}

// Printf sends a log event using debug level and no extra field.
// Arguments are handled in the manner of fmt.Printf.
func Printf(format string, v ...interface{}) {
	clog.Printf(format, v...)
}

func GetLevel() Level {
	return clog.level
}

// ???????????? ????????????Logger ?????????
func CopyDefault() *Logger {
	l := &Logger{
		w:      clog.w,
		level:  clog.level,
		random: clog.random,
	}
	if len(clog.preStr) > 0 {
		l.preStr = make([]byte, len(clog.preStr))
		copy(l.preStr, clog.preStr)
	}
	if len(clog.preHook) > 0 {
		l.preHook = make([]Hook, len(clog.preHook))
		copy(l.preHook, clog.preHook)
	}
	if len(clog.hooks) > 0 {
		l.hooks = make([]Hook, len(clog.hooks))
		copy(l.hooks, clog.hooks)
	}

	return l
}
