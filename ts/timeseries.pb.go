// Code generated by protoc-gen-gogo.
// source: cockroach/ts/timeseries.proto
// DO NOT EDIT!

/*
	Package ts is a generated protocol buffer package.

	It is generated from these files:
		cockroach/ts/timeseries.proto

	It has these top-level messages:
		TimeSeriesDatapoint
		TimeSeriesData
		Query
		TimeSeriesQueryRequest
		TimeSeriesQueryResponse
*/
package ts

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"

// skipping weak import gogoproto "github.com/cockroachdb/gogoproto"

import io "io"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// TimeSeriesQueryAggregator describes a set of aggregation functions which can
// be used to combine multiple datapoints into a single datapoint.
//
// Aggregators are used to "downsample" series by combining datapoints from the
// same series at different times. They are also used to "aggregate" values from
// different series, combining data points from different series at the same
// time.
type TimeSeriesQueryAggregator int32

const (
	// AVG returns the average value of datapoints.
	TimeSeriesQueryAggregator_AVG TimeSeriesQueryAggregator = 1
	// SUM returns the sum value of datapoints.
	TimeSeriesQueryAggregator_SUM TimeSeriesQueryAggregator = 2
	// MAX returns the maximum value of datapoints.
	TimeSeriesQueryAggregator_MAX TimeSeriesQueryAggregator = 3
	// MIN returns the minimum value of datapoints.
	TimeSeriesQueryAggregator_MIN TimeSeriesQueryAggregator = 4
)

var TimeSeriesQueryAggregator_name = map[int32]string{
	1: "AVG",
	2: "SUM",
	3: "MAX",
	4: "MIN",
}
var TimeSeriesQueryAggregator_value = map[string]int32{
	"AVG": 1,
	"SUM": 2,
	"MAX": 3,
	"MIN": 4,
}

func (x TimeSeriesQueryAggregator) Enum() *TimeSeriesQueryAggregator {
	p := new(TimeSeriesQueryAggregator)
	*p = x
	return p
}
func (x TimeSeriesQueryAggregator) String() string {
	return proto.EnumName(TimeSeriesQueryAggregator_name, int32(x))
}
func (x *TimeSeriesQueryAggregator) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(TimeSeriesQueryAggregator_value, data, "TimeSeriesQueryAggregator")
	if err != nil {
		return err
	}
	*x = TimeSeriesQueryAggregator(value)
	return nil
}

// TimeSeriesQueryDerivative describes a derivative function used to convert
// returned datapoints into a rate-of-change.
type TimeSeriesQueryDerivative int32

const (
	// NONE is the default value, and does not apply a derivative function.
	TimeSeriesQueryDerivative_NONE TimeSeriesQueryDerivative = 0
	// DERIVATIVE returns the first-order derivative of values in the time series.
	TimeSeriesQueryDerivative_DERIVATIVE TimeSeriesQueryDerivative = 1
	// NON_NEGATIVE_DERIVATIVE returns only non-negative values of the first-order
	// derivative; negative values are returned as zero. This should be used for
	// counters that monotonically increase, but might wrap or reset.
	TimeSeriesQueryDerivative_NON_NEGATIVE_DERIVATIVE TimeSeriesQueryDerivative = 2
)

var TimeSeriesQueryDerivative_name = map[int32]string{
	0: "NONE",
	1: "DERIVATIVE",
	2: "NON_NEGATIVE_DERIVATIVE",
}
var TimeSeriesQueryDerivative_value = map[string]int32{
	"NONE":                    0,
	"DERIVATIVE":              1,
	"NON_NEGATIVE_DERIVATIVE": 2,
}

func (x TimeSeriesQueryDerivative) Enum() *TimeSeriesQueryDerivative {
	p := new(TimeSeriesQueryDerivative)
	*p = x
	return p
}
func (x TimeSeriesQueryDerivative) String() string {
	return proto.EnumName(TimeSeriesQueryDerivative_name, int32(x))
}
func (x *TimeSeriesQueryDerivative) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(TimeSeriesQueryDerivative_value, data, "TimeSeriesQueryDerivative")
	if err != nil {
		return err
	}
	*x = TimeSeriesQueryDerivative(value)
	return nil
}

// TimeSeriesDatapoint is a single point of time series data; a value associated
// with a timestamp.
type TimeSeriesDatapoint struct {
	// The timestamp when this datapoint is located, expressed in nanoseconds
	// since the unix epoch.
	TimestampNanos int64 `protobuf:"varint,1,opt,name=timestamp_nanos" json:"timestamp_nanos"`
	// A floating point representation of the value of this datapoint.
	Value float64 `protobuf:"fixed64,2,opt,name=value" json:"value"`
}

func (m *TimeSeriesDatapoint) Reset()         { *m = TimeSeriesDatapoint{} }
func (m *TimeSeriesDatapoint) String() string { return proto.CompactTextString(m) }
func (*TimeSeriesDatapoint) ProtoMessage()    {}

// TimeSeriesData is a set of measurements of a single named variable at
// multiple points in time. This message contains a name and a source which, in
// combination, uniquely identify the time series being measured. Measurement
// data is represented as a repeated set of TimeSeriesDatapoint messages.
type TimeSeriesData struct {
	// A string which uniquely identifies the variable from which this data was
	// measured.
	Name string `protobuf:"bytes,1,opt,name=name" json:"name"`
	// A string which identifies the unique source from which the variable was measured.
	Source string `protobuf:"bytes,2,opt,name=source" json:"source"`
	// Datapoints representing one or more measurements taken from the variable.
	Datapoints []*TimeSeriesDatapoint `protobuf:"bytes,3,rep,name=datapoints" json:"datapoints,omitempty"`
}

func (m *TimeSeriesData) Reset()         { *m = TimeSeriesData{} }
func (m *TimeSeriesData) String() string { return proto.CompactTextString(m) }
func (*TimeSeriesData) ProtoMessage()    {}

// Each Query defines a specific metric to query over the time span of
// this request.
type Query struct {
	// The name of the time series to query.
	Name string `protobuf:"bytes,1,opt,name=name" json:"name"`
	// A downsampling aggregation function to apply to datapoints within the
	// same sample period.
	Downsampler *TimeSeriesQueryAggregator `protobuf:"varint,2,opt,name=downsampler,enum=cockroach.ts.TimeSeriesQueryAggregator,def=1" json:"downsampler,omitempty"`
	// An aggregation function used to combine timelike datapoints from the
	// different sources being queried.
	SourceAggregator *TimeSeriesQueryAggregator `protobuf:"varint,3,opt,name=source_aggregator,enum=cockroach.ts.TimeSeriesQueryAggregator,def=2" json:"source_aggregator,omitempty"`
	// If set to a value other than 'NONE', query will return a derivative
	// (rate of change) of the aggregated datapoints.
	Derivative *TimeSeriesQueryDerivative `protobuf:"varint,4,opt,name=derivative,enum=cockroach.ts.TimeSeriesQueryDerivative,def=0" json:"derivative,omitempty"`
	// An optional list of sources to restrict the time series query. If no
	// sources are provided, all available sources will be queried.
	Sources []string `protobuf:"bytes,5,rep,name=sources" json:"sources,omitempty"`
}

func (m *Query) Reset()         { *m = Query{} }
func (m *Query) String() string { return proto.CompactTextString(m) }
func (*Query) ProtoMessage()    {}

const Default_Query_Downsampler TimeSeriesQueryAggregator = TimeSeriesQueryAggregator_AVG
const Default_Query_SourceAggregator TimeSeriesQueryAggregator = TimeSeriesQueryAggregator_SUM
const Default_Query_Derivative TimeSeriesQueryDerivative = TimeSeriesQueryDerivative_NONE

func (m *Query) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Query) GetDownsampler() TimeSeriesQueryAggregator {
	if m != nil && m.Downsampler != nil {
		return *m.Downsampler
	}
	return Default_Query_Downsampler
}

func (m *Query) GetSourceAggregator() TimeSeriesQueryAggregator {
	if m != nil && m.SourceAggregator != nil {
		return *m.SourceAggregator
	}
	return Default_Query_SourceAggregator
}

func (m *Query) GetDerivative() TimeSeriesQueryDerivative {
	if m != nil && m.Derivative != nil {
		return *m.Derivative
	}
	return Default_Query_Derivative
}

func (m *Query) GetSources() []string {
	if m != nil {
		return m.Sources
	}
	return nil
}

// TimeSeriesQueryRequest is the standard incoming time series query request
// accepted from cockroach clients.
type TimeSeriesQueryRequest struct {
	// A timestamp in nanoseconds which defines the early bound of the time span
	// for this query.
	StartNanos int64 `protobuf:"varint,1,opt,name=start_nanos" json:"start_nanos"`
	// A timestamp in nanoseconds which defines the late bound of the time span
	// for this query. Must be greater than start_nanos.
	EndNanos int64 `protobuf:"varint,2,opt,name=end_nanos" json:"end_nanos"`
	// A set of Queries for this request. A request must have at least one
	// Query.
	Queries []Query `protobuf:"bytes,3,rep,name=queries" json:"queries"`
}

func (m *TimeSeriesQueryRequest) Reset()         { *m = TimeSeriesQueryRequest{} }
func (m *TimeSeriesQueryRequest) String() string { return proto.CompactTextString(m) }
func (*TimeSeriesQueryRequest) ProtoMessage()    {}

// TimeSeriesQueryResponse is the standard response for time series queries
// returned to cockroach clients.
type TimeSeriesQueryResponse struct {
	// A set of Results; there will be one result for each Query in the matching
	// TimeSeriesQueryRequest, in the same order. A Result will be present for
	// each Query even if there are zero datapoints to return.
	Results []*TimeSeriesQueryResponse_Result `protobuf:"bytes,1,rep,name=results" json:"results,omitempty"`
}

func (m *TimeSeriesQueryResponse) Reset()         { *m = TimeSeriesQueryResponse{} }
func (m *TimeSeriesQueryResponse) String() string { return proto.CompactTextString(m) }
func (*TimeSeriesQueryResponse) ProtoMessage()    {}

// Result is the data returned from a single metric query over a time span.
type TimeSeriesQueryResponse_Result struct {
	Query      `protobuf:"bytes,1,opt,name=query,embedded=query" json:"query"`
	Datapoints []*TimeSeriesDatapoint `protobuf:"bytes,2,rep,name=datapoints" json:"datapoints,omitempty"`
}

func (m *TimeSeriesQueryResponse_Result) Reset()         { *m = TimeSeriesQueryResponse_Result{} }
func (m *TimeSeriesQueryResponse_Result) String() string { return proto.CompactTextString(m) }
func (*TimeSeriesQueryResponse_Result) ProtoMessage()    {}

func init() {
	proto.RegisterType((*TimeSeriesDatapoint)(nil), "cockroach.ts.TimeSeriesDatapoint")
	proto.RegisterType((*TimeSeriesData)(nil), "cockroach.ts.TimeSeriesData")
	proto.RegisterType((*Query)(nil), "cockroach.ts.Query")
	proto.RegisterType((*TimeSeriesQueryRequest)(nil), "cockroach.ts.TimeSeriesQueryRequest")
	proto.RegisterType((*TimeSeriesQueryResponse)(nil), "cockroach.ts.TimeSeriesQueryResponse")
	proto.RegisterType((*TimeSeriesQueryResponse_Result)(nil), "cockroach.ts.TimeSeriesQueryResponse.Result")
	proto.RegisterEnum("cockroach.ts.TimeSeriesQueryAggregator", TimeSeriesQueryAggregator_name, TimeSeriesQueryAggregator_value)
	proto.RegisterEnum("cockroach.ts.TimeSeriesQueryDerivative", TimeSeriesQueryDerivative_name, TimeSeriesQueryDerivative_value)
}
func (m *TimeSeriesDatapoint) Marshal() (data []byte, err error) {
	size := m.Size()
	data = make([]byte, size)
	n, err := m.MarshalTo(data)
	if err != nil {
		return nil, err
	}
	return data[:n], nil
}

func (m *TimeSeriesDatapoint) MarshalTo(data []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	data[i] = 0x8
	i++
	i = encodeVarintTimeseries(data, i, uint64(m.TimestampNanos))
	data[i] = 0x11
	i++
	i = encodeFixed64Timeseries(data, i, uint64(math.Float64bits(float64(m.Value))))
	return i, nil
}

func (m *TimeSeriesData) Marshal() (data []byte, err error) {
	size := m.Size()
	data = make([]byte, size)
	n, err := m.MarshalTo(data)
	if err != nil {
		return nil, err
	}
	return data[:n], nil
}

func (m *TimeSeriesData) MarshalTo(data []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	data[i] = 0xa
	i++
	i = encodeVarintTimeseries(data, i, uint64(len(m.Name)))
	i += copy(data[i:], m.Name)
	data[i] = 0x12
	i++
	i = encodeVarintTimeseries(data, i, uint64(len(m.Source)))
	i += copy(data[i:], m.Source)
	if len(m.Datapoints) > 0 {
		for _, msg := range m.Datapoints {
			data[i] = 0x1a
			i++
			i = encodeVarintTimeseries(data, i, uint64(msg.Size()))
			n, err := msg.MarshalTo(data[i:])
			if err != nil {
				return 0, err
			}
			i += n
		}
	}
	return i, nil
}

func (m *Query) Marshal() (data []byte, err error) {
	size := m.Size()
	data = make([]byte, size)
	n, err := m.MarshalTo(data)
	if err != nil {
		return nil, err
	}
	return data[:n], nil
}

func (m *Query) MarshalTo(data []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	data[i] = 0xa
	i++
	i = encodeVarintTimeseries(data, i, uint64(len(m.Name)))
	i += copy(data[i:], m.Name)
	if m.Downsampler != nil {
		data[i] = 0x10
		i++
		i = encodeVarintTimeseries(data, i, uint64(*m.Downsampler))
	}
	if m.SourceAggregator != nil {
		data[i] = 0x18
		i++
		i = encodeVarintTimeseries(data, i, uint64(*m.SourceAggregator))
	}
	if m.Derivative != nil {
		data[i] = 0x20
		i++
		i = encodeVarintTimeseries(data, i, uint64(*m.Derivative))
	}
	if len(m.Sources) > 0 {
		for _, s := range m.Sources {
			data[i] = 0x2a
			i++
			l = len(s)
			for l >= 1<<7 {
				data[i] = uint8(uint64(l)&0x7f | 0x80)
				l >>= 7
				i++
			}
			data[i] = uint8(l)
			i++
			i += copy(data[i:], s)
		}
	}
	return i, nil
}

func (m *TimeSeriesQueryRequest) Marshal() (data []byte, err error) {
	size := m.Size()
	data = make([]byte, size)
	n, err := m.MarshalTo(data)
	if err != nil {
		return nil, err
	}
	return data[:n], nil
}

func (m *TimeSeriesQueryRequest) MarshalTo(data []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	data[i] = 0x8
	i++
	i = encodeVarintTimeseries(data, i, uint64(m.StartNanos))
	data[i] = 0x10
	i++
	i = encodeVarintTimeseries(data, i, uint64(m.EndNanos))
	if len(m.Queries) > 0 {
		for _, msg := range m.Queries {
			data[i] = 0x1a
			i++
			i = encodeVarintTimeseries(data, i, uint64(msg.Size()))
			n, err := msg.MarshalTo(data[i:])
			if err != nil {
				return 0, err
			}
			i += n
		}
	}
	return i, nil
}

func (m *TimeSeriesQueryResponse) Marshal() (data []byte, err error) {
	size := m.Size()
	data = make([]byte, size)
	n, err := m.MarshalTo(data)
	if err != nil {
		return nil, err
	}
	return data[:n], nil
}

func (m *TimeSeriesQueryResponse) MarshalTo(data []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Results) > 0 {
		for _, msg := range m.Results {
			data[i] = 0xa
			i++
			i = encodeVarintTimeseries(data, i, uint64(msg.Size()))
			n, err := msg.MarshalTo(data[i:])
			if err != nil {
				return 0, err
			}
			i += n
		}
	}
	return i, nil
}

func (m *TimeSeriesQueryResponse_Result) Marshal() (data []byte, err error) {
	size := m.Size()
	data = make([]byte, size)
	n, err := m.MarshalTo(data)
	if err != nil {
		return nil, err
	}
	return data[:n], nil
}

func (m *TimeSeriesQueryResponse_Result) MarshalTo(data []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	data[i] = 0xa
	i++
	i = encodeVarintTimeseries(data, i, uint64(m.Query.Size()))
	n1, err := m.Query.MarshalTo(data[i:])
	if err != nil {
		return 0, err
	}
	i += n1
	if len(m.Datapoints) > 0 {
		for _, msg := range m.Datapoints {
			data[i] = 0x12
			i++
			i = encodeVarintTimeseries(data, i, uint64(msg.Size()))
			n, err := msg.MarshalTo(data[i:])
			if err != nil {
				return 0, err
			}
			i += n
		}
	}
	return i, nil
}

func encodeFixed64Timeseries(data []byte, offset int, v uint64) int {
	data[offset] = uint8(v)
	data[offset+1] = uint8(v >> 8)
	data[offset+2] = uint8(v >> 16)
	data[offset+3] = uint8(v >> 24)
	data[offset+4] = uint8(v >> 32)
	data[offset+5] = uint8(v >> 40)
	data[offset+6] = uint8(v >> 48)
	data[offset+7] = uint8(v >> 56)
	return offset + 8
}
func encodeFixed32Timeseries(data []byte, offset int, v uint32) int {
	data[offset] = uint8(v)
	data[offset+1] = uint8(v >> 8)
	data[offset+2] = uint8(v >> 16)
	data[offset+3] = uint8(v >> 24)
	return offset + 4
}
func encodeVarintTimeseries(data []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		data[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	data[offset] = uint8(v)
	return offset + 1
}
func (m *TimeSeriesDatapoint) Size() (n int) {
	var l int
	_ = l
	n += 1 + sovTimeseries(uint64(m.TimestampNanos))
	n += 9
	return n
}

func (m *TimeSeriesData) Size() (n int) {
	var l int
	_ = l
	l = len(m.Name)
	n += 1 + l + sovTimeseries(uint64(l))
	l = len(m.Source)
	n += 1 + l + sovTimeseries(uint64(l))
	if len(m.Datapoints) > 0 {
		for _, e := range m.Datapoints {
			l = e.Size()
			n += 1 + l + sovTimeseries(uint64(l))
		}
	}
	return n
}

func (m *Query) Size() (n int) {
	var l int
	_ = l
	l = len(m.Name)
	n += 1 + l + sovTimeseries(uint64(l))
	if m.Downsampler != nil {
		n += 1 + sovTimeseries(uint64(*m.Downsampler))
	}
	if m.SourceAggregator != nil {
		n += 1 + sovTimeseries(uint64(*m.SourceAggregator))
	}
	if m.Derivative != nil {
		n += 1 + sovTimeseries(uint64(*m.Derivative))
	}
	if len(m.Sources) > 0 {
		for _, s := range m.Sources {
			l = len(s)
			n += 1 + l + sovTimeseries(uint64(l))
		}
	}
	return n
}

func (m *TimeSeriesQueryRequest) Size() (n int) {
	var l int
	_ = l
	n += 1 + sovTimeseries(uint64(m.StartNanos))
	n += 1 + sovTimeseries(uint64(m.EndNanos))
	if len(m.Queries) > 0 {
		for _, e := range m.Queries {
			l = e.Size()
			n += 1 + l + sovTimeseries(uint64(l))
		}
	}
	return n
}

func (m *TimeSeriesQueryResponse) Size() (n int) {
	var l int
	_ = l
	if len(m.Results) > 0 {
		for _, e := range m.Results {
			l = e.Size()
			n += 1 + l + sovTimeseries(uint64(l))
		}
	}
	return n
}

func (m *TimeSeriesQueryResponse_Result) Size() (n int) {
	var l int
	_ = l
	l = m.Query.Size()
	n += 1 + l + sovTimeseries(uint64(l))
	if len(m.Datapoints) > 0 {
		for _, e := range m.Datapoints {
			l = e.Size()
			n += 1 + l + sovTimeseries(uint64(l))
		}
	}
	return n
}

func sovTimeseries(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozTimeseries(x uint64) (n int) {
	return sovTimeseries(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *TimeSeriesDatapoint) Unmarshal(data []byte) error {
	l := len(data)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTimeseries
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := data[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: TimeSeriesDatapoint: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: TimeSeriesDatapoint: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field TimestampNanos", wireType)
			}
			m.TimestampNanos = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTimeseries
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				m.TimestampNanos |= (int64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 1 {
				return fmt.Errorf("proto: wrong wireType = %d for field Value", wireType)
			}
			var v uint64
			if (iNdEx + 8) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += 8
			v = uint64(data[iNdEx-8])
			v |= uint64(data[iNdEx-7]) << 8
			v |= uint64(data[iNdEx-6]) << 16
			v |= uint64(data[iNdEx-5]) << 24
			v |= uint64(data[iNdEx-4]) << 32
			v |= uint64(data[iNdEx-3]) << 40
			v |= uint64(data[iNdEx-2]) << 48
			v |= uint64(data[iNdEx-1]) << 56
			m.Value = float64(math.Float64frombits(v))
		default:
			iNdEx = preIndex
			skippy, err := skipTimeseries(data[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthTimeseries
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *TimeSeriesData) Unmarshal(data []byte) error {
	l := len(data)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTimeseries
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := data[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: TimeSeriesData: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: TimeSeriesData: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Name", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTimeseries
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTimeseries
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Name = string(data[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Source", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTimeseries
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTimeseries
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Source = string(data[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Datapoints", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTimeseries
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthTimeseries
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Datapoints = append(m.Datapoints, &TimeSeriesDatapoint{})
			if err := m.Datapoints[len(m.Datapoints)-1].Unmarshal(data[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTimeseries(data[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthTimeseries
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *Query) Unmarshal(data []byte) error {
	l := len(data)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTimeseries
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := data[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Query: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Query: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Name", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTimeseries
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTimeseries
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Name = string(data[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Downsampler", wireType)
			}
			var v TimeSeriesQueryAggregator
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTimeseries
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				v |= (TimeSeriesQueryAggregator(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.Downsampler = &v
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field SourceAggregator", wireType)
			}
			var v TimeSeriesQueryAggregator
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTimeseries
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				v |= (TimeSeriesQueryAggregator(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.SourceAggregator = &v
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Derivative", wireType)
			}
			var v TimeSeriesQueryDerivative
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTimeseries
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				v |= (TimeSeriesQueryDerivative(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.Derivative = &v
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Sources", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTimeseries
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTimeseries
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Sources = append(m.Sources, string(data[iNdEx:postIndex]))
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTimeseries(data[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthTimeseries
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *TimeSeriesQueryRequest) Unmarshal(data []byte) error {
	l := len(data)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTimeseries
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := data[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: TimeSeriesQueryRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: TimeSeriesQueryRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field StartNanos", wireType)
			}
			m.StartNanos = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTimeseries
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				m.StartNanos |= (int64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field EndNanos", wireType)
			}
			m.EndNanos = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTimeseries
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				m.EndNanos |= (int64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Queries", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTimeseries
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthTimeseries
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Queries = append(m.Queries, Query{})
			if err := m.Queries[len(m.Queries)-1].Unmarshal(data[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTimeseries(data[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthTimeseries
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *TimeSeriesQueryResponse) Unmarshal(data []byte) error {
	l := len(data)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTimeseries
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := data[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: TimeSeriesQueryResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: TimeSeriesQueryResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Results", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTimeseries
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthTimeseries
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Results = append(m.Results, &TimeSeriesQueryResponse_Result{})
			if err := m.Results[len(m.Results)-1].Unmarshal(data[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTimeseries(data[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthTimeseries
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *TimeSeriesQueryResponse_Result) Unmarshal(data []byte) error {
	l := len(data)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTimeseries
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := data[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Result: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Result: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Query", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTimeseries
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthTimeseries
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Query.Unmarshal(data[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Datapoints", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTimeseries
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthTimeseries
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Datapoints = append(m.Datapoints, &TimeSeriesDatapoint{})
			if err := m.Datapoints[len(m.Datapoints)-1].Unmarshal(data[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTimeseries(data[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthTimeseries
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipTimeseries(data []byte) (n int, err error) {
	l := len(data)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowTimeseries
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := data[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowTimeseries
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if data[iNdEx-1] < 0x80 {
					break
				}
			}
			return iNdEx, nil
		case 1:
			iNdEx += 8
			return iNdEx, nil
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowTimeseries
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			iNdEx += length
			if length < 0 {
				return 0, ErrInvalidLengthTimeseries
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowTimeseries
					}
					if iNdEx >= l {
						return 0, io.ErrUnexpectedEOF
					}
					b := data[iNdEx]
					iNdEx++
					innerWire |= (uint64(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				innerWireType := int(innerWire & 0x7)
				if innerWireType == 4 {
					break
				}
				next, err := skipTimeseries(data[start:])
				if err != nil {
					return 0, err
				}
				iNdEx = start + next
			}
			return iNdEx, nil
		case 4:
			return iNdEx, nil
		case 5:
			iNdEx += 4
			return iNdEx, nil
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
	}
	panic("unreachable")
}

var (
	ErrInvalidLengthTimeseries = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowTimeseries   = fmt.Errorf("proto: integer overflow")
)
