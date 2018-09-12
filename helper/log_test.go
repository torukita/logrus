package helper

import (
	"testing"
	"github.com/sirupsen/logrus"
)

func TestLogLevel(t *testing.T) {
	DefaultLogLevel = logrus.WarnLevel
	cases := []struct {
		in  string
		out logrus.Level
	}{
		{
			in: "panic",
			out: logrus.PanicLevel,
		},
		{
			in: "error",
			out: logrus.ErrorLevel,
		},
		{
			in: "DeBug",
			out: logrus.DebugLevel,
		},
		{
			in: "hogehoge",
			out: DefaultLogLevel,
		},
	}

	for _, c := range cases {
		got := LogLevel(c.in)
		if got != c.out {
			t.Fatalf("in: %s out: %d got: %d", c.in, c.out, got)
		}
	}
}

