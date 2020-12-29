package nakivo

import (
	"context"
	"net/http"
)

const (
	JobAction = "JobSummaryManagement"
)

type JobService service

type GroupInfo struct {
	Children []groupInfo `json:"children"`
}

type groupInfo struct {
	// Job group id
	Id int `json:"id"`

	// Job group vid
	Vid string `json:"vid"`

	// Job group display name
	Name string `json:"name"`

	// Job status
	Status string `json:"status"`

	// The number of jobs inside the group, grouped by hypervisor
	JobCount jobCount `json:"jobCount"`

	// The number of enabled jobs
	JobCountEnabled int `json:"jobCountEnabled"`

	// The number of jobs that don't violate the license
	JobCountLicensed int `json:"jobCountLicensed"`

	// The number of backups grouped by hypervisor
	HVTypeBackupCount hvTypeBackupCount `json:"hvTypeBackupCount,omitempty"`

	// AWS-specific: backup count which savepoints have a root volume.
	HVTypeBackupHasRootDiskCount hvTypeBackupHasRootDiskCount `json:"hvTypeBackupHasRootDiskCount,omitempty"`

	// The number of VMs processed by the jobs inside a group
	VMCount int `json:"vmCount"`

	// The number of disks processed by the jobs inside a group
	DiskCount int `json:"diskCount"`

	// Total size of the sources processed by the jobs inside a group (in bytes)
	SourcesSize int `json:"sourcesSize"`

	// Checks if the group is enabled
	IsEnabled bool `json:"isEnabled"`

	// Checks if the group is removed
	IsRemoved bool `json:"isRemoved"`

	// The number of currently running jobs
	CrJobRunning int `json:"crJobRunning"`

	// The number of currently processed vms
	CrVMRunning int `json:"crVmRunning"`

	// HasLastRun indicates if the job had the last run
	HasLastRun bool `json:"hasLastRun"`

	// Number of successful last runs
	LrJobOk int `json:"lrJobOk"`

	// The number of failed jobs
	LrJobFailed int `json:"lrJobFailed"`

	// The number of stopped jobs
	LrJobStopped int `json:"lrJobStopped"`

	// The IDs of the child jobs
	ChildJobIds []int `json:"childJobIds"`

	// The IDs of the direct children of a group
	ImmediateChildJobIds []int `json:"immediateChildJobIds"`

	// Info about the transporters involved
	Transporters []transporter `json:"transporters"`

	// Info about the storage involved
	Storages []storage `json:"storages"`
}

type jobCount struct {
	Replication     int `json:"REPLICATION"`
	Backup          int `json:"BACKUP"`
	RecoveryVMs     int `json:"RECOVERY_VMS"`
	RecoveryFiles   int `json:"RECOVERY_FILES"`
	RecoveryBackups int `json:"RECOVERY_BACKUPS"`
	BackupCopy      int `json:"BACKUP_COPY"`
	FlashBoot       int `json:"FLASH_BOOT"`
}

// TODO: no official documentation from nakivo
type hvTypeBackupCount struct {
}

// TODO: no official documentation from nakivo
type hvTypeBackupHasRootDiskCount struct {
}

type transporter struct {
	// IsAuto indicates if a transporter was assigned automatically
	IsAuto bool `json:"isAuto"`

	// UsedAsSource indicates if a transporter used as a source
	UsedAsSource bool `json:"usedAsSource"`

	// UsedAsTarget indicates if a transporter used as a target
	UsedAsTarget bool `json:"usedAsTarget"`

	// Maximum number of jobs supported by a transporter
	MaxLoadFactor int `json:"maxLoadFactor"`

	// Current number of jobs using a transporter
	CurrentTotalLoad int `json:"currentTotalLoad"`

	// Vid of a transporter
	Vid string `json:"vid"`

	// Display name of a transporter
	Name string `json:"name"`

	// Current state of a transporter
	State string `json:"state"`
}

type storage struct {
	// Storage vid
	Vid string `json:"vid"`

	// Storage display name
	Name string `json:"name"`

	// Full size of the storage
	Size int64 `json:"size"`

	// Free space on the storage
	Free int64 `json:"free"`

	// Used space on the storage
	Used int64 `json:"used"`

	// Storage status
	State string `json:"state"`

	// Online indicates if the storage is online
	Online bool `json:"online"`

	// AWS-specific: true if the ebs or ebs snapshot storage is used
	InfiniteSize bool `json:"infiniteSize"`

	// The type of storage
	Type string `json:"type"`
}

// List lists all jobs in a group
//
// TODO: Currently limited to list ALL jobs in a group. To list a limited set of group, the first
// element of the data element in the request should relflect a group id.
func (s *JobService) List(ctx context.Context, clientTimeOffset int, collectAllChildJobs bool) (*Response, *http.Response, error) {
	request := Request{
		Action: JobAction,
		Method: "getGroupInfo",
		Data:   []interface{}{[]interface{}{nil}, clientTimeOffset, collectAllChildJobs},
		Type:   "rpc",
		Tid:    1,
	}

	req, err := s.client.NewRequest(&request)
	if err != nil {
		return nil, nil, err
	}
	r := Response{Data: &GroupInfo{}}
	resp, err := s.client.Do(ctx, req, &r)
	if err != nil {
		return nil, resp, err
	}
	return &r, resp, nil
}
