package request

import "testing"

func TestRequest_KeepAlive_Integration(t *testing.T) {
	requestClient := NewRequest()

	t.Run("Should return true when the request is successful", func(t *testing.T) {
		if !requestClient.KeepAlive("https://www.google.com") {
			t.Fail()
		}
	})

	t.Run("Should return false when the request is unsuccessful", func(t *testing.T) {
		if requestClient.KeepAlive("https://www.wrongdomain.whatever") {
			t.Fail()
		}
	})
}
