// Autogenerated by Thrift Compiler (0.9.3)
// DO NOT EDIT UNLESS YOU ARE SURE THAT YOU KNOW WHAT YOU ARE DOING

package worker

import (
	"bytes"
	"fmt"
	"github.com/apache/thrift/lib/go/thrift"
	"github.com/twitter/scoot/bazel/execution/gen-go/bazel"
)

// (needed to ensure safety because of naive import list construction.)
var _ = thrift.ZERO
var _ = fmt.Printf
var _ = bytes.Equal

var _ = bazel.GoUnusedProtection__

type Worker interface {
	QueryWorker() (r *WorkerStatus, err error)
	// Parameters:
	//  - Cmd
	Run(cmd *RunCommand) (r *RunStatus, err error)
	// Parameters:
	//  - RunId
	Abort(runId string) (r *RunStatus, err error)
	// Parameters:
	//  - RunId
	Erase(runId string) (err error)
}

type WorkerClient struct {
	Transport       thrift.TTransport
	ProtocolFactory thrift.TProtocolFactory
	InputProtocol   thrift.TProtocol
	OutputProtocol  thrift.TProtocol
	SeqId           int32
}

func NewWorkerClientFactory(t thrift.TTransport, f thrift.TProtocolFactory) *WorkerClient {
	return &WorkerClient{Transport: t,
		ProtocolFactory: f,
		InputProtocol:   f.GetProtocol(t),
		OutputProtocol:  f.GetProtocol(t),
		SeqId:           0,
	}
}

func NewWorkerClientProtocol(t thrift.TTransport, iprot thrift.TProtocol, oprot thrift.TProtocol) *WorkerClient {
	return &WorkerClient{Transport: t,
		ProtocolFactory: nil,
		InputProtocol:   iprot,
		OutputProtocol:  oprot,
		SeqId:           0,
	}
}

func (p *WorkerClient) QueryWorker() (r *WorkerStatus, err error) {
	if err = p.sendQueryWorker(); err != nil {
		return
	}
	return p.recvQueryWorker()
}

func (p *WorkerClient) sendQueryWorker() (err error) {
	oprot := p.OutputProtocol
	if oprot == nil {
		oprot = p.ProtocolFactory.GetProtocol(p.Transport)
		p.OutputProtocol = oprot
	}
	p.SeqId++
	if err = oprot.WriteMessageBegin("QueryWorker", thrift.CALL, p.SeqId); err != nil {
		return
	}
	args := WorkerQueryWorkerArgs{}
	if err = args.Write(oprot); err != nil {
		return
	}
	if err = oprot.WriteMessageEnd(); err != nil {
		return
	}
	return oprot.Flush()
}

func (p *WorkerClient) recvQueryWorker() (value *WorkerStatus, err error) {
	iprot := p.InputProtocol
	if iprot == nil {
		iprot = p.ProtocolFactory.GetProtocol(p.Transport)
		p.InputProtocol = iprot
	}
	method, mTypeId, seqId, err := iprot.ReadMessageBegin()
	if err != nil {
		return
	}
	if method != "QueryWorker" {
		err = thrift.NewTApplicationException(thrift.WRONG_METHOD_NAME, "QueryWorker failed: wrong method name")
		return
	}
	if p.SeqId != seqId {
		err = thrift.NewTApplicationException(thrift.BAD_SEQUENCE_ID, "QueryWorker failed: out of sequence response")
		return
	}
	if mTypeId == thrift.EXCEPTION {
		error4 := thrift.NewTApplicationException(thrift.UNKNOWN_APPLICATION_EXCEPTION, "Unknown Exception")
		var error5 error
		error5, err = error4.Read(iprot)
		if err != nil {
			return
		}
		if err = iprot.ReadMessageEnd(); err != nil {
			return
		}
		err = error5
		return
	}
	if mTypeId != thrift.REPLY {
		err = thrift.NewTApplicationException(thrift.INVALID_MESSAGE_TYPE_EXCEPTION, "QueryWorker failed: invalid message type")
		return
	}
	result := WorkerQueryWorkerResult{}
	if err = result.Read(iprot); err != nil {
		return
	}
	if err = iprot.ReadMessageEnd(); err != nil {
		return
	}
	value = result.GetSuccess()
	return
}

// Parameters:
//  - Cmd
func (p *WorkerClient) Run(cmd *RunCommand) (r *RunStatus, err error) {
	if err = p.sendRun(cmd); err != nil {
		return
	}
	return p.recvRun()
}

func (p *WorkerClient) sendRun(cmd *RunCommand) (err error) {
	oprot := p.OutputProtocol
	if oprot == nil {
		oprot = p.ProtocolFactory.GetProtocol(p.Transport)
		p.OutputProtocol = oprot
	}
	p.SeqId++
	if err = oprot.WriteMessageBegin("Run", thrift.CALL, p.SeqId); err != nil {
		return
	}
	args := WorkerRunArgs{
		Cmd: cmd,
	}
	if err = args.Write(oprot); err != nil {
		return
	}
	if err = oprot.WriteMessageEnd(); err != nil {
		return
	}
	return oprot.Flush()
}

func (p *WorkerClient) recvRun() (value *RunStatus, err error) {
	iprot := p.InputProtocol
	if iprot == nil {
		iprot = p.ProtocolFactory.GetProtocol(p.Transport)
		p.InputProtocol = iprot
	}
	method, mTypeId, seqId, err := iprot.ReadMessageBegin()
	if err != nil {
		return
	}
	if method != "Run" {
		err = thrift.NewTApplicationException(thrift.WRONG_METHOD_NAME, "Run failed: wrong method name")
		return
	}
	if p.SeqId != seqId {
		err = thrift.NewTApplicationException(thrift.BAD_SEQUENCE_ID, "Run failed: out of sequence response")
		return
	}
	if mTypeId == thrift.EXCEPTION {
		error6 := thrift.NewTApplicationException(thrift.UNKNOWN_APPLICATION_EXCEPTION, "Unknown Exception")
		var error7 error
		error7, err = error6.Read(iprot)
		if err != nil {
			return
		}
		if err = iprot.ReadMessageEnd(); err != nil {
			return
		}
		err = error7
		return
	}
	if mTypeId != thrift.REPLY {
		err = thrift.NewTApplicationException(thrift.INVALID_MESSAGE_TYPE_EXCEPTION, "Run failed: invalid message type")
		return
	}
	result := WorkerRunResult{}
	if err = result.Read(iprot); err != nil {
		return
	}
	if err = iprot.ReadMessageEnd(); err != nil {
		return
	}
	value = result.GetSuccess()
	return
}

// Parameters:
//  - RunId
func (p *WorkerClient) Abort(runId string) (r *RunStatus, err error) {
	if err = p.sendAbort(runId); err != nil {
		return
	}
	return p.recvAbort()
}

func (p *WorkerClient) sendAbort(runId string) (err error) {
	oprot := p.OutputProtocol
	if oprot == nil {
		oprot = p.ProtocolFactory.GetProtocol(p.Transport)
		p.OutputProtocol = oprot
	}
	p.SeqId++
	if err = oprot.WriteMessageBegin("Abort", thrift.CALL, p.SeqId); err != nil {
		return
	}
	args := WorkerAbortArgs{
		RunId: runId,
	}
	if err = args.Write(oprot); err != nil {
		return
	}
	if err = oprot.WriteMessageEnd(); err != nil {
		return
	}
	return oprot.Flush()
}

func (p *WorkerClient) recvAbort() (value *RunStatus, err error) {
	iprot := p.InputProtocol
	if iprot == nil {
		iprot = p.ProtocolFactory.GetProtocol(p.Transport)
		p.InputProtocol = iprot
	}
	method, mTypeId, seqId, err := iprot.ReadMessageBegin()
	if err != nil {
		return
	}
	if method != "Abort" {
		err = thrift.NewTApplicationException(thrift.WRONG_METHOD_NAME, "Abort failed: wrong method name")
		return
	}
	if p.SeqId != seqId {
		err = thrift.NewTApplicationException(thrift.BAD_SEQUENCE_ID, "Abort failed: out of sequence response")
		return
	}
	if mTypeId == thrift.EXCEPTION {
		error8 := thrift.NewTApplicationException(thrift.UNKNOWN_APPLICATION_EXCEPTION, "Unknown Exception")
		var error9 error
		error9, err = error8.Read(iprot)
		if err != nil {
			return
		}
		if err = iprot.ReadMessageEnd(); err != nil {
			return
		}
		err = error9
		return
	}
	if mTypeId != thrift.REPLY {
		err = thrift.NewTApplicationException(thrift.INVALID_MESSAGE_TYPE_EXCEPTION, "Abort failed: invalid message type")
		return
	}
	result := WorkerAbortResult{}
	if err = result.Read(iprot); err != nil {
		return
	}
	if err = iprot.ReadMessageEnd(); err != nil {
		return
	}
	value = result.GetSuccess()
	return
}

// Parameters:
//  - RunId
func (p *WorkerClient) Erase(runId string) (err error) {
	if err = p.sendErase(runId); err != nil {
		return
	}
	return p.recvErase()
}

func (p *WorkerClient) sendErase(runId string) (err error) {
	oprot := p.OutputProtocol
	if oprot == nil {
		oprot = p.ProtocolFactory.GetProtocol(p.Transport)
		p.OutputProtocol = oprot
	}
	p.SeqId++
	if err = oprot.WriteMessageBegin("Erase", thrift.CALL, p.SeqId); err != nil {
		return
	}
	args := WorkerEraseArgs{
		RunId: runId,
	}
	if err = args.Write(oprot); err != nil {
		return
	}
	if err = oprot.WriteMessageEnd(); err != nil {
		return
	}
	return oprot.Flush()
}

func (p *WorkerClient) recvErase() (err error) {
	iprot := p.InputProtocol
	if iprot == nil {
		iprot = p.ProtocolFactory.GetProtocol(p.Transport)
		p.InputProtocol = iprot
	}
	method, mTypeId, seqId, err := iprot.ReadMessageBegin()
	if err != nil {
		return
	}
	if method != "Erase" {
		err = thrift.NewTApplicationException(thrift.WRONG_METHOD_NAME, "Erase failed: wrong method name")
		return
	}
	if p.SeqId != seqId {
		err = thrift.NewTApplicationException(thrift.BAD_SEQUENCE_ID, "Erase failed: out of sequence response")
		return
	}
	if mTypeId == thrift.EXCEPTION {
		error10 := thrift.NewTApplicationException(thrift.UNKNOWN_APPLICATION_EXCEPTION, "Unknown Exception")
		var error11 error
		error11, err = error10.Read(iprot)
		if err != nil {
			return
		}
		if err = iprot.ReadMessageEnd(); err != nil {
			return
		}
		err = error11
		return
	}
	if mTypeId != thrift.REPLY {
		err = thrift.NewTApplicationException(thrift.INVALID_MESSAGE_TYPE_EXCEPTION, "Erase failed: invalid message type")
		return
	}
	result := WorkerEraseResult{}
	if err = result.Read(iprot); err != nil {
		return
	}
	if err = iprot.ReadMessageEnd(); err != nil {
		return
	}
	return
}

type WorkerProcessor struct {
	processorMap map[string]thrift.TProcessorFunction
	handler      Worker
}

func (p *WorkerProcessor) AddToProcessorMap(key string, processor thrift.TProcessorFunction) {
	p.processorMap[key] = processor
}

func (p *WorkerProcessor) GetProcessorFunction(key string) (processor thrift.TProcessorFunction, ok bool) {
	processor, ok = p.processorMap[key]
	return processor, ok
}

func (p *WorkerProcessor) ProcessorMap() map[string]thrift.TProcessorFunction {
	return p.processorMap
}

func NewWorkerProcessor(handler Worker) *WorkerProcessor {

	self12 := &WorkerProcessor{handler: handler, processorMap: make(map[string]thrift.TProcessorFunction)}
	self12.processorMap["QueryWorker"] = &workerProcessorQueryWorker{handler: handler}
	self12.processorMap["Run"] = &workerProcessorRun{handler: handler}
	self12.processorMap["Abort"] = &workerProcessorAbort{handler: handler}
	self12.processorMap["Erase"] = &workerProcessorErase{handler: handler}
	return self12
}

func (p *WorkerProcessor) Process(iprot, oprot thrift.TProtocol) (success bool, err thrift.TException) {
	name, _, seqId, err := iprot.ReadMessageBegin()
	if err != nil {
		return false, err
	}
	if processor, ok := p.GetProcessorFunction(name); ok {
		return processor.Process(seqId, iprot, oprot)
	}
	iprot.Skip(thrift.STRUCT)
	iprot.ReadMessageEnd()
	x13 := thrift.NewTApplicationException(thrift.UNKNOWN_METHOD, "Unknown function "+name)
	oprot.WriteMessageBegin(name, thrift.EXCEPTION, seqId)
	x13.Write(oprot)
	oprot.WriteMessageEnd()
	oprot.Flush()
	return false, x13

}

type workerProcessorQueryWorker struct {
	handler Worker
}

func (p *workerProcessorQueryWorker) Process(seqId int32, iprot, oprot thrift.TProtocol) (success bool, err thrift.TException) {
	args := WorkerQueryWorkerArgs{}
	if err = args.Read(iprot); err != nil {
		iprot.ReadMessageEnd()
		x := thrift.NewTApplicationException(thrift.PROTOCOL_ERROR, err.Error())
		oprot.WriteMessageBegin("QueryWorker", thrift.EXCEPTION, seqId)
		x.Write(oprot)
		oprot.WriteMessageEnd()
		oprot.Flush()
		return false, err
	}

	iprot.ReadMessageEnd()
	result := WorkerQueryWorkerResult{}
	var retval *WorkerStatus
	var err2 error
	if retval, err2 = p.handler.QueryWorker(); err2 != nil {
		x := thrift.NewTApplicationException(thrift.INTERNAL_ERROR, "Internal error processing QueryWorker: "+err2.Error())
		oprot.WriteMessageBegin("QueryWorker", thrift.EXCEPTION, seqId)
		x.Write(oprot)
		oprot.WriteMessageEnd()
		oprot.Flush()
		return true, err2
	} else {
		result.Success = retval
	}
	if err2 = oprot.WriteMessageBegin("QueryWorker", thrift.REPLY, seqId); err2 != nil {
		err = err2
	}
	if err2 = result.Write(oprot); err == nil && err2 != nil {
		err = err2
	}
	if err2 = oprot.WriteMessageEnd(); err == nil && err2 != nil {
		err = err2
	}
	if err2 = oprot.Flush(); err == nil && err2 != nil {
		err = err2
	}
	if err != nil {
		return
	}
	return true, err
}

type workerProcessorRun struct {
	handler Worker
}

func (p *workerProcessorRun) Process(seqId int32, iprot, oprot thrift.TProtocol) (success bool, err thrift.TException) {
	args := WorkerRunArgs{}
	if err = args.Read(iprot); err != nil {
		iprot.ReadMessageEnd()
		x := thrift.NewTApplicationException(thrift.PROTOCOL_ERROR, err.Error())
		oprot.WriteMessageBegin("Run", thrift.EXCEPTION, seqId)
		x.Write(oprot)
		oprot.WriteMessageEnd()
		oprot.Flush()
		return false, err
	}

	iprot.ReadMessageEnd()
	result := WorkerRunResult{}
	var retval *RunStatus
	var err2 error
	if retval, err2 = p.handler.Run(args.Cmd); err2 != nil {
		x := thrift.NewTApplicationException(thrift.INTERNAL_ERROR, "Internal error processing Run: "+err2.Error())
		oprot.WriteMessageBegin("Run", thrift.EXCEPTION, seqId)
		x.Write(oprot)
		oprot.WriteMessageEnd()
		oprot.Flush()
		return true, err2
	} else {
		result.Success = retval
	}
	if err2 = oprot.WriteMessageBegin("Run", thrift.REPLY, seqId); err2 != nil {
		err = err2
	}
	if err2 = result.Write(oprot); err == nil && err2 != nil {
		err = err2
	}
	if err2 = oprot.WriteMessageEnd(); err == nil && err2 != nil {
		err = err2
	}
	if err2 = oprot.Flush(); err == nil && err2 != nil {
		err = err2
	}
	if err != nil {
		return
	}
	return true, err
}

type workerProcessorAbort struct {
	handler Worker
}

func (p *workerProcessorAbort) Process(seqId int32, iprot, oprot thrift.TProtocol) (success bool, err thrift.TException) {
	args := WorkerAbortArgs{}
	if err = args.Read(iprot); err != nil {
		iprot.ReadMessageEnd()
		x := thrift.NewTApplicationException(thrift.PROTOCOL_ERROR, err.Error())
		oprot.WriteMessageBegin("Abort", thrift.EXCEPTION, seqId)
		x.Write(oprot)
		oprot.WriteMessageEnd()
		oprot.Flush()
		return false, err
	}

	iprot.ReadMessageEnd()
	result := WorkerAbortResult{}
	var retval *RunStatus
	var err2 error
	if retval, err2 = p.handler.Abort(args.RunId); err2 != nil {
		x := thrift.NewTApplicationException(thrift.INTERNAL_ERROR, "Internal error processing Abort: "+err2.Error())
		oprot.WriteMessageBegin("Abort", thrift.EXCEPTION, seqId)
		x.Write(oprot)
		oprot.WriteMessageEnd()
		oprot.Flush()
		return true, err2
	} else {
		result.Success = retval
	}
	if err2 = oprot.WriteMessageBegin("Abort", thrift.REPLY, seqId); err2 != nil {
		err = err2
	}
	if err2 = result.Write(oprot); err == nil && err2 != nil {
		err = err2
	}
	if err2 = oprot.WriteMessageEnd(); err == nil && err2 != nil {
		err = err2
	}
	if err2 = oprot.Flush(); err == nil && err2 != nil {
		err = err2
	}
	if err != nil {
		return
	}
	return true, err
}

type workerProcessorErase struct {
	handler Worker
}

func (p *workerProcessorErase) Process(seqId int32, iprot, oprot thrift.TProtocol) (success bool, err thrift.TException) {
	args := WorkerEraseArgs{}
	if err = args.Read(iprot); err != nil {
		iprot.ReadMessageEnd()
		x := thrift.NewTApplicationException(thrift.PROTOCOL_ERROR, err.Error())
		oprot.WriteMessageBegin("Erase", thrift.EXCEPTION, seqId)
		x.Write(oprot)
		oprot.WriteMessageEnd()
		oprot.Flush()
		return false, err
	}

	iprot.ReadMessageEnd()
	result := WorkerEraseResult{}
	var err2 error
	if err2 = p.handler.Erase(args.RunId); err2 != nil {
		x := thrift.NewTApplicationException(thrift.INTERNAL_ERROR, "Internal error processing Erase: "+err2.Error())
		oprot.WriteMessageBegin("Erase", thrift.EXCEPTION, seqId)
		x.Write(oprot)
		oprot.WriteMessageEnd()
		oprot.Flush()
		return true, err2
	}
	if err2 = oprot.WriteMessageBegin("Erase", thrift.REPLY, seqId); err2 != nil {
		err = err2
	}
	if err2 = result.Write(oprot); err == nil && err2 != nil {
		err = err2
	}
	if err2 = oprot.WriteMessageEnd(); err == nil && err2 != nil {
		err = err2
	}
	if err2 = oprot.Flush(); err == nil && err2 != nil {
		err = err2
	}
	if err != nil {
		return
	}
	return true, err
}

// HELPER FUNCTIONS AND STRUCTURES

type WorkerQueryWorkerArgs struct {
}

func NewWorkerQueryWorkerArgs() *WorkerQueryWorkerArgs {
	return &WorkerQueryWorkerArgs{}
}

func (p *WorkerQueryWorkerArgs) Read(iprot thrift.TProtocol) error {
	if _, err := iprot.ReadStructBegin(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read error: ", p), err)
	}

	for {
		_, fieldTypeId, fieldId, err := iprot.ReadFieldBegin()
		if err != nil {
			return thrift.PrependError(fmt.Sprintf("%T field %d read error: ", p, fieldId), err)
		}
		if fieldTypeId == thrift.STOP {
			break
		}
		if err := iprot.Skip(fieldTypeId); err != nil {
			return err
		}
		if err := iprot.ReadFieldEnd(); err != nil {
			return err
		}
	}
	if err := iprot.ReadStructEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
	}
	return nil
}

func (p *WorkerQueryWorkerArgs) Write(oprot thrift.TProtocol) error {
	if err := oprot.WriteStructBegin("QueryWorker_args"); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err)
	}
	if err := oprot.WriteFieldStop(); err != nil {
		return thrift.PrependError("write field stop error: ", err)
	}
	if err := oprot.WriteStructEnd(); err != nil {
		return thrift.PrependError("write struct stop error: ", err)
	}
	return nil
}

func (p *WorkerQueryWorkerArgs) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("WorkerQueryWorkerArgs(%+v)", *p)
}

// Attributes:
//  - Success
type WorkerQueryWorkerResult struct {
	Success *WorkerStatus `thrift:"success,0" json:"success,omitempty"`
}

func NewWorkerQueryWorkerResult() *WorkerQueryWorkerResult {
	return &WorkerQueryWorkerResult{}
}

var WorkerQueryWorkerResult_Success_DEFAULT *WorkerStatus

func (p *WorkerQueryWorkerResult) GetSuccess() *WorkerStatus {
	if !p.IsSetSuccess() {
		return WorkerQueryWorkerResult_Success_DEFAULT
	}
	return p.Success
}
func (p *WorkerQueryWorkerResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *WorkerQueryWorkerResult) Read(iprot thrift.TProtocol) error {
	if _, err := iprot.ReadStructBegin(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read error: ", p), err)
	}

	for {
		_, fieldTypeId, fieldId, err := iprot.ReadFieldBegin()
		if err != nil {
			return thrift.PrependError(fmt.Sprintf("%T field %d read error: ", p, fieldId), err)
		}
		if fieldTypeId == thrift.STOP {
			break
		}
		switch fieldId {
		case 0:
			if err := p.readField0(iprot); err != nil {
				return err
			}
		default:
			if err := iprot.Skip(fieldTypeId); err != nil {
				return err
			}
		}
		if err := iprot.ReadFieldEnd(); err != nil {
			return err
		}
	}
	if err := iprot.ReadStructEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
	}
	return nil
}

func (p *WorkerQueryWorkerResult) readField0(iprot thrift.TProtocol) error {
	p.Success = &WorkerStatus{}
	if err := p.Success.Read(iprot); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T error reading struct: ", p.Success), err)
	}
	return nil
}

func (p *WorkerQueryWorkerResult) Write(oprot thrift.TProtocol) error {
	if err := oprot.WriteStructBegin("QueryWorker_result"); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err)
	}
	if err := p.writeField0(oprot); err != nil {
		return err
	}
	if err := oprot.WriteFieldStop(); err != nil {
		return thrift.PrependError("write field stop error: ", err)
	}
	if err := oprot.WriteStructEnd(); err != nil {
		return thrift.PrependError("write struct stop error: ", err)
	}
	return nil
}

func (p *WorkerQueryWorkerResult) writeField0(oprot thrift.TProtocol) (err error) {
	if p.IsSetSuccess() {
		if err := oprot.WriteFieldBegin("success", thrift.STRUCT, 0); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field begin error 0:success: ", p), err)
		}
		if err := p.Success.Write(oprot); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T error writing struct: ", p.Success), err)
		}
		if err := oprot.WriteFieldEnd(); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field end error 0:success: ", p), err)
		}
	}
	return err
}

func (p *WorkerQueryWorkerResult) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("WorkerQueryWorkerResult(%+v)", *p)
}

// Attributes:
//  - Cmd
type WorkerRunArgs struct {
	Cmd *RunCommand `thrift:"cmd,1" json:"cmd"`
}

func NewWorkerRunArgs() *WorkerRunArgs {
	return &WorkerRunArgs{}
}

var WorkerRunArgs_Cmd_DEFAULT *RunCommand

func (p *WorkerRunArgs) GetCmd() *RunCommand {
	if !p.IsSetCmd() {
		return WorkerRunArgs_Cmd_DEFAULT
	}
	return p.Cmd
}
func (p *WorkerRunArgs) IsSetCmd() bool {
	return p.Cmd != nil
}

func (p *WorkerRunArgs) Read(iprot thrift.TProtocol) error {
	if _, err := iprot.ReadStructBegin(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read error: ", p), err)
	}

	for {
		_, fieldTypeId, fieldId, err := iprot.ReadFieldBegin()
		if err != nil {
			return thrift.PrependError(fmt.Sprintf("%T field %d read error: ", p, fieldId), err)
		}
		if fieldTypeId == thrift.STOP {
			break
		}
		switch fieldId {
		case 1:
			if err := p.readField1(iprot); err != nil {
				return err
			}
		default:
			if err := iprot.Skip(fieldTypeId); err != nil {
				return err
			}
		}
		if err := iprot.ReadFieldEnd(); err != nil {
			return err
		}
	}
	if err := iprot.ReadStructEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
	}
	return nil
}

func (p *WorkerRunArgs) readField1(iprot thrift.TProtocol) error {
	p.Cmd = &RunCommand{}
	if err := p.Cmd.Read(iprot); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T error reading struct: ", p.Cmd), err)
	}
	return nil
}

func (p *WorkerRunArgs) Write(oprot thrift.TProtocol) error {
	if err := oprot.WriteStructBegin("Run_args"); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err)
	}
	if err := p.writeField1(oprot); err != nil {
		return err
	}
	if err := oprot.WriteFieldStop(); err != nil {
		return thrift.PrependError("write field stop error: ", err)
	}
	if err := oprot.WriteStructEnd(); err != nil {
		return thrift.PrependError("write struct stop error: ", err)
	}
	return nil
}

func (p *WorkerRunArgs) writeField1(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("cmd", thrift.STRUCT, 1); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 1:cmd: ", p), err)
	}
	if err := p.Cmd.Write(oprot); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T error writing struct: ", p.Cmd), err)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 1:cmd: ", p), err)
	}
	return err
}

func (p *WorkerRunArgs) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("WorkerRunArgs(%+v)", *p)
}

// Attributes:
//  - Success
type WorkerRunResult struct {
	Success *RunStatus `thrift:"success,0" json:"success,omitempty"`
}

func NewWorkerRunResult() *WorkerRunResult {
	return &WorkerRunResult{}
}

var WorkerRunResult_Success_DEFAULT *RunStatus

func (p *WorkerRunResult) GetSuccess() *RunStatus {
	if !p.IsSetSuccess() {
		return WorkerRunResult_Success_DEFAULT
	}
	return p.Success
}
func (p *WorkerRunResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *WorkerRunResult) Read(iprot thrift.TProtocol) error {
	if _, err := iprot.ReadStructBegin(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read error: ", p), err)
	}

	for {
		_, fieldTypeId, fieldId, err := iprot.ReadFieldBegin()
		if err != nil {
			return thrift.PrependError(fmt.Sprintf("%T field %d read error: ", p, fieldId), err)
		}
		if fieldTypeId == thrift.STOP {
			break
		}
		switch fieldId {
		case 0:
			if err := p.readField0(iprot); err != nil {
				return err
			}
		default:
			if err := iprot.Skip(fieldTypeId); err != nil {
				return err
			}
		}
		if err := iprot.ReadFieldEnd(); err != nil {
			return err
		}
	}
	if err := iprot.ReadStructEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
	}
	return nil
}

func (p *WorkerRunResult) readField0(iprot thrift.TProtocol) error {
	p.Success = &RunStatus{}
	if err := p.Success.Read(iprot); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T error reading struct: ", p.Success), err)
	}
	return nil
}

func (p *WorkerRunResult) Write(oprot thrift.TProtocol) error {
	if err := oprot.WriteStructBegin("Run_result"); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err)
	}
	if err := p.writeField0(oprot); err != nil {
		return err
	}
	if err := oprot.WriteFieldStop(); err != nil {
		return thrift.PrependError("write field stop error: ", err)
	}
	if err := oprot.WriteStructEnd(); err != nil {
		return thrift.PrependError("write struct stop error: ", err)
	}
	return nil
}

func (p *WorkerRunResult) writeField0(oprot thrift.TProtocol) (err error) {
	if p.IsSetSuccess() {
		if err := oprot.WriteFieldBegin("success", thrift.STRUCT, 0); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field begin error 0:success: ", p), err)
		}
		if err := p.Success.Write(oprot); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T error writing struct: ", p.Success), err)
		}
		if err := oprot.WriteFieldEnd(); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field end error 0:success: ", p), err)
		}
	}
	return err
}

func (p *WorkerRunResult) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("WorkerRunResult(%+v)", *p)
}

// Attributes:
//  - RunId
type WorkerAbortArgs struct {
	RunId string `thrift:"runId,1" json:"runId"`
}

func NewWorkerAbortArgs() *WorkerAbortArgs {
	return &WorkerAbortArgs{}
}

func (p *WorkerAbortArgs) GetRunId() string {
	return p.RunId
}
func (p *WorkerAbortArgs) Read(iprot thrift.TProtocol) error {
	if _, err := iprot.ReadStructBegin(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read error: ", p), err)
	}

	for {
		_, fieldTypeId, fieldId, err := iprot.ReadFieldBegin()
		if err != nil {
			return thrift.PrependError(fmt.Sprintf("%T field %d read error: ", p, fieldId), err)
		}
		if fieldTypeId == thrift.STOP {
			break
		}
		switch fieldId {
		case 1:
			if err := p.readField1(iprot); err != nil {
				return err
			}
		default:
			if err := iprot.Skip(fieldTypeId); err != nil {
				return err
			}
		}
		if err := iprot.ReadFieldEnd(); err != nil {
			return err
		}
	}
	if err := iprot.ReadStructEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
	}
	return nil
}

func (p *WorkerAbortArgs) readField1(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadString(); err != nil {
		return thrift.PrependError("error reading field 1: ", err)
	} else {
		p.RunId = v
	}
	return nil
}

func (p *WorkerAbortArgs) Write(oprot thrift.TProtocol) error {
	if err := oprot.WriteStructBegin("Abort_args"); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err)
	}
	if err := p.writeField1(oprot); err != nil {
		return err
	}
	if err := oprot.WriteFieldStop(); err != nil {
		return thrift.PrependError("write field stop error: ", err)
	}
	if err := oprot.WriteStructEnd(); err != nil {
		return thrift.PrependError("write struct stop error: ", err)
	}
	return nil
}

func (p *WorkerAbortArgs) writeField1(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("runId", thrift.STRING, 1); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 1:runId: ", p), err)
	}
	if err := oprot.WriteString(string(p.RunId)); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T.runId (1) field write error: ", p), err)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 1:runId: ", p), err)
	}
	return err
}

func (p *WorkerAbortArgs) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("WorkerAbortArgs(%+v)", *p)
}

// Attributes:
//  - Success
type WorkerAbortResult struct {
	Success *RunStatus `thrift:"success,0" json:"success,omitempty"`
}

func NewWorkerAbortResult() *WorkerAbortResult {
	return &WorkerAbortResult{}
}

var WorkerAbortResult_Success_DEFAULT *RunStatus

func (p *WorkerAbortResult) GetSuccess() *RunStatus {
	if !p.IsSetSuccess() {
		return WorkerAbortResult_Success_DEFAULT
	}
	return p.Success
}
func (p *WorkerAbortResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *WorkerAbortResult) Read(iprot thrift.TProtocol) error {
	if _, err := iprot.ReadStructBegin(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read error: ", p), err)
	}

	for {
		_, fieldTypeId, fieldId, err := iprot.ReadFieldBegin()
		if err != nil {
			return thrift.PrependError(fmt.Sprintf("%T field %d read error: ", p, fieldId), err)
		}
		if fieldTypeId == thrift.STOP {
			break
		}
		switch fieldId {
		case 0:
			if err := p.readField0(iprot); err != nil {
				return err
			}
		default:
			if err := iprot.Skip(fieldTypeId); err != nil {
				return err
			}
		}
		if err := iprot.ReadFieldEnd(); err != nil {
			return err
		}
	}
	if err := iprot.ReadStructEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
	}
	return nil
}

func (p *WorkerAbortResult) readField0(iprot thrift.TProtocol) error {
	p.Success = &RunStatus{}
	if err := p.Success.Read(iprot); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T error reading struct: ", p.Success), err)
	}
	return nil
}

func (p *WorkerAbortResult) Write(oprot thrift.TProtocol) error {
	if err := oprot.WriteStructBegin("Abort_result"); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err)
	}
	if err := p.writeField0(oprot); err != nil {
		return err
	}
	if err := oprot.WriteFieldStop(); err != nil {
		return thrift.PrependError("write field stop error: ", err)
	}
	if err := oprot.WriteStructEnd(); err != nil {
		return thrift.PrependError("write struct stop error: ", err)
	}
	return nil
}

func (p *WorkerAbortResult) writeField0(oprot thrift.TProtocol) (err error) {
	if p.IsSetSuccess() {
		if err := oprot.WriteFieldBegin("success", thrift.STRUCT, 0); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field begin error 0:success: ", p), err)
		}
		if err := p.Success.Write(oprot); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T error writing struct: ", p.Success), err)
		}
		if err := oprot.WriteFieldEnd(); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field end error 0:success: ", p), err)
		}
	}
	return err
}

func (p *WorkerAbortResult) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("WorkerAbortResult(%+v)", *p)
}

// Attributes:
//  - RunId
type WorkerEraseArgs struct {
	RunId string `thrift:"runId,1" json:"runId"`
}

func NewWorkerEraseArgs() *WorkerEraseArgs {
	return &WorkerEraseArgs{}
}

func (p *WorkerEraseArgs) GetRunId() string {
	return p.RunId
}
func (p *WorkerEraseArgs) Read(iprot thrift.TProtocol) error {
	if _, err := iprot.ReadStructBegin(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read error: ", p), err)
	}

	for {
		_, fieldTypeId, fieldId, err := iprot.ReadFieldBegin()
		if err != nil {
			return thrift.PrependError(fmt.Sprintf("%T field %d read error: ", p, fieldId), err)
		}
		if fieldTypeId == thrift.STOP {
			break
		}
		switch fieldId {
		case 1:
			if err := p.readField1(iprot); err != nil {
				return err
			}
		default:
			if err := iprot.Skip(fieldTypeId); err != nil {
				return err
			}
		}
		if err := iprot.ReadFieldEnd(); err != nil {
			return err
		}
	}
	if err := iprot.ReadStructEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
	}
	return nil
}

func (p *WorkerEraseArgs) readField1(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadString(); err != nil {
		return thrift.PrependError("error reading field 1: ", err)
	} else {
		p.RunId = v
	}
	return nil
}

func (p *WorkerEraseArgs) Write(oprot thrift.TProtocol) error {
	if err := oprot.WriteStructBegin("Erase_args"); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err)
	}
	if err := p.writeField1(oprot); err != nil {
		return err
	}
	if err := oprot.WriteFieldStop(); err != nil {
		return thrift.PrependError("write field stop error: ", err)
	}
	if err := oprot.WriteStructEnd(); err != nil {
		return thrift.PrependError("write struct stop error: ", err)
	}
	return nil
}

func (p *WorkerEraseArgs) writeField1(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("runId", thrift.STRING, 1); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 1:runId: ", p), err)
	}
	if err := oprot.WriteString(string(p.RunId)); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T.runId (1) field write error: ", p), err)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 1:runId: ", p), err)
	}
	return err
}

func (p *WorkerEraseArgs) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("WorkerEraseArgs(%+v)", *p)
}

type WorkerEraseResult struct {
}

func NewWorkerEraseResult() *WorkerEraseResult {
	return &WorkerEraseResult{}
}

func (p *WorkerEraseResult) Read(iprot thrift.TProtocol) error {
	if _, err := iprot.ReadStructBegin(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read error: ", p), err)
	}

	for {
		_, fieldTypeId, fieldId, err := iprot.ReadFieldBegin()
		if err != nil {
			return thrift.PrependError(fmt.Sprintf("%T field %d read error: ", p, fieldId), err)
		}
		if fieldTypeId == thrift.STOP {
			break
		}
		if err := iprot.Skip(fieldTypeId); err != nil {
			return err
		}
		if err := iprot.ReadFieldEnd(); err != nil {
			return err
		}
	}
	if err := iprot.ReadStructEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
	}
	return nil
}

func (p *WorkerEraseResult) Write(oprot thrift.TProtocol) error {
	if err := oprot.WriteStructBegin("Erase_result"); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err)
	}
	if err := oprot.WriteFieldStop(); err != nil {
		return thrift.PrependError("write field stop error: ", err)
	}
	if err := oprot.WriteStructEnd(); err != nil {
		return thrift.PrependError("write struct stop error: ", err)
	}
	return nil
}

func (p *WorkerEraseResult) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("WorkerEraseResult(%+v)", *p)
}
