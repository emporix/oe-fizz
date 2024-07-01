package openapi

func Clone(o *OpenAPI) (OpenAPI, error) {
	clone := OpenAPI{
		OpenAPI:       o.OpenAPI,
		Info:          cloneInfo(o.Info),
		Servers:       cloneServers(o.Servers),
		Paths:         clonePaths(o.Paths),
		Components:    cloneComponents(o.Components),
		Tags:          cloneTags(o.Tags),
		Security:      cloneSecurity(o.Security),
		XTagGroups:    cloneXTagGroups(o.XTagGroups),
		IncludeLabels: append([]string{}, o.IncludeLabels...),
	}
	return clone, nil
}

func cloneInfo(info *Info) *Info {
	if info == nil {
		return nil
	}
	return &Info{
		Title:          info.Title,
		Description:    info.Description,
		TermsOfService: info.TermsOfService,
		Contact:        cloneContact(info.Contact),
		License:        cloneLicense(info.License),
		Version:        info.Version,
		XLogo:          cloneXLogo(info.XLogo),
	}
}

func cloneContact(contact *Contact) *Contact {
	if contact == nil {
		return nil
	}
	return &Contact{
		Name:  contact.Name,
		URL:   contact.URL,
		Email: contact.Email,
	}
}

func cloneLicense(license *License) *License {
	if license == nil {
		return nil
	}
	return &License{
		Name: license.Name,
		URL:  license.URL,
	}
}

func cloneXLogo(xlogo *XLogo) *XLogo {
	if xlogo == nil {
		return nil
	}
	return &XLogo{
		URL:             xlogo.URL,
		BackgroundColor: xlogo.BackgroundColor,
		AltText:         xlogo.AltText,
		Href:            xlogo.Href,
	}
}

func cloneServers(servers []*Server) []*Server {
	if servers == nil {
		return nil
	}
	clonedServers := make([]*Server, len(servers))
	for i, server := range servers {
		clonedServers[i] = &Server{
			URL:         server.URL,
			Description: server.Description,
			Variables:   cloneServerVariables(server.Variables),
		}
	}
	return clonedServers
}

func cloneServerVariables(vars map[string]*ServerVariable) map[string]*ServerVariable {
	if vars == nil {
		return nil
	}
	clonedVars := make(map[string]*ServerVariable)
	for k, v := range vars {
		clonedVars[k] = &ServerVariable{
			Enum:        append([]string{}, v.Enum...),
			Default:     v.Default,
			Description: v.Description,
		}
	}
	return clonedVars
}

func clonePaths(paths Paths) Paths {
	if paths == nil {
		return nil
	}
	clonedPaths := make(Paths)
	for k, v := range paths {
		clonedPaths[k] = clonePathItem(v)
	}
	return clonedPaths
}

func clonePathItem(p *PathItem) *PathItem {
	if p == nil {
		return nil
	}
	return &PathItem{
		Ref:         p.Ref,
		Summary:     p.Summary,
		Description: p.Description,
		GET:         cloneOperation(p.GET),
		PUT:         cloneOperation(p.PUT),
		POST:        cloneOperation(p.POST),
		DELETE:      cloneOperation(p.DELETE),
		OPTIONS:     cloneOperation(p.OPTIONS),
		HEAD:        cloneOperation(p.HEAD),
		PATCH:       cloneOperation(p.PATCH),
		TRACE:       cloneOperation(p.TRACE),
		Servers:     cloneServers(p.Servers),
		Parameters:  cloneParameters(p.Parameters),
	}
}

func cloneOperation(op *Operation) *Operation {
	if op == nil {
		return nil
	}
	return &Operation{
		Tags:         cloneStringSlice(op.Tags),
		Summary:      op.Summary,
		Description:  op.Description,
		ID:           op.ID,
		Parameters:   cloneParameters(op.Parameters),
		RequestBody:  cloneRequestBody(op.RequestBody),
		Responses:    cloneResponses(op.Responses),
		Deprecated:   op.Deprecated,
		Servers:      cloneServers(op.Servers),
		Security:     cloneSecurity(op.Security),
		XCodeSamples: cloneXCodeSamples(op.XCodeSamples),
		XInternal:    op.XInternal,
		Labels:       cloneStringSlice(op.Labels),
	}
}

func cloneStringSlice(slice []string) []string {
	if slice == nil {
		return nil
	}
	return append([]string{}, slice...)
}

func cloneXCodeSamples(samples []*XCodeSample) []*XCodeSample {
	if samples == nil {
		return nil
	}
	clonedSamples := make([]*XCodeSample, len(samples))
	for i, sample := range samples {
		clonedSamples[i] = &XCodeSample{
			Lang:   sample.Lang,
			Label:  sample.Label,
			Source: sample.Source,
		}
	}
	return clonedSamples
}

func cloneParameters(params []*ParameterOrRef) []*ParameterOrRef {
	if params == nil {
		return nil
	}
	clonedParams := make([]*ParameterOrRef, len(params))
	for i, p := range params {
		clonedParams[i] = &ParameterOrRef{
			Parameter: cloneParameter(p.Parameter),
			Reference: cloneReference(p.Reference),
		}
	}
	return clonedParams
}

func cloneParameter(p *Parameter) *Parameter {
	if p == nil {
		return nil
	}
	return &Parameter{
		Name:            p.Name,
		In:              p.In,
		Description:     p.Description,
		Required:        p.Required,
		Deprecated:      p.Deprecated,
		AllowEmptyValue: p.AllowEmptyValue,
		Schema:          cloneSchemaOrRef(p.Schema),
		Style:           p.Style,
		Explode:         p.Explode,
	}
}

func cloneReference(ref *Reference) *Reference {
	if ref == nil {
		return nil
	}
	return &Reference{
		Ref: ref.Ref,
	}
}

func cloneRequestBody(rb *RequestBody) *RequestBody {
	if rb == nil {
		return nil
	}
	return &RequestBody{
		Description: rb.Description,
		Content:     cloneContentWithoutRef(rb.Content),
		Required:    rb.Required,
	}
}

func cloneContentWithoutRef(content map[string]*MediaType) map[string]*MediaType {
	if content == nil {
		return nil
	}
	clonedContent := make(map[string]*MediaType)
	for k, v := range content {
		clonedContent[k] = cloneMediaType(v)
	}
	return clonedContent
}

func cloneRequestBodyContent(content map[string]*MediaType) map[string]*MediaType {
	if content == nil {
		return nil
	}
	clonedContent := make(map[string]*MediaType)
	for k, v := range content {
		clonedContent[k] = cloneMediaType(v)
	}
	return clonedContent
}

func cloneContent(content map[string]*MediaTypeOrRef) map[string]*MediaTypeOrRef {
	if content == nil {
		return nil
	}
	clonedContent := make(map[string]*MediaTypeOrRef)
	for k, v := range content {
		clonedContent[k] = cloneMediaTypeOrRef(v)
	}
	return clonedContent
}

func cloneMediaTypeOrRef(mtor *MediaTypeOrRef) *MediaTypeOrRef {
	if mtor == nil {
		return nil
	}
	return &MediaTypeOrRef{
		MediaType: cloneMediaType(mtor.MediaType),
		Reference: cloneReference(mtor.Reference),
	}
}

func cloneMediaType(mt *MediaType) *MediaType {
	if mt == nil {
		return nil
	}
	return &MediaType{
		Schema:   cloneSchemaOrRef(mt.Schema),
		Example:  mt.Example,
		Examples: cloneExamples(mt.Examples),
		Encoding: cloneEncoding(mt.Encoding),
	}
}

func cloneExamples(examples map[string]*ExampleOrRef) map[string]*ExampleOrRef {
	if examples == nil {
		return nil
	}
	clonedExamples := make(map[string]*ExampleOrRef)
	for k, v := range examples {
		clonedExamples[k] = &ExampleOrRef{
			Example:   cloneExample(v.Example),
			Reference: cloneReference(v.Reference),
		}
	}
	return clonedExamples
}

func cloneExample(ex *Example) *Example {
	if ex == nil {
		return nil
	}
	return &Example{
		Summary:       ex.Summary,
		Description:   ex.Description,
		Value:         ex.Value,
		ExternalValue: ex.ExternalValue,
	}
}

func cloneEncoding(enc map[string]*Encoding) map[string]*Encoding {
	if enc == nil {
		return nil
	}
	clonedEncoding := make(map[string]*Encoding)
	for k, v := range enc {
		clonedEncoding[k] = &Encoding{
			ContentType:   v.ContentType,
			Headers:       cloneHeaders(v.Headers),
			Style:         v.Style,
			Explode:       v.Explode,
			AllowReserved: v.AllowReserved,
		}
	}
	return clonedEncoding
}

func cloneHeaders(headers map[string]*HeaderOrRef) map[string]*HeaderOrRef {
	if headers == nil {
		return nil
	}
	clonedHeaders := make(map[string]*HeaderOrRef)
	for k, v := range headers {
		clonedHeaders[k] = &HeaderOrRef{
			Header:    cloneHeader(v.Header),
			Reference: cloneReference(v.Reference),
		}
	}
	return clonedHeaders
}

func cloneHeader(h *Header) *Header {
	if h == nil {
		return nil
	}
	return &Header{
		Description:     h.Description,
		Required:        h.Required,
		Deprecated:      h.Deprecated,
		AllowEmptyValue: h.AllowEmptyValue,
		Schema:          cloneSchemaOrRef(h.Schema),
	}
}

func cloneSchemaOrRef(sor *SchemaOrRef) *SchemaOrRef {
	if sor == nil {
		return nil
	}
	return &SchemaOrRef{
		Schema:    cloneSchema(sor.Schema),
		Reference: cloneReference(sor.Reference),
	}
}

func cloneSchema(s *Schema) *Schema {
	if s == nil {
		return nil
	}
	return &Schema{
		Type:                 s.Type,
		AllOf:                cloneSchemaOrRef(s.AllOf),
		OneOf:                cloneSchemaOrRef(s.OneOf),
		AnyOf:                cloneSchemaOrRef(s.AnyOf),
		Items:                cloneSchemaOrRef(s.Items),
		Properties:           cloneSchemaProperties(s.Properties),
		AdditionalProperties: cloneSchemaOrRef(s.AdditionalProperties),
		Description:          s.Description,
		Format:               s.Format,
		Default:              s.Default,
		Example:              s.Example,
		Title:                s.Title,
		MultipleOf:           s.MultipleOf,
		Maximum:              s.Maximum,
		ExclusiveMaximum:     s.ExclusiveMaximum,
		Minimum:              s.Minimum,
		ExclusiveMinimum:     s.ExclusiveMinimum,
		MaxLength:            s.MaxLength,
		MinLength:            s.MinLength,
		Pattern:              s.Pattern,
		MaxItems:             s.MaxItems,
		MinItems:             s.MinItems,
		UniqueItems:          s.UniqueItems,
		MaxProperties:        s.MaxProperties,
		MinProperties:        s.MinProperties,
		Required:             cloneStringSlice(s.Required),
		Enum:                 cloneEnum(s.Enum),
		Nullable:             s.Nullable,
		Deprecated:           s.Deprecated,
	}
}

func cloneSchemaProperties(properties map[string]*SchemaOrRef) map[string]*SchemaOrRef {
	if properties == nil {
		return nil
	}
	clonedProperties := make(map[string]*SchemaOrRef)
	for k, v := range properties {
		clonedProperties[k] = cloneSchemaOrRef(v)
	}
	return clonedProperties
}

func cloneEnum(enum []interface{}) []interface{} {
	if enum == nil {
		return nil
	}
	clonedEnum := make([]interface{}, len(enum))
	copy(clonedEnum, enum)
	return clonedEnum
}

func cloneResponses(responses Responses) Responses {
	if responses == nil {
		return nil
	}
	clonedResponses := make(Responses)
	for k, v := range responses {
		clonedResponses[k] = cloneResponseOrRef(v)
	}
	return clonedResponses
}

func cloneResponseOrRef(ror *ResponseOrRef) *ResponseOrRef {
	if ror == nil {
		return nil
	}
	return &ResponseOrRef{
		Response:  cloneResponse(ror.Response),
		Reference: cloneReference(ror.Reference),
	}
}

func cloneResponse(r *Response) *Response {
	if r == nil {
		return nil
	}
	return &Response{
		Description: r.Description,
		Headers:     cloneHeaders(r.Headers),
		Content:     cloneContent(r.Content),
	}
}

func cloneComponents(c *Components) *Components {
	if c == nil {
		return nil
	}
	return &Components{
		Schemas:         cloneSchemas(c.Schemas),
		Responses:       cloneResponsesMap(c.Responses),
		Parameters:      cloneParametersMap(c.Parameters),
		Examples:        cloneExamplesMap(c.Examples),
		Headers:         cloneHeadersMap(c.Headers),
		SecuritySchemes: cloneSecuritySchemes(c.SecuritySchemes),
	}
}

func cloneSchemas(schemas map[string]*SchemaOrRef) map[string]*SchemaOrRef {
	if schemas == nil {
		return nil
	}
	clonedSchemas := make(map[string]*SchemaOrRef)
	for k, v := range schemas {
		clonedSchemas[k] = cloneSchemaOrRef(v)
	}
	return clonedSchemas
}

func cloneResponsesMap(responses map[string]*ResponseOrRef) map[string]*ResponseOrRef {
	if responses == nil {
		return nil
	}
	clonedResponses := make(map[string]*ResponseOrRef)
	for k, v := range responses {
		clonedResponses[k] = cloneResponseOrRef(v)
	}
	return clonedResponses
}

func cloneParametersMap(parameters map[string]*ParameterOrRef) map[string]*ParameterOrRef {
	if parameters == nil {
		return nil
	}
	clonedParameters := make(map[string]*ParameterOrRef)
	for k, v := range parameters {
		clonedParameters[k] = cloneParameterOrRef(v)
	}
	return clonedParameters
}

func cloneParameterOrRef(por *ParameterOrRef) *ParameterOrRef {
	if por == nil {
		return nil
	}
	return &ParameterOrRef{
		Parameter: cloneParameter(por.Parameter),
		Reference: cloneReference(por.Reference),
	}
}

func cloneExamplesMap(examples map[string]*ExampleOrRef) map[string]*ExampleOrRef {
	if examples == nil {
		return nil
	}
	clonedExamples := make(map[string]*ExampleOrRef)
	for k, v := range examples {
		clonedExamples[k] = cloneExampleOrRef(v)
	}
	return clonedExamples
}

func cloneExampleOrRef(eor *ExampleOrRef) *ExampleOrRef {
	if eor == nil {
		return nil
	}
	return &ExampleOrRef{
		Example:   cloneExample(eor.Example),
		Reference: cloneReference(eor.Reference),
	}
}

func cloneHeadersMap(headers map[string]*HeaderOrRef) map[string]*HeaderOrRef {
	if headers == nil {
		return nil
	}
	clonedHeaders := make(map[string]*HeaderOrRef)
	for k, v := range headers {
		clonedHeaders[k] = cloneHeaderOrRef(v)
	}
	return clonedHeaders
}

func cloneHeaderOrRef(hor *HeaderOrRef) *HeaderOrRef {
	if hor == nil {
		return nil
	}
	return &HeaderOrRef{
		Header:    cloneHeader(hor.Header),
		Reference: cloneReference(hor.Reference),
	}
}

func cloneSecuritySchemes(schemes map[string]*SecuritySchemeOrRef) map[string]*SecuritySchemeOrRef {
	if schemes == nil {
		return nil
	}
	clonedSchemes := make(map[string]*SecuritySchemeOrRef)
	for k, v := range schemes {
		clonedSchemes[k] = cloneSecuritySchemeOrRef(v)
	}
	return clonedSchemes
}

func cloneSecuritySchemeOrRef(sor *SecuritySchemeOrRef) *SecuritySchemeOrRef {
	if sor == nil {
		return nil
	}
	return &SecuritySchemeOrRef{
		SecurityScheme: cloneSecurityScheme(sor.SecurityScheme),
		Reference:      cloneReference(sor.Reference),
	}
}

func cloneSecurityScheme(ss *SecurityScheme) *SecurityScheme {
	if ss == nil {
		return nil
	}
	return &SecurityScheme{
		Type:             ss.Type,
		Scheme:           ss.Scheme,
		BearerFormat:     ss.BearerFormat,
		Description:      ss.Description,
		In:               ss.In,
		Name:             ss.Name,
		OpenIDConnectURL: ss.OpenIDConnectURL,
		Flows:            cloneOAuthFlows(ss.Flows),
	}
}

func cloneOAuthFlows(of *OAuthFlows) *OAuthFlows {
	if of == nil {
		return nil
	}
	return &OAuthFlows{
		Implicit:          cloneOAuthFlow(of.Implicit),
		Password:          cloneOAuthFlow(of.Password),
		ClientCredentials: cloneOAuthFlow(of.ClientCredentials),
		AuthorizationCode: cloneOAuthFlow(of.AuthorizationCode),
	}
}

func cloneOAuthFlow(of *OAuthFlow) *OAuthFlow {
	if of == nil {
		return nil
	}
	return &OAuthFlow{
		AuthorizationURL: of.AuthorizationURL,
		TokenURL:         of.TokenURL,
		RefreshURL:       of.RefreshURL,
		Scopes:           cloneStringMap(of.Scopes),
	}
}

func cloneStringMap(m map[string]string) map[string]string {
	if m == nil {
		return nil
	}
	clonedMap := make(map[string]string)
	for k, v := range m {
		clonedMap[k] = v
	}
	return clonedMap
}

func cloneTags(tags []*Tag) []*Tag {
	if tags == nil {
		return nil
	}
	clonedTags := make([]*Tag, len(tags))
	for i, t := range tags {
		clonedTags[i] = cloneTag(t)
	}
	return clonedTags
}

func cloneTag(t *Tag) *Tag {
	if t == nil {
		return nil
	}
	return &Tag{
		Name:        t.Name,
		Description: t.Description,
	}
}

func cloneSecurity(sec []*SecurityRequirement) []*SecurityRequirement {
	if sec == nil {
		return nil
	}
	clonedSec := make([]*SecurityRequirement, len(sec))
	for i, s := range sec {
		clonedSec[i] = cloneSecurityRequirement(s)
	}
	return clonedSec
}

func cloneSecurityRequirement(sr *SecurityRequirement) *SecurityRequirement {
	if sr == nil {
		return nil
	}
	clonedSR := make(SecurityRequirement)
	for k, v := range *sr {
		clonedSR[k] = cloneStringSlice(v)
	}
	return &clonedSR
}

func cloneXTagGroups(xtg []*XTagGroup) []*XTagGroup {
	if xtg == nil {
		return nil
	}
	clonedXTG := make([]*XTagGroup, len(xtg))
	for i, x := range xtg {
		clonedXTG[i] = cloneXTagGroup(x)
	}
	return clonedXTG
}

func cloneXTagGroup(x *XTagGroup) *XTagGroup {
	if x == nil {
		return nil
	}
	return &XTagGroup{
		Name: x.Name,
		Tags: append([]string{}, x.Tags...),
	}
}
