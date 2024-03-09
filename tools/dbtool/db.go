package dbtool

import (
	"github.com/zhanmengao/aihelp/tools/dbtool/internal/gen"
	"github.com/zhanmengao/aihelp/tools/internal/pbpkg"
	"log"
	"path"
	"path/filepath"

	"github.com/spf13/cobra"
)

type GenDBCode struct {
	protoDir      string
	pbPackagePath string
	genPackage    string
	outputDir     string
	pbGoDir       string
	buildTag      string
	tableProto    string
}

var d GenDBCode

var DBTool = &cobra.Command{
	Use:   d.Name(),
	Short: d.Synopsis(),
	Long:  d.Synopsis(),
	Run:   d.Execute,
}

func init() {
	DBTool.Flags().StringVar(&d.protoDir, "d", ".", "proto仓库目录")
	DBTool.Flags().StringVar(&d.pbPackagePath, "p", "", "pb的package name")
	DBTool.Flags().StringVar(&d.genPackage, "gen", "", "gen的package name")
	DBTool.Flags().StringVar(&d.outputDir, "o", "", "db.go输出目录")
	DBTool.Flags().StringVar(&d.pbGoDir, "go", "", ".pb.go目录")
	DBTool.Flags().StringVar(&d.buildTag, "build", "", "build tag")
	DBTool.Flags().StringVar(&d.tableProto, "table", "db.proto", "table.proto")
}

func (d *GenDBCode) Execute(cmd *cobra.Command, args []string) {
	if d.pbGoDir == "" {
		d.pbGoDir = path.Join(d.protoDir, "pb")
	}
	if d.pbPackagePath == "" {
		tableProto := path.Join(d.protoDir, "proto", d.tableProto)
		pbPackage, err := pbpkg.GetPBPackage(tableProto)
		if err != nil {
			panic(err)
		}
		d.pbPackagePath = pbPackage
	}
	d.pbPackagePath = path.Dir(d.pbPackagePath)
	d.outputDir = filepath.Join(d.outputDir, "gendb")
	if d.genPackage == "" {
		d.genPackage = d.pbPackagePath
	}
	log.Printf("%+v", d)
	gen.Run(d.pbGoDir, d.pbPackagePath, d.genPackage, d.outputDir, d.buildTag)
}

func (d *GenDBCode) Name() string {
	return "db"
}

func (d *GenDBCode) Synopsis() string {
	return "生成数据库访问代码"
}

func (d *GenDBCode) Usage() string {
	return "生成数据库访问代码"
}
