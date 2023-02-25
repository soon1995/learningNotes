package com.alsoon.studentms.servlet;

import java.io.IOException;
import java.util.ArrayList;

import javax.servlet.ServletException;
import javax.servlet.ServletRequest;
import javax.servlet.ServletResponse;
import javax.servlet.annotation.WebServlet;
import javax.servlet.http.HttpServlet;

import com.alsoon.studentms.bean.User;
import com.alsoon.studentms.service.UserService;
import com.alsoon.studentms.service.impl.UserServiceImpl;

@WebServlet("/searchUser")
public class SearchUserServlet extends HttpServlet {
  UserService userService = new UserServiceImpl();

  @Override
  public void service(ServletRequest req, ServletResponse res) throws ServletException, IOException {
    req.setCharacterEncoding("UTF-8");
    String searchValue = req.getParameter("searchValue");
    if (searchValue == null) {
      searchValue = "";
    }
    ArrayList<User> students = userService.getStudentsBySearchQuery(searchValue);
    req.setAttribute("students", students);
    req.getRequestDispatcher("dashboard.jsp").forward(req, res);
  }

}
