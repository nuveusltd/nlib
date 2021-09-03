package nlib

const (
	BashFontColor_RESET  = "\033[0m"
	BashFontColor_RED    = "\033[0;31m"
	BashFontColor_YELLOW = "\033[0;33m"
	BashFontColor_GREEN  = "\033[0;32m"

	NumValidCharList = 85
	GreenTick        = BashFontColor_GREEN + "\u2713" + BashFontColor_RESET
	ValidCharList    = "!#$%&()*+-.0123456789:;<>?@ABCDEFGHIJKLMNOPQRSTUVWXYZ[]_abcdefghijklmnopqrstuvwxyz{|}"
)

/*
Black        0;30     Dark Gray     1;30
Light Red     1;31
Light Green   1;32
Brown/Orange 0;33
Blue         0;34     Light Blue    1;34
Purple       0;35     Light Purple  1;35
Cyan         0;36     Light Cyan    1;36
Light Gray   0;37     White         1;37


*/
/*
1609-09-12 19:02:35 PM +03:00 Sep Sat PDT
--------------- + ------------ + ------------
Type            | Placeholder  |        Value
--------------- + ------------ + ------------
Year            | 2006         | 1609
Year            | 06           | 09
Month           | 01           | 09
Month           | 1            | 9
Month           | Jan          | Sep
Month           | January      | September
Day             | 02           | 12
Day             | 2            | 12
Week day        | Mon          | Sat
Week day        | Monday       | Saturday
Hours           | 03           | 07
Hours           | 3            | 7
Hours           | 15           | 19
Minutes         | 04           | 02
Minutes         | 4            | 2
Seconds         | 05           | 35
Seconds         | 5            | 35
AM or PM        | PM           | PM
Milliseconds     | .000         | .123
Microseconds    | .000000      | .123456
Nanoseconds     | .000000000   | .123456789
Timezone offset | -0700        | +0300
Timezone offset | -07:00       | +03:00
Timezone offset | Z0700        | +0300
Timezone offset | Z07:00       | +03:00
Timezone        | MST          | PDT
--------------- + ------------ + ------------

060102150405

*/
