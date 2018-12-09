package benchmarks

import (
	"errors"
	"fmt"
	"testing"
	"time"

	"github.com/goph/logur"
)

var (
	errExample = errors.New("fail")

	_messages = fakeMessages(1000)

	_tenInts    = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 0}
	_tenStrings = []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j"}
	_tenTimes   = []time.Time{
		time.Unix(0, 0),
		time.Unix(1, 0),
		time.Unix(2, 0),
		time.Unix(3, 0),
		time.Unix(4, 0),
		time.Unix(5, 0),
		time.Unix(6, 0),
		time.Unix(7, 0),
		time.Unix(8, 0),
		time.Unix(9, 0),
	}
	_oneUser = &user{
		Name:      "Jane Doe",
		Email:     "jane@test.com",
		CreatedAt: time.Date(1980, 1, 1, 12, 0, 0, 0, time.UTC),
	}
	_tenUsers = users{
		_oneUser,
		_oneUser,
		_oneUser,
		_oneUser,
		_oneUser,
		_oneUser,
		_oneUser,
		_oneUser,
		_oneUser,
		_oneUser,
	}
)

func fakeMessages(n int) []string {
	messages := make([]string, n)
	for i := range messages {
		messages[i] = fmt.Sprintf("Test logging, but use a somewhat realistic message length. (#%v)", i)
	}
	return messages
}

func getMessage(iter int) string {
	return _messages[iter%1000]
}

type users []*user

type user struct {
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
}

func fakeFields() logur.Fields {
	return logur.Fields{
		"int":     _tenInts[0],
		"ints":    _tenInts,
		"string":  _tenStrings[0],
		"strings": _tenStrings,
		"time":    _tenTimes[0],
		"times":   _tenTimes,
		"user1":   _oneUser,
		"user2":   _oneUser,
		"users":   _tenUsers,
		"error":   errExample,
	}
}

func BenchmarkDisabledWithoutFields(b *testing.B) {
	b.Log("Logging at a disabled level without any structured context.")

	b.Run(logrusPackage, func(b *testing.B) {
		logger := newDisabledLogrus()
		b.ResetTimer()
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				logger.Info(getMessage(0))
			}
		})
	})
}

func BenchmarkDisabledAccumulatedContext(b *testing.B) {
	b.Log("Logging at a disabled level with some accumulated context.")

	b.Run(logrusPackage, func(b *testing.B) {
		logger := newDisabledLogrus().WithFields(fakeFields())
		b.ResetTimer()
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				logger.Info(getMessage(0))
			}
		})
	})
}

func BenchmarkDisabledAddingFields(b *testing.B) {
	b.Log("Logging at a disabled level, adding context at each log site.")

	b.Run(logrusPackage, func(b *testing.B) {
		logger := newDisabledLogrus()
		b.ResetTimer()
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				logger.WithFields(fakeFields()).Info(getMessage(0))
			}
		})
	})
}

func BenchmarkWithoutFields(b *testing.B) {
	b.Log("Logging without any structured context.")

	b.Run(logrusPackage, func(b *testing.B) {
		logger := newLogrus()
		b.ResetTimer()
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				logger.Info(getMessage(0))
			}
		})
	})
}

func BenchmarkAccumulatedContext(b *testing.B) {
	b.Log("Logging with some accumulated context.")

	b.Run(logrusPackage, func(b *testing.B) {
		logger := newLogrus().WithFields(fakeFields())
		b.ResetTimer()
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				logger.Info(getMessage(0))
			}
		})
	})
}

func BenchmarkAddingFields(b *testing.B) {
	b.Log("Logging with additional context at each log site.")

	b.Run(logrusPackage, func(b *testing.B) {
		logger := newLogrus()
		b.ResetTimer()
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				logger.WithFields(fakeFields()).Info(getMessage(0))
			}
		})
	})
}
