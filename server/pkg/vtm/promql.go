package vtm

import (
	"fmt"
	"reflect"
	"strings"
)

type ArithmeticBinOperator string

const (
	Addition       ArithmeticBinOperator = "+"
	Subtraction    ArithmeticBinOperator = "-"
	Multiplication ArithmeticBinOperator = "*"
	Division       ArithmeticBinOperator = "/"
	Modulo         ArithmeticBinOperator = "%"
	Exponent       ArithmeticBinOperator = "^"
)

type TrigonometricBinOperator string

const (
	Atan2 TrigonometricBinOperator = "atan2"
)

type ComparisonBinOperator string

const (
	Equal          ComparisonBinOperator = "="
	NotEqual       ComparisonBinOperator = "!="
	GreaterThan    ComparisonBinOperator = ">"
	GreaterOrEqual ComparisonBinOperator = ">="
	LessThan       ComparisonBinOperator = "<"
	LessOrEqual    ComparisonBinOperator = "<="
)

type LogicBinOperator string

const (
	And    LogicBinOperator = "and"
	Or     LogicBinOperator = "or"
	Unless LogicBinOperator = "unless"
)

type VectorMatching string

const (
	On       VectorMatching = "on"
	Ignoring VectorMatching = "ignoring"
)

type GroupModifier string

const (
	GroupLeft  GroupModifier = "group_left"
	GroupRight GroupModifier = "group_right"
)

type AggregationOperator string

const (
	Sum         AggregationOperator = "sum"
	Min         AggregationOperator = "min"
	Max         AggregationOperator = "max"
	Avg         AggregationOperator = "avg"
	Group       AggregationOperator = "group"
	Stddev      AggregationOperator = "stddev"
	Stdvar      AggregationOperator = "stdvar"
	Count       AggregationOperator = "count"
	CountValues AggregationOperator = "count_values"
	ButtomK     AggregationOperator = "bottomk"
	TopK        AggregationOperator = "topk"
	Quantile    AggregationOperator = "quantile"
)

type AggregationMethod string

const (
	GroupBy      AggregationMethod = "by"
	GroupWithout AggregationMethod = "without"
)

type Compare struct {
	Op    ComparisonBinOperator
	Value float64
}

type Aggregation struct {
	Op     AggregationOperator
	AggWay AggregationMethod
	By     []string
}

type BinaryOp struct {
	ArithmeticBinOperator ArithmeticBinOperator
}

type Matcher string

const (
	EqualMatcher    Matcher = "="
	NotEqualMatcher Matcher = "!="
	LikeMatcher     Matcher = "=~"
	NotLikeMatcher  Matcher = "!~"
)

type Label struct {
	Name  string
	Value string
	Matcher
}

type PromQLBuilder struct {
	metricName string
	funcName   string
	labels     []Label
	window     string
	offset     string
	compOps    []Compare
	agg        Aggregation
}

func NewPromQLBuilder(metricName string) *PromQLBuilder {
	return &PromQLBuilder{metricName: metricName}
}

// see prometheus functions: https://prometheus.io/docs/prometheus/latest/querying/functions/
// WithFuncName, basicQL will be format as "funcName(basicQL)"
func (pb *PromQLBuilder) WithFuncName(funcName string) *PromQLBuilder {
	pb.funcName = funcName
	return pb
}

// WithLabels basicQL will be format as `{labelName1}={labelValue1},{labelName2}=~{labelValue2|labelValue2x}...{labelNameN}!={labelValueN}`
// and append to metric name
func (pb *PromQLBuilder) WithLabels(labels Label) *PromQLBuilder {
	pb.labels = append(pb.labels, labels)

	return pb
}

// WithOffset basicQL will be format as "offset {offset}"
func (pb *PromQLBuilder) WithOffset(offset string) *PromQLBuilder {
	pb.offset = offset
	return pb
}

// WithWindow basicQL will be format as "[window]"
func (pb *PromQLBuilder) WithWindow(window string) *PromQLBuilder {
	pb.window = window
	return pb
}

// WithComp basicQL will be format as "comp1 and comp2 and ... and compN"
// when compare length > 1, combine with "and" as default
func (pb *PromQLBuilder) WithComp(comp Compare) *PromQLBuilder {
	pb.compOps = append(pb.compOps, comp)
	return pb
}

// WithAgg basicQL will be format as "agg(basicQL) by/without (labelName1,labelName2,...,labelNameN)"
func (pb *PromQLBuilder) WithAgg(agg Aggregation) *PromQLBuilder {
	pb.agg = agg
	return pb
}

// Build build basicQL
func (pb *PromQLBuilder) Build() (string, error) {
	if pb.metricName == "" {
		return "", fmt.Errorf("metric name is required")
	}
	basicQL :=pb.metricName
	if len(pb.labels) > 0 {
		tmp := make([]string, len(pb.labels))
		for index, label := range pb.labels {
			tmp[index] = fmt.Sprintf(`%s%s"%s"`, label.Name, label.Matcher, label.Value)
		}
		basicQL += fmt.Sprintf("{%s}", strings.Join(tmp, ","))
	}

	if pb.window != "" {
		basicQL += fmt.Sprintf("[%s]", pb.window)
	}
	if len(pb.compOps) > 0 {
		tmp := make([]string, len(pb.compOps))
		for index, comp := range pb.compOps {
			tmp[index] = fmt.Sprintf("%s%s%.0f", basicQL, comp.Op, comp.Value)
		}
		fmt.Println(tmp)
		fmt.Println(len(tmp))
		basicQL = fmt.Sprintf(" %s", strings.Join(tmp, " and "))

	}
	if pb.funcName != "" {
		basicQL = fmt.Sprintf("%s(%s)", pb.funcName, basicQL)
	}

	if pb.offset != "" {
		basicQL += fmt.Sprintf(" offset %s", pb.offset)
	}

	if !reflect.DeepEqual(pb.agg, Aggregation{}) {
		if pb.agg.Op != "" {
			basicQL = fmt.Sprintf("%s(%s)", pb.agg.Op, basicQL)
		}
		if pb.agg.AggWay != "" && pb.agg.By != nil {
			basicQL = fmt.Sprintf("%s %s (%s)", basicQL, pb.agg.AggWay, strings.Join(pb.agg.By, ","))
		}
	}
	return basicQL, nil
}
