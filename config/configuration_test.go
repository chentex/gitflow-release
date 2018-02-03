package config

import (
	"os"
	"reflect"
	"testing"
)

func TestNewConfigure(t *testing.T) {
	tests := []struct {
		name string
		want Configure
	}{
		{"1", &Configurator{}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewConfigure(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewConfigure() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestConfigurator_InitConfig(t *testing.T) {
	type args struct {
		p Params
	}
	tests := []struct {
		name    string
		c       Configure
		args    args
		wantErr bool
	}{
		{"1", NewConfigure(), args{Params{Force: "false", CfgFile: ".config", InitialVersion: "0.1.0", VersionFile: "VERSION"}}, false},
		{"2", NewConfigure(), args{Params{Force: "false", CfgFile: ".config", InitialVersion: "0.1.0", VersionFile: "VERSION"}}, true},
		{"3", NewConfigure(), args{Params{Force: "true", CfgFile: ".config", InitialVersion: "0.1.0", VersionFile: "VERSION"}}, false},
		{"4", NewConfigure(), args{Params{Force: "true", CfgFile: ".config", InitialVersion: "0.1.0", VersionFile: "VERSION2"}}, false},
		{"5", NewConfigure(), args{Params{Force: "notbool", CfgFile: ".config", InitialVersion: "0.1.0", VersionFile: "VERSION2"}}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := tt.c
			if err := c.InitConfig(tt.args.p); (err != nil) != tt.wantErr {
				t.Errorf("Configurator.InitConfig() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
	err := os.Remove(".config")
	if err != nil {
		t.Fatalf("Cleaning files: %s", err)
	}
	err = os.Remove("VERSION")
	if err != nil {
		t.Fatalf("Cleaning files: %s", err)
	}
	err = os.Remove("VERSION2")
	if err != nil {
		t.Fatalf("Cleaning files: %s", err)
	}
}
