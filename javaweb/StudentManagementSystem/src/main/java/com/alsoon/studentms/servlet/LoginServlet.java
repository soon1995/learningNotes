package com.alsoon.studentms.servlet;

import java.io.IOException;
import java.util.ArrayList;

import javax.servlet.ServletException;
import javax.servlet.annotation.WebServlet;
import javax.servlet.http.HttpServlet;
import javax.servlet.http.HttpServletRequest;
import javax.servlet.http.HttpServletResponse;
import javax.servlet.http.HttpSession;

import com.alsoon.studentms.bean.User;
import com.alsoon.studentms.service.UserService;
import com.alsoon.studentms.service.impl.UserServiceImpl;

@WebServlet("/login")
public class LoginServlet extends HttpServlet {
  UserService userService = new UserServiceImpl();

  @Override
  protected void doGet(HttpServletRequest req, HttpServletResponse resp) throws ServletException, IOException {
    // not doing error handling
    Integer uid = Integer.parseInt(req.getParameter("u_id"));
    String upwd = req.getParameter("u_pwd");
    User user = new User();
    user.setU_id(uid);
    user.setU_pwd(upwd);
    User dbUser = userService.login(user);
    if (dbUser != null && dbUser.getU_id().equals(uid) && dbUser.getU_pwd().equals(upwd)) {
      // enter dashboard
      HttpSession session = req.getSession();
      session.setAttribute("user", dbUser);

      ArrayList<User> students = userService.getAllStudents();
      dbUser.setU_pwd(null);
      req.setAttribute("students", students);

      req.getRequestDispatcher("dashboard.jsp").forward(req, resp);
    } else {
      // relogin
    }
  }
}
