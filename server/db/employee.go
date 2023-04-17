package db

import (
	"time"

	"github.com/lib/pq"
)

type EmployeeTableSetting struct {
	Border              bool      `json:"border" description:"是否带有纵向边框"`
	Fit                 bool      `json:"fit" description:"列的宽度是否自撑开"`
	HighlightCurrentRow bool      `json:"highlight_current_row" description:"是否要高亮当前行"`
	Height              uint32    `description:"Table 的高度， 默认为自动高度"`
	TableLayout         string    `json:"table_layout" description:"表格布局, auto or fixed"`
	EmptyText           string    `json:"empty_text" description:"空数据时显示的文本内容"`
	RowKey              string    `json:"row_key" description:"行数据的 Key, 用来优化 Table 的渲染"`
	UpdatedAt           time.Time `json:"update_at" description:"更新时间"`
}

type EmployeeColum struct {
	Visible   bool      `json:"visible" description:"是否可显示"`
	Sortable  bool      `json:"sortable" description:"是否可排序"`
	Prop      string    `json:"prop" description:"列属性"`
	Label     string    `json:"label" description:"列标题"`
	Align     string    `json:"align" description:"对齐方式"`
	UpdatedAt time.Time `json:"update_at" description:"更新时间"`
}

type EmployeePaginationSetting struct {
	Layout             string        `json:"layout" description:"组件布局, 子组件名用逗号分隔 prev, pager, next, jumper, ->, total"`
	PrevText           string        `json:"prev_text" description:"替代图标显示的上一页文字"`
	NextText           string        `json:"next_text" description:"替代图标显示的下一页文字"`
	PageSizes          pq.Int32Array `gorm:"type:text[]" json:"page_sizes" description:"每页显示个数选择器的选项设置"`
	DefaultPageSize    uint          `json:"default_page_size" description:"每页显示条目数的初始值"`
	DefaultCurrentPage uint          `json:"default_current_page" description:"当前页数的初始值"`
	HideOnSinglePage   bool          `json:"hide_on_single_page" description:"只有一页时是否隐藏"`
	UpdatedAt          time.Time     `json:"update_at" description:"更新时间"`
}

type Employee struct {
	Id              uint64    `json:"id" description:"用户ID"`
	Name            string    `json:"name" description:"姓名"`
	Gender          string    `json:"gender" description:"性别"`
	Birthday        time.Time `json:"birthday" description:"出生日期"`
	Height          uint64    `json:"height" description:"身高"`
	Weight          uint64    `json:"weight" description:"体重"`
	Degree          string    `json:"degree" description:"学历"`
	Identifier      string    `json:"identifier" description:"身份证号"`
	Phone           string    `json:"phone" description:"手机号码"`
	PoliticalStatus string    `json:"political_status" description:"政治面貌"`
	SocialSecurity  string    `json:"social_security" description:"社保"`
	CurrentAddress  string    `json:"current_address" description:"现住址"`
	FirstWorkTime   time.Time `json:"first_work_time" description:"首次工作"`
	FormerEmployer  string    `json:"former_employer" description:"原单位"`
	Post            string    `json:"post" description:"岗位"`
	Salary          uint64    `json:"salary" description:"工资"`
	SecurityCard    string    `json:"security_card" description:"保安证"`
	Comments        string    `json:"comments" description:"备注"`
	UpdatedAt       time.Time `json:"updated_at" description:"更新时间"`
}
