<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <title>预防跨站脚本</title>
    <style type="text/css">
        body {
            padding: 10px 20px;
        }
        dl,dt,dd {
            padding: 0;
            margin: 0;
        }
        dt,dd {
            color: #333333;
            font-size: 14px;
        }
        p input {
            width: 100px;
        }

    </style>
</head>
<body>
<form action="/xss" method="post">
    <dl>
        <dt>正文：</dt>
        <dd>
            <textarea style="width:400px; height:200px" name="code"></textarea>
        </dd>
    </dl>
    <p><input type="submit" value="提交"></p>
    </form>
</body>
</html>
