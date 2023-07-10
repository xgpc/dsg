package etcd

import "testing"

func Test_parseServiceKey(t *testing.T) {
	type args struct {
		key string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		want1   string
		want2   int
		wantErr bool
	}{
		// TODO: Add test cases.
		{name: "ok", args: args{key: "api/192.168.0.1/8082"}, want: "api", want1: "192.168.0.1", want2: 8082, wantErr: false},
		{name: "不ok", args: args{key: "api/192.168.0.1/8082"}, want: "api", want1: "192.168.0.1", want2: 8081, wantErr: true},
		{name: "参数过多", args: args{key: "xxx/api/192.168.0.1/8082"}, want: "api", want1: "192.168.0.1", want2: 8081, wantErr: true},
		{name: "参数过少", args: args{key: "/192.168.0.1/8082"}, want: "api", want1: "192.168.0.1", want2: 8081, wantErr: true},
		{name: "参数不匹配", args: args{key: "api2/192.168.0.1/8082"}, want: "api", want1: "192.168.0.1", want2: 8081, wantErr: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1, got2, err := parseServiceKey(tt.args.key)
			if (err != nil) != tt.wantErr {
				t.Errorf("parseServiceKey() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("parseServiceKey() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("parseServiceKey() got1 = %v, want %v", got1, tt.want1)
			}
			if got2 != tt.want2 {
				t.Errorf("parseServiceKey() got2 = %v, want %v", got2, tt.want2)
			}
		})
	}
}
