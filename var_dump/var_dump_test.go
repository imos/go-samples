package var_dump

import (
  "github.com/imos/go/var_dump"
  "testing"
)

type TestStruct struct {
  Key               int32
  PublicAttribute   []string
  private_attribute *string
  Child             *TestStruct
}

func TestVarDumpExport(t *testing.T) {
  foo := "foo"
  data := TestStruct{
    Key: 12345,
    PublicAttribute: []string{"hoge", "piyo"},
    private_attribute: &foo,
    Child: &TestStruct{
      Key: 23456,
      PublicAttribute: []string{},
      private_attribute: nil,
      Child: nil,
    },
  }
  actual := var_dump.Export(data)
  expected := "" + // For gofmt
    "var_dump.TestStruct{\n" +
    "  Key: int32(12345),\n" +
    "  PublicAttribute: []string{\n" +
    "    string(\"hoge\"),\n" +
    "    string(\"piyo\"),\n" +
    "  },\n" +
    "  private_attribute: &string(\"foo\"),\n" +
    "  Child: &var_dump.TestStruct{\n" +
    "    Key: int32(23456),\n" +
    "    PublicAttribute: []string{},\n" +
    "    private_attribute: (*string)nil,\n" +
    "    Child: (*var_dump.TestStruct)nil,\n" +
    "  },\n" +
    "}"
  if actual != expected {
    t.Errorf("Output should be %#v, but %#v.", expected, actual)
  }
}
