// RFC5424 The Syslog Protocol: https://www.rfc-editor.org/rfc/rfc5424
package log

import (
	"fmt"
	"github.com/a-skua/busybox-go/option"
	"os"
	"strconv"
	"strings"
	"time"
)

type Logger interface {
	Emergency(msg ...any) error
	Alert(msg ...any) error
	Critical(msg ...any) error
	Error(msg ...any) error
	Warning(msg ...any) error
	Notice(msg ...any) error
	Informational(msg ...any) error
	Debug(msg ...any) error
}

type Facility uint8

const (
	FacilityKernelMessages Facility = iota
	FacilityUserLevelMessages
	FacilityMailSystem
	FacilitySystemDaemons
	FacilitySecurityOrAuthorizationMessages0
	FacilityMessagesGeneratedInternallyBySyslogd
	FacilityLinePrinterSubsystem
	FacilityNetworkNewsSubsystem
	FacilityUUCPSubsystem
	FacilityClockDaemon0
	FacilitySecurityOrAuthorizationMessages1
	FacilityFTPDaemon
	FacilityNTPSubsystem
	FacilityLogAudit
	FacilityLogAlert
	FacilityClockDaemon1
	FacilityLocalUse0
	FacilityLocalUse1
	FacilityLocalUse2
	FacilityLocalUse3
	FacilityLocalUse4
	FacilityLocalUse5
	FacilityLocalUse6
	FacilityLocalUse7
)

type Severity uint8

const (
	SeverityEmergency Severity = iota
	SeverityAlert
	SeverityCritical
	SeverityError
	SeverityWarning
	SeverityNotice
	SeverityInformational
	SeverityDebug
)

type Priority struct {
	Facility Facility
	Severity Severity
}

func NewPriority(f Facility, s Severity) Priority {
	return Priority{
		Facility: f,
		Severity: s,
	}
}

func (pri Priority) String() string {
	return "<" + strconv.Itoa(int(pri.Num())) + ">"
}

func (pri Priority) Num() uint8 {
	return uint8(pri.Facility)*8 + uint8(pri.Severity)
}

type Version uint8

type Timestamp time.Time

func TimestampNow() Timestamp {
	return Timestamp(time.Now())
}

func (t Timestamp) String() string {
	return time.Time(t).Format(time.RFC3339Nano)
}

func (t Timestamp) MarshalJSON() ([]byte, error) {
	return time.Time(t).MarshalJSON()
}

type HostName string

func (host HostName) String() string {
	return string(host)
}

type AppName string

func (app AppName) String() string {
	return string(app)
}

type ProcessID string

func (proc ProcessID) String() string {
	return string(proc)
}

type MessageID string

func (msg MessageID) String() string {
	return string(msg)
}

func optionToString[T fmt.Stringer](v option.Option[T]) string {
	if v.Valid {
		return v.Value.String()
	}
	return "-"
}

type Header struct {
	Priority  Priority
	Version   Version
	Timestamp option.Option[Timestamp]
	Host      option.Option[HostName]
	App       option.Option[AppName]
	ProcessID option.Option[ProcessID]
	MessageID option.Option[MessageID]
}

func NewHeader(pri Priority, ver Version, time option.Option[Timestamp], host option.Option[HostName], app option.Option[AppName], proc option.Option[ProcessID], msg option.Option[MessageID]) Header {
	return Header{
		Priority:  pri,
		Version:   ver,
		Timestamp: time,
		Host:      host,
		App:       app,
		ProcessID: proc,
		MessageID: msg,
	}
}

func (h Header) String() string {
	return h.Priority.String() +
		strconv.Itoa(int(h.Version)) +
		" " +
		optionToString(h.Timestamp) +
		" " +
		optionToString(h.Host) +
		" " +
		optionToString(h.App) +
		" " +
		optionToString(h.ProcessID) +
		" " +
		optionToString(h.MessageID)
}

type MetadataID string

func (id MetadataID) String() string {
	return string(id)
}

type MetadataName string

func (name MetadataName) String() string {
	return string(name)
}

type MetadataValue string

func (value MetadataValue) String() string {
	return string(value)
}

type MetadataParam struct {
	Name  MetadataName
	Value MetadataValue
}

func NewMetadataParam(name MetadataName, value MetadataValue) MetadataParam {
	return MetadataParam{
		Name:  name,
		Value: value,
	}
}

func (param MetadataParam) String() string {
	return param.Name.String() +
		"=\"" +
		strings.ReplaceAll(strings.ReplaceAll(strings.ReplaceAll(param.Value.String(), "\\", "\\\\"), "\"", "\\\""), "]", "\\]") +
		"\""
}

type Metadata struct {
	ID     MetadataID
	Params []MetadataParam
}

func NewMetadata(id MetadataID, params ...MetadataParam) Metadata {
	return Metadata{
		ID:     id,
		Params: params,
	}
}

func (meta Metadata) String() string {
	params := ""
	for _, param := range meta.Params {
		params += " " + param.String()
	}

	return "[" + meta.ID.String() + params + "]"
}

type Message struct {
	Header   Header
	Metadata []Metadata
	Message  []any
}

func (msg *Message) String() string {
	message := ""
	for _, msg := range msg.Message {
		message += " " + fmt.Sprint(msg)
	}

	metadata := ""
	if len(msg.Metadata) > 0 {
		for _, meta := range msg.Metadata {
			metadata += meta.String()
		}
	} else {
		metadata = "-"
	}

	return msg.Header.String() +
		" " +
		metadata +
		message
}

func NewMessage(head Header, meta []Metadata, msg ...any) *Message {
	return &Message{
		Header:   head,
		Metadata: meta,
		Message:  msg,
	}
}

type Writer interface {
	Write(*Message) error
}

type stdWriter struct{}

func (w stdWriter) Write(msg *Message) error {
	_, err := fmt.Fprintln(os.Stdout, msg)
	return err
}

func NewStderrWriter() Writer {
	return stdWriter{}
}

type Log struct {
	Facility Facility
	Version  Version
	AppName  option.Option[AppName]
	HostName option.Option[HostName]
	Proccess option.Option[ProcessID]
	Metadata []Metadata
	Writer   Writer
}

func NewDefaultLogger(app option.Option[AppName], host option.Option[HostName], proc option.Option[ProcessID]) Logger {
	return &Log{
		Facility: FacilityUserLevelMessages,
		Version:  1,
		AppName:  app,
		HostName: host,
		Proccess: proc,
		Metadata: []Metadata{},
		Writer:   stdWriter{},
	}
}

func (log *Log) write(severity Severity, msg []any) error {
	return log.Writer.Write(NewMessage(
		NewHeader(
			NewPriority(log.Facility, severity),
			log.Version,
			option.Some(TimestampNow()),
			log.HostName,
			log.AppName,
			log.Proccess,
			option.None[MessageID](),
		),
		log.Metadata,
		msg...,
	))
}

func (log *Log) Emergency(msg ...any) error {
	return log.write(SeverityEmergency, msg)
}

func (log *Log) Alert(msg ...any) error {
	return log.write(SeverityAlert, msg)
}

func (log *Log) Critical(msg ...any) error {
	return log.write(SeverityCritical, msg)
}

func (log *Log) Error(msg ...any) error {
	return log.write(SeverityError, msg)
}

func (log *Log) Warning(msg ...any) error {
	return log.write(SeverityWarning, msg)
}

func (log *Log) Notice(msg ...any) error {
	return log.write(SeverityNotice, msg)
}

func (log *Log) Informational(msg ...any) error {
	return log.write(SeverityInformational, msg)
}

func (log *Log) Debug(msg ...any) error {
	return log.write(SeverityDebug, msg)
}
