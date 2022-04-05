package logger

type Logger struct{}

func New() (*Logger, error) {
	logger := Logger{}

	err := initZap()
	if err != nil {
		return nil, err
	}

	return &logger, err
}

func (l Logger) Panic(msg string) {
	logger.Panic(msg)
}
func (l Logger) Error(msg string) {
	logger.Panic(msg)
}
