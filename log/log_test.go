package log

import (
	"github.com/a-skua/busybox/option"
	"testing"
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

func TestDefaultLogger(t *testing.T) {
	t.Skip("FIXME")
	log := NewDefaultLogger(option.None[AppName](), option.None[HostName](), option.None[ProcessID]())
	log.Emergency("foo", "bar", "baz")
	log.Alert("foo", "bar", "baz")
	log.Critical("foo", "bar", "baz")
	log.Error("foo", "bar", "baz")
	log.Warning("foo", "bar", "baz")
	log.Notice("foo", "bar", "baz")
	log.Informational("foo", "bar", "baz")
	log.Debug("foo", "bar", "baz")
}
