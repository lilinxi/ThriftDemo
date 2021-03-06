// Code generated by Thrift Compiler (0.14.2). DO NOT EDIT.

package example

import(
	"bytes"
	"context"
	"fmt"
	"time"
	"github.com/apache/thrift/lib/go/thrift"
)

// (needed to ensure safety because of naive import list construction.)
var _ = thrift.ZERO
var _ = fmt.Printf
var _ = context.Background
var _ = time.Now
var _ = bytes.Equal

// Attributes:
//  - Text
type Data struct {
  Text string `thrift:"text,1" db:"text" json:"text"`
}

func NewData() *Data {
  return &Data{}
}


func (p *Data) GetText() string {
  return p.Text
}
func (p *Data) Read(ctx context.Context, iprot thrift.TProtocol) error {
  if _, err := iprot.ReadStructBegin(ctx); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T read error: ", p), err)
  }


  for {
    _, fieldTypeId, fieldId, err := iprot.ReadFieldBegin(ctx)
    if err != nil {
      return thrift.PrependError(fmt.Sprintf("%T field %d read error: ", p, fieldId), err)
    }
    if fieldTypeId == thrift.STOP { break; }
    switch fieldId {
    case 1:
      if fieldTypeId == thrift.STRING {
        if err := p.ReadField1(ctx, iprot); err != nil {
          return err
        }
      } else {
        if err := iprot.Skip(ctx, fieldTypeId); err != nil {
          return err
        }
      }
    default:
      if err := iprot.Skip(ctx, fieldTypeId); err != nil {
        return err
      }
    }
    if err := iprot.ReadFieldEnd(ctx); err != nil {
      return err
    }
  }
  if err := iprot.ReadStructEnd(ctx); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
  }
  return nil
}

func (p *Data)  ReadField1(ctx context.Context, iprot thrift.TProtocol) error {
  if v, err := iprot.ReadString(ctx); err != nil {
  return thrift.PrependError("error reading field 1: ", err)
} else {
  p.Text = v
}
  return nil
}

func (p *Data) Write(ctx context.Context, oprot thrift.TProtocol) error {
  if err := oprot.WriteStructBegin(ctx, "Data"); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err) }
  if p != nil {
    if err := p.writeField1(ctx, oprot); err != nil { return err }
  }
  if err := oprot.WriteFieldStop(ctx); err != nil {
    return thrift.PrependError("write field stop error: ", err) }
  if err := oprot.WriteStructEnd(ctx); err != nil {
    return thrift.PrependError("write struct stop error: ", err) }
  return nil
}

func (p *Data) writeField1(ctx context.Context, oprot thrift.TProtocol) (err error) {
  if err := oprot.WriteFieldBegin(ctx, "text", thrift.STRING, 1); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write field begin error 1:text: ", p), err) }
  if err := oprot.WriteString(ctx, string(p.Text)); err != nil {
  return thrift.PrependError(fmt.Sprintf("%T.text (1) field write error: ", p), err) }
  if err := oprot.WriteFieldEnd(ctx); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write field end error 1:text: ", p), err) }
  return err
}

func (p *Data) Equals(other *Data) bool {
  if p == other {
    return true
  } else if p == nil || other == nil {
    return false
  }
  if p.Text != other.Text { return false }
  return true
}

func (p *Data) String() string {
  if p == nil {
    return "<nil>"
  }
  return fmt.Sprintf("Data(%+v)", *p)
}

type FormatData interface {
  // Parameters:
  //  - Data
  DoFormat(ctx context.Context, data *Data) (_r *Data, _err error)
}

type FormatDataClient struct {
  c thrift.TClient
  meta thrift.ResponseMeta
}

func NewFormatDataClientFactory(t thrift.TTransport, f thrift.TProtocolFactory) *FormatDataClient {
  return &FormatDataClient{
    c: thrift.NewTStandardClient(f.GetProtocol(t), f.GetProtocol(t)),
  }
}

func NewFormatDataClientProtocol(t thrift.TTransport, iprot thrift.TProtocol, oprot thrift.TProtocol) *FormatDataClient {
  return &FormatDataClient{
    c: thrift.NewTStandardClient(iprot, oprot),
  }
}

func NewFormatDataClient(c thrift.TClient) *FormatDataClient {
  return &FormatDataClient{
    c: c,
  }
}

func (p *FormatDataClient) Client_() thrift.TClient {
  return p.c
}

func (p *FormatDataClient) LastResponseMeta_() thrift.ResponseMeta {
  return p.meta
}

func (p *FormatDataClient) SetLastResponseMeta_(meta thrift.ResponseMeta) {
  p.meta = meta
}

// Parameters:
//  - Data
func (p *FormatDataClient) DoFormat(ctx context.Context, data *Data) (_r *Data, _err error) {
  var _args0 FormatDataDoFormatArgs
  _args0.Data = data
  var _result2 FormatDataDoFormatResult
  var _meta1 thrift.ResponseMeta
  _meta1, _err = p.Client_().Call(ctx, "do_format", &_args0, &_result2)
  p.SetLastResponseMeta_(_meta1)
  if _err != nil {
    return
  }
  return _result2.GetSuccess(), nil
}

type FormatDataProcessor struct {
  processorMap map[string]thrift.TProcessorFunction
  handler FormatData
}

func (p *FormatDataProcessor) AddToProcessorMap(key string, processor thrift.TProcessorFunction) {
  p.processorMap[key] = processor
}

func (p *FormatDataProcessor) GetProcessorFunction(key string) (processor thrift.TProcessorFunction, ok bool) {
  processor, ok = p.processorMap[key]
  return processor, ok
}

func (p *FormatDataProcessor) ProcessorMap() map[string]thrift.TProcessorFunction {
  return p.processorMap
}

func NewFormatDataProcessor(handler FormatData) *FormatDataProcessor {

  self3 := &FormatDataProcessor{handler:handler, processorMap:make(map[string]thrift.TProcessorFunction)}
  self3.processorMap["do_format"] = &formatDataProcessorDoFormat{handler:handler}
return self3
}

func (p *FormatDataProcessor) Process(ctx context.Context, iprot, oprot thrift.TProtocol) (success bool, err thrift.TException) {
  name, _, seqId, err2 := iprot.ReadMessageBegin(ctx)
  if err2 != nil { return false, thrift.WrapTException(err2) }
  if processor, ok := p.GetProcessorFunction(name); ok {
    return processor.Process(ctx, seqId, iprot, oprot)
  }
  iprot.Skip(ctx, thrift.STRUCT)
  iprot.ReadMessageEnd(ctx)
  x4 := thrift.NewTApplicationException(thrift.UNKNOWN_METHOD, "Unknown function " + name)
  oprot.WriteMessageBegin(ctx, name, thrift.EXCEPTION, seqId)
  x4.Write(ctx, oprot)
  oprot.WriteMessageEnd(ctx)
  oprot.Flush(ctx)
  return false, x4

}

type formatDataProcessorDoFormat struct {
  handler FormatData
}

func (p *formatDataProcessorDoFormat) Process(ctx context.Context, seqId int32, iprot, oprot thrift.TProtocol) (success bool, err thrift.TException) {
  args := FormatDataDoFormatArgs{}
  var err2 error
  if err2 = args.Read(ctx, iprot); err2 != nil {
    iprot.ReadMessageEnd(ctx)
    x := thrift.NewTApplicationException(thrift.PROTOCOL_ERROR, err2.Error())
    oprot.WriteMessageBegin(ctx, "do_format", thrift.EXCEPTION, seqId)
    x.Write(ctx, oprot)
    oprot.WriteMessageEnd(ctx)
    oprot.Flush(ctx)
    return false, thrift.WrapTException(err2)
  }
  iprot.ReadMessageEnd(ctx)

  tickerCancel := func() {}
  // Start a goroutine to do server side connectivity check.
  if thrift.ServerConnectivityCheckInterval > 0 {
    var cancel context.CancelFunc
    ctx, cancel = context.WithCancel(ctx)
    defer cancel()
    var tickerCtx context.Context
    tickerCtx, tickerCancel = context.WithCancel(context.Background())
    defer tickerCancel()
    go func(ctx context.Context, cancel context.CancelFunc) {
      ticker := time.NewTicker(thrift.ServerConnectivityCheckInterval)
      defer ticker.Stop()
      for {
        select {
        case <-ctx.Done():
          return
        case <-ticker.C:
          if !iprot.Transport().IsOpen() {
            cancel()
            return
          }
        }
      }
    }(tickerCtx, cancel)
  }

  result := FormatDataDoFormatResult{}
  var retval *Data
  if retval, err2 = p.handler.DoFormat(ctx, args.Data); err2 != nil {
    tickerCancel()
    if err2 == thrift.ErrAbandonRequest {
      return false, thrift.WrapTException(err2)
    }
    x := thrift.NewTApplicationException(thrift.INTERNAL_ERROR, "Internal error processing do_format: " + err2.Error())
    oprot.WriteMessageBegin(ctx, "do_format", thrift.EXCEPTION, seqId)
    x.Write(ctx, oprot)
    oprot.WriteMessageEnd(ctx)
    oprot.Flush(ctx)
    return true, thrift.WrapTException(err2)
  } else {
    result.Success = retval
  }
  tickerCancel()
  if err2 = oprot.WriteMessageBegin(ctx, "do_format", thrift.REPLY, seqId); err2 != nil {
    err = thrift.WrapTException(err2)
  }
  if err2 = result.Write(ctx, oprot); err == nil && err2 != nil {
    err = thrift.WrapTException(err2)
  }
  if err2 = oprot.WriteMessageEnd(ctx); err == nil && err2 != nil {
    err = thrift.WrapTException(err2)
  }
  if err2 = oprot.Flush(ctx); err == nil && err2 != nil {
    err = thrift.WrapTException(err2)
  }
  if err != nil {
    return
  }
  return true, err
}


// HELPER FUNCTIONS AND STRUCTURES

// Attributes:
//  - Data
type FormatDataDoFormatArgs struct {
  Data *Data `thrift:"data,1" db:"data" json:"data"`
}

func NewFormatDataDoFormatArgs() *FormatDataDoFormatArgs {
  return &FormatDataDoFormatArgs{}
}

var FormatDataDoFormatArgs_Data_DEFAULT *Data
func (p *FormatDataDoFormatArgs) GetData() *Data {
  if !p.IsSetData() {
    return FormatDataDoFormatArgs_Data_DEFAULT
  }
return p.Data
}
func (p *FormatDataDoFormatArgs) IsSetData() bool {
  return p.Data != nil
}

func (p *FormatDataDoFormatArgs) Read(ctx context.Context, iprot thrift.TProtocol) error {
  if _, err := iprot.ReadStructBegin(ctx); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T read error: ", p), err)
  }


  for {
    _, fieldTypeId, fieldId, err := iprot.ReadFieldBegin(ctx)
    if err != nil {
      return thrift.PrependError(fmt.Sprintf("%T field %d read error: ", p, fieldId), err)
    }
    if fieldTypeId == thrift.STOP { break; }
    switch fieldId {
    case 1:
      if fieldTypeId == thrift.STRUCT {
        if err := p.ReadField1(ctx, iprot); err != nil {
          return err
        }
      } else {
        if err := iprot.Skip(ctx, fieldTypeId); err != nil {
          return err
        }
      }
    default:
      if err := iprot.Skip(ctx, fieldTypeId); err != nil {
        return err
      }
    }
    if err := iprot.ReadFieldEnd(ctx); err != nil {
      return err
    }
  }
  if err := iprot.ReadStructEnd(ctx); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
  }
  return nil
}

func (p *FormatDataDoFormatArgs)  ReadField1(ctx context.Context, iprot thrift.TProtocol) error {
  p.Data = &Data{}
  if err := p.Data.Read(ctx, iprot); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T error reading struct: ", p.Data), err)
  }
  return nil
}

func (p *FormatDataDoFormatArgs) Write(ctx context.Context, oprot thrift.TProtocol) error {
  if err := oprot.WriteStructBegin(ctx, "do_format_args"); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err) }
  if p != nil {
    if err := p.writeField1(ctx, oprot); err != nil { return err }
  }
  if err := oprot.WriteFieldStop(ctx); err != nil {
    return thrift.PrependError("write field stop error: ", err) }
  if err := oprot.WriteStructEnd(ctx); err != nil {
    return thrift.PrependError("write struct stop error: ", err) }
  return nil
}

func (p *FormatDataDoFormatArgs) writeField1(ctx context.Context, oprot thrift.TProtocol) (err error) {
  if err := oprot.WriteFieldBegin(ctx, "data", thrift.STRUCT, 1); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write field begin error 1:data: ", p), err) }
  if err := p.Data.Write(ctx, oprot); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T error writing struct: ", p.Data), err)
  }
  if err := oprot.WriteFieldEnd(ctx); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write field end error 1:data: ", p), err) }
  return err
}

func (p *FormatDataDoFormatArgs) String() string {
  if p == nil {
    return "<nil>"
  }
  return fmt.Sprintf("FormatDataDoFormatArgs(%+v)", *p)
}

// Attributes:
//  - Success
type FormatDataDoFormatResult struct {
  Success *Data `thrift:"success,0" db:"success" json:"success,omitempty"`
}

func NewFormatDataDoFormatResult() *FormatDataDoFormatResult {
  return &FormatDataDoFormatResult{}
}

var FormatDataDoFormatResult_Success_DEFAULT *Data
func (p *FormatDataDoFormatResult) GetSuccess() *Data {
  if !p.IsSetSuccess() {
    return FormatDataDoFormatResult_Success_DEFAULT
  }
return p.Success
}
func (p *FormatDataDoFormatResult) IsSetSuccess() bool {
  return p.Success != nil
}

func (p *FormatDataDoFormatResult) Read(ctx context.Context, iprot thrift.TProtocol) error {
  if _, err := iprot.ReadStructBegin(ctx); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T read error: ", p), err)
  }


  for {
    _, fieldTypeId, fieldId, err := iprot.ReadFieldBegin(ctx)
    if err != nil {
      return thrift.PrependError(fmt.Sprintf("%T field %d read error: ", p, fieldId), err)
    }
    if fieldTypeId == thrift.STOP { break; }
    switch fieldId {
    case 0:
      if fieldTypeId == thrift.STRUCT {
        if err := p.ReadField0(ctx, iprot); err != nil {
          return err
        }
      } else {
        if err := iprot.Skip(ctx, fieldTypeId); err != nil {
          return err
        }
      }
    default:
      if err := iprot.Skip(ctx, fieldTypeId); err != nil {
        return err
      }
    }
    if err := iprot.ReadFieldEnd(ctx); err != nil {
      return err
    }
  }
  if err := iprot.ReadStructEnd(ctx); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
  }
  return nil
}

func (p *FormatDataDoFormatResult)  ReadField0(ctx context.Context, iprot thrift.TProtocol) error {
  p.Success = &Data{}
  if err := p.Success.Read(ctx, iprot); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T error reading struct: ", p.Success), err)
  }
  return nil
}

func (p *FormatDataDoFormatResult) Write(ctx context.Context, oprot thrift.TProtocol) error {
  if err := oprot.WriteStructBegin(ctx, "do_format_result"); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err) }
  if p != nil {
    if err := p.writeField0(ctx, oprot); err != nil { return err }
  }
  if err := oprot.WriteFieldStop(ctx); err != nil {
    return thrift.PrependError("write field stop error: ", err) }
  if err := oprot.WriteStructEnd(ctx); err != nil {
    return thrift.PrependError("write struct stop error: ", err) }
  return nil
}

func (p *FormatDataDoFormatResult) writeField0(ctx context.Context, oprot thrift.TProtocol) (err error) {
  if p.IsSetSuccess() {
    if err := oprot.WriteFieldBegin(ctx, "success", thrift.STRUCT, 0); err != nil {
      return thrift.PrependError(fmt.Sprintf("%T write field begin error 0:success: ", p), err) }
    if err := p.Success.Write(ctx, oprot); err != nil {
      return thrift.PrependError(fmt.Sprintf("%T error writing struct: ", p.Success), err)
    }
    if err := oprot.WriteFieldEnd(ctx); err != nil {
      return thrift.PrependError(fmt.Sprintf("%T write field end error 0:success: ", p), err) }
  }
  return err
}

func (p *FormatDataDoFormatResult) String() string {
  if p == nil {
    return "<nil>"
  }
  return fmt.Sprintf("FormatDataDoFormatResult(%+v)", *p)
}


