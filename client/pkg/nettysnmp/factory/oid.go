package factory

var SysDescr = ".1.3.6.1.2.1.1.1.0"
var SysObjectID = ".1.3.6.1.2.1.1.2.0"
var SysUpTime = ".1.3.6.1.2.1.1.3.0"
var SysName = ".1.3.6.1.2.1.1.5.0"

// IF-MIB and related
var IfIndex = ".1.3.6.1.2.1.2.2.1.1"
var IfDescr = ".1.3.6.1.2.1.2.2.1.2"
var IfType = ".1.3.6.1.2.1.2.2.1.3"
var IfMtu = ".1.3.6.1.2.1.2.2.1.4"
var IfSpeed = ".1.3.6.1.2.1.2.2.1.5"
var IfPhysAddr = ".1.3.6.1.2.1.2.2.1.6"
var IfAdminStatus = ".1.3.6.1.2.1.2.2.1.7"
var IfOperStatus = ".1.3.6.1.2.1.2.2.1.8"
var IfLastChange = ".1.3.6.1.2.1.2.2.1.9"
var IfAlias = ".1.3.6.1.2.1.31.1.1.1.18"
var IfHighSpeed = ".1.3.6.1.2.1.31.1.1.1.15"
var IfAdEntIfIndex = ".1.3.6.1.2.1.4.20.1.2"
var IfAdEntNetMask = ".1.3.6.1.2.1.4.20.1.3"

// LLDP-MIB
var LldpLocChassisId = ".1.0.8802.1.1.2.1.3.2.0"
var LldpLocPortId = ".1.0.8802.1.1.2.1.3.7.1.3"
var LldpLocPortDesc = ".1.0.8802.1.1.2.1.3.7.1.4"
var LldpRemChassisIdSubtype = ".1.0.8802.1.1.2.1.4.1.1.4"
var LldpRemChassisId = ".1.0.8802.1.1.2.1.4.1.1.5"
var LldpRemPortIdSubtype = ".1.0.8802.1.1.2.1.4.1.1.6"
var LldpRemPortId = ".1.0.8802.1.1.2.1.4.1.1.7"
var LldpRemPortDesc = ".1.0.8802.1.1.2.1.4.1.1.8"
var LldpRemSysName = ".1.0.8802.1.1.2.1.4.1.1.9"

// ENTITY-MIB

var EntPhysicalDescr = ".1.3.6.1.2.1.47.1.1.1.1.2"

var EntPhysicalClass = ".1.3.6.1.2.1.47.1.1.1.1.5"

var EntPhysicalName = ".1.3.6.1.2.1.47.1.1.1.1.7"
var EntPhysicalHardwareRev = ".1.3.6.1.2.1.47.1.1.1.1.8"
var EntPhysicalFirmwareRev = ".1.3.6.1.2.1.47.1.1.1.1.9"
var EntPhysicalSoftwareRev = ".1.3.6.1.2.1.47.1.1.1.1.10"
var EntPhysicalSerialNum = ".1.3.6.1.2.1.47.1.1.1.1.11"

// -------- stack Cisco ------- #
var CswSwitchRole = ".1.3.6.1.4.1.9.9.500.1.2.1.1.3"
var CswSwitchHwPriority = ".1.3.6.1.4.1.9.9.500.1.2.1.1.5"
var CswSwitchState = ".1.3.6.1.4.1.9.9.500.1.2.1.1.6"
var CswSwitchMacAddress = ".1.3.6.1.4.1.9.9.500.1.2.1.1.7"

// --------- stack H3C --------
var H3cStackMemberID = ".1.3.6.1.4.1.2011.10.2.91.2.1.1"
var H3cStackConfigMemberID = ".1.3.6.1.4.1.2011.10.2.91.2.1.2"
var H3cStackPriority = ".1.3.6.1.4.1.2011.10.2.91.2.1.3"

// ---------- RuiJie stack ----------
var ScMemberMacAddress = ".1.3.6.1.4.1.4881.1.1.10.2.31.1.2.1.1.1"
var ScMemberNumber = ".1.3.6.1.4.1.4881.1.1.10.2.31.1.2.1.1.2"
var ScMemberOperStatus = ".1.3.6.1.4.1.4881.1.1.10.2.31.1.2.1.1.3"
var ScMemberDeviceID = ".1.3.6.1.4.1.4881.1.1.10.2.31.1.2.1.1.4"
var ScMemberRowStatus = ".1.3.6.1.4.1.4881.1.1.10.2.31.1.2.1.1.5"

// PaloAlto Basic Info
var PanSysSwVersion = ".1.3.6.1.4.1.25461.2.1.2.1.1"
var PanSysSerialNumber = ".1.3.6.1.4.1.25461.2.1.2.1.3"

var Dot1dTpFdbAddress = ".1.3.6.1.2.1.17.4.3.1.1"     // The MAC address of the FDB entry. mac_address
var Dot1dTpFdbPort = ".1.3.6.1.2.1.17.4.3.1.2"        // The port number of the FDB entry. int
var Dot1dBasePortIfIndex = ".1.3.6.1.2.1.17.1.4.1.2"  // The ifIndex of the port. int
var IpNetToMediaPhysAddress = ".1.3.6.1.2.1.4.22.1.2" // The MAC address of the port.
var IpNetToMediaType = ".1.3.6.1.2.1.4.22.1.4" // 1: other(1), 2: invalid(2), 3: dynamic(3), 4: static(4)
 