package handler

import (
	"html/template"
	"net/http"
)

var page = template.Must(template.New("graphiql").Parse(`<!DOCTYPE html>
<html>
<head>
	<meta charset=utf-8/>
	<meta name="viewport" content="user-scalable=no, initial-scale=1.0, minimum-scale=1.0, maximum-scale=1.0, minimal-ui">
	<link rel="shortcut icon" href="https://graphcool-playground.netlify.com/favicon.png">
	<link rel="stylesheet" href="//cdn.jsdelivr.net/npm/graphql-playground-react@{{ .version }}/build/static/css/index.css"/>
	<link rel="shortcut icon" href="//cdn.jsdelivr.net/npm/graphql-playground-react@{{ .version }}/build/favicon.png"/>
	<script src="//cdn.jsdelivr.net/npm/graphql-playground-react@{{ .version }}/build/static/js/middleware.js"></script>
	<title>{{.title}}</title>
</head>
<body>
<style type="text/css">
	html { font-family: "Open Sans", sans-serif; overflow: hidden; }
	body { margin: 0; background: #172a3a; }
</style>
<a
    style="
    	position: absolute;
    	z-index: 9999;
    	margin-left: 82vw;
    	margin-right: 32px;
    	margin-top: 22px;
    	font-family: 'Arial';
    	font-style: normal;
    	font-weight: bold;
    	font-size: 12px;
    	line-height: 14px;
    	color: #4b5660;
	"
    target="_blank"
    rel="noreferrer"
    href="https://www.ntucenterprise.sg/contactus/vulnerability-disclosure-programme"
    alt="Vulnerability Disclosure Programme"
>
    Vulnerability Disclosure Programme
</a>
<div id="root"/>
<script type="text/javascript">
	window.addEventListener('load', function (event) {
		const root = document.getElementById('root');
		root.classList.add('playgroundIn');
		const wsProto = location.protocol == 'https:' ? 'wss:' : 'ws:'
		GraphQLPlayground.init(root, {
			endpoint: location.protocol + '//' + location.host + '{{.endpoint}}',
			subscriptionsEndpoint: wsProto + '//' + location.host + '{{.endpoint }}',
			settings: {
				'request.credentials': 'same-origin'
			}
		})
	})
</script>
</body>
</html>
`))

func Playground(title string, endpoint string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		err := page.Execute(w, map[string]string{
			"title":    title,
			"endpoint": endpoint,
			"version":  "1.7.8",
		})
		if err != nil {
			panic(err)
		}
	}
}
