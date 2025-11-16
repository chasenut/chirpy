package auth

import "testing"

func TestCreateHash(t *testing.T) {
	cases := []struct{
		input		string
		expected	string
	} {
		{
			input: "password",
			expected: "$argon2id$v=19$m=65536,t=1,p=12$YN/lYQmrNSwBzJ4fno7iXQ$It5vwQ6DQrqEXm+ajDf8/nyNN63ImjPV5ugjuBVnBas",
		},
	}
	for _, c := range cases {
		_, err := HashPassword(c.input)
		if err != nil {
			t.Errorf("HashPassword error: %s", err)
			t.Fail()
		}
//		if actual != c.expected {
//			t.Errorf("Got: %s\n Expected: %s", actual, c.expected)
//			t.Fail()
//		}
	}
}

func TestCheckPasswordHash(t *testing.T) {

}
