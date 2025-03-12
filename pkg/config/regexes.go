package config

import (
	_struct "github.com/dsecuredcom/yataf/pkg/struct"
	"strings"
)

func GetRegExConfig(Types []string) _struct.RegExConfig {
	Config := _struct.RegExConfig{
		Elements: []_struct.RegEx{
			{Type: "Cloud", Name: "AWS regions", Matcher: `((us|eu|ap|sa|ca|me|af)-(east|west|central|north|south).*)`},
			{Type: "Cloud", Name: "AWS Lambda", Matcher : `execute-api.*?aws.*?com` },
			{Type: "Cloud", Name: "AWS Cognito", Matcher : `cognito.*?aws.*?com`  },
			{Type: "Cloud", Name: "AWS Cognito pool", Matcher : `(?im)userPool` },
			{Type: "Cloud", Name: "Azure Tenant subdomain", Matcher : `[A-Za-z0-9-]+.onmicrosoft.com` },
			{Type: "Cloud", Name: "Azure Tenant Auth", Matcher : `[A-Za-z0-9-]+.ciamlogin.com` },
			{Type: "Cloud", Name: "Azure clientId", Matcher : `clientId:.*` },
			{Type: "Credentials", Name: "User/Pass Case#1", Matcher: `(?im)['|"](user|username|mail|email|access|ident|login)['|"](\s*?):(\s*?)['|"].*?['|"](\s*?),(\s*?)['|"](pass|password|passwordpassive|authkey|credentials|access_key|apikey|secret|access)['|"](\s*?):(\s*?)['|"].*?['|"]`},
			{Type: "Credentials", Name: "JS const/var", Matcher: `(?i)(var|const|let)(\s*?)(pass|password|passwordpassive|authkey|credentials|access_key|apikey|secret|access)(\s*?)(=|:)(\s*?)['|"].*?['|"]`},
			{Type: "Credentials", Name: "Slack Token", Matcher: `(xox[p|b|o|a]-[0-9]{12}-[0-9]{12}-[0-9]{12}-[a-z0-9]{32})`},
			{Type: "Credentials", Name: "RSA private key", Matcher: `-----BEGIN RSA PRIVATE KEY-----`},
			{Type: "Credentials", Name: "SSH (DSA) private key", Matcher: `-----BEGIN DSA PRIVATE KEY-----`},
			{Type: "Credentials", Name: "SSH (EC) private key", Matcher: `-----BEGIN EC PRIVATE KEY-----`},
			{Type: "Credentials", Name: "PGP private key block", Matcher: `-----BEGIN PGP PRIVATE KEY BLOCK-----`},
			{Type: "Credentials", Name: "Open SSH Private Key", Matcher: `-----BEGIN OPENSSH PRIVATE KEY-----`},
			{Type: "Credentials", Name: "Amazon Related Key ID", Matcher: `AKIA[0-9A-Z]{16}`},
			{Type: "Credentials", Name: "Amazon MWS Auth Token", Matcher: `amzn\.mws\.[0-9a-f]{8}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{12}`},
			{Type: "Credentials", Name: "Amazon Secret Ident", Matcher: `(?i)aws_secret_access_key|aws_secret_key|aws\.secret_key|aws\.secret_access_key`},
			{Type: "Credentials", Name: "Amazon Secret Key", Matcher: `(?im)['|"]aws_secret_access_key['|"](\s*?):(\s*?)['|"].*?['|"]`},
			{Type: "Credentials", Name: "Facebook Access Token", Matcher: `EAACEdEose0cBA[0-9A-Za-z]+`},
			{Type: "Credentials", Name: "Facebook OAuth", Matcher: `[f|F][a|A][c|C][e|E][b|B][o|O][o|O][k|K]['|"](\s*?):(\s*?)['|"][0-9a-f]{32}['|"]`},
			{Type: "Credentials", Name: "GitHub", Matcher: `[g|G][i|I][t|T][h|H][u|U][b|B]['|"](\s*?):(\s*?)['|"][0-9a-zA-Z]{35,40}['|"]`},
			{Type: "Credentials", Name: "Generic API Key", Matcher: `[a-zA-Z0-9\._]+[a|A][p|P][i|I][_]?[k|K][e|E][y|Y]['|"](\s*?):(\s*?)['|"][0-9a-zA-Z]{32,45}['|"]`},
			{Type: "Credentials", Name: "Generic Secret #1", Matcher: `[a-zA-Z0-9\._]+[s|S][e|E][c|C][r|R][e|E][t|T]['|"](\s*?):(\s*?)['|"][0-9a-zA-Z-=_.]{16,45}['|"]`},
			{Type: "Credentials", Name: "Generic Secret #2", Matcher: `[a-zA-Z0-9\._]+[s|S][e|E][c|C][r|R][e|E][t|T](\s*?)=(\s*?)['|"][0-9a-zA-Z-=_.]{16,45}['|"]`},
			{Type: "Credentials", Name: "Google Related Key", Matcher: `AIza[0-9A-Za-z\-_]{35}`},
			{Type: "Credentials", Name: "Google OAuth", Matcher: `[0-9]+-[0-9A-Za-z_]{32}\.apps\.googleusercontent\.com`},
			{Type: "Credentials", Name: "Google (GCP) Service-account", Matcher: `\"type\": \"service_account\"`},
			{Type: "Credentials", Name: "Google OAuth Access Token", Matcher: `ya29\.[0-9A-Za-z\-_]+`},
			{Type: "Credentials", Name: "Heroku API Key", Matcher: `[h|H][e|E][r|R][o|O][k|K][u|U].*[0-9A-F]{8}-[0-9A-F]{4}-[0-9A-F]{4}-[0-9A-F]{4}-[0-9A-F]{12}`},
			{Type: "Credentials", Name: "MailChimp API Key", Matcher: `[0-9a-f]{32}-us[0-9]{1,2}`},
			{Type: "Credentials", Name: "Mailgun API Key", Matcher: `key-[0-9a-zA-Z]{32}`},
			{Type: "Credentials", Name: "Password in URL", Matcher: `[a-zA-Z]{3,10}://[^/\s:@]{3,20}:[^/\s:@]{3,20}@.{1,100}[\"'\s]`},
			{Type: "Credentials", Name: "PayPal Braintree Access Token", Matcher: `access_token\$production\$[0-9a-z]{16}\$[0-9a-f]{32}`},
			{Type: "Credentials", Name: "Picatic API Key", Matcher: `sk_live_[0-9a-z]{32}`},
			{Type: "Credentials", Name: "Slack Webhook", Matcher: `https://hooks.slack.com/services/T[a-zA-Z0-9_]{8}/B[a-zA-Z0-9_]{8}/[a-zA-Z0-9_]{24}`},
			{Type: "Credentials", Name: "Potential Slack related #1", Matcher: `slack[_|.|-](access|api|secret|token)[_|.|-](token|key|id)`},
			{Type: "Credentials", Name: "Stripe API Key", Matcher: `sk_live_[0-9a-zA-Z]{24}`},
			{Type: "Credentials", Name: "Stripe Restricted API Key", Matcher: `rk_live_[0-9a-zA-Z]{24}`},
			{Type: "Credentials", Name: "Square Access Token", Matcher: `sq0atp-[0-9A-Za-z\-_]{22}`},
			{Type: "Credentials", Name: "Square OAuth Secret", Matcher: `sq0csp-[0-9A-Za-z\-_]{43}`},
			{Type: "Credentials", Name: "Twilio API Key", Matcher: `SK[0-9a-fA-F]{32}`},
			{Type: "Credentials", Name: "Twitter Access Token", Matcher: `[t|T][w|W][i|I][t|T][t|T][e|E][r|R]['|"](\s*?):(\s*?)[1-9][0-9]+-[0-9a-zA-Z]{40}`},
			{Type: "Credentials", Name: "Twitter OAuth", Matcher: `[t|T][w|W][i|I][t|T][t|T][e|E][r|R]['|"](\s*?):(\s*?)['|"][0-9a-zA-Z]{35,44}['|"]`},
			{Type: "Credentials", Name: "Shared Access Key", Matcher: `(?i)SharedAccessKey\=`},
			{Type: "Credentials", Name: "Authorization Header #1", Matcher: `(?i)["|'](Basic|Bearer|Digest|HOBA|Mutual|Negotiate|OAuth|SCRAM-SHA-1|SCRAM-SHA-256|vapid) [a-zA-Z0-9-\._~\+\/]+=*["|']`},
			{Type: "Credentials", Name: "Authorization Header #2", Matcher: `setHeader\(["|']Authorization["|']`},
			{Type: "Credentials", Name: "Authorization Header #3", Matcher: `(?i)Authorization:\s*(Basic|Bearer|Digest|HOBA|Mutual|Negotiate|OAuth|SCRAM-SHA-1|SCRAM-SHA-256|vapid) [a-zA-Z0-9-\._~\+\/]+=*`},
			{Type: "Credentials", Name: "Seen in the past tokens", Matcher: `(?i)['|"](DISCOVERY_IAM_APIKEY|appPassword|slackToken|slack_signing_secret|watson_assistant_api_key|pythonPassword)['|"]`},
			{Type: "Credentials", Name: "Secret indicator with _", Matcher: `(?i)['|"][a-zA-Z0-9\-]+[\.|\-|_](access-key|password|apikey|secret|access_key|secret-key|pwd|passwd|appsecret|app_secret)['|"](\s*?):(\s*?)['|"].*?['|"](\s*?)`},
			{Type: "Credentials", Name: "Credentials in URL", Matcher: `(?i)['|"](http|https|ftp|postgresql|smtp|ssh|mongodb|redis|mysql|mssql|jdbc|odbc)://[a-zA-Z0-9\_-]+:.*@[a-z0-9]+`},
			{Type: "Credentials", Name: "Credentials in module.exports", Matcher: `(?im)module.exports\s*?=\s*?.*{\s*?.*?(admin|secret|password|credentials).*?\s*?}`},
			{Type: "Credentials", Name: "Escaped credentials", Matcher: `(?im)\\['|"](admin|user|client|users|)[_\.]?(pass|password|passwd|secret|credentials|token)\\['|"]\s*?:\s*?\\['|"].*?\\['|"]`},
			{Type: "Credentials", Name: "Potential credentials in html table", Matcher: `(?i)<td>.*?(password|secret|token|keyfile|passwd|credentials|admin).*?</td>`},
			{Type: "Credentials", Name: "Credentials with \"this.\" prepended", Matcher: `(?i)this\.(apikey|secret|password|token|passwd|credentials|access_key|secret_key|pwd|appsecret|app_secret)\s*?=\s*?['|"].*?['|"]`},
			{Type: "Credentials", Name: "API Tokens with high entropy", Matcher: `(?i)token:\s*?['|"][a-zA-Z0-9]+['|"]`},

			{Type: "Urls-Paths", Name: "Cloudinary", Matcher: `cloudinary://.*`},
			{Type: "Urls-Paths", Name: "Firebase URL", Matcher: `[a-zA-Z0-9-]+\.firebaseio\.com`},
			{Type: "Urls-Paths", Name: "Urls and paths #1", Matcher: `(?i)[a-zA-Z]\.(get|post|fetch|patch|delete|option|put|ajax)\((\s*?)['|"][a-zA-Z0-9-_\/:\?]['|"]`},
			{Type: "Urls-Paths", Name: "Urls and paths #2", Matcher: `(?i)<(a|script).*(href|src)=['|"]\/[a-zA-Z0-9]+.*['|"].*>.*<\/(a|script)>`},
			{Type: "Urls-Paths", Name: "Urls and paths #3", Matcher: `(?i)<(a|script).*href=['|"]http[s]?:\/\/[a-zA-Z0-9\.\-\/?=]+['|"].*>.*<\/(a|script)>`},
			{Type: "Urls-Paths", Name: "Urls and paths #4", Matcher: `(?i)<(a|script).*href=['|"]\/\/[a-zA-Z0-9\.\-\/?=]+['|"].*>.*<\/(a|script)>`},
			{Type: "Urls-Paths", Name: "Urls and paths #5", Matcher: `(?i)<base.*href=['|"]\/.*\/>`},
			{Type: "Urls-Paths", Name: "Urls and paths #6", Matcher: `(?i)(url|path|file)['|"](\s*?):(\s*?)['|"].*?['|"]`},
			{Type: "Urls-Paths", Name: "Urls and paths #8", Matcher: `(?im)url: ['|"]\/.*?['|"]`},
			{Type: "Urls-Paths", Name: "Urls and paths #9", Matcher: `(?im)url: [a-zA-Z0-9]+(\s*?)\+(\s*?)['|"]\/.*?['|"]`},
			{Type: "Urls-Paths", Name: "Urls and paths #10", Matcher: `(?im)(axios|jquery|instance|http|\$).(get|post|delete|put|ajax)\(['|"](https|http|\/|\.\/).*?['|"]`},
			{Type: "Urls-Paths", Name: "Urls and paths #11", Matcher: `(?im)(axios|jquery|instance|http|\$).(get|post|delete|put|ajax)\([a-zA-Z0-9._]+,`},
			{Type: "Urls-Paths", Name: "Urls and paths #12", Matcher: `(?i)fetch\(['|"]\/.*?['|"],`},
			{Type: "Urls-Paths", Name: "Urls and paths #13", Matcher: `(?im)['|"][a-zA-Z0-9_\.]+['|"](\s*?):(\s*?)['|"](http[s]?:|//|\./).*?['|"]`},
			{Type: "Urls-Paths", Name: "Urls and paths #14", Matcher: `(?im)(axios|jquery|instance|http|\$)\.(get|post|delete)\(.*?\)(\s*?)\.then\(function`},
			{Type: "Urls-Paths", Name: "Urls and paths #15", Matcher: `(?im)\/(api|v1|v2|v3|alpha|apis)\/[a-zA-Z0-9_\-\/\?\${}\.]{3,75}`},
			{Type: "Urls-Paths", Name: "Urls and paths #16", Matcher: `(?im)(req|request)\.open\(['|"](GET|POST)['|"],\s*?['|"]\/.*?['|"]`},
			{Type: "Urls-Paths", Name: "Urls and paths #17", Matcher: `(?i)(const|let|var)\s*?[a-zA-Z0-9_]+\s*?=\s*?['|"](https|http|\/|\.\/).*?['|"]`},
			{Type: "Urls-Paths", Name: "Urls and paths #18", Matcher: `(?i)new Websocket\(['|"](https|http|\/|\.\/).*?['|"]\)`},
			// @TODO: In JS one can also use backticks instead of " or '... unfortunately it is not possible
			//		  to escape backticks (`) within backticks... The only solution is to use double quotes like below
			//		  new WebSocket(`ws://localhost:8080`) vs new WebSocket("ws://localhost:8080")
			{Type: "Urls-Paths", Name: "Urls and paths #19", Matcher: "(?i)new WebSocket\\(`.*?`"}, // <--- see difference here
			{Type: "Urls-Paths", Name: "Urls and paths #20", Matcher: `(?im)(axios|jquery|instance|http|\$).(get|post|delete|put|ajax)\([a-zA-Z0-9_]+\s*?\+\s*?['|"].*?['|"]`},
			{Type: "Urls-Paths", Name: "Urls and paths #21", Matcher: `(?im)(axios|jquery|instance|http|\$).(get|post|delete|put|ajax)\(['|"]\/.*?\s*?\+\s*?[a-zA-Z0-9_]+\s*?\+\s*?['|"]`},
			{Type: "Urls-Paths", Name: "Urls and paths #22", Matcher: `(?i)(path|pathname|file)\s*?(:|=)\s*?['|"]\/.*?['|"]`},
		},
	}

	if IsTypeAllowed(Types, "all") {
		return Config
	}

	FilteredConfig := _struct.RegExConfig{
		Elements: []_struct.RegEx{},
	}

	for _, Element := range Config.Elements {
		DesiredType := strings.ToLower(Element.Type)
		if IsTypeAllowed(Types, DesiredType) {
			FilteredConfig.Elements = append(FilteredConfig.Elements, Element)
		}
	}

	return FilteredConfig
}

func IsTypeAllowed(Types []string, DesiredType string) bool {

	for _, Type := range Types {
		Type = strings.ToLower(Type)
		DesiredType = strings.TrimSpace(DesiredType)
		if Type == DesiredType {
			return true
		}
	}
	return false
}
