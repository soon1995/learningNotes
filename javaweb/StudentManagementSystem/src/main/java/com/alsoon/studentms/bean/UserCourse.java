package com.alsoon.studentms.bean;

import java.io.Serializable;

import lombok.AllArgsConstructor;
import lombok.Data;
import lombok.NoArgsConstructor;

@Data
@NoArgsConstructor
@AllArgsConstructor
public class UserCourse implements Serializable{
  private Integer u_id;
  private Integer c_id;
  private String c_name;
  private String c_teacher;
  private Integer score;

}
