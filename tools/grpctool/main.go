package Run

import (
	"flag"
	"google.golang.org/protobuf/compiler/protogen"
	"google.golang.org/protobuf/types/pluginpb"
)

func Run(f func(plug *protogen.Plugin) error) {
	flag.Parse()
	protogen.Options{
		ParamFunc: flag.CommandLine.Set,
	}.Run(f)
}

func DefaultRun(plug *protogen.Plugin) error {
	plug.SupportedFeatures = uint64(pluginpb.CodeGeneratorResponse_FEATURE_PROTO3_OPTIONAL)
	for _, f := range plug.Files {
		if !f.Generate {
			continue
		}
		//generateFile(gen, f, *omitempty, *buildTag)
	}
	return nil
}
