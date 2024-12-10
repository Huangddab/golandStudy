<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <title>上传文件</title>
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
    </style>
</head>
<body>
<form enctype="multipart/form-data" action="/upload" method="post">
    <dl>
        <dt>上传文件：</dt>
        <dd>
            <input type="file" name="uploadfile">
        </dd>
    </dl>
    <p><input type="submit" value="提交"></p>
    </form>
</body>
</html>
