package com.alsoon.studentms.service.impl;

import java.util.ArrayList;

import com.alsoon.studentms.bean.UserCourse;
import com.alsoon.studentms.bean.UserEditVo;
import com.alsoon.studentms.dao.UserCourseDao;
import com.alsoon.studentms.dao.impl.UserCourseDaoImpl;
import com.alsoon.studentms.service.UserCourseService;

public class UserCourseServiceImpl implements UserCourseService{
  UserCourseDao userCourseDao = new UserCourseDaoImpl();

  @Override
  public ArrayList<UserCourse> getCourseDetailsByUserId(Integer id) {
    return userCourseDao.getCourseDetailsByUserId(id);
  }

  @Override
  public void updateUserCourseById(UserEditVo userEditVo) {
    userCourseDao.updateUserCourseById(userEditVo);
  }
  

}
