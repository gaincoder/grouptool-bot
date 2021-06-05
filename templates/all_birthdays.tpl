:information::balloon: <b>Geburtstage:</b>

{{range .}}{{.Date.Format "02.01.06"}}  {{ .Name }}
{{end}}
