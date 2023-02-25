package com.alsoon.studentms.service;

import java.util.ArrayList;

import com.alsoon.studentms.bean.User;

public interface UserService {

    User login(User user);

    ArrayList<User> getAllStudents();

    ArrayList<User> getStudentsBySearchQuery(String searchValue);

    void deleteUserById(Integer id);

    User getUserById(Integer id);
  
}
