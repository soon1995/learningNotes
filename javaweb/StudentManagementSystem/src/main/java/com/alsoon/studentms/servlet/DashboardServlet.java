package com.alsoon.studentms.servlet;

import java.io.IOException;
import java.util.ArrayList;

import javax.servlet.ServletException;
import javax.servlet.annotation.WebServlet;
import javax.servlet.http.HttpServlet;
import javax.servlet.http.HttpServletRequest;
import javax.servlet.http.HttpServletResponse;

import com.alsoon.studentms.bean.User;
import com.alsoon.studentms.service.UserService;
import com.alsoon.studentms.service.impl.UserServiceImpl;

@WebServlet("/dashboard")
public class DashboardServlet extends HttpServlet{
  UserService userService = new UserServiceImpl();
  @Override
  protected void doGet(HttpServletRequest req, HttpServletResponse resp) throws ServletException, IOException {
    doPost(req, resp);
  }

  @Override
  protected void doPost(HttpServletRequest req, HttpServletResponse resp) throws ServletException, IOException {
    ArrayList<User> students = userService.getAllStudents();
    req.setAttribute("students", students);
    req.getRequestDispatcher("dashboard.jsp").forward(req, resp);
      
  }
  
}
