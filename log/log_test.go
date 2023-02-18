package log

import (
	"github.com/a-skua/busybox-go/option"
	"testing"
	"time"
)

func TestConst_Facility(t *testing.T) {
	type test struct {
		name   string
		define Facility
		want   Facility
	}

	do := func(tt *test) {
		t.Run(tt.name, func(t *testing.T) {
			if tt.want != tt.define {
				t.Fatalf("want=%v, got=%v.", tt.want, tt.define)
			}
		})
	}

	tests := []*test{
		{
			name:   "kernel messages",
			define: FacilityKernelMessages,
			want:   0,
		},
		{
			name:   "user-level messages",
			define: FacilityUserLevelMessages,
			want:   1,
		},
		{
			name:   "mail system",
			define: FacilityMailSystem,
			want:   2,
		},
		{
			name:   "system daemons",
			define: FacilitySystemDaemons,
			want:   3,
		},
		{
			name:   "security/authorization messages",
			define: FacilitySecurityOrAuthorizationMessages0,
			want:   4,
		},
		{
			name:   "messages generated internally by syslogd",
			define: FacilityMessagesGeneratedInternallyBySyslogd,
			want:   5,
		},
		{
			name:   "line printer subsystem",
			define: FacilityLinePrinterSubsystem,
			want:   6,
		},
		{
			name:   "network news subsystem",
			define: FacilityNetworkNewsSubsystem,
			want:   7,
		},
		{
			name:   "UUCP subsystem",
			define: FacilityUUCPSubsystem,
			want:   8,
		},
		{
			name:   "clock daemon",
			define: FacilityClockDaemon0,
			want:   9,
		},
		{
			name:   "security/authorization messages",
			define: FacilitySecurityOrAuthorizationMessages1,
			want:   10,
		},
		{
			name:   "FTP daemon",
			define: FacilityFTPDaemon,
			want:   11,
		},
		{
			name:   "NTP subsystem",
			define: FacilityNTPSubsystem,
			want:   12,
		},
		{
			name:   "log audit",
			define: FacilityLogAudit,
			want:   13,
		},
		{
			name:   "log alert",
			define: FacilityLogAlert,
			want:   14,
		},
		{
			name:   "clock daemon",
			define: FacilityClockDaemon1,
			want:   15,
		},
		{
			name:   "local use 0",
			define: FacilityLocalUse0,
			want:   16,
		},
		{
			name:   "local use 1",
			define: FacilityLocalUse1,
			want:   17,
		},
		{
			name:   "local use 2",
			define: FacilityLocalUse2,
			want:   18,
		},
		{
			name:   "local use 3",
			define: FacilityLocalUse3,
			want:   19,
		},
		{
			name:   "local use 4",
			define: FacilityLocalUse4,
			want:   20,
		},
		{
			name:   "local use 5",
			define: FacilityLocalUse5,
			want:   21,
		},
		{
			name:   "local use 6",
			define: FacilityLocalUse6,
			want:   22,
		},
		{
			name:   "local use 7",
			define: FacilityLocalUse7,
			want:   23,
		},
	}

	for _, tt := range tests {
		do(tt)
	}
}

func TestConst_Severity(t *testing.T) {
	type test struct {
		name   string
		define Severity
		want   Severity
	}

	do := func(tt *test) {
		t.Run(tt.name, func(t *testing.T) {
			if tt.want != tt.define {
				t.Fatalf("want=%v, got=%v.", tt.want, tt.define)
			}
		})
	}

	tests := []*test{
		{
			name:   "Emergency",
			define: SeverityEmergency,
			want:   0,
		},
		{
			name:   "Alert",
			define: SeverityAlert,
			want:   1,
		},
		{
			name:   "Critical",
			define: SeverityCritical,
			want:   2,
		},
		{
			name:   "Error",
			define: SeverityError,
			want:   3,
		},
		{
			name:   "Warning",
			define: SeverityWarning,
			want:   4,
		},
		{
			name:   "Notice",
			define: SeverityNotice,
			want:   5,
		},
		{
			name:   "Informational",
			define: SeverityInformational,
			want:   6,
		},
		{
			name:   "Debug",
			define: SeverityDebug,
			want:   7,
		},
	}

	for _, tt := range tests {
		do(tt)
	}
}

func TestNewPriority(t *testing.T) {
	type args struct {
		facility Facility
		severity Severity
	}

	type test struct {
		name string
		args
		want Priority
	}

	do := func(tt *test) {
		t.Run(tt.name, func(t *testing.T) {
			got := NewPriority(tt.facility, tt.severity)
			if tt.want != got {
				t.Fatalf("want=%v, got=%v.", tt.want, got)
			}
		})
	}

	tests := []*test{
		{
			name: "true",
			args: args{
				facility: FacilityUserLevelMessages,
				severity: SeverityError,
			},
			want: Priority{
				Facility: FacilityUserLevelMessages,
				Severity: SeverityError,
			},
		},
	}

	for _, tt := range tests {
		do(tt)
	}
}

func TestPriority_Num(t *testing.T) {
	type test struct {
		name     string
		priority Priority
		want     uint8
	}

	do := func(tt *test) {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.priority.Num()
			if tt.want != got {
				t.Fatalf("want=%v, got=%v.", tt.want, got)
			}
		})
	}

	tests := []*test{
		{
			name:     "local use 4(Facility=20) / Notice(Severity=5)",
			priority: NewPriority(FacilityLocalUse4, SeverityNotice),
			want:     165,
		},
	}

	for _, tt := range tests {
		do(tt)
	}
}

func TestPriority_String(t *testing.T) {
	type test struct {
		name     string
		priority Priority
		want     string
	}

	do := func(tt *test) {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.priority.String()
			if tt.want != got {
				t.Fatalf("want=%v, got=%v.", tt.want, got)
			}
		})
	}

	tests := []*test{
		{
			name:     "local use 4(Facility=20) / Notice(Severity=5)",
			priority: NewPriority(FacilityLocalUse4, SeverityNotice),
			want:     "<165>",
		},
	}

	for _, tt := range tests {
		do(tt)
	}
}

func TestHeader_String(t *testing.T) {
	type test struct {
		name   string
		header Header
		want   string
	}

	do := func(tt *test) {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.header.String()
			if tt.want != got {
				t.Fatalf("want=%v, got=%v.", tt.want, got)
			}
		})
	}

	tests := []*test{
		{
			name: "true: minimum",
			header: NewHeader(
				NewPriority(FacilityLocalUse4, SeverityNotice),
				1,
				option.None[Timestamp](),
				option.None[HostName](),
				option.None[AppName](),
				option.None[ProcessID](),
				option.None[MessageID](),
			),
			want: "<165>1 - - - - -",
		},
		{
			name: "true: with timestamp",
			header: NewHeader(
				NewPriority(FacilityLocalUse4, SeverityNotice),
				1,
				option.Some(Timestamp(time.Date(2023, 2, 15, 12, 31, 56, 0, time.UTC))),
				option.None[HostName](),
				option.None[AppName](),
				option.None[ProcessID](),
				option.None[MessageID](),
			),
			want: "<165>1 2023-02-15T12:31:56Z - - - -",
		},
		{
			name: "true: with hostname",
			header: NewHeader(
				NewPriority(FacilityLocalUse4, SeverityNotice),
				1,
				option.None[Timestamp](),
				option.Some(HostName("localhost")),
				option.None[AppName](),
				option.None[ProcessID](),
				option.None[MessageID](),
			),
			want: "<165>1 - localhost - - -",
		},
		{
			name: "true: with appname",
			header: NewHeader(
				NewPriority(FacilityLocalUse4, SeverityNotice),
				1,
				option.None[Timestamp](),
				option.None[HostName](),
				option.Some(AppName("myapp")),
				option.None[ProcessID](),
				option.None[MessageID](),
			),
			want: "<165>1 - - myapp - -",
		},
		{
			name: "true: with processid",
			header: NewHeader(
				NewPriority(FacilityLocalUse4, SeverityNotice),
				1,
				option.None[Timestamp](),
				option.None[HostName](),
				option.None[AppName](),
				option.Some(ProcessID("my-process-id")),
				option.None[MessageID](),
			),
			want: "<165>1 - - - my-process-id -",
		},
		{
			name: "true: with messageid",
			header: NewHeader(
				NewPriority(FacilityLocalUse4, SeverityNotice),
				1,
				option.None[Timestamp](),
				option.None[HostName](),
				option.None[AppName](),
				option.None[ProcessID](),
				option.Some(MessageID("ID47")),
			),
			want: "<165>1 - - - - ID47",
		},
	}

	for _, tt := range tests {
		do(tt)
	}
}

func TestMetadataParam_String(t *testing.T) {
	type test struct {
		name  string
		param MetadataParam
		want  string
	}

	do := func(tt *test) {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.param.String()
			if tt.want != got {
				t.Fatalf("want=%v, got=%v.", tt.want, got)
			}
		})
	}

	tests := []*test{
		{
			name:  "normal",
			param: NewMetadataParam("id", "myid"),
			want:  "id=\"myid\"",
		},
		{
			name:  "escape \"",
			param: NewMetadataParam("id", "\""),
			want:  "id=\"\\\"\"",
		},
		{
			name:  "escape \\",
			param: NewMetadataParam("id", "\\"),
			want:  "id=\"\\\\\"",
		},
		{
			name:  "escape ]",
			param: NewMetadataParam("id", "]"),
			want:  "id=\"\\]\"",
		},
	}

	for _, tt := range tests {
		do(tt)
	}
}

func TestMetadata_String(t *testing.T) {
	type test struct {
		name     string
		metadata Metadata
		want     string
	}

	do := func(tt *test) {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.metadata.String()
			if tt.want != got {
				t.Fatalf("want=%v, got=%v.", tt.want, got)
			}
		})
	}

	tests := []*test{
		{
			name: "empty param",
			metadata: Metadata{
				ID:     "exampleSDID@32473",
				Params: []MetadataParam{},
			},
			want: "[exampleSDID@32473]",
		},
		{
			name: "with param",
			metadata: Metadata{
				ID: "exampleSDID@32473",
				Params: []MetadataParam{
					NewMetadataParam("eventID", "1011"),
				},
			},
			want: "[exampleSDID@32473 eventID=\"1011\"]",
		},
		{
			name: "with params",
			metadata: Metadata{
				ID: "exampleSDID@32473",
				Params: []MetadataParam{
					NewMetadataParam("eventID", "1011"),
					NewMetadataParam("eventSource", "Application"),
				},
			},
			want: "[exampleSDID@32473 eventID=\"1011\" eventSource=\"Application\"]",
		},
	}

	for _, tt := range tests {
		do(tt)
	}
}

func TestMessage_String(t *testing.T) {
	type test struct {
		name    string
		message *Message
		want    string
	}

	do := func(tt *test) {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.message.String()
			if tt.want != got {
				t.Fatalf("want=%v, got=%v.", tt.want, got)
			}
		})
	}

	tests := []*test{
		{
			name: "empty message",
			message: NewMessage(
				NewHeader(
					NewPriority(FacilityLocalUse4, SeverityNotice),
					1,
					option.None[Timestamp](),
					option.None[HostName](),
					option.None[AppName](),
					option.None[ProcessID](),
					option.None[MessageID](),
				),
				[]Metadata{},
			),
			want: "<165>1 - - - - - -",
		},
		{
			name: "with message",
			message: NewMessage(
				NewHeader(
					NewPriority(FacilityLocalUse4, SeverityNotice),
					1,
					option.None[Timestamp](),
					option.None[HostName](),
					option.None[AppName](),
					option.None[ProcessID](),
					option.None[MessageID](),
				),
				[]Metadata{},
				"foo",
				"bar",
				"baz",
			),
			want: "<165>1 - - - - - - foo bar baz",
		},
		{
			name: "with info",
			message: NewMessage(
				NewHeader(
					NewPriority(FacilityLocalUse4, SeverityNotice),
					1,
					option.Some(Timestamp(time.Date(2023, 02, 16, 12, 34, 56, 0, time.UTC))),
					option.Some(HostName("localhost")),
					option.Some(AppName("busybox")),
					option.None[ProcessID](),
					option.None[MessageID](),
				),
				[]Metadata{},
				"hello, syslog!",
			),
			want: "<165>1 2023-02-16T12:34:56Z localhost busybox - - - hello, syslog!",
		},
		{
			name: "with metadata",
			message: NewMessage(
				NewHeader(
					NewPriority(FacilityLocalUse4, SeverityNotice),
					1,
					option.Some(Timestamp(time.Date(2023, 02, 16, 12, 34, 56, 0, time.UTC))),
					option.Some(HostName("localhost")),
					option.Some(AppName("busybox")),
					option.None[ProcessID](),
					option.None[MessageID](),
				),
				[]Metadata{
					NewMetadata("exampleSDID@0"),
				},
				"hello, syslog!",
			),
			want: "<165>1 2023-02-16T12:34:56Z localhost busybox - - [exampleSDID@0] hello, syslog!",
		},
		{
			name: "with metadata",
			message: NewMessage(
				NewHeader(
					NewPriority(FacilityLocalUse4, SeverityNotice),
					1,
					option.Some(Timestamp(time.Date(2023, 02, 16, 12, 34, 56, 0, time.UTC))),
					option.Some(HostName("localhost")),
					option.Some(AppName("busybox")),
					option.None[ProcessID](),
					option.None[MessageID](),
				),
				[]Metadata{
					NewMetadata("exampleSDID@0"),
					NewMetadata("exampleSDID@1", NewMetadataParam("eventID", "1011"), NewMetadataParam("eventSource", "Application")),
				},
				"hello, syslog!",
			),
			want: "<165>1 2023-02-16T12:34:56Z localhost busybox - - [exampleSDID@0][exampleSDID@1 eventID=\"1011\" eventSource=\"Application\"] hello, syslog!",
		},
	}

	for _, tt := range tests {
		do(tt)
	}
}

func BenchmarkMessage_String(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = NewMessage(
			NewHeader(
				NewPriority(FacilityUserLevelMessages, SeverityInformational),
				1,
				option.Some(TimestampNow()),
				option.Some(HostName("localhost")),
				option.Some(AppName("benchmark")),
				option.None[ProcessID](),
				option.None[MessageID](),
			),
			[]Metadata{
				NewMetadata("benchmark@1"),
				NewMetadata("benchmark@2", NewMetadataParam("foo", "FOO"), NewMetadataParam("bar", "[BAR]")),
			},
			"Syslog!",
			"Benchmark!",
		).String()
	}
}

func Example_stderrWriter_Write() {
	w := NewStderrWriter()
	w.Write(NewMessage(
		NewHeader(
			NewPriority(FacilityLocalUse4, SeverityNotice),
			1,
			option.None[Timestamp](),
			option.None[HostName](),
			option.None[AppName](),
			option.None[ProcessID](),
			option.None[MessageID](),
		),
		[]Metadata{},
	))
	w.Write(NewMessage(
		NewHeader(
			NewPriority(FacilityLocalUse4, SeverityNotice),
			1,
			option.None[Timestamp](),
			option.None[HostName](),
			option.None[AppName](),
			option.None[ProcessID](),
			option.None[MessageID](),
		),
		[]Metadata{},
		"foo",
		"bar",
		"baz",
	))
	w.Write(NewMessage(
		NewHeader(
			NewPriority(FacilityLocalUse4, SeverityNotice),
			1,
			option.Some(Timestamp(time.Date(2023, 02, 16, 12, 34, 56, 0, time.UTC))),
			option.Some(HostName("localhost")),
			option.Some(AppName("busybox")),
			option.None[ProcessID](),
			option.None[MessageID](),
		),
		[]Metadata{},
		"hello, syslog!",
	))
	w.Write(NewMessage(
		NewHeader(
			NewPriority(FacilityLocalUse4, SeverityNotice),
			1,
			option.Some(Timestamp(time.Date(2023, 02, 16, 12, 34, 56, 0, time.UTC))),
			option.Some(HostName("localhost")),
			option.Some(AppName("busybox")),
			option.None[ProcessID](),
			option.None[MessageID](),
		),
		[]Metadata{
			NewMetadata("exampleSDID@0"),
		},
		"hello, syslog!",
	))
	w.Write(NewMessage(
		NewHeader(
			NewPriority(FacilityLocalUse4, SeverityNotice),
			1,
			option.Some(Timestamp(time.Date(2023, 02, 16, 12, 34, 56, 0, time.UTC))),
			option.Some(HostName("localhost")),
			option.Some(AppName("busybox")),
			option.None[ProcessID](),
			option.None[MessageID](),
		),
		[]Metadata{
			NewMetadata("exampleSDID@0"),
			NewMetadata("exampleSDID@1", NewMetadataParam("eventID", "1011"), NewMetadataParam("eventSource", "Application")),
		},
		"hello, syslog!",
	))
	// Output:
	// <165>1 - - - - - -
	// <165>1 - - - - - - foo bar baz
	// <165>1 2023-02-16T12:34:56Z localhost busybox - - - hello, syslog!
	// <165>1 2023-02-16T12:34:56Z localhost busybox - - [exampleSDID@0] hello, syslog!
	// <165>1 2023-02-16T12:34:56Z localhost busybox - - [exampleSDID@0][exampleSDID@1 eventID="1011" eventSource="Application"] hello, syslog!
}

func ExampleNewDefaultLogger() {
	log := NewDefaultLogger(option.None[AppName](), option.None[HostName](), option.None[ProcessID]())
	log.Emergency("hello, syslog!")
	log.Alert("hello, syslog!")
	log.Critical("hello, syslog!")
	log.Error("hello, syslog!")
	log.Warning("hello, syslog!")
	log.Notice("hello, syslog!")
	log.Informational("hello, syslog!")
	log.Debug("hello, syslog!")
}
