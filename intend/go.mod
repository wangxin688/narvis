module github.com/wangxin688/narvis/intend

go 1.23.0

require (
	github.com/samber/lo v1.47.0
	github.com/wangxin688/narvis/common v0.0.0-00010101000000-000000000000
)

require golang.org/x/text v0.16.0 // indirect

replace github.com/wangxin688/narvis/common => ../common

replace github.com/wangxin688/narvis/intend => ../intend
