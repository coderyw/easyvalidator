package main

import (
	"errors"
	"flag"
	"fmt"
	"github.com/coderyw/easyvalidator/bootstrap"
	"github.com/coderyw/easyvalidator/parser"
	"github.com/coderyw/easyvalidator/validator"
	"os"
	"path/filepath"
	"strings"
	// Reference the gen package to be friendly to vendoring tools,
	// as it is an indirect dependency.
	// (The temporary bootstrapping code uses it.)
	_ "github.com/coderyw/easyvalidator/gen"
)

// var buildTags = flag.String("build_tags", "", "build tags to add to generated file")
// var genBuildFlags = flag.String("gen_build_flags", "", "build flags when running the generator while bootstrapping")
var allStructs = flag.Bool("all", false, "generate validator for all structs in a file")
var excludeFileSuffix = flag.String("exclude_suff", "", "排除文件后缀名")
var includeFileSuffix = flag.String("include_suff", "", "只需要包含这个后缀的go文件")
var leaveTemps = flag.Bool("leave_temps", false, "do not delete temporary files")
var noformat = flag.Bool("noformat", false, "do not run 'gofmt -w' on output file")
var specifiedName = flag.String("output_filename", "", "specify the filename of the output")
var processPkg = flag.Bool("pkg", false, "process the whole package instead of just the given file")
var deleteBefore = flag.Bool("del_before", true, "开始处理前删除以前存在的文件")
var usages = flag.Bool("usages", false, "print usage information")

func generate(fname string) (err error) {
	fInfo, err := os.Stat(fname)
	if err != nil {
		return err
	}

	p := parser.Parser{
		AllStructs: *allStructs,
		ExSuff:     *excludeFileSuffix,
		WantSuff:   *includeFileSuffix,
	}
	if err := p.Parse(fname, fInfo.IsDir()); err != nil {
		return fmt.Errorf("Error parsing %v: %v", fname, err)
	}
	//fmt.Println(p)

	var outName string
	if fInfo.IsDir() {
		outName = filepath.Join(fname, p.PkgName+"_easyvalidator.go")
	} else {
		if s := strings.TrimSuffix(fname, ".go"); s == fname {
			return errors.New("Filename must end in '.go'")
		} else {
			outName = s + "_easyvalidator.go"
		}
	}
	if *deleteBefore { //删除文件
		os.Remove(outName)
	}

	if *specifiedName != "" {
		outName = *specifiedName
	}

	var trimmedBuildTags string
	//if *buildTags != "" {
	//	trimmedBuildTags = strings.TrimSpace(*buildTags)
	//}

	var trimmedGenBuildFlags string
	//if *genBuildFlags != "" {
	//	trimmedGenBuildFlags = strings.TrimSpace(*genBuildFlags)
	//}

	g := bootstrap.Generator{
		BuildTags:     trimmedBuildTags,
		GenBuildFlags: trimmedGenBuildFlags,
		PkgPath:       p.PkgPath,
		PkgName:       p.PkgName,
		Types:         p.StructNames,
		LeaveTemps:    *leaveTemps,
		OutName:       outName,
		NoFormat:      *noformat,
	}

	if err := g.Run(); err != nil {
		return fmt.Errorf("Bootstrap failed: %v", err)
	}
	return nil
}

func main() {
	flag.Parse()

	files := flag.Args()
	if usages != nil && *usages {
		printUsages()
		os.Exit(1)
	}

	gofile := os.Getenv("GOFILE")
	if *processPkg {
		gofile = filepath.Dir(gofile)
	}

	if len(files) == 0 && gofile != "" {
		files = []string{gofile}
	} else if len(files) == 0 {
		flag.Usage()
		os.Exit(1)
	}

	for _, fname := range files {
		if err := generate(fname); err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
	}
}

func printUsages() {
	out := flag.CommandLine.Output()
	buffer := strings.Builder{}

	buffer.WriteString("Usage of validator key:\n")
	buffer.WriteString("tag name: easy_valid\n")
	buffer.WriteString("example: \n\n")
	buffer.WriteString("type UserInfo struct { \n")
	buffer.WriteString("\tName string `json:\"name\" easy_valid:\"required\"` \n")
	buffer.WriteString("}\n\n")
	buffer.WriteString("单个属性校验(Explain of single validator key):\n")
	for _, v := range validator.ValidatorUsages {
		buffer.WriteString("\t")
		buffer.WriteString(v.Key)
		buffer.WriteString("\n")
		buffer.WriteString("\t\t")
		buffer.WriteString(v.Zh)
		buffer.WriteString("\n")
		buffer.WriteString("\t\t")
		buffer.WriteString(v.En)
		buffer.WriteString("\n")

	}
	buffer.WriteString("\n")
	buffer.WriteString("多个属性联合校验(Joint verification of multiple attributes):\n")
	buffer.WriteString("使用方式：Age=int_gte:1,int_lt:0|str_eq:12\n")
	buffer.WriteString("说明：要求属性0<Age<=1时，校验本属性字符串长度是否等于2\n")
	for _, v := range validator.ValidatorCrossUsages {
		buffer.WriteString("\t")
		buffer.WriteString(v.Key)
		buffer.WriteString("\n")
		buffer.WriteString("\t\t")
		buffer.WriteString(v.Zh)
		buffer.WriteString("\n")
		buffer.WriteString("\t\t")
		buffer.WriteString(v.En)
		buffer.WriteString("\n")

	}
	fmt.Fprintln(out, buffer.String())
	//flag.PrintDefaults()
	//fmt.Println("Usage: easyvalidator [flags] [files]")
}
