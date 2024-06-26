package main

const (
	OFFSET_MODE_QUIET = 0x08
	BIT_MODE_QUIET = 6

	OFFSET_MODE_GAMING = 0x0C
	BIT_MODE_GAMING = 4

	OFFSET_MODE_DEEP_CONTROL = 0x0D
	BIT_MODE_DEEP_CONTROL = 7

	OFFSET_MODE_AUTO = 0x0D
	BIT_MODE_AUTO = 0

	OFFSET_MODE_FIXED = 0x06
	BIT_MODE_FIXED = 4

	OFFSET_FAN0_RPM = 0xFC
	OFFSET_FAN1_RPM = 0xFE

	OFFSET_FAN0_PERCENT = 0xB3
	OFFSET_FAN1_PERCENT = 0xB4

	OFFSET_FAN0_TARGET_PERCENT = 0xB0
	OFFSET_FAN1_TARGET_PERCENT = 0xB1

	SPEED_PERCENT_FACTOR = 1 / 2.29
)
