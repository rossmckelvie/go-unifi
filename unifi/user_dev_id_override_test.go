package unifi

import (
	"encoding/json"
	"testing"
)

// UniFi Network 11 sends dev_id_override as a numeric string ("3263");
// earlier versions send a number. Both, plus the empty and null forms,
// must decode.
func TestUserUnmarshalDevIdOverride(t *testing.T) {
	cases := map[string]struct {
		body string
		want int
	}{
		"number":         {`{"mac":"aa:bb:cc:dd:ee:ff","dev_id_override":3263}`, 3263},
		"numeric string": {`{"mac":"aa:bb:cc:dd:ee:ff","dev_id_override":"3263"}`, 3263},
		"empty string":   {`{"mac":"aa:bb:cc:dd:ee:ff","dev_id_override":""}`, 0},
		"null":           {`{"mac":"aa:bb:cc:dd:ee:ff","dev_id_override":null}`, 0},
		"absent":         {`{"mac":"aa:bb:cc:dd:ee:ff"}`, 0},
	}
	for name, tc := range cases {
		t.Run(name, func(t *testing.T) {
			var u User
			if err := json.Unmarshal([]byte(tc.body), &u); err != nil {
				t.Fatalf("unmarshal: %v", err)
			}
			if u.DevIdOverride != tc.want {
				t.Fatalf("DevIdOverride = %d, want %d", u.DevIdOverride, tc.want)
			}
			if u.MAC != "aa:bb:cc:dd:ee:ff" {
				t.Fatalf("MAC = %q, other fields must still decode", u.MAC)
			}
		})
	}
}
