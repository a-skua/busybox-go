// RFC5424 The Syslog Protocol: https://www.rfc-editor.org/rfc/rfc5424
package log

import (
	"encoding/json"
	"github.com/a-skua/busybox/option"
	"io"
	"os"
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

func (pri Priority) Num() uint8 {
	return uint8(pri.Facility)*8 + uint8(pri.Severity)
}

type Version uint8

type Timestamp time.Time

func TimeNow() Timestamp {
	return Timestamp(time.Now())
}

func (t Timestamp) String() string {
	return time.Time(t).String()
}

func (t Timestamp) MarshalJSON() ([]byte, error) {
	return time.Time(t).MarshalJSON()
}

type HostName string

type AppName string

type ProcessID string

type MessageID string

type Header struct {
	Priority  Priority
	Version   Version
	Timestamp option.Option[Timestamp]
	HostName  option.Option[HostName]
	AppName   option.Option[AppName]
	ProcessID option.Option[ProcessID]
	MessageID option.Option[MessageID]
}

func NewHeader(pri Priority, ver Version, time option.Option[Timestamp], host option.Option[HostName], app option.Option[AppName], proc option.Option[ProcessID], msg option.Option[MessageID]) Header {
	return Header{
		Priority:  pri,
		Version:   ver,
		Timestamp: time,
		HostName:  host,
		AppName:   app,
		ProcessID: proc,
		MessageID: msg,
	}
}

type MetadataID string

type MetadataName string

type MetadataValue string

type MetadataParam struct {
	Name  MetadataName
	Value MetadataValue
}

type Metadata struct {
	ID     MetadataID
	Params []MetadataParam
}

type Message struct {
	Header   Header
	Metadata option.Option[Metadata]
	Message  []any
}

func NewMessage(head Header, meta option.Option[Metadata], msg []any) *Message {
	return &Message{
		Header:   head,
		Metadata: meta,
		Message:  msg,
	}
}

type Formatter interface {
	Format(*Message) ([]byte, error)
}

type JsonFormatter struct{}

func (fmt JsonFormatter) Format(msg *Message) ([]byte, error) {
	return json.Marshal(msg)
}

type Log struct {
	Facility Facility
	Version  Version
	AppName  option.Option[AppName]
	HostName option.Option[HostName]
	Proccess option.Option[ProcessID]
	Metadata option.Option[Metadata]
	Fmt      Formatter
	Output   io.Writer
}

func NewDefaultLogger(app option.Option[AppName], host option.Option[HostName], proc option.Option[ProcessID]) Logger {
	return &Log{
		Facility: FacilityUserLevelMessages,
		Version:  1,
		AppName:  app,
		HostName: host,
		Proccess: proc,
		Fmt:      JsonFormatter{},
		Output:   os.Stderr,
	}
}

func (log *Log) write(severity Severity, msg []any) error {
	bin, err := log.Fmt.Format(NewMessage(
		NewHeader(
			NewPriority(log.Facility, severity),
			log.Version,
			option.Some(TimeNow()),
			log.HostName,
			log.AppName,
			log.Proccess,
			option.None[MessageID](),
		),
		log.Metadata,
		msg,
	))
	if err != nil {
		return err
	}

	// FIXME
	log.Output.Write(bin)
	log.Output.Write([]byte("\n"))

	return nil
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
