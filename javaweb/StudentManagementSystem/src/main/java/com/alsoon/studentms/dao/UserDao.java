package com.alsoon.studentms.dao;

import java.util.ArrayList;

import com.alsoon.studentms.bean.User;

public interface UserDao {

  User login(User user);

  ArrayList<User> getAllStudents();

ArrayList<User> getStudentsBySearchQuery(String searchValue);

void deleteUserById(Integer id);

User getUserById(Integer id);

}
