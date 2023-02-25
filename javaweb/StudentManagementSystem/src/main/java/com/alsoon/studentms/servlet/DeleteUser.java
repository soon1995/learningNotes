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

@WebServlet("/delete")
public class DeleteUser extends HttpServlet {
  UserService userService = new UserServiceImpl();

  @Override
  protected void doGet(HttpServletRequest req, HttpServletResponse resp) throws ServletException, IOException {
    Integer id = Integer.parseInt(req.getParameter("id"));
    if (id != null) {
      userService.deleteUserById(id);
      ArrayList<User> students = userService.getAllStudents();
      req.setAttribute("students", students);
      req.getRequestDispatcher("/dashboard.jsp").forward(req, resp);
    } else {
      super.doGet(req, resp);
    }
  }

  
}
