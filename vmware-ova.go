package main

import (
	"errors"
	"fmt"
	"github.com/mitchellh/packer/builder/vmware/vmx"
	"github.com/mitchellh/packer/packer"
	"github.com/mitchellh/packer/packer/plugin"
	"os/exec"
	"path/filepath"
	//"reflect"
	//"strings"
)

type Builder struct {
	vmx      *vmx.Builder
	vmx_path string
	ova_path string
}

func (b *Builder) Prepare(raws ...interface{}) ([]string, error) {
	raw, ok := raws[0].(map[interface{}]interface{})
	if ok == false {
		return nil, errors.New("invalid configuration")
	}
	b.ova_path, ok = raw["source_path"].(string)
	if ok == false {
		return nil, errors.New("source_path must be defined in builder config")
	}
	b.vmx_path = ova_to_vmx_path(b.ova_path)
	raw["source_path"] = b.vmx_path
	raws[0] = raw
	b.vmx = new(vmx.Builder)
	return b.vmx.Prepare(raws...)
}

func (b *Builder) Run(ui packer.Ui, hook packer.Hook, cache packer.Cache) (packer.Artifact, error) {
	program, err := FindOvfTool()
	if err != nil {
		ui.Message(fmt.Sprintf("err: %s", err))
	}

	sourcetype := "--sourceType=OVA"
	targettype := "--targetType=VMX"
	cmd := exec.Command(program, sourcetype, targettype, b.ova_path, b.vmx_path)
	ui.Message(fmt.Sprintf("Converting OVA: %s to VMX: %s.", b.ova_path, b.vmx_path))
	cmd.Start()
	cmd.Wait()
	ui.Message(fmt.Sprintf("OVA conversion completed."))

	ui.Message(fmt.Sprintf("Starting VMX conversion."))
	return b.vmx.Run(ui, hook, cache)
}

func (b *Builder) Cancel() {
	b.vmx.Cancel()
}

func ova_to_vmx_path(ova_path string) string {
	basename := ova_path[:len(ova_path)-4]
	return basename + "/" + filepath.Base(basename+".vmx")
}

func main() {
	server, err := plugin.Server()
	if err != nil {
		panic(err)
	}
	server.RegisterBuilder(new(Builder))
	server.Serve()
}
