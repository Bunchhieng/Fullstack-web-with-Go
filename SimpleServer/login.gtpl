<html>
<head>
    <title></title>
</head>
<body>
<form action="/login" method="POST">
    Username:<input type="text" name="username">
    Password:<input type="text" name="password">
    <input type="hidden" name="token" value="{{.}}">
    <input type="submit" value="Login">
</form>
</body>
</html>