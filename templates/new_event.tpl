{{.User.Name}} hat eine neue Veranstaltung erstellt:
 {{if .Event.Public}}:globe_showing_europe_africa:{{ else }}:unlock:{{end}} <a href="{{ .Event.Url }}">{{ .Event.Name }}</a>
 Am {{.Event.Date.Format "02.01.06"}} um {{.Event.Date.Format "03:04"}} Uhr @ {{.Event.Location }}
