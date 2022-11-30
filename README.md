# go-set

Provides an implementation of a set; that is, a collection of a comparable type
that contains no duplicate elements. This implementation is not thread safe and
is backed by a map with empty structs as values (map[T]struct{}).
