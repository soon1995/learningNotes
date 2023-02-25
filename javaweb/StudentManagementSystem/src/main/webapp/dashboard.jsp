 <%@ page contentType="text/html;charset=UTF-8" language="java" %>
 <%@taglib prefix="c" uri="http://java.sun.com/jsp/jstl/core" %>
<html>
  <head>
    <title>DashBoard</title>
<style>
#head_{
  height:100px;
  width: 100%;
  background-color:slategray;
}
#admin {
  width: 120px;
  height: 50px;
  padding-top: 25px;
  margin-left: 90%;
}
#search{
  height: 70px;
  width:100%;
  background-color: cornflowerblue;
}
#inner_s {
  width: 30%;
  padding-top: 25px;
  padding-left: 45%;
}
#table {
  margin-left: 30%;
  margin-top: 30px;
}
td {
  text-align: center;
  height: 20px;
  width: 150px;
  border: black 1px solid;
  padding: 1px
}
</style>
  </head>
  <body>
    ${user.getU_name()}, welcome to Student Management System.<br>
    
    <div id="head_">
      <div id="admin">
        Name: <span style="color: red">${user.getU_name()}</span><br>
        Id: <span style="color: red">${user.getU_id()}</span>
      </div>
    </div>
    <div id="search">
      <div id="inner_s">
        <input type="text" style="font-size: 20px; height: 26px; width: 190px" id="searchId">&nbsp;&nbsp;
        <button style="font-size: 18px; height: 28px" onclick="search()">Search</button>
      </div>
    </div>

    <div>
      <table id="table" style="height: 30px; width: 700px; border: black 1px solid; border-collapse: collapse">
        <tr>
          <td>Id</td>
          <td>Name</td>
          <td>Phone</td>
          <td>Check Result</td>
          <td>Operation</td>
        </tr>
        <c:forEach items="${students}" var="student">
          <tr>
            <td>${student.getU_id()}</td>
            <td>${student.getU_name()}</td>
            <td>${student.getU_phone()}</td>
            <td><button style="color: chocolate" onclick="checkScore(${student.getU_id()})">Check</button></td>
            <td>
              <button style="color: chocolate" onclick="editUser(${student.getU_id()})">Edit</button>
              <button style="color: chocolate" onclick="deleteUser(${student.getU_id()}, '${student.getU_name()}')">Delete</button>
            </td>
          </tr>
        </c:forEach>
      </table>
      <span style="color: blue">${message}</span>
    </div>

    <script>
                      function search() {
                        let searchValue = document.getElementById('searchId').value;
                        window.location.href = "searchUser?searchValue=" + searchValue;
                      }

                      function checkScore(id) {
                        window.location.href = "checkScore?id=" + id;
                      }

                                               function editUser(id) {
                                                window.location.href = "editinfo?id=" + id;
                                               }

                                               function deleteUser(id, name) {
                                               let question = "Delete student '" + name + "'?";
                                                  var confirmDelete = confirm(question);
                                                  if (confirmDelete) {

                                               window.location.href = "delete?id=" + id;
                                               }
                                               }
    </script>
  </body>
</html>
