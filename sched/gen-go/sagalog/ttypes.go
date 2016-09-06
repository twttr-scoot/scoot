// Autogenerated by Thrift Compiler (0.9.3)
// DO NOT EDIT UNLESS YOU ARE SURE THAT YOU KNOW WHAT YOU ARE DOING

package sagalog

import (
	"bytes"
	"fmt"
	"github.com/apache/thrift/lib/go/thrift"
)

// (needed to ensure safety because of naive import list construction.)
var _ = thrift.ZERO
var _ = fmt.Printf
var _ = bytes.Equal

var GoUnusedProtection__ int

// Attributes:
//  - Argv
//  - EnvVars
//  - Timeout
//  - SnapshotId
type Command struct {
	Argv       []string          `thrift:"argv,1" json:"argv"`
	EnvVars    map[string]string `thrift:"envVars,2" json:"envVars"`
	Timeout    int64             `thrift:"timeout,3" json:"timeout"`
	SnapshotId string            `thrift:"snapshotId,4" json:"snapshotId"`
}

func NewCommand() *Command {
	return &Command{}
}

func (p *Command) GetArgv() []string {
	return p.Argv
}

func (p *Command) GetEnvVars() map[string]string {
	return p.EnvVars
}

func (p *Command) GetTimeout() int64 {
	return p.Timeout
}

func (p *Command) GetSnapshotId() string {
	return p.SnapshotId
}
func (p *Command) Read(iprot thrift.TProtocol) error {
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
		case 2:
			if err := p.readField2(iprot); err != nil {
				return err
			}
		case 3:
			if err := p.readField3(iprot); err != nil {
				return err
			}
		case 4:
			if err := p.readField4(iprot); err != nil {
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

func (p *Command) readField1(iprot thrift.TProtocol) error {
	_, size, err := iprot.ReadListBegin()
	if err != nil {
		return thrift.PrependError("error reading list begin: ", err)
	}
	tSlice := make([]string, 0, size)
	p.Argv = tSlice
	for i := 0; i < size; i++ {
		var _elem0 string
		if v, err := iprot.ReadString(); err != nil {
			return thrift.PrependError("error reading field 0: ", err)
		} else {
			_elem0 = v
		}
		p.Argv = append(p.Argv, _elem0)
	}
	if err := iprot.ReadListEnd(); err != nil {
		return thrift.PrependError("error reading list end: ", err)
	}
	return nil
}

func (p *Command) readField2(iprot thrift.TProtocol) error {
	_, _, size, err := iprot.ReadMapBegin()
	if err != nil {
		return thrift.PrependError("error reading map begin: ", err)
	}
	tMap := make(map[string]string, size)
	p.EnvVars = tMap
	for i := 0; i < size; i++ {
		var _key1 string
		if v, err := iprot.ReadString(); err != nil {
			return thrift.PrependError("error reading field 0: ", err)
		} else {
			_key1 = v
		}
		var _val2 string
		if v, err := iprot.ReadString(); err != nil {
			return thrift.PrependError("error reading field 0: ", err)
		} else {
			_val2 = v
		}
		p.EnvVars[_key1] = _val2
	}
	if err := iprot.ReadMapEnd(); err != nil {
		return thrift.PrependError("error reading map end: ", err)
	}
	return nil
}

func (p *Command) readField3(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadI64(); err != nil {
		return thrift.PrependError("error reading field 3: ", err)
	} else {
		p.Timeout = v
	}
	return nil
}

func (p *Command) readField4(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadString(); err != nil {
		return thrift.PrependError("error reading field 4: ", err)
	} else {
		p.SnapshotId = v
	}
	return nil
}

func (p *Command) Write(oprot thrift.TProtocol) error {
	if err := oprot.WriteStructBegin("Command"); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err)
	}
	if err := p.writeField1(oprot); err != nil {
		return err
	}
	if err := p.writeField2(oprot); err != nil {
		return err
	}
	if err := p.writeField3(oprot); err != nil {
		return err
	}
	if err := p.writeField4(oprot); err != nil {
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

func (p *Command) writeField1(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("argv", thrift.LIST, 1); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 1:argv: ", p), err)
	}
	if err := oprot.WriteListBegin(thrift.STRING, len(p.Argv)); err != nil {
		return thrift.PrependError("error writing list begin: ", err)
	}
	for _, v := range p.Argv {
		if err := oprot.WriteString(string(v)); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T. (0) field write error: ", p), err)
		}
	}
	if err := oprot.WriteListEnd(); err != nil {
		return thrift.PrependError("error writing list end: ", err)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 1:argv: ", p), err)
	}
	return err
}

func (p *Command) writeField2(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("envVars", thrift.MAP, 2); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 2:envVars: ", p), err)
	}
	if err := oprot.WriteMapBegin(thrift.STRING, thrift.STRING, len(p.EnvVars)); err != nil {
		return thrift.PrependError("error writing map begin: ", err)
	}
	for k, v := range p.EnvVars {
		if err := oprot.WriteString(string(k)); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T. (0) field write error: ", p), err)
		}
		if err := oprot.WriteString(string(v)); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T. (0) field write error: ", p), err)
		}
	}
	if err := oprot.WriteMapEnd(); err != nil {
		return thrift.PrependError("error writing map end: ", err)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 2:envVars: ", p), err)
	}
	return err
}

func (p *Command) writeField3(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("timeout", thrift.I64, 3); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 3:timeout: ", p), err)
	}
	if err := oprot.WriteI64(int64(p.Timeout)); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T.timeout (3) field write error: ", p), err)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 3:timeout: ", p), err)
	}
	return err
}

func (p *Command) writeField4(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("snapshotId", thrift.STRING, 4); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 4:snapshotId: ", p), err)
	}
	if err := oprot.WriteString(string(p.SnapshotId)); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T.snapshotId (4) field write error: ", p), err)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 4:snapshotId: ", p), err)
	}
	return err
}

func (p *Command) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("Command(%+v)", *p)
}

// Attributes:
//  - Command
type TaskDefinition struct {
	Command *Command `thrift:"command,1" json:"command"`
}

func NewTaskDefinition() *TaskDefinition {
	return &TaskDefinition{}
}

var TaskDefinition_Command_DEFAULT *Command

func (p *TaskDefinition) GetCommand() *Command {
	if !p.IsSetCommand() {
		return TaskDefinition_Command_DEFAULT
	}
	return p.Command
}
func (p *TaskDefinition) IsSetCommand() bool {
	return p.Command != nil
}

func (p *TaskDefinition) Read(iprot thrift.TProtocol) error {
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

func (p *TaskDefinition) readField1(iprot thrift.TProtocol) error {
	p.Command = &Command{}
	if err := p.Command.Read(iprot); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T error reading struct: ", p.Command), err)
	}
	return nil
}

func (p *TaskDefinition) Write(oprot thrift.TProtocol) error {
	if err := oprot.WriteStructBegin("TaskDefinition"); err != nil {
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

func (p *TaskDefinition) writeField1(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("command", thrift.STRUCT, 1); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 1:command: ", p), err)
	}
	if err := p.Command.Write(oprot); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T error writing struct: ", p.Command), err)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 1:command: ", p), err)
	}
	return err
}

func (p *TaskDefinition) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("TaskDefinition(%+v)", *p)
}

// Attributes:
//  - JobType
//  - Tasks
type JobDefinition struct {
	JobType string                     `thrift:"jobType,1" json:"jobType"`
	Tasks   map[string]*TaskDefinition `thrift:"tasks,2" json:"tasks"`
}

func NewJobDefinition() *JobDefinition {
	return &JobDefinition{}
}

func (p *JobDefinition) GetJobType() string {
	return p.JobType
}

func (p *JobDefinition) GetTasks() map[string]*TaskDefinition {
	return p.Tasks
}
func (p *JobDefinition) Read(iprot thrift.TProtocol) error {
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
		case 2:
			if err := p.readField2(iprot); err != nil {
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

func (p *JobDefinition) readField1(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadString(); err != nil {
		return thrift.PrependError("error reading field 1: ", err)
	} else {
		p.JobType = v
	}
	return nil
}

func (p *JobDefinition) readField2(iprot thrift.TProtocol) error {
	_, _, size, err := iprot.ReadMapBegin()
	if err != nil {
		return thrift.PrependError("error reading map begin: ", err)
	}
	tMap := make(map[string]*TaskDefinition, size)
	p.Tasks = tMap
	for i := 0; i < size; i++ {
		var _key3 string
		if v, err := iprot.ReadString(); err != nil {
			return thrift.PrependError("error reading field 0: ", err)
		} else {
			_key3 = v
		}
		_val4 := &TaskDefinition{}
		if err := _val4.Read(iprot); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T error reading struct: ", _val4), err)
		}
		p.Tasks[_key3] = _val4
	}
	if err := iprot.ReadMapEnd(); err != nil {
		return thrift.PrependError("error reading map end: ", err)
	}
	return nil
}

func (p *JobDefinition) Write(oprot thrift.TProtocol) error {
	if err := oprot.WriteStructBegin("JobDefinition"); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err)
	}
	if err := p.writeField1(oprot); err != nil {
		return err
	}
	if err := p.writeField2(oprot); err != nil {
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

func (p *JobDefinition) writeField1(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("jobType", thrift.STRING, 1); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 1:jobType: ", p), err)
	}
	if err := oprot.WriteString(string(p.JobType)); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T.jobType (1) field write error: ", p), err)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 1:jobType: ", p), err)
	}
	return err
}

func (p *JobDefinition) writeField2(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("tasks", thrift.MAP, 2); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 2:tasks: ", p), err)
	}
	if err := oprot.WriteMapBegin(thrift.STRING, thrift.STRUCT, len(p.Tasks)); err != nil {
		return thrift.PrependError("error writing map begin: ", err)
	}
	for k, v := range p.Tasks {
		if err := oprot.WriteString(string(k)); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T. (0) field write error: ", p), err)
		}
		if err := v.Write(oprot); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T error writing struct: ", v), err)
		}
	}
	if err := oprot.WriteMapEnd(); err != nil {
		return thrift.PrependError("error writing map end: ", err)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 2:tasks: ", p), err)
	}
	return err
}

func (p *JobDefinition) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("JobDefinition(%+v)", *p)
}
