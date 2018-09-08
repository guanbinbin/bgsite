<!DOCTYPE html>

<html>
<!---------------------------------------->
<head>
    <link href="https://fonts.googleapis.com/icon?family=Material+Icons" rel="stylesheet">
    <!--Import materialize.css-->
    <link type="text/css" rel="stylesheet" href="static/css/materialize.css"  media="screen,projection"/>
    <!--Let browser know website is optimized for mobile-->
    <meta name="viewport" content="width=device-width, initial-scale=1.0"/>
    <title>My Page</title>
    <meta http-equiv="Content-Type" content="text/html; charset=utf-8">
</head>

<body>

<nav class="nav-extended">
    <div class="nav-wrapper">
        <a href="/" class="brand-logo">Website on GO</a>
    </div>
    <div class="nav-content">
        <ul class="tabs tabs-transparent">
            <li class="tab"><a href="/nameold">NameOld</a></li>
            <li class="tab"><a href="/calc">Calc</a></li>
            <li class="tab"><a href="/session">Session</a></li>
        </ul>
    </div>
</nav>
<br /> <br /> <br />
<!---------------------------------------->


    <div class="container">
        <div class="row">
            {{.key}} , {{.sessid}}
            <br /> <br />
            {{.name}} , {{.pass}}
        </div>
    </div>


<!--JavaScript at end of body for optimized loading-->
<script type="text/javascript" src="static/js/materialize.js"></script>
</body>
</html>