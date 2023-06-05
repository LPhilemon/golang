func equal(x, y reflect.Value, seen map[comparison]bool) bool {
	if !x.IsValid() || !y.IsValid() {
	return x.IsValid() == y.IsValid()
	}
	if x.Type() != y.Type() {
	return false
	}
	// ...cycle check omitted (shown later)...
	switch x.Kind() {
	case reflect.Bool:
	return x.Bool() == y.Bool()
	case reflect.String:
	return x.String() == y.String()
	// ...numeric cases omitted for brevity...
	case reflect.Chan, reflect.UnsafePointer, reflect.Func:
	return x.Pointer() == y.Pointer()
	case reflect.Ptr, reflect.Interface:
	return equal(x.Elem(), y.Elem(), seen)
	case reflect.Array, reflect.Slice:
	if x.Len() != y.Len() {
	return false
	}
	for i := 0; i < x.Len(); i++ {
	if !equal(x.Index(i), y.Index(i), seen) {
	return false
	}
	}
	return true

	// ...struct and map cases omitted for brevity...
}
panic("unreachable")
}