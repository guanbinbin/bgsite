<title>Login</title>
<div class="container" >
    {{.msg}}
    <br />
    {{if .is_login}}
        <a href="/">Домой</a>
    {{else}}
        <a href="/session">Назад</a>
    {{end}}
</div>