package com.alsoon.studentms.servlet;

import java.io.IOException;
import java.lang.reflect.InvocationTargetException;
import javax.servlet.ServletException;
import javax.servlet.annotation.WebServlet;
import javax.servlet.http.HttpServlet;
import javax.servlet.http.HttpServletRequest;
import javax.servlet.http.HttpServletResponse;

import org.apache.commons.beanutils.BeanUtils;

import com.alsoon.studentms.bean.UserEditVo;
import com.alsoon.studentms.service.UserCourseService;
import com.alsoon.studentms.service.impl.UserCourseServiceImpl;

@WebServlet("/edit")
public class EditServlet extends HttpServlet{
  UserCourseService userCourseService = new UserCourseServiceImpl();
  @Override
  protected void doPost(HttpServletRequest req, HttpServletResponse resp) throws ServletException, IOException {
    UserEditVo userEditVo = new UserEditVo();
    try {
      BeanUtils.populate(userEditVo, req.getParameterMap());
      userCourseService.updateUserCourseById(userEditVo);
    } catch (IllegalAccessException | InvocationTargetException e) {
      e.printStackTrace();
    }
    req.setAttribute("message", "Successfully updated");
    resp.setStatus(400);
    req.getRequestDispatcher("/dashboard").forward(req, resp);
  }
  
}
