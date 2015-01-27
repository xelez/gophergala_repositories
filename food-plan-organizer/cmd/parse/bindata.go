package main

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
	"strings"
)

func bindata_read(data []byte, name string) ([]byte, error) {
	gz, err := gzip.NewReader(bytes.NewBuffer(data))
	if err != nil {
		return nil, fmt.Errorf("Read %q: %v", name, err)
	}

	var buf bytes.Buffer
	_, err = io.Copy(&buf, gz)
	gz.Close()

	if err != nil {
		return nil, fmt.Errorf("Read %q: %v", name, err)
	}

	return buf.Bytes(), nil
}

func migrations_01_create_foods_table_sql() ([]byte, error) {
	return bindata_read([]byte{
		0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x00, 0xff, 0x74, 0x92,
		0xc1, 0x4e, 0xc3, 0x30, 0x0c, 0x86, 0xef, 0x79, 0x8a, 0x1c, 0x41, 0xb0,
		0x27, 0xd8, 0x69, 0xd0, 0x1e, 0x90, 0xd0, 0x86, 0xa6, 0x21, 0x71, 0xab,
		0x42, 0xea, 0x76, 0x96, 0x68, 0x5c, 0x79, 0x46, 0x8c, 0xb7, 0xa7, 0x61,
		0x6a, 0xe3, 0x74, 0x5d, 0x6f, 0xcd, 0xf7, 0xfd, 0x8a, 0xf2, 0xdb, 0xab,
		0x95, 0x7d, 0xe8, 0xb0, 0x65, 0x27, 0x60, 0xdf, 0x7b, 0x63, 0x9e, 0xf7,
		0xe5, 0xe6, 0x50, 0xda, 0xc3, 0xe6, 0xe9, 0xb5, 0xb4, 0x0d, 0x51, 0x7d,
		0xb2, 0x77, 0xc6, 0x5a, 0xac, 0xed, 0xf5, 0x87, 0x41, 0xa0, 0x05, 0x7e,
		0x1c, 0x78, 0x34, 0xab, 0x96, 0xe9, 0xbb, 0xaf, 0x92, 0xaa, 0x78, 0x70,
		0x1d, 0x5c, 0xe5, 0x05, 0xce, 0x12, 0xe1, 0xe9, 0x48, 0x2c, 0xd5, 0x4c,
		0x19, 0xa1, 0xa7, 0xae, 0xa3, 0x30, 0xa3, 0x53, 0xd2, 0x23, 0x04, 0xc1,
		0x06, 0xbd, 0x12, 0x46, 0x18, 0x50, 0x98, 0x5a, 0x08, 0x55, 0xe3, 0xbc,
		0x10, 0x5f, 0x20, 0x83, 0xfb, 0x8a, 0xb0, 0x67, 0x12, 0xc0, 0x9c, 0x4d,
		0xb0, 0x71, 0x92, 0x03, 0x0d, 0xbd, 0xe3, 0x4f, 0x0a, 0xc7, 0xdf, 0x3a,
		0x76, 0x36, 0x6a, 0x11, 0x9a, 0xfb, 0xf5, 0xd4, 0xdf, 0xcb, 0xb6, 0x28,
		0x3f, 0x2e, 0xfd, 0x0d, 0x85, 0x9c, 0xed, 0x6e, 0x3b, 0x96, 0x89, 0xf5,
		0xa0, 0x2d, 0x58, 0x59, 0x83, 0x3a, 0x90, 0x81, 0xe5, 0xec, 0xff, 0xdb,
		0x53, 0x22, 0xfe, 0x2e, 0x8b, 0xaa, 0xe9, 0xa4, 0xa7, 0xc3, 0xe5, 0x90,
		0x9e, 0x40, 0x4a, 0xa9, 0xd3, 0x1b, 0x77, 0xcd, 0x66, 0xa3, 0x2e, 0xcc,
		0x49, 0xac, 0x4d, 0xaf, 0x61, 0x41, 0x3f, 0xc1, 0x14, 0xfb, 0xdd, 0x9b,
		0xde, 0xc2, 0xb5, 0xf9, 0x0b, 0x00, 0x00, 0xff, 0xff, 0x5b, 0xe6, 0xbf,
		0xde, 0xaa, 0x02, 0x00, 0x00,
	},
		"migrations/01-create-foods-table.sql",
	)
}

func migrations_02_create_food_groups_table_sql() ([]byte, error) {
	return bindata_read([]byte{
		0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x00, 0xff, 0xd2, 0xd5,
		0x55, 0xd0, 0xce, 0xcd, 0x4c, 0x2f, 0x4a, 0x2c, 0x49, 0x55, 0x08, 0x2d,
		0xe0, 0xe2, 0x72, 0x0e, 0x72, 0x75, 0x0c, 0x71, 0x55, 0x08, 0x71, 0x74,
		0xf2, 0x71, 0x55, 0x48, 0xcb, 0xcf, 0x4f, 0x89, 0x4f, 0x2f, 0xca, 0x2f,
		0x2d, 0x28, 0x56, 0xd0, 0xe0, 0x52, 0x50, 0xc8, 0x4c, 0x51, 0x00, 0x12,
		0x79, 0x25, 0xa9, 0xe9, 0xa9, 0x45, 0x3a, 0x40, 0x7e, 0x5e, 0x62, 0x6e,
		0xaa, 0x42, 0x49, 0x6a, 0x45, 0x09, 0x97, 0xa6, 0x35, 0x4c, 0xa7, 0xa7,
		0x9f, 0x8b, 0x6b, 0x04, 0xb2, 0xce, 0xf8, 0xcc, 0x94, 0x0a, 0x05, 0x7f,
		0x3f, 0x54, 0xc3, 0x32, 0x53, 0x80, 0x3a, 0xb8, 0x90, 0x2d, 0x77, 0xc9,
		0x2f, 0xcf, 0xe3, 0x72, 0x09, 0xf2, 0x0f, 0xc0, 0xb4, 0xdb, 0x9a, 0x0b,
		0x10, 0x00, 0x00, 0xff, 0xff, 0xd2, 0x68, 0x9c, 0x58, 0xa6, 0x00, 0x00,
		0x00,
	},
		"migrations/02-create-food-groups-table.sql",
	)
}

func migrations_03_create_languages_table_sql() ([]byte, error) {
	return bindata_read([]byte{
		0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x00, 0xff, 0xd2, 0xd5,
		0x55, 0xd0, 0xce, 0xcd, 0x4c, 0x2f, 0x4a, 0x2c, 0x49, 0x55, 0x08, 0x2d,
		0xe0, 0xe2, 0x72, 0x0e, 0x72, 0x75, 0x0c, 0x71, 0x55, 0x08, 0x71, 0x74,
		0xf2, 0x71, 0x55, 0xc8, 0x49, 0xcc, 0x4b, 0x2f, 0x4d, 0x4c, 0x4f, 0x2d,
		0x56, 0xd0, 0xe0, 0x52, 0x50, 0xc8, 0x2b, 0x2d, 0x29, 0xca, 0x4c, 0xcd,
		0x2b, 0x89, 0xcf, 0x4c, 0x51, 0x28, 0x49, 0xad, 0x28, 0xd1, 0x01, 0x8a,
		0xa5, 0x25, 0x26, 0x97, 0xe4, 0x17, 0xc5, 0x27, 0xe7, 0xa7, 0xa4, 0x82,
		0xc5, 0xb8, 0x34, 0xad, 0x61, 0x46, 0x78, 0xfa, 0xb9, 0xb8, 0x46, 0x20,
		0x8c, 0x88, 0x47, 0xd2, 0x5e, 0xa1, 0xe0, 0xef, 0x87, 0x6c, 0x38, 0x92,
		0x14, 0x6e, 0xfd, 0x48, 0x56, 0x61, 0x31, 0x02, 0x49, 0x16, 0x68, 0x04,
		0x17, 0xb2, 0xb7, 0x5c, 0xf2, 0xcb, 0xf3, 0xb8, 0x5c, 0x82, 0xfc, 0x03,
		0xd0, 0x7d, 0x65, 0xcd, 0x05, 0x08, 0x00, 0x00, 0xff, 0xff, 0xc0, 0x18,
		0x4a, 0xec, 0xfe, 0x00, 0x00, 0x00,
	},
		"migrations/03-create-languages-table.sql",
	)
}

func migrations_04_create_language_descriptions_table_sql() ([]byte, error) {
	return bindata_read([]byte{
		0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x00, 0xff, 0xd2, 0xd5,
		0x55, 0xd0, 0xce, 0xcd, 0x4c, 0x2f, 0x4a, 0x2c, 0x49, 0x55, 0x08, 0x2d,
		0xe0, 0xe2, 0x72, 0x0e, 0x72, 0x75, 0x0c, 0x71, 0x55, 0x08, 0x71, 0x74,
		0xf2, 0x71, 0x55, 0xc8, 0x49, 0xcc, 0x4b, 0x2f, 0x4d, 0x4c, 0x4f, 0x8d,
		0x4f, 0x49, 0x2d, 0x4e, 0x2e, 0xca, 0x2c, 0x28, 0xc9, 0xcc, 0xcf, 0x2b,
		0x56, 0xd0, 0xe0, 0x52, 0x50, 0x48, 0x4b, 0x4c, 0x2e, 0xc9, 0x2f, 0x8a,
		0x4f, 0xce, 0x4f, 0x49, 0x55, 0x28, 0x49, 0xad, 0x28, 0xd1, 0x01, 0x8a,
		0x21, 0x29, 0x02, 0x8b, 0x71, 0x69, 0x5a, 0xc3, 0x8c, 0xf3, 0xf4, 0x73,
		0x71, 0x8d, 0xc0, 0x6e, 0x5c, 0x7c, 0x66, 0x4a, 0x85, 0x82, 0xbf, 0x1f,
		0x2e, 0xbb, 0x90, 0x2c, 0x02, 0x1a, 0xc7, 0x85, 0xec, 0x5c, 0x97, 0xfc,
		0xf2, 0x3c, 0x2e, 0x97, 0x20, 0xff, 0x00, 0x7c, 0xae, 0xb5, 0xe6, 0x02,
		0x04, 0x00, 0x00, 0xff, 0xff, 0xa8, 0xec, 0x72, 0x88, 0xe2, 0x00, 0x00,
		0x00,
	},
		"migrations/04-create-language-descriptions-table.sql",
	)
}

func migrations_05_create_nutrients_table_sql() ([]byte, error) {
	return bindata_read([]byte{
		0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x00, 0xff, 0x74, 0x91,
		0xc1, 0x6a, 0x84, 0x30, 0x10, 0x86, 0xef, 0x79, 0x8a, 0x39, 0xb6, 0xb4,
		0xfb, 0x04, 0x7b, 0xda, 0xd6, 0x50, 0x16, 0x44, 0x8b, 0x58, 0xe8, 0x2d,
		0x58, 0x32, 0x4a, 0x40, 0x33, 0x32, 0x26, 0xd5, 0xc7, 0x6f, 0x2c, 0xba,
		0x06, 0xd1, 0x39, 0x85, 0x7c, 0xf3, 0x31, 0xfc, 0x33, 0x97, 0x0b, 0xbc,
		0x74, 0xa6, 0xe1, 0xca, 0x21, 0x7c, 0xf5, 0x42, 0xbc, 0x17, 0xf2, 0x56,
		0x4a, 0x28, 0x6f, 0x6f, 0xa9, 0x04, 0xeb, 0x1d, 0x1b, 0xb4, 0x6e, 0x80,
		0x27, 0x01, 0x50, 0x13, 0x69, 0x65, 0x34, 0x44, 0x65, 0xac, 0xc3, 0x06,
		0xf9, 0x35, 0xd0, 0xb5, 0x37, 0xea, 0x38, 0xa2, 0xbf, 0x55, 0xeb, 0xf1,
		0x9f, 0x86, 0x41, 0xe9, 0x8c, 0x3a, 0x63, 0x61, 0x57, 0x0f, 0x54, 0x4d,
		0x67, 0x48, 0x63, 0xc3, 0x88, 0x83, 0xa2, 0x5a, 0xd5, 0xe1, 0xa1, 0xa9,
		0x83, 0x7b, 0x56, 0xca, 0x0f, 0x59, 0xcc, 0xb4, 0xa5, 0x11, 0x59, 0x21,
		0x33, 0xb1, 0xfa, 0x21, 0x6f, 0xf5, 0x26, 0xfa, 0xbe, 0x3f, 0x42, 0xe2,
		0xf9, 0xba, 0x66, 0xbf, 0x67, 0x89, 0xfc, 0xde, 0xb2, 0xab, 0x25, 0xf7,
		0x04, 0x79, 0x16, 0x6f, 0x64, 0xf9, 0x3e, 0xf7, 0xa2, 0x8d, 0xec, 0xdd,
		0x08, 0x05, 0x5f, 0xc4, 0x47, 0x48, 0x68, 0xb4, 0x22, 0x29, 0xf2, 0xcf,
		0xfd, 0x0d, 0xae, 0xe2, 0x2f, 0x00, 0x00, 0xff, 0xff, 0x9b, 0xa5, 0x01,
		0x7c, 0xac, 0x01, 0x00, 0x00,
	},
		"migrations/05-create-nutrients-table.sql",
	)
}

func migrations_06_create_nutrient_definitions_table_sql() ([]byte, error) {
	return bindata_read([]byte{
		0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x00, 0xff, 0x7c, 0x8f,
		0x41, 0xcb, 0x82, 0x40, 0x10, 0x86, 0xef, 0xf3, 0x2b, 0xe6, 0xf8, 0x7d,
		0x94, 0xbf, 0xc0, 0x93, 0xe5, 0x1e, 0x82, 0xd0, 0x90, 0x0d, 0xbc, 0xc9,
		0xa2, 0x93, 0x0c, 0xe8, 0x28, 0xbb, 0x2b, 0xf5, 0xf3, 0x5b, 0x0f, 0x96,
		0x45, 0x35, 0xc7, 0xf7, 0x1d, 0x1e, 0x9e, 0x37, 0x8a, 0x70, 0xd3, 0x73,
		0x6b, 0x8d, 0x27, 0x3c, 0x8f, 0x00, 0xfb, 0x42, 0x25, 0x5a, 0xa1, 0x4e,
		0x76, 0x47, 0x85, 0x32, 0x79, 0xcb, 0x24, 0xbe, 0x6a, 0xe8, 0xc2, 0xc2,
		0x9e, 0x07, 0x71, 0xf8, 0x07, 0xf8, 0x2c, 0xb8, 0xc1, 0x70, 0x2c, 0x9e,
		0x5a, 0xb2, 0xdb, 0xd0, 0x4c, 0xe1, 0xcd, 0xe1, 0xe3, 0xb4, 0x2a, 0xf5,
		0x1c, 0x7b, 0xd3, 0x8a, 0xe9, 0xe9, 0x3d, 0x6e, 0xc8, 0xd5, 0x96, 0xc7,
		0x19, 0xfc, 0x1a, 0xd7, 0xdc, 0x9b, 0xae, 0x1a, 0x3b, 0x53, 0x93, 0x5b,
		0xf0, 0xf0, 0x1f, 0x2f, 0x7a, 0x87, 0x2c, 0x55, 0xe5, 0x47, 0xbd, 0x60,
		0x74, 0xc3, 0x3c, 0xfb, 0xa2, 0xbe, 0xf2, 0x0e, 0x30, 0x58, 0x8f, 0x4f,
		0x87, 0xab, 0x40, 0x5a, 0xe4, 0xa7, 0x1f, 0xdb, 0x63, 0xb8, 0x07, 0x00,
		0x00, 0xff, 0xff, 0xb6, 0xc9, 0xf0, 0xe9, 0x2f, 0x01, 0x00, 0x00,
	},
		"migrations/06-create-nutrient-definitions-table.sql",
	)
}

func migrations_07_create_weights_table_sql() ([]byte, error) {
	return bindata_read([]byte{
		0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x00, 0xff, 0x4c, 0x8f,
		0xc1, 0x8e, 0x82, 0x50, 0x0c, 0x45, 0xf7, 0xfd, 0x8a, 0x2e, 0x67, 0x32,
		0xc3, 0x17, 0xb0, 0x42, 0x79, 0x0b, 0x13, 0x03, 0x86, 0x60, 0xe2, 0x8e,
		0x10, 0xa9, 0xd8, 0x44, 0x0a, 0xbe, 0x57, 0x02, 0x9f, 0xef, 0x43, 0x04,
		0xe9, 0xf2, 0xe4, 0x36, 0xe7, 0xde, 0x20, 0xc0, 0xbf, 0x86, 0x6b, 0x5b,
		0x2a, 0xe1, 0xb9, 0x03, 0xd8, 0x67, 0x26, 0xca, 0x0d, 0xe6, 0xd1, 0xee,
		0x68, 0x70, 0x20, 0xae, 0xef, 0xea, 0xf0, 0x07, 0x10, 0xa5, 0x57, 0xcb,
		0x24, 0x5a, 0x70, 0x85, 0x2c, 0x4a, 0x35, 0xd9, 0x7f, 0x8f, 0x1d, 0x3d,
		0x71, 0x39, 0xa5, 0x51, 0x27, 0x56, 0x36, 0x6d, 0x2f, 0x3a, 0x33, 0x4b,
		0xe5, 0x63, 0x62, 0x15, 0xb9, 0xab, 0xe5, 0x4e, 0xb9, 0x95, 0x35, 0xe7,
		0xad, 0x4d, 0x31, 0x3b, 0xde, 0x39, 0xf8, 0x0d, 0x17, 0xff, 0x21, 0x89,
		0xcd, 0x65, 0xf1, 0x7b, 0xe5, 0x88, 0x69, 0xf2, 0xad, 0xb3, 0xe9, 0xe2,
		0x5f, 0x60, 0xbb, 0x21, 0x6e, 0x07, 0x81, 0x38, 0x4b, 0x4f, 0x9f, 0x09,
		0x6b, 0xb4, 0xa2, 0x1b, 0x0b, 0x4f, 0x7e, 0x17, 0xc2, 0x2b, 0x00, 0x00,
		0xff, 0xff, 0x55, 0x4c, 0xca, 0x0b, 0xf6, 0x00, 0x00, 0x00,
	},
		"migrations/07-create-weights-table.sql",
	)
}

// Asset loads and returns the asset for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func Asset(name string) ([]byte, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		return f()
	}
	return nil, fmt.Errorf("Asset %s not found", name)
}

// AssetNames returns the names of the assets.
func AssetNames() []string {
	names := make([]string, 0, len(_bindata))
	for name := range _bindata {
		names = append(names, name)
	}
	return names
}

// _bindata is a table, holding each asset generator, mapped to its name.
var _bindata = map[string]func() ([]byte, error){
	"migrations/01-create-foods-table.sql":                 migrations_01_create_foods_table_sql,
	"migrations/02-create-food-groups-table.sql":           migrations_02_create_food_groups_table_sql,
	"migrations/03-create-languages-table.sql":             migrations_03_create_languages_table_sql,
	"migrations/04-create-language-descriptions-table.sql": migrations_04_create_language_descriptions_table_sql,
	"migrations/05-create-nutrients-table.sql":             migrations_05_create_nutrients_table_sql,
	"migrations/06-create-nutrient-definitions-table.sql":  migrations_06_create_nutrient_definitions_table_sql,
	"migrations/07-create-weights-table.sql":               migrations_07_create_weights_table_sql,
}

// AssetDir returns the file names below a certain
// directory embedded in the file by go-bindata.
// For example if you run go-bindata on data/... and data contains the
// following hierarchy:
//     data/
//       foo.txt
//       img/
//         a.png
//         b.png
// then AssetDir("data") would return []string{"foo.txt", "img"}
// AssetDir("data/img") would return []string{"a.png", "b.png"}
// AssetDir("foo.txt") and AssetDir("notexist") would return an error
// AssetDir("") will return []string{"data"}.
func AssetDir(name string) ([]string, error) {
	node := _bintree
	if len(name) != 0 {
		cannonicalName := strings.Replace(name, "\\", "/", -1)
		pathList := strings.Split(cannonicalName, "/")
		for _, p := range pathList {
			node = node.Children[p]
			if node == nil {
				return nil, fmt.Errorf("Asset %s not found", name)
			}
		}
	}
	if node.Func != nil {
		return nil, fmt.Errorf("Asset %s not found", name)
	}
	rv := make([]string, 0, len(node.Children))
	for name := range node.Children {
		rv = append(rv, name)
	}
	return rv, nil
}

type _bintree_t struct {
	Func     func() ([]byte, error)
	Children map[string]*_bintree_t
}

var _bintree = &_bintree_t{nil, map[string]*_bintree_t{
	"migrations": &_bintree_t{nil, map[string]*_bintree_t{
		"01-create-foods-table.sql":                 &_bintree_t{migrations_01_create_foods_table_sql, map[string]*_bintree_t{}},
		"02-create-food-groups-table.sql":           &_bintree_t{migrations_02_create_food_groups_table_sql, map[string]*_bintree_t{}},
		"03-create-languages-table.sql":             &_bintree_t{migrations_03_create_languages_table_sql, map[string]*_bintree_t{}},
		"04-create-language-descriptions-table.sql": &_bintree_t{migrations_04_create_language_descriptions_table_sql, map[string]*_bintree_t{}},
		"05-create-nutrients-table.sql":             &_bintree_t{migrations_05_create_nutrients_table_sql, map[string]*_bintree_t{}},
		"06-create-nutrient-definitions-table.sql":  &_bintree_t{migrations_06_create_nutrient_definitions_table_sql, map[string]*_bintree_t{}},
		"07-create-weights-table.sql":               &_bintree_t{migrations_07_create_weights_table_sql, map[string]*_bintree_t{}},
	}},
}}