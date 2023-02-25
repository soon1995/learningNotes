package com.alsoon.studentms.service.impl;

import java.util.ArrayList;

import com.alsoon.studentms.bean.User;
import com.alsoon.studentms.dao.UserDao;
import com.alsoon.studentms.dao.impl.UserDaoImpl;
import com.alsoon.studentms.service.UserService;

public class UserServiceImpl implements UserService {

  UserDao userDao = new UserDaoImpl();

  @Override
  public User login(User user) {
    return userDao.login(user);
  }

  @Override
  public ArrayList<User> getAllStudents() {
    
    return userDao.getAllStudents();
  }

  @Override
  public ArrayList<User> getStudentsBySearchQuery(String searchValue) {
    return userDao.getStudentsBySearchQuery(searchValue);
  }

  @Override
  public void deleteUserById(Integer id) {
    userDao.deleteUserById(id); 
  }

  @Override
  public User getUserById(Integer id) {
    return userDao.getUserById(id);
  }

}
