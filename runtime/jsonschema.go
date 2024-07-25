package runtime

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"path/filepath"

	"github.com/naivary/specraft/definer"
	"github.com/naivary/specraft/generator"
	"github.com/naivary/specraft/schema"
	"github.com/naivary/specraft/utils/fsutil"
	"sigs.k8s.io/controller-tools/pkg/loader"
	"sigs.k8s.io/controller-tools/pkg/markers"
)

func JSONSchema(rootDir string, roots ...string) (Runtime[*schema.JSON], error) {
	jrt := &jsonSchemaRuntime{}
	if err := fsutil.MkdirAllIfNotExsting(rootDir, os.FileMode(0755)); err != nil {
		return nil, err
	}
	jrt.rootDir = rootDir

	reg := &markers.Registry{}
	jrt.reg = reg

	defn, err := definer.JSONSchema(reg)
	if err != nil {
		return nil, err
	}
	jrt.defn = defn

	col := &markers.Collector{
		Registry: reg,
	}
	jrt.col = col
	pkgs, err := jrt.loadPackages(roots...)
	if err != nil {
		return nil, err
	}
	jrt.pkgs = pkgs

	jrt.gen = generator.JSONSchema()
	jrt.schemas = make(map[string]*os.File)

	return jrt, nil
}

type jsonSchemaRuntime struct {
	reg *markers.Registry

	col *markers.Collector

	pkgs map[*loader.Package][]*markers.TypeInfo

	gen generator.Generator[*schema.JSON]

	defn definer.Definer[*schema.JSON]

	// schemas are all the generated files indexed by $id
	schemas map[string]*os.File

	rootDir string
}

func (jrt *jsonSchemaRuntime) Packages() map[*loader.Package][]*markers.TypeInfo {
	return jrt.pkgs
}

func (jrt *jsonSchemaRuntime) Registry() *markers.Registry {
	return jrt.reg
}

func (jrt *jsonSchemaRuntime) Schemas() map[string]*os.File {
	return jrt.schemas
}

func (jrt *jsonSchemaRuntime) Generate() error {
	for pkg, typeInfos := range jrt.pkgs {
		for _, typeInfo := range typeInfos {
			schm, err := jrt.gen.Generate(jrt.defn, pkg, typeInfo)
			if errors.Is(err, generator.ErrNonStructType) {
				continue
			}
			if err != nil {
				return err
			}
			err = jrt.writeToFile(typeInfo.Name, schm)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func (jrt *jsonSchemaRuntime) writeToFile(name string, s *schema.JSON) error {
	fileName := fmt.Sprintf("%s.json", name)
	file, err := jrt.createFile(fileName)
	if err != nil {
		return err
	}
	defer file.Close()
	err = json.NewEncoder(file).Encode(s)
	if err != nil {
		return err
	}
	jrt.schemas[s.ID] = file
	return nil
}

func (jrt *jsonSchemaRuntime) createFile(name string) (*os.File, error) {
	path := filepath.Join(jrt.rootDir, name)
	return os.Create(path)
}

func (jrt *jsonSchemaRuntime) loadPackages(roots ...string) (map[*loader.Package][]*markers.TypeInfo, error) {
	pkgs, err := loader.LoadRoots(roots...)
	if err != nil {
		return nil, err
	}
	typeInfos := make(map[*loader.Package][]*markers.TypeInfo, 0)
	for _, pkg := range pkgs {
		pkg.NeedTypesInfo()
		pkg.NeedSyntax()

		infos := make([]*markers.TypeInfo, 0)
		err := markers.EachType(jrt.col, pkg, func(info *markers.TypeInfo) {
			infos = append(infos, info)
		})
		if err != nil {
			return nil, err
		}
		typeInfos[pkg] = infos
	}

	return typeInfos, nil
}
