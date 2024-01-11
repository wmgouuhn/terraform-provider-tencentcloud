package sqlserver

const (
	SQLSERVER_CHARGE_TYPE_PREPAID  = "PREPAID"
	SQLSERVER_CHARGE_TYPE_POSTPAID = "POSTPAID_BY_HOUR"
)

var SQLSERVER_CHARGE_TYPE_NAME = map[string]string{
	"PRE":  SQLSERVER_CHARGE_TYPE_PREPAID,
	"POST": SQLSERVER_CHARGE_TYPE_POSTPAID,
	"ALL":  "ALL",
}

const (
	SQLSERVER_TASK_SUCCESS = 0
	SQLSERVER_TASK_FAIL    = 1
	SQLSERVER_TASK_RUNNING = 2
)

var SQLSERVER_CHARSET_LIST = []string{
	"Chinese_PRC_CI_AS",
	"Chinese_PRC_CS_AS",
	"Chinese_PRC_BIN",
	"Chinese_Taiwan_Stroke_CI_AS",
	"SQL_Latin1_General_CP1_CI_AS",
	"SQL_Latin1_General_CP1_CS_AS",
}

const (
	SQLSERVER_ACCOUNT_RW = "ReadWrite"
	SQLSERVER_ACCOUNT_RO = "ReadOnly"
)

var SQLSERVER_ACCOUNT_PRIVILEGE = []string{
	SQLSERVER_ACCOUNT_RW,
	SQLSERVER_ACCOUNT_RO,
}

const (
	SQLSERVER_DEFAULT_LIMIT  = 20
	SQLSERVER_DEFAULT_OFFSET = 0
)

const (
	SQLSERVER_DB_CREATING         = 1
	SQLSERVER_DB_RUNNING          = 2
	SQLSERVER_DB_MODIFYING        = 3
	SQLSERVER_DB_DELETING         = -1
	SQLSERVER_DB_UPGRADING        = 9
	SQLSERVER_DB_SWITCH_PENDING   = 13 // 实例修改中且待切换
	SQLSERVER_DB_PUBSUB_CREATING  = 14 // 订阅发布创建中
	SQLSERVER_DB_PUBSUB_MODIFYING = 15 // 订阅发布修改中
	SQLSERVER_DB_SWITCHING        = 16 // 实例修改中且切换中
	SQLSERVER_DB_RO_CREATING      = 17 // 创建RO副本中
)

var SQLSERVER_DB_STATUS = map[int64]string{
	SQLSERVER_DB_CREATING:         "creating",
	SQLSERVER_DB_RUNNING:          "running",
	SQLSERVER_DB_MODIFYING:        "modifying",
	SQLSERVER_DB_DELETING:         "deleting",
	SQLSERVER_DB_UPGRADING:        "upgrading",
	SQLSERVER_DB_SWITCH_PENDING:   "switch pending",
	SQLSERVER_DB_PUBSUB_CREATING:  "pubsub creating",
	SQLSERVER_DB_PUBSUB_MODIFYING: "pubsub modifying",
	SQLSERVER_DB_SWITCHING:        "switching",
	SQLSERVER_DB_RO_CREATING:      "ro creating",
}

var SQLSERVER_STATUS_WAITING = []int{
	SQLSERVER_DB_CREATING,
	SQLSERVER_DB_MODIFYING,
	SQLSERVER_DB_UPGRADING,
	SQLSERVER_DB_SWITCH_PENDING,
	SQLSERVER_DB_PUBSUB_CREATING,
	SQLSERVER_DB_PUBSUB_MODIFYING,
	SQLSERVER_DB_SWITCHING,
	SQLSERVER_DB_RO_CREATING,
}

const (
	SQLSERVER_BACKUP_CYCLETYPE_DAILY   = "daily"
	SQLSERVER_BACKUP_CYCLETYPE_WEEKLY  = "weekly"
	SQLSERVER_BACKUP_CYCLETYPE_MONTHLY = "monthly"
)

const (
	SQLSERVER_HA_FLAG_SINGLE  = "SINGLE"
	SQLSERVER_HA_FLAG_DAUL    = "MIRROR"
	SQLSERVER_HA_FLAG_CLUSTER = "ALWAYSON"
)

var SQLSERVER_HA_TYPE_FLAGS = map[string]string{
	SQLSERVER_HA_FLAG_SINGLE:  "SINGLE",
	SQLSERVER_HA_FLAG_DAUL:    "DUAL",
	SQLSERVER_HA_FLAG_CLUSTER: "CLUSTER",
}

const (
	BASIC_NETWORK = 0
	VPC_NEWORK    = 1
)

const (
	INTERNALERROR_DBERROR   = "InternalError.DBError"
	INSTANCE_STATUS_INVALID = "ResourceUnavailable.InstanceStatusInvalid"
)

const (
	CLOUD_PREMIUM = "CLOUD_PREMIUM"
	CLOUD_SSD     = "CLOUD_SSD"
	CLOUD_BSSD    = "CLOUD_BSSD"
	CLOUD_HSSD    = "CLOUD_HSSD"
)

const (
	SQLSERVER_MIGRATION_INIT       = 1
	SQLSERVER_MIGRATION_MIGRATING  = 4
	SQLSERVER_MIGRATION_FAILED     = 5
	SQLSERVER_MIGRATION_COMPLETED  = 6
	SQLSERVER_MIGRATION_TERMINATED = 8
)

const (
	SQLSERVER_BACKUP_RUNNING = 0
	SQLSERVER_BACKUP_SUCCESS = 1
	SQLSERVER_BACKUP_FAIL    = 2
)

const (
	SQLSERVER_BACKUP_STRATEGY_ALL    = 0
	SQLSERVER_BACKUP_STRATEGY_SINGEL = 1
)

const (
	SQLSERVER_CLONE_SUCCESS = 0
	SQLSERVER_CLONE_FAIL    = 1
	SQLSERVER_CLONE_RUNNING = 2
)

const (
	SQLSERVER_BSDBINSTANCE_STATUS_RUNNING = 1
	SQLSERVER_BSDBINSTANCE_STATUS_SUCCESS = 2
)

const (
	SQLSERVER_DB_VERSION_2016 = "201603"
	SQLSERVER_DB_VERSION_2017 = "201703"
	SQLSERVER_DB_VERSION_2019 = "201903"
)

const (
	SQLSERVER_DB_VERSION_NAME_2016 = "SQL Server 2016 Integration Services"
	SQLSERVER_DB_VERSION_NAME_2017 = "SQL Server 2017 Integration Services"
	SQLSERVER_DB_VERSION_NAME_2019 = "SQL Server 2019 Integration Services"
)

const (
	SQLSERVER_CLOUD_DB_VERSION_2008R2  = "2008R2"
	SQLSERVER_CLOUD_DB_VERSION_2012SP3 = "2012SP3"
	SQLSERVER_CLOUD_DB_VERSION_201202  = "201202"
	SQLSERVER_CLOUD_DB_VERSION_2014SP2 = "2014SP2"
	SQLSERVER_CLOUD_DB_VERSION_201402  = "201402"
	SQLSERVER_CLOUD_DB_VERSION_2016SP1 = "2016SP1"
	SQLSERVER_CLOUD_DB_VERSION_201602  = "201602"
	SQLSERVER_CLOUD_DB_VERSION_201603  = "201603"
	SQLSERVER_CLOUD_DB_VERSION_2017    = "2017"
	SQLSERVER_CLOUD_DB_VERSION_201702  = "201702"
	SQLSERVER_CLOUD_DB_VERSION_201703  = "201703"
	SQLSERVER_CLOUD_DB_VERSION_2019    = "2019"
	SQLSERVER_CLOUD_DB_VERSION_201902  = "201902"
	SQLSERVER_CLOUD_DB_VERSION_201903  = "201903"
	SQLSERVER_CLOUD_DB_VERSION_2022    = "2022"
)

const (
	SQLSERVER_CLOUD_DB_VERSION_NAME_2008R2  = "SQL Server 2008 R2 Enterprise"
	SQLSERVER_CLOUD_DB_VERSION_NAME_2012SP3 = "SQL Server 2012 Enterprise"
	SQLSERVER_CLOUD_DB_VERSION_NAME_201202  = "SQL Server 2012 Standard"
	SQLSERVER_CLOUD_DB_VERSION_NAME_2014SP2 = "SQL Server 2014 Enterprise"
	SQLSERVER_CLOUD_DB_VERSION_NAME_201402  = "SQL Server 2014 Standard"
	SQLSERVER_CLOUD_DB_VERSION_NAME_2016SP1 = "SQL Server 2016 Enterprise"
	SQLSERVER_CLOUD_DB_VERSION_NAME_201602  = "SQL Server 2016 Standard"
	SQLSERVER_CLOUD_DB_VERSION_NAME_201603  = "SQL Server 2016 Integration Services"
	SQLSERVER_CLOUD_DB_VERSION_NAME_2017    = "SQL Server 2017 Enterprise"
	SQLSERVER_CLOUD_DB_VERSION_NAME_201702  = "SQL Server 2017 Standard"
	SQLSERVER_CLOUD_DB_VERSION_NAME_201703  = "SQL Server 2017 Integration Services"
	SQLSERVER_CLOUD_DB_VERSION_NAME_2019    = "SQL Server 2019 Enterprise"
	SQLSERVER_CLOUD_DB_VERSION_NAME_201902  = "SQL Server 2019 Standard"
	SQLSERVER_CLOUD_DB_VERSION_NAME_201903  = "SQL Server 2019 Integration Services"
	SQLSERVER_CLOUD_DB_VERSION_NAME_2022    = "SQL Server 2022 Enterprise"
)

const (
	SQLSERVER_TYPE_PREPAID  = "PREPAID"
	SQLSERVER_TYPE_POSTPAID = "POSTPAID"
)

var SQLSERVER_MIGRATION_STATUS = map[int64]string{
	SQLSERVER_MIGRATION_INIT:       "init",
	SQLSERVER_MIGRATION_MIGRATING:  "migrating",
	SQLSERVER_MIGRATION_FAILED:     "failed",
	SQLSERVER_MIGRATION_COMPLETED:  "completed",
	SQLSERVER_MIGRATION_TERMINATED: "terminated",
}

const (
	CreateDefaultTimeout = 7200
	ReadDefaultTimeout   = 7200
	UpdateDefaultTimeout = 7200
	DeleteDefaultTimeout = 7200
)

const (
	SSL_TYPE_ENABLE  = "enable"
	SSL_TYPE_DISABLE = "disable"
	SSL_TYPE_RENEW   = "renew"
)

var SSL_TYPE = []string{
	SSL_TYPE_ENABLE,
	SSL_TYPE_DISABLE,
	SSL_TYPE_RENEW,
}