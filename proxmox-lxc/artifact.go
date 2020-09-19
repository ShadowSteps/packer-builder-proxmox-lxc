package proxmox_lxc

import (
	"fmt"
	"log"
	"strconv"

	"github.com/Telmate/proxmox-api-go/proxmox"
	"github.com/hashicorp/packer/packer"
)

type Artifact struct {
	templateID    int
	proxmoxClient *proxmox.Client

	// StateData should store data such as GeneratedData
	// to be shared with post-processors
	StateData map[string]interface{}
}

// Artifact implements packer.Artifact
var _ packer.Artifact = &Artifact{}

func (*Artifact) BuilderId() string {
	return BuilderId
}

func (*Artifact) Files() []string {
	return nil
}

func (a *Artifact) Id() string {
	return strconv.Itoa(a.templateID)
}

func (a *Artifact) String() string {
	return fmt.Sprintf("A template was created: %d", a.templateID)
}

func (a *Artifact) State(name string) interface{} {
	return a.StateData[name]
}

func (a *Artifact) Destroy() error {
	log.Printf("Destroying template: %d", a.templateID)
	_, err := a.proxmoxClient.DeleteVm(proxmox.NewVmRef(a.templateID))
	return err
}