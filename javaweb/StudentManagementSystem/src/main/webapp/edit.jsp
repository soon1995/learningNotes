 <%@ page contentType="text/html;charset=UTF-8" language="java" %>
 <%@taglib prefix="c" uri="http://java.sun.com/jsp/jstl/core" %>
<html>
  <head>
    <title>Score</title>
<style>
tr {
  border: 1px solid black;
}
</style>
  </head>
  <body>
    <div>
      <form name="editForm" action="edit?id=${student.getU_id()}" method="POST">
      <table id="table" style="height: 30px; width: 700px; border: black 1px solid; border-collapse: collapse">
        <input type="hidden" value="${student.getU_id()}" name="u_id"/>
        <tr>
          <td></td>
          <td>Detail</td>
        </tr>
          <tr>
            <td>Name</td>
            <td><input type="text" value="${student.getU_name()}" name="u_name"/></td>
          </tr>
          <tr>
            <td>Phone</td>
            <td><input type="text" value="${student.getU_phone()}" name="u_phone"/></td>
          </trconsole>
          <tr>
            <td colspan="2">Courses</td>
          </tr>
          <c:forEach items="${courses}" var="course">
            <tr>
              <input type="hidden" name="c_ids" value="${course.getC_id()}"/>
              <td>${course.getC_name()}</td>
              <td><input type="number" value="${course.getScore()}" name="c_scores"/></td>
            </tr>
          </c:forEach>
      </table>
      <button type="button" onclick="goDashboard()">Back</button>
      <button type="submit">Edit</button>
      </form>
    </div>
    <script>
              function goDashboard() {
                window.location.href = "dashboard";
              }
    </script>
  </body>
</html>
