module main

go 1.14

replace local.packages/printer => ./printer

replace local.packages/filereader => ./filereader

require (
	local.packages/filereader v0.0.0-00010101000000-000000000000
	local.packages/printer v0.0.0-00010101000000-000000000000
)
