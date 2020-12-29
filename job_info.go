package nakivo

import (
	"context"
	"net/http"
)

type JobInfo struct {
	Children []jobInfo `json:"children"`
}

type jobInfo struct {
	// Name of the job
	Name string `json:"name"`

	// Id of the job
	Id int `json:"id"`

	// Vid of the job
	Vid string `json:"vid"`

	// Platform type
	HvType string `json:"hvType"`

	// Number of machines by platform
	HVTypeBackupCount hvTypeBackupCount `json:"hvTypeBackupCount,omitempty"`

	// Special case for a replication job: true if from Backup, false if from VM
	FromBackup bool `json:"fromBackup"`

	// Site recovery specific: the type of recovery running for the last job run.
	// Possible values: TEST, RUN
	LrSiteRecoveryRunType string `json:"lrSiteRecoveryRunType"`

	// Site recovery specific: the type of recovery time objective for the last job run.
	// Possible values: MINUTE, HOUR
	LrRecoveryTimeObjectiveType string `json:"lrRecoveryTimeObjectiveType"`

	// Site recovery specific: the recovery time objective for the last job run
	LrRecoveryTimeObjective int `json:"lrRecoveryTimeObjective"`

	// Site recovery specific: the failover type for the last job run.
	// Possible values: PLANNED_FAILOVER, EMERGENCY_FAILOVER
	LrFailoverType string `json:"lrFailoverType"`

	// Site recovery specific: action failed for the last job run
	LrActionFailed int `json:"lrActionFailed"`

	// Site recovery specific: action skipped for the last job run
	LrActionSkipped int `json:"lrActionSkipped"`

	// Site recovery specific: action succeeded for the last job run
	LrActionSucceed int `json:"lrActionSucceed"`

	// Site recovery specific: action stopped for the last job run
	LrActionStopped int `json:"lrActionStopped"`

	// Path to the post-script
	PostScriptPath string `json:"postScriptPath"`

	// true if powering off the source VMs is needed
	PowerSourceVmsOff bool `json:"powerSourceVmsOff"`

	// Truncation mode of Microsoft SQL Server logging.
	// Possible values: NONE, ALWAYS, JOB_SUCCESS
	SqlLogTruncationMode string `json:"sqlLogTruncationMode"`

	// Type of full backup job run settings.
	// Possible values: ALWAYS, EVERY, EVERY2ND, FIRST, SECOND, THIRD, FOURTH, LAST, DAY, EVERY_JOB_RUNS
	FullBackupRunSettingsType string `json:"fullBackupRunSettingsType"`

	// Full backup mode.
	// Possible values: SYNTHETIC, ACTIVE
	FullBackupMode string `json:"fullBackupMode"`

	// Type of recovery time objective.
	// Possible values: MINUTE, HOUR
	RecoveryTimeObjectiveType string `json:"recoveryTimeObjectiveType"`

	// Recovery time objective
	RecoveryTimeObjective int `json:"recoveryTimeObjective"`

	// Specific for site recovery job: the type of site recovery running for the current job
	// run.
	// Possible values: TEST, RUN
	CrSiteRecoveryRunType string `json:"crSiteRecoveryRunType"`

	// Specific for site recovery job: the type of recovery time objective for the current job
	// run.
	// Possible values: MINUTE, HOUR
	CrRecoveryTimeObjectiveType string `json:"crRecoveryTimeObjectiveType"`

	// Specific for site recovery job: the recovery time objective for the current job run
	CrRecoveryTimeObjective int `json:"crRecoveryTimeObjective"`

	// Specific for site recovery job: the failover type for the current job run.
	// Possible values: PLANNED_FAILOVER, EMERGENCY_FAILOVER
	CrFailoverType string `json:"crFailoverType"`

	// List of actions
	//
	// TODO: no official documentation from nakivo
	Actions []interface{} `json:"action,omitempty"`

	// List of actions for the current job run
	//
	CrActionExecutions []interface{} `json:"crActionExecutions,omitempty"`

	// List of action for the last job run
	LrActionExecutions []interface{} `json:"lrActionExecutions,omitempty"`

	// AWS-specific: backup count which savepoints have a Root volume
	HVTypeBackupHasRootDiskCount hvTypeBackupHasRootDiskCount `json:"hvTypeBackupHasRootDiskCount,omitempty"`

	// Job status
	Status string `json:"status"`

	// Job type.
	// Possible values: REPLICATION, BACKUP, RECOVERY_VMS, RECOVERY_FILES, RECOVERY_OBJECTS,
	// BACKUP_COPY, FLASH_BOOT, REPLICA_FAILOVER, SITE_RECOVERY
	JobType string `json:"jobType"`

	// Date the job got added
	Added string `json:"added"`

	// Date the job got updated
	Updated string `json:"updated"`

	// The number of machines
	VmCount int `json:"vmCount"`

	// The number of disks
	DiskCount int `json:"diskCount"`

	// The total size of source machines
	SourcesSize int `json:"sourcesSize"`

	// Checks if the job is enabled
	IsEnabled bool `json:"isEnabled"`

	// Checks if the job does not violate license
	IsLicensed bool `json:"isLicensed"`

	// Checks if the job was edited
	IsEdited bool `json:"isEdited"`

	// Checks if the job was locked
	IsLocked bool `json:"isLocked"`

	// Checks if the job was removed
	IsRemoved bool `json:"isRemoved"`

	// The average job duration (in ms)
	AverageDurationMs int `json:"averageDurationMs"`

	// The number of job runs to calculate average job duration
	AverageDurationSampleCount int `json:"averageDurationSampleCount"`

	// Current state.
	// Possible values: WAITING_DEMAND, WAITING_SCHEDULE, RUNNING, OK, FAILED, STOPPED
	CrState string `json:"crState"`

	// Timestamp of the current job run
	CrDate string `json:"crDate"`

	// Relative timestamp (time passed from the job start run) of the current job run
	CrDateRelative int `json:"crDateRelative"`

	// The number of machines queued for processing for the current job run
	CrVmPlanned int `json:"crVmPlanned"`

	// The number of successfully processed machines during the current job run
	CrVmOk int `json:"crVmOk"`

	// The number of failed machines during the current job run
	CrVmFailed int `json:"crVmFailed"`

	// The number of machines stopped during the current job run
	CrVmStopped int `json:"crVmStopped"`

	// Current job run progress
	CrProgress int `json:"crProgress"`

	// true if the current job run was started manually (not on schedule)
	CrAdhoc bool `json:"crAdhoc"`

	// Checks if the job had a last run
	HasLastRun bool `json:"hasLastRun"`

	// The state of the last job run.
	// ossible values: WAITING_DEMAND, WAITING_SCHEDULE, RUNNING, OK, FAILED, STOPPED
	LrState string `json:"lrState"`

	// The date and time of the last job run start
	LrDate string `json:"lrDate"`

	// The date and time of the last job run end
	LrFinishDate string `json:"lrFinishDate"`

	// The speed of the last job run
	LrSpeed int64 `json:"lrSpeed"`

	// The duration of the last job run
	LrDurationMs int64 `json:"lrDurationMs"`

	// Data transferred during the last job run
	LrDataKb int64 `json:"lrDataKb"`

	// Number of successfully processed machines during the last job run
	LrVmOk int `json:"lrVmOk"`

	// Number of machines failed during the last job run
	LrVmFailed int `json:"lrVmFailed"`

	// Number of machines stopped during the last job run
	LrVmStopped int `json:"lrVmStopped"`

	// true if the last job run was started manually (not on schedule)
	LrAdhoc bool `json:"lrAdhoc"`

	// The compression ratio during the last job run
	LrCompressionRatio int `json:"lrCompressionRatio"`

	// A method used for forever-incremental backup.
	// Possible values: NONE, HYPERVISOR, DOUBLE_CHECK, PROPRIETARY
	DifferentialTrackingMode string `json:"differentialTrackingMode"`

	// The mode of execution of pre-job scripts.
	// Possible values: NEVER, ALWAYS
	PreScriptExecutionMode string `json:"preScriptExecutionmode"`

	// Job behavior: either to wait for the script to finish or proceed.
	// Possible values: NONE, WAIT, PROCEED
	PreScriptBehavior string `json:"preScriptBehavior"`

	// The job behavior on pre-job script failure.
	// Possible values: NONE, FAIL, SKIP
	PreScriptErrorMode string `json:"preScriptErrorMode"`

	// The path to the pre-job script
	PreScriptPath string `json:"preScriptPath"`

	// The mode of execution of post-job scripts.
	// Possible values: NEVER, ALWAYS
	PostScriptExecutionMode string `json:"postScriptExecutionMode"`

	// Job behavior: either to wait for the script to finish or proceed.
	// Possible values: NONE, WAIT, PROCEED
	PostScriptBehavior string `json:"postScriptBehavior"`

	// The job behavior on post-job script failure.
	// Possible values: NONE, FAIL, SKIP
	PostScriptErrorMode string `json:"postScriptErrorMode"`

	// For replication: thin disk or respect source.
	// Possible values: AUTO, FORCE_THIN
	ThinDiskMode string `json:"thinDiskMode"`

	// For AWS jobs only: a type of EBS volume.
	// Possible values: AUTO, FORCE_MAGNETIC
	EbsVolumeMode string `json:"ebsVolumeMode"`

	// For AWS jobs only: a type of a temporary volume
	TemporaryVolumeType string `json:"temporaryVolumeType"`

	// Network acceleration mode.
	// Possible values: NONE, AUTO, FAST, MEDIUM, BEST
	NetworkAccelerationMode string `json:"networkAccelerationMode"`

	// Encryption mode.
	// Possible values: NONE, NORMAL
	EncryptionMode string `json:"encryptionMode"`

	// Application-aware mode.
	// Possible values: NONE, VSS_IGNORE_ERRORS, VSS_FAIL_ON_ERRORS
	ApplicationAwareMode string `json:"applicationAwareMode"`

	// Retention policy options.
	RetentionPolicy retentionPolicy `json:"retentionPolicy"`

	// Recovery only. Defines if the recovered machines must be powered on after the job is
	// completed
	PowerVmsOn bool `json:"powerVmsOn"`

	// Recovery only. Defines if a new MAC-address should be generated for the recovered machine
	GenerateMac bool `json:"generateMac"`

	// Recovery type.
	// Possible values: SYNTHETIC, PRODUCTION
	RecoveryType string `json:"recoveryType"`

	// The mode of data transfer.
	// Possible values: AUTO, SAN, LAN, HOT_ADD
	TransporterMode string `json:"transporterMode"`

	// Mode of Microsoft Exchange log truncation.
	// Possible values: NONE, ALWAYS, JOB_SUCCESS
	ExchangelogTruncationMode string `json:"exchangeLogTruncationMode"`

	// The mode of screenshot verification.
	// Possible values: NEVER, ALWAYS
	ScreenshotVerificationMode string `json:"screenshotVerificationMode"`

	// Source objects
	Objects []object `json:"objects"`

	// Info about the transporters involved
	Transporters []transporter `json:"transporters"`

	// Info about the storage involved
	Storages []storage `json:"storages"`

	// Job schedules
	Schedules []schedule `json:"schedules"`

	// If the job is locked, lock reasons
	LockReasons []interface{} `json:"lockReasons,omitempty"`
}

type object struct {
	// VID of a job object
	Vid string `json:"vid"`

	// The state of the last job run.
	// Possible values: SCHEDULED, DEMAND, WAITING, RUNNING, STOPPED, FAILED, SUCCEEDED, SKIPPED
	LrState string `json:"lrState"`

	// The speed of the last job run
	LrSpeed int64 `json:"lrSpeed"`

	// The amount of uncompressed data during the last job run
	LrDataTransferredUncompressed int64 `json:"lrDataTransferredUncompressed"`

	// The duration of the last job run
	LrDuration int64 `json:"lrDuration"`

	// Current state.
	// Possible values: SCHEDULED, DEMAND, WAITING, RUNNING, STOPPED, FAILED, SUCCEEDED, SKIPPED
	CrState string `json:"lrState"`

	// The progress of the current job
	CrProgress int `json:"crProgress"`

	// Source object VID
	SourceVid string `json:"sourceVid"`

	// Source object display name
	SourceName string `json:"sourceName"`

	// Source object power state.
	// Possible values: ON, OFF, SUSPENDED, UNKNOWN
	SourcePowerState string `json:"sourcePowerState"`

	// Source object subtype
	SourceSubType string `json:"sourceSubType"`

	// Target object VID
	TargetVid string `json:"targetVid"`

	// Target object display name
	TargetName string `json:"targetName"`

	// Target object power state.
	// Possible values: ON, OFF, SUSPENDED, UNKNOWN
	TargetPowerState string `json:"targetPowerState"`

	// Target object subtype
	TargetSubType string `json:"targetSubType"`

	// The state of verification.
	// Possible values: SCHEDULED, DEMAND, WAITING, RUNNING, STOPPED, FAILED, SUCCEEDED, SKIPPED
	VerificationState string `json:"verificationState"`

	// The name of the screenshot used for verification
	ScreenshotName string `json:"screenshotName"`

	// The state of the flash boot.
	// Possible values: WAITING, STARTING, RUNNING_VM, FAILED, DISCARDING, DISCARDED
	FlashBootState string `json:"flashBootState"`

	// Bandwidth limit for the last job run
	LrBandwidthLimit int64 `json:"lrBandwidthLimit"`

	// Transporter modes for the last job run.
	// Possible values: AUTO, SAN, LAN, HOT_ADD
	LrTransporterModes []interface{} `json:"lrTransporterModes"`

	// Transporter modes for the current job run.
	// Possible values: AUTO, SAN, LAN, HOT_ADD
	CrTransporterModes []interface{} `json:"crTransporterModes"`

	// Screenshot path
	ScreenshotPath string `json:"screenshotPath"`

	// Job speed in bytes/s
	CrSpeed int64 `json:"crSpeed"`

	// Size of uncompressed transferred data in bytes
	CrDataTransferredUncompressed int64 `json:"crDataTransferredUncompressed"`

	// Data copy duration in milliseconds
	CrDuration int64 `json:"crDuration"`

	// Bandwidth limit in bits/s
	CrBandwidthLimit int64 `json:"crBandwidthLimit"`
}
type retentionPolicy struct {
	Mode           string `json:"retentionMode"`
	MaxCount       int    `json:"maxCount"`
	KeepDayCount   int    `json:"keepDayCount"`
	KeepWeekCount  int    `json:"keepWeekcount"`
	KeepMonthCount int    `json:"keepMonthCount"`
	KeepYearCount  int    `json:"keepYearCount"`
}

type schedule struct {
	// Enabled indicates ff a schedule is enabled
	Enabled bool `json:"enabled"`

	// The type of backup schedule
	// Possible values: DAILY, PERIODICALLY, NONE, MONTHLY_YEARLY, TRIGGER
	Type string `json:"type"`

	// Priority of the schedule
	Position int `json:"position"`

	// Start time. hh:mm:ss AM/PM
	StartTime string `json:"startTime"`

	// End time. null if not set
	EndTime string `json:"endTime"`

	// Timezone
	Timezone string `json:"timezone"`

	// A decimal representation of a bit mask for a day of a week. The lowest bit is Monday, the 7th is
	// Sunday. For example, if you need to run a job on weekdays, the bitmask will be 00011111 which is 31
	// in decimal
	// Possible values: 1-127
	On int `json:"on"`

	// If type is PERIODICALLY or TRIGGERED, defines the delay unit between jobs. For example, "Run job
	// every 30 minutes".
	// Possible values: DAY (for PERIODICALLY only), SECOND, MINUTE, HOUR
	EveryType string `json:"everyType"`

	// The number of delay units between job runs.
	// For example, "Run job every 30 minutes"
	Every int `json:"every"`

	// If type is MONTHLY_YEARLY , selects the number of a weekday in a month or a day number
	// Possible values: FIRST, SECOND, THIRD, FOURTH, LAST, DAY
	MonthlyEveryType string `json:"monthlyEveryType"`

	// If monthlyEveryType is DAY, selects the day number in a month
	DayOfMonth int `json:"dayOfMonth"`

	// If monthlyEveryType is either from LAST, selects the number of a weekday
	DayOfWeek int `json:"dayOfWeek"`

	// if type is MONTHLY_YEARLY, the number of the month when a job must be running
	Month int `json:"month"`

	// VID of the job that triggers the current one
	TriggerItem string `json:"triggerItem"`

	// Selects either to run the job immediately after the previous or within a delay. If DELAYED is
	// selected, the delay is defined by the everyType and every fields
	TriggerRunType string `json:"triggerRunType"`

	// Trigger job conditions
	// RUN_SUCCESS, RUN_FAILURE, RUN_STOP
	TriggerEvents []interface{} `json:"triggerEvents"`

	// Time and date of the next job run
	// YYYY-MM-DDTHH:MM:SS.SSSZ
	NextRun string `json:"nextRun"`

	// The date that the schedule is effective from. It can be null, so the schedule is effective from now
	EffectiveDate string `json:"effectiveDate"`

	// Name of the trigger item
	TriggerItemName string `json:"triggerItemName"`

	// Name of trigger item type
	TriggerItemTypeName string `json:"triggerItemTypeName"`

	// Time zone offset in ms
	TimezoneOffsetMs int `json:"timezoneOffsetMs"`

	// Relative time of the next job run in ms
	NextRunRelative int64 `json:"nextRunRelative"`
}

func (s *JobService) JobInfo(ctx context.Context, ids []int, clientTimeOffset int) (*Response, *http.Response, error) {
	request := Request{
		Action: JobAction,
		Method: "getJobInfo",
		Data:   []interface{}{ids, clientTimeOffset},
		Type:   "rpc",
		Tid:    1,
	}

	req, err := s.client.NewRequest(&request)
	if err != nil {
		return nil, nil, err
	}
	r := Response{Data: &JobInfo{}}
	resp, err := s.client.Do(ctx, req, &r)
	if err != nil {
		return nil, resp, err
	}
	return &r, resp, nil
}
