<title>Index</title>
<div class="container" >
{{.msg}} {{.name}}  <br/> <br />
    {{if .is_login}}
    <form action = "" method="POST">
        <button class="btn waves-effect waves-light" type="submit" name="quit" value="quit">Выйти
            <i class="material-icons left">power_settings_new</i>
        </button>
    </form>
    {{end}}
</div>