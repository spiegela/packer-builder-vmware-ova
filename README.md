# packer-builder-vmware-ova

This packer plugin leverages [VMware OVF Tool](http://www.vmware.com/support/developer/ovf) to build a Vagrant box based on a generic (not already Vagrant friendly) OVA file like those used in VMware virtual appliances.

This plugin is used to build the a couple boxes (ESXi 5.5/6.0 and vCSA 5.5/6.0) available in the [packer templates here](https://github.com/spiegela/packer-templates)

## Prerequisites

Software:

  * VMware OVF Tool

## Installation

Starting from Packer v0.7.0 there are new ways of installing plugins, [see the official Packer documentation](http://www.packer.io/docs/extend/plugins.html) for further instructions.

## Usage

In your JSON template add the following post processor:

```json
  "builders": [
    {
        "type": "vmware-ova",
        "source_path": "<some file>.ova"
    }
  ]
```

If you don't want to compile the code, you can [grab a release here](https://github.com/spiegela/packer-builder-vmware-ova/releases).
