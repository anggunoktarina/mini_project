package usecase

import (
	"mini_project/model"
	"testing"
)

func TestLoginUser(t *testing.T) {
	type args struct {
		user *model.User
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "Test Login User",
			args: args{
				user: &model.User{
					Name:     "testuser",
					Password: "testpass",
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := LoginUser(tt.args.user); (err != nil) != tt.wantErr {
				t.Errorf("LoginUser() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func LoginUser(user *model.User) error {
	panic("unimplemented")
}

// func TestLoginUser(t *testing.T) {
// 	type args struct {
// 		user *model.User
// 	}
// 	tests := []struct {
// 		name    string
// 		args    args
// 		wantErr bool
// 	}{
// 		// TODO: Add test cases.
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			if err := LoginUser(tt.args.user); (err != nil) != tt.wantErr {
// 				t.Errorf("LoginUser() error = %v, wantErr %v", err, tt.wantErr)
// 			}
// 		})
// 	}
// }

// func LoginUser(invalid, type *invalid type) {
// 	panic("unimplemented")
// }

// func TestCreateUser(t *testing.T) {
// 	type args struct {
// 		user *model.User
// 	}
// 	tests := []struct {
// 		name    string
// 		args    args
// 		wantErr bool
// 	}{
// 		// TODO: Add test cases.
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			if err := TestCreateUser(tt.args.user); (err != nil) != tt.wantErr {
// 				t.Errorf("CreateUser() error = %v, wantErr %v", err, tt.wantErr)
// 			}
// 		})
// 	}
// }
