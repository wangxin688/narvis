package biz

import (
	"fmt"
	"sort"

	"github.com/wangxin688/narvis/intend/devicerole"
	"github.com/wangxin688/narvis/server/features/topology/schemas"
	"github.com/wangxin688/narvis/server/models"
)

const lineTypeLine string = "line"
const lineTypeWireless string = "wireless"
const lineTypeCircuit string = "circuit"

func createDeviceNode(device *models.Device) *schemas.Node {
	dr := devicerole.GetDeviceRole(device.DeviceRole)
	return &schemas.Node{
		Id:           device.Id,
		Name:         device.Name,
		ManagementIp: device.ManagementIp,
		DeviceRole:   device.DeviceRole,
		Weight:       dr.Weight,
		Floor:        device.Floor,
	}
}

func createApNode(ap *models.AP) *schemas.Node {
	return &schemas.Node{
		Id:           ap.Id,
		Name:         ap.Name,
		ManagementIp: ap.ManagementIp,
		DeviceRole:   ap.DeviceRole,
		Weight:       0,
		Floor:        ap.Floor,
	}
}

func generateLineAlias(lldp *models.LLDPNeighbor) string {
	return fmt.Sprintf("%s->%s", lldp.LocalIfName, lldp.RemoteIfName)
}

func createBaseLine(lldp *models.LLDPNeighbor, localNode *schemas.Node, remoteNode *schemas.Node) *schemas.LineBase {
	if localNode.Weight > remoteNode.Weight {
		localNode, remoteNode = remoteNode, localNode
	}
	return &schemas.LineBase{
		Id:                       lldp.Id,
		LineAlias:                generateLineAlias(lldp),
		LineStatus:               "up",
		LocalDeviceId:            localNode.Id,
		LocalDeviceName:          localNode.Name,
		LocalDeviceManagementIp:  localNode.ManagementIp,
		LocalDeviceRole:          localNode.DeviceRole,
		LocalIfName:              lldp.LocalIfName,
		RemoteDeviceId:           remoteNode.Id,
		RemoteDeviceName:         remoteNode.Name,
		RemoteDeviceManagementIp: remoteNode.ManagementIp,
		RemoteDeviceRole:         remoteNode.DeviceRole,
		RemoteIfName:             &lldp.RemoteIfName,
	}
}

func createApLine(lldp *models.ApLLDPNeighbor, localNode *schemas.Node, remoteNode *schemas.Node) *schemas.LineBase {
	return &schemas.LineBase{
		Id:                       lldp.Id,
		LineAlias:                fmt.Sprintf("%s->%s", lldp.LocalIfName, remoteNode.Name),
		LineStatus:               "up",
		LocalDeviceId:            localNode.Id,
		LocalDeviceName:          localNode.Name,
		LocalDeviceManagementIp:  localNode.ManagementIp,
		LocalDeviceRole:          localNode.DeviceRole,
		LocalIfName:              lldp.LocalIfName,
		RemoteDeviceId:           remoteNode.Id,
		RemoteDeviceName:         remoteNode.Name,
		RemoteDeviceManagementIp: remoteNode.ManagementIp,
		RemoteDeviceRole:         remoteNode.DeviceRole,
	}
}

func deduplicateEdges(lldp []*models.LLDPNeighbor) []*models.LLDPNeighbor {
	// Create a set to store the unique edges
	uniqueEdges := make(map[string]struct{})
	result := make([]*models.LLDPNeighbor, 0)
	for _, edge := range lldp {
		// Sort the node IDs
		nodes := []string{edge.LocalDeviceId, edge.RemoteDeviceId}
		sort.Strings(nodes)
		edgeKey := nodes[0] + nodes[1]
		if _, ok := uniqueEdges[edgeKey]; !ok {
			// Add the edge to the set
			uniqueEdges[edgeKey] = struct{}{}
			result = append(result, edge)
		}
	}
	return result
}

// group by line bases by localDeviceId and remoteDeviceId
func groupByLineBases(lines []*schemas.LineBase) []*schemas.Line {
	groupedLines := make(map[string][]*schemas.LineBase)
	for _, line := range lines {
		nodes := []string{line.LocalDeviceId, line.RemoteDeviceId}
		sort.Strings(nodes)
		key := nodes[0] + nodes[1]
		groupedLines[key] = append(groupedLines[key], line)
	}
	result := make([]*schemas.Line, 0)
	for _, lines := range groupedLines {
		result = append(result, &schemas.Line{
			Source:   lines[0].LocalDeviceId,
			Target:   lines[0].RemoteDeviceId,
			Type:     lineTypeLine,
			LineInfo: lines,
		})
	}
	return result
}

func deviceLldpToGraph(lldp []*models.LLDPNeighbor, devices map[string]*models.Device) ([]*schemas.Node, []*schemas.Line) {
	nodes := make([]*schemas.Node, 0)
	baseLines := make([]*schemas.LineBase, 0)
	lines := make([]*schemas.Line, 0)
	if len(lldp) == 0 {
		return nodes, lines
	}
	for _, edge := range lldp {
		localNode := createDeviceNode(devices[edge.LocalDeviceId])
		remoteNode := createDeviceNode(devices[edge.RemoteDeviceId])
		line := createBaseLine(edge, localNode, remoteNode)
		nodes = append(nodes, localNode, remoteNode)
		baseLines = append(baseLines, line)
	}

	lines = groupByLineBases(baseLines)
	return nodes, lines
}

func apLldpToGraph(lldp []*models.ApLLDPNeighbor, devices map[string]*models.Device, aps map[string]*models.AP) ([]*schemas.Node, []*schemas.Line) {
	nodes := make([]*schemas.Node, 0)
	baseLines := make([]*schemas.LineBase, 0)
	lines := make([]*schemas.Line, 0)
	if len(lldp) == 0 {
		return nodes, lines
	}
	for _, edge := range lldp {
		localNode := createDeviceNode(devices[edge.LocalDeviceId])
		remoteNode := createApNode(aps[edge.RemoteApId])
		line := createApLine(edge, localNode, remoteNode)
		nodes = append(nodes, localNode, remoteNode)
		baseLines = append(baseLines, line)
	}

	lines = groupByLineBases(baseLines)
	return nodes, lines
}

// // if combo needed for aggregation devices, add it in future
// func getFloor(nodes []*schemas.Node) int {

// }
