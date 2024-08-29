# Circuits
A circuit represents a physical size-to-side cable connection between network devices.
It's only aims to management Internet or Intranet circuit type.
- **Internet**: Internet Dedicated line, ADSL line, PPPoE line
- **Intranet**: P2P Line, DarkFiber, MPLSVPN, IEPL and etc

## Fields

### Provider
The provider to which this circuit belongs.

### CID
The unique identifier circuit ID assigned from provider

### CircuitType
- Internet
- Intranet

### Status
> Admin status of the circuit.
- **Active**: when circuit status is active and ip_address is configured, ICMP monitor will be enable; if ASizeInterface is associated with circuit, traffic monitoring/alerting will be enable.
- **InActive**: Disable circuit ICMP monitoring/alerting and traffic monitoring/alerting

### Description
A brief description of the circuit


### BandWidth
the bandwidth of circuit, it will be used to calculate the real traffic utilization of the circuit.

### IPAddress
the ip address of the circuit. it will be treat as /32 netmask for IPv4 or /128 netmask for IPv6.IP pool is not support.
when ip address set, circuit will enable icmp ping monitor depends on circuit status, make sure icmp is not blocked by security policy if you want know the metrics of circuit.


### ASide
ASize is local size of the circuit, if circuit type is `Internet`, the gateway device, such as firewall, router or internet-switch which used to connected the internet line should be treated as a_device, and the interface as a_interface.
the ZSize is Provider network, which is not supported to edit.

### ZSize
when circuit type is `Intranet`, circuit will be treated as inter-connection of sites or buildings.


## ErrorCodes reference

