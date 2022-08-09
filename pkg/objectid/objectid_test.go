package objectid

import "testing"

func TestGetObjectId(t *testing.T) {
	oid := GetObjectId()
	t.Log(oid)

	for i := 0; i < 10; i++ {
		t.Log(GetObjectId())
	}
}
