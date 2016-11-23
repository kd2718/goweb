<!DOCTYPE html>

        <html>
        <head>

        </head>
        <body>
        <h1>Enter your email and password</h1>
        <form method="post">
                <input type="text" name="email" placeholder="name@email.com">
                <input name="password" type="password" placeholder="password">
                <input type="submit" name="submit" value="Login">

        </form>
        <hr>
        <h1>Or register!</h1>
        <form method="post" action="/register/">
                <input type="text" name="email" placeholder="name@email.com">
                <input name="password" type="password" placeholder="password">
                <input name="password2" type="password" placeholder="confirm password">
                <input type="submit" name="submit" value="Register">
        </form>
        </body>
</html>