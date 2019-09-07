package golang

import (
	"reflect"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestInitializationModel(t *testing.T) {
	recipe := SourceCode{
		PackageName: "typical",
		Imports: []Import{
			{"", "fmt"},
			{"tgo", "github.com/typical-go/typical-go"},
		},
		Structs: []Struct{
			{
				Name: "Config",
				Fields: []reflect.StructField{
					{Name: "Name", Type: reflect.TypeOf("something")},
				},
			},
		},
		Constructors: []string{
			"contructor1",
		},
	}

	// fmt.Println(model.String())

	var builder strings.Builder
	recipe.Write(&builder)

	require.Equal(t, "// Autogenerated by Typical-Go. DO NOT EDIT.\n\npackage typical\nimport  \"fmt\"\nimport tgo \"github.com/typical-go/typical-go\"\ntype Config struct{\nName string\n}\nfunc init() {\nContext.AddConstructor(contructor1)\n}\n",
		builder.String())
}

func TestBlank(t *testing.T) {
	testcase := []struct {
		src   SourceCode
		blank bool
	}{
		{src: SourceCode{}, blank: true},
		{src: SourceCode{Imports: []Import{Import{}}}, blank: false},
		{src: SourceCode{Structs: []Struct{Struct{}}}, blank: false},
		{src: SourceCode{Constructors: []string{""}}, blank: false},
		{src: SourceCode{MockTargets: []string{""}}, blank: false},
		{src: SourceCode{TestTargets: []string{""}}, blank: false},
	}

	for i, tt := range testcase {
		require.Equal(t, tt.blank, tt.src.Blank(), "%d: %+v", i, tt.src)
	}

}