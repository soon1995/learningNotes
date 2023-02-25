package  com.alsoon.studentms.servlet;

import java.io.IOException;
import java.util.ArrayList;

import javax.servlet.ServletException;
import javax.servlet.annotation.WebServlet;
import javax.servlet.http.HttpServlet;
import javax.servlet.http.HttpServletRequest;
import javax.servlet.http.HttpServletResponse;

import com.alsoon.studentms.bean.User;
import com.alsoon.studentms.bean.UserCourse;
import com.alsoon.studentms.service.UserCourseService;
import com.alsoon.studentms.service.UserService;
import com.alsoon.studentms.service.impl.UserCourseServiceImpl;
import com.alsoon.studentms.service.impl.UserServiceImpl;

@WebServlet("/editinfo")
public class ShowEditServlet extends HttpServlet {

  UserService userService = new UserServiceImpl();
  UserCourseService userCourseService = new UserCourseServiceImpl();

  @Override
  protected void doGet(HttpServletRequest req, HttpServletResponse resp) throws ServletException, IOException {
    Integer id = Integer.parseInt(req.getParameter("id"));
    if (id != null) {
      User student = userService.getUserById(id);
      if (student != null) {
        req.setAttribute("student", student);
        ArrayList<UserCourse> courses = userCourseService.getCourseDetailsByUserId(id); 
        req.setAttribute("courses", courses);
      }
      req.getRequestDispatcher("/edit.jsp").forward(req, resp);
    } else {
      super.doGet(req, resp);
    } 
  }
  
}
