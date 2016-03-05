package libvirt

import (
	"encoding/xml"
)

type defVolume struct {
	XMLName xml.Name  `xml:"volume"`
	Name  string  `xml:"name"`
	Target struct {
		Format struct {
			Type  string  `xml:"type,attr"`
		} `xml:"format"`
	} `xml:"target"`
	BackingStore struct {
		Path string `xml:"path"`
		Format struct {
			Type    string    `xml:"type,attr"`
		} `xml:"format"`
	} `xml:"backingStore"`
}

type defDisk struct {
	XMLName xml.Name  `xml:"disk"`
	Type    string    `xml:"type,attr"`
	Device  string    `xml:"device,attr"`
	Format struct {
		Type  string  `xml:"type,attr"`

	} `xml:"format"`
	Source struct {
		Pool string `xml:"pool,attr"`
		Volume string `xml:"volume,attr"`
	} `xml:"source"`
	Target struct {
		Dev  string  `xml:"dev,attr"`
		Bus  string  `xml:"bus,attr"`
	} `xml:"target"`
}

type defDomain struct {
	XMLName xml.Name  `xml:"domain"`
	Name    string    `xml:"name"`
	Type    string    `xml:"type,attr"`
	Os      defOs     `xml:"os"`
	Memory  defMemory `xml:"memory"`
	VCpu    defVCpu   `xml:"vcpu"`
	Devices struct {
		RootDisk defDisk `xml:"disk"`
	} `xml:"devices"`
	Spice struct {
		Type    string    `xml:"type,attr"`
		Autoport    bool    `xml:"autoport,attr"`
	} `xml:"graphics"`
	ChannelSpice struct {
		Type    string    `xml:"type,attr"`
		Target struct {
			Type    string    `xml:"type,attr"`
			Name    string    `xml:"name,attr"`
		} `xml:"target"`
	} `xml:"channel"`
}

type defOs struct {
	Type defOsType `xml:"type"`
}

type defOsType struct {
	Arch    string `xml:"arch,attr"`
	Machine string `xml:"machine,attr"`
	Name    string `xml:"chardata"`
}

type defMemory struct {
	Unit   string `xml:"unit,attr"`
	Amount int    `xml:"chardata"`
}

type defVCpu struct {
	Placement string `xml:"unit,attr"`
	Amount    int    `xml:"chardata"`
}

// Creates a domain definition with the defaults
// the provider uses
func newDomainDef() defDomain {
		// libvirt domain definition
	domainDef := defDomain{}
	domainDef.Type = "kvm"

	domainDef.Os = defOs{}
	domainDef.Os.Type = defOsType{}
	domainDef.Os.Type.Arch = "x86_64"
    domainDef.Os.Type.Machine = "pc-i440fx-2.4"
	domainDef.Os.Type.Name = "hvm"

	domainDef.Memory = defMemory{}
	domainDef.Memory.Unit = "MiB"
	domainDef.Memory.Amount = 512

	domainDef.VCpu = defVCpu{}
	domainDef.VCpu.Placement = "static"
	domainDef.VCpu.Amount = 1

	domainDef.Spice.Type = "spice"
	domainDef.Spice.Autoport = true
	domainDef.ChannelSpice.Type = "spicevmc"
	return domainDef
}
