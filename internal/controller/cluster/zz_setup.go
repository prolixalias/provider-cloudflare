// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0


package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/v2/pkg/controller"

	record "github.com/prolixalias/provider-cloudflare/internal/controller/cluster/dns/record"
connector "github.com/prolixalias/provider-cloudflare/internal/controller/cluster/sso/connector"
webhook "github.com/prolixalias/provider-cloudflare/internal/controller/cluster/stream/webhook"
deployment "github.com/prolixalias/provider-cloudflare/internal/controller/cluster/workers/deployment"
scriptsubdomain "github.com/prolixalias/provider-cloudflare/internal/controller/cluster/workers/scriptsubdomain"
tieredcaching "github.com/prolixalias/provider-cloudflare/internal/controller/cluster/argo/tieredcaching"
routingdns "github.com/prolixalias/provider-cloudflare/internal/controller/cluster/email/routingdns"
networkmonitoringrule "github.com/prolixalias/provider-cloudflare/internal/controller/cluster/magic/networkmonitoringrule"
transitsite "github.com/prolixalias/provider-cloudflare/internal/controller/cluster/magic/transitsite"
transitsitelan "github.com/prolixalias/provider-cloudflare/internal/controller/cluster/magic/transitsitelan"
download "github.com/prolixalias/provider-cloudflare/internal/controller/cluster/stream/download"
watermark "github.com/prolixalias/provider-cloudflare/internal/controller/cluster/stream/watermark"
sslsetting "github.com/prolixalias/provider-cloudflare/internal/controller/cluster/universal/sslsetting"
directoryservice "github.com/prolixalias/provider-cloudflare/internal/controller/cluster/connectivity/directoryservice"
firewall "github.com/prolixalias/provider-cloudflare/internal/controller/cluster/dns/firewall"
rulefirewall "github.com/prolixalias/provider-cloudflare/internal/controller/cluster/firewall/rule"
credentialcheck "github.com/prolixalias/provider-cloudflare/internal/controller/cluster/leaked/credentialcheck"
balancer "github.com/prolixalias/provider-cloudflare/internal/controller/cluster/load/balancer"
balancermonitor "github.com/prolixalias/provider-cloudflare/internal/controller/cluster/load/balancermonitor"
domainregistrar "github.com/prolixalias/provider-cloudflare/internal/controller/cluster/registrar/domain"
kv "github.com/prolixalias/provider-cloudflare/internal/controller/cluster/workers/kv"
shield "github.com/prolixalias/provider-cloudflare/internal/controller/cluster/api/shield"
tokenapi "github.com/prolixalias/provider-cloudflare/internal/controller/cluster/api/token"
user "github.com/prolixalias/provider-cloudflare/internal/controller/cluster/cloudflare/user"
routingcatchall "github.com/prolixalias/provider-cloudflare/internal/controller/cluster/email/routingcatchall"
bucketcors "github.com/prolixalias/provider-cloudflare/internal/controller/cluster/r2/bucketcors"
trustaccessgroup "github.com/prolixalias/provider-cloudflare/internal/controller/cluster/zero/trustaccessgroup"
trustdevicesettings "github.com/prolixalias/provider-cloudflare/internal/controller/cluster/zero/trustdevicesettings"
trustorganization "github.com/prolixalias/provider-cloudflare/internal/controller/cluster/zero/trustorganization"
routingrule "github.com/prolixalias/provider-cloudflare/internal/controller/cluster/email/routingrule"
certificate "github.com/prolixalias/provider-cloudflare/internal/controller/cluster/keyless/certificate"
tieredcache "github.com/prolixalias/provider-cloudflare/internal/controller/cluster/regional/tieredcache"
trustaccessmtlscertificate "github.com/prolixalias/provider-cloudflare/internal/controller/cluster/zero/trustaccessmtlscertificate"
trustdlpentry "github.com/prolixalias/provider-cloudflare/internal/controller/cluster/zero/trustdlpentry"
setting "github.com/prolixalias/provider-cloudflare/internal/controller/cluster/zone/setting"
rule "github.com/prolixalias/provider-cloudflare/internal/controller/cluster/access/rule"
ruleset "github.com/prolixalias/provider-cloudflare/internal/controller/cluster/cloudflare/ruleset"
hostname "github.com/prolixalias/provider-cloudflare/internal/controller/cluster/custom/hostname"
rulepage "github.com/prolixalias/provider-cloudflare/internal/controller/cluster/page/rule"
captionlanguage "github.com/prolixalias/provider-cloudflare/internal/controller/cluster/stream/captionlanguage"
key "github.com/prolixalias/provider-cloudflare/internal/controller/cluster/stream/key"
validationrules "github.com/prolixalias/provider-cloudflare/internal/controller/cluster/token/validationrules"
trustdlpcustomentry "github.com/prolixalias/provider-cloudflare/internal/controller/cluster/zero/trustdlpcustomentry"
shieldoperation "github.com/prolixalias/provider-cloudflare/internal/controller/cluster/api/shieldoperation"
queue "github.com/prolixalias/provider-cloudflare/internal/controller/cluster/cloudflare/queue"
onerequestmessage "github.com/prolixalias/provider-cloudflare/internal/controller/cluster/cloudforce/onerequestmessage"
pages "github.com/prolixalias/provider-cloudflare/internal/controller/cluster/custom/pages"
networkmonitoringconfiguration "github.com/prolixalias/provider-cloudflare/internal/controller/cluster/magic/networkmonitoringconfiguration"
transitsitewan "github.com/prolixalias/provider-cloudflare/internal/controller/cluster/magic/transitsitewan"
certificatemtls "github.com/prolixalias/provider-cloudflare/internal/controller/cluster/mtls/certificate"
roomsettings "github.com/prolixalias/provider-cloudflare/internal/controller/cluster/waiting/roomsettings"
list "github.com/prolixalias/provider-cloudflare/internal/controller/cluster/cloudflare/list"
wanipsectunnel "github.com/prolixalias/provider-cloudflare/internal/controller/cluster/magic/wanipsectunnel"
trustdextest "github.com/prolixalias/provider-cloudflare/internal/controller/cluster/zero/trustdextest"
trustgatewaysettings "github.com/prolixalias/provider-cloudflare/internal/controller/cluster/zero/trustgatewaysettings"
consumer "github.com/prolixalias/provider-cloudflare/internal/controller/cluster/queue/consumer"
audiotrack "github.com/prolixalias/provider-cloudflare/internal/controller/cluster/stream/audiotrack"
shieldschemavalidationsettings "github.com/prolixalias/provider-cloudflare/internal/controller/cluster/api/shieldschemavalidationsettings"
workflow "github.com/prolixalias/provider-cloudflare/internal/controller/cluster/cloudflare/workflow"
customdomain "github.com/prolixalias/provider-cloudflare/internal/controller/cluster/r2/customdomain"
hostnameregional "github.com/prolixalias/provider-cloudflare/internal/controller/cluster/regional/hostname"
validationconfig "github.com/prolixalias/provider-cloudflare/internal/controller/cluster/token/validationconfig"
trustlist "github.com/prolixalias/provider-cloudflare/internal/controller/cluster/zero/trustlist"
management "github.com/prolixalias/provider-cloudflare/internal/controller/cluster/bot/management"
account "github.com/prolixalias/provider-cloudflare/internal/controller/cluster/cloudflare/account"
stream "github.com/prolixalias/provider-cloudflare/internal/controller/cluster/cloudflare/stream"
zonetransfersoutgoing "github.com/prolixalias/provider-cloudflare/internal/controller/cluster/dns/zonetransfersoutgoing"
job "github.com/prolixalias/provider-cloudflare/internal/controller/cluster/logpush/job"
manageddomain "github.com/prolixalias/provider-cloudflare/internal/controller/cluster/r2/manageddomain"
trustaccessshortlivedcertificate "github.com/prolixalias/provider-cloudflare/internal/controller/cluster/zero/trustaccessshortlivedcertificate"
trustdevicemanagednetworks "github.com/prolixalias/provider-cloudflare/internal/controller/cluster/zero/trustdevicemanagednetworks"
roomrules "github.com/prolixalias/provider-cloudflare/internal/controller/cluster/waiting/roomrules"
customdomainworkers "github.com/prolixalias/provider-cloudflare/internal/controller/cluster/workers/customdomain"
route "github.com/prolixalias/provider-cloudflare/internal/controller/cluster/workers/route"
trustdlpcustomprofile "github.com/prolixalias/provider-cloudflare/internal/controller/cluster/zero/trustdlpcustomprofile"
trustdlpdataset "github.com/prolixalias/provider-cloudflare/internal/controller/cluster/zero/trustdlpdataset"
item "github.com/prolixalias/provider-cloudflare/internal/controller/cluster/list/item"
rules "github.com/prolixalias/provider-cloudflare/internal/controller/cluster/snippet/rules"
analyticsrule "github.com/prolixalias/provider-cloudflare/internal/controller/cluster/web/analyticsrule"
trustaccesscustompage "github.com/prolixalias/provider-cloudflare/internal/controller/cluster/zero/trustaccesscustompage"
trustdevicecustomprofilelocaldomainfallback "github.com/prolixalias/provider-cloudflare/internal/controller/cluster/zero/trustdevicecustomprofilelocaldomainfallback"
trustnetworkhostnameroute "github.com/prolixalias/provider-cloudflare/internal/controller/cluster/zero/trustnetworkhostnameroute"
shieldoperationschemavalidationsettings "github.com/prolixalias/provider-cloudflare/internal/controller/cluster/api/shieldoperationschemavalidationsettings"
transforms "github.com/prolixalias/provider-cloudflare/internal/controller/cluster/managed/transforms"
kvnamespace "github.com/prolixalias/provider-cloudflare/internal/controller/cluster/workers/kvnamespace"
script "github.com/prolixalias/provider-cloudflare/internal/controller/cluster/workers/script"
trustaccesstag "github.com/prolixalias/provider-cloudflare/internal/controller/cluster/zero/trustaccesstag"
trustdevicecustomprofile "github.com/prolixalias/provider-cloudflare/internal/controller/cluster/zero/trustdevicecustomprofile"
trustdevicedefaultprofilelocaldomainfallback "github.com/prolixalias/provider-cloudflare/internal/controller/cluster/zero/trustdevicedefaultprofilelocaldomainfallback"
cachereserve "github.com/prolixalias/provider-cloudflare/internal/controller/cluster/zone/cachereserve"
subscription "github.com/prolixalias/provider-cloudflare/internal/controller/cluster/account/subscription"
bucketeventnotification "github.com/prolixalias/provider-cloudflare/internal/controller/cluster/r2/bucketeventnotification"
roomevent "github.com/prolixalias/provider-cloudflare/internal/controller/cluster/waiting/roomevent"
trusttunnelcloudflaredconfig "github.com/prolixalias/provider-cloudflare/internal/controller/cluster/zero/trusttunnelcloudflaredconfig"
originpulls "github.com/prolixalias/provider-cloudflare/internal/controller/cluster/authenticated/originpulls"
healthcheck "github.com/prolixalias/provider-cloudflare/internal/controller/cluster/cloudflare/healthcheck"
variant "github.com/prolixalias/provider-cloudflare/internal/controller/cluster/image/variant"
wanstaticroute "github.com/prolixalias/provider-cloudflare/internal/controller/cluster/magic/wanstaticroute"
scheduledtest "github.com/prolixalias/provider-cloudflare/internal/controller/cluster/observatory/scheduledtest"
application "github.com/prolixalias/provider-cloudflare/internal/controller/cluster/spectrum/application"
sfuapp "github.com/prolixalias/provider-cloudflare/internal/controller/cluster/calls/sfuapp"
worker "github.com/prolixalias/provider-cloudflare/internal/controller/cluster/cloudflare/worker"
onerequest "github.com/prolixalias/provider-cloudflare/internal/controller/cluster/cloudforce/onerequest"
retention "github.com/prolixalias/provider-cloudflare/internal/controller/cluster/logpull/retention"
limit "github.com/prolixalias/provider-cloudflare/internal/controller/cluster/rate/limit"
version "github.com/prolixalias/provider-cloudflare/internal/controller/cluster/worker/version"
crontrigger "github.com/prolixalias/provider-cloudflare/internal/controller/cluster/workers/crontrigger"
trustaccessinfrastructuretarget "github.com/prolixalias/provider-cloudflare/internal/controller/cluster/zero/trustaccessinfrastructuretarget"
pack "github.com/prolixalias/provider-cloudflare/internal/controller/cluster/certificate/pack"
balancerpool "github.com/prolixalias/provider-cloudflare/internal/controller/cluster/load/balancerpool"
wangretunnel "github.com/prolixalias/provider-cloudflare/internal/controller/cluster/magic/wangretunnel"
trustdevicepostureintegration "github.com/prolixalias/provider-cloudflare/internal/controller/cluster/zero/trustdevicepostureintegration"
trustgatewaypolicy "github.com/prolixalias/provider-cloudflare/internal/controller/cluster/zero/trustgatewaypolicy"
trustgatewayproxyendpoint "github.com/prolixalias/provider-cloudflare/internal/controller/cluster/zero/trustgatewayproxyendpoint"
trusttunnelcloudflaredvirtualnetwork "github.com/prolixalias/provider-cloudflare/internal/controller/cluster/zero/trusttunnelcloudflaredvirtualnetwork"
hostnamefallbackorigin "github.com/prolixalias/provider-cloudflare/internal/controller/cluster/custom/hostnamefallbackorigin"
zonetransferspeer "github.com/prolixalias/provider-cloudflare/internal/controller/cluster/dns/zonetransferspeer"
securityblocksender "github.com/prolixalias/provider-cloudflare/internal/controller/cluster/email/securityblocksender"
domain "github.com/prolixalias/provider-cloudflare/internal/controller/cluster/pages/domain"
trustaccessaicontrolsmcpserver "github.com/prolixalias/provider-cloudflare/internal/controller/cluster/zero/trustaccessaicontrolsmcpserver"
trustdeviceposturerule "github.com/prolixalias/provider-cloudflare/internal/controller/cluster/zero/trustdeviceposturerule"
trustdlpintegrationentry "github.com/prolixalias/provider-cloudflare/internal/controller/cluster/zero/trustdlpintegrationentry"
trustgatewaylogging "github.com/prolixalias/provider-cloudflare/internal/controller/cluster/zero/trustgatewaylogging"
scanningexpression "github.com/prolixalias/provider-cloudflare/internal/controller/cluster/content/scanningexpression"
routingsettings "github.com/prolixalias/provider-cloudflare/internal/controller/cluster/email/routingsettings"
validationschemas "github.com/prolixalias/provider-cloudflare/internal/controller/cluster/schema/validationschemas"
room "github.com/prolixalias/provider-cloudflare/internal/controller/cluster/waiting/room"
trustaccesskeyconfiguration "github.com/prolixalias/provider-cloudflare/internal/controller/cluster/zero/trustaccesskeyconfiguration"
trusttunnelwarpconnector "github.com/prolixalias/provider-cloudflare/internal/controller/cluster/zero/trusttunnelwarpconnector"
dnssettingszone "github.com/prolixalias/provider-cloudflare/internal/controller/cluster/zone/dnssettings"
subscriptionzone "github.com/prolixalias/provider-cloudflare/internal/controller/cluster/zone/subscription"
originpullscertificate "github.com/prolixalias/provider-cloudflare/internal/controller/cluster/authenticated/originpullscertificate"
securitytrusteddomains "github.com/prolixalias/provider-cloudflare/internal/controller/cluster/email/securitytrusteddomains"
addressmap "github.com/prolixalias/provider-cloudflare/internal/controller/cluster/addressmap/addressmap"
shielddiscoveryoperation "github.com/prolixalias/provider-cloudflare/internal/controller/cluster/api/shielddiscoveryoperation"
snippet "github.com/prolixalias/provider-cloudflare/internal/controller/cluster/cloudflare/snippet"
zone "github.com/prolixalias/provider-cloudflare/internal/controller/cluster/cloudflare/zone"
config "github.com/prolixalias/provider-cloudflare/internal/controller/cluster/hyperdrive/config"
widget "github.com/prolixalias/provider-cloudflare/internal/controller/cluster/turnstile/widget"
trusttunnelcloudflaredroute "github.com/prolixalias/provider-cloudflare/internal/controller/cluster/zero/trusttunnelcloudflaredroute"
connectorrules "github.com/prolixalias/provider-cloudflare/internal/controller/cluster/cloud/connectorrules"
onerequestasset "github.com/prolixalias/provider-cloudflare/internal/controller/cluster/cloudforce/onerequestasset"
credentialcheckrule "github.com/prolixalias/provider-cloudflare/internal/controller/cluster/leaked/credentialcheckrule"
transitsiteacl "github.com/prolixalias/provider-cloudflare/internal/controller/cluster/magic/transitsiteacl"
bucketlock "github.com/prolixalias/provider-cloudflare/internal/controller/cluster/r2/bucketlock"
validationsettings "github.com/prolixalias/provider-cloudflare/internal/controller/cluster/schema/validationsettings"
trustaccessidentityprovider "github.com/prolixalias/provider-cloudflare/internal/controller/cluster/zero/trustaccessidentityprovider"
trustaccesspolicy "github.com/prolixalias/provider-cloudflare/internal/controller/cluster/zero/trustaccesspolicy"
turnapp "github.com/prolixalias/provider-cloudflare/internal/controller/cluster/calls/turnapp"
image "github.com/prolixalias/provider-cloudflare/internal/controller/cluster/cloudflare/image"
ssl "github.com/prolixalias/provider-cloudflare/internal/controller/cluster/custom/ssl"
zonetransfersincoming "github.com/prolixalias/provider-cloudflare/internal/controller/cluster/dns/zonetransfersincoming"
securityimpersonationregistry "github.com/prolixalias/provider-cloudflare/internal/controller/cluster/email/securityimpersonationregistry"
forplatformsdispatchnamespace "github.com/prolixalias/provider-cloudflare/internal/controller/cluster/workers/forplatformsdispatchnamespace"
trustdnslocation "github.com/prolixalias/provider-cloudflare/internal/controller/cluster/zero/trustdnslocation"
lockdown "github.com/prolixalias/provider-cloudflare/internal/controller/cluster/zone/lockdown"
smartrouting "github.com/prolixalias/provider-cloudflare/internal/controller/cluster/argo/smartrouting"
tlssetting "github.com/prolixalias/provider-cloudflare/internal/controller/cluster/hostname/tlssetting"
cacertificate "github.com/prolixalias/provider-cloudflare/internal/controller/cluster/origin/cacertificate"
normalizationsettings "github.com/prolixalias/provider-cloudflare/internal/controller/cluster/url/normalizationsettings"
trustaccessapplication "github.com/prolixalias/provider-cloudflare/internal/controller/cluster/zero/trustaccessapplication"
trustgatewaycertificate "github.com/prolixalias/provider-cloudflare/internal/controller/cluster/zero/trustgatewaycertificate"
trustriskbehavior "github.com/prolixalias/provider-cloudflare/internal/controller/cluster/zero/trustriskbehavior"
ipprefix "github.com/prolixalias/provider-cloudflare/internal/controller/cluster/byo/ipprefix"
database "github.com/prolixalias/provider-cloudflare/internal/controller/cluster/d1/database"
transitconnector "github.com/prolixalias/provider-cloudflare/internal/controller/cluster/magic/transitconnector"
providerconfig "github.com/prolixalias/provider-cloudflare/internal/controller/cluster/providerconfig"
bucket "github.com/prolixalias/provider-cloudflare/internal/controller/cluster/r2/bucket"
cache "github.com/prolixalias/provider-cloudflare/internal/controller/cluster/tiered/cache"
analyticssite "github.com/prolixalias/provider-cloudflare/internal/controller/cluster/web/analyticssite"
trustriskscoringintegration "github.com/prolixalias/provider-cloudflare/internal/controller/cluster/zero/trustriskscoringintegration"
ownershipchallenge "github.com/prolixalias/provider-cloudflare/internal/controller/cluster/logpush/ownershipchallenge"
policy "github.com/prolixalias/provider-cloudflare/internal/controller/cluster/notification/policy"
trustaccessservicetoken "github.com/prolixalias/provider-cloudflare/internal/controller/cluster/zero/trustaccessservicetoken"
trustdlppredefinedprofile "github.com/prolixalias/provider-cloudflare/internal/controller/cluster/zero/trustdlppredefinedprofile"
cachevariants "github.com/prolixalias/provider-cloudflare/internal/controller/cluster/zone/cachevariants"
dnssec "github.com/prolixalias/provider-cloudflare/internal/controller/cluster/zone/dnssec"
dnssettings "github.com/prolixalias/provider-cloudflare/internal/controller/cluster/account/dnssettings"
token "github.com/prolixalias/provider-cloudflare/internal/controller/cluster/account/token"
zonetransferstsig "github.com/prolixalias/provider-cloudflare/internal/controller/cluster/dns/zonetransferstsig"
routingaddress "github.com/prolixalias/provider-cloudflare/internal/controller/cluster/email/routingaddress"
bucketlifecycle "github.com/prolixalias/provider-cloudflare/internal/controller/cluster/r2/bucketlifecycle"
validationoperationsettings "github.com/prolixalias/provider-cloudflare/internal/controller/cluster/schema/validationoperationsettings"
liveinput "github.com/prolixalias/provider-cloudflare/internal/controller/cluster/stream/liveinput"
hold "github.com/prolixalias/provider-cloudflare/internal/controller/cluster/zone/hold"
dnssettingsinternalview "github.com/prolixalias/provider-cloudflare/internal/controller/cluster/account/dnssettingsinternalview"
snippets "github.com/prolixalias/provider-cloudflare/internal/controller/cluster/cloudflare/snippets"
scanning "github.com/prolixalias/provider-cloudflare/internal/controller/cluster/content/scanning"
shieldpolicy "github.com/prolixalias/provider-cloudflare/internal/controller/cluster/page/shieldpolicy"
agentblockingrule "github.com/prolixalias/provider-cloudflare/internal/controller/cluster/user/agentblockingrule"
trustdevicedefaultprofile "github.com/prolixalias/provider-cloudflare/internal/controller/cluster/zero/trustdevicedefaultprofile"
trusttunnelcloudflared "github.com/prolixalias/provider-cloudflare/internal/controller/cluster/zero/trusttunnelcloudflared"
member "github.com/prolixalias/provider-cloudflare/internal/controller/cluster/account/member"
shieldschema "github.com/prolixalias/provider-cloudflare/internal/controller/cluster/api/shieldschema"
originpullssettings "github.com/prolixalias/provider-cloudflare/internal/controller/cluster/authenticated/originpullssettings"
onerequestpriority "github.com/prolixalias/provider-cloudflare/internal/controller/cluster/cloudforce/onerequestpriority"
bucketsippy "github.com/prolixalias/provider-cloudflare/internal/controller/cluster/r2/bucketsippy"
trustdevicedefaultprofilecertificates "github.com/prolixalias/provider-cloudflare/internal/controller/cluster/zero/trustdevicedefaultprofilecertificates"
filter "github.com/prolixalias/provider-cloudflare/internal/controller/cluster/cloudflare/filter"
organization "github.com/prolixalias/provider-cloudflare/internal/controller/cluster/cloudflare/organization"
project "github.com/prolixalias/provider-cloudflare/internal/controller/cluster/pages/project"
hostnameweb3 "github.com/prolixalias/provider-cloudflare/internal/controller/cluster/web3/hostname"
zonetransfersacl "github.com/prolixalias/provider-cloudflare/internal/controller/cluster/dns/zonetransfersacl"
policywebhooks "github.com/prolixalias/provider-cloudflare/internal/controller/cluster/notification/policywebhooks"
profile "github.com/prolixalias/provider-cloudflare/internal/controller/cluster/organization/profile"
tls "github.com/prolixalias/provider-cloudflare/internal/controller/cluster/total/tls"
trustaccessaicontrolsmcpportal "github.com/prolixalias/provider-cloudflare/internal/controller/cluster/zero/trustaccessaicontrolsmcpportal"
trustaccessmtlshostnamesettings "github.com/prolixalias/provider-cloudflare/internal/controller/cluster/zero/trustaccessmtlshostnamesettings"
trustdlppredefinedentry "github.com/prolixalias/provider-cloudflare/internal/controller/cluster/zero/trustdlppredefinedentry"

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