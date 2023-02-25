package com.alsoon.studentms.dao.impl;

import java.sql.Connection;
import java.sql.PreparedStatement;
import java.sql.ResultSet;
import java.util.ArrayList;

import com.alsoon.studentms.bean.UserCourse;
import com.alsoon.studentms.bean.UserEditVo;
import com.alsoon.studentms.dao.UserCourseDao;
import com.alsoon.studentms.util.ConnectionFactory;
import com.alsoon.studentms.util.UserIsDeleted;

public class UserCourseDaoImpl implements UserCourseDao{

  @Override
  public ArrayList<UserCourse> getCourseDetailsByUserId(Integer id) {
    ArrayList<UserCourse> courses = new ArrayList<>();

    Connection conn = ConnectionFactory.getConnection();
    PreparedStatement preparedStatement = null;
    ResultSet res = null;
    String sql = "SELECT * FROM course c, user_course uc WHERE c.c_id = uc.c_id AND uc.u_id = ?";
    try {
      preparedStatement = conn.prepareStatement(sql);
      preparedStatement.setInt(1, id);
      res = preparedStatement.executeQuery();
      while (res.next()) {
        UserCourse uc = new UserCourse();
        uc.setU_id(id);
        uc.setC_id(res.getInt("c_id"));
        uc.setC_name(res.getString("c_name"));
        uc.setScore(res.getInt("score"));
        uc.setC_teacher(res.getString("c_teacher"));
        courses.add(uc);
      }
    } catch (Exception e) {
      e.printStackTrace();
    } finally {
      ConnectionFactory.close(conn, preparedStatement, res);
    }
    return courses;
  }

  @Override
  public void updateUserCourseById(UserEditVo userEditVo) {
    Connection conn = ConnectionFactory.getConnection();
    PreparedStatement preparedStatement = null;
    ResultSet res = null;
    String sql = "UPDATE user SET u_name=?, u_phone=? WHERE u_id=? AND u_isdeleted=?";
    try {
      preparedStatement = conn.prepareStatement(sql);
      preparedStatement.setString(1, userEditVo.getU_name());
      preparedStatement.setString(2, userEditVo.getU_phone());
      preparedStatement.setInt(3, userEditVo.getU_id());
      preparedStatement.setInt(4, UserIsDeleted.NOT_DELETED);
      preparedStatement.executeUpdate();
    } catch (Exception e) {
      e.printStackTrace();
    } finally {
      ConnectionFactory.close(null, preparedStatement, res);
    }
    sql = "UPDATE user_course SET score=? WHERE c_id=? AND u_id=?";
    try {
      preparedStatement = conn.prepareStatement(sql);
      for (int i = 0; i < userEditVo.getC_ids().length; i++) {
        preparedStatement.setInt(1, Integer.parseInt(userEditVo.getC_scores()[i]));
        preparedStatement.setInt(2, Integer.parseInt(userEditVo.getC_ids()[i]));
        preparedStatement.setInt(3, userEditVo.getU_id());
        preparedStatement.addBatch();
      }
      preparedStatement.executeBatch();
    } catch (Exception e) {
      e.printStackTrace();
    } finally {
      ConnectionFactory.close(conn, preparedStatement, res);
    }
  }
  
}
