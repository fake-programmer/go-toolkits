package cfg

import "testing"

func TestSetCfgFile(t *testing.T) {
	type args struct {
		filePath string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
		{name: "1", args: args{"testfile"}, wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := SetCfgFile(tt.args.filePath); (err != nil) != tt.wantErr {
				t.Errorf("SetCfgFile() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
