package plistparser

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"log"
	"reflect"
	"sort"
	"strconv"
)

/// Writes out values. Only accepts primitives, []interface{} and map[string]interface{}
func Write(val interface{}, writer io.Writer) (err error) {
	w := bufio.NewWriter(writer)
	err = write(val, w, 0)
	w.WriteRune('\n')
	w.Flush()
	return
}

func stringNeedsQuoting(str string) bool {
	if str == "" {
		return true
	}
	for i, c := range str {
		if !isIdent(c, i) {
			return true

		}
	}
	return false
}

// Quotes the string if needed
func quoteStringIfNeeded(str string) string {
	if stringNeedsQuoting(str) {
		return strconv.Quote(str)
	}
	return str
}

var stringType = reflect.TypeOf("")

func write(val interface{}, writer *bufio.Writer, indent int) (err error) {
	writeIndent := func() {
		for i := 0; i < indent; i++ {
			writer.WriteString("  ")
		}
	}

	switch val := val.(type) {
	case string:
		if _, err = writer.WriteString(quoteStringIfNeeded(val)); err != nil {
			return
		}
		return
	case []byte:
		if _, err = writer.WriteRune('<'); err != nil {
			return
		}

		for i, v := range val {
			if i != 0 {
				if _, err = writer.WriteRune(' '); err != nil {
					return
				}
			}

			if _, err = writer.WriteString(strconv.FormatInt(int64(v), 16)); err != nil {
				return
			}
		}

		if _, err = writer.WriteRune('>'); err != nil {
			return
		}
		return
	}

	reflectVal := reflect.ValueOf(val)

	t := reflectVal.Type()
	switch t.Kind() {
	case reflect.Slice:
		if _, err = writer.WriteString("(\n"); err != nil {
			return
		}

		indent++
		l := reflectVal.Len()
		for i := 0; i < l; i++ {
			writeIndent()
			write(reflectVal.Index(i).Interface(), writer, indent)
			writer.WriteString(",\n")
		}

		indent--

		writeIndent()
		if _, err = writer.WriteString(")"); err != nil {
			return
		}

		return
	case reflect.Map:
		keyType := t.Key()
		if !keyType.ConvertibleTo(stringType) {
			log.Panicln("Can't have key that isn't convertable to a string type")
		}

		if _, err = writer.WriteString("{\n"); err != nil {
			return
		}

		var keys []string
		for _, k := range reflectVal.MapKeys() {
			keys = append(keys, k.Convert(stringType).Interface().(string))
		}

		sort.Strings(keys)

		indent++
		for _, k := range keys {
			v := reflectVal.MapIndex(reflect.ValueOf(k).Convert(keyType)).Interface()
			writeIndent()
			writer.WriteString(quoteStringIfNeeded(k))
			writer.WriteString(" = ")
			write(v, writer, indent)
			writer.WriteString(";\n")
		}
		indent--

		writeIndent()
		if _, err = writer.WriteString("}"); err != nil {
			return
		}
	case reflect.String:
		if _, err = writer.WriteString(quoteStringIfNeeded(reflectVal.Convert(stringType).Interface().(string))); err != nil {
			return
		}

		return
	default:
		err = fmt.Errorf("Could not serialize type of value %+v", val)
		return
	}
	return

}

func WriteString(val interface{}) (str string, err error) {
	buff := bytes.NewBufferString("")

	if err = Write(val, buff); err != nil {
		return
	}

	str = buff.String()
	return
}
