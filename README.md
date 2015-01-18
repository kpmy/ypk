halt/assert
===
Original idea from Oberon/Component Pascal

Usage:
===

assert.For(boolean condition, error code, error messages) - checks condition, panic if false

halt.As(error code, error messages) - unconditional panic with messages

assert.For(err == nil, 20, "ooops")
halt.As(100, "invariant violated")

