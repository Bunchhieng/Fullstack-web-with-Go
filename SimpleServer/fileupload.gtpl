<html>
<head>
    <title></title>
</head>
<body>
<form enctype="multipart/form-data" action="/upload" method="POST">
    <input type="file" name="uploadfile">
    <input type="hidden" name="token" value="{{.}}">
    <input type="submit" value="Upload">
</form>
</body>
</html>