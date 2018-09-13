<title>Calc</title>
<div class="container" >
    <form action="/calc" method="POST" >
            <div class="row">
                <div class="col s3">
                    <input id = "firsta" type="number" name="firsta" />  <label for="firsta"> Первое число </label>
                </div>
                <div class="col s1">
                    <input id = "action" type="text" name="action" />  <label for="action"> +, -, /, * </label>
                </div>
                <div class="col s3">
                    <input id = "seconda" type="number" name="seconda" />  <label for="seconda"> Второе число </label>
                </div>
                <div class="col s1">
                    =
                </div>
                <div class="col s4">
                {{.reply}}
                </div>
            </div>
        <button class="btn waves-effect waves-light" type="submit" name="action">Submit
            <i class="material-icons right">send</i>
        </button>
    </form>
</div>