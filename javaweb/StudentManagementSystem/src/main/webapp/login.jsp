<%@ page contentType="text/html;charset=UTF-8" language="java" %>
<html>
  <head>
    <title>Login</title>
  </head>
  <body>
    <div>
      Username:&nbsp;<input type="text" id="u_id"><br>
      Password:&nbsp;<input type="password" id="u_pwd"><br>
      <button type="button" id="btn_01">Reset</button>
      <button type="submit" id="btn_02">Login</button><br>
      <span id="errormsg" style="color: red"></span>
    </div>
  </body>
  <script>
    // let btn01 = document.getElementById("btn_01");
    // btn01.onclick = function() {
      // let id_str = document.getElementById("u_id");
      //                       id_str.setValue("bc");
      // if (id_str == null) {

      //                       return;
      //                       }
      //                       let pwd_str = document.getElementById("u_pwd").value;
      //                       if (id_str == null) {
      //                       return;
      //                       }

      //                       }
    let btn02 = document.getElementById("btn_02");
    btn02.onclick = function() {
      let id_str = document.getElementById("u_id").value;
           if (id_str == null || "" == id_str) {
                            document.getElementById("errormsg").innerHTML = "username cannot be empty";
                            return;
                            }
                            let pwd_str = document.getElementById("u_pwd").value;
                            if (pwd_str == null || "" == pwd_str) {
                            document.getElementById("errormsg").innerHTML = "password cannot be empty";
                            return;
                            }
                            document.getElementById("errormsg").innerHTML = "";
                            window.location.href = "login?u_id=" + id_str + "&u_pwd=" + pwd_str;
                            
                            }
  </script>
</html>
<%-- <%@ page contentType="text/html; charset=UTF-8" language="java" %>
<html>
  <head>
    <title>Login</title>
  </head>
  <body>
    <div>

    </div>
  </body>
</html>
--%>
