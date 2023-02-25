 <%@ page contentType="text/html;charset=UTF-8" language="java" %>
 <%@taglib prefix="c" uri="http://java.sun.com/jsp/jstl/core" %>
<html>
  <head>
    <title>Score</title>
  </head>
  <body>
    <div>
      <table id="table" style="height: 30px; width: 700px; border: black 1px solid; border-collapse: collapse">
        <tr>
          <td>Course Name</td>
          <td>Score</td>
        </tr>
        <c:forEach items="${courses}" var="course">
          <tr>
            <td>${course.getC_name()}</td>
            <td>${course.getScore()}</td>
          </tr>
        </c:forEach>
      </table>
    </div>
  </body>
</html>
