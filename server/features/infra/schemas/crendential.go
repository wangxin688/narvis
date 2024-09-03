package schemas

type CliCredentialCreate struct {
	Username string  `json:"username" binding:"required"`
	Password string  `json:"password" binding:"required"`
	Port     *uint16 `json:"port" binding:"omitempty,gte=1,lte=65535"` // default is 22
	DeviceId *string `json:"deviceId" binding:"uuid"`        // if deviceId is empty, treat as create credential for all devices, else create credential for specified device
}

type CliCredentialUpdate struct {
	Username *string `json:"username" binding:"omitempty"`
	Password *string `json:"password" binding:"omitempty"`
	Port     *uint16 `json:"port" binding:"omitempty,gte=1,lte=65535"`
}

type CliCredential struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Port     uint16 `json:"port"`
}

type SnmpV2CredentialCreate struct {
	Community      string  `json:"community" binding:"required"`
	MaxRepetitions *uint8  `json:"maxRepetitions" binding:"omitempty, gte=10,lte=200"` // default is 50
	Timeout        *uint8  `json:"timeout" binding:"omitempty,gte=1,lte=30"`           // default is 5
	Port           *uint16 `json:"port" binding:"omitempty,gte=1,lte=65535"`           // default is 161
	DeviceId       *string `json:"deviceId" binding:"uuid"`                  // if deviceId is empty, treat as create credential for all devices, else create credential for specified device
}

func (d *SnmpV2CredentialCreate) SetDefaultValue() {
	if d.Community == "" {
		d.Community = "public"
	}

	if d.MaxRepetitions == nil {
		d.MaxRepetitions = new(uint8)
		*d.MaxRepetitions = 50
	}

	if d.Timeout == nil {
		d.Timeout = new(uint8)
		*d.Timeout = 5
	}

	if d.Port == nil {
		d.Port = new(uint16)
		*d.Port = 161
	}
}

type SnmpV2CredentialUpdate struct {
	Community      *string `json:"community" binding:"omitempty"`
	MaxRepetitions *uint8  `json:"maxRepetitions" binding:"omitempty, gte=10,lte=200"`
	Timeout        *uint8  `json:"timeout" binding:"omitempty,gte=1,lte=30"`
	Port           *uint16 `json:"port" binding:"omitempty,gte=1,lte=65535"`
}

type SnmpV2Credential struct {
	Community      string `json:"community"`
	MaxRepetitions uint8  `json:"maxRepetitions"`
	Timeout        uint8  `json:"timeout"`
	Port           uint16 `json:"port"`
}

type RestconfCredentialCreate struct {
	Url      string  `json:"url" binding:"required,http_url"`
	Username string  `json:"username" binding:"required"`
	Password string  `json:"password" binding:"required"`
	DeviceId string `json:"deviceId" binding:"uuid"`
}

type RestconfCredentialUpdate struct {
	Url      *string `json:"url" binding:"omitempty,http_url"`
	Username *string `json:"username" binding:"omitempty"`
	Password *string `json:"password" binding:"omitempty"`
}

type RestconfCredential struct {
	Url      string `json:"url"`
	Username string `json:"username"`
	Password string `json:"password"`
}