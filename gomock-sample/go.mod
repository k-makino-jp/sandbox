module example

go 1.14

replace local.packages/printer => ./printer

replace local.packages/filereader => ./filereader

require (
	github.com/golang/mock v1.4.4
	local.packages/filereader v0.0.0-00010101000000-000000000000
	local.packages/printer v0.0.0-00010101000000-000000000000
)
