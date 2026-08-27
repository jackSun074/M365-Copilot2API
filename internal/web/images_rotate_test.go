package web

import "testing"

// When an account's upstream refuses image generation with a quota message,
// the gateway marks it image-limited so future requests fail over to another
// account instead of hammering the exhausted one.
func TestMarkImageLimitedBlocksAvailability(t *testing.T) {
	h := newAccountHealth()
	if h.ImageLimited("acc-1") {
		t.Fatal("fresh account should not be image-limited")
	}
	h.MarkImageLimited("acc-1")
	if !h.ImageLimited("acc-1") {
		t.Fatal("MarkImageLimited must flag the account")
	}
	// MarkImageLimited also puts the account into general cooldown so the
	// scheduler skips it.
	if h.Available("acc-1") {
		t.Fatal("image-limited account must be in cooldown / unavailable")
	}
	// An unrelated account is unaffected.
	if h.ImageLimited("acc-2") || !h.Available("acc-2") {
		t.Fatal("unrelated account should remain available")
	}
}

