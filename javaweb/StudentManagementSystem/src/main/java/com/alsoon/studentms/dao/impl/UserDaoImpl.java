package com.alsoon.studentms.dao.impl;

import java.sql.Connection;
import java.sql.PreparedStatement;
import java.sql.ResultSet;
import java.util.ArrayList;

import com.alsoon.studentms.bean.User;
import com.alsoon.studentms.dao.UserDao;
import com.alsoon.studentms.util.ConnectionFactory;
import com.alsoon.studentms.util.UserIsDeleted;
import com.alsoon.studentms.util.UserRole;

public class UserDaoImpl implements UserDao {

  @Override
  public User login(User user) {
    Integer id = user.getU_id();
    String pwd = user.getU_pwd();

    Connection connection = ConnectionFactory.getConnection();
    PreparedStatement preparedStatement = null;
    ResultSet res = null;
    String sql = "SELECT * FROM user WHERE u_id=? AND u_pwd=? AND u_isdeleted=?";
    try {
      preparedStatement = connection.prepareStatement(sql);
      preparedStatement.setInt(1, id);
      preparedStatement.setString(2, pwd);
      preparedStatement.setInt(3, UserIsDeleted.NOT_DELETED);
      res = preparedStatement.executeQuery();

      while (res.next()) {
        User u = new User();
        u.setU_id(res.getInt("u_id"));
        u.setU_name(res.getString("u_name"));
        u.setU_pwd(res.getString("u_pwd"));
        u.setU_phone(res.getString("u_phone"));
        u.setU_role(res.getInt("u_role"));
        u.setU_isdeleted(res.getInt("u_isdeleted"));
        return u;
      }

    } catch (Exception e) {
      e.printStackTrace();
    } finally {
      ConnectionFactory.close(connection, preparedStatement, res);
    }
    return null;
  }

  @Override
  public ArrayList<User> getAllStudents() {
    Connection connection = ConnectionFactory.getConnection();
    PreparedStatement preparedStatement = null;
    ResultSet res = null;
    String sql = "SELECT * FROM user WHERE u_role=? AND u_isdeleted=?";
    int role = UserRole.STUDENT; 
    ArrayList<User> students = new ArrayList<>();
    try {
      preparedStatement = connection.prepareStatement(sql);
      preparedStatement.setInt(1, role);
      preparedStatement.setInt(2, UserIsDeleted.NOT_DELETED);
      res = preparedStatement.executeQuery();

      while(res.next()) {
        User u = new User();
        u.setU_id(res.getInt("u_id"));
        u.setU_name(res.getString("u_name"));
        u.setU_pwd(res.getString("u_pwd"));
        u.setU_phone(res.getString("u_phone"));
        u.setU_role(res.getInt("u_role"));
        u.setU_isdeleted(res.getInt("u_isdeleted"));
        students.add(u);
      }
    } catch (Exception e) {
      e.printStackTrace();
    } finally {
      ConnectionFactory.close(connection, preparedStatement, res);
    }
    return students;
  }

  @Override
  public ArrayList<User> getStudentsBySearchQuery(String searchValue) {
    searchValue = "%" + searchValue + "%";

    Connection connection = ConnectionFactory.getConnection();
    PreparedStatement preparedStatement = null;
    ResultSet res = null;
    String sql = "SELECT * FROM user WHERE u_role=? AND (u_name LIKE ? OR u_id LIKE ?) AND u_isdeleted=?";
    int role = UserRole.STUDENT; 
    ArrayList<User> students = new ArrayList<>();
    try {
      preparedStatement = connection.prepareStatement(sql);
      preparedStatement.setInt(1, role);
      preparedStatement.setString(2, searchValue);
      preparedStatement.setString(3, searchValue);
      preparedStatement.setInt(4, UserIsDeleted.NOT_DELETED);
      res = preparedStatement.executeQuery();

      while(res.next()) {
        User u = new User();
        u.setU_id(res.getInt("u_id"));
        u.setU_name(res.getString("u_name"));
        u.setU_pwd(res.getString("u_pwd"));
        u.setU_phone(res.getString("u_phone"));
        u.setU_role(res.getInt("u_role"));
        u.setU_isdeleted(res.getInt("u_isdeleted"));
        students.add(u);
      }
    } catch (Exception e) {
      e.printStackTrace();
    } finally {
      ConnectionFactory.close(connection, preparedStatement, res);
    }
    return students;
  }

  @Override
  public void deleteUserById(Integer id) {
    Connection connection = ConnectionFactory.getConnection();
    PreparedStatement currentStatement = null;
    ResultSet res = null;
    String sql1 = "DELETE FROM user_course WHERE u_id=?";
    String sql2 = "DELETE FROM user WHERE u_id=?";
    try {
      currentStatement = connection.prepareStatement(sql1);
      currentStatement.setInt(1, id);
      currentStatement.executeUpdate();
    } catch (Exception e) {
      e.printStackTrace();
    } finally {
      ConnectionFactory.close(null, currentStatement, res);
    }

    try {
      currentStatement = connection.prepareStatement(sql2);
      currentStatement.setInt(1, id);
      currentStatement.executeUpdate();
    } catch (Exception e) {
      e.printStackTrace();
    } finally {
      ConnectionFactory.close(connection, currentStatement, res);
    }
  }

  @Override
  public User getUserById(Integer id) {
    Connection connection = ConnectionFactory.getConnection();
    PreparedStatement currentStatement = null;
    ResultSet res = null;
    String sql = "SELECT * FROM user WHERE u_id=?";
    try {
      currentStatement = connection.prepareStatement(sql);
      currentStatement.setInt(1, id);
      res = currentStatement.executeQuery();
      while (res.next()) {
        User u = new User();
        u.setU_id(res.getInt("u_id"));
        u.setU_name(res.getString("u_name"));
        u.setU_pwd(res.getString("u_pwd"));
        u.setU_phone(res.getString("u_phone"));
        u.setU_role(res.getInt("u_role"));
        u.setU_isdeleted(res.getInt("u_isdeleted"));
        return u;
      }
    } catch (Exception e) {
      e.printStackTrace();
    } finally {
      ConnectionFactory.close(null, currentStatement, res);
    }

    return null;
  }
  
}
