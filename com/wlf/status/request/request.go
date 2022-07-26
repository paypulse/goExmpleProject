package request

/**
* 진위 확인
 */

// 주민 등록 확인
type IdCard struct {
	Identity   string `json:"identity"`
	IssueDate  string `json:"issueDate"`
	UserName   string `json:"userName"`
	SecretMode bool   `json:"secret_mode"`
}

// 운전면허증 확인
type DriverCard struct {
	JuminNo    string `json:"juminNo"`
	UserName   string `json:"userName"`
	BirthDate  string `json:"birthDate"`
	LicenseNo  string `json:"licenseNo"`
	SecretMode bool   `json:"secret_mode"`
}

// 한국 여권 확인
type PassPort struct {
	UserName       string `json:"userName"`
	PassportNo     string `json:"passportNo"`
	IssueDate      string `json:"issueDate"`
	ExpirationDate string `json:"expirationDate"`
	BirthDate      string `json:"birthDate"`
	SecretMode     bool   `json:"secret_mode"`
}

// 외국인 여권
type PassPortOverSeas struct {
	PassportNo  string `json:"passportNo"`
	Nationality string `json:"nationality"`
	BirthDate   string `json:"birthDate"`
	SecretMode  bool   `json:"secret_mode"`
}

// 외국인 등록증
type Alien struct {
	IssueNo    string `json:"issueNo"`
	IssueDate  string `json:"issueDate"`
	SecretMode bool   `json:"secret_mode"`
}

// 사업자등록 및 휴폐업 조회
type BusinessRegistraction struct {
	BizNo string `json:"biz_no"`
}
