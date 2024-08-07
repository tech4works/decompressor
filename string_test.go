package decompressor

import (
	"reflect"
	"testing"
)

func TestToString(t *testing.T) {
	for _, tt := range initTestCases() {
		t.Run(tt.name, func(t *testing.T) {
			defer func() {
				r := recover()
				if (r != nil) != tt.wantErr {
					t.Errorf("ToString() should panic = %v, but it did = %v", tt.wantErr, r != nil)
				}
			}()

			if gotRes := ToString(tt.compressType, tt.arg); !reflect.DeepEqual([]byte(gotRes), tt.want) {
				t.Errorf("ToString() = %v, want %v", gotRes, tt.want)
			}
		})
	}
}

func TestToStringWithErr(t *testing.T) {
	for _, tt := range initTestCases() {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ToStringWithErr(tt.compressType, tt.arg)
			if (err != nil) != tt.wantErr {
				t.Errorf("ToStringWithErr() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if err == nil && !reflect.DeepEqual([]byte(got), tt.want) {
				t.Errorf("ToStringWithErr() = %v, want %v", got, tt.want)
			}
		})
	}
}
