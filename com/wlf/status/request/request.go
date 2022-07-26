package request

/**
* 진위 확인
 */

// 주민 등록 확인
type IdCard struct {
	Identity   string `json:"identity"`
	IssueDate  string `json:"issueDate"`
	UserName   string `json:"userName"`
	SecretMode string `json:"secret_mode"`
}

// 운전면허증 확인
type DriverCard struct {
	JuminNo    string `json:"jumoinNo"`
	UserName   string `json:"userName"`
	BirthDate  string `json:"birthDate"`
	LicenseNo  string `json:"licenseNo"`
	SecretMode string `json:"secret_mode"`
}

// 한국 여권 확인
type PassPort struct {
	UserName       string `json:"userName"`
	PassportNo     string `json:"passportNo"`
	IssueDate      string `json:"issueDate"`
	ExpirationDate string `json:"expirationDate"`
	BirthDate      string `json:"birthDate"`
	SecretMode     string `json:"secret_mode"`
}

//외국인 여권

//외국인 등록증

//사업자등록 및 휴폐업 조회
