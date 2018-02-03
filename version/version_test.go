package version

import (
	"reflect"
	"testing"

	fm "github.com/chentex/go-fm"
)

const (
	vFile      = "fixtures/VERSION.TEST"
	vAlphaFile = "fixtures/VERSION.ALPHA.TEST"
	vBetaFile  = "fixtures/VERSION.BETA.TEST"
)

func TestNewVersioner(t *testing.T) {
	tests := []struct {
		name string
		want Versioner
	}{
		{"1", &Manager{}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewVersioner(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewVersioner() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestManager_BumpVersion(t *testing.T) {
	f := fm.NewFileManager()
	type args struct {
		versionFile string
		bumpType    string
		alpha       bool
		beta        bool
	}
	tests := []struct {
		name        string
		args        args
		wantErr     bool
		wantVersion string
	}{
		{"1", args{vFile, patchBump, true, true}, true, "0.1.0"},
		{"2", args{vFile, "other", false, false}, true, "0.1.0"},
		{"3", args{vFile, patchBump, false, false}, false, "0.1.1"},
		{"4", args{vFile, patchBump, false, false}, false, "0.1.2"},
		{"5", args{vFile, minorBump, false, false}, false, "0.2.0"},
		{"6", args{vFile, patchBump, false, false}, false, "0.2.1"},
		{"7", args{vFile, minorBump, false, false}, false, "0.3.0"},
		{"8", args{vFile, majorBump, false, false}, false, "1.0.0"},
		{"9", args{vFile, minorBump, false, false}, false, "1.1.0"},
		{"10", args{vFile, patchBump, false, false}, false, "1.1.1"},
		{"11", args{vFile, majorBump, false, false}, false, "2.0.0"},
		{"12", args{vFile, patchBump, true, false}, false, "2.0.1-alpha"},
		{"13", args{vFile, patchBump, false, true}, false, "2.0.2-beta"},
		{"14", args{vFile, minorBump, false, false}, false, "2.1.0"},
		{"15", args{vAlphaFile, patchBump, true, false}, false, "0.1.1-alpha"},
		{"16", args{vAlphaFile, minorBump, true, false}, false, "0.2.0-alpha"},
		{"17", args{vAlphaFile, majorBump, false, false}, false, "1.0.0"},
		{"18", args{vBetaFile, patchBump, false, true}, false, "0.1.1-beta"},
		{"19", args{vBetaFile, minorBump, false, true}, false, "0.2.0-beta"},
		{"20", args{vBetaFile, majorBump, false, false}, false, "1.0.0"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &Manager{}
			if err := m.BumpVersion(tt.args.versionFile, tt.args.bumpType, tt.args.alpha, tt.args.beta); (err != nil) != tt.wantErr {
				t.Errorf("Manager.BumpVersion() error = %v, wantErr %v", err, tt.wantErr)
			}
			if !tt.wantErr {
				version, err := f.OpenFile(tt.args.versionFile)
				if err != nil {
					t.Fatal(err)
				}
				if tt.wantVersion != version {
					t.Errorf("Manager.BumpVersion() version = %v, wantVersion %v", version, tt.wantVersion)
				}
			}
		})
	}
	err := f.WriteFile(vFile, []byte("0.1.0"), 0644)
	if err != nil {
		t.Fatal(err)
	}
	err = f.WriteFile(vAlphaFile, []byte("0.1.0-alpha"), 0644)
	if err != nil {
		t.Fatal(err)
	}
	err = f.WriteFile(vBetaFile, []byte("0.1.0-beta"), 0644)
	if err != nil {
		t.Fatal(err)
	}
}
