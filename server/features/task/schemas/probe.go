package schemas

type DNSProbe struct {
	Server    string   `json:"server"`
	ProbeList []string `json:"probeList"`
}

type RadiusProbe struct {
	Server []string `json:"server"`
	Port   int      `json:"port"`
}

type DHCPProbe struct {
	Server []string `json:"server"`
	Port   int      `json:"port"`
}

type HTTPProbe struct {
}

type ICMPProbe struct {

}

type TCPProbe struct {

}

type UDPProbe struct {
	
}