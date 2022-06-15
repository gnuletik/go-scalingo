package scalingo

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"github.com/Scalingo/go-scalingo/v5/debug"
)

type Event struct {
	ID          string                 `json:"id"`
	AppID       string                 `json:"app_id"`
	CreatedAt   time.Time              `json:"created_at"`
	User        EventUser              `json:"user"`
	Type        EventTypeName          `json:"type"`
	AppName     string                 `json:"app_name"`
	RawTypeData json.RawMessage        `json:"type_data"`
	TypeData    map[string]interface{} `json:"-"`
}

func (ev *Event) GetEvent() *Event {
	return ev
}

func (ev *Event) TypeDataPtr() interface{} {
	return ev.TypeData
}

func (ev *Event) String() string {
	return fmt.Sprintf("Unknown event %v on app %v", ev.Type, ev.AppName)
}

func (ev *Event) When() string {
	return ev.CreatedAt.Format("Mon Jan 02 2006 15:04:05")
}

func (ev *Event) Who() string {
	return fmt.Sprintf("%s (%s)", ev.User.Username, ev.User.Email)
}

func (ev *Event) PrintableType() string {
	return strings.Title(strings.Replace(string(ev.Type), "_", " ", -1))
}

type DetailedEvent interface {
	fmt.Stringer
	GetEvent() *Event
	PrintableType() string
	When() string
	Who() string
	TypeDataPtr() interface{}
}

type Events []DetailedEvent

type EventUser struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	ID       string `json:"id"`
}

type EventTypeName string

const (
	EventNewUser                 EventTypeName = "new_user"
	EventNewApp                  EventTypeName = "new_app"
	EventEditApp                 EventTypeName = "edit_app"
	EventDeleteApp               EventTypeName = "delete_app"
	EventRenameApp               EventTypeName = "rename_app"
	EventTransferApp             EventTypeName = "transfer_app"
	EventRestart                 EventTypeName = "restart"
	EventScale                   EventTypeName = "scale"
	EventStopApp                 EventTypeName = "stop_app"
	EventCrash                   EventTypeName = "crash"
	EventRepeatedCrash           EventTypeName = "repeated_crash"
	EventDeployment              EventTypeName = "deployment"
	EventLinkSCM                 EventTypeName = "link_scm"
	EventUnlinkSCM               EventTypeName = "unlink_scm"
	EventNewIntegration          EventTypeName = "new_integration"
	EventDeleteIntegration       EventTypeName = "delete_integration"
	EventAuthorizeGithub         EventTypeName = "authorize_github"
	EventRevokeGithub            EventTypeName = "revoke_github"
	EventRun                     EventTypeName = "run"
	EventNewDomain               EventTypeName = "new_domain"
	EventEditDomain              EventTypeName = "edit_domain"
	EventDeleteDomain            EventTypeName = "delete_domain"
	EventUpgradeDatabase         EventTypeName = "upgrade_database"
	EventNewAddon                EventTypeName = "new_addon"
	EventUpgradeAddon            EventTypeName = "upgrade_addon"
	EventDeleteAddon             EventTypeName = "delete_addon"
	EventResumeAddon             EventTypeName = "resume_addon"
	EventSuspendAddon            EventTypeName = "suspend_addon"
	EventNewCollaborator         EventTypeName = "new_collaborator"
	EventAcceptCollaborator      EventTypeName = "accept_collaborator"
	EventDeleteCollaborator      EventTypeName = "delete_collaborator"
	EventNewVariable             EventTypeName = "new_variable"
	EventEditVariable            EventTypeName = "edit_variable"
	EventEditVariables           EventTypeName = "edit_variables"
	EventDeleteVariable          EventTypeName = "delete_variable"
	EventAddCredit               EventTypeName = "add_credit"
	EventAddPaymentMethod        EventTypeName = "add_payment_method"
	EventAddVoucher              EventTypeName = "add_voucher"
	EventNewKey                  EventTypeName = "new_key"
	EventEditKey                 EventTypeName = "edit_key"
	EventDeleteKey               EventTypeName = "delete_key"
	EventPaymentAttempt          EventTypeName = "payment_attempt"
	EventNewAlert                EventTypeName = "new_alert"
	EventAlert                   EventTypeName = "alert"
	EventDeleteAlert             EventTypeName = "delete_alert"
	EventNewAutoscaler           EventTypeName = "new_autoscaler"
	EventEditAutoscaler          EventTypeName = "edit_autoscaler"
	EventDeleteAutoscaler        EventTypeName = "delete_autoscaler"
	EventAddonUpdated            EventTypeName = "addon_updated"
	EventStartRegionMigration    EventTypeName = "start_region_migration"
	EventNewLogDrain             EventTypeName = "new_log_drain"
	EventDeleteLogDrain          EventTypeName = "delete_log_drain"
	EventNewAddonLogDrain        EventTypeName = "new_addon_log_drain"
	EventDeleteAddonLogDrain     EventTypeName = "delete_addon_log_drain"
	EventNewNotifier             EventTypeName = "new_notifier"
	EventEditNotifier            EventTypeName = "edit_notifier"
	EventDeleteNotifier          EventTypeName = "delete_notifier"
	EventEditHDSContact          EventTypeName = "edit_hds_contact"
	EventCreateDataAccessConsent EventTypeName = "create_data_access_consent"
	EventNewToken                EventTypeName = "new_token"
	EventRegenerateToken         EventTypeName = "regenerate_token"
	EventDeleteToken             EventTypeName = "delete_token"
	EventTfaEnabled              EventTypeName = "tfa_enabled"
	EventTfaDisabled             EventTypeName = "tfa_disabled"
	EventLoginSuccess            EventTypeName = "login_success"
	EventLoginFailure            EventTypeName = "login_failure"
	EventLoginLock               EventTypeName = "login_lock"
	EventLoginUnlockSuccess      EventTypeName = "login_unlock_success"
	EventPasswordResetQuery      EventTypeName = "password_reset_query"
	EventPasswordResetSuccess    EventTypeName = "password_reset_success"

	// EventLinkGithub and EventUnlinkGithub events are kept for
	// retro-compatibility. They are replaced by SCM events.
	EventLinkGithub   EventTypeName = "link_github"
	EventUnlinkGithub EventTypeName = "unlink_github"
)

type EventNewUserType struct {
	Event
	TypeData EventNewUserTypeData `json:"type_data"`
}

func (ev *EventNewUserType) String() string {
	return "You joined Scalingo. Hooray!"
}

type EventNewUserTypeData struct {
}

type EventLinkGithubType struct {
	Event
	TypeData EventLinkGithubTypeData `json:"type_data"`
}

func (ev *EventLinkGithubType) String() string {
	return fmt.Sprintf("app has been linked to Github repository '%s'", ev.TypeData.RepoName)
}

type EventLinkGithubTypeData struct {
	RepoName       string `json:"repo_name"`
	LinkerUsername string `json:"linker_username"`
	GithubSource   string `json:"github_source"`
}

type EventUnlinkGithubType struct {
	Event
	TypeData EventUnlinkGithubTypeData `json:"type_data"`
}

func (ev *EventUnlinkGithubType) String() string {
	return fmt.Sprintf("app has been unlinked from Github repository '%s'", ev.TypeData.RepoName)
}

type EventUnlinkGithubTypeData struct {
	RepoName         string `json:"repo_name"`
	UnlinkerUsername string `json:"unlinker_username"`
	GithubSource     string `json:"github_source"`
}

type EventLinkSCMType struct {
	Event
	TypeData EventLinkSCMTypeData `json:"type_data"`
}

func (ev *EventLinkSCMType) String() string {
	return fmt.Sprintf("app has been linked to repository '%s'", ev.TypeData.RepoName)
}

type EventLinkSCMTypeData struct {
	RepoName       string `json:"repo_name"`
	LinkerUsername string `json:"linker_username"`
	Source         string `json:"source"`
}

type EventUnlinkSCMType struct {
	Event
	TypeData EventUnlinkSCMTypeData `json:"type_data"`
}

func (ev *EventUnlinkSCMType) String() string {
	return fmt.Sprintf("app has been unlinked from repository '%s'", ev.TypeData.RepoName)
}

type EventUnlinkSCMTypeData struct {
	RepoName         string `json:"repo_name"`
	UnlinkerUsername string `json:"unlinker_username"`
	Source           string `json:"source"`
}

type EventRunType struct {
	Event
	TypeData EventRunTypeData `json:"type_data"`
}

func (ev *EventRunType) String() string {
	return fmt.Sprintf("one-off container with command '%s'", ev.TypeData.Command)
}

type EventRunTypeData struct {
	Command    string `json:"command"`
	AuditLogID string `json:"audit_log_id"`
}

type EventNewDomainType struct {
	Event
	TypeData EventNewDomainTypeData `json:"type_data"`
}

func (ev *EventNewDomainType) String() string {
	return fmt.Sprintf("'%s' has been associated", ev.TypeData.Hostname)
}

type EventNewDomainTypeData struct {
	Hostname string `json:"hostname"`
	SSL      bool   `json:"ssl"`
}

type EventEditDomainType struct {
	Event
	TypeData EventEditDomainTypeData `json:"type_data"`
}

func (ev *EventEditDomainType) String() string {
	t := ev.TypeData
	res := fmt.Sprintf("'%s' modified", t.Hostname)
	if !t.SSL && t.OldSSL {
		res += ", TLS certificate has been removed"
	} else if !t.SSL && t.OldSSL {
		res += ", TLS certificate has been added"
	} else if t.SSL && t.OldSSL {
		res += ", TLS certificate has been changed"
	}
	return res
}

type EventEditDomainTypeData struct {
	Hostname string `json:"hostname"`
	SSL      bool   `json:"ssl"`
	OldSSL   bool   `json:"old_ssl"`
}

type EventDeleteDomainType struct {
	Event
	TypeData EventDeleteDomainTypeData `json:"type_data"`
}

func (ev *EventDeleteDomainType) String() string {
	return fmt.Sprintf("'%s' has been disassociated", ev.TypeData.Hostname)
}

type EventDeleteDomainTypeData struct {
	Hostname string `json:"hostname"`
}

type EventCollaborator struct {
	EventUser
	Inviter EventUser `json:"inviter"`
}

type EventNewCollaboratorType struct {
	Event
	TypeData EventNewCollaboratorTypeData `json:"type_data"`
}

func (ev *EventNewCollaboratorType) String() string {
	return fmt.Sprintf(
		"'%s' has been invited",
		ev.TypeData.Collaborator.Email,
	)
}

type EventNewCollaboratorTypeData struct {
	Collaborator EventCollaborator `json:"collaborator"`
}

type EventAcceptCollaboratorType struct {
	Event
	TypeData EventAcceptCollaboratorTypeData `json:"type_data"`
}

func (ev *EventAcceptCollaboratorType) String() string {
	return fmt.Sprintf(
		"'%s' (%s) has accepted the collaboration",
		ev.TypeData.Collaborator.Username,
		ev.TypeData.Collaborator.Email,
	)
}

// Inviter is filled there
type EventAcceptCollaboratorTypeData struct {
	Collaborator EventCollaborator `json:"collaborator"`
}

type EventDeleteCollaboratorType struct {
	Event
	TypeData EventDeleteCollaboratorTypeData `json:"type_data"`
}

func (ev *EventDeleteCollaboratorType) String() string {
	return fmt.Sprintf(
		"'%s' (%s) is not a collaborator anymore",
		ev.TypeData.Collaborator.Username,
		ev.TypeData.Collaborator.Email,
	)
}

type EventDeleteCollaboratorTypeData struct {
	Collaborator EventCollaborator `json:"collaborator"`
}

type EventUpgradeDatabaseType struct {
	Event
	TypeData EventUpgradeDatabaseTypeData `json:"type_data"`
}

type EventUpgradeDatabaseTypeData struct {
	AddonName  string `json:"addon_name"`
	OldVersion string `json:"old_version"`
	NewVersion string `json:"new_version"`
}

func (ev *EventUpgradeDatabaseType) String() string {
	return fmt.Sprintf(
		"'%s' upgraded from v%s to v%s",
		ev.TypeData.AddonName, ev.TypeData.OldVersion, ev.TypeData.NewVersion,
	)
}

func (ev *EventUpgradeDatabaseType) Who() string {
	if ev.TypeData.AddonName != "" {
		return fmt.Sprintf("Addon %s", ev.TypeData.AddonName)
	} else {
		return ev.Event.Who()
	}
}

type EventVariable struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

type EventNewVariableType struct {
	Event
	TypeData EventNewVariableTypeData `json:"type_data"`
}

func (ev *EventNewVariableType) String() string {
	return fmt.Sprintf("'%s' added to the environment", ev.TypeData.Name)
}

func (ev *EventNewVariableType) Who() string {
	if ev.TypeData.AddonName != "" {
		return fmt.Sprintf("Addon %s", ev.TypeData.AddonName)
	} else {
		return ev.Event.Who()
	}
}

type EventNewVariableTypeData struct {
	AddonName string `json:"addon_name"`
	EventVariable
}

type EventVariables []EventVariable

func (evs EventVariables) Names() string {
	names := []string{}
	for _, e := range evs {
		names = append(names, e.Name)
	}
	return strings.Join(names, ", ")
}

type EventEditVariableType struct {
	Event
	TypeData EventEditVariableTypeData `json:"type_data"`
}

func (ev *EventEditVariableType) String() string {
	return fmt.Sprintf("'%s' modified", ev.TypeData.Name)
}

type EventEditVariableTypeData struct {
	EventVariable
	OldValue  string `json:"old_value"`
	AddonName string `json:"addon_name"`
}

type EventEditVariablesType struct {
	Event
	TypeData EventEditVariablesTypeData `json:"type_data"`
}

func (ev *EventEditVariablesType) String() string {
	res := "environment changes:"
	if len(ev.TypeData.NewVars) > 0 {
		res += fmt.Sprintf(" %s added", ev.TypeData.NewVars.Names())
	}
	if len(ev.TypeData.UpdatedVars) > 0 {
		res += fmt.Sprintf(" %s modified", ev.TypeData.UpdatedVars.Names())
	}
	if len(ev.TypeData.DeletedVars) > 0 {
		res += fmt.Sprintf(" %s removed", ev.TypeData.DeletedVars.Names())
	}
	return res
}

func (ev *EventEditVariableType) Who() string {
	if ev.TypeData.AddonName != "" {
		return fmt.Sprintf("Addon %s", ev.TypeData.AddonName)
	} else {
		return ev.Event.Who()
	}
}

type EventEditVariablesTypeData struct {
	NewVars     EventVariables `json:"new_vars"`
	UpdatedVars EventVariables `json:"updated_vars"`
	DeletedVars EventVariables `json:"deleted_vars"`
}

type EventDeleteVariableType struct {
	Event
	TypeData EventDeleteVariableTypeData `json:"type_data"`
}

func (ev *EventDeleteVariableType) String() string {
	return fmt.Sprintf("'%s' removed from environment", ev.TypeData.Name)
}

type EventDeleteVariableTypeData struct {
	EventVariable
}

type EventPaymentAttemptTypeData struct {
	Amount        float32 `json:"amount"`
	PaymentMethod string  `json:"payment_method"`
	Status        string  `json:"status"`
}

type EventPaymentAttemptType struct {
	Event
	TypeData EventPaymentAttemptTypeData `json:"type_data"`
}

func (ev *EventPaymentAttemptType) String() string {
	res := "Payment attempt of "
	res += fmt.Sprintf("%0.2f€", ev.TypeData.Amount)
	res += " with your "
	if ev.TypeData.PaymentMethod == "credit" {
		res += "credits"
	} else {
		res += "card"
	}
	if ev.TypeData.Status == "new" {
		res += " (pending)"
	} else if ev.TypeData.Status == "paid" {
		res += " (success)"
	} else {
		res += " (fail)"
	}
	return res
}

type EventNewAutoscalerTypeData struct {
	ContainerType string  `json:"container_type"`
	MinContainers int     `json:"min_containers,string"`
	MaxContainers int     `json:"max_containers,string"`
	Metric        string  `json:"metric"`
	Target        float64 `json:"target"`
	TargetText    string  `json:"target_text"`
}

type EventNewAutoscalerType struct {
	Event
	TypeData EventNewAutoscalerTypeData `json:"type_data"`
}

func (ev *EventNewAutoscalerType) String() string {
	d := ev.TypeData
	return fmt.Sprintf("Autoscaler created about %s on container %s (target: %s)", d.Metric, d.ContainerType, d.TargetText)
}

type EventEditAutoscalerTypeData struct {
	ContainerType string  `json:"container_type"`
	MinContainers int     `json:"min_containers,string"`
	MaxContainers int     `json:"max_containers,string"`
	Metric        string  `json:"metric"`
	Target        float64 `json:"target"`
	TargetText    string  `json:"target_text"`
}

type EventEditAutoscalerType struct {
	Event
	TypeData EventEditAutoscalerTypeData `json:"type_data"`
}

func (ev *EventEditAutoscalerType) String() string {
	d := ev.TypeData
	return fmt.Sprintf("Autoscaler edited about %s on container %s (target: %s)", d.Metric, d.ContainerType, d.TargetText)
}

type EventDeleteAutoscalerTypeData struct {
	ContainerType string `json:"container_type"`
	Metric        string `json:"metric"`
}

type EventDeleteAutoscalerType struct {
	Event
	TypeData EventDeleteAutoscalerTypeData `json:"type_data"`
}

func (ev *EventDeleteAutoscalerType) String() string {
	d := ev.TypeData
	return fmt.Sprintf("Alert deleted about %s on container %s", d.Metric, d.ContainerType)
}

type EventStartRegionMigrationTypeData struct {
	MigrationID string `json:"migration_id"`
	Destination string `json:"destination"`
	Source      string `json:"source"`
	DstAppName  string `json:"dst_app_name"`
}

type EventStartRegionMigrationType struct {
	Event
	TypeData EventStartRegionMigrationTypeData `json:"type_data"`
}

func (ev *EventStartRegionMigrationType) String() string {
	return fmt.Sprintf("Application region migration started from %s to %s/%s", ev.TypeData.Source, ev.TypeData.Destination, ev.TypeData.DstAppName)
}

// New log drain
type EventNewLogDrainTypeData struct {
	URL string `json:"url"`
}

type EventNewLogDrainType struct {
	Event
	TypeData EventNewLogDrainTypeData `json:"type_data"`
}

func (ev *EventNewLogDrainType) String() string {
	return fmt.Sprintf("Log drain added on %s app", ev.AppName)
}

// Delete log drain
type EventDeleteLogDrainTypeData struct {
	URL string `json:"url"`
}

type EventDeleteLogDrainType struct {
	Event
	TypeData EventDeleteLogDrainTypeData `json:"type_data"`
}

func (ev *EventDeleteLogDrainType) String() string {
	return fmt.Sprintf("Log drain deleted on %s app", ev.AppName)
}

// New addon log drain
type EventNewAddonLogDrainTypeData struct {
	URL       string `json:"url"`
	AddonUUID string `json:"addon_uuid"`
	AddonName string `json:"addon_name"`
}

type EventNewAddonLogDrainType struct {
	Event
	TypeData EventNewAddonLogDrainTypeData `json:"type_data"`
}

func (ev *EventNewAddonLogDrainType) String() string {
	return fmt.Sprintf("Log drain added for %s addon on %s app", ev.TypeData.AddonName, ev.AppName)
}

// Delete addon log drain
type EventDeleteAddonLogDrainTypeData struct {
	URL       string `json:"url"`
	AddonUUID string `json:"addon_uuid"`
	AddonName string `json:"addon_name"`
}

type EventDeleteAddonLogDrainType struct {
	Event
	TypeData EventDeleteAddonLogDrainTypeData `json:"type_data"`
}

func (ev *EventDeleteAddonLogDrainType) String() string {
	return fmt.Sprintf("Log drain deleted on %s addon for %s app", ev.TypeData.AddonName, ev.AppName)
}

// New notifier
type EventNewNotifierTypeData struct {
	NotifierName     string                 `json:"notifier_name"`
	Active           bool                   `json:"active"`
	SendAllEvents    bool                   `json:"send_all_events"`
	SelectedEvents   []string               `json:"selected_events"`
	NotifierType     string                 `json:"notifier_type"`
	NotifierTypeData map[string]interface{} `json:"notifier_type_data"`
	PlatformName     string                 `json:"platform_name"`
}

type EventNewNotifierType struct {
	Event
	TypeData EventNewNotifierTypeData `json:"type_data"`
}

var NotifierPlatformNames = map[string]string{
	"email":       "E-mail",
	"rocker_chat": "Rocket Chat",
	"slack":       "Slack",
	"webhook":     "Webhook",
}

func (ev *EventNewNotifierType) String() string {
	d := ev.TypeData
	platformName, ok := NotifierPlatformNames[d.PlatformName]
	if !ok {
		platformName = "unknown"
	}
	return fmt.Sprintf("Notifier '%s' created for the platform '%s' on %s app", d.NotifierName, platformName, ev.AppName)
}

// Edit notifier
type EventEditNotifierTypeData struct {
	NotifierName     string                 `json:"notifier_name"`
	Active           bool                   `json:"active"`
	SendAllEvents    bool                   `json:"send_all_events"`
	SelectedEvents   []string               `json:"selected_events"`
	NotifierType     string                 `json:"notifier_type"`
	NotifierTypeData map[string]interface{} `json:"notifier_type_data"`
	PlatformName     string                 `json:"platform_name"`
}

type EventEditNotifierType struct {
	Event
	TypeData EventEditNotifierTypeData `json:"type_data"`
}

func (ev *EventEditNotifierType) String() string {
	d := ev.TypeData
	return fmt.Sprintf("Notifier '%s' edited on %s app", d.NotifierName, ev.AppName)
}

// Delete notifier
type EventDeleteNotifierTypeData struct {
	NotifierName     string                 `json:"notifier_name"`
	Active           bool                   `json:"active"`
	SendAllEvents    bool                   `json:"send_all_events"`
	SelectedEvents   []string               `json:"selected_events"`
	NotifierType     string                 `json:"notifier_type"`
	NotifierTypeData map[string]interface{} `json:"notifier_type_data"`
	PlatformName     string                 `json:"platform_name"`
}

type EventDeleteNotifierType struct {
	Event
	TypeData EventDeleteNotifierTypeData `json:"type_data"`
}

func (ev *EventDeleteNotifierType) String() string {
	d := ev.TypeData
	return fmt.Sprintf("Notifier '%s' deleted on %s app", d.NotifierName, ev.AppName)
}

// Edit hds_contact
type EventEditHDSContactTypeData struct {
	Name           string `json:"name"`
	Email          string `json:"email"`
	PhoneNumber    string `json:"phone_number"`
	Company        string `json:"company"`
	AddressLine1   string `json:"address_line1"`
	AddressLine2   string `json:"address_line2"`
	AddressCity    string `json:"address_city"`
	AddressZip     string `json:"address_zip"`
	AddressCountry string `json:"address_country"`
	Notes          string `json:"notes"`
}

type EventEditHDSContactType struct {
	Event
	TypeData EventEditHDSContactTypeData `json:"type_data"`
}

func (ev *EventEditHDSContactType) String() string {
	d := ev.TypeData
	return fmt.Sprintf("Contact health Professional '%s' edited on %s app", d.Name, ev.AppName)
}

// Create data_access_consent
type EventCreateDataAccessConsentTypeData struct {
	EndAt      time.Time `json:"end_at"`
	Databases  bool      `json:"databases"`
	Containers bool      `json:"containers"`
}

type EventCreateDataAccessConsentType struct {
	Event
	TypeData EventCreateDataAccessConsentTypeData `json:"type_data"`
}

func (ev *EventCreateDataAccessConsentType) String() string {
	d := ev.TypeData
	res := "Additional access "
	if d.Containers {
		res += "to application runtime environment, "
	}
	if d.Databases {
		res += "to databases metadata and monitoring data, "
	}
	res += fmt.Sprintf("created on %s app", ev.AppName)
	return res
}

// Enable Tfa
type EventTfaEnabledTypeData struct {
	Provider string `json:"provider"`
}

type EventTfaEnabledType struct {
	Event
	TypeData EventTfaEnabledTypeData `json:"type_data"`
}

func (ev *EventTfaEnabledType) String() string {
	return fmt.Sprintf("Two factor authentication enabled by %s", ev.TypeData.Provider)
}

// Disable Tfa
type EventTfaDisabledTypeData struct {
}

type EventTfaDisabledType struct {
	Event
	TypeData EventTfaDisabledTypeData `json:"type_data"`
}

func (ev *EventTfaDisabledType) String() string {
	return "Two factor authentication disabled"
}

// Security events
type EventSecurityTypeData struct {
	RemoteIP string `json:"remote_ip"`
}

type EventLoginSuccessType struct {
	Event
	TypeData EventLoginSuccessTypeData `json:"type_data"`
}
type EventLoginSuccessTypeData EventSecurityTypeData

func (ev *EventLoginSuccessType) String() string {
	return fmt.Sprintf("Successful login from %v", ev.TypeData.RemoteIP)
}

type EventLoginFailureType struct {
	Event
	TypeData EventLoginFailureTypeData `json:"type_data"`
}
type EventLoginFailureTypeData EventSecurityTypeData

func (ev *EventLoginFailureType) String() string {
	return fmt.Sprintf("Failed login attempt from %v", ev.TypeData.RemoteIP)
}

type EventLoginLockType struct {
	Event
	TypeData EventLoginLockTypeData `json:"type_data"`
}
type EventLoginLockTypeData EventSecurityTypeData

func (ev *EventLoginLockType) String() string {
	return fmt.Sprintf("Account is locked")
}

type EventLoginUnlockSuccessType struct {
	Event
	TypeData EventLoginUnlockSuccessTypeData `json:"type_data"`
}
type EventLoginUnlockSuccessTypeData EventSecurityTypeData

func (ev *EventLoginUnlockSuccessType) String() string {
	return fmt.Sprintf("Account unlocked from %v", ev.TypeData.RemoteIP)
}

type EventPasswordResetQueryType struct {
	Event
	TypeData EventPasswordResetQueryTypeData `json:"type_data"`
}
type EventPasswordResetQueryTypeData EventSecurityTypeData

func (ev *EventPasswordResetQueryType) String() string {
	return fmt.Sprintf("Password reset process initiated from %v", ev.TypeData.RemoteIP)
}

type EventPasswordResetSuccessType struct {
	Event
	TypeData EventPasswordResetSuccessTypeData `json:"type_data"`
}
type EventPasswordResetSuccessTypeData EventSecurityTypeData

func (ev *EventPasswordResetSuccessType) String() string {
	return fmt.Sprintf("Password changed from from %v", ev.TypeData.RemoteIP)
}

func (pev *Event) Specialize() DetailedEvent {
	var e DetailedEvent
	ev := *pev
	switch ev.Type {
	case EventNewUser:
		e = &EventNewUserType{Event: ev}
	case EventNewApp:
		e = &EventNewAppType{Event: ev}
	case EventEditApp:
		e = &EventEditAppType{Event: ev}
	case EventDeleteApp:
		e = &EventDeleteAppType{Event: ev}
	case EventRenameApp:
		e = &EventRenameAppType{Event: ev}
	case EventTransferApp:
		e = &EventTransferAppType{Event: ev}
	case EventRestart:
		e = &EventRestartType{Event: ev}
	case EventStopApp:
		e = &EventStopAppType{Event: ev}
	case EventScale:
		e = &EventScaleType{Event: ev}
	case EventCrash:
		e = &EventCrashType{Event: ev}
	case EventLinkSCM:
		e = &EventLinkSCMType{Event: ev}
	case EventUnlinkSCM:
		e = &EventUnlinkSCMType{Event: ev}
	case EventNewIntegration:
		e = &EventNewIntegrationType{Event: ev}
	case EventDeleteIntegration:
		e = &EventDeleteIntegrationType{Event: ev}
	case EventAuthorizeGithub:
		e = &EventAuthorizeGithubType{Event: ev}
	case EventRevokeGithub:
		e = &EventRevokeGithubType{Event: ev}
	case EventDeployment:
		e = &EventDeploymentType{Event: ev}
	case EventRun:
		e = &EventRunType{Event: ev}
	case EventNewDomain:
		e = &EventNewDomainType{Event: ev}
	case EventEditDomain:
		e = &EventEditDomainType{Event: ev}
	case EventDeleteDomain:
		e = &EventDeleteDomainType{Event: ev}
	case EventUpgradeDatabase:
		e = &EventUpgradeDatabaseType{Event: ev}
	case EventNewAddon:
		e = &EventNewAddonType{Event: ev}
	case EventUpgradeAddon:
		e = &EventUpgradeAddonType{Event: ev}
	case EventDeleteAddon:
		e = &EventDeleteAddonType{Event: ev}
	case EventResumeAddon:
		e = &EventResumeAddonType{Event: ev}
	case EventSuspendAddon:
		e = &EventSuspendAddonType{Event: ev}
	case EventNewCollaborator:
		e = &EventNewCollaboratorType{Event: ev}
	case EventAcceptCollaborator:
		e = &EventAcceptCollaboratorType{Event: ev}
	case EventDeleteCollaborator:
		e = &EventDeleteCollaboratorType{Event: ev}
	case EventNewVariable:
		e = &EventNewVariableType{Event: ev}
	case EventEditVariable:
		e = &EventEditVariableType{Event: ev}
	case EventEditVariables:
		e = &EventEditVariablesType{Event: ev}
	case EventDeleteVariable:
		e = &EventDeleteVariableType{Event: ev}
	case EventAddCredit:
		e = &EventAddCreditType{Event: ev}
	case EventAddPaymentMethod:
		e = &EventAddPaymentMethodType{Event: ev}
	case EventAddVoucher:
		e = &EventAddVoucherType{Event: ev}
	case EventNewKey:
		e = &EventNewKeyType{Event: ev}
	case EventEditKey:
		e = &EventEditKeyType{Event: ev}
	case EventDeleteKey:
		e = &EventDeleteKeyType{Event: ev}
	case EventPaymentAttempt:
		e = &EventPaymentAttemptType{Event: ev}
	case EventNewAlert:
		e = &EventNewAlertType{Event: ev}
	case EventAlert:
		e = &EventAlertType{Event: ev}
	case EventDeleteAlert:
		e = &EventDeleteAlertType{Event: ev}
	case EventNewAutoscaler:
		e = &EventNewAutoscalerType{Event: ev}
	case EventEditAutoscaler:
		e = &EventEditAutoscalerType{Event: ev}
	case EventDeleteAutoscaler:
		e = &EventDeleteAutoscalerType{Event: ev}
	case EventAddonUpdated:
		e = &EventAddonUpdatedType{Event: ev}
	case EventStartRegionMigration:
		e = &EventStartRegionMigrationType{Event: ev}
	case EventNewLogDrain:
		e = &EventNewLogDrainType{Event: ev}
	case EventDeleteLogDrain:
		e = &EventDeleteLogDrainType{Event: ev}
	case EventNewAddonLogDrain:
		e = &EventNewAddonLogDrainType{Event: ev}
	case EventDeleteAddonLogDrain:
		e = &EventDeleteAddonLogDrainType{Event: ev}
	case EventNewNotifier:
		e = &EventNewNotifierType{Event: ev}
	case EventEditNotifier:
		e = &EventEditNotifierType{Event: ev}
	case EventDeleteNotifier:
		e = &EventDeleteNotifierType{Event: ev}
	case EventEditHDSContact:
		e = &EventEditHDSContactType{Event: ev}
	case EventCreateDataAccessConsent:
		e = &EventCreateDataAccessConsentType{Event: ev}
	case EventNewToken:
		e = &EventNewTokenType{Event: ev}
	case EventRegenerateToken:
		e = &EventRegenerateTokenType{Event: ev}
	case EventDeleteToken:
		e = &EventDeleteTokenType{Event: ev}
	case EventTfaEnabled:
		e = &EventTfaEnabledType{Event: ev}
	case EventTfaDisabled:
		e = &EventTfaDisabledType{Event: ev}
	// Deprecated events. Replaced by equivalent with SCM in the name instead of
	// Github
	case EventLinkGithub:
		e = &EventLinkGithubType{Event: ev}
	case EventUnlinkGithub:
		e = &EventUnlinkGithubType{Event: ev}
	case EventLoginSuccess:
		e = &EventLoginSuccessType{Event: ev}
	case EventLoginFailure:
		e = &EventLoginFailureType{Event: ev}
	case EventLoginLock:
		e = &EventLoginLockType{Event: ev}
	case EventLoginUnlockSuccess:
		e = &EventLoginUnlockSuccessType{Event: ev}
	case EventPasswordResetQuery:
		e = &EventPasswordResetQueryType{Event: ev}
	case EventPasswordResetSuccess:
		e = &EventPasswordResetSuccessType{Event: ev}
	default:
		return pev
	}
	err := json.Unmarshal(pev.RawTypeData, e.TypeDataPtr())
	if err != nil {
		debug.Printf("error reading the data: %+v\n", err)
		return pev
	}
	return e
}
