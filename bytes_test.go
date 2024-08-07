package decompressor

import (
	"reflect"
	"testing"
)

func TestToBytes(t *testing.T) {
	for _, tt := range initTestCases() {
		t.Run(tt.name, func(t *testing.T) {
			defer func() {
				r := recover()
				if (r != nil) != tt.wantErr {
					t.Errorf("ToBytes() should panic = %v, but it did = %v", tt.wantErr, r != nil)
				}
			}()

			if gotRes := ToBytes(tt.compressType, tt.arg); !reflect.DeepEqual(gotRes, tt.want) {
				t.Errorf("ToBytes() = %v, want %v", gotRes, tt.want)
			}
		})
	}
}

func TestToBytesWithErr(t *testing.T) {
	for _, tt := range initTestCases() {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ToBytesWithErr(tt.compressType, tt.arg)
			if (err != nil) != tt.wantErr {
				t.Errorf("ToBytesWithErr() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ToBytesWithErr() = %v, want %v", got, tt.want)
			}
		})
	}
}
