package cassandra

import "regexp"

var hidePasswordRegexp = regexp.MustCompile(`^(cassandra://[^/:]+):[^:@]+@`)
