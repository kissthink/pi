package db

var user = []byte("user")
var device = []byte("device")
var general = []byte("general")

var buckets = [][]byte{
	user,
	device,
	general,
}
