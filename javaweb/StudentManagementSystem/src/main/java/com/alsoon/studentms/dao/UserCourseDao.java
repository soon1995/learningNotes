package com.alsoon.studentms.dao;

import java.util.ArrayList;

import com.alsoon.studentms.bean.UserCourse;
import com.alsoon.studentms.bean.UserEditVo;

public interface UserCourseDao {

    ArrayList<UserCourse> getCourseDetailsByUserId(Integer id);

    void updateUserCourseById(UserEditVo userEditVo);
  
}
