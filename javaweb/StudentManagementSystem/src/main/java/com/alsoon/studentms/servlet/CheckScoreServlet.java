package com.alsoon.studentms.servlet;

import java.io.IOException;
import java.util.ArrayList;

import javax.servlet.ServletException;
import javax.servlet.annotation.WebServlet;
import javax.servlet.http.HttpServlet;
import javax.servlet.http.HttpServletRequest;
import javax.servlet.http.HttpServletResponse;

import com.alsoon.studentms.bean.UserCourse;
import com.alsoon.studentms.service.UserCourseService;
import com.alsoon.studentms.service.impl.UserCourseServiceImpl;

@WebServlet("/checkScore")
public class CheckScoreServlet extends HttpServlet {
  UserCourseService userCourseService = new UserCourseServiceImpl();

  @Override
  protected void doGet(HttpServletRequest req, HttpServletResponse resp) throws ServletException, IOException {
    Integer id = Integer.parseInt(req.getParameter("id"));
    if (id != null) {
      ArrayList<UserCourse> userCourses = userCourseService.getCourseDetailsByUserId(id);
      req.setAttribute("courses", userCourses);
      req.getRequestDispatcher("/score.jsp").forward(req, resp);
    } else {

    }
  }

}
