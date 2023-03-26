package v1beta

import "testing"

func TestIsEmailValid(t *testing.T) {
	type args struct {
		email string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// create test cases here. email with domain, email without domain, email with invalid domain, email with invalid characters, etc.
		{
			name: "email with domain",
			args: args{
				email: "devs@jobsika.cm",
			},
			want: true,
		},
		{
			name: "email with sub domain",
			args: args{
				email: "devs@jobs.jobsika.cm",
			},
			want: true,
		},
		{
			name: "email with multiple spaces",
			args: args{
				email: "devs @ jobsika.cm",
			},
			want: false,
		},
		{
			name: "email without domain",
			args: args{
				email: "devs@jobsika",
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsEmailValid(tt.args.email); got != tt.want {
				t.Errorf("IsEmailValid() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsPhoneNumberValid(t *testing.T) {
	type args struct {
		phone string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "valid phone number",
			args: args{
				phone: "+1-123-456-7890",
			},
			want: true,
		},
		{
			name: "valid phone number ((123) 456-7890)",
			args: args{
				phone: "(123) 456-7890",
			},
			want: true,
		},
		{
			name: "valid phone number (123.456.7890)",
			args: args{
				phone: "123.456.7890",
			},
			want: true,
		},
		{
			name: "valid phone number (123 456 7890)",
			args: args{
				phone: "123 456 7890",
			},
			want: true,
		},
		{
			name: "valid phone number (+1234567890)",
			args: args{
				phone: "+1234567890",
			},
			want: true,
		},

		// invalid phone numbers
		{
			name: "invalid phone number (abc-123-4567)",
			args: args{
				phone: "abc-123-4567",
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsPhoneNumberValid(tt.args.phone); got != tt.want {
				t.Errorf("IsPhoneNumberValid() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFormatTags(t *testing.T) {
	type args struct {
		tagsString string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{name: "Valid tags", args: args{tagsString: "go,linux,gcp"}, want: "go,linux,gcp"},
		{name: "tags with duplicate", args: args{tagsString: "go,go,linux,gcp,go"}, want: "go,linux,gcp"},
		{name: "tags with leading/trailing spaces", args: args{tagsString: " go ,linux ,gcp"}, want: "go,linux,gcp"},
		{name: "tags with empty element", args: args{tagsString: " ,go,  ,linux,  ,gcp,  "}, want: "go,linux,gcp"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FormatTags(tt.args.tagsString); got != tt.want {
				t.Errorf("FormatTags() = %v, want %v", got, tt.want)
			}
		})
	}
}
