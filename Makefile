all: nullserver/nullserver simplecache

diff: main.strings null.strings
	/usr/bin/diff \
		main.strings null.strings |\
		grep '^<' |\
		grep -v '^< http' |\
		grep -v '^< time' |\
		grep -v '^< net/http' |\
		grep -v '^< bufio\.' |\
		grep -v '^< type\.\.'

run: all
	./simplecache

nullserver/nullserver: nullserver/main.go
	go build -o nullserver/nullserver nullserver/main.go

simplecache: main.go
	go build

clean:
	rm -f simplecache nullserver/nullserver null.strings main.strings

null.strings: nullserver/nullserver
	strings nullserver/nullserver > null.strings

main.strings: simplecache
	strings simplecache > main.strings
