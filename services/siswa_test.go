package services

import "testing"

func Test_score(t *testing.T) {
	type args struct {
		nilai int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases. 
		//untuk unit test hanya perlu mengubah part ini saja.
		{
			name: "<= 40 == c", //untuk nama ini bisa diisi apa aja sesuai kasus yang ingin dites.
			args: args{
				nilai: 16, //masukan data yang ingin dicoba
			},
			want: "C", //hasil output
		}, //jangan lupa untuk menambahkan koma diantara kasus tes.
		{
			name: ">40 dan <= 70 == b",
			args: args{
				nilai: 55,
			},
			want: "B",
		},
		{
			name: "> 70 == a",
			args: args{
				nilai: 97,
			},
			want: "A",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := score(tt.args.nilai); got != tt.want {
				t.Errorf("score() = %v, want %v", got, tt.want)
			}
		})
	}
}
