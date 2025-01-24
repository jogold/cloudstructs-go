// High-level constructs for AWS CDK
package cloudstructs

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterClass(
		"cloudstructs.CodeCommitMirror",
		reflect.TypeOf((*CodeCommitMirror)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_CodeCommitMirror{}
			_jsii_.InitJsiiProxy(&j.Type__constructsConstruct)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"cloudstructs.CodeCommitMirrorProps",
		reflect.TypeOf((*CodeCommitMirrorProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"cloudstructs.CodeCommitMirrorSourceRepository",
		reflect.TypeOf((*CodeCommitMirrorSourceRepository)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "plainTextUrl", GoGetter: "PlainTextUrl"},
			_jsii_.MemberProperty{JsiiProperty: "secretUrl", GoGetter: "SecretUrl"},
		},
		func() interface{} {
			return &jsiiProxy_CodeCommitMirrorSourceRepository{}
		},
	)
	_jsii_.RegisterEnum(
		"cloudstructs.DmarcAlignment",
		reflect.TypeOf((*DmarcAlignment)(nil)).Elem(),
		map[string]interface{}{
			"RELAXED": DmarcAlignment_RELAXED,
			"STRICT": DmarcAlignment_STRICT,
		},
	)
	_jsii_.RegisterEnum(
		"cloudstructs.DmarcPolicy",
		reflect.TypeOf((*DmarcPolicy)(nil)).Elem(),
		map[string]interface{}{
			"NONE": DmarcPolicy_NONE,
			"QUARANTINE": DmarcPolicy_QUARANTINE,
			"REJECT": DmarcPolicy_REJECT,
		},
	)
	_jsii_.RegisterClass(
		"cloudstructs.DmarcReporter",
		reflect.TypeOf((*DmarcReporter)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_DmarcReporter{}
			_jsii_.InitJsiiProxy(&j.Type__constructsConstruct)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"cloudstructs.DmarcReporterProps",
		reflect.TypeOf((*DmarcReporterProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"cloudstructs.EcsServiceRoller",
		reflect.TypeOf((*EcsServiceRoller)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_EcsServiceRoller{}
			_jsii_.InitJsiiProxy(&j.Type__constructsConstruct)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"cloudstructs.EcsServiceRollerProps",
		reflect.TypeOf((*EcsServiceRollerProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"cloudstructs.EmailReceiver",
		reflect.TypeOf((*EmailReceiver)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "bucket", GoGetter: "Bucket"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "topic", GoGetter: "Topic"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_EmailReceiver{}
			_jsii_.InitJsiiProxy(&j.Type__constructsConstruct)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"cloudstructs.EmailReceiverProps",
		reflect.TypeOf((*EmailReceiverProps)(nil)).Elem(),
	)
	_jsii_.RegisterInterface(
		"cloudstructs.IStateMachine",
		reflect.TypeOf((*IStateMachine)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "stateMachineArn", GoGetter: "StateMachineArn"},
		},
		func() interface{} {
			return &jsiiProxy_IStateMachine{}
		},
	)
	_jsii_.RegisterClass(
		"cloudstructs.MjmlTemplate",
		reflect.TypeOf((*MjmlTemplate)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "templateName", GoGetter: "TemplateName"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_MjmlTemplate{}
			_jsii_.InitJsiiProxy(&j.Type__constructsConstruct)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"cloudstructs.MjmlTemplateProps",
		reflect.TypeOf((*MjmlTemplateProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"cloudstructs.RollTrigger",
		reflect.TypeOf((*RollTrigger)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "rule", GoGetter: "Rule"},
			_jsii_.MemberProperty{JsiiProperty: "schedule", GoGetter: "Schedule"},
		},
		func() interface{} {
			return &jsiiProxy_RollTrigger{}
		},
	)
	_jsii_.RegisterClass(
		"cloudstructs.SamlFederatedPrincipal",
		reflect.TypeOf((*SamlFederatedPrincipal)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addToAssumeRolePolicy", GoMethod: "AddToAssumeRolePolicy"},
			_jsii_.MemberMethod{JsiiMethod: "addToPolicy", GoMethod: "AddToPolicy"},
			_jsii_.MemberMethod{JsiiMethod: "addToPrincipalPolicy", GoMethod: "AddToPrincipalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "assumeRoleAction", GoGetter: "AssumeRoleAction"},
			_jsii_.MemberProperty{JsiiProperty: "conditions", GoGetter: "Conditions"},
			_jsii_.MemberMethod{JsiiMethod: "dedupeString", GoMethod: "DedupeString"},
			_jsii_.MemberProperty{JsiiProperty: "federated", GoGetter: "Federated"},
			_jsii_.MemberProperty{JsiiProperty: "grantPrincipal", GoGetter: "GrantPrincipal"},
			_jsii_.MemberProperty{JsiiProperty: "policyFragment", GoGetter: "PolicyFragment"},
			_jsii_.MemberProperty{JsiiProperty: "principalAccount", GoGetter: "PrincipalAccount"},
			_jsii_.MemberMethod{JsiiMethod: "toJSON", GoMethod: "ToJSON"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberMethod{JsiiMethod: "withConditions", GoMethod: "WithConditions"},
			_jsii_.MemberMethod{JsiiMethod: "withSessionTags", GoMethod: "WithSessionTags"},
		},
		func() interface{} {
			j := jsiiProxy_SamlFederatedPrincipal{}
			_jsii_.InitJsiiProxy(&j.Type__awsiamFederatedPrincipal)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"cloudstructs.SamlIdentityProvider",
		reflect.TypeOf((*SamlIdentityProvider)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "samlIdentityProviderArn", GoGetter: "SamlIdentityProviderArn"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_SamlIdentityProvider{}
			_jsii_.InitJsiiProxy(&j.Type__constructsConstruct)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"cloudstructs.SamlIdentityProviderProps",
		reflect.TypeOf((*SamlIdentityProviderProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"cloudstructs.SlackApp",
		reflect.TypeOf((*SlackApp)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "appId", GoGetter: "AppId"},
			_jsii_.MemberProperty{JsiiProperty: "clientId", GoGetter: "ClientId"},
			_jsii_.MemberProperty{JsiiProperty: "clientSecret", GoGetter: "ClientSecret"},
			_jsii_.MemberProperty{JsiiProperty: "credentials", GoGetter: "Credentials"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "signingSecret", GoGetter: "SigningSecret"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "verificationToken", GoGetter: "VerificationToken"},
		},
		func() interface{} {
			j := jsiiProxy_SlackApp{}
			_jsii_.InitJsiiProxy(&j.Type__constructsConstruct)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"cloudstructs.SlackAppManifest",
		reflect.TypeOf((*SlackAppManifest)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "render", GoMethod: "Render"},
		},
		func() interface{} {
			return &jsiiProxy_SlackAppManifest{}
		},
	)
	_jsii_.RegisterStruct(
		"cloudstructs.SlackAppManifestAppHome",
		reflect.TypeOf((*SlackAppManifestAppHome)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"cloudstructs.SlackAppManifestDefinition",
		reflect.TypeOf((*SlackAppManifestDefinition)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "render", GoMethod: "Render"},
		},
		func() interface{} {
			return &jsiiProxy_SlackAppManifestDefinition{}
		},
	)
	_jsii_.RegisterStruct(
		"cloudstructs.SlackAppManifestEventSubscriptions",
		reflect.TypeOf((*SlackAppManifestEventSubscriptions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cloudstructs.SlackAppManifestInteractivity",
		reflect.TypeOf((*SlackAppManifestInteractivity)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cloudstructs.SlackAppManifestOauthConfig",
		reflect.TypeOf((*SlackAppManifestOauthConfig)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cloudstructs.SlackAppManifestProps",
		reflect.TypeOf((*SlackAppManifestProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cloudstructs.SlackAppManifestSettings",
		reflect.TypeOf((*SlackAppManifestSettings)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cloudstructs.SlackAppManifestShortcut",
		reflect.TypeOf((*SlackAppManifestShortcut)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"cloudstructs.SlackAppManifestShortcutType",
		reflect.TypeOf((*SlackAppManifestShortcutType)(nil)).Elem(),
		map[string]interface{}{
			"MESSAGE": SlackAppManifestShortcutType_MESSAGE,
			"GLOBAL": SlackAppManifestShortcutType_GLOBAL,
		},
	)
	_jsii_.RegisterStruct(
		"cloudstructs.SlackAppManifestSlashCommand",
		reflect.TypeOf((*SlackAppManifestSlashCommand)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cloudstructs.SlackAppManifestWorkflowStep",
		reflect.TypeOf((*SlackAppManifestWorkflowStep)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cloudstructs.SlackAppProps",
		reflect.TypeOf((*SlackAppProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"cloudstructs.SlackEvents",
		reflect.TypeOf((*SlackEvents)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "eventBus", GoGetter: "EventBus"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_SlackEvents{}
			_jsii_.InitJsiiProxy(&j.Type__constructsConstruct)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"cloudstructs.SlackEventsProps",
		reflect.TypeOf((*SlackEventsProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"cloudstructs.SlackTextract",
		reflect.TypeOf((*SlackTextract)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_SlackTextract{}
			_jsii_.InitJsiiProxy(&j.Type__constructsConstruct)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"cloudstructs.SlackTextractProps",
		reflect.TypeOf((*SlackTextractProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cloudstructs.SlackkAppManifestBotUser",
		reflect.TypeOf((*SlackkAppManifestBotUser)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"cloudstructs.SslServerTest",
		reflect.TypeOf((*SslServerTest)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "alarmTopic", GoGetter: "AlarmTopic"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_SslServerTest{}
			_jsii_.InitJsiiProxy(&j.Type__constructsConstruct)
			return &j
		},
	)
	_jsii_.RegisterEnum(
		"cloudstructs.SslServerTestGrade",
		reflect.TypeOf((*SslServerTestGrade)(nil)).Elem(),
		map[string]interface{}{
			"A_PLUS": SslServerTestGrade_A_PLUS,
			"A": SslServerTestGrade_A,
			"A_MINUS": SslServerTestGrade_A_MINUS,
			"B": SslServerTestGrade_B,
			"C": SslServerTestGrade_C,
			"D": SslServerTestGrade_D,
			"E": SslServerTestGrade_E,
			"F": SslServerTestGrade_F,
		},
	)
	_jsii_.RegisterStruct(
		"cloudstructs.SslServerTestProps",
		reflect.TypeOf((*SslServerTestProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"cloudstructs.StateMachineCustomResourceProvider",
		reflect.TypeOf((*StateMachineCustomResourceProvider)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "serviceToken", GoGetter: "ServiceToken"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_StateMachineCustomResourceProvider{}
			_jsii_.InitJsiiProxy(&j.Type__constructsConstruct)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"cloudstructs.StateMachineCustomResourceProviderProps",
		reflect.TypeOf((*StateMachineCustomResourceProviderProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"cloudstructs.StaticWebsite",
		reflect.TypeOf((*StaticWebsite)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "bucket", GoGetter: "Bucket"},
			_jsii_.MemberProperty{JsiiProperty: "distribution", GoGetter: "Distribution"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_StaticWebsite{}
			_jsii_.InitJsiiProxy(&j.Type__constructsConstruct)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"cloudstructs.StaticWebsiteProps",
		reflect.TypeOf((*StaticWebsiteProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"cloudstructs.ToolkitCleaner",
		reflect.TypeOf((*ToolkitCleaner)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_ToolkitCleaner{}
			_jsii_.InitJsiiProxy(&j.Type__constructsConstruct)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"cloudstructs.ToolkitCleanerProps",
		reflect.TypeOf((*ToolkitCleanerProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"cloudstructs.UrlShortener",
		reflect.TypeOf((*UrlShortener)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "api", GoGetter: "Api"},
			_jsii_.MemberProperty{JsiiProperty: "apiEndpoint", GoGetter: "ApiEndpoint"},
			_jsii_.MemberMethod{JsiiMethod: "grantInvoke", GoMethod: "GrantInvoke"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_UrlShortener{}
			_jsii_.InitJsiiProxy(&j.Type__constructsConstruct)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"cloudstructs.UrlShortenerProps",
		reflect.TypeOf((*UrlShortenerProps)(nil)).Elem(),
	)
}
