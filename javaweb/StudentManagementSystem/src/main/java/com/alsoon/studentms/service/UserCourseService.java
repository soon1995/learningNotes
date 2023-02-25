package com.alsoon.studentms.service;

import java.util.ArrayList;

import com.alsoon.studentms.bean.UserCourse;
import com.alsoon.studentms.bean.UserEditVo;


public interface UserCourseService {

    ArrayList<UserCourse> getCourseDetailsByUserId(Integer id);

    void updateUserCourseById(UserEditVo userEditVo);
  
}
