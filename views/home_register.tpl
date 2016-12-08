{{template "base.tpl" .}}

{{define "content"}}
<h1>Register!</h1>
<ul>
        {{range $key, $val := .errors}}
        <li>
    {{$key}} -- {{$val}}
                </li>
    {{end}}
</ul>
        <form method="post">
                {{.Form | renderform}}
                <input type="submit" name="submit" value="Register">
        </form>
        <p><a href="/">Go Back</a></p>
{{end}}