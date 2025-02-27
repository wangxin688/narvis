package sysobjectid

import (
	manufacturer "github.com/wangxin688/narvis/intend/model/manufacturer"
	platform "github.com/wangxin688/narvis/intend/model/platform"
	"github.com/wangxin688/narvis/intend/netdisco/devicemodel"
)

func FortiNetDeviceModel(sysObjId string) *devicemodel.DeviceModel {
	stringPlatform := string(platform.FortiNet)
	oidMap := map[string]map[string]string{
		".1.3.6.1.4.1.12356.100":          {"platform": stringPlatform, "model": "FORTIGATE 100"},
		".1.3.6.1.4.1.12356.10004":        {"platform": stringPlatform, "model": "FORTIWEB 1000D"},
		".1.3.6.1.4.1.12356.1001":         {"platform": stringPlatform, "model": "FORTIGATE 100A"},
		".1.3.6.1.4.1.12356.101.1.10":     {"platform": stringPlatform, "model": "FORTIGATE ONE"},
		".1.3.6.1.4.1.12356.101.1.1000":   {"platform": stringPlatform, "model": "FORTIGATE 100"},
		".1.3.6.1.4.1.12356.101.1.10000":  {"platform": stringPlatform, "model": "FORTIGATE 1000"},
		".1.3.6.1.4.1.12356.101.1.10001":  {"platform": stringPlatform, "model": "FORTIGATE 1000A"},
		".1.3.6.1.4.1.12356.101.1.10002":  {"platform": stringPlatform, "model": "FORTIGATE 1000AFA2"},
		".1.3.6.1.4.1.12356.101.1.10003":  {"platform": stringPlatform, "model": "FORTIGATE 1000ALENC"},
		".1.3.6.1.4.1.12356.101.1.10004":  {"platform": stringPlatform, "model": "FORTIGATE 1000C"},
		".1.3.6.1.4.1.12356.101.1.10005":  {"platform": stringPlatform, "model": "FORTIGATE 1000D"},
		".1.3.6.1.4.1.12356.101.1.10006":  {"platform": stringPlatform, "model": "FORTIGATE 1000E"},
		".1.3.6.1.4.1.12356.101.1.10007":  {"platform": stringPlatform, "model": "FORTIGATE 1101E"},
		".1.3.6.1.4.1.12356.101.1.1001":   {"platform": stringPlatform, "model": "FORTIGATE 100F"},
		".1.3.6.1.4.1.12356.101.1.1002":   {"platform": stringPlatform, "model": "FORTIGATE 110C"},
		".1.3.6.1.4.1.12356.101.1.1003":   {"platform": stringPlatform, "model": "FORTIGATE 111C"},
		".1.3.6.1.4.1.12356.101.1.1004":   {"platform": stringPlatform, "model": "FORTIGATE-100D"},
		".1.3.6.1.4.1.12356.101.1.1005":   {"platform": stringPlatform, "model": "FORTIGATE RUGGED 100C"},
		".1.3.6.1.4.1.12356.101.1.1006":   {"platform": stringPlatform, "model": "FORTIGATE 140E-POE"},
		".1.3.6.1.4.1.12356.101.1.1010":   {"platform": stringPlatform, "model": "FORTIGATE 101F"},
		".1.3.6.1.4.1.12356.101.1.1041":   {"platform": stringPlatform, "model": "FORTIGATE 100E"},
		".1.3.6.1.4.1.12356.101.1.1042":   {"platform": stringPlatform, "model": "FORTIGATE 100EF"},
		".1.3.6.1.4.1.12356.101.1.1043":   {"platform": stringPlatform, "model": "FORTIGATE 101E"},
		".1.3.6.1.4.1.12356.101.1.12000":  {"platform": stringPlatform, "model": "FORTIGATE 1200D"},
		".1.3.6.1.4.1.12356.101.1.12400":  {"platform": stringPlatform, "model": "FORTIGATE 1240B"},
		".1.3.6.1.4.1.12356.101.1.1401":   {"platform": stringPlatform, "model": "FORTIGATE 140D"},
		".1.3.6.1.4.1.12356.101.1.1402":   {"platform": stringPlatform, "model": "FORTIGATE 140P"},
		".1.3.6.1.4.1.12356.101.1.1403":   {"platform": stringPlatform, "model": "FORTIGATE 140T"},
		".1.3.6.1.4.1.12356.101.1.15000":  {"platform": stringPlatform, "model": "FORTIGATE 1500D"},
		".1.3.6.1.4.1.12356.101.1.15001":  {"platform": stringPlatform, "model": "FORTIGATE 1500DT"},
		".1.3.6.1.4.1.12356.101.1.15002":  {"platform": stringPlatform, "model": "FORTIGATE 1801F"},
		".1.3.6.1.4.1.12356.101.1.18000":  {"platform": stringPlatform, "model": "FORTIGATE 2201E/2200E"},
		".1.3.6.1.4.1.12356.101.1.18001":  {"platform": stringPlatform, "model": "FORTIGATE 2201E"},
		".1.3.6.1.4.1.12356.101.1.18003":  {"platform": stringPlatform, "model": "FORTIGATE 1800F"},
		".1.3.6.1.4.1.12356.101.1.20":     {"platform": stringPlatform, "model": "FORTIGATE VM"},
		".1.3.6.1.4.1.12356.101.1.2000":   {"platform": stringPlatform, "model": "FORTIGATE 200"},
		".1.3.6.1.4.1.12356.101.1.20000":  {"platform": stringPlatform, "model": "FORTIGATE 2000"},
		".1.3.6.1.4.1.12356.101.1.2001":   {"platform": stringPlatform, "model": "FORTIGATE 200A"},
		".1.3.6.1.4.1.12356.101.1.2002":   {"platform": stringPlatform, "model": "FORTIGATE 224B"},
		".1.3.6.1.4.1.12356.101.1.2003":   {"platform": stringPlatform, "model": "FORTIGATE 200B"},
		".1.3.6.1.4.1.12356.101.1.2004":   {"platform": stringPlatform, "model": "FORTIGATE 200BPOE"},
		".1.3.6.1.4.1.12356.101.1.2005":   {"platform": stringPlatform, "model": "FORTIGATE 200D"},
		".1.3.6.1.4.1.12356.101.1.2006":   {"platform": stringPlatform, "model": "FORTIGATE 240D"},
		".1.3.6.1.4.1.12356.101.1.2007":   {"platform": stringPlatform, "model": "FORTIGATE 201E"},
		".1.3.6.1.4.1.12356.101.1.2008":   {"platform": stringPlatform, "model": "FORTIGATE 240DP"},
		".1.3.6.1.4.1.12356.101.1.2009":   {"platform": stringPlatform, "model": "FORTIGATE 200E"},
		".1.3.6.1.4.1.12356.101.1.2010":   {"platform": stringPlatform, "model": "FORTIGATE 201E"},
		".1.3.6.1.4.1.12356.101.1.2012":   {"platform": stringPlatform, "model": "FORTIGATE 201F"},
		".1.3.6.1.4.1.12356.101.1.2013":   {"platform": stringPlatform, "model": "FORTIGATE 280D-POE"},
		".1.3.6.1.4.1.12356.101.1.210":    {"platform": stringPlatform, "model": "FORTIWIFI 20C"},
		".1.3.6.1.4.1.12356.101.1.212":    {"platform": stringPlatform, "model": "FORTIGATE 20C"},
		".1.3.6.1.4.1.12356.101.1.213":    {"platform": stringPlatform, "model": "FORTIWIFI 20CA"},
		".1.3.6.1.4.1.12356.101.1.214":    {"platform": stringPlatform, "model": "FORTIGATE 20CA"},
		".1.3.6.1.4.1.12356.101.1.25000":  {"platform": stringPlatform, "model": "FORTIGATE 2500E"},
		".1.3.6.1.4.1.12356.101.1.26001":  {"platform": stringPlatform, "model": "FORTIGATE 2601F"},
		".1.3.6.1.4.1.12356.101.1.30":     {"platform": stringPlatform, "model": "FORTIGATE VM64"},
		".1.3.6.1.4.1.12356.101.1.3000":   {"platform": stringPlatform, "model": "FORTIGATE 300"},
		".1.3.6.1.4.1.12356.101.1.30000":  {"platform": stringPlatform, "model": "FORTIGATE 3000"},
		".1.3.6.1.4.1.12356.101.1.30001":  {"platform": stringPlatform, "model": "FORTIGATE 3301E"},
		".1.3.6.1.4.1.12356.101.1.30002":  {"platform": stringPlatform, "model": "FORTIGATE 3300"},
		".1.3.6.1.4.1.12356.101.1.3001":   {"platform": stringPlatform, "model": "FORTIGATE 100A"},
		".1.3.6.1.4.1.12356.101.1.3002":   {"platform": stringPlatform, "model": "FORTIGATE 310B"},
		".1.3.6.1.4.1.12356.101.1.3003":   {"platform": stringPlatform, "model": "FORTIGATE 300D"},
		".1.3.6.1.4.1.12356.101.1.3004":   {"platform": stringPlatform, "model": "FORTIGATE 311B"},
		".1.3.6.1.4.1.12356.101.1.3005":   {"platform": stringPlatform, "model": "FORTIGATE 300C"},
		".1.3.6.1.4.1.12356.101.1.3006":   {"platform": stringPlatform, "model": "FORTIGATE 300D"},
		".1.3.6.1.4.1.12356.101.1.3007":   {"platform": stringPlatform, "model": "FORTIGATE 300E"},
		".1.3.6.1.4.1.12356.101.1.3008":   {"platform": stringPlatform, "model": "FORTIGATE 301E"},
		".1.3.6.1.4.1.12356.101.1.30160":  {"platform": stringPlatform, "model": "FORTIGATE 3016B"},
		".1.3.6.1.4.1.12356.101.1.302":    {"platform": stringPlatform, "model": "FORTIGATE 30B"},
		".1.3.6.1.4.1.12356.101.1.304":    {"platform": stringPlatform, "model": "FORTIGATE 30D"},
		".1.3.6.1.4.1.12356.101.1.30400":  {"platform": stringPlatform, "model": "FORTIGATE 3040B"},
		".1.3.6.1.4.1.12356.101.1.30401":  {"platform": stringPlatform, "model": "FORTIGATE 3140B"},
		".1.3.6.1.4.1.12356.101.1.305":    {"platform": stringPlatform, "model": "FORTIGATE 30DPOE"},
		".1.3.6.1.4.1.12356.101.1.306":    {"platform": stringPlatform, "model": "FORTIGATE 30E"},
		".1.3.6.1.4.1.12356.101.1.31":     {"platform": stringPlatform, "model": "FORTIGATE VM64 VMX"},
		".1.3.6.1.4.1.12356.101.1.310":    {"platform": stringPlatform, "model": "FORTIWIFI 30B"},
		".1.3.6.1.4.1.12356.101.1.31000":  {"platform": stringPlatform, "model": "FORTIGATE 3100D"},
		".1.3.6.1.4.1.12356.101.1.314":    {"platform": stringPlatform, "model": "FORTIWIFI 30D"},
		".1.3.6.1.4.1.12356.101.1.315":    {"platform": stringPlatform, "model": "FORTIWIFI 30DPOE"},
		".1.3.6.1.4.1.12356.101.1.316":    {"platform": stringPlatform, "model": "FORTIWIFI 30E"},
		".1.3.6.1.4.1.12356.101.1.32":     {"platform": stringPlatform, "model": "FORTIGATE VM64 SVM"},
		".1.3.6.1.4.1.12356.101.1.32000":  {"platform": stringPlatform, "model": "FORTIGATE 3200D"},
		".1.3.6.1.4.1.12356.101.1.32401":  {"platform": stringPlatform, "model": "FORTIGATE 3240C"},
		".1.3.6.1.4.1.12356.101.1.34011":  {"platform": stringPlatform, "model": "FORTIGATE 3401E"},
		".1.3.6.1.4.1.12356.101.1.36000":  {"platform": stringPlatform, "model": "FORTIGATE 3600"},
		".1.3.6.1.4.1.12356.101.1.36001":  {"platform": stringPlatform, "model": "FORTIGATE 3600E"},
		".1.3.6.1.4.1.12356.101.1.36003":  {"platform": stringPlatform, "model": "FORTIGATE 3600A"},
		".1.3.6.1.4.1.12356.101.1.36004":  {"platform": stringPlatform, "model": "FORTIGATE 3600C"},
		".1.3.6.1.4.1.12356.101.1.36011":  {"platform": stringPlatform, "model": "FORTIGATE 3601E"},
		".1.3.6.1.4.1.12356.101.1.37000":  {"platform": stringPlatform, "model": "FORTIGATE 3700D"},
		".1.3.6.1.4.1.12356.101.1.37001":  {"platform": stringPlatform, "model": "FORTIGATE 3700DX"},
		".1.3.6.1.4.1.12356.101.1.38001":  {"platform": stringPlatform, "model": "FORTIGATE 3800D"},
		".1.3.6.1.4.1.12356.101.1.38100":  {"platform": stringPlatform, "model": "FORTIGATE 3810A"},
		".1.3.6.1.4.1.12356.101.1.38101":  {"platform": stringPlatform, "model": "FORTIGATE 3810D"},
		".1.3.6.1.4.1.12356.101.1.38150":  {"platform": stringPlatform, "model": "FORTIGATE 3815D"},
		".1.3.6.1.4.1.12356.101.1.39500":  {"platform": stringPlatform, "model": "FORTIGATE 3950B"},
		".1.3.6.1.4.1.12356.101.1.39501":  {"platform": stringPlatform, "model": "FORTIGATE 3951B"},
		".1.3.6.1.4.1.12356.101.1.39601":  {"platform": stringPlatform, "model": "FORTIGATE 3960E"},
		".1.3.6.1.4.1.12356.101.1.39801":  {"platform": stringPlatform, "model": "FORTIGATE 3980E"},
		".1.3.6.1.4.1.12356.101.1.40":     {"platform": stringPlatform, "model": "FORTIGATE VM64 XEN"},
		".1.3.6.1.4.1.12356.101.1.4000":   {"platform": stringPlatform, "model": "FORTIGATE 400"},
		".1.3.6.1.4.1.12356.101.1.40000":  {"platform": stringPlatform, "model": "FORTIGATE 4000"},
		".1.3.6.1.4.1.12356.101.1.4001":   {"platform": stringPlatform, "model": "FORTIGATE 400A"},
		".1.3.6.1.4.1.12356.101.1.4004":   {"platform": stringPlatform, "model": "FORTIGATE 400D"},
		".1.3.6.1.4.1.12356.101.1.4008":   {"platform": stringPlatform, "model": "FORTIGATE 401E"},
		".1.3.6.1.4.1.12356.101.1.410":    {"platform": stringPlatform, "model": "FORTIGATE 40C"},
		".1.3.6.1.4.1.12356.101.1.411":    {"platform": stringPlatform, "model": "FORTIWIFI 40C"},
		".1.3.6.1.4.1.12356.101.1.443":    {"platform": stringPlatform, "model": "FORTIGATE 40F"},
		".1.3.6.1.4.1.12356.101.1.45":     {"platform": stringPlatform, "model": "FORTIGATE VM64 AWS"},
		".1.3.6.1.4.1.12356.101.1.46":     {"platform": stringPlatform, "model": "FORTIGATE VM64 AWS ONDEMAND"},
		".1.3.6.1.4.1.12356.101.1.500":    {"platform": stringPlatform, "model": "FORTIGATE FGT50A"},
		".1.3.6.1.4.1.12356.101.1.5000":   {"platform": stringPlatform, "model": "FORTIGATE 500"},
		".1.3.6.1.4.1.12356.101.1.50000":  {"platform": stringPlatform, "model": "FORTIGATE 5000"},
		".1.3.6.1.4.1.12356.101.1.50001":  {"platform": stringPlatform, "model": "FORTIGATE 5002FB2"},
		".1.3.6.1.4.1.12356.101.1.5001":   {"platform": stringPlatform, "model": "FORTIGATE 500A"},
		".1.3.6.1.4.1.12356.101.1.50010":  {"platform": stringPlatform, "model": "FORTIGATE 5001"},
		".1.3.6.1.4.1.12356.101.1.50011":  {"platform": stringPlatform, "model": "FORTIGATE 5001A"},
		".1.3.6.1.4.1.12356.101.1.50012":  {"platform": stringPlatform, "model": "FORTIGATE 5001FA2"},
		".1.3.6.1.4.1.12356.101.1.50013":  {"platform": stringPlatform, "model": "FORTIGATE 5001B"},
		".1.3.6.1.4.1.12356.101.1.50014":  {"platform": stringPlatform, "model": "FORTIGATE 5001C"},
		".1.3.6.1.4.1.12356.101.1.50015":  {"platform": stringPlatform, "model": "FORTIGATE 5001D"},
		".1.3.6.1.4.1.12356.101.1.50021":  {"platform": stringPlatform, "model": "FORTIGATE 5002A"},
		".1.3.6.1.4.1.12356.101.1.50023":  {"platform": stringPlatform, "model": "FORTISWITCH 5023B"},
		".1.3.6.1.4.1.12356.101.1.5004":   {"platform": stringPlatform, "model": "FORTIGATE 500D"},
		".1.3.6.1.4.1.12356.101.1.50040":  {"platform": stringPlatform, "model": "FORTIGATE 5004"},
		".1.3.6.1.4.1.12356.101.1.5005":   {"platform": stringPlatform, "model": "FORTIGATE 500E"},
		".1.3.6.1.4.1.12356.101.1.50050":  {"platform": stringPlatform, "model": "FORTIGATE 5005"},
		".1.3.6.1.4.1.12356.101.1.50051":  {"platform": stringPlatform, "model": "FORTIGATE 5005FA2"},
		".1.3.6.1.4.1.12356.101.1.5006":   {"platform": stringPlatform, "model": "FORTIGATE 501E"},
		".1.3.6.1.4.1.12356.101.1.501":    {"platform": stringPlatform, "model": "FORTIGATE FGT50AM"},
		".1.3.6.1.4.1.12356.101.1.502":    {"platform": stringPlatform, "model": "FORTIGATE 50B"},
		".1.3.6.1.4.1.12356.101.1.503":    {"platform": stringPlatform, "model": "FORTIWIFI 50B"},
		".1.3.6.1.4.1.12356.101.1.504":    {"platform": stringPlatform, "model": "FORTIGATE 51B"},
		".1.3.6.1.4.1.12356.101.1.505":    {"platform": stringPlatform, "model": "FORTIGATE 50E"},
		".1.3.6.1.4.1.12356.101.1.506":    {"platform": stringPlatform, "model": "FORTIWIFI 50E"},
		".1.3.6.1.4.1.12356.101.1.510":    {"platform": stringPlatform, "model": "FORTIWIFI 50B"},
		".1.3.6.1.4.1.12356.101.1.51010":  {"platform": stringPlatform, "model": "FORTIGATE 5101C"},
		".1.3.6.1.4.1.12356.101.1.515":    {"platform": stringPlatform, "model": "FORTIGATE 51E"},
		".1.3.6.1.4.1.12356.101.1.516":    {"platform": stringPlatform, "model": "FORTIWIFI 51E"},
		".1.3.6.1.4.1.12356.101.1.60":     {"platform": stringPlatform, "model": "FORTIGATE VM64 KVM"},
		".1.3.6.1.4.1.12356.101.1.600":    {"platform": stringPlatform, "model": "FORTIGATE 60"},
		".1.3.6.1.4.1.12356.101.1.6003":   {"platform": stringPlatform, "model": "FORTIGATE 600C"},
		".1.3.6.1.4.1.12356.101.1.6004":   {"platform": stringPlatform, "model": "FORTIGATE 600D"},
		".1.3.6.1.4.1.12356.101.1.6005":   {"platform": stringPlatform, "model": "FORTIGATE 600E"},
		".1.3.6.1.4.1.12356.101.1.6006":   {"platform": stringPlatform, "model": "FORTIGATE 601E"},
		".1.3.6.1.4.1.12356.101.1.601":    {"platform": stringPlatform, "model": "FORTIGATE 60M"},
		".1.3.6.1.4.1.12356.101.1.602":    {"platform": stringPlatform, "model": "FORTIGATE 60ADSL"},
		".1.3.6.1.4.1.12356.101.1.603":    {"platform": stringPlatform, "model": "FORTIGATE 60B"},
		".1.3.6.1.4.1.12356.101.1.605":    {"platform": stringPlatform, "model": "FORTIGATE 60C"},
		".1.3.6.1.4.1.12356.101.1.61":     {"platform": stringPlatform, "model": "FORTIGATE VM64NPU"},
		".1.3.6.1.4.1.12356.101.1.610":    {"platform": stringPlatform, "model": "FORTIWIFI 60"},
		".1.3.6.1.4.1.12356.101.1.611":    {"platform": stringPlatform, "model": "FORTIWIFI 60A"},
		".1.3.6.1.4.1.12356.101.1.612":    {"platform": stringPlatform, "model": "FORTIWIFI 60AM"},
		".1.3.6.1.4.1.12356.101.1.613":    {"platform": stringPlatform, "model": "FORTIWIFI 60B"},
		".1.3.6.1.4.1.12356.101.1.615":    {"platform": stringPlatform, "model": "FORTIGATE 60C"},
		".1.3.6.1.4.1.12356.101.1.616":    {"platform": stringPlatform, "model": "FORTIWIFI 60CM"},
		".1.3.6.1.4.1.12356.101.1.617":    {"platform": stringPlatform, "model": "FORTIWIFI 60CA"},
		".1.3.6.1.4.1.12356.101.1.618":    {"platform": stringPlatform, "model": "FORTIWIFI 60CB"},
		".1.3.6.1.4.1.12356.101.1.619":    {"platform": stringPlatform, "model": "FORTIWIFI 6XMB"},
		".1.3.6.1.4.1.12356.101.1.6200":   {"platform": stringPlatform, "model": "FORTIGATE 620B"},
		".1.3.6.1.4.1.12356.101.1.6201":   {"platform": stringPlatform, "model": "FORTIGATE 600D"},
		".1.3.6.1.4.1.12356.101.1.621":    {"platform": stringPlatform, "model": "FORTIGATE 60CP"},
		".1.3.6.1.4.1.12356.101.1.6210":   {"platform": stringPlatform, "model": "FORTIGATE 621B"},
		".1.3.6.1.4.1.12356.101.1.622":    {"platform": stringPlatform, "model": "FORTIGATE 60CSFP"},
		".1.3.6.1.4.1.12356.101.1.624":    {"platform": stringPlatform, "model": "FORTIGATE 60D"},
		".1.3.6.1.4.1.12356.101.1.625":    {"platform": stringPlatform, "model": "FORTIGATE 60D"},
		".1.3.6.1.4.1.12356.101.1.626":    {"platform": stringPlatform, "model": "FORTIWIFI 60D"},
		".1.3.6.1.4.1.12356.101.1.627":    {"platform": stringPlatform, "model": "FORTIWIFI 60DP"},
		".1.3.6.1.4.1.12356.101.1.628":    {"platform": stringPlatform, "model": "FORTIGATE SOC3"},
		".1.3.6.1.4.1.12356.101.1.630":    {"platform": stringPlatform, "model": "FORTIGATE 90D"},
		".1.3.6.1.4.1.12356.101.1.631":    {"platform": stringPlatform, "model": "FORTIGATE 90DPOE"},
		".1.3.6.1.4.1.12356.101.1.632":    {"platform": stringPlatform, "model": "FORTIWIFI 90D"},
		".1.3.6.1.4.1.12356.101.1.633":    {"platform": stringPlatform, "model": "FORTIWIFI 90DPOE"},
		".1.3.6.1.4.1.12356.101.1.634":    {"platform": stringPlatform, "model": "FORTIGATE 94DPOE"},
		".1.3.6.1.4.1.12356.101.1.635":    {"platform": stringPlatform, "model": "FORTIGATE 98DPOE"},
		".1.3.6.1.4.1.12356.101.1.636":    {"platform": stringPlatform, "model": "FORTIGATE 92D"},
		".1.3.6.1.4.1.12356.101.1.637":    {"platform": stringPlatform, "model": "FORTIWIFI 92D"},
		".1.3.6.1.4.1.12356.101.1.638":    {"platform": stringPlatform, "model": "FORTIGATE RUGGED 90D"},
		".1.3.6.1.4.1.12356.101.1.639":    {"platform": stringPlatform, "model": "FORTIWIFI 60E"},
		".1.3.6.1.4.1.12356.101.1.640":    {"platform": stringPlatform, "model": "FORTIGATE 61E"},
		".1.3.6.1.4.1.12356.101.1.641":    {"platform": stringPlatform, "model": "FORTIGATE 60E"},
		".1.3.6.1.4.1.12356.101.1.643":    {"platform": stringPlatform, "model": "FORTIGATE RUGGED 60D"},
		".1.3.6.1.4.1.12356.101.1.644":    {"platform": stringPlatform, "model": "FORTIGATE RUGGED 60F"},
		".1.3.6.1.4.1.12356.101.1.645":    {"platform": stringPlatform, "model": "FORTIGATE 61F"},
		".1.3.6.1.4.1.12356.101.1.649":    {"platform": stringPlatform, "model": "FORTIWIFI 61E"},
		".1.3.6.1.4.1.12356.101.1.65":     {"platform": stringPlatform, "model": "FORTIGATE VM64GCP"},
		".1.3.6.1.4.1.12356.101.1.70":     {"platform": stringPlatform, "model": "FORTIGATE VM64 HW"},
		".1.3.6.1.4.1.12356.101.1.700":    {"platform": stringPlatform, "model": "FORTIGATE 70D"},
		".1.3.6.1.4.1.12356.101.1.701":    {"platform": stringPlatform, "model": "FORTIGATE 70DPOE"},
		".1.3.6.1.4.1.12356.101.1.800":    {"platform": stringPlatform, "model": "FORTIGATE 80C"},
		".1.3.6.1.4.1.12356.101.1.8000":   {"platform": stringPlatform, "model": "FORTIGATE 800"},
		".1.3.6.1.4.1.12356.101.1.8001":   {"platform": stringPlatform, "model": "FORTIGATE 800F"},
		".1.3.6.1.4.1.12356.101.1.8003":   {"platform": stringPlatform, "model": "FORTIGATE 800C"},
		".1.3.6.1.4.1.12356.101.1.8004":   {"platform": stringPlatform, "model": "FORTIGATE 800D"},
		".1.3.6.1.4.1.12356.101.1.801":    {"platform": stringPlatform, "model": "FORTIGATE 80CM"},
		".1.3.6.1.4.1.12356.101.1.802":    {"platform": stringPlatform, "model": "FORTIGATE 82C"},
		".1.3.6.1.4.1.12356.101.1.803":    {"platform": stringPlatform, "model": "FORTIGATE 80D"},
		".1.3.6.1.4.1.12356.101.1.810":    {"platform": stringPlatform, "model": "FORTIWIFI 80CM"},
		".1.3.6.1.4.1.12356.101.1.811":    {"platform": stringPlatform, "model": "FORTIWIFI 81CM"},
		".1.3.6.1.4.1.12356.101.1.842":    {"platform": stringPlatform, "model": "FORTIGATE 80E"},
		".1.3.6.1.4.1.12356.101.1.843":    {"platform": stringPlatform, "model": "FORTIGATE 81E"},
		".1.3.6.1.4.1.12356.101.1.844":    {"platform": stringPlatform, "model": "FORTIGATE 81POE"},
		".1.3.6.1.4.1.12356.101.1.846":    {"platform": stringPlatform, "model": "FORTIGATE 80F"},
		".1.3.6.1.4.1.12356.101.1.900":    {"platform": stringPlatform, "model": "FORTIGATE 900D"},
		".1.3.6.1.4.1.12356.101.1.90000":  {"platform": stringPlatform, "model": "FORTIGATE FOSVM64"},
		".1.3.6.1.4.1.12356.101.1.90010":  {"platform": stringPlatform, "model": "FORTIGATE AZURE VM"},
		".1.3.6.1.4.1.12356.101.1.90018":  {"platform": stringPlatform, "model": "FORTIGATE VM64GCPONDEMAND"},
		".1.3.6.1.4.1.12356.101.1.90019":  {"platform": stringPlatform, "model": "FORTIGATE VM64ALI"},
		".1.3.6.1.4.1.12356.101.1.90020":  {"platform": stringPlatform, "model": "FORTIGATE VM64ALIONDEMAND"},
		".1.3.6.1.4.1.12356.101.1.90060":  {"platform": stringPlatform, "model": "FORTIGATE FOSVM64KVM"},
		".1.3.6.1.4.1.12356.101.1.90081":  {"platform": stringPlatform, "model": "FORTIGATE VM64"},
		".1.3.6.1.4.1.12356.101.1.940":    {"platform": stringPlatform, "model": "FORTIGATE 90E"},
		".1.3.6.1.4.1.12356.101.1.941":    {"platform": stringPlatform, "model": "FORTIGATE 91E"},
		".1.3.6.1.4.1.12356.102.1.1000":   {"platform": stringPlatform, "model": "FORTIANALYZER 100"},
		".1.3.6.1.4.1.12356.102.1.10001":  {"platform": stringPlatform, "model": "FAZ1000B"},
		".1.3.6.1.4.1.12356.102.1.10002":  {"platform": stringPlatform, "model": "FORTIANALYZER 1000B"},
		".1.3.6.1.4.1.12356.102.1.1001":   {"platform": stringPlatform, "model": "FORTIANALYZER 100A"},
		".1.3.6.1.4.1.12356.102.1.1002":   {"platform": stringPlatform, "model": "FORTIANALYZER 100B"},
		".1.3.6.1.4.1.12356.102.1.1003":   {"platform": stringPlatform, "model": "FORTIANALYZER 100C"},
		".1.3.6.1.4.1.12356.102.1.20000":  {"platform": stringPlatform, "model": "FORTIANALYZER 2000"},
		".1.3.6.1.4.1.12356.102.1.20001":  {"platform": stringPlatform, "model": "FORTIANALYZER 2000A"},
		".1.3.6.1.4.1.12356.102.1.20002":  {"platform": stringPlatform, "model": "FORTIANALYZER-2000B"},
		".1.3.6.1.4.1.12356.102.1.30005":  {"platform": stringPlatform, "model": "FORTIANALYZER 3000E"},
		".1.3.6.1.4.1.12356.102.1.4000":   {"platform": stringPlatform, "model": "FORTIANALYZER 400"},
		".1.3.6.1.4.1.12356.102.1.40000":  {"platform": stringPlatform, "model": "FORTIANALYZER 4000"},
		".1.3.6.1.4.1.12356.102.1.40001":  {"platform": stringPlatform, "model": "FORTIANALYZER 4000A"},
		".1.3.6.1.4.1.12356.102.1.40002":  {"platform": stringPlatform, "model": "FORTIANALYZER 4000B"},
		".1.3.6.1.4.1.12356.102.1.4002":   {"platform": stringPlatform, "model": "FORTIANALYZER 400B"},
		".1.3.6.1.4.1.12356.102.1.8000":   {"platform": stringPlatform, "model": "FORTIANALYZER 800"},
		".1.3.6.1.4.1.12356.102.1.8002":   {"platform": stringPlatform, "model": "FORTIANALYZER 800B"},
		".1.3.6.1.4.1.12356.103.1.1000":   {"platform": stringPlatform, "model": "FORTIMANAGER 100"},
		".1.3.6.1.4.1.12356.103.1.10003":  {"platform": stringPlatform, "model": "FORTIMANAGER 1000C"},
		".1.3.6.1.4.1.12356.103.1.10004":  {"platform": stringPlatform, "model": "FORTIGATE 1000D"},
		".1.3.6.1.4.1.12356.103.1.10006":  {"platform": stringPlatform, "model": "FORTIMANAGER 1000F"},
		".1.3.6.1.4.1.12356.103.1.1001":   {"platform": stringPlatform, "model": "FORTIMANAGER VM"},
		".1.3.6.1.4.1.12356.103.1.1003":   {"platform": stringPlatform, "model": "FORTIMANAGER 100C"},
		".1.3.6.1.4.1.12356.103.1.20000":  {"platform": stringPlatform, "model": "FORTIMANAGER 2000XL"},
		".1.3.6.1.4.1.12356.103.1.20005":  {"platform": stringPlatform, "model": "FORTIMANAGER 2000E"},
		".1.3.6.1.4.1.12356.103.1.2004":   {"platform": stringPlatform, "model": "FORTIMANAGER 200D"},
		".1.3.6.1.4.1.12356.103.1.2005":   {"platform": stringPlatform, "model": "FORTIMANAGER 200E"},
		".1.3.6.1.4.1.12356.103.1.2006":   {"platform": stringPlatform, "model": "FORTIMANAGER 200F"},
		".1.3.6.1.4.1.12356.103.1.30000":  {"platform": stringPlatform, "model": "FORTIMANAGER 3000"},
		".1.3.6.1.4.1.12356.103.1.30002":  {"platform": stringPlatform, "model": "FORTIMANAGER 3000B"},
		".1.3.6.1.4.1.12356.103.1.30003":  {"platform": stringPlatform, "model": "FORTIMANAGER 3000C"},
		".1.3.6.1.4.1.12356.103.1.30006":  {"platform": stringPlatform, "model": "FORTIMANAGER 3000F"},
		".1.3.6.1.4.1.12356.103.1.3004":   {"platform": stringPlatform, "model": "FORTIMANAGER 300D"},
		".1.3.6.1.4.1.12356.103.1.3005":   {"platform": stringPlatform, "model": "FORTIMANAGER 300E"},
		".1.3.6.1.4.1.12356.103.1.3006":   {"platform": stringPlatform, "model": "FORTIMANAGER 300F"},
		".1.3.6.1.4.1.12356.103.1.30905":  {"platform": stringPlatform, "model": "FORTIMANAGER 3900E"},
		".1.3.6.1.4.1.12356.103.1.39005":  {"platform": stringPlatform, "model": "FORTIMANAGER 3900E"},
		".1.3.6.1.4.1.12356.103.1.4000":   {"platform": stringPlatform, "model": "FORTIMANAGER 400"},
		".1.3.6.1.4.1.12356.103.1.40004":  {"platform": stringPlatform, "model": "FORTIMANAGER 4000D"},
		".1.3.6.1.4.1.12356.103.1.40005":  {"platform": stringPlatform, "model": "FORTIMANAGER 4000E"},
		".1.3.6.1.4.1.12356.103.1.4001":   {"platform": stringPlatform, "model": "FORTIMANAGER 400A"},
		".1.3.6.1.4.1.12356.103.1.4002":   {"platform": stringPlatform, "model": "FORTIMANAGER 400B"},
		".1.3.6.1.4.1.12356.103.1.4003":   {"platform": stringPlatform, "model": "FORTIMANAGER 400C"},
		".1.3.6.1.4.1.12356.103.1.4005":   {"platform": stringPlatform, "model": "FORTIMANAGER 400E"},
		".1.3.6.1.4.1.12356.103.1.50011":  {"platform": stringPlatform, "model": "FORTIMANAGER 5001A"},
		".1.3.6.1.4.1.12356.103.1.64":     {"platform": stringPlatform, "model": "FORTIMANAGER"},
		".1.3.6.1.4.1.12356.103.3.1000":   {"platform": stringPlatform, "model": "FORTIANALYZER 100"},
		".1.3.6.1.4.1.12356.103.3.10002":  {"platform": stringPlatform, "model": "FORTIANALYZER 1000B"},
		".1.3.6.1.4.1.12356.103.3.10003":  {"platform": stringPlatform, "model": "FORTIANALYZER 1000C"},
		".1.3.6.1.4.1.12356.103.3.10004":  {"platform": stringPlatform, "model": "FORTIANALYZER 1000D"},
		".1.3.6.1.4.1.12356.103.3.10005":  {"platform": stringPlatform, "model": "FORTIANALYZER 1000E"},
		".1.3.6.1.4.1.12356.103.3.10006":  {"platform": stringPlatform, "model": "FORTIANALYZER 1000F"},
		".1.3.6.1.4.1.12356.103.3.1001":   {"platform": stringPlatform, "model": "FORTIANALYZER 100A"},
		".1.3.6.1.4.1.12356.103.3.1002":   {"platform": stringPlatform, "model": "FORTIANALYZER 100B"},
		".1.3.6.1.4.1.12356.103.3.1003":   {"platform": stringPlatform, "model": "FORTIANALYZER 100C"},
		".1.3.6.1.4.1.12356.103.3.20":     {"platform": stringPlatform, "model": "FORTIANALYZER VM"},
		".1.3.6.1.4.1.12356.103.3.20000":  {"platform": stringPlatform, "model": "FORTIANALYZER 2000"},
		".1.3.6.1.4.1.12356.103.3.20001":  {"platform": stringPlatform, "model": "FORTIANALYZER 2000A"},
		".1.3.6.1.4.1.12356.103.3.20002":  {"platform": stringPlatform, "model": "FORTIANALYZER 2000B"},
		".1.3.6.1.4.1.12356.103.3.20005":  {"platform": stringPlatform, "model": "FORTIANALYZER 2000E"},
		".1.3.6.1.4.1.12356.103.3.2004":   {"platform": stringPlatform, "model": "FORTIANALYZER 200D"},
		".1.3.6.1.4.1.12356.103.3.2005":   {"platform": stringPlatform, "model": "FORTIANALYZER 200E"},
		".1.3.6.1.4.1.12356.103.3.2006":   {"platform": stringPlatform, "model": "FORTIANALYZER 200F"},
		".1.3.6.1.4.1.12356.103.3.30004":  {"platform": stringPlatform, "model": "FORTIANALYZER 3000D"},
		".1.3.6.1.4.1.12356.103.3.30005":  {"platform": stringPlatform, "model": "FORTIANALYZER 3000E"},
		".1.3.6.1.4.1.12356.103.3.30006":  {"platform": stringPlatform, "model": "FORTIANALYZER 3000F"},
		".1.3.6.1.4.1.12356.103.3.3004":   {"platform": stringPlatform, "model": "FORTIANALYZER 300D"},
		".1.3.6.1.4.1.12356.103.3.3006":   {"platform": stringPlatform, "model": "FORTIANALYZER 300F"},
		".1.3.6.1.4.1.12356.103.3.35005":  {"platform": stringPlatform, "model": "FORTIGATE 3500E"},
		".1.3.6.1.4.1.12356.103.3.35006":  {"platform": stringPlatform, "model": "FORTIANALYZER 3500F"},
		".1.3.6.1.4.1.12356.103.3.37006":  {"platform": stringPlatform, "model": "FORTIANALYZER 3700F"},
		".1.3.6.1.4.1.12356.103.3.39005":  {"platform": stringPlatform, "model": "FORTIANALYZER 3900E"},
		".1.3.6.1.4.1.12356.103.3.4000":   {"platform": stringPlatform, "model": "FORTIANALYZER 400"},
		".1.3.6.1.4.1.12356.103.3.40000":  {"platform": stringPlatform, "model": "FORTIANALYZER 4000"},
		".1.3.6.1.4.1.12356.103.3.40001":  {"platform": stringPlatform, "model": "FORTIANALYZER 4000A"},
		".1.3.6.1.4.1.12356.103.3.40002":  {"platform": stringPlatform, "model": "FORTIANALYZER 4000B"},
		".1.3.6.1.4.1.12356.103.3.4002":   {"platform": stringPlatform, "model": "FORTIANALYZER 400B"},
		".1.3.6.1.4.1.12356.103.3.4003":   {"platform": stringPlatform, "model": "FORTIANALYZER 400C"},
		".1.3.6.1.4.1.12356.103.3.4005":   {"platform": stringPlatform, "model": "FORTIANALYZER 400E"},
		".1.3.6.1.4.1.12356.103.3.64":     {"platform": stringPlatform, "model": "FORTIANALYZER VM"},
		".1.3.6.1.4.1.12356.103.3.8000":   {"platform": stringPlatform, "model": "FORTIANALYZER 800"},
		".1.3.6.1.4.1.12356.103.3.8002":   {"platform": stringPlatform, "model": "FORTIANALYZER 800B"},
		".1.3.6.1.4.1.12356.103.3.8006":   {"platform": stringPlatform, "model": "FORTIANALYZER 800F"},
		".1.3.6.1.4.1.12356.105":          {"platform": stringPlatform, "model": "FORTIMAIL-VM-KVM"},
		".1.3.6.1.4.1.12356.106.1.10241":  {"platform": stringPlatform, "model": "FORTISWITCH1024D"},
		".1.3.6.1.4.1.12356.106.1.10481":  {"platform": stringPlatform, "model": "FORTISWITCH1048D"},
		".1.3.6.1.4.1.12356.106.1.1083":   {"platform": stringPlatform, "model": "FORTISWITCH 108E-POE"},
		".1.3.6.1.4.1.12356.106.1.2242":   {"platform": stringPlatform, "model": "FORTISWITCH 224D-FPOE"},
		".1.3.6.1.4.1.12356.106.1.50030":  {"platform": stringPlatform, "model": "FORTISWITCH 5003A"},
		".1.3.6.1.4.1.12356.106.1.50031":  {"platform": stringPlatform, "model": "FORTISWITCH 5003B"},
		".1.3.6.1.4.1.12356.106.1.51030":  {"platform": stringPlatform, "model": "FORTISWITCH5003A-CONTROLLER"},
		".1.3.6.1.4.1.12356.107.1.10002":  {"platform": stringPlatform, "model": "FORTIWEB 1000B"},
		".1.3.6.1.4.1.12356.107.1.10003":  {"platform": stringPlatform, "model": "FORTIWEB 1000C"},
		".1.3.6.1.4.1.12356.107.1.10004":  {"platform": stringPlatform, "model": "FORTIWEB 1000D"},
		".1.3.6.1.4.1.12356.107.1.10005":  {"platform": stringPlatform, "model": "FORTIWEB 2000E"},
		".1.3.6.1.4.1.12356.107.1.10006":  {"platform": stringPlatform, "model": "FORTIWEB 1000E"},
		".1.3.6.1.4.1.12356.107.1.1004":   {"platform": stringPlatform, "model": "FORTIWEB 100D"},
		".1.3.6.1.4.1.12356.107.1.30003":  {"platform": stringPlatform, "model": "FORTIWEB 3000C"},
		".1.3.6.1.4.1.12356.107.1.30004":  {"platform": stringPlatform, "model": "FORTIWEB 3000C FSX"},
		".1.3.6.1.4.1.12356.107.1.30005":  {"platform": stringPlatform, "model": "FORTIWEB 3000D"},
		".1.3.6.1.4.1.12356.107.1.30006":  {"platform": stringPlatform, "model": "FORTIWEB 3000D-FSX"},
		".1.3.6.1.4.1.12356.107.1.30007":  {"platform": stringPlatform, "model": "FORTIWEB 3000E"},
		".1.3.6.1.4.1.12356.107.1.30008":  {"platform": stringPlatform, "model": "FORTIWEB 3010E"},
		".1.3.6.1.4.1.12356.107.1.40003":  {"platform": stringPlatform, "model": "FORTIWEB 4000C"},
		".1.3.6.1.4.1.12356.107.1.40004":  {"platform": stringPlatform, "model": "FORTIWEB 4000D"},
		".1.3.6.1.4.1.12356.107.1.40005":  {"platform": stringPlatform, "model": "FORTIWEB 4000E"},
		".1.3.6.1.4.1.12356.107.1.4002":   {"platform": stringPlatform, "model": "FORTIWEB 400B"},
		".1.3.6.1.4.1.12356.107.1.4003":   {"platform": stringPlatform, "model": "FORTIWEB 400C"},
		".1.3.6.1.4.1.12356.107.1.4004":   {"platform": stringPlatform, "model": "FORTIWEB 400D"},
		".1.3.6.1.4.1.12356.107.1.50001":  {"platform": stringPlatform, "model": "FORTIWEB VM"},
		".1.3.6.1.4.1.12356.107.1.50002":  {"platform": stringPlatform, "model": "FORTIWEB XENOPENSOURCE"},
		".1.3.6.1.4.1.12356.107.1.50003":  {"platform": stringPlatform, "model": "FORTIWEB XENSERVER"},
		".1.3.6.1.4.1.12356.107.1.50004":  {"platform": stringPlatform, "model": "FORTIWEB XENAWS"},
		".1.3.6.1.4.1.12356.107.1.50005":  {"platform": stringPlatform, "model": "FORTIWEB HYPERV"},
		".1.3.6.1.4.1.12356.107.1.50006":  {"platform": stringPlatform, "model": "FORTIWEB KVM"},
		".1.3.6.1.4.1.12356.107.1.50007":  {"platform": stringPlatform, "model": "FORTIWEB AZURE"},
		".1.3.6.1.4.1.12356.107.1.50008":  {"platform": stringPlatform, "model": "FORTIWEB VM PAYG"},
		".1.3.6.1.4.1.12356.107.1.50009":  {"platform": stringPlatform, "model": "FORTIWEB KVM PAYG"},
		".1.3.6.1.4.1.12356.107.1.6004":   {"platform": stringPlatform, "model": "FORTIWEB 600D"},
		".1.3.6.1.4.1.12356.111":          {"platform": stringPlatform, "model": "FORTIDDOS 900B"},
		".1.3.6.1.4.1.12356.112.100.10":   {"platform": stringPlatform, "model": "FORTIADC DEV"},
		".1.3.6.1.4.1.12356.112.100.1001": {"platform": stringPlatform, "model": "FORTIADC 1000D"},
		".1.3.6.1.4.1.12356.112.100.1003": {"platform": stringPlatform, "model": "FORTIADC 1000F"},
		".1.3.6.1.4.1.12356.112.100.101":  {"platform": stringPlatform, "model": "FORTIADC 100D"},
		".1.3.6.1.4.1.12356.112.100.103":  {"platform": stringPlatform, "model": "FORTIADC 100F"},
		".1.3.6.1.4.1.12356.112.100.1203": {"platform": stringPlatform, "model": "FORTIADC 1200F"},
		".1.3.6.1.4.1.12356.112.100.1501": {"platform": stringPlatform, "model": "FORTIADC 1500D"},
		".1.3.6.1.4.1.12356.112.100.20":   {"platform": stringPlatform, "model": "FORTIADC KVM"},
		".1.3.6.1.4.1.12356.112.100.2001": {"platform": stringPlatform, "model": "FORTIADC 2000D"},
		".1.3.6.1.4.1.12356.112.100.2003": {"platform": stringPlatform, "model": "FORTIADC 2000F"},
		".1.3.6.1.4.1.12356.112.100.201":  {"platform": stringPlatform, "model": "FORTIADC 200D"},
		".1.3.6.1.4.1.12356.112.100.203":  {"platform": stringPlatform, "model": "FORTIADC 200F"},
		".1.3.6.1.4.1.12356.112.100.2203": {"platform": stringPlatform, "model": "FORTIADC 2200F"},
		".1.3.6.1.4.1.12356.112.100.30":   {"platform": stringPlatform, "model": "FORTIADC VM"},
		".1.3.6.1.4.1.12356.112.100.301":  {"platform": stringPlatform, "model": "FORTIADC 300D"},
		".1.3.6.1.4.1.12356.112.100.302":  {"platform": stringPlatform, "model": "FORTIADC 300E"},
		".1.3.6.1.4.1.12356.112.100.303":  {"platform": stringPlatform, "model": "FORTIADC 300F"},
		".1.3.6.1.4.1.12356.112.100.4001": {"platform": stringPlatform, "model": "FORTIADC 4000D"},
		".1.3.6.1.4.1.12356.112.100.4003": {"platform": stringPlatform, "model": "FORTIADC 4000F"},
		".1.3.6.1.4.1.12356.112.100.401":  {"platform": stringPlatform, "model": "FORTIADC 400D"},
		".1.3.6.1.4.1.12356.112.100.403":  {"platform": stringPlatform, "model": "FORTIADC 400F"},
		".1.3.6.1.4.1.12356.112.100.4203": {"platform": stringPlatform, "model": "FORTIADC 4200F"},
		".1.3.6.1.4.1.12356.112.100.5003": {"platform": stringPlatform, "model": "FORTIADC 5000F"},
		".1.3.6.1.4.1.12356.112.100.63":   {"platform": stringPlatform, "model": "FORTIADC 60F"},
		".1.3.6.1.4.1.12356.112.100.701":  {"platform": stringPlatform, "model": "FORTIADC 700D"},
		".1.3.6.1.4.1.12356.113.100.101":  {"platform": stringPlatform, "model": "FORTIAUTHENTICATOR"},
		".1.3.6.1.4.1.12356.118.1.30006":  {"platform": stringPlatform, "model": "FORTISANDBOX3000E"},
		".1.3.6.1.4.1.12356.118.1.30007":  {"platform": stringPlatform, "model": "FORTISANDBOX"},
		".1.3.6.1.4.1.12356.15.200":       {"platform": stringPlatform, "model": "FORTIGATE 200"},
		".1.3.6.1.4.1.12356.15.300":       {"platform": stringPlatform, "model": "FORTIGATE 300"},
		".1.3.6.1.4.1.12356.15.60":        {"platform": stringPlatform, "model": "FORTIGATE 60"},
		".1.3.6.1.4.1.12356.1688":         {"platform": stringPlatform, "model": "FORTIMAIL 2000A"},
		".1.3.6.1.4.1.12356.2001":         {"platform": stringPlatform, "model": "FORTIGATE 200A"},
		".1.3.6.1.4.1.12356.201":          {"platform": stringPlatform, "model": "FORTIGATE 200A"},
		".1.3.6.1.4.1.12356.300":          {"platform": stringPlatform, "model": "FORTIGATE 300"},
		".1.3.6.1.4.1.12356.3001":         {"platform": stringPlatform, "model": "FWDB"},
		".1.3.6.1.4.1.12356.3002":         {"platform": stringPlatform, "model": "FORTIGATE 310B"},
		".1.3.6.1.4.1.12356.301":          {"platform": stringPlatform, "model": "FORTIGATE 300A"},
		".1.3.6.1.4.1.12356.3600":         {"platform": stringPlatform, "model": "FORTIGATE 3600"},
		".1.3.6.1.4.1.12356.36000":        {"platform": stringPlatform, "model": "FGT-SSF-SSL"},
		".1.3.6.1.4.1.12356.40004":        {"platform": stringPlatform, "model": "FORTIWEB 4000D"},
		".1.3.6.1.4.1.12356.50051":        {"platform": stringPlatform, "model": "NETSOL"},
		".1.3.6.1.4.1.12356.600":          {"platform": stringPlatform, "model": "FG60-620B"},
		".1.3.6.1.4.1.12356.603":          {"platform": stringPlatform, "model": "FORTIGATE 60"},
		".1.3.6.1.4.1.12356.610":          {"platform": stringPlatform, "model": "FORTIWIFI 60"},
		".1.3.6.1.4.1.12356.611":          {"platform": stringPlatform, "model": "FORTIWIFI 60A"},
		".1.3.6.1.4.1.12356.612":          {"platform": stringPlatform, "model": "FORTIWIFI 60AM"},
		".1.3.6.1.4.1.12356.800":          {"platform": stringPlatform, "model": "FORTIGATE 800"},
		".1.3.6.1.4.1.12356.8001":         {"platform": stringPlatform, "model": "FORTIGATE 800F"},
		".1.3.6.1.4.1.12356.801":          {"platform": stringPlatform, "model": "FORTIGATE 800A"},
		".1.3.6.1.4.1.12356.101.1.441":    {"platform": stringPlatform, "model": "FORTIGATE 40F"},
	}

	data, ok := oidMap[sysObjId]
	if !ok {
		return &devicemodel.DeviceModel{
			Platform:     platform.FortiNet,
			Manufacturer: manufacturer.FortiNet,
			DeviceModel:  devicemodel.UnknownDeviceModel,
		}
	}

	return &devicemodel.DeviceModel{
		Platform:     platform.Platform(data["platform"]),
		Manufacturer: manufacturer.FortiNet,
		DeviceModel:  data["model"],
	}
}
