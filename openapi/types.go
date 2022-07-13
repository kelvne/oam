package openapi

type (
	Definition struct {
		OpenAPI      string                `yaml:"openapi" json:"openapi"`
		Info         *Info                 `yaml:"info" json:"info"`
		Servers      []*Server             `yaml:"servers,omitempty" json:"servers,omitempty"`
		Paths        Paths                 `yaml:"paths" json:"paths"`
		Components   Components            `yaml:"components,omitempty" json:"components,omitempty"`
		Security     []map[string][]string `yaml:"security,omitempty" json:"security,omitempty"`
		Tags         []*Tag                `yaml:"tags,omitempty" json:"tags,omitempty"`
		ExternalDocs *ExternalDocs         `yaml:"externalDocs,omitempty" json:"externalDocs,omitempty"`
	}

	Info struct {
		Title          string   `yaml:"title" json:"title"`
		Description    string   `yaml:"description,omitempty" json:"description,omitempty"`
		TermsOfService string   `yaml:"termsOfService,omitempty" json:"termsOfService,omitempty"`
		Contact        *Contact `yaml:"contact,omitempty" json:"contact,omitempty"`
		License        *License `yaml:"license,omitempty" json:"license,omitempty"`
		Version        string   `yaml:"version" json:"version"`
	}

	Contact struct {
		Name  string `yaml:"name,omitempty" json:"name,omitempty"`
		Url   string `yaml:"url,omitempty" json:"url,omitempty"`
		Email string `yaml:"email,omitempty" json:"email,omitempty"`
	}

	License struct {
		Name string `yaml:"name" json:"name"`
		Url  string `yaml:"url,omitempty" json:"url,omitempty"`
	}

	Server struct {
		Url         string                     `yaml:"url" json:"url"`
		Description string                     `yaml:"description,omitempty" json:"description,omitempty"`
		Variables   map[string]*ServerVariable `yaml:"variables,omitempty" json:"variables,omitempty"`
	}

	ServerVariable struct {
		Enum        []string `yaml:"enum,omitempty" json:"enum,omitempty"`
		Default     string   `yaml:"default" json:"default"`
		Description string   `yaml:"description,omitempty" json:"description,omitempty"`
	}

	Paths map[string]*Path

	Path struct {
		Ref         string           `yaml:"$ref,omitempty" json:"$ref,omitempty"`
		Summary     string           `yaml:"summary,omitempty" json:"summary,omitempty"`
		Description string           `yaml:"description,omitempty" json:"description,omitempty"`
		Get         *PathOperation   `yaml:"get,omitempty" json:"get,omitempty"`
		Put         *PathOperation   `yaml:"put,omitempty" json:"put,omitempty"`
		Post        *PathOperation   `yaml:"post,omitempty" json:"post,omitempty"`
		Delete      *PathOperation   `yaml:"delete,omitempty" json:"delete,omitempty"`
		Options     *PathOperation   `yaml:"options,omitempty" json:"options,omitempty"`
		Head        *PathOperation   `yaml:"head,omitempty" json:"head,omitempty"`
		Patch       *PathOperation   `yaml:"patch,omitempty" json:"patch,omitempty"`
		Trace       *PathOperation   `yaml:"trace,omitempty" json:"trace,omitempty"`
		Servers     []*Server        `yaml:"servers,omitempty" json:"servers,omitempty"`
		Parameters  []*PathParameter `yaml:"parameters,omitempty" json:"parameters,omitempty"`
	}

	PathOperation struct {
		Tags         []string               `yaml:"tags,omitempty" json:"tags,omitempty"`
		Summary      string                 `yaml:"summary,omitempty" json:"summary,omitempty"`
		Description  string                 `yaml:"description,omitempty" json:"description,omitempty"`
		ExternalDocs *ExternalDocs          `yaml:"externalDocs,omitempty" json:"externalDocs,omitempty"`
		OperationId  string                 `yaml:"operationId,omitempty" json:"operationId,omitempty"`
		Parameters   []*PathParameter       `yaml:"parameters,omitempty" json:"parameters,omitempty"`
		RequestBody  *RequestBody           `yaml:"requestBody,omitempty" json:"requestBody,omitempty"`
		Responses    Responses              `yaml:"responses" json:"responses"`
		Callbacks    map[string]interface{} `yaml:"callbacks,omitempty" json:"callbacks,omitempty"`
		Deprecated   bool                   `yaml:"deprecated,omitempty" json:"deprecated,omitempty"`
		Security     []map[string][]string  `yaml:"security,omitempty" json:"security,omitempty"`
		Servers      []*Server              `yaml:"servers,omitempty" json:"servers,omitempty"`
	}

	RequestBody struct {
		Description string                `yaml:"description,omitempty" json:"description,omitempty"`
		Content     map[string]*MediaType `yaml:"content,omitempty" json:"content,omitempty"`
		Required    bool                  `yaml:"required,omitempty" json:"required,omitempty"`
	}

	MediaType struct {
		Schema   Schema               `yaml:"schema,omitempty" json:"schema,omitempty"`
		Example  interface{}          `yaml:"example,omitempty" json:"example,omitempty"`
		Examples map[string]*Example  `yaml:"examples,omitempty" json:"examples,omitempty"`
		Encoding map[string]*Encoding `yaml:"encoding,omitempty" json:"encoding,omitempty"`
	}

	Encoding struct {
		ContentType   string             `yaml:"contentType,omitempty" json:"contentType,omitempty"`
		Headers       map[string]*Header `yaml:"headers,omitempty" json:"headers,omitempty"`
		Style         string             `yaml:"style,omitempty" json:"style,omitempty"`
		Explode       bool               `yaml:"explode,omitempty" json:"explode,omitempty"`
		AllowReserved bool               `yaml:"allowReserved,omitempty" json:"allowReserved,omitempty"`
	}

	Example struct {
		Ref           string      `yaml:"$ref,omitempty" json:"$ref,omitempty"`
		Summary       string      `yaml:"summary,omitempty" json:"summary,omitempty"`
		Description   string      `yaml:"description,omitempty" json:"description,omitempty"`
		Value         interface{} `yaml:"value,omitempty" json:"value,omitempty"`
		ExternalValue string      `yaml:"externalValue,omitempty" json:"externalValue,omitempty"`
	}

	PathParameter struct {
		Ref             string              `yaml:"$ref,omitempty" json:"$ref,omitempty"`
		Name            string              `yaml:"name,omitempty" json:"name,omitempty"`
		In              string              `yaml:"in,omitempty" json:"in,omitempty"`
		Description     string              `yaml:"description,omitempty" json:"description,omitempty"`
		Required        bool                `yaml:"required,omitempty" json:"required,omitempty"`
		Deprecated      bool                `yaml:"deprecated,omitempty" json:"deprecated,omitempty"`
		AllowEmptyValue bool                `yaml:"allowEmptyValue,omitempty" json:"allowEmptyValue,omitempty"`
		Style           string              `yaml:"style,omitempty" json:"style,omitempty"`
		Schema          Schema              `yaml:"schema,omitempty" json:"schema,omitempty"`
		Example         interface{}         `yaml:"example,omitempty" json:"example,omitempty"`
		Examples        map[string]*Example `yaml:"examples,omitempty" json:"examples,omitempty"`
	}

	Header struct {
		Ref             string `yaml:"$ref,omitempty" json:"$ref,omitempty"`
		Description     string `yaml:"description,omitempty" json:"description,omitempty"`
		Required        bool   `yaml:"required,omitempty" json:"required,omitempty"`
		Deprecated      bool   `yaml:"deprecated,omitempty" json:"deprecated,omitempty"`
		AllowEmptyValue bool   `yaml:"allowEmptyValue,omitempty" json:"allowEmptyValue,omitempty"`
	}

	Response struct {
		Ref         string                 `yaml:"$ref,omitempty" json:"$ref,omitempty"`
		Description string                 `yaml:"description,omitempty" json:"description,omitempty"`
		Headers     map[string]*Header     `yaml:"headers,omitempty" json:"headers,omitempty"`
		Content     map[string]*MediaType  `yaml:"content,omitempty" json:"content,omitempty"`
		Links       map[string]interface{} `yaml:"links,omitempty" json:"links,omitempty"`
	}

	Components struct {
		Schemas         Schemas                    `yaml:"schemas,omitempty" json:"schemas,omitempty"`
		Responses       Responses                  `yaml:"responses,omitempty" json:"responses,omitempty"`
		Parameters      map[string]*PathParameter  `yaml:"parameters,omitempty" json:"parameters,omitempty"`
		Examples        map[string]*Example        `yaml:"examples,omitempty" json:"examples,omitempty"`
		RequestBodies   map[string]*RequestBody    `yaml:"requestBodies,omitempty" json:"requestBodies,omitempty"`
		Headers         map[string]*Header         `yaml:"headers,omitempty" json:"headers,omitempty"`
		SecuritySchemes map[string]*SecurityScheme `yaml:"securitySchemes,omitempty" json:"securitySchemes,omitempty"`
		Links           map[string]interface{}     `yaml:"links,omitempty" json:"links,omitempty"`
		Callbacks       map[string]interface{}     `yaml:"callbacks,omitempty" json:"callbacks,omitempty"`
	}

	SecurityScheme struct {
		Type             string                 `yaml:"type,omitempty" json:"type,omitempty"`
		Description      string                 `yaml:"description,omitempty" json:"description,omitempty"`
		Name             string                 `yaml:"name,omitempty" json:"name,omitempty"`
		In               string                 `yaml:"in,omitempty" json:"in,omitempty"`
		Scheme           string                 `yaml:"scheme,omitempty" json:"scheme,omitempty"`
		BearerFormat     string                 `yaml:"bearerFormat,omitempty" json:"bearerFormat,omitempty"`
		Flows            map[string]interface{} `yaml:"flows,omitempty" json:"flows,omitempty"`
		OpenIdConnectUrl string                 `yaml:"openIdConnectUrl,omitempty" json:"openIdConnectUrl,omitempty"`
	}

	Tag struct {
		Name         string        `yaml:"name" json:"name"`
		Description  string        `yaml:"description,omitempty" json:"description,omitempty"`
		ExternalDocs *ExternalDocs `yaml:"externalDocs,omitempty" json:"externalDocs,omitempty"`
	}

	ExternalDocs struct {
		Description string `yaml:"description,omitempty" json:"description,omitempty"`
		Url         string `yaml:"url" json:"url"`
	}

	Responses map[string]*Response

	Schemas          map[string]Schema
	Schema           map[string]interface{}
	SchemaProperties map[string]interface{}
	SchemaProperty   map[string]interface{}
)
