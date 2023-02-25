package com.soon.test;

import java.util.ArrayList;

import org.junit.Assert;

import com.alsoon.studentms.bean.User;
import com.alsoon.studentms.bean.UserCourse;
import com.alsoon.studentms.bean.UserEditVo;
import com.alsoon.studentms.dao.UserCourseDao;
import com.alsoon.studentms.dao.UserDao;
import com.alsoon.studentms.dao.impl.UserCourseDaoImpl;
import com.alsoon.studentms.dao.impl.UserDaoImpl;

public class Test {
  @org.junit.Test
  public void getUser() {
    UserDao dao = new UserDaoImpl();
    User user = new User();
    user.setU_id(1001);
    user.setU_pwd("111111");
    User dbuser = dao.login(user);

    Assert.assertTrue(dbuser != null);
  }

  @org.junit.Test
  public void getUserFail() {
    UserDao dao = new UserDaoImpl();
    User user = new User();
    user.setU_id(0001);
    user.setU_pwd("111111");
    User dbuser = dao.login(user);
    Assert.assertTrue(dbuser == null);
  }

  @org.junit.Test
  public void getAllStudents() {
    UserDao dao = new UserDaoImpl();
    ArrayList<User> arrayList = dao.getAllStudents();
    Assert.assertTrue(arrayList.size() != 0);
  }

  @org.junit.Test
  public void getStudentsBySearchQuery() {
    UserDao dao = new UserDaoImpl();
    String searchQuery = "on";
    ArrayList<User> arrayList = dao.getStudentsBySearchQuery(searchQuery);
    Assert.assertEquals(2, arrayList.size());
  }

  @org.junit.Test
  public void getCourseDetailByUserId() {
    UserCourseDao dao = new UserCourseDaoImpl();
    ArrayList<UserCourse> arrayList = dao.getCourseDetailsByUserId(1001);
    System.out.println(arrayList);
    Assert.assertEquals(2, arrayList.size());
  }

  @org.junit.Test
  public void deleteUserByUserId() {
    UserDao dao = new UserDaoImpl();
    dao.deleteUserById(1006);
    Assert.assertTrue(true);
  }

  @org.junit.Test
  public void updateUserCourseByUserId() {
    UserCourseDao dao = new UserCourseDaoImpl();
    UserEditVo userEditVo = new UserEditVo();
    userEditVo.setU_id(1004);
    userEditVo.setU_name("Hello1");
    userEditVo.setU_phone("98777765411");
    userEditVo.setC_ids(new String[]{"1002", "1001"});
    userEditVo.setC_scores(new String[]{"33", "44"});
    dao.updateUserCourseById(userEditVo);
  }
}
