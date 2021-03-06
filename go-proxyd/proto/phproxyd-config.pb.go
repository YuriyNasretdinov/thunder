// Code generated by protoc-gen-gogo.
// source: phproxyd-config.proto
// DO NOT EDIT!

/*
Package badoo_phproxyd is a generated protocol buffer package.

It is generated from these files:
	phproxyd-config.proto

It has these top-level messages:
	PhproxydConfig
*/
package badoo_phproxyd

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type PhproxydConfig struct {
	Scripts          *PhproxydConfigScriptsT `protobuf:"bytes,2,req,name=scripts" json:"scripts,omitempty"`
	XXX_unrecognized []byte                  `json:"-"`
}

func (m *PhproxydConfig) Reset()         { *m = PhproxydConfig{} }
func (m *PhproxydConfig) String() string { return proto.CompactTextString(m) }
func (*PhproxydConfig) ProtoMessage()    {}

func (m *PhproxydConfig) GetScripts() *PhproxydConfigScriptsT {
	if m != nil {
		return m.Scripts
	}
	return nil
}

type PhproxydConfigScriptsT struct {
	PhpBin             *string `protobuf:"bytes,1,req,name=php_bin" json:"php_bin,omitempty"`
	PhpScriptsDir      *string `protobuf:"bytes,2,req,name=php_scripts_dir" json:"php_scripts_dir,omitempty"`
	FileStdoutTemplate *string `protobuf:"bytes,3,req,name=file_stdout_template" json:"file_stdout_template,omitempty"`
	FileStderrTemplate *string `protobuf:"bytes,4,req,name=file_stderr_template" json:"file_stderr_template,omitempty"`
	MemStdoutTemplate  *string `protobuf:"bytes,5,req,name=mem_stdout_template" json:"mem_stdout_template,omitempty"`
	MemStderrTemplate  *string `protobuf:"bytes,6,req,name=mem_stderr_template" json:"mem_stderr_template,omitempty"`
	XXX_unrecognized   []byte  `json:"-"`
}

func (m *PhproxydConfigScriptsT) Reset()         { *m = PhproxydConfigScriptsT{} }
func (m *PhproxydConfigScriptsT) String() string { return proto.CompactTextString(m) }
func (*PhproxydConfigScriptsT) ProtoMessage()    {}

func (m *PhproxydConfigScriptsT) GetPhpBin() string {
	if m != nil && m.PhpBin != nil {
		return *m.PhpBin
	}
	return ""
}

func (m *PhproxydConfigScriptsT) GetPhpScriptsDir() string {
	if m != nil && m.PhpScriptsDir != nil {
		return *m.PhpScriptsDir
	}
	return ""
}

func (m *PhproxydConfigScriptsT) GetFileStdoutTemplate() string {
	if m != nil && m.FileStdoutTemplate != nil {
		return *m.FileStdoutTemplate
	}
	return ""
}

func (m *PhproxydConfigScriptsT) GetFileStderrTemplate() string {
	if m != nil && m.FileStderrTemplate != nil {
		return *m.FileStderrTemplate
	}
	return ""
}

func (m *PhproxydConfigScriptsT) GetMemStdoutTemplate() string {
	if m != nil && m.MemStdoutTemplate != nil {
		return *m.MemStdoutTemplate
	}
	return ""
}

func (m *PhproxydConfigScriptsT) GetMemStderrTemplate() string {
	if m != nil && m.MemStderrTemplate != nil {
		return *m.MemStderrTemplate
	}
	return ""
}
