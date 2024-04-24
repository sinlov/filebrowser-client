package file_browser_log

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestShowLogLineNo(t *testing.T) {
	// mock ShowLogLineNo
	type args struct {
		format string
		v      []interface{}
	}
	tests := []struct {
		name     string
		openLine bool
		args     args
	}{
		{
			name:     "test verbose open line",
			openLine: true,
			args: args{
				format: "test verbose",
				v:      nil,
			},
		},
		{
			name:     "test verbose close line",
			openLine: false,
			args: args{
				format: "test verbose",
				v:      nil,
			},
		},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {

			// do ShowLogLineNo
			ShowLogLineNo(tc.openLine)

			// verify ShowLogLineNo
			Verbosef(tc.args.format, tc.args.v...)
		})
	}
}

func TestSetLogLineDeep(t *testing.T) {
	// mock SetLogLineDeep
	tests := []struct {
		name string
		deep uint
		msg  string
	}{
		{
			name: "sample",
			deep: 3,
			msg:  "sample",
		},
		{
			name: "zero",
			deep: 0,
			msg:  "zero",
		},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {

			// do SetLogLineDeep
			ShowLogLineNo(true)
			SetLogLineDeep(tc.deep)
			Info(tc.msg)
		})
	}
}

func TestDebug(t *testing.T) {
	// mock Debug
	type args struct {
		msg string
	}
	tests := []struct {
		name        string
		args        args
		isOpenDebug bool
	}{
		{
			name: "sample",
			args: args{
				msg: "sample",
			},
			isOpenDebug: true,
		},
		{
			name: "close debug",
			args: args{
				msg: "sample",
			},
			isOpenDebug: false,
		},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {

			// do Debug
			if tc.isOpenDebug {
				OpenDebug()
			}
			Debug(tc.args.msg)
		})
	}
}

func TestDebugf(t *testing.T) {
	type args struct {
		format string
		v      []interface{}
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "test debug",
			args: args{
				format: "test debug",
				v:      nil,
			},
		},
	}
	OpenDebug()
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Debugf(tt.args.format, tt.args.v...)
		})
	}
}

func TestDebugJson(t *testing.T) {
	type args struct {
		d interface{}
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "test debug json",
			args: args{
				d: struct {
					Name string
					Age  int
				}{
					Name: "foo",
					Age:  18,
				},
			},
		},
	}
	OpenDebug()
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			OpenDebug()
			DebugJson(tt.args.d)
		})
	}
}

func TestDebugJsonf(t *testing.T) {
	type args struct {
		d      interface{}
		format string
		v      []interface{}
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "test debug jsonf",
			args: args{
				d: struct {
					Name string
					Age  int
				}{
					Name: "foo",
					Age:  18,
				},
				format: "test debug jsonf",
				v:      nil,
			},
		},
	}
	OpenDebug()
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			DebugJsonf(tt.args.d, tt.args.format, tt.args.v...)
		})
	}
}

func TestVerbosef(t *testing.T) {
	type args struct {
		format string
		v      []interface{}
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "test verbose",
			args: args{
				format: "test verbose",
				v:      nil,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Verbosef(tt.args.format, tt.args.v...)
		})
	}
}

func TestVerboseJson(t *testing.T) {
	type args struct {
		d interface{}
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "test verbose json",
			args: args{
				d: struct {
					Name string
					Age  int
				}{
					Name: "foo",
					Age:  18,
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ShowLogLineNo(true)
			VerboseJson(tt.args.d)
		})
	}
}

func TestVerboseJsonf(t *testing.T) {
	type args struct {
		d      interface{}
		format string
		v      []interface{}
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "test verbose jsonf",
			args: args{
				d: struct {
					Name string
					Age  int
				}{
					Name: "foo",
					Age:  18,
				},
				format: "test verbose jsonf",
				v:      nil,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			VerboseJsonf(tt.args.d, tt.args.format, tt.args.v...)
		})
	}
}

func TestInfo(t *testing.T) {
	// mock Debug
	type args struct {
		msg string
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "sample",
			args: args{
				msg: "sample",
			},
		},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {

			// do Debug
			Info(tc.args.msg)
		})
	}
}

func TestInfof(t *testing.T) {
	type args struct {
		format string
		v      []interface{}
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "test info",
			args: args{
				format: "test info",
				v:      nil,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Infof(tt.args.format, tt.args.v...)
		})
	}
}

func TestWarnf(t *testing.T) {
	type args struct {
		format string
		v      []interface{}
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "test warn",
			args: args{
				format: "test warn",
				v:      nil,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ShowLogLineNo(false)
			Warnf(tt.args.format, tt.args.v...)
		})
	}
}

func TestError(t *testing.T) {
	type args struct {
		err error
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "test error",
			args: args{
				err: fmt.Errorf("new test error"),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ShowLogLineNo(true)
			Error(tt.args.err)
		})
	}
}

func TestErrorf(t *testing.T) {
	type args struct {
		err    error
		format string
		v      []interface{}
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "test error",
			args: args{
				err:    fmt.Errorf("new test error"),
				format: "test error",
				v:      nil,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Errorf(tt.args.err, tt.args.format, tt.args.v...)
		})
	}
}

func TestPanicf(t *testing.T) {
	// mock TestPanicName

	if !assert.Panics(t, func() {
		// do TestPanicName
		Panicf("this is Panicf: %s", "panic desc")
	}) {
		// verify TestPanicName
		t.Fatalf("TestPanicName should panic")
	}
}
