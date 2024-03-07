package gen

import (
	"log"
	"os"
)

func Run(pbGoPath, pbPackage, genPackage, outputDir, buildTag string) {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	st, err := os.Stat(pbGoPath)
	if err != nil {
		panic(err)
	}
	if !st.IsDir() {
		panic(pbGoPath + " not a dir")
	}
	meta := NewDBMeta()
	err = meta.ParseComments(pbGoPath)
	if err != nil {
		panic(err)
	}
	_ = os.RemoveAll(outputDir)
	if err = os.MkdirAll(outputDir, os.ModePerm); err != nil {
		panic(err)
	}
	meta.PkgPath = pbPackage
	meta.outputDir = outputDir
	meta.buildTag = buildTag
	meta.GenPkg = genPackage

	meta.Generate(pbGoPath)
}
