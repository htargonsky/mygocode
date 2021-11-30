package main

import (
	"fmt"
)

type VirtualMachine struct{
    ip string
    hostName string
    diskgb int
    ram int
}

// use a pointer here as we are making a change
func (v *VirtualMachine) increaseRam(additionalRam int){
    v.ram = v.ram + additionalRam    
}

func (v *VirtualMachine) changeHost(newHostName string){
    v.hostName = newHostName   
}

func (v VirtualMachine)display(){
    fmt.Println("Primary IP Address:", v.ip)    // primary IP address
    fmt.Println("Hostname:", v.hostName)        // hostname
    fmt.Println("Disk GB:", v.diskgb)           // diskgb
    fmt.Println("RAM:", v.ram)                  // ram
}
func main() {

	vm := VirtualMachine{ip: "129.10.43.124", hostName:  "optum.com", ram: 1024}

	vm.display()

	vm.increaseRam(1024)

	vm.changeHost("uhg.com")
	vm.display()
}
