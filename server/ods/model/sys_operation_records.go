package model

import (
    "github.com/andphp/go-gin/goby/utils"
    "time"
)

type SysOperationRecordModel struct { 
    ID            int64        `gorm:"column:id" json:"id"`                      // 
    IP            string       `gorm:"column:ip" json:"ip"`                      // 请求ip
    Method        string       `gorm:"column:method" json:"method"`              // 请求方法
    Path          string       `gorm:"column:path" json:"path"`                  // 请求路径
    Status        int          `gorm:"column:status" json:"status"`              // 请求状态
    Latency       int          `gorm:"column:latency" json:"latency"`            // 延迟
    Agent         string       `gorm:"column:agent" json:"agent"`                // 代理
    ErrorMessage  string       `gorm:"column:error_message" json:"errorMessage"` // 错误信息
    RequestBody   string       `gorm:"column:request_body" json:"requestBody"`   // 请求Body
    ResponseBody  string       `gorm:"column:response_body" json:"responseBody"` // 响应Body
    UserID        int64        `gorm:"column:user_id" json:"userID"`             // 用户ID
    Level         int          `gorm:"column:level" json:"level"`                // 1:增删改查，2：增删改，3：删改
    CreatedAt     utils.JSONTime `gorm:"column:created_at" json:"createdAt"`       // 
    UpdatedAt     utils.JSONTime `gorm:"column:updated_at" json:"updatedAt"`       // 
    DeletedAt     utils.JSONTime `gorm:"column:deleted_at" json:"deletedAt"`       // 删除时间 null未删除
}

func (model *SysOperationRecordModel) TableName() string {
	return "sys_operation_records"
}

var ColumnSysOperationRecord = struct {
    ID                   string
	IP                   string
	Method               string
	Path                 string
	Status               string
	Latency              string
	Agent                string
	ErrorMessage         string
	RequestBody          string
	ResponseBody         string
	UserID               string
	Level                string
	CreatedAt            string
	UpdatedAt            string
	DeletedAt            string
	
}{
    ID                  : "id",
    IP                  : "ip",
    Method              : "method",
    Path                : "path",
    Status              : "status",
    Latency             : "latency",
    Agent               : "agent",
    ErrorMessage        : "error_message",
    RequestBody         : "request_body",
    ResponseBody        : "response_body",
    UserID              : "user_id",
    Level               : "level",
    CreatedAt           : "created_at",
    UpdatedAt           : "updated_at",
    DeletedAt           : "deleted_at",
    
}

type SysOperationRecordModelOptions struct {
    apply func(*SysOperationRecordModel)
}

func NewSysOperationRecordModel(opts ...*SysOperationRecordModelOptions) *SysOperationRecordModel {
    m := &SysOperationRecordModel{}
    for _, opt := range opts{
        opt.apply(m)
    }
    return m
}

func SysOperationRecord_ID(id int64) *SysOperationRecordModelOptions {
    return &SysOperationRecordModelOptions{func(model *SysOperationRecordModel) {
        model.ID = id
    }}
}

func SysOperationRecord_IP(ip string) *SysOperationRecordModelOptions {
    return &SysOperationRecordModelOptions{func(model *SysOperationRecordModel) {
        model.IP = ip
    }}
}

func SysOperationRecord_Method(method string) *SysOperationRecordModelOptions {
    return &SysOperationRecordModelOptions{func(model *SysOperationRecordModel) {
        model.Method = method
    }}
}

func SysOperationRecord_Path(path string) *SysOperationRecordModelOptions {
    return &SysOperationRecordModelOptions{func(model *SysOperationRecordModel) {
        model.Path = path
    }}
}

func SysOperationRecord_Status(status int) *SysOperationRecordModelOptions {
    return &SysOperationRecordModelOptions{func(model *SysOperationRecordModel) {
        model.Status = status
    }}
}

func SysOperationRecord_Latency(latency int) *SysOperationRecordModelOptions {
    return &SysOperationRecordModelOptions{func(model *SysOperationRecordModel) {
        model.Latency = latency
    }}
}

func SysOperationRecord_Agent(agent string) *SysOperationRecordModelOptions {
    return &SysOperationRecordModelOptions{func(model *SysOperationRecordModel) {
        model.Agent = agent
    }}
}

func SysOperationRecord_ErrorMessage(error_message string) *SysOperationRecordModelOptions {
    return &SysOperationRecordModelOptions{func(model *SysOperationRecordModel) {
        model.ErrorMessage = error_message
    }}
}

func SysOperationRecord_RequestBody(request_body string) *SysOperationRecordModelOptions {
    return &SysOperationRecordModelOptions{func(model *SysOperationRecordModel) {
        model.RequestBody = request_body
    }}
}

func SysOperationRecord_ResponseBody(response_body string) *SysOperationRecordModelOptions {
    return &SysOperationRecordModelOptions{func(model *SysOperationRecordModel) {
        model.ResponseBody = response_body
    }}
}

func SysOperationRecord_UserID(user_id int64) *SysOperationRecordModelOptions {
    return &SysOperationRecordModelOptions{func(model *SysOperationRecordModel) {
        model.UserID = user_id
    }}
}

func SysOperationRecord_Level(level int) *SysOperationRecordModelOptions {
    return &SysOperationRecordModelOptions{func(model *SysOperationRecordModel) {
        model.Level = level
    }}
}

func SysOperationRecord_CreatedAt(created_at time.Time) *SysOperationRecordModelOptions {
    return &SysOperationRecordModelOptions{func(model *SysOperationRecordModel) {
        model.CreatedAt = utils.JSONTime(created_at)
    }}
}

func SysOperationRecord_UpdatedAt(updated_at time.Time) *SysOperationRecordModelOptions {
    return &SysOperationRecordModelOptions{func(model *SysOperationRecordModel) {
        model.UpdatedAt = utils.JSONTime(updated_at)
    }}
}

func SysOperationRecord_DeletedAt(deleted_at time.Time) *SysOperationRecordModelOptions {
    return &SysOperationRecordModelOptions{func(model *SysOperationRecordModel) {
        model.DeletedAt = utils.JSONTime(deleted_at)
    }}
}


func(model *SysOperationRecordModel) WithID(id int64) *SysOperationRecordModel  {
    model.ID = id
    return model
}

func(model *SysOperationRecordModel) WithIP(ip string) *SysOperationRecordModel  {
    model.IP = ip
    return model
}

func(model *SysOperationRecordModel) WithMethod(method string) *SysOperationRecordModel  {
    model.Method = method
    return model
}

func(model *SysOperationRecordModel) WithPath(path string) *SysOperationRecordModel  {
    model.Path = path
    return model
}

func(model *SysOperationRecordModel) WithStatus(status int) *SysOperationRecordModel  {
    model.Status = status
    return model
}

func(model *SysOperationRecordModel) WithLatency(latency int) *SysOperationRecordModel  {
    model.Latency = latency
    return model
}

func(model *SysOperationRecordModel) WithAgent(agent string) *SysOperationRecordModel  {
    model.Agent = agent
    return model
}

func(model *SysOperationRecordModel) WithErrorMessage(error_message string) *SysOperationRecordModel  {
    model.ErrorMessage = error_message
    return model
}

func(model *SysOperationRecordModel) WithRequestBody(request_body string) *SysOperationRecordModel  {
    model.RequestBody = request_body
    return model
}

func(model *SysOperationRecordModel) WithResponseBody(response_body string) *SysOperationRecordModel  {
    model.ResponseBody = response_body
    return model
}

func(model *SysOperationRecordModel) WithUserID(user_id int64) *SysOperationRecordModel  {
    model.UserID = user_id
    return model
}

func(model *SysOperationRecordModel) WithLevel(level int) *SysOperationRecordModel  {
    model.Level = level
    return model
}

func(model *SysOperationRecordModel) WithCreatedAt(created_at time.Time) *SysOperationRecordModel  {
    model.CreatedAt = utils.JSONTime(created_at)
    return model
}

func(model *SysOperationRecordModel) WithUpdatedAt(updated_at time.Time) *SysOperationRecordModel  {
    model.UpdatedAt = utils.JSONTime(updated_at)
    return model
}

func(model *SysOperationRecordModel) WithDeletedAt(deleted_at time.Time) *SysOperationRecordModel  {
    model.DeletedAt = utils.JSONTime(deleted_at)
    return model
}









