// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0


package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/v2/pkg/controller"

	organization "github.com/prolixalias/provider-cloudflare/internal/controller/cluster/cloudflare/organization"
policy "github.com/prolixalias/provider-cloudflare/internal/controller/cluster/notification/policy"
limit "github.com/prolixalias/provider-cloudflare/internal/controller/cluster/rate/limit"
trustdevicesettings "github.com/prolixalias/provider-cloudflare/internal/controller/cluster/zero/trustdevicesettings"
user "github.com/prolixalias/provider-cloudflare/internal/controller/cluster/cloudflare/user"
hostname "github.com/prolixalias/provider-cloudflare/internal/controller/cluster/custom/hostname"
database "github.com/prolixalias/provider-cloudflare/internal/controller/cluster/d1/database"
transitsite "github.com/prolixalias/provider-cloudflare/internal/controller/cluster/magic/transitsite"
connector "github.com/prolixalias/provider-cloudflare/internal/controller/cluster/sso/connector"
watermark "github.com/prolixalias/provider-cloudflare/internal/controller/cluster/stream/watermark"
hostnameweb3 "github.com/prolixalias/provider-cloudflare/internal/controller/cluster/web3/hostname"
trustaccessaicontrolsmcpportal "github.com/prolixalias/provider-cloudflare/internal/controller/cluster/zero/trustaccessaicontrolsmcpportal"
scanningexpression "github.com/prolixalias/provider-cloudflare/internal/controller/cluster/content/scanningexpression"
zonetransfersacl "github.com/prolixalias/provider-cloudflare/internal/controller/cluster/dns/zonetransfersacl"
wanstaticroute "github.com/prolixalias/provider-cloudflare/internal/controller/cluster/magic/wanstaticroute"
cacertificate "github.com/prolixalias/provider-cloudflare/internal/controller/cluster/origin/cacertificate"
tieredcache "github.com/prolixalias/provider-cloudflare/internal/controller/cluster/regional/tieredcache"
trustaccessaicontrolsmcpserver "github.com/prolixalias/provider-cloudflare/internal/controller/cluster/zero/trustaccessaicontrolsmcpserver"
trustdlpcustomentry "github.com/prolixalias/provider-cloudflare/internal/controller/cluster/zero/trustdlpcustomentry"
trustdlppredefinedentry "github.com/prolixalias/provider-cloudflare/internal/controller/cluster/zero/trustdlppredefinedentry"
filter "github.com/prolixalias/provider-cloudflare/internal/controller/cluster/cloudflare/filter"
shielddiscoveryoperation "github.com/prolixalias/provider-cloudflare/internal/controller/cluster/api/shielddiscoveryoperation"
ipprefix "github.com/prolixalias/provider-cloudflare/internal/controller/cluster/byo/ipprefix"
zonetransfersincoming "github.com/prolixalias/provider-cloudflare/internal/controller/cluster/dns/zonetransfersincoming"
zonetransferstsig "github.com/prolixalias/provider-cloudflare/internal/controller/cluster/dns/zonetransferstsig"
application "github.com/prolixalias/provider-cloudflare/internal/controller/cluster/spectrum/application"
validationrules "github.com/prolixalias/provider-cloudflare/internal/controller/cluster/token/validationrules"
trustaccesscustompage "github.com/prolixalias/provider-cloudflare/internal/controller/cluster/zero/trustaccesscustompage"
record "github.com/prolixalias/provider-cloudflare/internal/controller/cluster/dns/record"
list "github.com/prolixalias/provider-cloudflare/internal/controller/cluster/cloudflare/list"
download "github.com/prolixalias/provider-cloudflare/internal/controller/cluster/stream/download"
deployment "github.com/prolixalias/provider-cloudflare/internal/controller/cluster/workers/deployment"
trustaccesskeyconfiguration "github.com/prolixalias/provider-cloudflare/internal/controller/cluster/zero/trustaccesskeyconfiguration"
trusttunnelcloudflaredroute "github.com/prolixalias/provider-cloudflare/internal/controller/cluster/zero/trusttunnelcloudflaredroute"
dnssettingszone "github.com/prolixalias/provider-cloudflare/internal/controller/cluster/zone/dnssettings"
subscriptionzone "github.com/prolixalias/provider-cloudflare/internal/controller/cluster/zone/subscription"
shieldschemavalidationsettings "github.com/prolixalias/provider-cloudflare/internal/controller/cluster/api/shieldschemavalidationsettings"
snippet "github.com/prolixalias/provider-cloudflare/internal/controller/cluster/cloudflare/snippet"
zonetransferspeer "github.com/prolixalias/provider-cloudflare/internal/controller/cluster/dns/zonetransferspeer"
balancerpool "github.com/prolixalias/provider-cloudflare/internal/controller/cluster/load/balancerpool"
bucketlifecycle "github.com/prolixalias/provider-cloudflare/internal/controller/cluster/r2/bucketlifecycle"
tls "github.com/prolixalias/provider-cloudflare/internal/controller/cluster/total/tls"
sslsetting "github.com/prolixalias/provider-cloudflare/internal/controller/cluster/universal/sslsetting"
roomevent "github.com/prolixalias/provider-cloudflare/internal/controller/cluster/waiting/roomevent"
snippets "github.com/prolixalias/provider-cloudflare/internal/controller/cluster/cloudflare/snippets"
onerequestmessage "github.com/prolixalias/provider-cloudflare/internal/controller/cluster/cloudforce/onerequestmessage"
zonetransfersoutgoing "github.com/prolixalias/provider-cloudflare/internal/controller/cluster/dns/zonetransfersoutgoing"
routingrule "github.com/prolixalias/provider-cloudflare/internal/controller/cluster/email/routingrule"
providerconfig "github.com/prolixalias/provider-cloudflare/internal/controller/cluster/providerconfig"
bucketeventnotification "github.com/prolixalias/provider-cloudflare/internal/controller/cluster/r2/bucketeventnotification"
analyticssite "github.com/prolixalias/provider-cloudflare/internal/controller/cluster/web/analyticssite"
script "github.com/prolixalias/provider-cloudflare/internal/controller/cluster/workers/script"
management "github.com/prolixalias/provider-cloudflare/internal/controller/cluster/bot/management"
balancer "github.com/prolixalias/provider-cloudflare/internal/controller/cluster/load/balancer"
trustaccessservicetoken "github.com/prolixalias/provider-cloudflare/internal/controller/cluster/zero/trustaccessservicetoken"
trustdevicedefaultprofile "github.com/prolixalias/provider-cloudflare/internal/controller/cluster/zero/trustdevicedefaultprofile"
trustdlpentry "github.com/prolixalias/provider-cloudflare/internal/controller/cluster/zero/trustdlpentry"
trustgatewaycertificate "github.com/prolixalias/provider-cloudflare/internal/controller/cluster/zero/trustgatewaycertificate"
trustgatewayproxyendpoint "github.com/prolixalias/provider-cloudflare/internal/controller/cluster/zero/trustgatewayproxyendpoint"
trustlist "github.com/prolixalias/provider-cloudflare/internal/controller/cluster/zero/trustlist"
subscription "github.com/prolixalias/provider-cloudflare/internal/controller/cluster/account/subscription"
directoryservice "github.com/prolixalias/provider-cloudflare/internal/controller/cluster/connectivity/directoryservice"
routingcatchall "github.com/prolixalias/provider-cloudflare/internal/controller/cluster/email/routingcatchall"
ownershipchallenge "github.com/prolixalias/provider-cloudflare/internal/controller/cluster/logpush/ownershipchallenge"
trustaccessmtlshostnamesettings "github.com/prolixalias/provider-cloudflare/internal/controller/cluster/zero/trustaccessmtlshostnamesettings"
trustaccessshortlivedcertificate "github.com/prolixalias/provider-cloudflare/internal/controller/cluster/zero/trustaccessshortlivedcertificate"
trustdlppredefinedprofile "github.com/prolixalias/provider-cloudflare/internal/controller/cluster/zero/trustdlppredefinedprofile"
transitsitewan "github.com/prolixalias/provider-cloudflare/internal/controller/cluster/magic/transitsitewan"
bucketsippy "github.com/prolixalias/provider-cloudflare/internal/controller/cluster/r2/bucketsippy"
manageddomain "github.com/prolixalias/provider-cloudflare/internal/controller/cluster/r2/manageddomain"
normalizationsettings "github.com/prolixalias/provider-cloudflare/internal/controller/cluster/url/normalizationsettings"
trustnetworkhostnameroute "github.com/prolixalias/provider-cloudflare/internal/controller/cluster/zero/trustnetworkhostnameroute"
originpulls "github.com/prolixalias/provider-cloudflare/internal/controller/cluster/authenticated/originpulls"
onerequestpriority "github.com/prolixalias/provider-cloudflare/internal/controller/cluster/cloudforce/onerequestpriority"
validationsettings "github.com/prolixalias/provider-cloudflare/internal/controller/cluster/schema/validationsettings"
trustaccessidentityprovider "github.com/prolixalias/provider-cloudflare/internal/controller/cluster/zero/trustaccessidentityprovider"
trusttunnelcloudflaredvirtualnetwork "github.com/prolixalias/provider-cloudflare/internal/controller/cluster/zero/trusttunnelcloudflaredvirtualnetwork"
workflow "github.com/prolixalias/provider-cloudflare/internal/controller/cluster/cloudflare/workflow"
trustaccessgroup "github.com/prolixalias/provider-cloudflare/internal/controller/cluster/zero/trustaccessgroup"
member "github.com/prolixalias/provider-cloudflare/internal/controller/cluster/account/member"
shieldoperationschemavalidationsettings "github.com/prolixalias/provider-cloudflare/internal/controller/cluster/api/shieldoperationschemavalidationsettings"
stream "github.com/prolixalias/provider-cloudflare/internal/controller/cluster/cloudflare/stream"
transitsiteacl "github.com/prolixalias/provider-cloudflare/internal/controller/cluster/magic/transitsiteacl"
wanipsectunnel "github.com/prolixalias/provider-cloudflare/internal/controller/cluster/magic/wanipsectunnel"
trustaccessapplication "github.com/prolixalias/provider-cloudflare/internal/controller/cluster/zero/trustaccessapplication"
rule "github.com/prolixalias/provider-cloudflare/internal/controller/cluster/access/rule"
routingsettings "github.com/prolixalias/provider-cloudflare/internal/controller/cluster/email/routingsettings"
variant "github.com/prolixalias/provider-cloudflare/internal/controller/cluster/image/variant"
credentialcheck "github.com/prolixalias/provider-cloudflare/internal/controller/cluster/leaked/credentialcheck"
kvnamespace "github.com/prolixalias/provider-cloudflare/internal/controller/cluster/workers/kvnamespace"
route "github.com/prolixalias/provider-cloudflare/internal/controller/cluster/workers/route"
trustaccesstag "github.com/prolixalias/provider-cloudflare/internal/controller/cluster/zero/trustaccesstag"
trustdevicedefaultprofilelocaldomainfallback "github.com/prolixalias/provider-cloudflare/internal/controller/cluster/zero/trustdevicedefaultprofilelocaldomainfallback"
widget "github.com/prolixalias/provider-cloudflare/internal/controller/cluster/turnstile/widget"
analyticsrule "github.com/prolixalias/provider-cloudflare/internal/controller/cluster/web/analyticsrule"
hold "github.com/prolixalias/provider-cloudflare/internal/controller/cluster/zone/hold"
scanning "github.com/prolixalias/provider-cloudflare/internal/controller/cluster/content/scanning"
routingaddress "github.com/prolixalias/provider-cloudflare/internal/controller/cluster/email/routingaddress"
config "github.com/prolixalias/provider-cloudflare/internal/controller/cluster/hyperdrive/config"
rulepage "github.com/prolixalias/provider-cloudflare/internal/controller/cluster/page/rule"
liveinput "github.com/prolixalias/provider-cloudflare/internal/controller/cluster/stream/liveinput"
agentblockingrule "github.com/prolixalias/provider-cloudflare/internal/controller/cluster/user/agentblockingrule"
tokenapi "github.com/prolixalias/provider-cloudflare/internal/controller/cluster/api/token"
customdomain "github.com/prolixalias/provider-cloudflare/internal/controller/cluster/r2/customdomain"
crontrigger "github.com/prolixalias/provider-cloudflare/internal/controller/cluster/workers/crontrigger"
trustdnslocation "github.com/prolixalias/provider-cloudflare/internal/controller/cluster/zero/trustdnslocation"
account "github.com/prolixalias/provider-cloudflare/internal/controller/cluster/cloudflare/account"
securitytrusteddomains "github.com/prolixalias/provider-cloudflare/internal/controller/cluster/email/securitytrusteddomains"
project "github.com/prolixalias/provider-cloudflare/internal/controller/cluster/pages/project"
domainregistrar "github.com/prolixalias/provider-cloudflare/internal/controller/cluster/registrar/domain"
version "github.com/prolixalias/provider-cloudflare/internal/controller/cluster/worker/version"
trustaccessinfrastructuretarget "github.com/prolixalias/provider-cloudflare/internal/controller/cluster/zero/trustaccessinfrastructuretarget"
dnssettingsinternalview "github.com/prolixalias/provider-cloudflare/internal/controller/cluster/account/dnssettingsinternalview"
securityimpersonationregistry "github.com/prolixalias/provider-cloudflare/internal/controller/cluster/email/securityimpersonationregistry"
credentialcheckrule "github.com/prolixalias/provider-cloudflare/internal/controller/cluster/leaked/credentialcheckrule"
profile "github.com/prolixalias/provider-cloudflare/internal/controller/cluster/organization/profile"
bucketcors "github.com/prolixalias/provider-cloudflare/internal/controller/cluster/r2/bucketcors"
rulefirewall "github.com/prolixalias/provider-cloudflare/internal/controller/cluster/firewall/rule"
balancermonitor "github.com/prolixalias/provider-cloudflare/internal/controller/cluster/load/balancermonitor"
retention "github.com/prolixalias/provider-cloudflare/internal/controller/cluster/logpull/retention"
bucket "github.com/prolixalias/provider-cloudflare/internal/controller/cluster/r2/bucket"
roomrules "github.com/prolixalias/provider-cloudflare/internal/controller/cluster/waiting/roomrules"
lockdown "github.com/prolixalias/provider-cloudflare/internal/controller/cluster/zone/lockdown"
originpullscertificate "github.com/prolixalias/provider-cloudflare/internal/controller/cluster/authenticated/originpullscertificate"
queue "github.com/prolixalias/provider-cloudflare/internal/controller/cluster/cloudflare/queue"
onerequestasset "github.com/prolixalias/provider-cloudflare/internal/controller/cluster/cloudforce/onerequestasset"
hostnamefallbackorigin "github.com/prolixalias/provider-cloudflare/internal/controller/cluster/custom/hostnamefallbackorigin"
certificate "github.com/prolixalias/provider-cloudflare/internal/controller/cluster/keyless/certificate"
item "github.com/prolixalias/provider-cloudflare/internal/controller/cluster/list/item"
trustdevicemanagednetworks "github.com/prolixalias/provider-cloudflare/internal/controller/cluster/zero/trustdevicemanagednetworks"
trustdlpintegrationentry "github.com/prolixalias/provider-cloudflare/internal/controller/cluster/zero/trustdlpintegrationentry"
connectorrules "github.com/prolixalias/provider-cloudflare/internal/controller/cluster/cloud/connectorrules"
networkmonitoringrule "github.com/prolixalias/provider-cloudflare/internal/controller/cluster/magic/networkmonitoringrule"
domain "github.com/prolixalias/provider-cloudflare/internal/controller/cluster/pages/domain"
hostnameregional "github.com/prolixalias/provider-cloudflare/internal/controller/cluster/regional/hostname"
audiotrack "github.com/prolixalias/provider-cloudflare/internal/controller/cluster/stream/audiotrack"
trustriskbehavior "github.com/prolixalias/provider-cloudflare/internal/controller/cluster/zero/trustriskbehavior"
trustriskscoringintegration "github.com/prolixalias/provider-cloudflare/internal/controller/cluster/zero/trustriskscoringintegration"
trusttunnelcloudflaredconfig "github.com/prolixalias/provider-cloudflare/internal/controller/cluster/zero/trusttunnelcloudflaredconfig"
worker "github.com/prolixalias/provider-cloudflare/internal/controller/cluster/cloudflare/worker"
networkmonitoringconfiguration "github.com/prolixalias/provider-cloudflare/internal/controller/cluster/magic/networkmonitoringconfiguration"
certificatemtls "github.com/prolixalias/provider-cloudflare/internal/controller/cluster/mtls/certificate"
scheduledtest "github.com/prolixalias/provider-cloudflare/internal/controller/cluster/observatory/scheduledtest"
trustaccesspolicy "github.com/prolixalias/provider-cloudflare/internal/controller/cluster/zero/trustaccesspolicy"
trustdevicecustomprofile "github.com/prolixalias/provider-cloudflare/internal/controller/cluster/zero/trustdevicecustomprofile"
trustdevicedefaultprofilecertificates "github.com/prolixalias/provider-cloudflare/internal/controller/cluster/zero/trustdevicedefaultprofilecertificates"
dnssettings "github.com/prolixalias/provider-cloudflare/internal/controller/cluster/account/dnssettings"
shieldschema "github.com/prolixalias/provider-cloudflare/internal/controller/cluster/api/shieldschema"
tieredcaching "github.com/prolixalias/provider-cloudflare/internal/controller/cluster/argo/tieredcaching"
ruleset "github.com/prolixalias/provider-cloudflare/internal/controller/cluster/cloudflare/ruleset"
onerequest "github.com/prolixalias/provider-cloudflare/internal/controller/cluster/cloudforce/onerequest"
trustdextest "github.com/prolixalias/provider-cloudflare/internal/controller/cluster/zero/trustdextest"
dnssec "github.com/prolixalias/provider-cloudflare/internal/controller/cluster/zone/dnssec"
originpullssettings "github.com/prolixalias/provider-cloudflare/internal/controller/cluster/authenticated/originpullssettings"
pages "github.com/prolixalias/provider-cloudflare/internal/controller/cluster/custom/pages"
tlssetting "github.com/prolixalias/provider-cloudflare/internal/controller/cluster/hostname/tlssetting"
job "github.com/prolixalias/provider-cloudflare/internal/controller/cluster/logpush/job"
wangretunnel "github.com/prolixalias/provider-cloudflare/internal/controller/cluster/magic/wangretunnel"
validationoperationsettings "github.com/prolixalias/provider-cloudflare/internal/controller/cluster/schema/validationoperationsettings"
forplatformsdispatchnamespace "github.com/prolixalias/provider-cloudflare/internal/controller/cluster/workers/forplatformsdispatchnamespace"
trustdlpdataset "github.com/prolixalias/provider-cloudflare/internal/controller/cluster/zero/trustdlpdataset"
transitsitelan "github.com/prolixalias/provider-cloudflare/internal/controller/cluster/magic/transitsitelan"
key "github.com/prolixalias/provider-cloudflare/internal/controller/cluster/stream/key"
trustdlpcustomprofile "github.com/prolixalias/provider-cloudflare/internal/controller/cluster/zero/trustdlpcustomprofile"
trustgatewaylogging "github.com/prolixalias/provider-cloudflare/internal/controller/cluster/zero/trustgatewaylogging"
trustorganization "github.com/prolixalias/provider-cloudflare/internal/controller/cluster/zero/trustorganization"
securityblocksender "github.com/prolixalias/provider-cloudflare/internal/controller/cluster/email/securityblocksender"
trustdevicecustomprofilelocaldomainfallback "github.com/prolixalias/provider-cloudflare/internal/controller/cluster/zero/trustdevicecustomprofilelocaldomainfallback"
trustdevicepostureintegration "github.com/prolixalias/provider-cloudflare/internal/controller/cluster/zero/trustdevicepostureintegration"
trustgatewaypolicy "github.com/prolixalias/provider-cloudflare/internal/controller/cluster/zero/trustgatewaypolicy"
setting "github.com/prolixalias/provider-cloudflare/internal/controller/cluster/zone/setting"
healthcheck "github.com/prolixalias/provider-cloudflare/internal/controller/cluster/cloudflare/healthcheck"
zone "github.com/prolixalias/provider-cloudflare/internal/controller/cluster/cloudflare/zone"
ssl "github.com/prolixalias/provider-cloudflare/internal/controller/cluster/custom/ssl"
bucketlock "github.com/prolixalias/provider-cloudflare/internal/controller/cluster/r2/bucketlock"
cache "github.com/prolixalias/provider-cloudflare/internal/controller/cluster/tiered/cache"
validationconfig "github.com/prolixalias/provider-cloudflare/internal/controller/cluster/token/validationconfig"
trusttunnelcloudflared "github.com/prolixalias/provider-cloudflare/internal/controller/cluster/zero/trusttunnelcloudflared"
consumer "github.com/prolixalias/provider-cloudflare/internal/controller/cluster/queue/consumer"
token "github.com/prolixalias/provider-cloudflare/internal/controller/cluster/account/token"
addressmap "github.com/prolixalias/provider-cloudflare/internal/controller/cluster/addressmap/addressmap"
transforms "github.com/prolixalias/provider-cloudflare/internal/controller/cluster/managed/transforms"
policywebhooks "github.com/prolixalias/provider-cloudflare/internal/controller/cluster/notification/policywebhooks"
shieldpolicy "github.com/prolixalias/provider-cloudflare/internal/controller/cluster/page/shieldpolicy"
validationschemas "github.com/prolixalias/provider-cloudflare/internal/controller/cluster/schema/validationschemas"
captionlanguage "github.com/prolixalias/provider-cloudflare/internal/controller/cluster/stream/captionlanguage"
shield "github.com/prolixalias/provider-cloudflare/internal/controller/cluster/api/shield"
image "github.com/prolixalias/provider-cloudflare/internal/controller/cluster/cloudflare/image"
webhook "github.com/prolixalias/provider-cloudflare/internal/controller/cluster/stream/webhook"
customdomainworkers "github.com/prolixalias/provider-cloudflare/internal/controller/cluster/workers/customdomain"
kv "github.com/prolixalias/provider-cloudflare/internal/controller/cluster/workers/kv"
scriptsubdomain "github.com/prolixalias/provider-cloudflare/internal/controller/cluster/workers/scriptsubdomain"
smartrouting "github.com/prolixalias/provider-cloudflare/internal/controller/cluster/argo/smartrouting"
transitconnector "github.com/prolixalias/provider-cloudflare/internal/controller/cluster/magic/transitconnector"
room "github.com/prolixalias/provider-cloudflare/internal/controller/cluster/waiting/room"
trustaccessmtlscertificate "github.com/prolixalias/provider-cloudflare/internal/controller/cluster/zero/trustaccessmtlscertificate"
trustdeviceposturerule "github.com/prolixalias/provider-cloudflare/internal/controller/cluster/zero/trustdeviceposturerule"
trusttunnelwarpconnector "github.com/prolixalias/provider-cloudflare/internal/controller/cluster/zero/trusttunnelwarpconnector"
cachevariants "github.com/prolixalias/provider-cloudflare/internal/controller/cluster/zone/cachevariants"
firewall "github.com/prolixalias/provider-cloudflare/internal/controller/cluster/dns/firewall"
sfuapp "github.com/prolixalias/provider-cloudflare/internal/controller/cluster/calls/sfuapp"
pack "github.com/prolixalias/provider-cloudflare/internal/controller/cluster/certificate/pack"
routingdns "github.com/prolixalias/provider-cloudflare/internal/controller/cluster/email/routingdns"
rules "github.com/prolixalias/provider-cloudflare/internal/controller/cluster/snippet/rules"
roomsettings "github.com/prolixalias/provider-cloudflare/internal/controller/cluster/waiting/roomsettings"
trustgatewaysettings "github.com/prolixalias/provider-cloudflare/internal/controller/cluster/zero/trustgatewaysettings"
cachereserve "github.com/prolixalias/provider-cloudflare/internal/controller/cluster/zone/cachereserve"
shieldoperation "github.com/prolixalias/provider-cloudflare/internal/controller/cluster/api/shieldoperation"
turnapp "github.com/prolixalias/provider-cloudflare/internal/controller/cluster/calls/turnapp"

)

// Setup creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		rule.Setup,
		dnssettings.Setup,
		dnssettingsinternalview.Setup,
		member.Setup,
		subscription.Setup,
		token.Setup,
		addressmap.Setup,
		shield.Setup,
		shielddiscoveryoperation.Setup,
		shieldoperation.Setup,
		shieldoperationschemavalidationsettings.Setup,
		shieldschema.Setup,
		shieldschemavalidationsettings.Setup,
		tokenapi.Setup,
		smartrouting.Setup,
		tieredcaching.Setup,
		originpulls.Setup,
		originpullscertificate.Setup,
		originpullssettings.Setup,
		management.Setup,
		ipprefix.Setup,
		sfuapp.Setup,
		turnapp.Setup,
		pack.Setup,
		connectorrules.Setup,
		account.Setup,
		filter.Setup,
		healthcheck.Setup,
		image.Setup,
		list.Setup,
		organization.Setup,
		queue.Setup,
		ruleset.Setup,
		snippet.Setup,
		snippets.Setup,
		stream.Setup,
		user.Setup,
		worker.Setup,
		workflow.Setup,
		zone.Setup,
		onerequest.Setup,
		onerequestasset.Setup,
		onerequestmessage.Setup,
		onerequestpriority.Setup,
		directoryservice.Setup,
		scanning.Setup,
		scanningexpression.Setup,
		hostname.Setup,
		hostnamefallbackorigin.Setup,
		pages.Setup,
		ssl.Setup,
		database.Setup,
		firewall.Setup,
		record.Setup,
		zonetransfersacl.Setup,
		zonetransfersincoming.Setup,
		zonetransfersoutgoing.Setup,
		zonetransferspeer.Setup,
		zonetransferstsig.Setup,
		routingaddress.Setup,
		routingcatchall.Setup,
		routingdns.Setup,
		routingrule.Setup,
		routingsettings.Setup,
		securityblocksender.Setup,
		securityimpersonationregistry.Setup,
		securitytrusteddomains.Setup,
		rulefirewall.Setup,
		tlssetting.Setup,
		config.Setup,
		variant.Setup,
		certificate.Setup,
		credentialcheck.Setup,
		credentialcheckrule.Setup,
		item.Setup,
		balancer.Setup,
		balancermonitor.Setup,
		balancerpool.Setup,
		retention.Setup,
		job.Setup,
		ownershipchallenge.Setup,
		networkmonitoringconfiguration.Setup,
		networkmonitoringrule.Setup,
		transitconnector.Setup,
		transitsite.Setup,
		transitsiteacl.Setup,
		transitsitelan.Setup,
		transitsitewan.Setup,
		wangretunnel.Setup,
		wanipsectunnel.Setup,
		wanstaticroute.Setup,
		transforms.Setup,
		certificatemtls.Setup,
		policy.Setup,
		policywebhooks.Setup,
		scheduledtest.Setup,
		profile.Setup,
		cacertificate.Setup,
		rulepage.Setup,
		shieldpolicy.Setup,
		domain.Setup,
		project.Setup,
		providerconfig.Setup,
		consumer.Setup,
		bucket.Setup,
		bucketcors.Setup,
		bucketeventnotification.Setup,
		bucketlifecycle.Setup,
		bucketlock.Setup,
		bucketsippy.Setup,
		customdomain.Setup,
		manageddomain.Setup,
		limit.Setup,
		hostnameregional.Setup,
		tieredcache.Setup,
		domainregistrar.Setup,
		validationoperationsettings.Setup,
		validationschemas.Setup,
		validationsettings.Setup,
		rules.Setup,
		application.Setup,
		connector.Setup,
		audiotrack.Setup,
		captionlanguage.Setup,
		download.Setup,
		key.Setup,
		liveinput.Setup,
		watermark.Setup,
		webhook.Setup,
		cache.Setup,
		validationconfig.Setup,
		validationrules.Setup,
		tls.Setup,
		widget.Setup,
		sslsetting.Setup,
		normalizationsettings.Setup,
		agentblockingrule.Setup,
		room.Setup,
		roomevent.Setup,
		roomrules.Setup,
		roomsettings.Setup,
		analyticsrule.Setup,
		analyticssite.Setup,
		hostnameweb3.Setup,
		version.Setup,
		crontrigger.Setup,
		customdomainworkers.Setup,
		deployment.Setup,
		forplatformsdispatchnamespace.Setup,
		kv.Setup,
		kvnamespace.Setup,
		route.Setup,
		script.Setup,
		scriptsubdomain.Setup,
		trustaccessaicontrolsmcpportal.Setup,
		trustaccessaicontrolsmcpserver.Setup,
		trustaccessapplication.Setup,
		trustaccesscustompage.Setup,
		trustaccessgroup.Setup,
		trustaccessidentityprovider.Setup,
		trustaccessinfrastructuretarget.Setup,
		trustaccesskeyconfiguration.Setup,
		trustaccessmtlscertificate.Setup,
		trustaccessmtlshostnamesettings.Setup,
		trustaccesspolicy.Setup,
		trustaccessservicetoken.Setup,
		trustaccessshortlivedcertificate.Setup,
		trustaccesstag.Setup,
		trustdevicecustomprofile.Setup,
		trustdevicecustomprofilelocaldomainfallback.Setup,
		trustdevicedefaultprofile.Setup,
		trustdevicedefaultprofilecertificates.Setup,
		trustdevicedefaultprofilelocaldomainfallback.Setup,
		trustdevicemanagednetworks.Setup,
		trustdevicepostureintegration.Setup,
		trustdeviceposturerule.Setup,
		trustdevicesettings.Setup,
		trustdextest.Setup,
		trustdlpcustomentry.Setup,
		trustdlpcustomprofile.Setup,
		trustdlpdataset.Setup,
		trustdlpentry.Setup,
		trustdlpintegrationentry.Setup,
		trustdlppredefinedentry.Setup,
		trustdlppredefinedprofile.Setup,
		trustdnslocation.Setup,
		trustgatewaycertificate.Setup,
		trustgatewaylogging.Setup,
		trustgatewaypolicy.Setup,
		trustgatewayproxyendpoint.Setup,
		trustgatewaysettings.Setup,
		trustlist.Setup,
		trustnetworkhostnameroute.Setup,
		trustorganization.Setup,
		trustriskbehavior.Setup,
		trustriskscoringintegration.Setup,
		trusttunnelcloudflared.Setup,
		trusttunnelcloudflaredconfig.Setup,
		trusttunnelcloudflaredroute.Setup,
		trusttunnelcloudflaredvirtualnetwork.Setup,
		trusttunnelwarpconnector.Setup,
		cachereserve.Setup,
		cachevariants.Setup,
		dnssec.Setup,
		dnssettingszone.Setup,
		hold.Setup,
		lockdown.Setup,
		setting.Setup,
		subscriptionzone.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}

// SetupGated creates all controllers with the supplied logger and adds them to
// the supplied manager gated.
func SetupGated(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		rule.SetupGated,
		dnssettings.SetupGated,
		dnssettingsinternalview.SetupGated,
		member.SetupGated,
		subscription.SetupGated,
		token.SetupGated,
		addressmap.SetupGated,
		shield.SetupGated,
		shielddiscoveryoperation.SetupGated,
		shieldoperation.SetupGated,
		shieldoperationschemavalidationsettings.SetupGated,
		shieldschema.SetupGated,
		shieldschemavalidationsettings.SetupGated,
		tokenapi.SetupGated,
		smartrouting.SetupGated,
		tieredcaching.SetupGated,
		originpulls.SetupGated,
		originpullscertificate.SetupGated,
		originpullssettings.SetupGated,
		management.SetupGated,
		ipprefix.SetupGated,
		sfuapp.SetupGated,
		turnapp.SetupGated,
		pack.SetupGated,
		connectorrules.SetupGated,
		account.SetupGated,
		filter.SetupGated,
		healthcheck.SetupGated,
		image.SetupGated,
		list.SetupGated,
		organization.SetupGated,
		queue.SetupGated,
		ruleset.SetupGated,
		snippet.SetupGated,
		snippets.SetupGated,
		stream.SetupGated,
		user.SetupGated,
		worker.SetupGated,
		workflow.SetupGated,
		zone.SetupGated,
		onerequest.SetupGated,
		onerequestasset.SetupGated,
		onerequestmessage.SetupGated,
		onerequestpriority.SetupGated,
		directoryservice.SetupGated,
		scanning.SetupGated,
		scanningexpression.SetupGated,
		hostname.SetupGated,
		hostnamefallbackorigin.SetupGated,
		pages.SetupGated,
		ssl.SetupGated,
		database.SetupGated,
		firewall.SetupGated,
		record.SetupGated,
		zonetransfersacl.SetupGated,
		zonetransfersincoming.SetupGated,
		zonetransfersoutgoing.SetupGated,
		zonetransferspeer.SetupGated,
		zonetransferstsig.SetupGated,
		routingaddress.SetupGated,
		routingcatchall.SetupGated,
		routingdns.SetupGated,
		routingrule.SetupGated,
		routingsettings.SetupGated,
		securityblocksender.SetupGated,
		securityimpersonationregistry.SetupGated,
		securitytrusteddomains.SetupGated,
		rulefirewall.SetupGated,
		tlssetting.SetupGated,
		config.SetupGated,
		variant.SetupGated,
		certificate.SetupGated,
		credentialcheck.SetupGated,
		credentialcheckrule.SetupGated,
		item.SetupGated,
		balancer.SetupGated,
		balancermonitor.SetupGated,
		balancerpool.SetupGated,
		retention.SetupGated,
		job.SetupGated,
		ownershipchallenge.SetupGated,
		networkmonitoringconfiguration.SetupGated,
		networkmonitoringrule.SetupGated,
		transitconnector.SetupGated,
		transitsite.SetupGated,
		transitsiteacl.SetupGated,
		transitsitelan.SetupGated,
		transitsitewan.SetupGated,
		wangretunnel.SetupGated,
		wanipsectunnel.SetupGated,
		wanstaticroute.SetupGated,
		transforms.SetupGated,
		certificatemtls.SetupGated,
		policy.SetupGated,
		policywebhooks.SetupGated,
		scheduledtest.SetupGated,
		profile.SetupGated,
		cacertificate.SetupGated,
		rulepage.SetupGated,
		shieldpolicy.SetupGated,
		domain.SetupGated,
		project.SetupGated,
		providerconfig.SetupGated,
		consumer.SetupGated,
		bucket.SetupGated,
		bucketcors.SetupGated,
		bucketeventnotification.SetupGated,
		bucketlifecycle.SetupGated,
		bucketlock.SetupGated,
		bucketsippy.SetupGated,
		customdomain.SetupGated,
		manageddomain.SetupGated,
		limit.SetupGated,
		hostnameregional.SetupGated,
		tieredcache.SetupGated,
		domainregistrar.SetupGated,
		validationoperationsettings.SetupGated,
		validationschemas.SetupGated,
		validationsettings.SetupGated,
		rules.SetupGated,
		application.SetupGated,
		connector.SetupGated,
		audiotrack.SetupGated,
		captionlanguage.SetupGated,
		download.SetupGated,
		key.SetupGated,
		liveinput.SetupGated,
		watermark.SetupGated,
		webhook.SetupGated,
		cache.SetupGated,
		validationconfig.SetupGated,
		validationrules.SetupGated,
		tls.SetupGated,
		widget.SetupGated,
		sslsetting.SetupGated,
		normalizationsettings.SetupGated,
		agentblockingrule.SetupGated,
		room.SetupGated,
		roomevent.SetupGated,
		roomrules.SetupGated,
		roomsettings.SetupGated,
		analyticsrule.SetupGated,
		analyticssite.SetupGated,
		hostnameweb3.SetupGated,
		version.SetupGated,
		crontrigger.SetupGated,
		customdomainworkers.SetupGated,
		deployment.SetupGated,
		forplatformsdispatchnamespace.SetupGated,
		kv.SetupGated,
		kvnamespace.SetupGated,
		route.SetupGated,
		script.SetupGated,
		scriptsubdomain.SetupGated,
		trustaccessaicontrolsmcpportal.SetupGated,
		trustaccessaicontrolsmcpserver.SetupGated,
		trustaccessapplication.SetupGated,
		trustaccesscustompage.SetupGated,
		trustaccessgroup.SetupGated,
		trustaccessidentityprovider.SetupGated,
		trustaccessinfrastructuretarget.SetupGated,
		trustaccesskeyconfiguration.SetupGated,
		trustaccessmtlscertificate.SetupGated,
		trustaccessmtlshostnamesettings.SetupGated,
		trustaccesspolicy.SetupGated,
		trustaccessservicetoken.SetupGated,
		trustaccessshortlivedcertificate.SetupGated,
		trustaccesstag.SetupGated,
		trustdevicecustomprofile.SetupGated,
		trustdevicecustomprofilelocaldomainfallback.SetupGated,
		trustdevicedefaultprofile.SetupGated,
		trustdevicedefaultprofilecertificates.SetupGated,
		trustdevicedefaultprofilelocaldomainfallback.SetupGated,
		trustdevicemanagednetworks.SetupGated,
		trustdevicepostureintegration.SetupGated,
		trustdeviceposturerule.SetupGated,
		trustdevicesettings.SetupGated,
		trustdextest.SetupGated,
		trustdlpcustomentry.SetupGated,
		trustdlpcustomprofile.SetupGated,
		trustdlpdataset.SetupGated,
		trustdlpentry.SetupGated,
		trustdlpintegrationentry.SetupGated,
		trustdlppredefinedentry.SetupGated,
		trustdlppredefinedprofile.SetupGated,
		trustdnslocation.SetupGated,
		trustgatewaycertificate.SetupGated,
		trustgatewaylogging.SetupGated,
		trustgatewaypolicy.SetupGated,
		trustgatewayproxyendpoint.SetupGated,
		trustgatewaysettings.SetupGated,
		trustlist.SetupGated,
		trustnetworkhostnameroute.SetupGated,
		trustorganization.SetupGated,
		trustriskbehavior.SetupGated,
		trustriskscoringintegration.SetupGated,
		trusttunnelcloudflared.SetupGated,
		trusttunnelcloudflaredconfig.SetupGated,
		trusttunnelcloudflaredroute.SetupGated,
		trusttunnelcloudflaredvirtualnetwork.SetupGated,
		trusttunnelwarpconnector.SetupGated,
		cachereserve.SetupGated,
		cachevariants.SetupGated,
		dnssec.SetupGated,
		dnssettingszone.SetupGated,
		hold.SetupGated,
		lockdown.SetupGated,
		setting.SetupGated,
		subscriptionzone.SetupGated,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}