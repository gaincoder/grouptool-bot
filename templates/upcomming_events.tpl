{{if not .}}Sorry, leider sind gerade keine Veranstaltungen geplant :crying_face:{{ else }}:calendar: <b>Kommende Veranstaltungen:</b>

{{range .}}{{if .Public}}:globe_showing_europe_africa:{{ else }}:unlock:{{end}} {{.Date.Format "02.01"}}    <a href="{{ .Url }}">{{ .Name }}</a>
{{end}}{{end}}
