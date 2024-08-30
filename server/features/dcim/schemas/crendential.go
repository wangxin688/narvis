package schemas

type DeviceCliCredentialCreate struct {
	Username string  `json:"username" binding:"required"`
	Password string  `json:"password" binding:"required"`
	Port     *uint16 `json:"port" binding:"omitempty,gte=1,lte=65535"` // default is 22
	DeviceId *string `json:"deviceId" binding:"omitempty,uuid"`       // if deviceId is empty, treat as create credential for all devices, else create credential for specified device
}

type DeviceCliCredentialUpdate struct {
	Username *string `json:"username" binding:"omitempty"`
	Password *string `json:"password" binding:"omitempty"`
	Port     *uint16 `json:"port" binding:"omitempty,gte=1,lte=65535"`
	DeviceId *string `json:"deviceId" binding:"omitempty,uuid"`
}

type DeviceCliCredential struct {
	Username string  `json:"username"`
	Password string  `json:"password"`
	Port     *uint16 `json:"port"`
}

type DeviceSnmpV2CredentialCreate struct {
	Community      string  `json:"community" binding:"required"`
	MaxRepetitions *uint8  `json:"maxRepetitions" binding:"omitempty, gte=10,lte=200"` // default is 50
	Timeout        *uint8  `json:"timeout" binding:"omitempty,gte=1,lte=30"`            // default is 5
	Port           *uint16 `json:"port" binding:"omitempty,gte=1,lte=65535"`            // default is 161
	DeviceId       *string `json:"deviceId" binding:"omitempty,uuid"`                  // if deviceId is empty, treat as create credential for all devices, else create credential for specified device
}

type DeviceSnmpV2CredentialUpdate struct {
	Community      *string `json:"community" binding:"omitempty"`
	MaxRepetitions *uint8  `json:"maxRepetitions" binding:"omitempty, gte=10,lte=200"`
	Timeout        *uint8  `json:"timeout" binding:"omitempty,gte=1,lte=30"`
	Port           *uint16 `json:"port" binding:"omitempty,gte=1,lte=65535"`
	DeviceId       *string `json:"deviceId" binding:"omitempty,uuid"`
}

type DeviceSnmpV2Credential struct {
	Community      string  `json:"community"`
	MaxRepetitions *uint8  `json:"maxRepetitions"`
	Timeout        *uint8  `json:"timeout"`
	Port           *uint16 `json:"port"`
}

type DeviceRestconfCredentialCreate struct {
	Url      string  `json:"url" binding:"required,http_url"`
	Username string  `json:"username" binding:"required"`
	Password string  `json:"password" binding:"required"`
	DeviceId *string `json:"deviceId" binding:"omitempty,uuid"` // if deviceId is empty, treat as create credential for all devices, else create credential for specified device
}

type DeviceRestconfCredentialUpdate struct {
	Url      *string `json:"url" binding:"omitempty,http_url"`
	Username *string `json:"username" binding:"omitempty"`
	Password *string `json:"password" binding:"omitempty"`
	DeviceId *string `json:"deviceId" binding:"omitempty,uuid"`
}
