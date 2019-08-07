package main

var templIndex string = `
<!-- <h1>Статус спутниковых терминалов</h1> -->
<div id="maintableID">
<table class="maintable">
    <tr>
        {{/* {{range $k, $v := .}} {{if eq $k "IP адрес"}}
        <th>{{$k}}</th>
        <th class="small-font">{{$v.Note}}</th>
        <th>{{$v.SiteID}}</th>
        <th class="small-font">{{$v.StatusNow.Code}}</th>
        {{range $s := $v.StatusOfHour}}
        <th>
            {{$s.Code}}
        </th>
        {{end}} {{end}} {{end}} */}}
        <th>IP адрес</th>
        <th class="small-font">Описание</th>
        <th>SiteID</th>
        <th class="small-font">тек.стат</th>
        <th>00</th><th>01</th><th>02</th><th>03</th><th>04</th><th>05</th><th>06</th><th>07</th><th>08</th><th>09</th><th>10</th><th>11</th><th>12</th><th>13</th><th>14</th><th>15</th><th>16</th><th>17</th><th>18</th><th>19</th><th>20</th><th>21</th><th>22</th><th>23</th>

    </tr>
    {{range $k, $v := .}} {{if ne $k "IP адрес"}}
    <tr>
        <td><a class="b-href" href="http://{{$k}}/">{{$k}}</a></td>
        <td class="small-font">
            <a class="b-href" href="/edit?IP={{$k}}">
                <img width="18" height="18" alt="Редактирование" src="data:image/png;base64,iVBORw0KGgoAAAANSUhEUgAAADIAAAAyCAYAAAAeP4ixAAAAAXNSR0IArs4c6QAAAARnQU1BAACxjwv8YQUAAAAJcEhZcwAADsMAAA7DAcdvqGQAAAF1SURBVGhD7Zi9SsRAFEZT2YlPoOADWNhZaKtWPoE2oo+hiK3aijaCraAoWGrpD76LtZ16PtyBECfJ7BqYO3IPHFgWstwTZmZ3UzmOUyRreImv+IxnuITFMI13+NWigqbQNDOoux8LqHuOZkmNCJpdZrsYG7jNUzTLAcaGjvmCpkmNeUMTaE8s/7z8RUrMBWYnbOwPXNcbEfpiVjErzdNpkpgbzErbETtOjK7X52SjLSKYEmM+ItgVs41FRAS7YrIxbkTQVMykEUHFrGBW/hohdX1ReyKmRwyBR4w0EfGEseFS9Ygh8IiRHjEE/yJCzKN+Tusf2ifGBu3SRESTTYwN26bJiIAeycSGbqrlaDZCHGFs8LpmIhaw7en3Ih6ilo3pCHGC99j3KH8L6weAueX0iBosJeYWTUaIdwx3uS9mB01GzGKISInR94y5CLGBzZC+GJPsYSxEFhVzhc0A7ZkHPEYdzUWgI/Ua91HLbA4dx3GcAaiqb1gXvuW/ej2aAAAAAElFTkSuQmCC">                </a>
            <a class="b-href" href="/edit?IP={{$k}}">{{$v.Note}}</a>
        </td>
        <td>{{$v.SiteID}}</td>
        <!-- Текущий статус -->
        {{if eq $v.StatusNow.Code "√"}}
            <td class="bggreen">√</td>
        {{else if eq $v.StatusNow.Code "X"}}
            <td class="bgred">X</td>
        {{else}}
        <td class="bggrey">O</td>
        {{end}}
        <!-- Основная таблица статусов спутниковых станций -->
        {{range $s := $v.StatusOfHour}}
            {{if eq $s.NumFail 0}}
                {{if eq $s.NumPass 0}}
                    <td class="bggrey">O</td>
                {{else if ne $s.NumPass 0}}
                    <td class="bggreen">√</td>
                {{end}}
            {{else if ne $s.NumFail 0}}
                {{if eq $s.NumPass 0}}
                    <td class="bgred">X</td>
                {{else if ne $s.NumPass 0}}
                    {{if gt $s.NumPass $s.NumFail}}
                        <td class="bgyellowgreen">√</td>
                    {{else if lt $s.NumPass $s.NumFail}}
                        <td class="bgpalevioletred">X</td>
                    {{else}}
                        <td class="bgyellow">√</td>
                    {{end}}
                {{end}}

            {{end}}
        {{end}} 
    </tr>
    {{end}} {{end}}
</table>
</div><div></div>

<div id="legendID">
<table class="legend">
    <tr>
        <td class="bggreen">√</td>
        <td>- Станция онлайн</td>
        <td class="bgred">X</td>
        <td>- Станция офлайн</td>
        <td class="bggrey">O</td>
        <td>- Станция не проверялась</td>
    </tr>
</table>
</div>
`
var templHeader string = `

<!DOCTYPE html>
<html lang="ru">

<head>
    <meta http-equiv="Refresh" content="300" />
    <meta charset="UTF-8">
    <meta http-equiv="X-UA-Compatible" content="IE=edge">
    <meta name="viewport" content="width=device-width, initial-scale=1, shrink-to-fit=no">
    <title>Состояние спутниковых терминалов</title>
    {{/* <link href="/assets/css/style.css" rel="stylesheet"> */}}
    <style type="text/css">
        *{
            margin: 0px;
            padding: 0px;
        }

        body{
            background-color: #f1f1f1
        
        }

        header {
            display: grid;
            grid-template-columns: 1fr 1fr 1fr 9fr;
			margin-top: 1rem;
        }

        .submenu{
            font-size: 0.9em;
        }

        .submenu li{
            padding:0.3rem;
        }

        .main{
            display: grid;
            grid-template-columns: 1fr 7fr;
            margin: 1rem 0;
        }

        .maintable {
            empty-cells: show;
            margin: auto;
            width: 96%;
            border: 1px solid black;
            border-collapse: collapse;
            box-shadow: black 8px 7px 9px 1px;
        }

        .maintable td, th {
            border: 1px solid black;
            font-size: 0.7em;
        }
        .legend td, th {
            border: 1px solid black;
            font-size: 0.7em;
        }

        .maintable a{
            color: #0c0229;
        }

        #legendID{
            padding-top: 2rem;
        }


        .note{
            font-size: 0.8em
        }

        .small-font {
            font-size: 0.7em
        }


        .bggreen {
            background-color: green;
            text-align: center;
        }

        .bgred {
            text-align: center;
            background-color: red;
        }

        .bgyellow {
            text-align: center;
            background-color: yellow;
        }

        .bgyellowgreen {
            text-align: center;
            background-color: yellowgreen;
        }

        .bgpalevioletred {
            text-align: center;
            background-color: palevioletred;
        }

        .bggrey {
            text-align: center;
            background-color: grey;
        }

        .legend {
            empty-cells: show;
            margin: auto;
            align-self: center;
            width: 50%;
        }

        footer {
            padding-top: 1rem;
            padding-bottom: 1rem;
            background-color: #02020245;
            display: grid;
            margin: auto;
            grid-template-columns: 10% 30% 26% 34%;
            color: #ffffff;
            align-content: center;
            align-items: center;
        }

        footer a{
            color: #ffffff;
        }

        </style>
    {{/* <link href="/assets/css/panda.css" rel="stylesheet"> */}}
    <style type="text/css">
                /* footer ul.footer-menu>li {
            display: inline-block;
            padding: 40px;
        } */

        .panda {
            position: relative;
            display: block;
            width: 30px;
            height: 30px;
        }

        .head {
            position: absolute;
            top: 20%;
            left: 20%;
            width: 60%;
            height: 60%;
            background: #A6BECF;
            border-radius: 50%;
            z-index: 2;
        }

        .head-copy {
            width: 100%;
            height: 100%;
            position: absolute;
            background: #A6BECF;
            border-radius: 50%;
            z-index: 2;
        }

        .ear-left {
            position: absolute;
            width: 60%;
            height: 65%;
            left: -20%;
            top: 5%;
            background: #A6BECF;
            border-radius: 50%;
            z-index: 1;
        }

        .ear-right {
            position: absolute;
            width: 60%;
            height: 65%;
            right: -20%;
            top: 5%;
            background: #A6BECF;
            border-radius: 50%;
            z-index: 1;
        }

        .inner-ear {
            position: absolute;
            border-radius: 50%;
            width: 80%;
            height: 80%;
            top: 10%;
            left: 10%;
            background: #819CAF;
        }

        .eye-left {
            position: absolute;
            background: white;
            width: 30%;
            height: 33%;
            top: 25%;
            left: 21%;
            border-radius: 50%;
            z-index: 3;
        }

        .eye-right {
            position: absolute;
            background: white;
            width: 30%;
            height: 33%;
            top: 25%;
            right: 21%;
            border-radius: 50%;
            z-index: 2;
        }

        .pupil {
            position: absolute;
            width: 28%;
            height: 30%;
            top: 35%;
            left: 36%;
            border-radius: 50%;
            background: #27354A;
        }

        .nose {
            position: absolute;
            background: #BE845F;
            width: 25%;
            height: 42.5%;
            left: 37%;
            top: 45%;
            border-radius: 50px;
            z-index: 4;
        }

        .hair-left {
            position: absolute;
            top: -8%;
            left: 30%;
            width: 20%;
            height: 20%;
            background: #A6BECF;
            -webkit-clip-path: polygon(50% 0%, 0% 100%, 100% 100%);
            clip-path: polygon(50% 0%, 0% 100%, 100% 100%);
        }

        .hair-right {
            position: absolute;
            top: -4%;
            left: 48%;
            width: 20%;
            height: 20%;
            background: #A6BECF;
            -webkit-clip-path: polygon(50% 0%, 0% 100%, 100% 100%);
            clip-path: polygon(50% 0%, 0% 100%, 100% 100%);
        }
    </style>

</head>

<body>
<header>
<a class="navbar-brand " href="/">
    <div class="panda">
        <!-- Круглая голова -->
        <div class="head">
            <!-- Круглая копия головы -->
            <div class="head-copy"></div>
            <!-- Левое ухо ~ светло-серое -->
            <div class="ear-left">
                <!-- Внутреннее ухо ~ тёмно-серое -->
                <div class="inner-ear"></div>
            </div>
            <!-- Правое ухо ~ светло-серое -->
            <div class="ear-right">
                <!-- Внутреннее ухо ~ Тёмно-серое -->
                <div class="inner-ear"></div>
            </div>
            <!-- Левый глаз ~ белый -->
            <div class="eye-left">
                <!-- Зрачок ~ чёрный -->
                <div class="pupil">
                </div>
            </div>
            <!-- Правый глаз ~ белый -->
            <div class="eye-right">
                <!-- Зрачок ~ чёрный -->
                <div class="pupil">
                </div>
            </div>
            <!-- Нос ~ коричневый -->
            <div class="nose">
            </div>
            <!-- Волосы ~ светло-серые -->
            <div class="hair-left"></div>
            <div class="hair-right"></div>
            <!-- Конец головы -->
        </div>
        <!-- Конец невидимого поля -->
    </div>
</a>
<a href="/">SatMon</a>

<a href="/">Главная </a>
</header>
<hr>
<div class="main">
<ul class="submenu">
    <li><a href="/">Перейти на главную </a></li>
    <li><a href="/checknow">Проверить сейчас</a></li>
    <li><a href="/write ">Добавить сервер</a></li>
    <li><a href="/report ">Открыть архив отчетов</a></li>
    <li><a href="/getreport ">Создать отчёт</a></li>
    <li><a href="/reloadconf">Перечитать конфигурацию</a></li>
    <li><a href="/saveconf">Сохранить конфигурацию</a></li>
    <li><a href="/loaddefaultconf">Загрузить конфигурацию по-умолчанию</a></li>
</ul>

`
var templFooter string = `
</div>
</body>
<br>
<br>
<footer>
    <div class="panda">
        <!-- Круглая голова -->
        <div class="head">
            <!-- Круглая копия головы -->
            <div class="head-copy"></div>
            <!-- Левое ухо ~ светло-серое -->
            <div class="ear-left">
                <!-- Внутреннее ухо ~ тёмно-серое -->
                <div class="inner-ear"></div>
            </div>
            <!-- Правое ухо ~ светло-серое -->
            <div class="ear-right">
                <!-- Внутреннее ухо ~ Тёмно-серое -->
                <div class="inner-ear"></div>
            </div>
            <!-- Левый глаз ~ белый -->
            <div class="eye-left">
                <!-- Зрачок ~ чёрный -->
                <div class="pupil">
                </div>
            </div>
            <!-- Правый глаз ~ белый -->
            <div class="eye-right">
                <!-- Зрачок ~ чёрный -->
                <div class="pupil">
                </div>
            </div>
            <!-- Нос ~ коричневый -->
            <div class="nose">
            </div>
            <!-- Волосы ~ светло-серые -->
            <div class="hair-left"></div>
            <div class="hair-right"></div>
            <!-- Конец головы -->
        </div>
        <!-- Конец невидимого поля -->
    </div>

    <body><a class="b-href" href="#">SatMon <i>© 2018</i> by Vladislav Vegner</a></body>
    <aside><a class="b-href" href="mailto:vegner.vs@uttist.ru">Email: vegner.vs@uttist.ru</a></aside>
    <aside>Создано с использованием <a class="b-href" href="https://icons8.com">Icon pack by Icons8</a></aside>
</footer>

</html>

`
var templEdit string = `
<div class="row">
    <div class="col-xs-2">
        {{ if .Post.Id }}
        <a href="/delete/{{.Post.Id}}">Delete</a>
        <br/>
        <a href="/view/{{.Post.Id}}" target="_blank">View</a> {{ end }}
    </div>
    <div class="col-xs-4">
        <form role="form" method="POST" action="/SavePost">
            <!-- <input type="hidden" name="id" value="{{.Post.Id}}" /> -->
            <div class="form-group">
                <label>Title</label>
                <input type="text" class="form-control" id="IP" name="IP" value="{{.IP}}" />
            </div>
            <div class="form-group">
                <label>Content</label>
                <textarea id="Note" name="Note">{{.Note}}</textarea>
            </div>
            <div class="form-group">
                <label>SiteID</label>
                <input type="text" class="form-control" id="SiteID" name="title" value="{{.SiteID}}" />
            </div>

            <button type="submit" class="btn btn-default">Submit</button>
        </form>
    </div>
    <div class="col-xs-6" id="md_html">
        {{. }}
    </div>
</div>

`
var templReport string = `
<p>Отчет сформирован и находится по <a href="/report/{{.}}" onclick="window.location(this.href); close_window(); return false;">ссылке</a>.</p>
<p>Все отчёты находятся <a href="/report/">тут</a>.</p>

`
var templWrite string = `
<div class="container">
    <div class="row">
        <div class="col-xs-2"></div>
        <div class="col-md-6">
            <form role="form" method="POST" action="/addserver">
                <div class="form-group">
                    <label>IP-адрес</label>
                    <input type="text" class="form-control" id="IP" name="IP" value="{{.IP}}" />
                </div>
                <div class="form-group">
                    <label>Описание</label>
                    <input type="text" class="form-control" id="Note" name="Note" value="{{.Note}}" />
                </div>
                <div class="form-group">
                    <label>SiteID</label>
                    <input type="text" class="form-control" id="SiteID" name="SiteID" value="{{.SiteID}}" />
                </div>
                <button type="submit" class="btn btn-default">Принять изменения</button>
            </form>
        </div>
        <div class="col-md-1">
            {{if .IP}}<a class="b-href" href="/delete?IP={{.IP}}">Delete</a>{{end}}
        </div>
    </div>
</div>

`
